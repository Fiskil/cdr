# ErrorV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code of the error encountered. Where the error is specific to the respondent, an application-specific error code, expressed as a string value. If the error is application-specific, the URN code that the specific error extends must be provided in the _meta_ object. Otherwise, the value is the error code URN. | 
**Title** | **string** | A short, human-readable summary of the problem that **MUST NOT** change from occurrence to occurrence of the problem represented by the error code. | 
**Detail** | **string** | A human-readable explanation specific to this occurrence of the problem. | 
**Meta** | Pointer to [**ErrorV2Meta**](ErrorV2Meta.md) |  | [optional] 

## Methods

### NewErrorV2

`func NewErrorV2(code string, title string, detail string, ) *ErrorV2`

NewErrorV2 instantiates a new ErrorV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorV2WithDefaults

`func NewErrorV2WithDefaults() *ErrorV2`

NewErrorV2WithDefaults instantiates a new ErrorV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorV2) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorV2) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorV2) SetCode(v string)`

SetCode sets Code field to given value.


### GetTitle

`func (o *ErrorV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorV2) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ErrorV2) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorV2) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorV2) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetMeta

`func (o *ErrorV2) GetMeta() ErrorV2Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorV2) GetMetaOk() (*ErrorV2Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorV2) SetMeta(v ErrorV2Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorV2) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


