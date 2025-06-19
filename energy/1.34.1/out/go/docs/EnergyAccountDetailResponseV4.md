# EnergyAccountDetailResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyAccountDetailV4**](EnergyAccountDetailV4.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEnergyAccountDetailResponseV4

`func NewEnergyAccountDetailResponseV4(data EnergyAccountDetailV4, links Links, ) *EnergyAccountDetailResponseV4`

NewEnergyAccountDetailResponseV4 instantiates a new EnergyAccountDetailResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyAccountDetailResponseV4WithDefaults

`func NewEnergyAccountDetailResponseV4WithDefaults() *EnergyAccountDetailResponseV4`

NewEnergyAccountDetailResponseV4WithDefaults instantiates a new EnergyAccountDetailResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyAccountDetailResponseV4) GetData() EnergyAccountDetailV4`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyAccountDetailResponseV4) GetDataOk() (*EnergyAccountDetailV4, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyAccountDetailResponseV4) SetData(v EnergyAccountDetailV4)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyAccountDetailResponseV4) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyAccountDetailResponseV4) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyAccountDetailResponseV4) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyAccountDetailResponseV4) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyAccountDetailResponseV4) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyAccountDetailResponseV4) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnergyAccountDetailResponseV4) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


