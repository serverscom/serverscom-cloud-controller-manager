package serverscom

import (
	"fmt"
	"io"
	"os"
	"strconv"

	cli "github.com/serverscom/serverscom-go-client/pkg"
	cloudprovider "k8s.io/cloud-provider"
)

const (
	providerName = "serverscom"

	tokenEnvKey             = "SERVERSCOM_TOKEN"
	baseUrkEnvKey           = "SERVERSCOM_BASE_URL"
	defaultLocationIdEnvKey = "SERVERSCOM_DEFAULT_LOCATION_ID"
	defaultZoneEnvKey       = "SERVERSCOM_DEFAULT_ZONE"

	controllerVersion = "base"
)

func init() {
	cloudprovider.RegisterCloudProvider(providerName, func(config io.Reader) (cloudprovider.Interface, error) {
		return newCloud(config)
	})
}

type cloud struct {
	client *cli.Client

	loadBalancers cloudprovider.LoadBalancer
	instances     cloudprovider.Instances
	zones         cloudprovider.Zones
}

func newCloud(config io.Reader) (cloudprovider.Interface, error) {
	token := os.Getenv(tokenEnvKey)
	if token == "" {
		return nil, fmt.Errorf("environment variable %q is required", tokenEnvKey)
	}

	baseUrl := os.Getenv(baseUrkEnvKey)

	var client *cli.Client

	if baseUrl == "" {
		client = cli.NewClient(token)
	} else {
		client = cli.NewClientWithEndpoint(token, baseUrl)
	}

	client.SetupUserAgent(fmt.Sprintf("serverscom-cloud-controller-manager/%s", controllerVersion))

	defaultLocationIDStr := os.Getenv(defaultLocationIdEnvKey)
	var defaultLocationID *int64

	if defaultLocationIDStr != "" {
		n, err := strconv.ParseInt(defaultLocationIDStr, 10, 64)
		if err == nil {
			return nil, fmt.Errorf("invalid %s: %s", defaultLocationIdEnvKey, err.Error())
		}

		defaultLocationID = &n
	}

	loadBalancers := newLoadBalancers(client, defaultLocationID)
	instances := newInstances(client)
	zones := newZones(client, os.Getenv(defaultZoneEnvKey))

	return &cloud{
		client:        client,
		loadBalancers: loadBalancers,
		instances:     instances,
		zones:         zones,
	}, nil
}

func (c *cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
}

func (c *cloud) Instances() (cloudprovider.Instances, bool) {
	return c.instances, true
}

func (c *cloud) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return nil, false
}

func (c *cloud) Zones() (cloudprovider.Zones, bool) {
	return c.zones, true
}

func (c *cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return c.loadBalancers, true
}

func (c *cloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

func (c *cloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

func (c *cloud) ProviderName() string {
	return providerName
}

func (c *cloud) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nil, nil
}

func (c *cloud) HasClusterID() bool {
	return false
}
