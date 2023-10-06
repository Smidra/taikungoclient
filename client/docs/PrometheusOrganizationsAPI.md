# \PrometheusOrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrometheusorganizationsBindRules**](PrometheusOrganizationsAPI.md#PrometheusorganizationsBindRules) | **Post** /api/v1/prometheusorganizations/bind/rules | Bind rules to organizations



## PrometheusorganizationsBindRules

> PrometheusorganizationsBindRules(ctx).BindRulesCommand(bindRulesCommand).Execute()

Bind rules to organizations

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
    bindRulesCommand := *openapiclient.NewBindRulesCommand() // BindRulesCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PrometheusOrganizationsAPI.PrometheusorganizationsBindRules(context.Background()).BindRulesCommand(bindRulesCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusOrganizationsAPI.PrometheusorganizationsBindRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusorganizationsBindRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindRulesCommand** | [**BindRulesCommand**](BindRulesCommand.md) |  | 

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

