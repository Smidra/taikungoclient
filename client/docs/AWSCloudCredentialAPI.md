# \AWSCloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwsCreate**](AWSCloudCredentialAPI.md#AwsCreate) | **Post** /api/v1/aws/create | Add Aws credentials
[**AwsDeviceNames**](AWSCloudCredentialAPI.md#AwsDeviceNames) | **Post** /api/v1/aws/device-names | Aws device name list
[**AwsList**](AWSCloudCredentialAPI.md#AwsList) | **Get** /api/v1/aws/list | Retrieve list of aws cloud credentials
[**AwsOwners**](AWSCloudCredentialAPI.md#AwsOwners) | **Get** /api/v1/aws/owners | Retrieve aws verified owner list
[**AwsRegionlist**](AWSCloudCredentialAPI.md#AwsRegionlist) | **Post** /api/v1/aws/regions | Retrieve aws regions list
[**AwsUpdate**](AWSCloudCredentialAPI.md#AwsUpdate) | **Post** /api/v1/aws/update | Update AWS credentials
[**AwsValidateOwners**](AWSCloudCredentialAPI.md#AwsValidateOwners) | **Post** /api/v1/aws/validate-owners | Aws validate owners
[**AwsZones**](AWSCloudCredentialAPI.md#AwsZones) | **Post** /api/v1/aws/zones | Fetch Aws zones



## AwsCreate

> ApiResponse AwsCreate(ctx).CreateAwsCloudCommand(createAwsCloudCommand).Execute()

Add Aws credentials

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
    createAwsCloudCommand := *openapiclient.NewCreateAwsCloudCommand() // CreateAwsCloudCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialAPI.AwsCreate(context.Background()).CreateAwsCloudCommand(createAwsCloudCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAwsCloudCommand** | [**CreateAwsCloudCommand**](CreateAwsCloudCommand.md) |  | 

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


## AwsDeviceNames

> []string AwsDeviceNames(ctx).AwsBlockDeviceMappingsCommand(awsBlockDeviceMappingsCommand).Execute()

Aws device name list

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
    awsBlockDeviceMappingsCommand := *openapiclient.NewAwsBlockDeviceMappingsCommand() // AwsBlockDeviceMappingsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialAPI.AwsDeviceNames(context.Background()).AwsBlockDeviceMappingsCommand(awsBlockDeviceMappingsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsDeviceNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsDeviceNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsDeviceNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsDeviceNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsBlockDeviceMappingsCommand** | [**AwsBlockDeviceMappingsCommand**](AwsBlockDeviceMappingsCommand.md) |  | 

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


## AwsList

> AwsCredentialList AwsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of aws cloud credentials

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
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialAPI.AwsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsList`: AwsCredentialList
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsListRequest struct via the builder pattern


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

[**AwsCredentialList**](AwsCredentialList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsOwners

> []CommonStringBasedDropdownDto AwsOwners(ctx).Execute()

Retrieve aws verified owner list

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialAPI.AwsOwners(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsOwners`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsOwners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAwsOwnersRequest struct via the builder pattern


### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsRegionlist

> []AwsRegionDto AwsRegionlist(ctx).RegionListCommand(regionListCommand).Execute()

Retrieve aws regions list

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
    regionListCommand := *openapiclient.NewRegionListCommand() // RegionListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialAPI.AwsRegionlist(context.Background()).RegionListCommand(regionListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsRegionlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsRegionlist`: []AwsRegionDto
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsRegionlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsRegionlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionListCommand** | [**RegionListCommand**](RegionListCommand.md) |  | 

### Return type

[**[]AwsRegionDto**](AwsRegionDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsUpdate

> AwsUpdate(ctx).UpdateAwsCommand(updateAwsCommand).Execute()

Update AWS credentials

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
    updateAwsCommand := *openapiclient.NewUpdateAwsCommand() // UpdateAwsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCloudCredentialAPI.AwsUpdate(context.Background()).UpdateAwsCommand(updateAwsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAwsCommand** | [**UpdateAwsCommand**](UpdateAwsCommand.md) |  | 

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


## AwsValidateOwners

> AwsValidateOwners(ctx).AwsValidateOwnerCommand(awsValidateOwnerCommand).Execute()

Aws validate owners

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
    awsValidateOwnerCommand := *openapiclient.NewAwsValidateOwnerCommand() // AwsValidateOwnerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCloudCredentialAPI.AwsValidateOwners(context.Background()).AwsValidateOwnerCommand(awsValidateOwnerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsValidateOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsValidateOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsValidateOwnerCommand** | [**AwsValidateOwnerCommand**](AwsValidateOwnerCommand.md) |  | 

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


## AwsZones

> AzResult AwsZones(ctx).AmazonAvailabilityZonesCommand(amazonAvailabilityZonesCommand).Execute()

Fetch Aws zones

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
    amazonAvailabilityZonesCommand := *openapiclient.NewAmazonAvailabilityZonesCommand() // AmazonAvailabilityZonesCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCloudCredentialAPI.AwsZones(context.Background()).AmazonAvailabilityZonesCommand(amazonAvailabilityZonesCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsZones`: AzResult
    fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amazonAvailabilityZonesCommand** | [**AmazonAvailabilityZonesCommand**](AmazonAvailabilityZonesCommand.md) |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

