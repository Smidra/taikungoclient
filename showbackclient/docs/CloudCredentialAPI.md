# \CloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudcredentialsAllFlavors**](CloudCredentialAPI.md#CloudcredentialsAllFlavors) | **Get** /api/v1/cloudcredentials/flavors/{cloudId} | 
[**CloudcredentialsDashboardList**](CloudCredentialAPI.md#CloudcredentialsDashboardList) | **Get** /api/v1/cloudcredentials/list | Retrieve all cloud credentials
[**CloudcredentialsDelete**](CloudCredentialAPI.md#CloudcredentialsDelete) | **Delete** /api/v1/cloudcredentials/{cloudId} | Remove cloud credential by cloud Id
[**CloudcredentialsExceeded**](CloudCredentialAPI.md#CloudcredentialsExceeded) | **Get** /api/v1/cloudcredentials/exceeded-quotas | Retrieve cloud credentials exceeded quotas
[**CloudcredentialsForCli**](CloudCredentialAPI.md#CloudcredentialsForCli) | **Get** /api/v1/cloudcredentials/cli | Retrieve cloud credentials for CLI
[**CloudcredentialsForProject**](CloudCredentialAPI.md#CloudcredentialsForProject) | **Get** /api/v1/cloudcredentials/details | Retrieve cloud credential details by cloud Id
[**CloudcredentialsLockManager**](CloudCredentialAPI.md#CloudcredentialsLockManager) | **Post** /api/v1/cloudcredentials/lockmanager | Lock/Unlock cloud credential
[**CloudcredentialsMakeDefault**](CloudCredentialAPI.md#CloudcredentialsMakeDefault) | **Post** /api/v1/cloudcredentials/makedefault | Make cloud credentials default
[**CloudcredentialsOrgList**](CloudCredentialAPI.md#CloudcredentialsOrgList) | **Get** /api/v1/cloudcredentials | Retrieve a list of cloud credentials by organization Id



## CloudcredentialsAllFlavors

> AllFlavorsList CloudcredentialsAllFlavors(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    cloudId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    startRam := float64(1.2) // float64 |  (optional)
    endRam := float64(1.2) // float64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsAllFlavors(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsAllFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudcredentialsAllFlavors`: AllFlavorsList
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsAllFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsAllFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**AllFlavorsList**](AllFlavorsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsDashboardList

> CredentialsChart CloudcredentialsDashboardList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve all cloud credentials

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsDashboardList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudcredentialsDashboardList`: CredentialsChart
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsDashboardList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**CredentialsChart**](CredentialsChart.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsDelete

> CloudcredentialsDelete(ctx, cloudId).Execute()

Remove cloud credential by cloud Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    cloudId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudCredentialAPI.CloudcredentialsDelete(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsDeleteRequest struct via the builder pattern


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


## CloudcredentialsExceeded

> ExceededQuotaList CloudcredentialsExceeded(ctx).OrganizationId(organizationId).Execute()

Retrieve cloud credentials exceeded quotas

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsExceeded(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsExceeded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudcredentialsExceeded`: ExceededQuotaList
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsExceeded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsExceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**ExceededQuotaList**](ExceededQuotaList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsForCli

> CredentialsForCli CloudcredentialsForCli(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Execute()

Retrieve cloud credentials for CLI

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsForCli(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsForCli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudcredentialsForCli`: CredentialsForCli
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsForCli`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsForCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**CredentialsForCli**](CredentialsForCli.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsForProject

> CredentialsForProjectList CloudcredentialsForProject(ctx).CloudId(cloudId).Execute()

Retrieve cloud credential details by cloud Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    cloudId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsForProject(context.Background()).CloudId(cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudcredentialsForProject`: CredentialsForProjectList
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsForProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudId** | **int32** |  | 

### Return type

[**CredentialsForProjectList**](CredentialsForProjectList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsLockManager

> CloudcredentialsLockManager(ctx).CloudLockManagerCommand(cloudLockManagerCommand).Execute()

Lock/Unlock cloud credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    cloudLockManagerCommand := *openapiclient.NewCloudLockManagerCommand() // CloudLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudCredentialAPI.CloudcredentialsLockManager(context.Background()).CloudLockManagerCommand(cloudLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudLockManagerCommand** | [**CloudLockManagerCommand**](CloudLockManagerCommand.md) |  | 

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


## CloudcredentialsMakeDefault

> CloudcredentialsMakeDefault(ctx).CredentialMakeDefaultCommand(credentialMakeDefaultCommand).Execute()

Make cloud credentials default

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    credentialMakeDefaultCommand := *openapiclient.NewCredentialMakeDefaultCommand() // CredentialMakeDefaultCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudCredentialAPI.CloudcredentialsMakeDefault(context.Background()).CredentialMakeDefaultCommand(credentialMakeDefaultCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsMakeDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsMakeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialMakeDefaultCommand** | [**CredentialMakeDefaultCommand**](CredentialMakeDefaultCommand.md) |  | 

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


## CloudcredentialsOrgList

> []CloudCredentialsForOrganizationEntity CloudcredentialsOrgList(ctx).IsAdmin(isAdmin).OrganizationId(organizationId).Search(search).IsInfra(isInfra).Execute()

Retrieve a list of cloud credentials by organization Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/showbackclient"
)

func main() {
    isAdmin := true // bool | 
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    isInfra := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsOrgList(context.Background()).IsAdmin(isAdmin).OrganizationId(organizationId).Search(search).IsInfra(isInfra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsOrgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudcredentialsOrgList`: []CloudCredentialsForOrganizationEntity
    fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsOrgList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsOrgListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isAdmin** | **bool** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **isInfra** | **bool** |  | 

### Return type

[**[]CloudCredentialsForOrganizationEntity**](CloudCredentialsForOrganizationEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

