# \S3CredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3credentialsCreate**](S3CredentialsAPI.md#S3credentialsCreate) | **Post** /api/v1/s3credentials | Add s3 credential
[**S3credentialsDelete**](S3CredentialsAPI.md#S3credentialsDelete) | **Delete** /api/v1/s3credentials/{id} | Delete s3 credential
[**S3credentialsDropdown**](S3CredentialsAPI.md#S3credentialsDropdown) | **Get** /api/v1/s3credentials | Retrieve all S3 credentials for organization
[**S3credentialsList**](S3CredentialsAPI.md#S3credentialsList) | **Get** /api/v1/s3credentials/list | Retrieve all S3 credentials
[**S3credentialsLockManagement**](S3CredentialsAPI.md#S3credentialsLockManagement) | **Post** /api/v1/s3credentials/lockmanager | Lock/unlock s3 credential
[**S3credentialsMakeDeafult**](S3CredentialsAPI.md#S3credentialsMakeDeafult) | **Post** /api/v1/s3credentials/makedefault | Make default s3 credential
[**S3credentialsUpdate**](S3CredentialsAPI.md#S3credentialsUpdate) | **Put** /api/v1/s3credentials | Update s3 credential



## S3credentialsCreate

> ApiResponse S3credentialsCreate(ctx).BackupCredentialsCreateCommand(backupCredentialsCreateCommand).Execute()

Add s3 credential

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
    backupCredentialsCreateCommand := *openapiclient.NewBackupCredentialsCreateCommand() // BackupCredentialsCreateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3CredentialsAPI.S3credentialsCreate(context.Background()).BackupCredentialsCreateCommand(backupCredentialsCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsAPI.S3credentialsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3credentialsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsAPI.S3credentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3credentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupCredentialsCreateCommand** | [**BackupCredentialsCreateCommand**](BackupCredentialsCreateCommand.md) |  | 

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


## S3credentialsDelete

> S3credentialsDelete(ctx, id).Execute()

Delete s3 credential

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3CredentialsAPI.S3credentialsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsAPI.S3credentialsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiS3credentialsDeleteRequest struct via the builder pattern


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


## S3credentialsDropdown

> []BackupCredentialsForOrganizationEntity S3credentialsDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all S3 credentials for organization

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
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3CredentialsAPI.S3credentialsDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsAPI.S3credentialsDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3credentialsDropdown`: []BackupCredentialsForOrganizationEntity
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsAPI.S3credentialsDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3credentialsDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]BackupCredentialsForOrganizationEntity**](BackupCredentialsForOrganizationEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3credentialsList

> BackupCredentials S3credentialsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieve all S3 credentials

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3CredentialsAPI.S3credentialsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsAPI.S3credentialsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3credentialsList`: BackupCredentials
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsAPI.S3credentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3credentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**BackupCredentials**](BackupCredentials.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3credentialsLockManagement

> S3credentialsLockManagement(ctx).BackupLockManagerCommand(backupLockManagerCommand).Execute()

Lock/unlock s3 credential

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
    backupLockManagerCommand := *openapiclient.NewBackupLockManagerCommand() // BackupLockManagerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3CredentialsAPI.S3credentialsLockManagement(context.Background()).BackupLockManagerCommand(backupLockManagerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsAPI.S3credentialsLockManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3credentialsLockManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupLockManagerCommand** | [**BackupLockManagerCommand**](BackupLockManagerCommand.md) |  | 

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


## S3credentialsMakeDeafult

> S3credentialsMakeDeafult(ctx).BackupMakeDefaultCommand(backupMakeDefaultCommand).Execute()

Make default s3 credential

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
    backupMakeDefaultCommand := *openapiclient.NewBackupMakeDefaultCommand() // BackupMakeDefaultCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3CredentialsAPI.S3credentialsMakeDeafult(context.Background()).BackupMakeDefaultCommand(backupMakeDefaultCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsAPI.S3credentialsMakeDeafult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3credentialsMakeDeafultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupMakeDefaultCommand** | [**BackupMakeDefaultCommand**](BackupMakeDefaultCommand.md) |  | 

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


## S3credentialsUpdate

> S3credentialsUpdate(ctx).BackupCredentialsUpdateCommand(backupCredentialsUpdateCommand).Execute()

Update s3 credential

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
    backupCredentialsUpdateCommand := *openapiclient.NewBackupCredentialsUpdateCommand() // BackupCredentialsUpdateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3CredentialsAPI.S3credentialsUpdate(context.Background()).BackupCredentialsUpdateCommand(backupCredentialsUpdateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsAPI.S3credentialsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3credentialsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupCredentialsUpdateCommand** | [**BackupCredentialsUpdateCommand**](BackupCredentialsUpdateCommand.md) |  | 

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

