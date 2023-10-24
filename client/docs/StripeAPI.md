# \StripeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StripeSubscriptionItemList**](StripeAPI.md#StripeSubscriptionItemList) | **Get** /api/v1/stripe/{subscriptionId} | 



## StripeSubscriptionItemList

> []StripeSubscriptionItemDto StripeSubscriptionItemList(ctx, subscriptionId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/itera-io/taikungoclient/client"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeAPI.StripeSubscriptionItemList(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeAPI.StripeSubscriptionItemList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StripeSubscriptionItemList`: []StripeSubscriptionItemDto
    fmt.Fprintf(os.Stdout, "Response from `StripeAPI.StripeSubscriptionItemList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStripeSubscriptionItemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StripeSubscriptionItemDto**](StripeSubscriptionItemDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

