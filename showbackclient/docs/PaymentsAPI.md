# \PaymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentBillingInfo**](PaymentsAPI.md#PaymentBillingInfo) | **Get** /api/v1/payment/billing-info | Get billing info for organization
[**PaymentCardinfo**](PaymentsAPI.md#PaymentCardinfo) | **Get** /api/v1/payment/cardinfo | Get card information
[**PaymentClear**](PaymentsAPI.md#PaymentClear) | **Post** /api/v1/payment/clear | Clear payment
[**PaymentCreateCustomer**](PaymentsAPI.md#PaymentCreateCustomer) | **Post** /api/v1/payment/createcustomer | Create customer
[**PaymentFinalPrice**](PaymentsAPI.md#PaymentFinalPrice) | **Post** /api/v1/payment/finalprice | Fetch final price
[**PaymentGetStripeInvoices**](PaymentsAPI.md#PaymentGetStripeInvoices) | **Get** /api/v1/payment/stripeinvoices/{subscriptionId} | 
[**PaymentPay**](PaymentsAPI.md#PaymentPay) | **Post** /api/v1/payment/pay | Pay invoice
[**PaymentUpdateCard**](PaymentsAPI.md#PaymentUpdateCard) | **Post** /api/v1/payment/updatecard | Update payment card
[**PaymentWebhook**](PaymentsAPI.md#PaymentWebhook) | **Post** /api/v1/payment/webhook | Listen to payment webhook



## PaymentBillingInfo

> BillingInfoDto PaymentBillingInfo(ctx).Execute()

Get billing info for organization

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
    resp, r, err := apiClient.PaymentsAPI.PaymentBillingInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentBillingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentBillingInfo`: BillingInfoDto
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentBillingInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentBillingInfoRequest struct via the builder pattern


### Return type

[**BillingInfoDto**](BillingInfoDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentCardinfo

> CardInformationDto PaymentCardinfo(ctx).OrganizationId(organizationId).Execute()

Get card information

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
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.PaymentCardinfo(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentCardinfo`: CardInformationDto
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**CardInformationDto**](CardInformationDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentClear

> PaymentClear(ctx).Body(body).Execute()

Clear payment

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
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsAPI.PaymentClear(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentClear``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentClearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## PaymentCreateCustomer

> PaymentCreateCustomer(ctx).CreateStripeCustomerCommand(createStripeCustomerCommand).Execute()

Create customer

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
    createStripeCustomerCommand := *openapiclient.NewCreateStripeCustomerCommand() // CreateStripeCustomerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsAPI.PaymentCreateCustomer(context.Background()).CreateStripeCustomerCommand(createStripeCustomerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentCreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStripeCustomerCommand** | [**CreateStripeCustomerCommand**](CreateStripeCustomerCommand.md) |  | 

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


## PaymentFinalPrice

> FinalPriceDto PaymentFinalPrice(ctx).FinalPriceCommand(finalPriceCommand).Execute()

Fetch final price

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
    finalPriceCommand := *openapiclient.NewFinalPriceCommand() // FinalPriceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.PaymentFinalPrice(context.Background()).FinalPriceCommand(finalPriceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentFinalPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentFinalPrice`: FinalPriceDto
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentFinalPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentFinalPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **finalPriceCommand** | [**FinalPriceCommand**](FinalPriceCommand.md) |  | 

### Return type

[**FinalPriceDto**](FinalPriceDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentGetStripeInvoices

> StripeInvoices PaymentGetStripeInvoices(ctx, subscriptionId).Execute()



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
    subscriptionId := "subscriptionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.PaymentGetStripeInvoices(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentGetStripeInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentGetStripeInvoices`: StripeInvoices
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentGetStripeInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentGetStripeInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StripeInvoices**](StripeInvoices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentPay

> map[string]interface{} PaymentPay(ctx).PayInvoiceCommand(payInvoiceCommand).Execute()

Pay invoice

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
    payInvoiceCommand := *openapiclient.NewPayInvoiceCommand() // PayInvoiceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.PaymentPay(context.Background()).PayInvoiceCommand(payInvoiceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentPay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentPay`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentPay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payInvoiceCommand** | [**PayInvoiceCommand**](PayInvoiceCommand.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentUpdateCard

> PaymentUpdateCard(ctx).ChangeCardCommand(changeCardCommand).Execute()

Update payment card

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
    changeCardCommand := *openapiclient.NewChangeCardCommand() // ChangeCardCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsAPI.PaymentUpdateCard(context.Background()).ChangeCardCommand(changeCardCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentUpdateCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentUpdateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeCardCommand** | [**ChangeCardCommand**](ChangeCardCommand.md) |  | 

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


## PaymentWebhook

> PaymentWebhook(ctx).Body(body).Execute()

Listen to payment webhook

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
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsAPI.PaymentWebhook(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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

