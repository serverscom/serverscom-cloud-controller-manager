package serverscom

import (
	"context"
	"fmt"

	cli "github.com/serverscom/serverscom-go-client/pkg"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
)

type instances struct {
	client *cli.Client
}

func newInstances(client *cli.Client) cloudprovider.Instances {
	return &instances{client: client}
}

func (i *instances) NodeAddresses(ctx context.Context, nodeName types.NodeName) ([]v1.NodeAddress, error) {
	cloudInstances, err := i.client.CloudComputingInstances.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return nil, fmt.Errorf("can't get list of cloud instances: %s", err.Error())
	}

	for _, instance := range cloudInstances {
		if instance.Name == string(nodeName) {
			return collectCloudInstanceAddresses(&instance), nil
		}
	}

	hosts, err := i.client.Hosts.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return nil, fmt.Errorf("can't get list of hosts: %s", err.Error())
	}

	for _, host := range hosts {
		if host.Title == string(nodeName) {
			return collectHostAddresses(&host), nil
		}
	}

	return nil, fmt.Errorf("can't find instance by name: %s", string(nodeName))
}

func (i *instances) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	instanceType, instanceID, err := parseProviderID(providerID)
	if err != nil {
		return nil, err
	}

	switch instanceType {
	case cloudInstanceType:
		cloudInstance, err := i.client.CloudComputingInstances.Get(ctx, instanceID)
		if err != nil {
			return nil, fmt.Errorf("can't get cloud instance: %s", err.Error())
		}

		return collectCloudInstanceAddresses(cloudInstance), nil
	case dedicatedServerType:
		host, err := i.client.Hosts.GetDedicatedServer(ctx, instanceID)
		if err != nil {
			return nil, fmt.Errorf("can't get dedicated server: %s", err.Error())
		}

		return collectHostAddresses(&cli.Host{
			Title:              host.Title,
			PrivateIPv4Address: host.PrivateIPv4Address,
			PublicIPv4Address:  host.PublicIPv4Address,
		}), nil
	case kubernetesBaremetalNodeType:
		host, err := i.client.Hosts.GetKubernetesBaremetalNode(ctx, instanceID)
		if err != nil {
			return nil, fmt.Errorf("can't get kubernetes baremetal node %s", err.Error())
		}

		return collectHostAddresses(&cli.Host{
			Title:              host.Title,
			PrivateIPv4Address: host.PrivateIPv4Address,
			PublicIPv4Address:  host.PublicIPv4Address,
		}), nil
	default:
		return nil, fmt.Errorf("invalid instance type: %s", instanceType)
	}
}

func (i *instances) ExternalID(ctx context.Context, nodeName types.NodeName) (string, error) {
	return i.InstanceID(ctx, nodeName)
}

func (i *instances) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	cloudInstances, err := i.client.CloudComputingInstances.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return "", fmt.Errorf("can't get list of cloud instances: %s", err.Error())
	}

	for _, instance := range cloudInstances {
		if instance.Name == string(nodeName) || anyMatch(string(nodeName), instance.PrivateIPv4Address, instance.PublicIPv4Address) {
			return buildExternalID(cloudInstanceType, instance.ID), nil
		}
	}

	hosts, err := i.client.Hosts.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return "", fmt.Errorf("can't get list of hosts: %s", err.Error())
	}

	for _, host := range hosts {
		if host.Title == string(nodeName) || anyMatch(string(nodeName), host.PrivateIPv4Address, host.PublicIPv4Address) {
			switch host.Type {
			case "dedicated_server":
				return buildExternalID(dedicatedServerType, host.ID), nil
			case "kubernetes_baremetal_node":
				return buildExternalID(kubernetesBaremetalNodeType, host.ID), nil
			default:
				return "", fmt.Errorf("unknown host type: %s", host.Type)
			}
		}
	}

	return "", fmt.Errorf("can't find instance by name: %s", string(nodeName))
}

func (i *instances) InstanceType(ctx context.Context, nodeName types.NodeName) (string, error) {
	cloudInstances, err := i.client.CloudComputingInstances.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return "", fmt.Errorf("can't get list of cloud instances: %s", err.Error())
	}

	for _, instance := range cloudInstances {
		if instance.Name == string(nodeName) {
			return cloudInstanceType, nil
		}
	}

	hosts, err := i.client.Hosts.Collection().
		SetPerPage(100).
		SetParam(searchPatternParamKey, string(nodeName)).
		Collect(ctx)

	if err != nil {
		return "", fmt.Errorf("can't get list of hosts: %s", err.Error())
	}

	for _, host := range hosts {
		if host.Title == string(nodeName) {
			switch host.Type {
			case "dedicated_server":
				return dedicatedServerType, nil
			case "kubernetes_baremetal_node":
				return kubernetesBaremetalNodeType, nil
			default:
				return "", fmt.Errorf("unknown host type: %s", host.Type)
			}
		}
	}

	return "", fmt.Errorf("can't find instance by name: %s", string(nodeName))
}

func (i *instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	instanceType, _, err := parseProviderID(providerID)
	if err != nil {
		return "", err
	}

	return instanceType, nil
}

func (i *instances) AddSSHKeyToAllInstances(_ context.Context, _ string, _ []byte) error {
	return cloudprovider.NotImplemented
}

func (i *instances) CurrentNodeName(_ context.Context, hostname string) (types.NodeName, error) {
	return types.NodeName(hostname), nil
}

func (i *instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	instanceType, instanceID, err := parseProviderID(providerID)
	if err != nil {
		return false, err
	}

	switch instanceType {
	case cloudInstanceType:
		cloudInstance, err := i.client.CloudComputingInstances.Get(ctx, instanceID)
		if err != nil {
			if isNotFoundError(err) {
				return false, nil
			}

			return false, fmt.Errorf("can't get cloud instance: %s", err.Error())
		}

		return cloudInstance != nil, nil
	case dedicatedServerType:
		host, err := i.client.Hosts.GetDedicatedServer(ctx, instanceID)
		if err != nil {
			if isNotFoundError(err) {
				return false, nil
			}

			return false, fmt.Errorf("can't get dedicated server: %s", err.Error())
		}

		return host != nil, nil
	case kubernetesBaremetalNodeType:
		host, err := i.client.Hosts.GetKubernetesBaremetalNode(ctx, instanceID)
		if err != nil {
			if isNotFoundError(err) {
				return false, nil
			}

			return false, fmt.Errorf("can't get kubernetes baremetal node: %s", err.Error())
		}

		return host != nil, nil
	default:
		return false, fmt.Errorf("invalid instance type: %s", instanceType)
	}
}

func (i *instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	instanceType, instanceID, err := parseProviderID(providerID)
	if err != nil {
		return true, err
	}

	switch instanceType {
	case cloudInstanceType:
		cloudInstance, err := i.client.CloudComputingInstances.Get(ctx, instanceID)
		if err != nil {
			return true, fmt.Errorf("can't get cloud instance: %s", err.Error())
		}

		return cloudInstance.Status == "SWITCHED_OFF", nil
	case dedicatedServerType:
		host, err := i.client.Hosts.GetDedicatedServer(ctx, instanceID)
		if err != nil {
			return true, fmt.Errorf("can't get dedicated server: %s", err.Error())
		}

		return host.PowerStatus == "powered_off", nil
	case kubernetesBaremetalNodeType:
		host, err := i.client.Hosts.GetKubernetesBaremetalNode(ctx, instanceID)
		if err != nil {
			return true, fmt.Errorf("can't get kubernetes baremetal node %s", err.Error())
		}

		return host.PowerStatus == "powered_off", nil
	default:
		return true, fmt.Errorf("invalid instance type: %s", instanceType)
	}
}
