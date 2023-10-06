# \DnsServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsserversCreate**](DnsServersAPI.md#DnsserversCreate) | **Post** /api/v1/dnsservers/create | Create dns servers for access profile
[**DnsserversDelete**](DnsServersAPI.md#DnsserversDelete) | **Delete** /api/v1/dnsservers/{id} | Delete dns server
[**DnsserversEdit**](DnsServersAPI.md#DnsserversEdit) | **Put** /api/v1/dnsservers/edit/{id} | Edit dns server
[**DnsserversList**](DnsServersAPI.md#DnsserversList) | **Get** /api/v1/dnsservers/{accessProfileId} | List dn servers by profile id



## DnsserversCreate

> ApiResponse DnsserversCreate(ctx).CreateDnsServerCommand(createDnsServerCommand).Execute()

Create dns servers for access profile

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
    createDnsServerCommand := *openapiclient.NewCreateDnsServerCommand() // CreateDnsServerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsServersAPI.DnsserversCreate(context.Background()).CreateDnsServerCommand(createDnsServerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersAPI.DnsserversCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsserversCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DnsServersAPI.DnsserversCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsserversCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDnsServerCommand** | [**CreateDnsServerCommand**](CreateDnsServerCommand.md) |  | 

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


## DnsserversDelete

> DnsserversDelete(ctx, id).Execute()

Delete dns server

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DnsServersAPI.DnsserversDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersAPI.DnsserversDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDnsserversDeleteRequest struct via the builder pattern


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


## DnsserversEdit

> DnsserversEdit(ctx, id).DnsNtpAddressEditDto(dnsNtpAddressEditDto).Execute()

Edit dns server

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
    id := int32(56) // int32 | 
    dnsNtpAddressEditDto := *openapiclient.NewDnsNtpAddressEditDto() // DnsNtpAddressEditDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DnsServersAPI.DnsserversEdit(context.Background(), id).DnsNtpAddressEditDto(dnsNtpAddressEditDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersAPI.DnsserversEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDnsserversEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsNtpAddressEditDto** | [**DnsNtpAddressEditDto**](DnsNtpAddressEditDto.md) |  | 

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


## DnsserversList

> []DnsServersListDto DnsserversList(ctx, accessProfileId).Search(search).Execute()

List dn servers by profile id

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
    accessProfileId := int32(56) // int32 | 
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsServersAPI.DnsserversList(context.Background(), accessProfileId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsServersAPI.DnsserversList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsserversList`: []DnsServersListDto
    fmt.Fprintf(os.Stdout, "Response from `DnsServersAPI.DnsserversList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsserversListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]DnsServersListDto**](DnsServersListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

