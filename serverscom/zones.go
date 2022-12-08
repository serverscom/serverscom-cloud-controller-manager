package serverscom

import (
	"context"
	"fmt"

	serverscom "github.com/serverscom/serverscom-go-client/pkg"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
)

type zones struct {
	client      *serverscom.Client
	defaultZone string
}

func newZones(client *serverscom.Client, defaultZone string) *zones {
	return &zones{client, defaultZone}
}

func (z zones) GetZone(_ context.Context) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{Region: z.defaultZone}, nil
}

func (z zones) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
	instanceType, instanceID, err := parseProviderID(providerID)
	if err != nil {
		return cloudprovider.Zone{}, err
	}

	switch instanceType {
	case cloudInstanceType:
		cloudInstance, err := z.client.CloudComputingInstances.Get(ctx, instanceID)
		if err != nil {
			return cloudprovider.Zone{}, fmt.Errorf("can't get cloud instance: %s", err.Error())
		}

		return cloudprovider.Zone{Region: cloudInstance.RegionCode}, nil
	case dedicatedServerType:
		host, err := z.client.Hosts.GetDedicatedServer(ctx, instanceID)
		if err != nil {
			return cloudprovider.Zone{}, fmt.Errorf("can't get dedicated server: %s", err.Error())
		}

		return cloudprovider.Zone{Region: host.LocationCode}, nil
	case kubernetesBaremetalNodeType:
		host, err := z.client.Hosts.GetKubernetesBaremetalNode(ctx, instanceID)
		if err != nil {
			return cloudprovider.Zone{}, fmt.Errorf("can't get kubernetes baremetal node %s", err.Error())
		}

		return cloudprovider.Zone{Region: host.LocationCode}, nil
	default:
		return cloudprovider.Zone{}, fmt.Errorf("invalid instance type: %s", instanceType)
	}
}

func (z zones) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	cloudInstances, err := z.client.CloudComputingInstances.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return cloudprovider.Zone{}, fmt.Errorf("can't get list of cloud instances: %s", err.Error())
	}

	for _, instance := range cloudInstances {
		if instance.Name == string(nodeName) {
			return cloudprovider.Zone{Region: instance.RegionCode}, nil
		}
	}

	hosts, err := z.client.Hosts.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return cloudprovider.Zone{}, fmt.Errorf("can't get list of hosts: %s", err.Error())
	}

	for _, host := range hosts {
		if host.Title == string(nodeName) {
			return cloudprovider.Zone{Region: host.LocationCode}, nil
		}
	}

	return cloudprovider.Zone{}, fmt.Errorf("can't find instance by name: %s", string(nodeName))
}
