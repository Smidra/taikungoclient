# \OrganizationSubscriptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsubcriptionsList**](OrganizationSubscriptionsAPI.md#OrganizationsubcriptionsList) | **Get** /api/v1/organizationsubcriptions | Retrieve all organization subscriptions
[**OrganizationsubcriptionsUpdate**](OrganizationSubscriptionsAPI.md#OrganizationsubcriptionsUpdate) | **Post** /api/v1/organizationsubcriptions/update | Update subscription



## OrganizationsubcriptionsList

> []OrganizationSubscriptionDto OrganizationsubcriptionsList(ctx).Execute()

Retrieve all organization subscriptions

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationSubscriptionsAPI.OrganizationsubcriptionsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSubscriptionsAPI.OrganizationsubcriptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsubcriptionsList`: []OrganizationSubscriptionDto
    fmt.Fprintf(os.Stdout, "Response from `OrganizationSubscriptionsAPI.OrganizationsubcriptionsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsubcriptionsListRequest struct via the builder pattern


### Return type

[**[]OrganizationSubscriptionDto**](OrganizationSubscriptionDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsubcriptionsUpdate

> LeaveTaikunDto OrganizationsubcriptionsUpdate(ctx).UpdateOrganizationSubscriptionCommand(updateOrganizationSubscriptionCommand).Execute()

Update subscription

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
    updateOrganizationSubscriptionCommand := *openapiclient.NewUpdateOrganizationSubscriptionCommand() // UpdateOrganizationSubscriptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationSubscriptionsAPI.OrganizationsubcriptionsUpdate(context.Background()).UpdateOrganizationSubscriptionCommand(updateOrganizationSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSubscriptionsAPI.OrganizationsubcriptionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsubcriptionsUpdate`: LeaveTaikunDto
    fmt.Fprintf(os.Stdout, "Response from `OrganizationSubscriptionsAPI.OrganizationsubcriptionsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsubcriptionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationSubscriptionCommand** | [**UpdateOrganizationSubscriptionCommand**](UpdateOrganizationSubscriptionCommand.md) |  | 

### Return type

[**LeaveTaikunDto**](LeaveTaikunDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

