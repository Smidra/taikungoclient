# \PrometheusRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrometheusrulesBindOrganizations**](PrometheusRulesAPI.md#PrometheusrulesBindOrganizations) | **Post** /api/v1/prometheusrules/bind/organizations | Bind organizations to prometheus rule
[**PrometheusrulesCreate**](PrometheusRulesAPI.md#PrometheusrulesCreate) | **Post** /api/v1/prometheusrules | Add prometheus rule
[**PrometheusrulesDelete**](PrometheusRulesAPI.md#PrometheusrulesDelete) | **Delete** /api/v1/prometheusrules/{id} | Remove prometheus rule
[**PrometheusrulesDetails**](PrometheusRulesAPI.md#PrometheusrulesDetails) | **Get** /api/v1/prometheusrules/details | Retrieve all prometheus rules with detailed info
[**PrometheusrulesList**](PrometheusRulesAPI.md#PrometheusrulesList) | **Get** /api/v1/prometheusrules | Retrieve a list of prometheus rules
[**PrometheusrulesUpdate**](PrometheusRulesAPI.md#PrometheusrulesUpdate) | **Put** /api/v1/prometheusrules/edit/{id} | Edit prometheus rule



## PrometheusrulesBindOrganizations

> PrometheusrulesBindOrganizations(ctx).BindPrometheusOrganizationsCommand(bindPrometheusOrganizationsCommand).Execute()

Bind organizations to prometheus rule

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
    bindPrometheusOrganizationsCommand := *openapiclient.NewBindPrometheusOrganizationsCommand() // BindPrometheusOrganizationsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PrometheusRulesAPI.PrometheusrulesBindOrganizations(context.Background()).BindPrometheusOrganizationsCommand(bindPrometheusOrganizationsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusRulesAPI.PrometheusrulesBindOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusrulesBindOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindPrometheusOrganizationsCommand** | [**BindPrometheusOrganizationsCommand**](BindPrometheusOrganizationsCommand.md) |  | 

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


## PrometheusrulesCreate

> ApiResponse PrometheusrulesCreate(ctx).RuleCreateCommand(ruleCreateCommand).Execute()

Add prometheus rule

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
    ruleCreateCommand := *openapiclient.NewRuleCreateCommand() // RuleCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrometheusRulesAPI.PrometheusrulesCreate(context.Background()).RuleCreateCommand(ruleCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusRulesAPI.PrometheusrulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrometheusrulesCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PrometheusRulesAPI.PrometheusrulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusrulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleCreateCommand** | [**RuleCreateCommand**](RuleCreateCommand.md) |  | 

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


## PrometheusrulesDelete

> PrometheusrulesDelete(ctx, id).Execute()

Remove prometheus rule

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PrometheusRulesAPI.PrometheusrulesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusRulesAPI.PrometheusrulesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPrometheusrulesDeleteRequest struct via the builder pattern


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


## PrometheusrulesDetails

> []CommonDropdownIsBoundDto PrometheusrulesDetails(ctx).OrganizationId(organizationId).Execute()

Retrieve all prometheus rules with detailed info

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
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrometheusRulesAPI.PrometheusrulesDetails(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusRulesAPI.PrometheusrulesDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrometheusrulesDetails`: []CommonDropdownIsBoundDto
    fmt.Fprintf(os.Stdout, "Response from `PrometheusRulesAPI.PrometheusrulesDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusrulesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**[]CommonDropdownIsBoundDto**](CommonDropdownIsBoundDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrometheusrulesList

> PrometheusRulesList PrometheusrulesList(ctx).Limit(limit).Offset(offset).PartnerId(partnerId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve a list of prometheus rules

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    partnerId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrometheusRulesAPI.PrometheusrulesList(context.Background()).Limit(limit).Offset(offset).PartnerId(partnerId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusRulesAPI.PrometheusrulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrometheusrulesList`: PrometheusRulesList
    fmt.Fprintf(os.Stdout, "Response from `PrometheusRulesAPI.PrometheusrulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusrulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **partnerId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**PrometheusRulesList**](PrometheusRulesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrometheusrulesUpdate

> PrometheusrulesUpdate(ctx, id).RuleForUpdateDto(ruleForUpdateDto).Execute()

Edit prometheus rule

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
    id := int32(56) // int32 | 
    ruleForUpdateDto := *openapiclient.NewRuleForUpdateDto() // RuleForUpdateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PrometheusRulesAPI.PrometheusrulesUpdate(context.Background(), id).RuleForUpdateDto(ruleForUpdateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusRulesAPI.PrometheusrulesUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPrometheusrulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleForUpdateDto** | [**RuleForUpdateDto**](RuleForUpdateDto.md) |  | 

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

