# \KeycloakAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeycloakCreate**](KeycloakAPI.md#KeycloakCreate) | **Post** /api/v1/keycloak/create | Create keycloak configuration for organization
[**KeycloakDelete**](KeycloakAPI.md#KeycloakDelete) | **Post** /api/v1/keycloak/delete | Delete keycloak configuration
[**KeycloakEdit**](KeycloakAPI.md#KeycloakEdit) | **Post** /api/v1/keycloak/edit | Edit keycloak configuration for organization
[**KeycloakList**](KeycloakAPI.md#KeycloakList) | **Get** /api/v1/keycloak | Get keycloak configuration



## KeycloakCreate

> KeycloakCreate(ctx).KeycloakCreateCommand(keycloakCreateCommand).Execute()

Create keycloak configuration for organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    keycloakCreateCommand := *openapiclient.NewKeycloakCreateCommand() // KeycloakCreateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeycloakAPI.KeycloakCreate(context.Background()).KeycloakCreateCommand(keycloakCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakAPI.KeycloakCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keycloakCreateCommand** | [**KeycloakCreateCommand**](KeycloakCreateCommand.md) |  | 

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


## KeycloakDelete

> KeycloakDelete(ctx).KeycloakDeleteCommand(keycloakDeleteCommand).Execute()

Delete keycloak configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    keycloakDeleteCommand := *openapiclient.NewKeycloakDeleteCommand() // KeycloakDeleteCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeycloakAPI.KeycloakDelete(context.Background()).KeycloakDeleteCommand(keycloakDeleteCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakAPI.KeycloakDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keycloakDeleteCommand** | [**KeycloakDeleteCommand**](KeycloakDeleteCommand.md) |  | 

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


## KeycloakEdit

> KeycloakEdit(ctx).KeycloakEditCommand(keycloakEditCommand).Execute()

Edit keycloak configuration for organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {
    keycloakEditCommand := *openapiclient.NewKeycloakEditCommand() // KeycloakEditCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeycloakAPI.KeycloakEdit(context.Background()).KeycloakEditCommand(keycloakEditCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakAPI.KeycloakEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keycloakEditCommand** | [**KeycloakEditCommand**](KeycloakEditCommand.md) |  | 

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


## KeycloakList

> KeycloakListDto KeycloakList(ctx).Execute()

Get keycloak configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Smidra/taikungoclient"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeycloakAPI.KeycloakList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeycloakAPI.KeycloakList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeycloakList`: KeycloakListDto
    fmt.Fprintf(os.Stdout, "Response from `KeycloakAPI.KeycloakList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeycloakListRequest struct via the builder pattern


### Return type

[**KeycloakListDto**](KeycloakListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

