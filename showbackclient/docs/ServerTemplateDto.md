# ServerTemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**CloudRole**](CloudRole.md) |  | [optional] 
**Flavor** | Pointer to **NullableString** |  | [optional] 
**DiskSize** | Pointer to **float64** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewServerTemplateDto

`func NewServerTemplateDto() *ServerTemplateDto`

NewServerTemplateDto instantiates a new ServerTemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTemplateDtoWithDefaults

`func NewServerTemplateDtoWithDefaults() *ServerTemplateDto`

NewServerTemplateDtoWithDefaults instantiates a new ServerTemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ServerTemplateDto) GetRole() CloudRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServerTemplateDto) GetRoleOk() (*CloudRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServerTemplateDto) SetRole(v CloudRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ServerTemplateDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetFlavor

`func (o *ServerTemplateDto) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *ServerTemplateDto) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *ServerTemplateDto) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *ServerTemplateDto) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *ServerTemplateDto) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *ServerTemplateDto) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetDiskSize

`func (o *ServerTemplateDto) GetDiskSize() float64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ServerTemplateDto) GetDiskSizeOk() (*float64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ServerTemplateDto) SetDiskSize(v float64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ServerTemplateDto) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetCount

`func (o *ServerTemplateDto) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ServerTemplateDto) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ServerTemplateDto) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ServerTemplateDto) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


