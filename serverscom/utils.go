package serverscom

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"k8s.io/klog/v2"

	"maps"

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

// sanitizeLabelValue sanitazes label value according to:
// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
//
// rules:
// must be 63 characters or less (can be empty),
// unless empty, must begin and end with an alphanumeric character ([a-z0-9A-Z]),
// could contain dashes (-), underscores (_), dots (.), and alphanumerics between.
//
// returns empty string if no valid chars found
func sanitizeLabelValue(value string) string {
	if len(value) == 0 {
		return value
	}

	runes := []rune(value)

	// replace any invalid char to '-'
	for i, r := range runes {
		if !isValidLabelChar(r) {
			runes[i] = '-'
		}
	}

	start := 0
	for start < len(runes) && !isAlphaNumeric(runes[start]) {
		start++
	}

	if start == len(runes) {
		return ""
	}

	end := len(runes) - 1
	for end >= 0 && !isAlphaNumeric(runes[end]) {
		end--
	}

	runes = runes[start : end+1]

	if len(runes) > 63 {
		runes = runes[:63]

		// check that after truncate we still have valid chars in the end
		if !isAlphaNumeric(runes[len(runes)-1]) {
			lastValid := len(runes) - 1
			for lastValid >= 0 && !isAlphaNumeric(runes[lastValid]) {
				lastValid--
			}

			if lastValid >= 0 {
				runes = runes[:lastValid+1]
			} else {
				return ""
			}
		}
	}

	return string(runes)
}

func isValidLabelChar(r rune) bool {
	return isAlphaNumeric(r) || r == '-' || r == '_' || r == '.'
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

// mergeDefaultLabels merge existing labels with default ones.
// To ensure that we can add default labels, existing labels will sorted by keys and truncated to 64 - count of default labels.
// 64 - max supported labels for resource.
func mergeDefaultLabels(existing, defaultLabels map[string]string) map[string]string {
	if defaultLabels == nil {
		return existing
	}
	if existing == nil {
		return defaultLabels
	}
	// max labels - count of default labels
	truncateTo := 64 - len(defaultLabels)

	for k := range defaultLabels {
		delete(existing, k)
	}

	if len(existing) > truncateTo {
		keys := make([]string, 0, len(existing))
		for k := range existing {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		truncated := make(map[string]string, 64)
		for i := range truncateTo {
			k := keys[i]
			truncated[k] = existing[k]
		}
		existing = truncated
	}

	maps.Copy(existing, defaultLabels)

	return existing
}
