# \ProjectAppsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectappAutosync**](ProjectAppsAPI.md#ProjectappAutosync) | **Post** /api/v1/projectapp/autosync | AutoSync management
[**ProjectappDelete**](ProjectAppsAPI.md#ProjectappDelete) | **Delete** /api/v1/projectapp/uninstall/{projectAppId} | Uninstall application
[**ProjectappDetails**](ProjectAppsAPI.md#ProjectappDetails) | **Get** /api/v1/projectapp/{id} | Retrieve project app&#39;s details
[**ProjectappInstall**](ProjectAppsAPI.md#ProjectappInstall) | **Post** /api/v1/projectapp/install | Install an application
[**ProjectappList**](ProjectAppsAPI.md#ProjectappList) | **Get** /api/v1/projectapp/list | Retrieve all project apps according to current organization
[**ProjectappLockManager**](ProjectAppsAPI.md#ProjectappLockManager) | **Post** /api/v1/projectapp/lockmanager | Lock/Unlock project app
[**ProjectappSync**](ProjectAppsAPI.md#ProjectappSync) | **Post** /api/v1/projectapp/sync | Sync an application



## ProjectappAutosync

> ProjectappAutosync(ctx).AutoSyncManagementCommand(autoSyncManagementCommand).Execute()

AutoSync management

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
    autoSyncManagementCommand := *openapiclient.NewAutoSyncManagementCommand() // AutoSyncManagementCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectAppsAPI.ProjectappAutosync(context.Background()).AutoSyncManagementCommand(autoSyncManagementCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppsAPI.ProjectappAutosync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectappAutosyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoSyncManagementCommand** | [**AutoSyncManagementCommand**](AutoSyncManagementCommand.md) |  | 

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


## ProjectappDelete

> ProjectappDelete(ctx, projectAppId).Execute()

Uninstall application

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
    projectAppId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectAppsAPI.ProjectappDelete(context.Background(), projectAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppsAPI.ProjectappDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectappDeleteRequest struct via the builder pattern


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


## ProjectappDetails

> ProjectAppDetailsDto ProjectappDetails(ctx, id).Execute()

Retrieve project app's details

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAppsAPI.ProjectappDetails(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppsAPI.ProjectappDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectappDetails`: ProjectAppDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectAppsAPI.ProjectappDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectappDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectAppDetailsDto**](ProjectAppDetailsDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectappInstall

> ApiResponse ProjectappInstall(ctx).CreateProjectAppCommand(createProjectAppCommand).Execute()

Install an application

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
    createProjectAppCommand := *openapiclient.NewCreateProjectAppCommand() // CreateProjectAppCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAppsAPI.ProjectappInstall(context.Background()).CreateProjectAppCommand(createProjectAppCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppsAPI.ProjectappInstall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectappInstall`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectAppsAPI.ProjectappInstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectappInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectAppCommand** | [**CreateProjectAppCommand**](CreateProjectAppCommand.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectappList

> ProjectAppList ProjectappList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve all project apps according to current organization

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAppsAPI.ProjectappList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppsAPI.ProjectappList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectappList`: ProjectAppList
    fmt.Fprintf(os.Stdout, "Response from `ProjectAppsAPI.ProjectappList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectappListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**ProjectAppList**](ProjectAppList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectappLockManager

> ProjectappLockManager(ctx).LockProjectAppCommand(lockProjectAppCommand).Execute()

Lock/Unlock project app

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
    lockProjectAppCommand := *openapiclient.NewLockProjectAppCommand() // LockProjectAppCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectAppsAPI.ProjectappLockManager(context.Background()).LockProjectAppCommand(lockProjectAppCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppsAPI.ProjectappLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectappLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lockProjectAppCommand** | [**LockProjectAppCommand**](LockProjectAppCommand.md) |  | 

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


## ProjectappSync

> ProjectappSync(ctx).SyncProjectAppCommand(syncProjectAppCommand).Execute()

Sync an application

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
    syncProjectAppCommand := *openapiclient.NewSyncProjectAppCommand() // SyncProjectAppCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectAppsAPI.ProjectappSync(context.Background()).SyncProjectAppCommand(syncProjectAppCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppsAPI.ProjectappSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectappSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncProjectAppCommand** | [**SyncProjectAppCommand**](SyncProjectAppCommand.md) |  | 

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

