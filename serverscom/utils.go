package serverscom

import (
	"fmt"
	"k8s.io/klog/v2"
	"regexp"
	"strings"

	cli "github.com/serverscom/serverscom-go-client/pkg"
	v1 "k8s.io/api/core/v1"
)

const (
	cloudInstanceType           = "cloud-instance"
	dedicatedServerType         = "dedicated-server"
	kubernetesBaremetalNodeType = "kubernetes-baremetal-node"

	searchPatternParamKey = "search_pattern"
	typeParamKey          = "type"
)

var providerIDRE = regexp.MustCompile(`^` + providerName + `://([^/]+)/([^/]+)$`)

func isNotFoundError(err error) bool {
	switch err.(type) {
	case *cli.NotFoundError:
		return true
	default:
		return false
	}
}

func parseProviderID(providerID string) (string, string, error) {
	providerPrefix := providerName + "://"

	if !strings.HasPrefix(providerID, providerPrefix) {
		klog.Infof(" make sure your cluster configured for an external cloud provider")
		return "", "", fmt.Errorf("missing prefix %s: %s", providerPrefix, providerID)
	}

	matches := providerIDRE.FindStringSubmatch(providerID)

	if len(matches) != 3 {
		return "", "", fmt.Errorf("error splitting providerID: %s", providerID)
	}

	return strings.ReplaceAll(matches[1], "_", "-"), matches[2], nil
}

func collectCloudInstanceAddresses(cloudInstance *cli.CloudComputingInstance) []v1.NodeAddress {
	var addresses []v1.NodeAddress

	addresses = append(addresses, v1.NodeAddress{Address: cloudInstance.Name, Type: v1.NodeHostName})

	if cloudInstance.PrivateIPv4Address != nil {
		addresses = append(
			addresses,
			v1.NodeAddress{
				Address: *cloudInstance.PrivateIPv4Address,
				Type:    v1.NodeInternalIP,
			})
	}

	if cloudInstance.PublicIPv4Address != nil {
		addresses = append(
			addresses,
			v1.NodeAddress{
				Address: *cloudInstance.PublicIPv4Address,
				Type:    v1.NodeExternalIP,
			})
	}

	if cloudInstance.PublicIPv6Address != nil {
		addresses = append(
			addresses,
			v1.NodeAddress{
				Address: *cloudInstance.PublicIPv6Address,
				Type:    v1.NodeExternalIP,
			})
	}

	return addresses
}

func collectHostAddresses(host *cli.Host) []v1.NodeAddress {
	var addresses []v1.NodeAddress

	addresses = append(addresses, v1.NodeAddress{Address: host.Title, Type: v1.NodeHostName})

	if host.PrivateIPv4Address != nil {
		addresses = append(
			addresses,
			v1.NodeAddress{
				Address: *host.PrivateIPv4Address,
				Type:    v1.NodeInternalIP,
			})
	}

	if host.PublicIPv4Address != nil {
		addresses = append(
			addresses,
			v1.NodeAddress{
				Address: *host.PublicIPv4Address,
				Type:    v1.NodeExternalIP,
			})
	}

	return addresses
}

func buildExternalID(instanceType, ID string) string {
	return fmt.Sprintf("%s/%s", instanceType, ID)
}

func getLoadBalancerName(srv *v1.Service, clusterName string) string {
	ret := "a" + string(srv.UID)
	ret = strings.Replace(ret, "-", "", -1)
	if len(ret) > 32 {
		ret = ret[:32]
	}
	return fmt.Sprintf("service-%s-%s", clusterName, ret)
}

func anyMatch(str string, matches ...*string) bool {
	for _, m := range matches {
		if m == nil {
			continue
		}

		if *m == str {
			return true
		}
	}

	return false
}
