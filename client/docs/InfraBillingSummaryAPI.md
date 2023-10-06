# \InfraBillingSummaryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfraBillingSummaryCreate**](InfraBillingSummaryAPI.md#InfraBillingSummaryCreate) | **Post** /api/v1/infra-billing-summary/create | Add infra billing summary
[**InfraBillingSummaryList**](InfraBillingSummaryAPI.md#InfraBillingSummaryList) | **Get** /api/v1/infra-billing-summary/list | Retrieve infra billing info



## InfraBillingSummaryCreate

> InfraBillingSummaryCreate(ctx).InfraBillingSummariesCreateCommand(infraBillingSummariesCreateCommand).Execute()

Add infra billing summary

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
    infraBillingSummariesCreateCommand := *openapiclient.NewInfraBillingSummariesCreateCommand() // InfraBillingSummariesCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InfraBillingSummaryAPI.InfraBillingSummaryCreate(context.Background()).InfraBillingSummariesCreateCommand(infraBillingSummariesCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraBillingSummaryAPI.InfraBillingSummaryCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInfraBillingSummaryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **infraBillingSummariesCreateCommand** | [**InfraBillingSummariesCreateCommand**](InfraBillingSummariesCreateCommand.md) |  | 

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


## InfraBillingSummaryList

> InfraBillingInfo InfraBillingSummaryList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).Execute()

Retrieve infra billing info

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfraBillingSummaryAPI.InfraBillingSummaryList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraBillingSummaryAPI.InfraBillingSummaryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfraBillingSummaryList`: InfraBillingInfo
    fmt.Fprintf(os.Stdout, "Response from `InfraBillingSummaryAPI.InfraBillingSummaryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInfraBillingSummaryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**InfraBillingInfo**](InfraBillingInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

