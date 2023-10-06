# \ProjectGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectgroupsBind**](ProjectGroupsAPI.md#ProjectgroupsBind) | **Post** /api/v1/projectgroups/bind | Bind User groups
[**ProjectgroupsCreate**](ProjectGroupsAPI.md#ProjectgroupsCreate) | **Post** /api/v1/projectgroups/create | Add Project groups
[**ProjectgroupsDelete**](ProjectGroupsAPI.md#ProjectgroupsDelete) | **Delete** /api/v1/projectgroups | Remove Project group(s)
[**ProjectgroupsEdit**](ProjectGroupsAPI.md#ProjectgroupsEdit) | **Put** /api/v1/projectgroups/update/{projectGroupId} | Update project groups
[**ProjectgroupsList**](ProjectGroupsAPI.md#ProjectgroupsList) | **Get** /api/v1/projectgroups/list | Retrieve list of Project groups
[**ProjectgroupsListByUserId**](ProjectGroupsAPI.md#ProjectgroupsListByUserId) | **Get** /api/v1/projectgroups/list/{userGroupId} | Retrieve list of Project groups by user group id for dropdown
[**ProjectgroupsProjectList**](ProjectGroupsAPI.md#ProjectgroupsProjectList) | **Get** /api/v1/projectgroups/{projectGroupId}/projects | Retrieve list of projects by project group id



## ProjectgroupsBind

> ProjectgroupsBind(ctx).BindUserGroupsToProjectGroupCommand(bindUserGroupsToProjectGroupCommand).Execute()

Bind User groups

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
    bindUserGroupsToProjectGroupCommand := *openapiclient.NewBindUserGroupsToProjectGroupCommand() // BindUserGroupsToProjectGroupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectGroupsAPI.ProjectgroupsBind(context.Background()).BindUserGroupsToProjectGroupCommand(bindUserGroupsToProjectGroupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectGroupsAPI.ProjectgroupsBind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectgroupsBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindUserGroupsToProjectGroupCommand** | [**BindUserGroupsToProjectGroupCommand**](BindUserGroupsToProjectGroupCommand.md) |  | 

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


## ProjectgroupsCreate

> ApiResponse ProjectgroupsCreate(ctx).CreateProjectGroupCommand(createProjectGroupCommand).Execute()

Add Project groups

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
    createProjectGroupCommand := *openapiclient.NewCreateProjectGroupCommand() // CreateProjectGroupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectGroupsAPI.ProjectgroupsCreate(context.Background()).CreateProjectGroupCommand(createProjectGroupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectGroupsAPI.ProjectgroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectgroupsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectGroupsAPI.ProjectgroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectgroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectGroupCommand** | [**CreateProjectGroupCommand**](CreateProjectGroupCommand.md) |  | 

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


## ProjectgroupsDelete

> ProjectgroupsDelete(ctx).RequestBody(requestBody).Execute()

Remove Project group(s)

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
    requestBody := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectGroupsAPI.ProjectgroupsDelete(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectGroupsAPI.ProjectgroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectgroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** |  | 

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


## ProjectgroupsEdit

> ProjectgroupsEdit(ctx, projectGroupId).UpdateProjectGroupDto(updateProjectGroupDto).Execute()

Update project groups

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
    projectGroupId := int32(56) // int32 | 
    updateProjectGroupDto := *openapiclient.NewUpdateProjectGroupDto() // UpdateProjectGroupDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectGroupsAPI.ProjectgroupsEdit(context.Background(), projectGroupId).UpdateProjectGroupDto(updateProjectGroupDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectGroupsAPI.ProjectgroupsEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectgroupsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProjectGroupDto** | [**UpdateProjectGroupDto**](UpdateProjectGroupDto.md) |  | 

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


## ProjectgroupsList

> ProjectGroupList ProjectgroupsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of Project groups

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectGroupsAPI.ProjectgroupsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectGroupsAPI.ProjectgroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectgroupsList`: ProjectGroupList
    fmt.Fprintf(os.Stdout, "Response from `ProjectGroupsAPI.ProjectgroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectgroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**ProjectGroupList**](ProjectGroupList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectgroupsListByUserId

> []CommonDropdownDto ProjectgroupsListByUserId(ctx, userGroupId).Execute()

Retrieve list of Project groups by user group id for dropdown

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
    userGroupId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectGroupsAPI.ProjectgroupsListByUserId(context.Background(), userGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectGroupsAPI.ProjectgroupsListByUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectgroupsListByUserId`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectGroupsAPI.ProjectgroupsListByUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectgroupsListByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectgroupsProjectList

> []ProjectListForProjectGroupDto ProjectgroupsProjectList(ctx, projectGroupId).OrganizationId(organizationId).Execute()

Retrieve list of projects by project group id

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
    projectGroupId := int32(56) // int32 | 
    organizationId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectGroupsAPI.ProjectgroupsProjectList(context.Background(), projectGroupId).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectGroupsAPI.ProjectgroupsProjectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectgroupsProjectList`: []ProjectListForProjectGroupDto
    fmt.Fprintf(os.Stdout, "Response from `ProjectGroupsAPI.ProjectgroupsProjectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectgroupsProjectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 

### Return type

[**[]ProjectListForProjectGroupDto**](ProjectListForProjectGroupDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

