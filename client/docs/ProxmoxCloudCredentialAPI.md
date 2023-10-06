# \ProxmoxCloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProxmoxBridgeList**](ProxmoxCloudCredentialAPI.md#ProxmoxBridgeList) | **Post** /api/v1/proxmox/bridge-list | Fetch proxmox bridge list
[**ProxmoxCreate**](ProxmoxCloudCredentialAPI.md#ProxmoxCreate) | **Post** /api/v1/proxmox/create | Add Proxmox credentials
[**ProxmoxHypervisorList**](ProxmoxCloudCredentialAPI.md#ProxmoxHypervisorList) | **Post** /api/v1/proxmox/hypervisor-list | Fetch proxmox hypervisor list
[**ProxmoxList**](ProxmoxCloudCredentialAPI.md#ProxmoxList) | **Get** /api/v1/proxmox/list | Retrieve list of proxmox cloud credentials
[**ProxmoxStorageList**](ProxmoxCloudCredentialAPI.md#ProxmoxStorageList) | **Post** /api/v1/proxmox/storage-list | Fetch proxmox storage list
[**ProxmoxUpdate**](ProxmoxCloudCredentialAPI.md#ProxmoxUpdate) | **Post** /api/v1/proxmox/update | Update proxmox credentials
[**ProxmoxUpdateHypervisors**](ProxmoxCloudCredentialAPI.md#ProxmoxUpdateHypervisors) | **Post** /api/v1/proxmox/update/hypervisors | Update proxmox credentials
[**ProxmoxUpdateIpAddresses**](ProxmoxCloudCredentialAPI.md#ProxmoxUpdateIpAddresses) | **Post** /api/v1/proxmox/update/ip-addresses | Update proxmox network used ip addresses
[**ProxmoxVmTemplateList**](ProxmoxCloudCredentialAPI.md#ProxmoxVmTemplateList) | **Post** /api/v1/proxmox/vm-template-list | Fetch proxmox vm template list



## ProxmoxBridgeList

> []string ProxmoxBridgeList(ctx).BridgeListCommand(bridgeListCommand).Execute()

Fetch proxmox bridge list

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
    bridgeListCommand := *openapiclient.NewBridgeListCommand() // BridgeListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxBridgeList(context.Background()).BridgeListCommand(bridgeListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxBridgeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxmoxBridgeList`: []string
    fmt.Fprintf(os.Stdout, "Response from `ProxmoxCloudCredentialAPI.ProxmoxBridgeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxBridgeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridgeListCommand** | [**BridgeListCommand**](BridgeListCommand.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxmoxCreate

> ApiResponse ProxmoxCreate(ctx).CreateProxmoxCommand(createProxmoxCommand).Execute()

Add Proxmox credentials

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
    createProxmoxCommand := *openapiclient.NewCreateProxmoxCommand() // CreateProxmoxCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxCreate(context.Background()).CreateProxmoxCommand(createProxmoxCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxmoxCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ProxmoxCloudCredentialAPI.ProxmoxCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProxmoxCommand** | [**CreateProxmoxCommand**](CreateProxmoxCommand.md) |  | 

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


## ProxmoxHypervisorList

> []ProxmoxHypervisorDto ProxmoxHypervisorList(ctx).HypervisorListCommand(hypervisorListCommand).Execute()

Fetch proxmox hypervisor list

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
    hypervisorListCommand := *openapiclient.NewHypervisorListCommand() // HypervisorListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxHypervisorList(context.Background()).HypervisorListCommand(hypervisorListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxHypervisorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxmoxHypervisorList`: []ProxmoxHypervisorDto
    fmt.Fprintf(os.Stdout, "Response from `ProxmoxCloudCredentialAPI.ProxmoxHypervisorList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxHypervisorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hypervisorListCommand** | [**HypervisorListCommand**](HypervisorListCommand.md) |  | 

### Return type

[**[]ProxmoxHypervisorDto**](ProxmoxHypervisorDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxmoxList

> ProxmoxList ProxmoxList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of proxmox cloud credentials

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
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxmoxList`: ProxmoxList
    fmt.Fprintf(os.Stdout, "Response from `ProxmoxCloudCredentialAPI.ProxmoxList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxListRequest struct via the builder pattern


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

[**ProxmoxList**](ProxmoxList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxmoxStorageList

> []string ProxmoxStorageList(ctx).StorageListCommand(storageListCommand).Execute()

Fetch proxmox storage list

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
    storageListCommand := *openapiclient.NewStorageListCommand() // StorageListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxStorageList(context.Background()).StorageListCommand(storageListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxStorageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxmoxStorageList`: []string
    fmt.Fprintf(os.Stdout, "Response from `ProxmoxCloudCredentialAPI.ProxmoxStorageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxStorageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageListCommand** | [**StorageListCommand**](StorageListCommand.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxmoxUpdate

> ProxmoxUpdate(ctx).UpdateProxmoxCommand(updateProxmoxCommand).Execute()

Update proxmox credentials

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
    updateProxmoxCommand := *openapiclient.NewUpdateProxmoxCommand() // UpdateProxmoxCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxUpdate(context.Background()).UpdateProxmoxCommand(updateProxmoxCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProxmoxCommand** | [**UpdateProxmoxCommand**](UpdateProxmoxCommand.md) |  | 

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


## ProxmoxUpdateHypervisors

> ProxmoxUpdateHypervisors(ctx).UpdateHypervisorsCommand(updateHypervisorsCommand).Execute()

Update proxmox credentials

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
    updateHypervisorsCommand := *openapiclient.NewUpdateHypervisorsCommand() // UpdateHypervisorsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxUpdateHypervisors(context.Background()).UpdateHypervisorsCommand(updateHypervisorsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxUpdateHypervisors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxUpdateHypervisorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateHypervisorsCommand** | [**UpdateHypervisorsCommand**](UpdateHypervisorsCommand.md) |  | 

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


## ProxmoxUpdateIpAddresses

> ProxmoxUpdateIpAddresses(ctx).UpdateUsedIpAddressesCommand(updateUsedIpAddressesCommand).Execute()

Update proxmox network used ip addresses

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
    updateUsedIpAddressesCommand := *openapiclient.NewUpdateUsedIpAddressesCommand() // UpdateUsedIpAddressesCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxUpdateIpAddresses(context.Background()).UpdateUsedIpAddressesCommand(updateUsedIpAddressesCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxUpdateIpAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxUpdateIpAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUsedIpAddressesCommand** | [**UpdateUsedIpAddressesCommand**](UpdateUsedIpAddressesCommand.md) |  | 

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


## ProxmoxVmTemplateList

> []CommonDropdownDto ProxmoxVmTemplateList(ctx).VmTemplateListCommand(vmTemplateListCommand).Execute()

Fetch proxmox vm template list

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
    vmTemplateListCommand := *openapiclient.NewVmTemplateListCommand() // VmTemplateListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProxmoxCloudCredentialAPI.ProxmoxVmTemplateList(context.Background()).VmTemplateListCommand(vmTemplateListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProxmoxCloudCredentialAPI.ProxmoxVmTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProxmoxVmTemplateList`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `ProxmoxCloudCredentialAPI.ProxmoxVmTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxmoxVmTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vmTemplateListCommand** | [**VmTemplateListCommand**](VmTemplateListCommand.md) |  | 

### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

