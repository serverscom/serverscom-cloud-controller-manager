package serverscom

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	cli "github.com/serverscom/serverscom-go-client/pkg"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
)

const (
	loadBalancerActiveStatus = "active"

	loadBalancerNameAnnotation          = "servers.com/load-balancer-name"
	loadBalancerHostnameAnnotation      = "servers.com/load-balancer-hostname"
	loadBalancerLocationIdAnnotation    = "servers.com/load-balancer-location-id"
	loadBalancerProxyProtocolAnnotation = "servers.com/proxy-protocol"
)

type loadBalancers struct {
	client            *cli.Client
	defaultLocationID *int64
}

func newLoadBalancers(client *cli.Client, defaultLocationID *int64) cloudprovider.LoadBalancer {
	return &loadBalancers{client: client, defaultLocationID: defaultLocationID}
}

func (l *loadBalancers) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (*v1.LoadBalancerStatus, bool, error) {
	loadBalancer, err := l.findLoadBalancerByName(ctx, clusterName, service)

	if err != nil && isNotFoundError(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	if loadBalancer.Status != loadBalancerActiveStatus {
		return nil, true, fmt.Errorf("load balancer is not active, current status: %s", loadBalancer.Status)
	}

	return l.buildResult(service, loadBalancer), true, nil
}

func (l *loadBalancers) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	name, ok := service.Annotations[loadBalancerNameAnnotation]
	if !ok {
		return getLoadBalancerName(service, clusterName)
	}

	return name
}

func (l *loadBalancers) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	loadBalancer, err := l.findLoadBalancerByName(ctx, clusterName, service)
	if err != nil {
		if !isNotFoundError(err) {
			return nil, err
		}
	}

	vhostZones, upstreamZones, err := l.buildZones(service, nodes)
	if err != nil {
		return nil, err
	}

	if loadBalancer == nil {
		locationID, err := l.extractLocationID(service)
		if err != nil {
			return nil, err
		}

		input := cli.L4LoadBalancerCreateInput{}
		input.VHostZones = vhostZones
		input.UpstreamZones = upstreamZones
		input.LocationID = locationID
		input.Name = l.GetLoadBalancerName(ctx, clusterName, service)

		loadBalancer, err = l.client.LoadBalancers.CreateL4LoadBalancer(ctx, input)
		if err != nil {
			return nil, err
		}

		if loadBalancer.Status != loadBalancerActiveStatus {
			return nil, fmt.Errorf("load balancer is not active, current status: %s", loadBalancer.Status)
		}
	} else {
		name := l.GetLoadBalancerName(ctx, clusterName, service)

		input := cli.L4LoadBalancerUpdateInput{}
		input.VHostZones = vhostZones
		input.UpstreamZones = upstreamZones
		input.Name = &name

		loadBalancer, err = l.client.LoadBalancers.UpdateL4LoadBalancer(ctx, loadBalancer.ID, input)
		if err != nil {
			return nil, err
		}

		if loadBalancer.Status != loadBalancerActiveStatus {
			return nil, fmt.Errorf("load balancer is not active, current status: %s", loadBalancer.Status)
		}
	}

	return l.buildResult(service, loadBalancer), nil
}

func (l *loadBalancers) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	loadBalancer, err := l.findLoadBalancerByName(ctx, clusterName, service)
	if err != nil {
		return err
	}

	vhostZones, upstreamZones, err := l.buildZones(service, nodes)
	if err != nil {
		return err
	}

	name := l.GetLoadBalancerName(ctx, clusterName, service)

	input := cli.L4LoadBalancerUpdateInput{}
	input.VHostZones = vhostZones
	input.UpstreamZones = upstreamZones
	input.Name = &name

	_, err = l.client.LoadBalancers.UpdateL4LoadBalancer(ctx, loadBalancer.ID, input)
	if err != nil {
		return err
	}

	if loadBalancer.Status != loadBalancerActiveStatus {
		return fmt.Errorf("load balancer is not active, current status: %s", loadBalancer.Status)
	}

	return nil
}

func (l *loadBalancers) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	loadBalancer, err := l.findLoadBalancerByName(ctx, clusterName, service)
	if err != nil {
		if isNotFoundError(err) {
			return nil
		}

		return err
	}

	return l.client.LoadBalancers.DeleteL4LoadBalancer(ctx, loadBalancer.ID)
}

func (l *loadBalancers) findLoadBalancerByName(ctx context.Context, clusterName string, service *v1.Service) (*cli.L4LoadBalancer, error) {
	name := l.GetLoadBalancerName(ctx, clusterName, service)

	loadBalancers, err := l.client.LoadBalancers.Collection().
		SetPerPage(100).
		SetParam(typeParamKey, "l4").
		SetParam(searchPatternParamKey, name).
		Collect(ctx)

	if err != nil {
		return nil, err
	}

	if len(loadBalancers) == 0 {
		return nil, &cli.NotFoundError{
			StatusCode: 404,
			ErrorCode:  "NOT_FOUND",
			Message:    "Empty load balancers list",
		}
	}

	var currentLoadBalancer *cli.LoadBalancer

	for _, loadBalancer := range loadBalancers {
		if loadBalancer.Name == name {
			if currentLoadBalancer != nil {
				return nil, fmt.Errorf("found more than one load balancer with the same name: %s", name)
			}

			currentLoadBalancer = &loadBalancer
		}
	}

	return l.client.LoadBalancers.GetL4LoadBalancer(ctx, currentLoadBalancer.ID)
}

func (l *loadBalancers) buildZones(service *v1.Service, nodes []*v1.Node) ([]cli.L4VHostZoneInput, []cli.L4UpstreamZoneInput, error) {
	var vhostZoneInputs []cli.L4VHostZoneInput
	var upstreamZoneInputs []cli.L4UpstreamZoneInput

	for _, port := range service.Spec.Ports {
		if port.Protocol != "TCP" && port.Protocol != "UDP" {
			return nil, nil, fmt.Errorf("only TCP and UDP protocols is supported, got: %q", port.Protocol)
		}

		id := fmt.Sprintf("k8s-nodes-%d-%s", port.Port, strings.ToLower(string(port.Protocol)))

		vhostZoneInput := cli.L4VHostZoneInput{}
		vhostZoneInput.ID = id
		vhostZoneInput.UpstreamID = id
		vhostZoneInput.Ports = append(vhostZoneInput.Ports, port.Port)

		if port.Protocol == "UDP" {
			vhostZoneInput.UDP = true
		}

		proxyProtocolEnabled, ok := service.Annotations[loadBalancerProxyProtocolAnnotation]
		if ok && (proxyProtocolEnabled == "true" || proxyProtocolEnabled == "True") {
			vhostZoneInput.ProxyProtocol = true
		} else {
			vhostZoneInput.ProxyProtocol = false
		}

		upstreamZoneInput := cli.L4UpstreamZoneInput{}
		upstreamZoneInput.ID = id

		for _, node := range nodes {
			for _, address := range node.Status.Addresses {
				if address.Type != v1.NodeInternalIP {
					continue
				}

				upstreamZoneInput.Upstreams = append(upstreamZoneInput.Upstreams, cli.L4UpstreamInput{
					IP:     address.Address,
					Weight: 1,
					Port:   port.NodePort,
				})
			}
		}

		vhostZoneInputs = append(vhostZoneInputs, vhostZoneInput)
		upstreamZoneInputs = append(upstreamZoneInputs, upstreamZoneInput)
	}

	return vhostZoneInputs, upstreamZoneInputs, nil
}

func (l *loadBalancers) extractLocationID(service *v1.Service) (int64, error) {
	var locationID int64
	var err error

	customLocationId, ok := service.Annotations[loadBalancerLocationIdAnnotation]
	if !ok {
		if l.defaultLocationID != nil {
			locationID = *l.defaultLocationID
		} else {
			return 0, fmt.Errorf("no default location id is set, no annotation %s is set, can't create the load balancer", loadBalancerLocationIdAnnotation)
		}
	} else {
		locationID, err = strconv.ParseInt(customLocationId, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid annotation %s: %s", loadBalancerLocationIdAnnotation, err.Error())
		}
	}

	return locationID, nil
}

func (l *loadBalancers) buildResult(service *v1.Service, loadBalancer *cli.L4LoadBalancer) *v1.LoadBalancerStatus {
	hostname, ok := service.Annotations[loadBalancerHostnameAnnotation]
	if ok && hostname != "" {
		return &v1.LoadBalancerStatus{
			Ingress: []v1.LoadBalancerIngress{{Hostname: hostname}},
		}
	}

	var ingresses []v1.LoadBalancerIngress

	for _, ip := range loadBalancer.ExternalAddresses {
		ingresses = append(ingresses, v1.LoadBalancerIngress{IP: ip})
	}

	return &v1.LoadBalancerStatus{Ingress: ingresses}
}
