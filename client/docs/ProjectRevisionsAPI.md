# \ProjectRevisionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectrevisionsEdit**](ProjectRevisionsAPI.md#ProjectrevisionsEdit) | **Put** /api/v1/projectrevisions/edit/{projectId} | Update project revision by ProjectId for poller



## ProjectrevisionsEdit

> ProjectrevisionsEdit(ctx, projectId).ProjectRevisionUpdateDto(projectRevisionUpdateDto).Execute()

Update project revision by ProjectId for poller

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
    projectId := int32(56) // int32 | 
    projectRevisionUpdateDto := *openapiclient.NewProjectRevisionUpdateDto() // ProjectRevisionUpdateDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectRevisionsAPI.ProjectrevisionsEdit(context.Background(), projectId).ProjectRevisionUpdateDto(projectRevisionUpdateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectRevisionsAPI.ProjectrevisionsEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectrevisionsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectRevisionUpdateDto** | [**ProjectRevisionUpdateDto**](ProjectRevisionUpdateDto.md) |  | 

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

