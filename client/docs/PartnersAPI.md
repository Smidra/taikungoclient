# \PartnersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartnerAddWhitelistDomain**](PartnersAPI.md#PartnerAddWhitelistDomain) | **Post** /api/v1/partner/add/whitelist/domain | Add white list domain
[**PartnerBecomeAPartner**](PartnersAPI.md#PartnerBecomeAPartner) | **Post** /api/v1/partner/become-a-partner | Become a partner
[**PartnerBindOrganizations**](PartnersAPI.md#PartnerBindOrganizations) | **Post** /api/v1/partner/bindorganizations | Bind organizations to a partner
[**PartnerContactUs**](PartnersAPI.md#PartnerContactUs) | **Post** /api/v1/partner/contact-us | Contact with us
[**PartnerCreate**](PartnersAPI.md#PartnerCreate) | **Post** /api/v1/partner/create | 
[**PartnerDeleteWhitelistDomain**](PartnersAPI.md#PartnerDeleteWhitelistDomain) | **Post** /api/v1/partner/delete/whitelist/domain | Delete white list domain
[**PartnerDetails**](PartnersAPI.md#PartnerDetails) | **Get** /api/v1/partner/details | Details of partners
[**PartnerDropdown**](PartnersAPI.md#PartnerDropdown) | **Get** /api/v1/partner/list | Get partners dropdown
[**PartnerInfo**](PartnersAPI.md#PartnerInfo) | **Get** /api/v1/partner/info | Get partner&#39;s registration info
[**PartnerList**](PartnersAPI.md#PartnerList) | **Get** /api/v1/partner | Get partners
[**PartnerUpdate**](PartnersAPI.md#PartnerUpdate) | **Put** /api/v1/partner/update/{id} | 



## PartnerAddWhitelistDomain

> PartnerAddWhitelistDomain(ctx).WhiteListDomainCreateCommand(whiteListDomainCreateCommand).Execute()

Add white list domain

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
    whiteListDomainCreateCommand := *openapiclient.NewWhiteListDomainCreateCommand() // WhiteListDomainCreateCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnersAPI.PartnerAddWhitelistDomain(context.Background()).WhiteListDomainCreateCommand(whiteListDomainCreateCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerAddWhitelistDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerAddWhitelistDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whiteListDomainCreateCommand** | [**WhiteListDomainCreateCommand**](WhiteListDomainCreateCommand.md) |  | 

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


## PartnerBecomeAPartner

> PartnerBecomeAPartner(ctx).BecomePartnerCommand(becomePartnerCommand).Execute()

Become a partner

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
    becomePartnerCommand := *openapiclient.NewBecomePartnerCommand() // BecomePartnerCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnersAPI.PartnerBecomeAPartner(context.Background()).BecomePartnerCommand(becomePartnerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerBecomeAPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerBecomeAPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **becomePartnerCommand** | [**BecomePartnerCommand**](BecomePartnerCommand.md) |  | 

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


## PartnerBindOrganizations

> PartnerBindOrganizations(ctx).BindOrganizationsCommand(bindOrganizationsCommand).Execute()

Bind organizations to a partner

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
    bindOrganizationsCommand := *openapiclient.NewBindOrganizationsCommand() // BindOrganizationsCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnersAPI.PartnerBindOrganizations(context.Background()).BindOrganizationsCommand(bindOrganizationsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerBindOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerBindOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindOrganizationsCommand** | [**BindOrganizationsCommand**](BindOrganizationsCommand.md) |  | 

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


## PartnerContactUs

> PartnerContactUs(ctx).ContactUsCommand(contactUsCommand).Execute()

Contact with us

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
    contactUsCommand := *openapiclient.NewContactUsCommand() // ContactUsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnersAPI.PartnerContactUs(context.Background()).ContactUsCommand(contactUsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerContactUs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerContactUsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactUsCommand** | [**ContactUsCommand**](ContactUsCommand.md) |  | 

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


## PartnerCreate

> PartnerCreate(ctx).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).Logo(logo).BackgroundImage(backgroundImage).Execute()



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
    allowRegistration := true // bool |  (optional)
    requiredUserApproval := true // bool |  (optional)
    paymentEnabled := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    domain := "domain_example" // string |  (optional)
    link := "link_example" // string |  (optional)
    phone := "phone_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    country := "country_example" // string |  (optional)
    city := "city_example" // string |  (optional)
    vatNumber := "vatNumber_example" // string |  (optional)
    address := "address_example" // string |  (optional)
    logo := os.NewFile(1234, "some_file") // *os.File |  (optional)
    backgroundImage := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnersAPI.PartnerCreate(context.Background()).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).Logo(logo).BackgroundImage(backgroundImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowRegistration** | **bool** |  | 
 **requiredUserApproval** | **bool** |  | 
 **paymentEnabled** | **bool** |  | 
 **name** | **string** |  | 
 **domain** | **string** |  | 
 **link** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **country** | **string** |  | 
 **city** | **string** |  | 
 **vatNumber** | **string** |  | 
 **address** | **string** |  | 
 **logo** | ***os.File** |  | 
 **backgroundImage** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerDeleteWhitelistDomain

> PartnerDeleteWhitelistDomain(ctx).WhiteListDomainDeleteCommand(whiteListDomainDeleteCommand).Execute()

Delete white list domain

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
    whiteListDomainDeleteCommand := *openapiclient.NewWhiteListDomainDeleteCommand() // WhiteListDomainDeleteCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnersAPI.PartnerDeleteWhitelistDomain(context.Background()).WhiteListDomainDeleteCommand(whiteListDomainDeleteCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerDeleteWhitelistDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerDeleteWhitelistDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whiteListDomainDeleteCommand** | [**WhiteListDomainDeleteCommand**](WhiteListDomainDeleteCommand.md) |  | 

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


## PartnerDetails

> PartnerDetailsDto PartnerDetails(ctx).Execute()

Details of partners

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersAPI.PartnerDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerDetails`: PartnerDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.PartnerDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPartnerDetailsRequest struct via the builder pattern


### Return type

[**PartnerDetailsDto**](PartnerDetailsDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerDropdown

> []PartnerEntity PartnerDropdown(ctx).Search(search).Execute()

Get partners dropdown

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersAPI.PartnerDropdown(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerDropdown`: []PartnerEntity
    fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.PartnerDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]PartnerEntity**](PartnerEntity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerInfo

> PartnerRecordDto PartnerInfo(ctx).Domain(domain).Execute()

Get partner's registration info

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
    domain := "domain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersAPI.PartnerInfo(context.Background()).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerInfo`: PartnerRecordDto
    fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.PartnerInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** |  | 

### Return type

[**PartnerRecordDto**](PartnerRecordDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerList

> PartnersList PartnerList(ctx).Offset(offset).Limit(limit).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Execute()

Get partners

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnersAPI.PartnerList(context.Background()).Offset(offset).Limit(limit).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartnerList`: PartnersList
    fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.PartnerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 

### Return type

[**PartnersList**](PartnersList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerUpdate

> PartnerUpdate(ctx, id).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Logo(logo).BackgroundImage(backgroundImage).Execute()



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
    id := int32(56) // int32 | 
    name := "name_example" // string |  (optional)
    domain := "domain_example" // string |  (optional)
    link := "link_example" // string |  (optional)
    phone := "phone_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    country := "country_example" // string |  (optional)
    city := "city_example" // string |  (optional)
    vatNumber := "vatNumber_example" // string |  (optional)
    address := "address_example" // string |  (optional)
    allowRegistration := true // bool |  (optional)
    requiredUserApproval := true // bool |  (optional)
    paymentEnabled := true // bool |  (optional)
    logo := os.NewFile(1234, "some_file") // *os.File |  (optional)
    backgroundImage := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartnersAPI.PartnerUpdate(context.Background(), id).Name(name).Domain(domain).Link(link).Phone(phone).Email(email).Country(country).City(city).VatNumber(vatNumber).Address(address).AllowRegistration(allowRegistration).RequiredUserApproval(requiredUserApproval).PaymentEnabled(paymentEnabled).Logo(logo).BackgroundImage(backgroundImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.PartnerUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPartnerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **domain** | **string** |  | 
 **link** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **country** | **string** |  | 
 **city** | **string** |  | 
 **vatNumber** | **string** |  | 
 **address** | **string** |  | 
 **allowRegistration** | **bool** |  | 
 **requiredUserApproval** | **bool** |  | 
 **paymentEnabled** | **bool** |  | 
 **logo** | ***os.File** |  | 
 **backgroundImage** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

