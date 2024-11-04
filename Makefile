GO_CLIENT_PATH="./vendor/github.com/serverscom/serverscom-go-client/pkg"

generate:
	go mod tidy
	go mod vendor
	mockgen --destination ./serverscom/testing/collection_mock.go --package=serverscom_testing --source ${GO_CLIENT_PATH}/collection.go
	mockgen --destination ./serverscom/testing/cloud_computing_instances_mock.go --package=serverscom_testing --source ${GO_CLIENT_PATH}/cloud_computing_instances.go
	mockgen --destination ./serverscom/testing/hosts_mock.go --package=serverscom_testing --source ${GO_CLIENT_PATH}/hosts.go
	mockgen --destination ./serverscom/testing/load_balancers_mock.go --package=serverscom_testing --source ${GO_CLIENT_PATH}/load_balancers.go
	sed -i '' 's|github.com/serverscom/cloud-controller-manager/vendor/||' ./serverscom/testing/collection_mock.go ./serverscom/testing/cloud_computing_instances_mock.go ./serverscom/testing/hosts_mock.go ./serverscom/testing/load_balancers_mock.go
