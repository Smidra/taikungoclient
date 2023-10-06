# \ProjectTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjecttemplateCreate**](ProjectTemplatesAPI.md#ProjecttemplateCreate) | **Post** /api/v1/projecttemplate/create | Create project from template
[**ProjecttemplateDelete**](ProjectTemplatesAPI.md#ProjecttemplateDelete) | **Delete** /api/v1/projecttemplate/{id} | Delete project template by Id
[**ProjecttemplateDropdown**](ProjectTemplatesAPI.md#ProjecttemplateDropdown) | **Get** /api/v1/projecttemplate/list | Retrieve project template by organization Id
[**ProjecttemplateList**](ProjectTemplatesAPI.md#ProjecttemplateList) | **Get** /api/v1/projecttemplate | Retrieve all project templates



## ProjecttemplateCreate

> ProjecttemplateCreate(ctx).CreateProjectFromTemplateCommand(createProjectFromTemplateCommand).Execute()

Create project from template

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
    createProjectFromTemplateCommand := *openapiclient.NewCreateProjectFromTemplateCommand() // CreateProjectFromTemplateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectTemplatesAPI.ProjecttemplateCreate(context.Background()).CreateProjectFromTemplateCommand(createProjectFromTemplateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.ProjecttemplateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjecttemplateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectFromTemplateCommand** | [**CreateProjectFromTemplateCommand**](CreateProjectFromTemplateCommand.md) |  | 

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


## ProjecttemplateDelete

> ProjecttemplateDelete(ctx, id).Execute()

Delete project template by Id

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
    r, err := apiClient.ProjectTemplatesAPI.ProjecttemplateDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.ProjecttemplateDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjecttemplateDeleteRequest struct via the builder pattern


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


## ProjecttemplateDropdown

> []ProjectTemplateDropdownListDto ProjecttemplateDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve project template by organization Id

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTemplatesAPI.ProjecttemplateDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.ProjecttemplateDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjecttemplateDropdown`: []ProjectTemplateDropdownListDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.ProjecttemplateDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjecttemplateDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]ProjectTemplateDropdownListDto**](ProjectTemplateDropdownListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjecttemplateList

> ProjectTemplateList ProjecttemplateList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve all project templates

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
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectTemplatesAPI.ProjecttemplateList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.ProjecttemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjecttemplateList`: ProjectTemplateList
    fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.ProjecttemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjecttemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**ProjectTemplateList**](ProjectTemplateList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

