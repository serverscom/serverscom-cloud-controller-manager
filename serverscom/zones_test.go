package serverscom

import (
	"context"
	"errors"
	"testing"

	. "github.com/onsi/gomega"
	serverscom_testing "github.com/serverscom/cloud-controller-manager/serverscom/testing"
	cli "github.com/serverscom/serverscom-go-client/pkg"
	gomock "go.uber.org/mock/gomock"
)

func TestZones_GetZoneByProviderIDWithAutoscaleNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	node := cli.KubernetesClusterNode{
		ID:           "9V8Mjm2p",
		ClusterID:    "X4bR2mQ9",
		Hostname:     "my-super-node1",
		LocationCode: "ams1",
	}

	clustersService := serverscom_testing.NewMockKubernetesClustersService(ctrl)
	clustersService.EXPECT().GetNode(ctx, "X4bR2mQ9", "9V8Mjm2p").Return(&node, nil)

	client := cli.NewClient("some")
	client.KubernetesClusters = clustersService

	zones := newZones(client, "default-zone")
	zone, err := zones.GetZoneByProviderID(ctx, "serverscom://kubernetes-autoscale-node/X4bR2mQ9/9V8Mjm2p")

	g.Expect(err).To(BeNil())
	g.Expect(zone.Region).To(Equal("ams1"))
}

func TestZones_GetZoneByProviderIDWithAutoscaleNodeError(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	clustersService := serverscom_testing.NewMockKubernetesClustersService(ctrl)
	clustersService.EXPECT().GetNode(ctx, "X4bR2mQ9", "9V8Mjm2p").Return(nil, errors.New("some error"))

	client := cli.NewClient("some")
	client.KubernetesClusters = clustersService

	zones := newZones(client, "default-zone")
	zone, err := zones.GetZoneByProviderID(ctx, "serverscom://kubernetes-autoscale-node/X4bR2mQ9/9V8Mjm2p")

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("can't get kubernetes autoscale node: some error"))
	g.Expect(zone.Region).To(BeEmpty())
}

func TestZones_GetZoneByProviderIDWithAutoscaleNodeMissingClusterID(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	// no API call is expected: the cluster ID is required to address the node
	clustersService := serverscom_testing.NewMockKubernetesClustersService(ctrl)

	client := cli.NewClient("some")
	client.KubernetesClusters = clustersService

	zones := newZones(client, "default-zone")
	zone, err := zones.GetZoneByProviderID(ctx, "serverscom://kubernetes-autoscale-node/9V8Mjm2p")

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal(`providerID "serverscom://kubernetes-autoscale-node/9V8Mjm2p" for node type kubernetes-autoscale-node is missing the cluster ID, expected serverscom://kubernetes-autoscale-node/<cluster_id>/<node_id>`))
	g.Expect(zone.Region).To(BeEmpty())
}

func TestZones_GetZoneByProviderIDWithUnknownType(t *testing.T) {
	g := NewGomegaWithT(t)

	ctx := context.TODO()

	zones := newZones(cli.NewClient("some"), "default-zone")
	zone, err := zones.GetZoneByProviderID(ctx, "serverscom://some-new-type/a")

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("invalid instance type: some-new-type"))
	g.Expect(zone.Region).To(BeEmpty())
}
