# BindSubscriptionResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** |  | [optional] 
**PaymentIntentClientSecret** | Pointer to **NullableString** |  | [optional] 
**PaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**InvoiceFailureCode** | Pointer to **NullableString** |  | [optional] 
**InvoiceFailureMessage** | Pointer to **NullableString** |  | [optional] 
**InvoiceFailureReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBindSubscriptionResponseDto

`func NewBindSubscriptionResponseDto() *BindSubscriptionResponseDto`

NewBindSubscriptionResponseDto instantiates a new BindSubscriptionResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindSubscriptionResponseDtoWithDefaults

`func NewBindSubscriptionResponseDtoWithDefaults() *BindSubscriptionResponseDto`

NewBindSubscriptionResponseDtoWithDefaults instantiates a new BindSubscriptionResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BindSubscriptionResponseDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BindSubscriptionResponseDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BindSubscriptionResponseDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BindSubscriptionResponseDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BindSubscriptionResponseDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BindSubscriptionResponseDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPaymentIntentClientSecret

`func (o *BindSubscriptionResponseDto) GetPaymentIntentClientSecret() string`

GetPaymentIntentClientSecret returns the PaymentIntentClientSecret field if non-nil, zero value otherwise.

### GetPaymentIntentClientSecretOk

`func (o *BindSubscriptionResponseDto) GetPaymentIntentClientSecretOk() (*string, bool)`

GetPaymentIntentClientSecretOk returns a tuple with the PaymentIntentClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentClientSecret

`func (o *BindSubscriptionResponseDto) SetPaymentIntentClientSecret(v string)`

SetPaymentIntentClientSecret sets PaymentIntentClientSecret field to given value.

### HasPaymentIntentClientSecret

`func (o *BindSubscriptionResponseDto) HasPaymentIntentClientSecret() bool`

HasPaymentIntentClientSecret returns a boolean if a field has been set.

### SetPaymentIntentClientSecretNil

`func (o *BindSubscriptionResponseDto) SetPaymentIntentClientSecretNil(b bool)`

 SetPaymentIntentClientSecretNil sets the value for PaymentIntentClientSecret to be an explicit nil

### UnsetPaymentIntentClientSecret
`func (o *BindSubscriptionResponseDto) UnsetPaymentIntentClientSecret()`

UnsetPaymentIntentClientSecret ensures that no value is present for PaymentIntentClientSecret, not even an explicit nil
### GetPaymentIntentId

`func (o *BindSubscriptionResponseDto) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *BindSubscriptionResponseDto) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *BindSubscriptionResponseDto) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *BindSubscriptionResponseDto) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### SetPaymentIntentIdNil

`func (o *BindSubscriptionResponseDto) SetPaymentIntentIdNil(b bool)`

 SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil

### UnsetPaymentIntentId
`func (o *BindSubscriptionResponseDto) UnsetPaymentIntentId()`

UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
### GetInvoiceFailureCode

`func (o *BindSubscriptionResponseDto) GetInvoiceFailureCode() string`

GetInvoiceFailureCode returns the InvoiceFailureCode field if non-nil, zero value otherwise.

### GetInvoiceFailureCodeOk

`func (o *BindSubscriptionResponseDto) GetInvoiceFailureCodeOk() (*string, bool)`

GetInvoiceFailureCodeOk returns a tuple with the InvoiceFailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFailureCode

`func (o *BindSubscriptionResponseDto) SetInvoiceFailureCode(v string)`

SetInvoiceFailureCode sets InvoiceFailureCode field to given value.

### HasInvoiceFailureCode

`func (o *BindSubscriptionResponseDto) HasInvoiceFailureCode() bool`

HasInvoiceFailureCode returns a boolean if a field has been set.

### SetInvoiceFailureCodeNil

`func (o *BindSubscriptionResponseDto) SetInvoiceFailureCodeNil(b bool)`

 SetInvoiceFailureCodeNil sets the value for InvoiceFailureCode to be an explicit nil

### UnsetInvoiceFailureCode
`func (o *BindSubscriptionResponseDto) UnsetInvoiceFailureCode()`

UnsetInvoiceFailureCode ensures that no value is present for InvoiceFailureCode, not even an explicit nil
### GetInvoiceFailureMessage

`func (o *BindSubscriptionResponseDto) GetInvoiceFailureMessage() string`

GetInvoiceFailureMessage returns the InvoiceFailureMessage field if non-nil, zero value otherwise.

### GetInvoiceFailureMessageOk

`func (o *BindSubscriptionResponseDto) GetInvoiceFailureMessageOk() (*string, bool)`

GetInvoiceFailureMessageOk returns a tuple with the InvoiceFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFailureMessage

`func (o *BindSubscriptionResponseDto) SetInvoiceFailureMessage(v string)`

SetInvoiceFailureMessage sets InvoiceFailureMessage field to given value.

### HasInvoiceFailureMessage

`func (o *BindSubscriptionResponseDto) HasInvoiceFailureMessage() bool`

HasInvoiceFailureMessage returns a boolean if a field has been set.

### SetInvoiceFailureMessageNil

`func (o *BindSubscriptionResponseDto) SetInvoiceFailureMessageNil(b bool)`

 SetInvoiceFailureMessageNil sets the value for InvoiceFailureMessage to be an explicit nil

### UnsetInvoiceFailureMessage
`func (o *BindSubscriptionResponseDto) UnsetInvoiceFailureMessage()`

UnsetInvoiceFailureMessage ensures that no value is present for InvoiceFailureMessage, not even an explicit nil
### GetInvoiceFailureReason

`func (o *BindSubscriptionResponseDto) GetInvoiceFailureReason() string`

GetInvoiceFailureReason returns the InvoiceFailureReason field if non-nil, zero value otherwise.

### GetInvoiceFailureReasonOk

`func (o *BindSubscriptionResponseDto) GetInvoiceFailureReasonOk() (*string, bool)`

GetInvoiceFailureReasonOk returns a tuple with the InvoiceFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFailureReason

`func (o *BindSubscriptionResponseDto) SetInvoiceFailureReason(v string)`

SetInvoiceFailureReason sets InvoiceFailureReason field to given value.

### HasInvoiceFailureReason

`func (o *BindSubscriptionResponseDto) HasInvoiceFailureReason() bool`

HasInvoiceFailureReason returns a boolean if a field has been set.

### SetInvoiceFailureReasonNil

`func (o *BindSubscriptionResponseDto) SetInvoiceFailureReasonNil(b bool)`

 SetInvoiceFailureReasonNil sets the value for InvoiceFailureReason to be an explicit nil

### UnsetInvoiceFailureReason
`func (o *BindSubscriptionResponseDto) UnsetInvoiceFailureReason()`

UnsetInvoiceFailureReason ensures that no value is present for InvoiceFailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


