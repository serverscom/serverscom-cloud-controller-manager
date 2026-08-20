package serverscom

import (
	"testing"

	. "github.com/onsi/gomega"
	cli "github.com/serverscom/serverscom-go-client/pkg"
	v1 "k8s.io/api/core/v1"
)

func TestParseProviderID(t *testing.T) {
	tests := []struct {
		name       string
		providerID string
		expected   providerIDInfo
		errMsg     string
	}{
		{
			name:       "cloud instance",
			providerID: "serverscom://cloud-instance/G0aK9O7v",
			expected:   providerIDInfo{nodeType: cloudInstanceType, instanceID: "G0aK9O7v"},
		},
		{
			name:       "dedicated server",
			providerID: "serverscom://dedicated-server/a",
			expected:   providerIDInfo{nodeType: dedicatedServerType, instanceID: "a"},
		},
		{
			name:       "kubernetes baremetal node",
			providerID: "serverscom://kubernetes-baremetal-node/a",
			expected:   providerIDInfo{nodeType: kubernetesBaremetalNodeType, instanceID: "a"},
		},
		{
			name:       "underscores in node type are normalized",
			providerID: "serverscom://kubernetes_baremetal_node/a",
			expected:   providerIDInfo{nodeType: kubernetesBaremetalNodeType, instanceID: "a"},
		},
		{
			name:       "autoscale node",
			providerID: "serverscom://kubernetes-autoscale-node/X4bR2mQ9/9V8Mjm2p",
			expected: providerIDInfo{
				nodeType:   kubernetesAutoscaleNodeType,
				clusterID:  "X4bR2mQ9",
				instanceID: "9V8Mjm2p",
			},
		},
		{
			name:       "autoscale node with underscores in node type",
			providerID: "serverscom://kubernetes_autoscale_node/X4bR2mQ9/9V8Mjm2p",
			expected: providerIDInfo{
				nodeType:   kubernetesAutoscaleNodeType,
				clusterID:  "X4bR2mQ9",
				instanceID: "9V8Mjm2p",
			},
		},
		{
			name:       "unknown node type is passed through",
			providerID: "serverscom://some-new-type/a",
			expected:   providerIDInfo{nodeType: "some-new-type", instanceID: "a"},
		},
		{
			name:       "autoscale node without cluster id",
			providerID: "serverscom://kubernetes-autoscale-node/9V8Mjm2p",
			errMsg:     `providerID "serverscom://kubernetes-autoscale-node/9V8Mjm2p" for node type kubernetes-autoscale-node is missing the cluster ID, expected serverscom://kubernetes-autoscale-node/<cluster_id>/<node_id>`,
		},
		{
			name:       "four part form is rejected for non autoscale types",
			providerID: "serverscom://cloud-instance/X4bR2mQ9/9V8Mjm2p",
			errMsg:     "error splitting providerID: serverscom://cloud-instance/X4bR2mQ9/9V8Mjm2p",
		},
		{
			name:       "missing prefix",
			providerID: "cloud-instance/a",
			errMsg:     "missing prefix serverscom://: cloud-instance/a",
		},
		{
			name:       "wrong prefix",
			providerID: "aws://cloud-instance/a",
			errMsg:     "missing prefix serverscom://: aws://cloud-instance/a",
		},
		{
			name:       "node type only",
			providerID: "serverscom://cloud-instance",
			errMsg:     "error splitting providerID: serverscom://cloud-instance",
		},
		{
			name:       "too many segments",
			providerID: "serverscom://kubernetes-autoscale-node/a/b/c",
			errMsg:     "error splitting providerID: serverscom://kubernetes-autoscale-node/a/b/c",
		},
		{
			name:       "empty instance id",
			providerID: "serverscom://cloud-instance/",
			errMsg:     "error splitting providerID: serverscom://cloud-instance/",
		},
		{
			name:       "empty cluster id",
			providerID: "serverscom://kubernetes-autoscale-node//9V8Mjm2p",
			errMsg:     "error splitting providerID: serverscom://kubernetes-autoscale-node//9V8Mjm2p",
		},
		{
			name:       "empty node type",
			providerID: "serverscom:///a",
			errMsg:     "error splitting providerID: serverscom:///a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			info, err := parseProviderID(tt.providerID)

			if tt.errMsg != "" {
				g.Expect(err).NotTo(BeNil())
				g.Expect(err.Error()).To(Equal(tt.errMsg))
				g.Expect(info).To(Equal(providerIDInfo{}))

				return
			}

			g.Expect(err).To(BeNil())
			g.Expect(info).To(Equal(tt.expected))
		})
	}
}

func TestCollectKubernetesClusterNodeAddresses(t *testing.T) {
	tests := []struct {
		name     string
		node     cli.KubernetesClusterNode
		expected []v1.NodeAddress
	}{
		{
			name: "hostname and both addresses",
			node: cli.KubernetesClusterNode{
				Hostname:           "my-super-node1",
				PrivateIPv4Address: "127.0.0.1",
				PublicIPv4Address:  "127.0.0.2",
			},
			expected: []v1.NodeAddress{
				{Address: "my-super-node1", Type: v1.NodeHostName},
				{Address: "127.0.0.1", Type: v1.NodeInternalIP},
				{Address: "127.0.0.2", Type: v1.NodeExternalIP},
			},
		},
		{
			name: "private address only",
			node: cli.KubernetesClusterNode{
				Hostname:           "my-super-node1",
				PrivateIPv4Address: "127.0.0.1",
			},
			expected: []v1.NodeAddress{
				{Address: "my-super-node1", Type: v1.NodeHostName},
				{Address: "127.0.0.1", Type: v1.NodeInternalIP},
			},
		},
		{
			name: "public address only",
			node: cli.KubernetesClusterNode{
				Hostname:          "my-super-node1",
				PublicIPv4Address: "127.0.0.2",
			},
			expected: []v1.NodeAddress{
				{Address: "my-super-node1", Type: v1.NodeHostName},
				{Address: "127.0.0.2", Type: v1.NodeExternalIP},
			},
		},
		{
			name: "hostname only",
			node: cli.KubernetesClusterNode{Hostname: "my-super-node1"},
			expected: []v1.NodeAddress{
				{Address: "my-super-node1", Type: v1.NodeHostName},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			g.Expect(collectKubernetesClusterNodeAddresses(&tt.node)).To(Equal(tt.expected))
		})
	}
}
