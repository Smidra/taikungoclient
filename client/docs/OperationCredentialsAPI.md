# \OperationCredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpscredentialsCreate**](OperationCredentialsAPI.md#OpscredentialsCreate) | **Post** /api/v1/opscredentials | Add Operation credentials
[**OpscredentialsDelete**](OperationCredentialsAPI.md#OpscredentialsDelete) | **Delete** /api/v1/opscredentials/{id} | Remove Operation credential by Id
[**OpscredentialsList**](OperationCredentialsAPI.md#OpscredentialsList) | **Get** /api/v1/opscredentials/list | Retrieve all operation credentials
[**OpscredentialsListByOrganizationId**](OperationCredentialsAPI.md#OpscredentialsListByOrganizationId) | **Get** /api/v1/opscredentials | Retrieve operation credentials by organization Id
[**OpscredentialsLockManager**](OperationCredentialsAPI.md#OpscredentialsLockManager) | **Post** /api/v1/opscredentials/lockmanager | Lock/Unlock operation credential
[**OpscredentialsMakeDefault**](OperationCredentialsAPI.md#OpscredentialsMakeDefault) | **Post** /api/v1/opscredentials/makedefault | Choose default operation credential
[**OpscredentialsMetricNames**](OperationCredentialsAPI.md#OpscredentialsMetricNames) | **Get** /api/v1/opscredentials/{id}/metric/names | Fetch prometheus metric names



## OpscredentialsCreate

> ApiResponse OpscredentialsCreate(ctx).OperationCredentialsCreateCommand(operationCredentialsCreateCommand).Execute()

Add Operation credentials

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
    operationCredentialsCreateCommand := *openapiclient.NewOperationCredentialsCreateCommand() // OperationCredentialsCreateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationCredentialsAPI.OpscredentialsCreate(context.Background()).OperationCredentialsCreateCommand(operationCredentialsCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsAPI.OpscredentialsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpscredentialsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `OperationCredentialsAPI.OpscredentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCredentialsCreateCommand** | [**OperationCredentialsCreateCommand**](OperationCredentialsCreateCommand.md) |  | 

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


## OpscredentialsDelete

> OpscredentialsDelete(ctx, id).Execute()

Remove Operation credential by Id

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
    r, err := apiClient.OperationCredentialsAPI.OpscredentialsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsAPI.OpscredentialsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsDeleteRequest struct via the builder pattern


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


## OpscredentialsList

> OperationCredentials OpscredentialsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve all operation credentials

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationCredentialsAPI.OpscredentialsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsAPI.OpscredentialsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpscredentialsList`: OperationCredentials
    fmt.Fprintf(os.Stdout, "Response from `OperationCredentialsAPI.OpscredentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**OperationCredentials**](OperationCredentials.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpscredentialsListByOrganizationId

> []OperationCredentialsForOrganizationEntity OpscredentialsListByOrganizationId(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve operation credentials by organization Id

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
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationCredentialsAPI.OpscredentialsListByOrganizationId(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsAPI.OpscredentialsListByOrganizationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpscredentialsListByOrganizationId`: []OperationCredentialsForOrganizationEntity
    fmt.Fprintf(os.Stdout, "Response from `OperationCredentialsAPI.OpscredentialsListByOrganizationId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsListByOrganizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]OperationCredentialsForOrganizationEntity**](OperationCredentialsForOrganizationEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpscredentialsLockManager

> OpscredentialsLockManager(ctx).OperationCredentialLockManagerCommand(operationCredentialLockManagerCommand).Execute()

Lock/Unlock operation credential

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
    operationCredentialLockManagerCommand := *openapiclient.NewOperationCredentialLockManagerCommand() // OperationCredentialLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OperationCredentialsAPI.OpscredentialsLockManager(context.Background()).OperationCredentialLockManagerCommand(operationCredentialLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsAPI.OpscredentialsLockManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCredentialLockManagerCommand** | [**OperationCredentialLockManagerCommand**](OperationCredentialLockManagerCommand.md) |  | 

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


## OpscredentialsMakeDefault

> OpscredentialsMakeDefault(ctx).OperationCredentialsMakeDefaultCommand(operationCredentialsMakeDefaultCommand).Execute()

Choose default operation credential

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
    operationCredentialsMakeDefaultCommand := *openapiclient.NewOperationCredentialsMakeDefaultCommand() // OperationCredentialsMakeDefaultCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OperationCredentialsAPI.OpscredentialsMakeDefault(context.Background()).OperationCredentialsMakeDefaultCommand(operationCredentialsMakeDefaultCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsAPI.OpscredentialsMakeDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsMakeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCredentialsMakeDefaultCommand** | [**OperationCredentialsMakeDefaultCommand**](OperationCredentialsMakeDefaultCommand.md) |  | 

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


## OpscredentialsMetricNames

> interface{} OpscredentialsMetricNames(ctx, id).Execute()

Fetch prometheus metric names

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
    resp, r, err := apiClient.OperationCredentialsAPI.OpscredentialsMetricNames(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationCredentialsAPI.OpscredentialsMetricNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpscredentialsMetricNames`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OperationCredentialsAPI.OpscredentialsMetricNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpscredentialsMetricNamesRequest struct via the builder pattern


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

