/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the ImageByIdCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageByIdCommand{}

// ImageByIdCommand struct for ImageByIdCommand
type ImageByIdCommand struct {
	CloudId *int32         `json:"cloudId,omitempty"`
	ImageId NullableString `json:"imageId,omitempty"`
}

// NewImageByIdCommand instantiates a new ImageByIdCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageByIdCommand() *ImageByIdCommand {
	this := ImageByIdCommand{}
	return &this
}

// NewImageByIdCommandWithDefaults instantiates a new ImageByIdCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageByIdCommandWithDefaults() *ImageByIdCommand {
	this := ImageByIdCommand{}
	return &this
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *ImageByIdCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId) {
		var ret int32
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageByIdCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudId) {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *ImageByIdCommand) HasCloudId() bool {
	if o != nil && !IsNil(o.CloudId) {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given int32 and assigns it to the CloudId field.
func (o *ImageByIdCommand) SetCloudId(v int32) {
	o.CloudId = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageByIdCommand) GetImageId() string {
	if o == nil || IsNil(o.ImageId.Get()) {
		var ret string
		return ret
	}
	return *o.ImageId.Get()
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageByIdCommand) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageId.Get(), o.ImageId.IsSet()
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageByIdCommand) HasImageId() bool {
	if o != nil && o.ImageId.IsSet() {
		return true
	}

	return false
}

// SetImageId gets a reference to the given NullableString and assigns it to the ImageId field.
func (o *ImageByIdCommand) SetImageId(v string) {
	o.ImageId.Set(&v)
}

// SetImageIdNil sets the value for ImageId to be an explicit nil
func (o *ImageByIdCommand) SetImageIdNil() {
	o.ImageId.Set(nil)
}

// UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
func (o *ImageByIdCommand) UnsetImageId() {
	o.ImageId.Unset()
}

func (o ImageByIdCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageByIdCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudId) {
		toSerialize["cloudId"] = o.CloudId
	}
	if o.ImageId.IsSet() {
		toSerialize["imageId"] = o.ImageId.Get()
	}
	return toSerialize, nil
}

type NullableImageByIdCommand struct {
	value *ImageByIdCommand
	isSet bool
}

func (v NullableImageByIdCommand) Get() *ImageByIdCommand {
	return v.value
}

func (v *NullableImageByIdCommand) Set(val *ImageByIdCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableImageByIdCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableImageByIdCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageByIdCommand(val *ImageByIdCommand) *NullableImageByIdCommand {
	return &NullableImageByIdCommand{value: val, isSet: true}
}

func (v NullableImageByIdCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageByIdCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
