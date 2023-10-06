# \AiManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiManagementDisable**](AiManagementAPI.md#AiManagementDisable) | **Post** /api/v1/ai-management/disable | Disable ai
[**AiManagementEnable**](AiManagementAPI.md#AiManagementEnable) | **Post** /api/v1/ai-management/enable | Enable ai



## AiManagementDisable

> AiManagementDisable(ctx).DisableAiCommand(disableAiCommand).Execute()

Disable ai

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
    disableAiCommand := *openapiclient.NewDisableAiCommand() // DisableAiCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AiManagementAPI.AiManagementDisable(context.Background()).DisableAiCommand(disableAiCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiManagementAPI.AiManagementDisable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiManagementDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableAiCommand** | [**DisableAiCommand**](DisableAiCommand.md) |  | 

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


## AiManagementEnable

> AiManagementEnable(ctx).EnableAiCommand(enableAiCommand).Execute()

Enable ai

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
    enableAiCommand := *openapiclient.NewEnableAiCommand() // EnableAiCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AiManagementAPI.AiManagementEnable(context.Background()).EnableAiCommand(enableAiCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiManagementAPI.AiManagementEnable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiManagementEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAiCommand** | [**EnableAiCommand**](EnableAiCommand.md) |  | 

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

