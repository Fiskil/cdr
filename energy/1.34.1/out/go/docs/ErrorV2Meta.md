# ErrorV2Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urn** | Pointer to **string** | The CDR error code URN which the application-specific error code extends. Mandatory if the error _code_ is an application-specific error rather than a standardised error code. | [optional] 

## Methods

### NewErrorV2Meta

`func NewErrorV2Meta() *ErrorV2Meta`

NewErrorV2Meta instantiates a new ErrorV2Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorV2MetaWithDefaults

`func NewErrorV2MetaWithDefaults() *ErrorV2Meta`

NewErrorV2MetaWithDefaults instantiates a new ErrorV2Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrn

`func (o *ErrorV2Meta) GetUrn() string`

GetUrn returns the Urn field if non-nil, zero value otherwise.

### GetUrnOk

`func (o *ErrorV2Meta) GetUrnOk() (*string, bool)`

GetUrnOk returns a tuple with the Urn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrn

`func (o *ErrorV2Meta) SetUrn(v string)`

SetUrn sets Urn field to given value.

### HasUrn

`func (o *ErrorV2Meta) HasUrn() bool`

HasUrn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


