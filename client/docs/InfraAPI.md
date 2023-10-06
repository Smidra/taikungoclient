# \InfraAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfraCreate**](InfraAPI.md#InfraCreate) | **Post** /api/v1/infra/create | Create infra product
[**InfraDetails**](InfraAPI.md#InfraDetails) | **Get** /api/v1/infra/details | Retrieve infra details
[**InfraOrganizationsList**](InfraAPI.md#InfraOrganizationsList) | **Get** /api/v1/infra/organizations-list | Retrieve infra products list
[**InfraProductList**](InfraAPI.md#InfraProductList) | **Get** /api/v1/infra/list | Retrieve infra products list



## InfraCreate

> InfraCreate(ctx).CreateInfraProductCommand(createInfraProductCommand).Execute()

Create infra product

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
    createInfraProductCommand := *openapiclient.NewCreateInfraProductCommand() // CreateInfraProductCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InfraAPI.InfraCreate(context.Background()).CreateInfraProductCommand(createInfraProductCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraAPI.InfraCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInfraCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInfraProductCommand** | [**CreateInfraProductCommand**](CreateInfraProductCommand.md) |  | 

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


## InfraDetails

> OpenstackQuotaList InfraDetails(ctx).OrganizationId(organizationId).Execute()

Retrieve infra details

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfraAPI.InfraDetails(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraAPI.InfraDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfraDetails`: OpenstackQuotaList
    fmt.Fprintf(os.Stdout, "Response from `InfraAPI.InfraDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInfraDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**OpenstackQuotaList**](OpenstackQuotaList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfraOrganizationsList

> []InfraOrganizationsListDto InfraOrganizationsList(ctx).Execute()

Retrieve infra products list

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfraAPI.InfraOrganizationsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraAPI.InfraOrganizationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfraOrganizationsList`: []InfraOrganizationsListDto
    fmt.Fprintf(os.Stdout, "Response from `InfraAPI.InfraOrganizationsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfraOrganizationsListRequest struct via the builder pattern


### Return type

[**[]InfraOrganizationsListDto**](InfraOrganizationsListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfraProductList

> []InfraProductDto InfraProductList(ctx).Execute()

Retrieve infra products list

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfraAPI.InfraProductList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraAPI.InfraProductList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfraProductList`: []InfraProductDto
    fmt.Fprintf(os.Stdout, "Response from `InfraAPI.InfraProductList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfraProductListRequest struct via the builder pattern


### Return type

[**[]InfraProductDto**](InfraProductDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

