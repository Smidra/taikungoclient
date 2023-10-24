# \BackupPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupByName**](BackupPolicyAPI.md#BackupByName) | **Get** /api/v1/backup/{projectId}/{name} | 
[**BackupClearProject**](BackupPolicyAPI.md#BackupClearProject) | **Post** /api/v1/backup/clear/project | Delete unfinished backup for project
[**BackupCreate**](BackupPolicyAPI.md#BackupCreate) | **Post** /api/v1/backup/create | Add backup policy
[**BackupDeleteBackup**](BackupPolicyAPI.md#BackupDeleteBackup) | **Post** /api/v1/backup/delete/backup | Remove policy backup
[**BackupDeleteBackupLocation**](BackupPolicyAPI.md#BackupDeleteBackupLocation) | **Post** /api/v1/backup/delete/location | Remove backup location from project
[**BackupDeleteRestore**](BackupPolicyAPI.md#BackupDeleteRestore) | **Post** /api/v1/backup/delete/restore | Remove policy restore
[**BackupDeleteSchedule**](BackupPolicyAPI.md#BackupDeleteSchedule) | **Post** /api/v1/backup/delete/schedule | Remove policy schedule
[**BackupDescribeBackup**](BackupPolicyAPI.md#BackupDescribeBackup) | **Get** /api/v1/backup/describe/backup/{projectId}/{name} | 
[**BackupDescribeRestore**](BackupPolicyAPI.md#BackupDescribeRestore) | **Get** /api/v1/backup/describe/restore/{projectId}/{name} | 
[**BackupDescribeSchedule**](BackupPolicyAPI.md#BackupDescribeSchedule) | **Get** /api/v1/backup/describe/schedule/{projectId}/{name} | 
[**BackupDisableBackup**](BackupPolicyAPI.md#BackupDisableBackup) | **Post** /api/v1/backup/disablebackup | Disable backup by the projectId
[**BackupEnableBackup**](BackupPolicyAPI.md#BackupEnableBackup) | **Post** /api/v1/backup/enablebackup | Enable backup by the projectId
[**BackupImportBackupStorage**](BackupPolicyAPI.md#BackupImportBackupStorage) | **Post** /api/v1/backup/location | Import backup storage from source project to target project
[**BackupListAllBackupStorages**](BackupPolicyAPI.md#BackupListAllBackupStorages) | **Get** /api/v1/backup/location/{projectId} | List all backup locations
[**BackupListAllBackups**](BackupPolicyAPI.md#BackupListAllBackups) | **Get** /api/v1/backup/backups/{projectId} | List all backups
[**BackupListAllDeleteBackupRequests**](BackupPolicyAPI.md#BackupListAllDeleteBackupRequests) | **Get** /api/v1/backup/delete-requests/{projectId} | List all delete backup requests
[**BackupListAllRestores**](BackupPolicyAPI.md#BackupListAllRestores) | **Get** /api/v1/backup/restores/{projectId} | List all restores
[**BackupListAllSchedules**](BackupPolicyAPI.md#BackupListAllSchedules) | **Get** /api/v1/backup/schedules/{projectId} | List all schedules
[**BackupRestoreBackup**](BackupPolicyAPI.md#BackupRestoreBackup) | **Post** /api/v1/backup/restore | Restore backup



## BackupByName

> BackupDto BackupByName(ctx, projectId, name).Execute()



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
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupByName(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupByName`: BackupDto
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BackupDto**](BackupDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupClearProject

> BackupClearProject(ctx).ClearProjectBackupCommand(clearProjectBackupCommand).Execute()

Delete unfinished backup for project

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
    clearProjectBackupCommand := *openapiclient.NewClearProjectBackupCommand() // ClearProjectBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupClearProject(context.Background()).ClearProjectBackupCommand(clearProjectBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupClearProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupClearProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clearProjectBackupCommand** | [**ClearProjectBackupCommand**](ClearProjectBackupCommand.md) |  | 

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


## BackupCreate

> BackupCreate(ctx).CreateBackupPolicyCommand(createBackupPolicyCommand).Execute()

Add backup policy

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
    createBackupPolicyCommand := *openapiclient.NewCreateBackupPolicyCommand() // CreateBackupPolicyCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupCreate(context.Background()).CreateBackupPolicyCommand(createBackupPolicyCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBackupPolicyCommand** | [**CreateBackupPolicyCommand**](CreateBackupPolicyCommand.md) |  | 

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


## BackupDeleteBackup

> BackupDeleteBackup(ctx).DeleteBackupCommand(deleteBackupCommand).Execute()

Remove policy backup

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
    deleteBackupCommand := *openapiclient.NewDeleteBackupCommand() // DeleteBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupDeleteBackup(context.Background()).DeleteBackupCommand(deleteBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupCommand** | [**DeleteBackupCommand**](DeleteBackupCommand.md) |  | 

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


## BackupDeleteBackupLocation

> BackupDeleteBackupLocation(ctx).DeleteBackupStorageLocationCommand(deleteBackupStorageLocationCommand).Execute()

Remove backup location from project

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
    deleteBackupStorageLocationCommand := *openapiclient.NewDeleteBackupStorageLocationCommand() // DeleteBackupStorageLocationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupDeleteBackupLocation(context.Background()).DeleteBackupStorageLocationCommand(deleteBackupStorageLocationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteBackupLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteBackupLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupStorageLocationCommand** | [**DeleteBackupStorageLocationCommand**](DeleteBackupStorageLocationCommand.md) |  | 

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


## BackupDeleteRestore

> BackupDeleteRestore(ctx).DeleteRestoreCommand(deleteRestoreCommand).Execute()

Remove policy restore

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
    deleteRestoreCommand := *openapiclient.NewDeleteRestoreCommand() // DeleteRestoreCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupDeleteRestore(context.Background()).DeleteRestoreCommand(deleteRestoreCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRestoreCommand** | [**DeleteRestoreCommand**](DeleteRestoreCommand.md) |  | 

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


## BackupDeleteSchedule

> BackupDeleteSchedule(ctx).DeleteScheduleCommand(deleteScheduleCommand).Execute()

Remove policy schedule

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
    deleteScheduleCommand := *openapiclient.NewDeleteScheduleCommand() // DeleteScheduleCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupDeleteSchedule(context.Background()).DeleteScheduleCommand(deleteScheduleCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteScheduleCommand** | [**DeleteScheduleCommand**](DeleteScheduleCommand.md) |  | 

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


## BackupDescribeBackup

> string BackupDescribeBackup(ctx, projectId, name).Execute()



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
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupDescribeBackup(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDescribeBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupDescribeBackup`: string
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupDescribeBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDescribeRestore

> string BackupDescribeRestore(ctx, projectId, name).Execute()



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
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupDescribeRestore(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDescribeRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupDescribeRestore`: string
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupDescribeRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDescribeSchedule

> string BackupDescribeSchedule(ctx, projectId, name).Execute()



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
    projectId := int32(56) // int32 | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupDescribeSchedule(context.Background(), projectId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDescribeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupDescribeSchedule`: string
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupDescribeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupDisableBackup

> BackupDisableBackup(ctx).DisableBackupCommand(disableBackupCommand).Execute()

Disable backup by the projectId

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
    disableBackupCommand := *openapiclient.NewDisableBackupCommand() // DisableBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupDisableBackup(context.Background()).DisableBackupCommand(disableBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDisableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDisableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableBackupCommand** | [**DisableBackupCommand**](DisableBackupCommand.md) |  | 

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


## BackupEnableBackup

> BackupEnableBackup(ctx).EnableBackupCommand(enableBackupCommand).Execute()

Enable backup by the projectId

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
    enableBackupCommand := *openapiclient.NewEnableBackupCommand() // EnableBackupCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupEnableBackup(context.Background()).EnableBackupCommand(enableBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupEnableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupEnableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableBackupCommand** | [**EnableBackupCommand**](EnableBackupCommand.md) |  | 

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


## BackupImportBackupStorage

> BackupImportBackupStorage(ctx).ImportBackupStorageLocationCommand(importBackupStorageLocationCommand).Execute()

Import backup storage from source project to target project

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
    importBackupStorageLocationCommand := *openapiclient.NewImportBackupStorageLocationCommand() // ImportBackupStorageLocationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupImportBackupStorage(context.Background()).ImportBackupStorageLocationCommand(importBackupStorageLocationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupImportBackupStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupImportBackupStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importBackupStorageLocationCommand** | [**ImportBackupStorageLocationCommand**](ImportBackupStorageLocationCommand.md) |  | 

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


## BackupListAllBackupStorages

> ListAllBackupStorageLocations BackupListAllBackupStorages(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

List all backup locations

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupListAllBackupStorages(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllBackupStorages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllBackupStorages`: ListAllBackupStorageLocations
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllBackupStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllBackupStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListAllBackupStorageLocations**](ListAllBackupStorageLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllBackups

> ListAllBackups BackupListAllBackups(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

List all backups

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupListAllBackups(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllBackups`: ListAllBackups
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListAllBackups**](ListAllBackups.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllDeleteBackupRequests

> ListAllDeleteBackupRequests BackupListAllDeleteBackupRequests(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

List all delete backup requests

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupListAllDeleteBackupRequests(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllDeleteBackupRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllDeleteBackupRequests`: ListAllDeleteBackupRequests
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllDeleteBackupRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllDeleteBackupRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListAllDeleteBackupRequests**](ListAllDeleteBackupRequests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllRestores

> ListAllRestores BackupListAllRestores(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

List all restores

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupListAllRestores(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllRestores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllRestores`: ListAllRestores
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllRestores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListAllRestores**](ListAllRestores.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllSchedules

> ListAllSchedules BackupListAllSchedules(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

List all schedules

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
    projectId := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupPolicyAPI.BackupListAllSchedules(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupListAllSchedules`: ListAllSchedules
    fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListAllSchedules**](ListAllSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreBackup

> BackupRestoreBackup(ctx).RestoreBackupCommand(restoreBackupCommand).Execute()

Restore backup

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
    restoreBackupCommand := *openapiclient.NewRestoreBackupCommand() // RestoreBackupCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackupPolicyAPI.BackupRestoreBackup(context.Background()).RestoreBackupCommand(restoreBackupCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupRestoreBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreBackupCommand** | [**RestoreBackupCommand**](RestoreBackupCommand.md) |  | 

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

