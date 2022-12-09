package serverscom

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	serverscom_testing "github.com/serverscom/cloud-controller-manager/serverscom/testing"
	cli "github.com/serverscom/serverscom-go-client/pkg"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
)

func TestInstances_NodeAddressesWithCloudInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"
	publicIPv6 := "0:0:0:0:0:0:0:1"

	cloudInstance := cli.CloudComputingInstance{
		ID:                 "a",
		Name:               nodeName,
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
		PublicIPv6Address:  &publicIPv6,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	service := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{cloudInstance}, nil)

	service.EXPECT().Collection().Return(instanceCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = service

	instances := newInstances(client)
	addresses, err := instances.NodeAddresses(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(addresses).NotTo(BeNil())
	g.Expect(len(addresses)).To(Equal(4))

	g.Expect(addresses[0].Address).To(Equal(nodeName))
	g.Expect(addresses[0].Type).To(Equal(v1.NodeHostName))

	g.Expect(addresses[1].Address).To(Equal(privateIPv4))
	g.Expect(addresses[1].Type).To(Equal(v1.NodeInternalIP))

	g.Expect(addresses[2].Address).To(Equal(publicIPv4))
	g.Expect(addresses[2].Type).To(Equal(v1.NodeExternalIP))

	g.Expect(addresses[3].Address).To(Equal(publicIPv6))
	g.Expect(addresses[3].Type).To(Equal(v1.NodeExternalIP))
}

func TestInstances_NodeAddressesWithHost(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"

	host := cli.Host{
		ID:                 "a",
		Title:              nodeName,
		Type:               "dedicated_server",
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	hostsCollection := serverscom_testing.NewMockCollection[cli.Host](ctrl)

	instancesService := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	hostsService := serverscom_testing.NewMockHostsService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{}, nil)

	hostsCollection.EXPECT().SetPerPage(100).Return(hostsCollection)
	hostsCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(hostsCollection)
	hostsCollection.EXPECT().Collect(ctx).Return([]cli.Host{host}, nil)

	instancesService.EXPECT().Collection().Return(instanceCollection)
	hostsService.EXPECT().Collection().Return(hostsCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = instancesService
	client.Hosts = hostsService

	instances := newInstances(client)
	addresses, err := instances.NodeAddresses(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(addresses).NotTo(BeNil())
	g.Expect(len(addresses)).To(Equal(3))

	g.Expect(addresses[0].Address).To(Equal(nodeName))
	g.Expect(addresses[0].Type).To(Equal(v1.NodeHostName))

	g.Expect(addresses[1].Address).To(Equal(privateIPv4))
	g.Expect(addresses[1].Type).To(Equal(v1.NodeInternalIP))

	g.Expect(addresses[2].Address).To(Equal(publicIPv4))
	g.Expect(addresses[2].Type).To(Equal(v1.NodeExternalIP))
}

func TestInstances_NodeAddressesNotFound(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	hostsCollection := serverscom_testing.NewMockCollection[cli.Host](ctrl)

	instancesService := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	hostsService := serverscom_testing.NewMockHostsService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{}, nil)

	hostsCollection.EXPECT().SetPerPage(100).Return(hostsCollection)
	hostsCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(hostsCollection)
	hostsCollection.EXPECT().Collect(ctx).Return([]cli.Host{}, nil)

	instancesService.EXPECT().Collection().Return(instanceCollection)
	hostsService.EXPECT().Collection().Return(hostsCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = instancesService
	client.Hosts = hostsService

	instances := newInstances(client)
	addresses, err := instances.NodeAddresses(ctx, types.NodeName(nodeName))

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("can't find instance by name: my-super-node1"))
	g.Expect(addresses).To(BeNil())
}

func TestInstances_NodeAddressesByProviderIDWithCloudInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"
	publicIPv6 := "0:0:0:0:0:0:0:1"

	cloudInstance := cli.CloudComputingInstance{
		ID:                 "a",
		Name:               nodeName,
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
		PublicIPv6Address:  &publicIPv6,
	}

	instancesService := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	instancesService.EXPECT().Get(ctx, "a").Return(&cloudInstance, nil)

	client := cli.NewClient("some")
	client.CloudComputingInstances = instancesService

	instances := newInstances(client)
	addresses, err := instances.NodeAddressesByProviderID(ctx, "serverscom://cloud-instance/a")

	g.Expect(err).To(BeNil())
	g.Expect(addresses).NotTo(BeNil())
	g.Expect(len(addresses)).To(Equal(4))

	g.Expect(addresses[0].Address).To(Equal(nodeName))
	g.Expect(addresses[0].Type).To(Equal(v1.NodeHostName))

	g.Expect(addresses[1].Address).To(Equal(privateIPv4))
	g.Expect(addresses[1].Type).To(Equal(v1.NodeInternalIP))

	g.Expect(addresses[2].Address).To(Equal(publicIPv4))
	g.Expect(addresses[2].Type).To(Equal(v1.NodeExternalIP))

	g.Expect(addresses[3].Address).To(Equal(publicIPv6))
	g.Expect(addresses[3].Type).To(Equal(v1.NodeExternalIP))
}

func TestInstances_NodeAddressesByProviderIDWithDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"

	dedicatedServer := cli.DedicatedServer{
		ID:                 "a",
		Title:              nodeName,
		Type:               "dedicated_server",
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
	}

	hostsService := serverscom_testing.NewMockHostsService(ctrl)
	hostsService.EXPECT().GetDedicatedServer(ctx, "a").Return(&dedicatedServer, nil)

	client := cli.NewClient("some")
	client.Hosts = hostsService

	instances := newInstances(client)
	addresses, err := instances.NodeAddressesByProviderID(ctx, "serverscom://dedicated-server/a")

	g.Expect(err).To(BeNil())
	g.Expect(addresses).NotTo(BeNil())
	g.Expect(len(addresses)).To(Equal(3))

	g.Expect(addresses[0].Address).To(Equal(nodeName))
	g.Expect(addresses[0].Type).To(Equal(v1.NodeHostName))

	g.Expect(addresses[1].Address).To(Equal(privateIPv4))
	g.Expect(addresses[1].Type).To(Equal(v1.NodeInternalIP))

	g.Expect(addresses[2].Address).To(Equal(publicIPv4))
	g.Expect(addresses[2].Type).To(Equal(v1.NodeExternalIP))
}

func TestInstances_NodeAddressesByProviderIDWithKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"

	kubernetesBaremetalNode := cli.KubernetesBaremetalNode{
		ID:                 "a",
		Title:              nodeName,
		Type:               "kubernetes_baremetal_node",
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
	}

	hostsService := serverscom_testing.NewMockHostsService(ctrl)
	hostsService.EXPECT().GetKubernetesBaremetalNode(ctx, "a").Return(&kubernetesBaremetalNode, nil)

	client := cli.NewClient("some")
	client.Hosts = hostsService

	instances := newInstances(client)
	addresses, err := instances.NodeAddressesByProviderID(ctx, "serverscom://kubernetes-baremetal-node/a")

	g.Expect(err).To(BeNil())
	g.Expect(addresses).NotTo(BeNil())
	g.Expect(len(addresses)).To(Equal(3))

	g.Expect(addresses[0].Address).To(Equal(nodeName))
	g.Expect(addresses[0].Type).To(Equal(v1.NodeHostName))

	g.Expect(addresses[1].Address).To(Equal(privateIPv4))
	g.Expect(addresses[1].Type).To(Equal(v1.NodeInternalIP))

	g.Expect(addresses[2].Address).To(Equal(publicIPv4))
	g.Expect(addresses[2].Type).To(Equal(v1.NodeExternalIP))
}

func TestInstances_InstanceIDWithCloudInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"
	publicIPv6 := "0:0:0:0:0:0:0:1"

	cloudInstance := cli.CloudComputingInstance{
		ID:                 "a",
		Name:               nodeName,
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
		PublicIPv6Address:  &publicIPv6,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	service := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{cloudInstance}, nil)

	service.EXPECT().Collection().Return(instanceCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = service

	instances := newInstances(client)
	providerID, err := instances.InstanceID(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(providerID).To(Equal("cloud-instance/a"))
}

func TestInstances_InstanceIDWithDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"

	host := cli.Host{
		ID:                 "a",
		Title:              nodeName,
		Type:               "dedicated_server",
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	hostsCollection := serverscom_testing.NewMockCollection[cli.Host](ctrl)

	instancesService := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	hostsService := serverscom_testing.NewMockHostsService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{}, nil)

	hostsCollection.EXPECT().SetPerPage(100).Return(hostsCollection)
	hostsCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(hostsCollection)
	hostsCollection.EXPECT().Collect(ctx).Return([]cli.Host{host}, nil)

	instancesService.EXPECT().Collection().Return(instanceCollection)
	hostsService.EXPECT().Collection().Return(hostsCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = instancesService
	client.Hosts = hostsService

	instances := newInstances(client)
	providerID, err := instances.InstanceID(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(providerID).To(Equal("dedicated-server/a"))
}

func TestInstances_InstanceIDWithKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"

	host := cli.Host{
		ID:                 "a",
		Title:              nodeName,
		Type:               "kubernetes_baremetal_node",
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	hostsCollection := serverscom_testing.NewMockCollection[cli.Host](ctrl)

	instancesService := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	hostsService := serverscom_testing.NewMockHostsService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{}, nil)

	hostsCollection.EXPECT().SetPerPage(100).Return(hostsCollection)
	hostsCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(hostsCollection)
	hostsCollection.EXPECT().Collect(ctx).Return([]cli.Host{host}, nil)

	instancesService.EXPECT().Collection().Return(instanceCollection)
	hostsService.EXPECT().Collection().Return(hostsCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = instancesService
	client.Hosts = hostsService

	instances := newInstances(client)
	providerID, err := instances.InstanceID(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(providerID).To(Equal("kubernetes-baremetal-node/a"))
}

func TestInstances_InstanceTypeWithCloudInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"
	publicIPv6 := "0:0:0:0:0:0:0:1"

	cloudInstance := cli.CloudComputingInstance{
		ID:                 "a",
		Name:               nodeName,
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
		PublicIPv6Address:  &publicIPv6,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	service := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{cloudInstance}, nil)

	service.EXPECT().Collection().Return(instanceCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = service

	instances := newInstances(client)
	instanceType, err := instances.InstanceType(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(instanceType).To(Equal("cloud-instance"))
}

func TestInstances_InstanceTypeWithDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"

	host := cli.Host{
		ID:                 "a",
		Title:              nodeName,
		Type:               "dedicated_server",
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	hostsCollection := serverscom_testing.NewMockCollection[cli.Host](ctrl)

	instancesService := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	hostsService := serverscom_testing.NewMockHostsService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{}, nil)

	hostsCollection.EXPECT().SetPerPage(100).Return(hostsCollection)
	hostsCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(hostsCollection)
	hostsCollection.EXPECT().Collect(ctx).Return([]cli.Host{host}, nil)

	instancesService.EXPECT().Collection().Return(instanceCollection)
	hostsService.EXPECT().Collection().Return(hostsCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = instancesService
	client.Hosts = hostsService

	instances := newInstances(client)
	instanceType, err := instances.InstanceType(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(instanceType).To(Equal("dedicated-server"))
}

func TestInstances_InstanceTypeWithKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	nodeName := "my-super-node1"
	ctx := context.TODO()

	privateIPv4 := "127.0.0.1"
	publicIPv4 := "127.0.0.2"

	host := cli.Host{
		ID:                 "a",
		Title:              nodeName,
		Type:               "kubernetes_baremetal_node",
		PrivateIPv4Address: &privateIPv4,
		PublicIPv4Address:  &publicIPv4,
	}

	instanceCollection := serverscom_testing.NewMockCollection[cli.CloudComputingInstance](ctrl)
	hostsCollection := serverscom_testing.NewMockCollection[cli.Host](ctrl)

	instancesService := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	hostsService := serverscom_testing.NewMockHostsService(ctrl)

	instanceCollection.EXPECT().SetPerPage(100).Return(instanceCollection)
	instanceCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(instanceCollection)
	instanceCollection.EXPECT().Collect(ctx).Return([]cli.CloudComputingInstance{}, nil)

	hostsCollection.EXPECT().SetPerPage(100).Return(hostsCollection)
	hostsCollection.EXPECT().SetParam(searchPatternParamKey, nodeName).Return(hostsCollection)
	hostsCollection.EXPECT().Collect(ctx).Return([]cli.Host{host}, nil)

	instancesService.EXPECT().Collection().Return(instanceCollection)
	hostsService.EXPECT().Collection().Return(hostsCollection)

	client := cli.NewClient("some")
	client.CloudComputingInstances = instancesService
	client.Hosts = hostsService

	instances := newInstances(client)
	instanceType, err := instances.InstanceType(ctx, types.NodeName(nodeName))

	g.Expect(err).To(BeNil())
	g.Expect(instanceType).To(Equal("kubernetes-baremetal-node"))
}

func TestInstances_InstanceTypeByProviderIDWithCloudInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	client := cli.NewClient("some")

	instances := newInstances(client)
	instanceType, err := instances.InstanceTypeByProviderID(context.TODO(), "serverscom://cloud-instance/a")

	g.Expect(err).To(BeNil())
	g.Expect(instanceType).To(Equal("cloud-instance"))
}

func TestInstances_InstanceTypeByProviderIDWithDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	client := cli.NewClient("some")

	instances := newInstances(client)
	instanceType, err := instances.InstanceTypeByProviderID(context.TODO(), "serverscom://dedicated-server/a")

	g.Expect(err).To(BeNil())
	g.Expect(instanceType).To(Equal("dedicated-server"))
}

func TestInstances_InstanceTypeByProviderIDWithKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	client := cli.NewClient("some")

	instances := newInstances(client)
	instanceType, err := instances.InstanceTypeByProviderID(context.TODO(), "serverscom://kubernetes-baremetal-node/a")

	g.Expect(err).To(BeNil())
	g.Expect(instanceType).To(Equal("kubernetes-baremetal-node"))
}

func TestInstances_AddSSHKeyToAllInstances(t *testing.T) {
	g := NewGomegaWithT(t)

	client := cli.NewClient("some")

	instances := newInstances(client)
	err := instances.AddSSHKeyToAllInstances(context.TODO(), "root", []byte{})

	g.Expect(err).NotTo(BeNil())
	g.Expect(err).To(Equal(cloudprovider.NotImplemented))
}

func TestInstances_CurrentNodeName(t *testing.T) {
	g := NewGomegaWithT(t)

	client := cli.NewClient("some")

	instances := newInstances(client)
	nodeName, err := instances.CurrentNodeName(context.TODO(), "my-super-node1")

	g.Expect(err).To(BeNil())
	g.Expect(nodeName).To(Equal(types.NodeName("my-super-node1")))
}

func TestInstances_InstanceExistsByProviderIDWithCloudInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	cloudInstance := cli.CloudComputingInstance{ID: "a"}

	service := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	service.EXPECT().Get(ctx, "a").Return(&cloudInstance, nil)
	service.EXPECT().Get(ctx, "b").Return(nil, &cli.NotFoundError{})

	client := cli.NewClient("some")
	client.CloudComputingInstances = service

	instances := newInstances(client)
	exists, err := instances.InstanceExistsByProviderID(ctx, "serverscom://cloud-instance/a")

	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(true))

	exists, err = instances.InstanceExistsByProviderID(ctx, "serverscom://cloud-instance/b")

	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(false))
}

func TestInstances_InstanceExistsByProviderIDWithDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	dedicatedServer := cli.DedicatedServer{ID: "a"}

	service := serverscom_testing.NewMockHostsService(ctrl)
	service.EXPECT().GetDedicatedServer(ctx, "a").Return(&dedicatedServer, nil)
	service.EXPECT().GetDedicatedServer(ctx, "b").Return(nil, &cli.NotFoundError{})

	client := cli.NewClient("some")
	client.Hosts = service

	instances := newInstances(client)
	exists, err := instances.InstanceExistsByProviderID(ctx, "serverscom://dedicated-server/a")

	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(true))

	exists, err = instances.InstanceExistsByProviderID(ctx, "serverscom://dedicated-server/b")

	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(false))
}

func TestInstances_InstanceExistsByProviderIDWithKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	kubernetesBaremetalNode := cli.KubernetesBaremetalNode{ID: "a"}

	service := serverscom_testing.NewMockHostsService(ctrl)
	service.EXPECT().GetKubernetesBaremetalNode(ctx, "a").Return(&kubernetesBaremetalNode, nil)
	service.EXPECT().GetKubernetesBaremetalNode(ctx, "b").Return(nil, &cli.NotFoundError{})

	client := cli.NewClient("some")
	client.Hosts = service

	instances := newInstances(client)
	exists, err := instances.InstanceExistsByProviderID(ctx, "serverscom://kubernetes-baremetal-node/a")

	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(true))

	exists, err = instances.InstanceExistsByProviderID(ctx, "serverscom://kubernetes-baremetal-node/b")

	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(false))
}

func TestInstances_InstanceShutdownByProviderIDWithCloudInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	cloudInstanceA := cli.CloudComputingInstance{ID: "a", Status: "SWITCHED_OFF"}
	cloudInstanceB := cli.CloudComputingInstance{ID: "b", Status: "ACTIVE"}

	service := serverscom_testing.NewMockCloudComputingInstancesService(ctrl)
	service.EXPECT().Get(ctx, "a").Return(&cloudInstanceA, nil)
	service.EXPECT().Get(ctx, "b").Return(&cloudInstanceB, nil)

	client := cli.NewClient("some")
	client.CloudComputingInstances = service

	instances := newInstances(client)
	isShutdown, err := instances.InstanceShutdownByProviderID(ctx, "serverscom://cloud-instance/a")

	g.Expect(err).To(BeNil())
	g.Expect(isShutdown).To(Equal(true))

	isShutdown, err = instances.InstanceShutdownByProviderID(ctx, "serverscom://cloud-instance/b")

	g.Expect(err).To(BeNil())
	g.Expect(isShutdown).To(Equal(false))
}

func TestInstances_InstanceShutdownByProviderIDWithDedicatedServer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	dedicatedServerA := cli.DedicatedServer{ID: "a", PowerStatus: "powered_off"}
	dedicatedServerB := cli.DedicatedServer{ID: "b", PowerStatus: "powered_on"}

	service := serverscom_testing.NewMockHostsService(ctrl)
	service.EXPECT().GetDedicatedServer(ctx, "a").Return(&dedicatedServerA, nil)
	service.EXPECT().GetDedicatedServer(ctx, "b").Return(&dedicatedServerB, nil)

	client := cli.NewClient("some")
	client.Hosts = service

	instances := newInstances(client)
	isShutdown, err := instances.InstanceShutdownByProviderID(ctx, "serverscom://dedicated-server/a")

	g.Expect(err).To(BeNil())
	g.Expect(isShutdown).To(Equal(true))

	isShutdown, err = instances.InstanceShutdownByProviderID(ctx, "serverscom://dedicated-server/b")

	g.Expect(err).To(BeNil())
	g.Expect(isShutdown).To(Equal(false))
}

func TestInstances_InstanceShutdownByProviderIDWithKubernetesBaremetalNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	kubernetesBaremetalNodeA := cli.KubernetesBaremetalNode{ID: "a", PowerStatus: "powered_off"}
	kubernetesBaremetalNodeB := cli.KubernetesBaremetalNode{ID: "b", PowerStatus: "powered_on"}

	service := serverscom_testing.NewMockHostsService(ctrl)
	service.EXPECT().GetKubernetesBaremetalNode(ctx, "a").Return(&kubernetesBaremetalNodeA, nil)
	service.EXPECT().GetKubernetesBaremetalNode(ctx, "b").Return(&kubernetesBaremetalNodeB, nil)

	client := cli.NewClient("some")
	client.Hosts = service

	instances := newInstances(client)
	isShutdown, err := instances.InstanceShutdownByProviderID(ctx, "serverscom://kubernetes-baremetal-node/a")

	g.Expect(err).To(BeNil())
	g.Expect(isShutdown).To(Equal(true))

	isShutdown, err = instances.InstanceShutdownByProviderID(ctx, "serverscom://kubernetes-baremetal-node/b")

	g.Expect(err).To(BeNil())
	g.Expect(isShutdown).To(Equal(false))
}
