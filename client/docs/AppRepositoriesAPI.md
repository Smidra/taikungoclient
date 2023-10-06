# \AppRepositoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoryAvailableList**](AppRepositoriesAPI.md#RepositoryAvailableList) | **Get** /api/v1/repository/available | Retrieve available repositories
[**RepositoryBind**](AppRepositoriesAPI.md#RepositoryBind) | **Post** /api/v1/repository/bind | Bind repo to organization
[**RepositoryDelete**](AppRepositoriesAPI.md#RepositoryDelete) | **Post** /api/v1/repository/delete | Delete repo from organization
[**RepositoryDisable**](AppRepositoriesAPI.md#RepositoryDisable) | **Post** /api/v1/repository/disable | Disable repo from organization
[**RepositoryImport**](AppRepositoriesAPI.md#RepositoryImport) | **Post** /api/v1/repository/import | Import repo to artifact
[**RepositoryRecommendedList**](AppRepositoriesAPI.md#RepositoryRecommendedList) | **Get** /api/v1/repository/recommended | Retrieve taikun recommended repositories
[**RepositoryUnbind**](AppRepositoriesAPI.md#RepositoryUnbind) | **Post** /api/v1/repository/unbind | Unbind repo from organization



## RepositoryAvailableList

> AppRepositoryList RepositoryAvailableList(ctx).IsPrivate(isPrivate).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve available repositories

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/client"
)

func main() {
    isPrivate := true // bool | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := "id_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRepositoriesAPI.RepositoryAvailableList(context.Background()).IsPrivate(isPrivate).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryAvailableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryAvailableList`: AppRepositoryList
    fmt.Fprintf(os.Stdout, "Response from `AppRepositoriesAPI.RepositoryAvailableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryAvailableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPrivate** | **bool** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **string** |  | 

### Return type

[**AppRepositoryList**](AppRepositoryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryBind

> RepositoryBind(ctx).BindAppRepositoryCommand(bindAppRepositoryCommand).Execute()

Bind repo to organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/client"
)

func main() {
    bindAppRepositoryCommand := *openapiclient.NewBindAppRepositoryCommand() // BindAppRepositoryCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppRepositoriesAPI.RepositoryBind(context.Background()).BindAppRepositoryCommand(bindAppRepositoryCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryBind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindAppRepositoryCommand** | [**BindAppRepositoryCommand**](BindAppRepositoryCommand.md) |  | 

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


## RepositoryDelete

> RepositoryDelete(ctx).DeleteRepositoryCommand(deleteRepositoryCommand).Execute()

Delete repo from organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/client"
)

func main() {
    deleteRepositoryCommand := *openapiclient.NewDeleteRepositoryCommand() // DeleteRepositoryCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppRepositoriesAPI.RepositoryDelete(context.Background()).DeleteRepositoryCommand(deleteRepositoryCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRepositoryCommand** | [**DeleteRepositoryCommand**](DeleteRepositoryCommand.md) |  | 

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


## RepositoryDisable

> RepositoryDisable(ctx).DisableRepoCommand(disableRepoCommand).Execute()

Disable repo from organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/client"
)

func main() {
    disableRepoCommand := *openapiclient.NewDisableRepoCommand() // DisableRepoCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppRepositoriesAPI.RepositoryDisable(context.Background()).DisableRepoCommand(disableRepoCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryDisable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableRepoCommand** | [**DisableRepoCommand**](DisableRepoCommand.md) |  | 

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


## RepositoryImport

> RepositoryImport(ctx).ImportRepoCommand(importRepoCommand).Execute()

Import repo to artifact

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/client"
)

func main() {
    importRepoCommand := *openapiclient.NewImportRepoCommand() // ImportRepoCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppRepositoriesAPI.RepositoryImport(context.Background()).ImportRepoCommand(importRepoCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importRepoCommand** | [**ImportRepoCommand**](ImportRepoCommand.md) |  | 

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


## RepositoryRecommendedList

> []ArtifactRepositoryDto RepositoryRecommendedList(ctx).Execute()

Retrieve taikun recommended repositories

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRepositoriesAPI.RepositoryRecommendedList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryRecommendedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryRecommendedList`: []ArtifactRepositoryDto
    fmt.Fprintf(os.Stdout, "Response from `AppRepositoriesAPI.RepositoryRecommendedList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryRecommendedListRequest struct via the builder pattern


### Return type

[**[]ArtifactRepositoryDto**](ArtifactRepositoryDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryUnbind

> RepositoryUnbind(ctx).UnbindAppRepositoryCommand(unbindAppRepositoryCommand).Execute()

Unbind repo from organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient/client"
)

func main() {
    unbindAppRepositoryCommand := *openapiclient.NewUnbindAppRepositoryCommand() // UnbindAppRepositoryCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppRepositoriesAPI.RepositoryUnbind(context.Background()).UnbindAppRepositoryCommand(unbindAppRepositoryCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryUnbind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryUnbindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unbindAppRepositoryCommand** | [**UnbindAppRepositoryCommand**](UnbindAppRepositoryCommand.md) |  | 

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

