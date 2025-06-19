# RequestAccountIdListV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**RequestAccountIdListV1Data**](RequestAccountIdListV1Data.md) |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRequestAccountIdListV1

`func NewRequestAccountIdListV1(data RequestAccountIdListV1Data, ) *RequestAccountIdListV1`

NewRequestAccountIdListV1 instantiates a new RequestAccountIdListV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAccountIdListV1WithDefaults

`func NewRequestAccountIdListV1WithDefaults() *RequestAccountIdListV1`

NewRequestAccountIdListV1WithDefaults instantiates a new RequestAccountIdListV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RequestAccountIdListV1) GetData() RequestAccountIdListV1Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestAccountIdListV1) GetDataOk() (*RequestAccountIdListV1Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestAccountIdListV1) SetData(v RequestAccountIdListV1Data)`

SetData sets Data field to given value.


### GetMeta

`func (o *RequestAccountIdListV1) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RequestAccountIdListV1) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RequestAccountIdListV1) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RequestAccountIdListV1) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


