# \ProjectAppParamsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectappparamEdit**](ProjectAppParamsAPI.md#ProjectappparamEdit) | **Put** /api/v1/projectappparam/edit/{projectAppId} | Edit project app params



## ProjectappparamEdit

> ProjectappparamEdit(ctx, projectAppId).EditProjectAppParamsDto(editProjectAppParamsDto).Execute()

Edit project app params

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
    projectAppId := int32(56) // int32 | 
    editProjectAppParamsDto := []openapiclient.EditProjectAppParamsDto{*openapiclient.NewEditProjectAppParamsDto()} // []EditProjectAppParamsDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectAppParamsAPI.ProjectappparamEdit(context.Background(), projectAppId).EditProjectAppParamsDto(editProjectAppParamsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAppParamsAPI.ProjectappparamEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectappparamEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editProjectAppParamsDto** | [**[]EditProjectAppParamsDto**](EditProjectAppParamsDto.md) |  | 

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

