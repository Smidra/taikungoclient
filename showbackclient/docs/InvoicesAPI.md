# \InvoicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoicesCreate**](InvoicesAPI.md#InvoicesCreate) | **Post** /api/v1/invoices/create | Create invoice
[**InvoicesDownload**](InvoicesAPI.md#InvoicesDownload) | **Post** /api/v1/invoices/download | Download invoice
[**InvoicesList**](InvoicesAPI.md#InvoicesList) | **Get** /api/v1/invoices/list | Invoices list
[**InvoicesUpdate**](InvoicesAPI.md#InvoicesUpdate) | **Put** /api/v1/invoices/update/{invoiceId} | Update invoice



## InvoicesCreate

> int32 InvoicesCreate(ctx).CreateInvoiceCommand(createInvoiceCommand).Execute()

Create invoice

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
    createInvoiceCommand := *openapiclient.NewCreateInvoiceCommand() // CreateInvoiceCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.InvoicesCreate(context.Background()).CreateInvoiceCommand(createInvoiceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoicesCreate`: int32
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvoiceCommand** | [**CreateInvoiceCommand**](CreateInvoiceCommand.md) |  | 

### Return type

**int32**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesDownload

> CsvExporter InvoicesDownload(ctx).DownloadInvoiceCommand(downloadInvoiceCommand).Execute()

Download invoice

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
    downloadInvoiceCommand := *openapiclient.NewDownloadInvoiceCommand() // DownloadInvoiceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.InvoicesDownload(context.Background()).DownloadInvoiceCommand(downloadInvoiceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoicesDownload`: CsvExporter
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadInvoiceCommand** | [**DownloadInvoiceCommand**](DownloadInvoiceCommand.md) |  | 

### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesList

> Invoices InvoicesList(ctx).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).Search(search).FilterBy(filterBy).SortBy(sortBy).SortDirection(sortDirection).OrganizationId(organizationId).PartnerId(partnerId).Execute()

Invoices list

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
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    partnerId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.InvoicesList(context.Background()).Offset(offset).Limit(limit).StartDate(startDate).EndDate(endDate).Search(search).FilterBy(filterBy).SortBy(sortBy).SortDirection(sortDirection).OrganizationId(organizationId).PartnerId(partnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoicesList`: Invoices
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **search** | **string** |  | 
 **filterBy** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **organizationId** | **int32** |  | 
 **partnerId** | **int32** |  | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesUpdate

> InvoicesUpdate(ctx, invoiceId).UpdateInvoiceDto(updateInvoiceDto).Execute()

Update invoice

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
    invoiceId := int32(56) // int32 | 
    updateInvoiceDto := *openapiclient.NewUpdateInvoiceDto() // UpdateInvoiceDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InvoicesAPI.InvoicesUpdate(context.Background(), invoiceId).UpdateInvoiceDto(updateInvoiceDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInvoiceDto** | [**UpdateInvoiceDto**](UpdateInvoiceDto.md) |  | 

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

