# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAccessProfiles**](SearchAPI.md#SearchAccessProfiles) | **Post** /api/v1/search/access-profiles | Global search for access-profiles
[**SearchBackupCredentials**](SearchAPI.md#SearchBackupCredentials) | **Post** /api/v1/search/backup-credentials | Global search for backup-credentials
[**SearchBillingCredentials**](SearchAPI.md#SearchBillingCredentials) | **Post** /api/v1/search/billing-credentials | Global search for billing-credentials
[**SearchCloudCredentials**](SearchAPI.md#SearchCloudCredentials) | **Post** /api/v1/search/cloud-credentials | Global search for cloud-credentials
[**SearchConfigMaps**](SearchAPI.md#SearchConfigMaps) | **Post** /api/v1/search/config-maps | Global search for config-maps
[**SearchDaemonSets**](SearchAPI.md#SearchDaemonSets) | **Post** /api/v1/search/daemon-sets | Global search for daemon-sets
[**SearchDeployments**](SearchAPI.md#SearchDeployments) | **Post** /api/v1/search/deployments | Global search for deployments
[**SearchIngress**](SearchAPI.md#SearchIngress) | **Post** /api/v1/search/ingress | Global search for ingress
[**SearchKubernetesProfiles**](SearchAPI.md#SearchKubernetesProfiles) | **Post** /api/v1/search/kubernetes-profiles | Global search for kubernetes-profiles
[**SearchNodes**](SearchAPI.md#SearchNodes) | **Post** /api/v1/search/nodes | Global search for nodes
[**SearchOrganizations**](SearchAPI.md#SearchOrganizations) | **Post** /api/v1/search/organizations | Global search for organizations
[**SearchPartners**](SearchAPI.md#SearchPartners) | **Post** /api/v1/search/partners | Global search for partners
[**SearchPods**](SearchAPI.md#SearchPods) | **Post** /api/v1/search/pods | Global search for pods
[**SearchProjects**](SearchAPI.md#SearchProjects) | **Post** /api/v1/search/projects | Global search for projects
[**SearchPrometheusRules**](SearchAPI.md#SearchPrometheusRules) | **Post** /api/v1/search/prometheus-rules | Global search for prometheus-rules
[**SearchPvcs**](SearchAPI.md#SearchPvcs) | **Post** /api/v1/search/pvcs | Global search for pvcs
[**SearchSecrets**](SearchAPI.md#SearchSecrets) | **Post** /api/v1/search/secrets | Global search for secrets
[**SearchServers**](SearchAPI.md#SearchServers) | **Post** /api/v1/search/servers | Global search for servers
[**SearchServices**](SearchAPI.md#SearchServices) | **Post** /api/v1/search/services | Global search for services
[**SearchStandAloneProfiles**](SearchAPI.md#SearchStandAloneProfiles) | **Post** /api/v1/search/stand-alone-profiles | Global search for stand-alone-profiles
[**SearchSts**](SearchAPI.md#SearchSts) | **Post** /api/v1/search/sts | Global search for sts
[**SearchUsers**](SearchAPI.md#SearchUsers) | **Post** /api/v1/search/users | Global search for users



## SearchAccessProfiles

> AccessProfilesSearchList SearchAccessProfiles(ctx).AccessProfilesSearchCommand(accessProfilesSearchCommand).Execute()

Global search for access-profiles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    accessProfilesSearchCommand := *openapiclient.NewAccessProfilesSearchCommand() // AccessProfilesSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchAccessProfiles(context.Background()).AccessProfilesSearchCommand(accessProfilesSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchAccessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAccessProfiles`: AccessProfilesSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchAccessProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessProfilesSearchCommand** | [**AccessProfilesSearchCommand**](AccessProfilesSearchCommand.md) |  | 

### Return type

[**AccessProfilesSearchList**](AccessProfilesSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBackupCredentials

> BackupCredentialsSearchList SearchBackupCredentials(ctx).BackupCredentialsSearchCommand(backupCredentialsSearchCommand).Execute()

Global search for backup-credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    backupCredentialsSearchCommand := *openapiclient.NewBackupCredentialsSearchCommand() // BackupCredentialsSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchBackupCredentials(context.Background()).BackupCredentialsSearchCommand(backupCredentialsSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchBackupCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchBackupCredentials`: BackupCredentialsSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchBackupCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBackupCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupCredentialsSearchCommand** | [**BackupCredentialsSearchCommand**](BackupCredentialsSearchCommand.md) |  | 

### Return type

[**BackupCredentialsSearchList**](BackupCredentialsSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBillingCredentials

> BillingCredentialsSearchList SearchBillingCredentials(ctx).BillingCredentialsSearchCommand(billingCredentialsSearchCommand).Execute()

Global search for billing-credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    billingCredentialsSearchCommand := *openapiclient.NewBillingCredentialsSearchCommand() // BillingCredentialsSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchBillingCredentials(context.Background()).BillingCredentialsSearchCommand(billingCredentialsSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchBillingCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchBillingCredentials`: BillingCredentialsSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchBillingCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBillingCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingCredentialsSearchCommand** | [**BillingCredentialsSearchCommand**](BillingCredentialsSearchCommand.md) |  | 

### Return type

[**BillingCredentialsSearchList**](BillingCredentialsSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudCredentials

> CloudCredentialsSearchList SearchCloudCredentials(ctx).CloudCredentialsSearchCommand(cloudCredentialsSearchCommand).Execute()

Global search for cloud-credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    cloudCredentialsSearchCommand := *openapiclient.NewCloudCredentialsSearchCommand() // CloudCredentialsSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchCloudCredentials(context.Background()).CloudCredentialsSearchCommand(cloudCredentialsSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCloudCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudCredentials`: CloudCredentialsSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCloudCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCredentialsSearchCommand** | [**CloudCredentialsSearchCommand**](CloudCredentialsSearchCommand.md) |  | 

### Return type

[**CloudCredentialsSearchList**](CloudCredentialsSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchConfigMaps

> ConfigMapSearchList SearchConfigMaps(ctx).ConfigMapSearchCommand(configMapSearchCommand).Execute()

Global search for config-maps

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    configMapSearchCommand := *openapiclient.NewConfigMapSearchCommand() // ConfigMapSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchConfigMaps(context.Background()).ConfigMapSearchCommand(configMapSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchConfigMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchConfigMaps`: ConfigMapSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchConfigMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchConfigMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configMapSearchCommand** | [**ConfigMapSearchCommand**](ConfigMapSearchCommand.md) |  | 

### Return type

[**ConfigMapSearchList**](ConfigMapSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDaemonSets

> DaemonSetSearchList SearchDaemonSets(ctx).DaemonSetSearchCommand(daemonSetSearchCommand).Execute()

Global search for daemon-sets

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    daemonSetSearchCommand := *openapiclient.NewDaemonSetSearchCommand() // DaemonSetSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchDaemonSets(context.Background()).DaemonSetSearchCommand(daemonSetSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchDaemonSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDaemonSets`: DaemonSetSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchDaemonSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDaemonSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **daemonSetSearchCommand** | [**DaemonSetSearchCommand**](DaemonSetSearchCommand.md) |  | 

### Return type

[**DaemonSetSearchList**](DaemonSetSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeployments

> DeploymentSearchList SearchDeployments(ctx).DeploymentSearchCommand(deploymentSearchCommand).Execute()

Global search for deployments

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    deploymentSearchCommand := *openapiclient.NewDeploymentSearchCommand() // DeploymentSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchDeployments(context.Background()).DeploymentSearchCommand(deploymentSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDeployments`: DeploymentSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentSearchCommand** | [**DeploymentSearchCommand**](DeploymentSearchCommand.md) |  | 

### Return type

[**DeploymentSearchList**](DeploymentSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIngress

> IngressSearchList SearchIngress(ctx).IngressSearchCommand(ingressSearchCommand).Execute()

Global search for ingress

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    ingressSearchCommand := *openapiclient.NewIngressSearchCommand() // IngressSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchIngress(context.Background()).IngressSearchCommand(ingressSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchIngress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIngress`: IngressSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchIngress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingressSearchCommand** | [**IngressSearchCommand**](IngressSearchCommand.md) |  | 

### Return type

[**IngressSearchList**](IngressSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchKubernetesProfiles

> KubernetesProfilesSearchList SearchKubernetesProfiles(ctx).KubernetesProfilesSearchCommand(kubernetesProfilesSearchCommand).Execute()

Global search for kubernetes-profiles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    kubernetesProfilesSearchCommand := *openapiclient.NewKubernetesProfilesSearchCommand() // KubernetesProfilesSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchKubernetesProfiles(context.Background()).KubernetesProfilesSearchCommand(kubernetesProfilesSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchKubernetesProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchKubernetesProfiles`: KubernetesProfilesSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchKubernetesProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchKubernetesProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesProfilesSearchCommand** | [**KubernetesProfilesSearchCommand**](KubernetesProfilesSearchCommand.md) |  | 

### Return type

[**KubernetesProfilesSearchList**](KubernetesProfilesSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchNodes

> NodesSearchList SearchNodes(ctx).NodesSearchCommand(nodesSearchCommand).Execute()

Global search for nodes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    nodesSearchCommand := *openapiclient.NewNodesSearchCommand() // NodesSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchNodes(context.Background()).NodesSearchCommand(nodesSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchNodes`: NodesSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodesSearchCommand** | [**NodesSearchCommand**](NodesSearchCommand.md) |  | 

### Return type

[**NodesSearchList**](NodesSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrganizations

> OrganizationSearchList SearchOrganizations(ctx).OrganizationSearchCommand(organizationSearchCommand).Execute()

Global search for organizations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    organizationSearchCommand := *openapiclient.NewOrganizationSearchCommand() // OrganizationSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchOrganizations(context.Background()).OrganizationSearchCommand(organizationSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchOrganizations`: OrganizationSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationSearchCommand** | [**OrganizationSearchCommand**](OrganizationSearchCommand.md) |  | 

### Return type

[**OrganizationSearchList**](OrganizationSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPartners

> PartnersSearchList SearchPartners(ctx).PartnersSearchCommand(partnersSearchCommand).Execute()

Global search for partners

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    partnersSearchCommand := *openapiclient.NewPartnersSearchCommand() // PartnersSearchCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchPartners(context.Background()).PartnersSearchCommand(partnersSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPartners`: PartnersSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partnersSearchCommand** | [**PartnersSearchCommand**](PartnersSearchCommand.md) |  | 

### Return type

[**PartnersSearchList**](PartnersSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPods

> PodsSearchList SearchPods(ctx).PodsSearchCommand(podsSearchCommand).Execute()

Global search for pods

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    podsSearchCommand := *openapiclient.NewPodsSearchCommand() // PodsSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchPods(context.Background()).PodsSearchCommand(podsSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPods`: PodsSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **podsSearchCommand** | [**PodsSearchCommand**](PodsSearchCommand.md) |  | 

### Return type

[**PodsSearchList**](PodsSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProjects

> ProjectsSearchList SearchProjects(ctx).ProjectsSearchCommand(projectsSearchCommand).Execute()

Global search for projects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    projectsSearchCommand := *openapiclient.NewProjectsSearchCommand() // ProjectsSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchProjects(context.Background()).ProjectsSearchCommand(projectsSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProjects`: ProjectsSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsSearchCommand** | [**ProjectsSearchCommand**](ProjectsSearchCommand.md) |  | 

### Return type

[**ProjectsSearchList**](ProjectsSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPrometheusRules

> PrometheusRulesSearchList SearchPrometheusRules(ctx).PrometheusRulesSearchCommand(prometheusRulesSearchCommand).Execute()

Global search for prometheus-rules

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    prometheusRulesSearchCommand := *openapiclient.NewPrometheusRulesSearchCommand() // PrometheusRulesSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchPrometheusRules(context.Background()).PrometheusRulesSearchCommand(prometheusRulesSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPrometheusRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPrometheusRules`: PrometheusRulesSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPrometheusRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPrometheusRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusRulesSearchCommand** | [**PrometheusRulesSearchCommand**](PrometheusRulesSearchCommand.md) |  | 

### Return type

[**PrometheusRulesSearchList**](PrometheusRulesSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPvcs

> PvcSearchList SearchPvcs(ctx).PvcSearchCommand(pvcSearchCommand).Execute()

Global search for pvcs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    pvcSearchCommand := *openapiclient.NewPvcSearchCommand() // PvcSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchPvcs(context.Background()).PvcSearchCommand(pvcSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPvcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPvcs`: PvcSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPvcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pvcSearchCommand** | [**PvcSearchCommand**](PvcSearchCommand.md) |  | 

### Return type

[**PvcSearchList**](PvcSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSecrets

> SecretSearchList SearchSecrets(ctx).SecretSearchCommand(secretSearchCommand).Execute()

Global search for secrets

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    secretSearchCommand := *openapiclient.NewSecretSearchCommand() // SecretSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchSecrets(context.Background()).SecretSearchCommand(secretSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecrets`: SecretSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secretSearchCommand** | [**SecretSearchCommand**](SecretSearchCommand.md) |  | 

### Return type

[**SecretSearchList**](SecretSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchServers

> ServersSearchList SearchServers(ctx).ServersSearchCommand(serversSearchCommand).Execute()

Global search for servers

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    serversSearchCommand := *openapiclient.NewServersSearchCommand() // ServersSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchServers(context.Background()).ServersSearchCommand(serversSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchServers`: ServersSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serversSearchCommand** | [**ServersSearchCommand**](ServersSearchCommand.md) |  | 

### Return type

[**ServersSearchList**](ServersSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchServices

> ServiceSearchList SearchServices(ctx).ServiceSearchCommand(serviceSearchCommand).Execute()

Global search for services

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    serviceSearchCommand := *openapiclient.NewServiceSearchCommand() // ServiceSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchServices(context.Background()).ServiceSearchCommand(serviceSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchServices`: ServiceSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceSearchCommand** | [**ServiceSearchCommand**](ServiceSearchCommand.md) |  | 

### Return type

[**ServiceSearchList**](ServiceSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStandAloneProfiles

> StandAloneProfilesSearchList SearchStandAloneProfiles(ctx).StandAloneProfilesSearchCommand(standAloneProfilesSearchCommand).Execute()

Global search for stand-alone-profiles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    standAloneProfilesSearchCommand := *openapiclient.NewStandAloneProfilesSearchCommand() // StandAloneProfilesSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchStandAloneProfiles(context.Background()).StandAloneProfilesSearchCommand(standAloneProfilesSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchStandAloneProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchStandAloneProfiles`: StandAloneProfilesSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchStandAloneProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStandAloneProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standAloneProfilesSearchCommand** | [**StandAloneProfilesSearchCommand**](StandAloneProfilesSearchCommand.md) |  | 

### Return type

[**StandAloneProfilesSearchList**](StandAloneProfilesSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSts

> StsSearchList SearchSts(ctx).StsSearchCommand(stsSearchCommand).Execute()

Global search for sts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    stsSearchCommand := *openapiclient.NewStsSearchCommand() // StsSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchSts(context.Background()).StsSearchCommand(stsSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSts`: StsSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stsSearchCommand** | [**StsSearchCommand**](StsSearchCommand.md) |  | 

### Return type

[**StsSearchList**](StsSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> UsersSearchList SearchUsers(ctx).UsersSearchCommand(usersSearchCommand).Execute()

Global search for users

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/smidra/taikungoclient"
)

func main() {
    usersSearchCommand := *openapiclient.NewUsersSearchCommand() // UsersSearchCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchUsers(context.Background()).UsersSearchCommand(usersSearchCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsers`: UsersSearchList
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersSearchCommand** | [**UsersSearchCommand**](UsersSearchCommand.md) |  | 

### Return type

[**UsersSearchList**](UsersSearchList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

