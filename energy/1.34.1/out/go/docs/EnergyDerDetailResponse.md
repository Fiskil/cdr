# EnergyDerDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyDerRecord**](EnergyDerRecord.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEnergyDerDetailResponse

`func NewEnergyDerDetailResponse(data EnergyDerRecord, links Links, ) *EnergyDerDetailResponse`

NewEnergyDerDetailResponse instantiates a new EnergyDerDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyDerDetailResponseWithDefaults

`func NewEnergyDerDetailResponseWithDefaults() *EnergyDerDetailResponse`

NewEnergyDerDetailResponseWithDefaults instantiates a new EnergyDerDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyDerDetailResponse) GetData() EnergyDerRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyDerDetailResponse) GetDataOk() (*EnergyDerRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyDerDetailResponse) SetData(v EnergyDerRecord)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyDerDetailResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyDerDetailResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyDerDetailResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyDerDetailResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyDerDetailResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyDerDetailResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnergyDerDetailResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


