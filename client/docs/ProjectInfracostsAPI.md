# \ProjectInfracostsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectinfracostsDelete**](ProjectInfracostsAPI.md#ProjectinfracostsDelete) | **Delete** /api/v1/projectinfracosts/{projectId} | Delete the project infracost
[**ProjectinfracostsDetails**](ProjectInfracostsAPI.md#ProjectinfracostsDetails) | **Get** /api/v1/projectinfracosts/{projectId} | Project Infracost details
[**ProjectinfracostsUpsert**](ProjectInfracostsAPI.md#ProjectinfracostsUpsert) | **Post** /api/v1/projectinfracosts/upsert/{projectId} | Upsert project infracost by ProjectId



## ProjectinfracostsDelete

> ProjectinfracostsDelete(ctx, projectId).Execute()

Delete the project infracost

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectInfracostsAPI.ProjectinfracostsDelete(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectInfracostsAPI.ProjectinfracostsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectinfracostsDeleteRequest struct via the builder pattern


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


## ProjectinfracostsDetails

> EstimatedInfracost ProjectinfracostsDetails(ctx, projectId).Execute()

Project Infracost details

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectInfracostsAPI.ProjectinfracostsDetails(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectInfracostsAPI.ProjectinfracostsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectinfracostsDetails`: EstimatedInfracost
    fmt.Fprintf(os.Stdout, "Response from `ProjectInfracostsAPI.ProjectinfracostsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectinfracostsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EstimatedInfracost**](EstimatedInfracost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectinfracostsUpsert

> ProjectinfracostsUpsert(ctx, projectId).ProjectInfracostUpsertDto(projectInfracostUpsertDto).Execute()

Upsert project infracost by ProjectId

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
    projectId := int32(56) // int32 | 
    projectInfracostUpsertDto := *openapiclient.NewProjectInfracostUpsertDto() // ProjectInfracostUpsertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectInfracostsAPI.ProjectinfracostsUpsert(context.Background(), projectId).ProjectInfracostUpsertDto(projectInfracostUpsertDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectInfracostsAPI.ProjectinfracostsUpsert``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectinfracostsUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectInfracostUpsertDto** | [**ProjectInfracostUpsertDto**](ProjectInfracostUpsertDto.md) |  | 

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

