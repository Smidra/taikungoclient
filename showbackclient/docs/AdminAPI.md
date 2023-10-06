# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAddBalance**](AdminAPI.md#AdminAddBalance) | **Post** /api/v1/admin/organizations/add/balance | Add balance for organization
[**AdminBillingOperations**](AdminAPI.md#AdminBillingOperations) | **Post** /api/v1/admin/cloudcredentials/billing | Billing operations: enable/disable billing 
[**AdminCreateUser**](AdminAPI.md#AdminCreateUser) | **Post** /api/v1/admin/users/create | User creation for admin
[**AdminDeleteOrg**](AdminAPI.md#AdminDeleteOrg) | **Post** /api/v1/admin/organizations/delete | Delete organization
[**AdminKeycloakList**](AdminAPI.md#AdminKeycloakList) | **Get** /api/v1/admin/keycloak/list | Keycloak list for admin
[**AdminMakeCsm**](AdminAPI.md#AdminMakeCsm) | **Post** /api/v1/admin/users/make/csm | User csm update for admin
[**AdminMakeOwner**](AdminAPI.md#AdminMakeOwner) | **Post** /api/v1/admin/users/make/owner | User choose owner for admin
[**AdminOrganizations**](AdminAPI.md#AdminOrganizations) | **Get** /api/v1/admin/organizations/list |  Organizations for admin
[**AdminProjectList**](AdminAPI.md#AdminProjectList) | **Get** /api/v1/admin/projects/list | Projects for admin
[**AdminUpdateProject**](AdminAPI.md#AdminUpdateProject) | **Post** /api/v1/admin/projects/update/version | Projects update for admin
[**AdminUpdateProjectKube**](AdminAPI.md#AdminUpdateProjectKube) | **Post** /api/v1/admin/projects/update/kubeconfig | Projects update kube for admin
[**AdminUpdateUser**](AdminAPI.md#AdminUpdateUser) | **Post** /api/v1/admin/users/update/password | User password update for admin
[**AdminUpdateUserKube**](AdminAPI.md#AdminUpdateUserKube) | **Post** /api/v1/admin/projects/update/userkube | Projects update kube for admin
[**AdminUpdateUsers**](AdminAPI.md#AdminUpdateUsers) | **Post** /api/v1/admin/users/update/email | User email update for admin
[**AdminUsersList**](AdminAPI.md#AdminUsersList) | **Get** /api/v1/admin/users/list | Users for admin



## AdminAddBalance

> AdminAddBalance(ctx).AdminAddBalanceCommand(adminAddBalanceCommand).Execute()

Add balance for organization

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
    adminAddBalanceCommand := *openapiclient.NewAdminAddBalanceCommand() // AdminAddBalanceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminAddBalance(context.Background()).AdminAddBalanceCommand(adminAddBalanceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminAddBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAddBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminAddBalanceCommand** | [**AdminAddBalanceCommand**](AdminAddBalanceCommand.md) |  | 

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


## AdminBillingOperations

> AdminBillingOperations(ctx).AdminBillingOperationCommand(adminBillingOperationCommand).Execute()

Billing operations: enable/disable billing 

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
    adminBillingOperationCommand := *openapiclient.NewAdminBillingOperationCommand() // AdminBillingOperationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminBillingOperations(context.Background()).AdminBillingOperationCommand(adminBillingOperationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminBillingOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminBillingOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminBillingOperationCommand** | [**AdminBillingOperationCommand**](AdminBillingOperationCommand.md) |  | 

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


## AdminCreateUser

> AdminCreateUser(ctx).AdminUserCreateCommand(adminUserCreateCommand).Execute()

User creation for admin

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
    adminUserCreateCommand := *openapiclient.NewAdminUserCreateCommand() // AdminUserCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminCreateUser(context.Background()).AdminUserCreateCommand(adminUserCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUserCreateCommand** | [**AdminUserCreateCommand**](AdminUserCreateCommand.md) |  | 

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


## AdminDeleteOrg

> AdminDeleteOrg(ctx).AdminOrganizationsDeleteCommand(adminOrganizationsDeleteCommand).Execute()

Delete organization

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
    adminOrganizationsDeleteCommand := *openapiclient.NewAdminOrganizationsDeleteCommand() // AdminOrganizationsDeleteCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminDeleteOrg(context.Background()).AdminOrganizationsDeleteCommand(adminOrganizationsDeleteCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminOrganizationsDeleteCommand** | [**AdminOrganizationsDeleteCommand**](AdminOrganizationsDeleteCommand.md) |  | 

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


## AdminKeycloakList

> AdminProjectsList AdminKeycloakList(ctx).Limit(limit).Offset(offset).Execute()

Keycloak list for admin

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.AdminKeycloakList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminKeycloakList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminKeycloakList`: AdminProjectsList
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminKeycloakList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminKeycloakListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**AdminProjectsList**](AdminProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminMakeCsm

> AdminMakeCsm(ctx).MakeCsmCommand(makeCsmCommand).Execute()

User csm update for admin

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
    makeCsmCommand := *openapiclient.NewMakeCsmCommand() // MakeCsmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminMakeCsm(context.Background()).MakeCsmCommand(makeCsmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminMakeCsm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminMakeCsmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeCsmCommand** | [**MakeCsmCommand**](MakeCsmCommand.md) |  | 

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


## AdminMakeOwner

> AdminMakeOwner(ctx).MakeOwnerCommand(makeOwnerCommand).Execute()

User choose owner for admin

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
    makeOwnerCommand := *openapiclient.NewMakeOwnerCommand() // MakeOwnerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminMakeOwner(context.Background()).MakeOwnerCommand(makeOwnerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminMakeOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminMakeOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeOwnerCommand** | [**MakeOwnerCommand**](MakeOwnerCommand.md) |  | 

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


## AdminOrganizations

> AdminOrganizationsList AdminOrganizations(ctx).Limit(limit).Offset(offset).PartnerId(partnerId).Search(search).Execute()

 Organizations for admin

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
    partnerId := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.AdminOrganizations(context.Background()).Limit(limit).Offset(offset).PartnerId(partnerId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminOrganizations`: AdminOrganizationsList
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **partnerId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**AdminOrganizationsList**](AdminOrganizationsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminProjectList

> AdminProjectsList AdminProjectList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()

Projects for admin

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.AdminProjectList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminProjectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminProjectList`: AdminProjectsList
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminProjectList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminProjectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**AdminProjectsList**](AdminProjectsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUpdateProject

> AdminUpdateProject(ctx).AdminProjectUpdateCommand(adminProjectUpdateCommand).Execute()

Projects update for admin

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
    adminProjectUpdateCommand := *openapiclient.NewAdminProjectUpdateCommand() // AdminProjectUpdateCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminUpdateProject(context.Background()).AdminProjectUpdateCommand(adminProjectUpdateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminProjectUpdateCommand** | [**AdminProjectUpdateCommand**](AdminProjectUpdateCommand.md) |  | 

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


## AdminUpdateProjectKube

> AdminUpdateProjectKube(ctx).AdminUpdateProjectKubeConfigCommand(adminUpdateProjectKubeConfigCommand).Execute()

Projects update kube for admin

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
    adminUpdateProjectKubeConfigCommand := *openapiclient.NewAdminUpdateProjectKubeConfigCommand() // AdminUpdateProjectKubeConfigCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminUpdateProjectKube(context.Background()).AdminUpdateProjectKubeConfigCommand(adminUpdateProjectKubeConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateProjectKube``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateProjectKubeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUpdateProjectKubeConfigCommand** | [**AdminUpdateProjectKubeConfigCommand**](AdminUpdateProjectKubeConfigCommand.md) |  | 

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


## AdminUpdateUser

> AdminUpdateUser(ctx).AdminUsersUpdatePasswordCommand(adminUsersUpdatePasswordCommand).Execute()

User password update for admin

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
    adminUsersUpdatePasswordCommand := *openapiclient.NewAdminUsersUpdatePasswordCommand() // AdminUsersUpdatePasswordCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminUpdateUser(context.Background()).AdminUsersUpdatePasswordCommand(adminUsersUpdatePasswordCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUsersUpdatePasswordCommand** | [**AdminUsersUpdatePasswordCommand**](AdminUsersUpdatePasswordCommand.md) |  | 

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


## AdminUpdateUserKube

> AdminUpdateUserKube(ctx).AdminUpdateUserKubeConfigCommand(adminUpdateUserKubeConfigCommand).Execute()

Projects update kube for admin

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
    adminUpdateUserKubeConfigCommand := *openapiclient.NewAdminUpdateUserKubeConfigCommand() // AdminUpdateUserKubeConfigCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminUpdateUserKube(context.Background()).AdminUpdateUserKubeConfigCommand(adminUpdateUserKubeConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateUserKube``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateUserKubeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUpdateUserKubeConfigCommand** | [**AdminUpdateUserKubeConfigCommand**](AdminUpdateUserKubeConfigCommand.md) |  | 

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


## AdminUpdateUsers

> AdminUpdateUsers(ctx).AdminUsersUpdateEmailCommand(adminUsersUpdateEmailCommand).Execute()

User email update for admin

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
    adminUsersUpdateEmailCommand := *openapiclient.NewAdminUsersUpdateEmailCommand() // AdminUsersUpdateEmailCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.AdminUpdateUsers(context.Background()).AdminUsersUpdateEmailCommand(adminUsersUpdateEmailCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUsersUpdateEmailCommand** | [**AdminUsersUpdateEmailCommand**](AdminUsersUpdateEmailCommand.md) |  | 

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


## AdminUsersList

> AdminUsersList AdminUsersList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()

Users for admin

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.AdminUsersList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUsersList`: AdminUsersList
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**AdminUsersList**](AdminUsersList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

