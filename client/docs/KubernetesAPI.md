# \KubernetesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesAddK8sAlert**](KubernetesAPI.md#KubernetesAddK8sAlert) | **Post** /api/v1/kubernetes/alert/{projectId} | Add k8s alert
[**KubernetesAddK8sEvents**](KubernetesAPI.md#KubernetesAddK8sEvents) | **Post** /api/v1/kubernetes/event/{projectId} | Add k8s event
[**KubernetesAlertList**](KubernetesAPI.md#KubernetesAlertList) | **Get** /api/v1/kubernetes/{projectId}/alerts | Retrieve a list of alerts for project
[**KubernetesCli**](KubernetesAPI.md#KubernetesCli) | **Post** /api/v1/kubernetes/cli | Execute k8s web socket namespaced pod
[**KubernetesConfigMapList**](KubernetesAPI.md#KubernetesConfigMapList) | **Get** /api/v1/kubernetes/{projectId}/configmap | Retrieve a list of k8s config map for all namespaces
[**KubernetesCrdList**](KubernetesAPI.md#KubernetesCrdList) | **Get** /api/v1/kubernetes/{projectId}/crd | Retrieve a list of crd
[**KubernetesCronJobList**](KubernetesAPI.md#KubernetesCronJobList) | **Get** /api/v1/kubernetes/{projectId}/cronjobs | Retrieve a list of k8s cron jobs for all namespaces
[**KubernetesDaemonSetList**](KubernetesAPI.md#KubernetesDaemonSetList) | **Get** /api/v1/kubernetes/{projectId}/daemonset | Retrieve list of k8s daemonset
[**KubernetesDashboardList**](KubernetesAPI.md#KubernetesDashboardList) | **Get** /api/v1/kubernetes/{projectId}/dashboard | Retrieve a list of crd
[**KubernetesDeploymentList**](KubernetesAPI.md#KubernetesDeploymentList) | **Get** /api/v1/kubernetes/{projectId}/deployment | Retrieve a list of k8s deployment for all namespaces
[**KubernetesDescribeConfigMap**](KubernetesAPI.md#KubernetesDescribeConfigMap) | **Post** /api/v1/kubernetes/describe/configmap | Describe configmap
[**KubernetesDescribeCrd**](KubernetesAPI.md#KubernetesDescribeCrd) | **Post** /api/v1/kubernetes/describe/crd | Describe crd
[**KubernetesDescribeCronjob**](KubernetesAPI.md#KubernetesDescribeCronjob) | **Post** /api/v1/kubernetes/describe/cronjob | Describe cronjob
[**KubernetesDescribeDaemonSet**](KubernetesAPI.md#KubernetesDescribeDaemonSet) | **Post** /api/v1/kubernetes/describe/daemonset | Describe daemonset
[**KubernetesDescribeDeployment**](KubernetesAPI.md#KubernetesDescribeDeployment) | **Post** /api/v1/kubernetes/describe/deployment | Describe deployment
[**KubernetesDescribeIngress**](KubernetesAPI.md#KubernetesDescribeIngress) | **Post** /api/v1/kubernetes/describe/ingress | Describe ingress
[**KubernetesDescribeJob**](KubernetesAPI.md#KubernetesDescribeJob) | **Post** /api/v1/kubernetes/describe/job | Describe job
[**KubernetesDescribeNetworkPolicy**](KubernetesAPI.md#KubernetesDescribeNetworkPolicy) | **Post** /api/v1/kubernetes/describe/network-policy | Describe network policy
[**KubernetesDescribeNode**](KubernetesAPI.md#KubernetesDescribeNode) | **Post** /api/v1/kubernetes/describe/node | Describe node
[**KubernetesDescribePdb**](KubernetesAPI.md#KubernetesDescribePdb) | **Post** /api/v1/kubernetes/describe/pdb | Describe pdb
[**KubernetesDescribePod**](KubernetesAPI.md#KubernetesDescribePod) | **Post** /api/v1/kubernetes/describe/pod | Describe pod
[**KubernetesDescribePvc**](KubernetesAPI.md#KubernetesDescribePvc) | **Post** /api/v1/kubernetes/describe/pvc | Describe pvc
[**KubernetesDescribeSecret**](KubernetesAPI.md#KubernetesDescribeSecret) | **Post** /api/v1/kubernetes/describe/secret | Describe secret
[**KubernetesDescribeService**](KubernetesAPI.md#KubernetesDescribeService) | **Post** /api/v1/kubernetes/describe/service | Describe service
[**KubernetesDescribeStorageClass**](KubernetesAPI.md#KubernetesDescribeStorageClass) | **Post** /api/v1/kubernetes/describe/storageclass | Describe storage class
[**KubernetesDescribeSts**](KubernetesAPI.md#KubernetesDescribeSts) | **Post** /api/v1/kubernetes/describe/sts | Describe stateful set
[**KubernetesDownload**](KubernetesAPI.md#KubernetesDownload) | **Get** /api/v1/kubernetes/{projectId}/download | Download kube config file
[**KubernetesExport**](KubernetesAPI.md#KubernetesExport) | **Get** /api/v1/kubernetes/export | Export
[**KubernetesGetSupportedList**](KubernetesAPI.md#KubernetesGetSupportedList) | **Get** /api/v1/kubernetes/supported/list | Retrieve Taikun supported kubernetes versions
[**KubernetesHelmReleaseList**](KubernetesAPI.md#KubernetesHelmReleaseList) | **Get** /api/v1/kubernetes/{projectId}/helmreleases | Retrieve a list of k8s helm releases for all namespaces
[**KubernetesIngressList**](KubernetesAPI.md#KubernetesIngressList) | **Get** /api/v1/kubernetes/{projectId}/ingress | Retrieve a list of k8s ingress for all namespaces
[**KubernetesInteractiveShell**](KubernetesAPI.md#KubernetesInteractiveShell) | **Post** /api/v1/kubernetes/interactive-shell | Produce interactive shell command
[**KubernetesJobsList**](KubernetesAPI.md#KubernetesJobsList) | **Get** /api/v1/kubernetes/{projectId}/jobs | Retrieve a list of k8s jobs for all namespaces
[**KubernetesKillPod**](KubernetesAPI.md#KubernetesKillPod) | **Post** /api/v1/kubernetes/{projectId}/deletepod/{metadataName}/{namespace} | 
[**KubernetesKubeConfig**](KubernetesAPI.md#KubernetesKubeConfig) | **Get** /api/v1/kubernetes/{projectId}/kubeconfig | Retrieve kube config file
[**KubernetesNamespaceList**](KubernetesAPI.md#KubernetesNamespaceList) | **Get** /api/v1/kubernetes/{projectId}/namespaces | Retrieve kube config file
[**KubernetesNetworkPolicyList**](KubernetesAPI.md#KubernetesNetworkPolicyList) | **Get** /api/v1/kubernetes/{projectId}/network-policies | Retrieve a list of k8s network-policies for all namespaces
[**KubernetesNodeList**](KubernetesAPI.md#KubernetesNodeList) | **Get** /api/v1/kubernetes/{projectId}/node | Retrieve a list of k8s node
[**KubernetesOverview**](KubernetesAPI.md#KubernetesOverview) | **Get** /api/v1/kubernetes/overview | Overview kubernetes nodes and pods by organization id
[**KubernetesPatchCrd**](KubernetesAPI.md#KubernetesPatchCrd) | **Post** /api/v1/kubernetes/patch/crd | Patch crd
[**KubernetesPatchCronJob**](KubernetesAPI.md#KubernetesPatchCronJob) | **Post** /api/v1/kubernetes/patch/cronjob | Patch cron-job
[**KubernetesPatchIngress**](KubernetesAPI.md#KubernetesPatchIngress) | **Post** /api/v1/kubernetes/patch/ingress | Patch ingress
[**KubernetesPatchJob**](KubernetesAPI.md#KubernetesPatchJob) | **Post** /api/v1/kubernetes/patch/job | Patch job
[**KubernetesPatchNode**](KubernetesAPI.md#KubernetesPatchNode) | **Post** /api/v1/kubernetes/patch/node | Patch node
[**KubernetesPatchPdb**](KubernetesAPI.md#KubernetesPatchPdb) | **Post** /api/v1/kubernetes/patch/pdb | Patch pdb
[**KubernetesPatchPod**](KubernetesAPI.md#KubernetesPatchPod) | **Post** /api/v1/kubernetes/patch/pod | Patch pod
[**KubernetesPatchPvc**](KubernetesAPI.md#KubernetesPatchPvc) | **Post** /api/v1/kubernetes/patch/pvc | Patch pvc
[**KubernetesPatchSecret**](KubernetesAPI.md#KubernetesPatchSecret) | **Post** /api/v1/kubernetes/patch/secret | Patch secret
[**KubernetesPatchSts**](KubernetesAPI.md#KubernetesPatchSts) | **Post** /api/v1/kubernetes/patch/sts | Patch sts
[**KubernetesPdbList**](KubernetesAPI.md#KubernetesPdbList) | **Get** /api/v1/kubernetes/{projectId}/pdb | Retrieve a list of k8s pdb for all namespaces
[**KubernetesPodList**](KubernetesAPI.md#KubernetesPodList) | **Get** /api/v1/kubernetes/{projectId}/pod | Retrieve a list of k8s pod for all namespaces
[**KubernetesPodLogs**](KubernetesAPI.md#KubernetesPodLogs) | **Post** /api/v1/kubernetes/podLogs | Retrieve k8s pod logs
[**KubernetesPvcList**](KubernetesAPI.md#KubernetesPvcList) | **Get** /api/v1/kubernetes/{projectId}/pvc | Retrieve a list of k8s pvc for all namespaces
[**KubernetesQuota**](KubernetesAPI.md#KubernetesQuota) | **Get** /api/v1/kubernetes/{projectId}/quota | K8s quota usage
[**KubernetesRemovealerts**](KubernetesAPI.md#KubernetesRemovealerts) | **Post** /api/v1/kubernetes/removealerts | Remove k8s alerts
[**KubernetesRestartDaemonSet**](KubernetesAPI.md#KubernetesRestartDaemonSet) | **Post** /api/v1/kubernetes/restart/daemonset | Restart daemon set
[**KubernetesRestartDeployment**](KubernetesAPI.md#KubernetesRestartDeployment) | **Post** /api/v1/kubernetes/restart/deployment | Restart deployment
[**KubernetesRestartSts**](KubernetesAPI.md#KubernetesRestartSts) | **Post** /api/v1/kubernetes/restart/sts | Restart stateful set
[**KubernetesSecretList**](KubernetesAPI.md#KubernetesSecretList) | **Get** /api/v1/kubernetes/{projectId}/secret | Retrieve a list of k8s secret for all namespaces
[**KubernetesServiceList**](KubernetesAPI.md#KubernetesServiceList) | **Get** /api/v1/kubernetes/{projectId}/service | Retrieve a list of k8s service for all namespaces
[**KubernetesSilenceManager**](KubernetesAPI.md#KubernetesSilenceManager) | **Post** /api/v1/kubernetes/silencemanager | Silence management for k8s alerts
[**KubernetesStorageClassList**](KubernetesAPI.md#KubernetesStorageClassList) | **Get** /api/v1/kubernetes/{projectId}/storageclass | Retrieve a list of k8s storageclass for all namespaces
[**KubernetesStsList**](KubernetesAPI.md#KubernetesStsList) | **Get** /api/v1/kubernetes/{projectId}/sts | Retrieve a list of k8s sts for all namespaces
[**KubernetesUpdateAlert**](KubernetesAPI.md#KubernetesUpdateAlert) | **Put** /api/v1/kubernetes/updatealert/{alertId} | Update k8s alert



## KubernetesAddK8sAlert

> KubernetesAddK8sAlert(ctx, projectId).CreateAlertDto(createAlertDto).Execute()

Add k8s alert

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    createAlertDto := *openapiclient.NewCreateAlertDto() // CreateAlertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesAddK8sAlert(context.Background(), projectId).CreateAlertDto(createAlertDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesAddK8sAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesAddK8sAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAlertDto** | [**CreateAlertDto**](CreateAlertDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesAddK8sEvents

> KubernetesAddK8sEvents(ctx, projectId).KubernetesEventCreateDto(kubernetesEventCreateDto).Execute()

Add k8s event

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    kubernetesEventCreateDto := *openapiclient.NewKubernetesEventCreateDto() // KubernetesEventCreateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesAddK8sEvents(context.Background(), projectId).KubernetesEventCreateDto(kubernetesEventCreateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesAddK8sEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesAddK8sEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesEventCreateDto** | [**KubernetesEventCreateDto**](KubernetesEventCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesAlertList

> KubernetesAlertList KubernetesAlertList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()

Retrieve a list of alerts for project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesAlertList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesAlertList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesAlertList`: KubernetesAlertList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesAlertList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesAlertListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **type_** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**KubernetesAlertList**](KubernetesAlertList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCli

> string KubernetesCli(ctx).KubernetesCliCommand(kubernetesCliCommand).Execute()

Execute k8s web socket namespaced pod

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    kubernetesCliCommand := *openapiclient.NewKubernetesCliCommand() // KubernetesCliCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesCli(context.Background()).KubernetesCliCommand(kubernetesCliCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesCli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesCli`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesCli`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCliCommand** | [**KubernetesCliCommand**](KubernetesCliCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesConfigMapList

> ConfigMaps KubernetesConfigMapList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s config map for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesConfigMapList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesConfigMapList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesConfigMapList`: ConfigMaps
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesConfigMapList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesConfigMapListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**ConfigMaps**](ConfigMaps.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCrdList

> interface{} KubernetesCrdList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of crd

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesCrdList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesCrdList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesCrdList`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesCrdList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCrdListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCronJobList

> KubernetesCronJobsList KubernetesCronJobList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s cron jobs for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesCronJobList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesCronJobList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesCronJobList`: KubernetesCronJobsList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesCronJobList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCronJobListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**KubernetesCronJobsList**](KubernetesCronJobsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDaemonSetList

> DaemonSets KubernetesDaemonSetList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve list of k8s daemonset

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDaemonSetList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDaemonSetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDaemonSetList`: DaemonSets
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDaemonSetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDaemonSetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**DaemonSets**](DaemonSets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDashboardList

> KubernetesDashboardDto KubernetesDashboardList(ctx, projectId).Execute()

Retrieve a list of crd

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDashboardList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDashboardList`: KubernetesDashboardDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDashboardList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesDashboardDto**](KubernetesDashboardDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDeploymentList

> Deployments KubernetesDeploymentList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s deployment for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDeploymentList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDeploymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDeploymentList`: Deployments
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDeploymentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDeploymentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Deployments**](Deployments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeConfigMap

> string KubernetesDescribeConfigMap(ctx).DescribeConfigMapCommand(describeConfigMapCommand).Execute()

Describe configmap

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeConfigMapCommand := *openapiclient.NewDescribeConfigMapCommand() // DescribeConfigMapCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeConfigMap(context.Background()).DescribeConfigMapCommand(describeConfigMapCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeConfigMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeConfigMap`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeConfigMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeConfigMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeConfigMapCommand** | [**DescribeConfigMapCommand**](DescribeConfigMapCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeCrd

> string KubernetesDescribeCrd(ctx).DescribeCrdCommand(describeCrdCommand).Execute()

Describe crd

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeCrdCommand := *openapiclient.NewDescribeCrdCommand() // DescribeCrdCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeCrd(context.Background()).DescribeCrdCommand(describeCrdCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeCrd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeCrd`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeCrd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeCrdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeCrdCommand** | [**DescribeCrdCommand**](DescribeCrdCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeCronjob

> string KubernetesDescribeCronjob(ctx).DescribeCronJobCommand(describeCronJobCommand).Execute()

Describe cronjob

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeCronJobCommand := *openapiclient.NewDescribeCronJobCommand() // DescribeCronJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeCronjob(context.Background()).DescribeCronJobCommand(describeCronJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeCronjob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeCronjob`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeCronjob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeCronjobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeCronJobCommand** | [**DescribeCronJobCommand**](DescribeCronJobCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeDaemonSet

> string KubernetesDescribeDaemonSet(ctx).DescribeDaemonSetCommand(describeDaemonSetCommand).Execute()

Describe daemonset

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeDaemonSetCommand := *openapiclient.NewDescribeDaemonSetCommand() // DescribeDaemonSetCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeDaemonSet(context.Background()).DescribeDaemonSetCommand(describeDaemonSetCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeDaemonSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeDaemonSet`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeDaemonSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeDaemonSetCommand** | [**DescribeDaemonSetCommand**](DescribeDaemonSetCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeDeployment

> string KubernetesDescribeDeployment(ctx).DescribeDeploymentCommand(describeDeploymentCommand).Execute()

Describe deployment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeDeploymentCommand := *openapiclient.NewDescribeDeploymentCommand() // DescribeDeploymentCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeDeployment(context.Background()).DescribeDeploymentCommand(describeDeploymentCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeDeployment`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeDeploymentCommand** | [**DescribeDeploymentCommand**](DescribeDeploymentCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeIngress

> string KubernetesDescribeIngress(ctx).DescribeIngressCommand(describeIngressCommand).Execute()

Describe ingress

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeIngressCommand := *openapiclient.NewDescribeIngressCommand() // DescribeIngressCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeIngress(context.Background()).DescribeIngressCommand(describeIngressCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeIngress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeIngress`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeIngress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeIngressCommand** | [**DescribeIngressCommand**](DescribeIngressCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeJob

> string KubernetesDescribeJob(ctx).DescribeJobCommand(describeJobCommand).Execute()

Describe job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeJobCommand := *openapiclient.NewDescribeJobCommand() // DescribeJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeJob(context.Background()).DescribeJobCommand(describeJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeJob`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeJobCommand** | [**DescribeJobCommand**](DescribeJobCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeNetworkPolicy

> string KubernetesDescribeNetworkPolicy(ctx).DescribeNetworkPolicyCommand(describeNetworkPolicyCommand).Execute()

Describe network policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeNetworkPolicyCommand := *openapiclient.NewDescribeNetworkPolicyCommand() // DescribeNetworkPolicyCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeNetworkPolicy(context.Background()).DescribeNetworkPolicyCommand(describeNetworkPolicyCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeNetworkPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeNetworkPolicy`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeNetworkPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeNetworkPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeNetworkPolicyCommand** | [**DescribeNetworkPolicyCommand**](DescribeNetworkPolicyCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeNode

> string KubernetesDescribeNode(ctx).DescribeNodeCommand(describeNodeCommand).Execute()

Describe node

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeNodeCommand := *openapiclient.NewDescribeNodeCommand() // DescribeNodeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeNode(context.Background()).DescribeNodeCommand(describeNodeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeNode`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeNodeCommand** | [**DescribeNodeCommand**](DescribeNodeCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePdb

> string KubernetesDescribePdb(ctx).DescribePodDisruptionCommand(describePodDisruptionCommand).Execute()

Describe pdb

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describePodDisruptionCommand := *openapiclient.NewDescribePodDisruptionCommand() // DescribePodDisruptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribePdb(context.Background()).DescribePodDisruptionCommand(describePodDisruptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribePdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePdb`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribePdb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describePodDisruptionCommand** | [**DescribePodDisruptionCommand**](DescribePodDisruptionCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePod

> string KubernetesDescribePod(ctx).DescribePodCommand(describePodCommand).Execute()

Describe pod

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describePodCommand := *openapiclient.NewDescribePodCommand() // DescribePodCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribePod(context.Background()).DescribePodCommand(describePodCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribePod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePod`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribePod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describePodCommand** | [**DescribePodCommand**](DescribePodCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribePvc

> string KubernetesDescribePvc(ctx).DescribePvcCommand(describePvcCommand).Execute()

Describe pvc

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describePvcCommand := *openapiclient.NewDescribePvcCommand() // DescribePvcCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribePvc(context.Background()).DescribePvcCommand(describePvcCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribePvc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribePvc`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribePvc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribePvcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describePvcCommand** | [**DescribePvcCommand**](DescribePvcCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeSecret

> string KubernetesDescribeSecret(ctx).DescribeSecretCommand(describeSecretCommand).Execute()

Describe secret

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeSecretCommand := *openapiclient.NewDescribeSecretCommand() // DescribeSecretCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeSecret(context.Background()).DescribeSecretCommand(describeSecretCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeSecret`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeSecretCommand** | [**DescribeSecretCommand**](DescribeSecretCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeService

> string KubernetesDescribeService(ctx).DescribeServiceCommand(describeServiceCommand).Execute()

Describe service

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeServiceCommand := *openapiclient.NewDescribeServiceCommand() // DescribeServiceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeService(context.Background()).DescribeServiceCommand(describeServiceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeService`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeServiceCommand** | [**DescribeServiceCommand**](DescribeServiceCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeStorageClass

> string KubernetesDescribeStorageClass(ctx).DescribeStorageClassCommand(describeStorageClassCommand).Execute()

Describe storage class

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeStorageClassCommand := *openapiclient.NewDescribeStorageClassCommand() // DescribeStorageClassCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeStorageClass(context.Background()).DescribeStorageClassCommand(describeStorageClassCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeStorageClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeStorageClass`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeStorageClass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeStorageClassCommand** | [**DescribeStorageClassCommand**](DescribeStorageClassCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeSts

> string KubernetesDescribeSts(ctx).DescribeStsCommand(describeStsCommand).Execute()

Describe stateful set

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    describeStsCommand := *openapiclient.NewDescribeStsCommand() // DescribeStsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeSts(context.Background()).DescribeStsCommand(describeStsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDescribeSts`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeSts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeStsCommand** | [**DescribeStsCommand**](DescribeStsCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDownload

> interface{} KubernetesDownload(ctx, projectId).Execute()

Download kube config file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesDownload(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesDownload`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesExport

> CsvExporter KubernetesExport(ctx).ProjectId(projectId).Execute()

Export

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesExport(context.Background()).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesExport`: CsvExporter
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 

### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetSupportedList

> []KubernetesVersionListDto KubernetesGetSupportedList(ctx).Execute()

Retrieve Taikun supported kubernetes versions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesGetSupportedList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesGetSupportedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesGetSupportedList`: []KubernetesVersionListDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesGetSupportedList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetSupportedListRequest struct via the builder pattern


### Return type

[**[]KubernetesVersionListDto**](KubernetesVersionListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesHelmReleaseList

> HelmReleasesList KubernetesHelmReleaseList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s helm releases for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesHelmReleaseList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesHelmReleaseList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesHelmReleaseList`: HelmReleasesList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesHelmReleaseList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesHelmReleaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**HelmReleasesList**](HelmReleasesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesIngressList

> Ingresses KubernetesIngressList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s ingress for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesIngressList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesIngressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesIngressList`: Ingresses
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesIngressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesIngressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Ingresses**](Ingresses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesInteractiveShell

> string KubernetesInteractiveShell(ctx).InteractiveShellSendCommand(interactiveShellSendCommand).Execute()

Produce interactive shell command

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    interactiveShellSendCommand := *openapiclient.NewInteractiveShellSendCommand() // InteractiveShellSendCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesInteractiveShell(context.Background()).InteractiveShellSendCommand(interactiveShellSendCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesInteractiveShell``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesInteractiveShell`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesInteractiveShell`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesInteractiveShellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactiveShellSendCommand** | [**InteractiveShellSendCommand**](InteractiveShellSendCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesJobsList

> KubernetesJobList KubernetesJobsList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s jobs for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesJobsList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesJobsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesJobsList`: KubernetesJobList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesJobsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesJobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**KubernetesJobList**](KubernetesJobList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesKillPod

> KubernetesKillPod(ctx, projectId, metadataName, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    metadataName := "metadataName_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesKillPod(context.Background(), projectId, metadataName, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesKillPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**metadataName** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesKillPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesKubeConfig

> KubeConfigResponse KubernetesKubeConfig(ctx, projectId).Execute()

Retrieve kube config file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesKubeConfig(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesKubeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesKubeConfig`: KubeConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesKubeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesKubeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubeConfigResponse**](KubeConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNamespaceList

> []string KubernetesNamespaceList(ctx, projectId).Execute()

Retrieve kube config file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesNamespaceList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesNamespaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesNamespaceList`: []string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesNamespaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNamespaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNetworkPolicyList

> NetworkPolicies KubernetesNetworkPolicyList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s network-policies for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesNetworkPolicyList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesNetworkPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesNetworkPolicyList`: NetworkPolicies
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesNetworkPolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNetworkPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**NetworkPolicies**](NetworkPolicies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNodeList

> interface{} KubernetesNodeList(ctx, projectId).SearchId(searchId).Execute()

Retrieve a list of k8s node

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    searchId := "searchId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesNodeList(context.Background(), projectId).SearchId(searchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesNodeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesNodeList`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesNodeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNodeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchId** | **string** |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesOverview

> []KubernetesOverviewDto KubernetesOverview(ctx).OrganizationId(organizationId).Execute()

Overview kubernetes nodes and pods by organization id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesOverview(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesOverview`: []KubernetesOverviewDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**[]KubernetesOverviewDto**](KubernetesOverviewDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchCrd

> KubernetesPatchCrd(ctx).PatchCrdCommand(patchCrdCommand).Execute()

Patch crd

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchCrdCommand := *openapiclient.NewPatchCrdCommand() // PatchCrdCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchCrd(context.Background()).PatchCrdCommand(patchCrdCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchCrd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchCrdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchCrdCommand** | [**PatchCrdCommand**](PatchCrdCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchCronJob

> KubernetesPatchCronJob(ctx).PatchCronJobCommand(patchCronJobCommand).Execute()

Patch cron-job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchCronJobCommand := *openapiclient.NewPatchCronJobCommand() // PatchCronJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchCronJob(context.Background()).PatchCronJobCommand(patchCronJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchCronJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchCronJobCommand** | [**PatchCronJobCommand**](PatchCronJobCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchIngress

> KubernetesPatchIngress(ctx).PatchIngressCommand(patchIngressCommand).Execute()

Patch ingress

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchIngressCommand := *openapiclient.NewPatchIngressCommand() // PatchIngressCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchIngress(context.Background()).PatchIngressCommand(patchIngressCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchIngress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchIngressCommand** | [**PatchIngressCommand**](PatchIngressCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchJob

> KubernetesPatchJob(ctx).PatchJobCommand(patchJobCommand).Execute()

Patch job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchJobCommand := *openapiclient.NewPatchJobCommand() // PatchJobCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchJob(context.Background()).PatchJobCommand(patchJobCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchJobCommand** | [**PatchJobCommand**](PatchJobCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchNode

> KubernetesPatchNode(ctx).PatchNodeCommand(patchNodeCommand).Execute()

Patch node

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchNodeCommand := *openapiclient.NewPatchNodeCommand() // PatchNodeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchNode(context.Background()).PatchNodeCommand(patchNodeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchNodeCommand** | [**PatchNodeCommand**](PatchNodeCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchPdb

> KubernetesPatchPdb(ctx).PatchPdbCommand(patchPdbCommand).Execute()

Patch pdb

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchPdbCommand := *openapiclient.NewPatchPdbCommand() // PatchPdbCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchPdb(context.Background()).PatchPdbCommand(patchPdbCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchPdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPdbCommand** | [**PatchPdbCommand**](PatchPdbCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchPod

> KubernetesPatchPod(ctx).PatchPodCommand(patchPodCommand).Execute()

Patch pod

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchPodCommand := *openapiclient.NewPatchPodCommand() // PatchPodCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchPod(context.Background()).PatchPodCommand(patchPodCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPodCommand** | [**PatchPodCommand**](PatchPodCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchPvc

> KubernetesPatchPvc(ctx).PatchPvcCommand(patchPvcCommand).Execute()

Patch pvc

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchPvcCommand := *openapiclient.NewPatchPvcCommand() // PatchPvcCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchPvc(context.Background()).PatchPvcCommand(patchPvcCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchPvc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchPvcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchPvcCommand** | [**PatchPvcCommand**](PatchPvcCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchSecret

> KubernetesPatchSecret(ctx).PatchSecretCommand(patchSecretCommand).Execute()

Patch secret

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchSecretCommand := *openapiclient.NewPatchSecretCommand() // PatchSecretCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchSecret(context.Background()).PatchSecretCommand(patchSecretCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchSecretCommand** | [**PatchSecretCommand**](PatchSecretCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchSts

> KubernetesPatchSts(ctx).PatchStsCommand(patchStsCommand).Execute()

Patch sts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    patchStsCommand := *openapiclient.NewPatchStsCommand() // PatchStsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesPatchSts(context.Background()).PatchStsCommand(patchStsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchStsCommand** | [**PatchStsCommand**](PatchStsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPdbList

> PodDisruptions KubernetesPdbList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pdb for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesPdbList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPdbList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPdbList`: PodDisruptions
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesPdbList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPdbListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**PodDisruptions**](PodDisruptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPodList

> Pods KubernetesPodList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pod for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesPodList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPodList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPodList`: Pods
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesPodList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPodListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Pods**](Pods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPodLogs

> interface{} KubernetesPodLogs(ctx).KubernetesPodLogsCommand(kubernetesPodLogsCommand).Execute()

Retrieve k8s pod logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    kubernetesPodLogsCommand := *openapiclient.NewKubernetesPodLogsCommand() // KubernetesPodLogsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesPodLogs(context.Background()).KubernetesPodLogsCommand(kubernetesPodLogsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPodLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPodLogs`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesPodLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPodLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesPodLogsCommand** | [**KubernetesPodLogsCommand**](KubernetesPodLogsCommand.md) |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPvcList

> Pvcs KubernetesPvcList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s pvc for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesPvcList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPvcList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesPvcList`: Pvcs
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesPvcList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPvcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Pvcs**](Pvcs.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesQuota

> KubernetesQuotaListDto KubernetesQuota(ctx, projectId).Execute()

K8s quota usage

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesQuota(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesQuota`: KubernetesQuotaListDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesQuotaListDto**](KubernetesQuotaListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRemovealerts

> KubernetesRemovealerts(ctx).DeleteAlertCommand(deleteAlertCommand).Execute()

Remove k8s alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    deleteAlertCommand := *openapiclient.NewDeleteAlertCommand() // DeleteAlertCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesRemovealerts(context.Background()).DeleteAlertCommand(deleteAlertCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesRemovealerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRemovealertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAlertCommand** | [**DeleteAlertCommand**](DeleteAlertCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRestartDaemonSet

> KubernetesRestartDaemonSet(ctx).RestartDaemonSetCommand(restartDaemonSetCommand).Execute()

Restart daemon set

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    restartDaemonSetCommand := *openapiclient.NewRestartDaemonSetCommand() // RestartDaemonSetCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesRestartDaemonSet(context.Background()).RestartDaemonSetCommand(restartDaemonSetCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesRestartDaemonSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restartDaemonSetCommand** | [**RestartDaemonSetCommand**](RestartDaemonSetCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRestartDeployment

> KubernetesRestartDeployment(ctx).RestartDeploymentCommand(restartDeploymentCommand).Execute()

Restart deployment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    restartDeploymentCommand := *openapiclient.NewRestartDeploymentCommand() // RestartDeploymentCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesRestartDeployment(context.Background()).RestartDeploymentCommand(restartDeploymentCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesRestartDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restartDeploymentCommand** | [**RestartDeploymentCommand**](RestartDeploymentCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRestartSts

> KubernetesRestartSts(ctx).RestartStsCommand(restartStsCommand).Execute()

Restart stateful set

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    restartStsCommand := *openapiclient.NewRestartStsCommand() // RestartStsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesRestartSts(context.Background()).RestartStsCommand(restartStsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesRestartSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRestartStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restartStsCommand** | [**RestartStsCommand**](RestartStsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesSecretList

> Secrets KubernetesSecretList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s secret for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesSecretList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesSecretList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesSecretList`: Secrets
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesSecretList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesSecretListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Secrets**](Secrets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesServiceList

> Services KubernetesServiceList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s service for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesServiceList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesServiceList`: Services
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**Services**](Services.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesSilenceManager

> KubernetesSilenceManager(ctx).SilenceOperationsCommand(silenceOperationsCommand).Execute()

Silence management for k8s alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    silenceOperationsCommand := *openapiclient.NewSilenceOperationsCommand() // SilenceOperationsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesSilenceManager(context.Background()).SilenceOperationsCommand(silenceOperationsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesSilenceManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesSilenceManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **silenceOperationsCommand** | [**SilenceOperationsCommand**](SilenceOperationsCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesStorageClassList

> StorageClasses KubernetesStorageClassList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s storageclass for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesStorageClassList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesStorageClassList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesStorageClassList`: StorageClasses
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesStorageClassList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesStorageClassListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**StorageClasses**](StorageClasses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesStsList

> StsList KubernetesStsList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve a list of k8s sts for all namespaces

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.KubernetesStsList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesStsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesStsList`: StsList
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesStsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesStsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**StsList**](StsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesUpdateAlert

> KubernetesUpdateAlert(ctx, alertId).UpdateKubernetesAlertDto(updateKubernetesAlertDto).Execute()

Update k8s alert

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    alertId := int32(56) // int32 | 
    updateKubernetesAlertDto := *openapiclient.NewUpdateKubernetesAlertDto() // UpdateKubernetesAlertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.KubernetesUpdateAlert(context.Background(), alertId).UpdateKubernetesAlertDto(updateKubernetesAlertDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesUpdateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKubernetesAlertDto** | [**UpdateKubernetesAlertDto**](UpdateKubernetesAlertDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

