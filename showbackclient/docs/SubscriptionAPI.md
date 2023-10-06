# \SubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionBind**](SubscriptionAPI.md#SubscriptionBind) | **Post** /api/v1/subscription/bind | Bind subscription
[**SubscriptionBoundList**](SubscriptionAPI.md#SubscriptionBoundList) | **Get** /api/v1/subscription/boundlist | Retrieve subscription for organization bound
[**SubscriptionDelete**](SubscriptionAPI.md#SubscriptionDelete) | **Post** /api/v1/subscription/delete | Delete subscription package
[**SubscriptionList**](SubscriptionAPI.md#SubscriptionList) | **Get** /api/v1/subscription | Retrieve private subscription list for partners
[**SubscriptionPublic**](SubscriptionAPI.md#SubscriptionPublic) | **Get** /api/v1/subscription/public | Retrieve subscription for organization bound
[**SubscriptionSubscription**](SubscriptionAPI.md#SubscriptionSubscription) | **Post** /api/v1/subscription/create | Add new subscription package
[**SubscriptionUpdate**](SubscriptionAPI.md#SubscriptionUpdate) | **Post** /api/v1/subscription/update | Update subscription



## SubscriptionBind

> BindSubscriptionResponseDto SubscriptionBind(ctx).BindSubscriptionCommand(bindSubscriptionCommand).Execute()

Bind subscription

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
    bindSubscriptionCommand := *openapiclient.NewBindSubscriptionCommand() // BindSubscriptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.SubscriptionBind(context.Background()).BindSubscriptionCommand(bindSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionBind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionBind`: BindSubscriptionResponseDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionBind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindSubscriptionCommand** | [**BindSubscriptionCommand**](BindSubscriptionCommand.md) |  | 

### Return type

[**BindSubscriptionResponseDto**](BindSubscriptionResponseDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionBoundList

> []ListForOrganizationEditDto SubscriptionBoundList(ctx).Search(search).Execute()

Retrieve subscription for organization bound

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.SubscriptionBoundList(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionBoundList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionBoundList`: []ListForOrganizationEditDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionBoundList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionBoundListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]ListForOrganizationEditDto**](ListForOrganizationEditDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionDelete

> SubscriptionDelete(ctx).DeleteSubscriptionCommand(deleteSubscriptionCommand).Execute()

Delete subscription package

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
    deleteSubscriptionCommand := *openapiclient.NewDeleteSubscriptionCommand() // DeleteSubscriptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionAPI.SubscriptionDelete(context.Background()).DeleteSubscriptionCommand(deleteSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSubscriptionCommand** | [**DeleteSubscriptionCommand**](DeleteSubscriptionCommand.md) |  | 

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


## SubscriptionList

> map[string]interface{} SubscriptionList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

Retrieve private subscription list for partners

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.SubscriptionList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPublic

> []ListForLandingPageDto SubscriptionPublic(ctx).Execute()

Retrieve subscription for organization bound

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.SubscriptionPublic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionPublic`: []ListForLandingPageDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPublicRequest struct via the builder pattern


### Return type

[**[]ListForLandingPageDto**](ListForLandingPageDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionSubscription

> SubscriptionSubscription(ctx).CreateSubscriptionCommand(createSubscriptionCommand).Execute()

Add new subscription package

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
    createSubscriptionCommand := *openapiclient.NewCreateSubscriptionCommand() // CreateSubscriptionCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionAPI.SubscriptionSubscription(context.Background()).CreateSubscriptionCommand(createSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubscriptionCommand** | [**CreateSubscriptionCommand**](CreateSubscriptionCommand.md) |  | 

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


## SubscriptionUpdate

> SubscriptionUpdate(ctx).UpdateSubscriptionCommand(updateSubscriptionCommand).Execute()

Update subscription

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
    updateSubscriptionCommand := *openapiclient.NewUpdateSubscriptionCommand() // UpdateSubscriptionCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionAPI.SubscriptionUpdate(context.Background()).UpdateSubscriptionCommand(updateSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSubscriptionCommand** | [**UpdateSubscriptionCommand**](UpdateSubscriptionCommand.md) |  | 

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

