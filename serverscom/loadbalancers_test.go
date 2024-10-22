package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	serverscom_testing "github.com/serverscom/cloud-controller-manager/serverscom/testing"
	cli "github.com/serverscom/serverscom-go-client/pkg"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

func TestLoadBalancers_GetLoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"
	locationID := int64(1)

	balancer := cli.LoadBalancer{
		ID:   "a",
		Name: balancerName,
	}

	ctx := context.TODO()

	collection.EXPECT().SetPerPage(100).Return(collection)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection)
	collection.EXPECT().SetParam("type", "l4").Return(collection)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{balancer}, nil)

	service.EXPECT().Collection().Return(collection)
	service.EXPECT().GetL4LoadBalancer(ctx, "a").Return(&cli.L4LoadBalancer{Name: balancerName, Status: "active", ExternalAddresses: []string{"127.0.0.1", "127.0.0.2"}}, nil)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"

	balancerInterface := newLoadBalancers(client, &locationID)
	status, exists, err := balancerInterface.GetLoadBalancer(ctx, "cluster", &srv)

	g.Expect(err).To(BeNil())
	g.Expect(status).NotTo(BeNil())
	g.Expect(len(status.Ingress)).To(Equal(2))
	g.Expect(status.Ingress[0].IP).To(Equal("127.0.0.1"))
	g.Expect(status.Ingress[1].IP).To(Equal("127.0.0.2"))
	g.Expect(exists).To(Equal(true))
}

func TestLoadBalancers_GetLoadBalancerNonActive(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"
	locationID := int64(1)

	balancer := cli.LoadBalancer{
		ID:   "a",
		Name: balancerName,
	}

	ctx := context.TODO()

	collection.EXPECT().SetPerPage(100).Return(collection)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection)
	collection.EXPECT().SetParam("type", "l4").Return(collection)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{balancer}, nil)

	service.EXPECT().Collection().Return(collection)
	service.EXPECT().GetL4LoadBalancer(ctx, "a").Return(&cli.L4LoadBalancer{Name: balancerName, Status: "in_process", ExternalAddresses: []string{"127.0.0.1", "127.0.0.2"}}, nil)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"

	balancerInterface := newLoadBalancers(client, &locationID)
	status, exists, err := balancerInterface.GetLoadBalancer(ctx, "cluster", &srv)

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("load balancer is not active, current status: in_process"))
	g.Expect(status).To(BeNil())
	g.Expect(exists).To(Equal(true))
}

func TestLoadBalancers_GetLoadBalancerEmptyList(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"
	locationID := int64(1)

	ctx := context.TODO()

	collection.EXPECT().SetPerPage(100).Return(collection)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection)
	collection.EXPECT().SetParam("type", "l4").Return(collection)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{}, nil)

	service.EXPECT().Collection().Return(collection)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"

	balancerInterface := newLoadBalancers(client, &locationID)
	status, exists, err := balancerInterface.GetLoadBalancer(ctx, "cluster", &srv)

	g.Expect(err).To(BeNil())
	g.Expect(status).To(BeNil())
	g.Expect(exists).To(Equal(false))
}

func TestLoadBalancers_GetLoadBalancerName(t *testing.T) {
	g := NewGomegaWithT(t)

	locationID := int64(1)
	client := cli.NewClient("some")
	ctx := context.TODO()

	srv := v1.Service{}
	srv.UID = "123"

	balancerInterface := newLoadBalancers(client, &locationID)
	name := balancerInterface.GetLoadBalancerName(ctx, "cluster", &srv)

	g.Expect(name).To(Equal("service-cluster-a123"))
}

func TestLoadBalancers_GetLoadBalancerNameWithAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	locationID := int64(1)
	client := cli.NewClient("some")
	ctx := context.TODO()

	srv := v1.Service{}
	srv.UID = "123"
	srv.Annotations = map[string]string{}
	srv.Annotations[loadBalancerNameAnnotation] = "my-awesome-balancer"

	balancerInterface := newLoadBalancers(client, &locationID)
	name := balancerInterface.GetLoadBalancerName(ctx, "cluster", &srv)

	g.Expect(name).To(Equal("my-awesome-balancer"))
}

func TestLoadBalancers_EnsureLoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"
	locationID := int64(1)
	sharedCluster := true

	balancer := cli.LoadBalancer{
		ID:   "a",
		Name: balancerName,
	}

	l4Balancer := cli.L4LoadBalancer{
		ID:                "a",
		Name:              balancerName,
		Status:            "active",
		ExternalAddresses: []string{"127.0.0.1", "127.0.0.2"},
	}

	ctx := context.TODO()

	input := cli.L4LoadBalancerUpdateInput{
		Name:          &balancerName,
		SharedCluster: &sharedCluster,
		VHostZones: []cli.L4VHostZoneInput{
			{
				ID:            "k8s-nodes-80-tcp",
				UDP:           false,
				ProxyProtocol: false,
				Ports:         []int32{80},
				Description:   nil,
				UpstreamID:    "k8s-nodes-80-tcp",
			},
			{
				ID:            "k8s-nodes-11211-udp",
				UDP:           true,
				ProxyProtocol: false,
				Ports:         []int32{11211},
				Description:   nil,
				UpstreamID:    "k8s-nodes-11211-udp",
			},
		},
		UpstreamZones: []cli.L4UpstreamZoneInput{
			{
				ID:         "k8s-nodes-80-tcp",
				Method:     nil,
				UDP:        false,
				HCInterval: nil,
				HCJitter:   nil,
				Upstreams: []cli.L4UpstreamInput{
					{
						IP:     "127.0.0.100",
						Port:   30200,
						Weight: 1,
					},
				},
			},
			{
				ID:         "k8s-nodes-11211-udp",
				Method:     nil,
				UDP:        false,
				HCInterval: nil,
				HCJitter:   nil,
				Upstreams: []cli.L4UpstreamInput{
					{
						IP:     "127.0.0.100",
						Port:   30201,
						Weight: 1,
					},
				},
			},
		},
	}

	collection.EXPECT().SetPerPage(100).Return(collection)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection)
	collection.EXPECT().SetParam("type", "l4").Return(collection)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{balancer}, nil)

	service.EXPECT().Collection().Return(collection)
	service.EXPECT().GetL4LoadBalancer(ctx, "a").Return(&l4Balancer, nil)
	service.EXPECT().UpdateL4LoadBalancer(ctx, "a", input).Return(&l4Balancer, nil)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"
	srv.Spec.Ports = []v1.ServicePort{
		{Port: 80, Protocol: "TCP", NodePort: 30200},
		{Port: 11211, Protocol: "UDP", NodePort: 30201},
	}

	node := v1.Node{}
	node.Status.Addresses = []v1.NodeAddress{
		{Address: "node1.example.com", Type: v1.NodeHostName},
		{Address: "127.0.0.50", Type: v1.NodeExternalIP},
		{Address: "127.0.0.100", Type: v1.NodeInternalIP},
	}

	balancerInterface := newLoadBalancers(client, &locationID)
	status, err := balancerInterface.EnsureLoadBalancer(ctx, "cluster", &srv, []*v1.Node{&node})

	g.Expect(err).To(BeNil())
	g.Expect(status).NotTo(BeNil())
}

func TestLoadBalancers_EnsureLoadBalancerWithCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"
	locationID := int64(1)

	l4Balancer := cli.L4LoadBalancer{
		ID:                "a",
		Name:              balancerName,
		Status:            "active",
		ExternalAddresses: []string{"127.0.0.1", "127.0.0.2"},
	}

	ctx := context.TODO()

	input := cli.L4LoadBalancerCreateInput{
		Name:       balancerName,
		LocationID: locationID,
		VHostZones: []cli.L4VHostZoneInput{
			{
				ID:            "k8s-nodes-80-tcp",
				UDP:           false,
				ProxyProtocol: false,
				Ports:         []int32{80},
				Description:   nil,
				UpstreamID:    "k8s-nodes-80-tcp",
			},
			{
				ID:            "k8s-nodes-11211-udp",
				UDP:           true,
				ProxyProtocol: false,
				Ports:         []int32{11211},
				Description:   nil,
				UpstreamID:    "k8s-nodes-11211-udp",
			},
		},
		UpstreamZones: []cli.L4UpstreamZoneInput{
			{
				ID:         "k8s-nodes-80-tcp",
				Method:     nil,
				UDP:        false,
				HCInterval: nil,
				HCJitter:   nil,
				Upstreams: []cli.L4UpstreamInput{
					{
						IP:     "127.0.0.100",
						Port:   30200,
						Weight: 1,
					},
				},
			},
			{
				ID:         "k8s-nodes-11211-udp",
				Method:     nil,
				UDP:        false,
				HCInterval: nil,
				HCJitter:   nil,
				Upstreams: []cli.L4UpstreamInput{
					{
						IP:     "127.0.0.100",
						Port:   30201,
						Weight: 1,
					},
				},
			},
		},
	}

	collection.EXPECT().SetPerPage(100).Return(collection).Times(2)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection).Times(2)
	collection.EXPECT().SetParam("type", "l4").Return(collection).Times(2)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{}, nil).Times(2)

	service.EXPECT().Collection().Return(collection).Times(2)
	service.EXPECT().CreateL4LoadBalancer(ctx, input).Return(&l4Balancer, nil)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"
	srv.Spec.Ports = []v1.ServicePort{
		{Port: 80, Protocol: "TCP", NodePort: 30200},
		{Port: 11211, Protocol: "UDP", NodePort: 30201},
	}

	node := v1.Node{}
	node.Status.Addresses = []v1.NodeAddress{
		{Address: "node1.example.com", Type: v1.NodeHostName},
		{Address: "127.0.0.50", Type: v1.NodeExternalIP},
		{Address: "127.0.0.100", Type: v1.NodeInternalIP},
	}

	balancerInterface := newLoadBalancers(client, &locationID)
	status, err := balancerInterface.EnsureLoadBalancer(ctx, "cluster", &srv, []*v1.Node{&node})

	g.Expect(err).To(BeNil())
	g.Expect(status).NotTo(BeNil())

	// with cluster-id annotation
	clusterID := "some-hash-id"
	srv.Annotations = map[string]string{
		"servers.com/cluster-id": clusterID,
	}

	input.ClusterID = &clusterID
	service.EXPECT().CreateL4LoadBalancer(ctx, input).Return(&l4Balancer, nil)
	status, err = balancerInterface.EnsureLoadBalancer(ctx, "cluster", &srv, []*v1.Node{&node})

	g.Expect(err).To(BeNil())
	g.Expect(status).NotTo(BeNil())
}

func TestLoadBalancers_UpdateLoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"
	locationID := int64(1)
	sharedCluster := true

	balancer := cli.LoadBalancer{
		ID:   "a",
		Name: balancerName,
	}

	l4Balancer := cli.L4LoadBalancer{
		ID:                "a",
		Name:              balancerName,
		Status:            "active",
		ExternalAddresses: []string{"127.0.0.1", "127.0.0.2"},
	}

	ctx := context.TODO()

	input := cli.L4LoadBalancerUpdateInput{
		Name:          &balancerName,
		SharedCluster: &sharedCluster,
		VHostZones: []cli.L4VHostZoneInput{
			{
				ID:            "k8s-nodes-80-tcp",
				UDP:           false,
				ProxyProtocol: false,
				Ports:         []int32{80},
				Description:   nil,
				UpstreamID:    "k8s-nodes-80-tcp",
			},
			{
				ID:            "k8s-nodes-11211-udp",
				UDP:           true,
				ProxyProtocol: false,
				Ports:         []int32{11211},
				Description:   nil,
				UpstreamID:    "k8s-nodes-11211-udp",
			},
		},
		UpstreamZones: []cli.L4UpstreamZoneInput{
			{
				ID:         "k8s-nodes-80-tcp",
				Method:     nil,
				UDP:        false,
				HCInterval: nil,
				HCJitter:   nil,
				Upstreams: []cli.L4UpstreamInput{
					{
						IP:     "127.0.0.100",
						Port:   30200,
						Weight: 1,
					},
				},
			},
			{
				ID:         "k8s-nodes-11211-udp",
				Method:     nil,
				UDP:        false,
				HCInterval: nil,
				HCJitter:   nil,
				Upstreams: []cli.L4UpstreamInput{
					{
						IP:     "127.0.0.100",
						Port:   30201,
						Weight: 1,
					},
				},
			},
		},
	}

	collection.EXPECT().SetPerPage(100).Return(collection).Times(2)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection).Times(2)
	collection.EXPECT().SetParam("type", "l4").Return(collection).Times(2)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{balancer}, nil).Times(2)

	service.EXPECT().Collection().Return(collection).Times(2)
	service.EXPECT().GetL4LoadBalancer(ctx, "a").Return(&l4Balancer, nil).Times(2)
	service.EXPECT().UpdateL4LoadBalancer(ctx, "a", input).Return(&l4Balancer, nil)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"
	srv.Spec.Ports = []v1.ServicePort{
		{Port: 80, Protocol: "TCP", NodePort: 30200},
		{Port: 11211, Protocol: "UDP", NodePort: 30201},
	}

	node := v1.Node{}
	node.Status.Addresses = []v1.NodeAddress{
		{Address: "node1.example.com", Type: v1.NodeHostName},
		{Address: "127.0.0.50", Type: v1.NodeExternalIP},
		{Address: "127.0.0.100", Type: v1.NodeInternalIP},
	}

	balancerInterface := newLoadBalancers(client, &locationID)
	status, err := balancerInterface.EnsureLoadBalancer(ctx, "cluster", &srv, []*v1.Node{&node})

	g.Expect(err).To(BeNil())
	g.Expect(status).NotTo(BeNil())

	// with cluster-id annotation
	clusterID := "some-hash-id"
	srv.Annotations = map[string]string{
		"servers.com/cluster-id": clusterID,
	}

	input.ClusterID = &clusterID
	input.SharedCluster = nil
	service.EXPECT().UpdateL4LoadBalancer(ctx, "a", input).Return(&l4Balancer, nil)
	status, err = balancerInterface.EnsureLoadBalancer(ctx, "cluster", &srv, []*v1.Node{&node})

	g.Expect(err).To(BeNil())
	g.Expect(status).NotTo(BeNil())
}

func TestLoadBalancers_EnsureLoadBalancerDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"

	balancer := cli.LoadBalancer{
		ID:   "a",
		Name: balancerName,
	}

	ctx := context.TODO()

	collection.EXPECT().SetPerPage(100).Return(collection)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection)
	collection.EXPECT().SetParam("type", "l4").Return(collection)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{balancer}, nil)

	service.EXPECT().Collection().Return(collection)
	service.EXPECT().GetL4LoadBalancer(ctx, "a").Return(&cli.L4LoadBalancer{ID: "a", Name: balancerName, Status: "in_process", ExternalAddresses: []string{"127.0.0.1", "127.0.0.2"}}, nil)
	service.EXPECT().DeleteL4LoadBalancer(ctx, "a").Return(nil)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"
	srv.Spec.Ports = []v1.ServicePort{
		{Port: 80, Protocol: "TCP", NodePort: 30200},
		{Port: 11211, Protocol: "UDP", NodePort: 30201},
	}

	balancerInterface := newLoadBalancers(client, nil)
	err := balancerInterface.EnsureLoadBalancerDeleted(ctx, "cluster", &srv)

	g.Expect(err).To(BeNil())
}

func TestLoadBalancers_EnsureLoadBalancerDeletedWhenBalancerAlreadyDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collection := serverscom_testing.NewMockCollection[cli.LoadBalancer](ctrl)
	service := serverscom_testing.NewMockLoadBalancersService(ctrl)

	balancerName := "service-cluster-a123"

	ctx := context.TODO()

	collection.EXPECT().SetPerPage(100).Return(collection)
	collection.EXPECT().SetParam("search_pattern", balancerName).Return(collection)
	collection.EXPECT().SetParam("type", "l4").Return(collection)
	collection.EXPECT().Collect(ctx).Return([]cli.LoadBalancer{}, nil)

	service.EXPECT().Collection().Return(collection)

	client := cli.NewClient("some")
	client.LoadBalancers = service

	srv := v1.Service{}
	srv.UID = "123"
	srv.Spec.Ports = []v1.ServicePort{
		{Port: 80, Protocol: "TCP", NodePort: 30200},
		{Port: 11211, Protocol: "UDP", NodePort: 30201},
	}

	balancerInterface := newLoadBalancers(client, nil)
	err := balancerInterface.EnsureLoadBalancerDeleted(ctx, "cluster", &srv)

	g.Expect(err).To(BeNil())
}
