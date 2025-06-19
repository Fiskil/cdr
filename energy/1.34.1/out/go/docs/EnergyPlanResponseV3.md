# EnergyPlanResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyPlanDetailV3**](EnergyPlanDetailV3.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEnergyPlanResponseV3

`func NewEnergyPlanResponseV3(data EnergyPlanDetailV3, links Links, ) *EnergyPlanResponseV3`

NewEnergyPlanResponseV3 instantiates a new EnergyPlanResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanResponseV3WithDefaults

`func NewEnergyPlanResponseV3WithDefaults() *EnergyPlanResponseV3`

NewEnergyPlanResponseV3WithDefaults instantiates a new EnergyPlanResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyPlanResponseV3) GetData() EnergyPlanDetailV3`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyPlanResponseV3) GetDataOk() (*EnergyPlanDetailV3, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyPlanResponseV3) SetData(v EnergyPlanDetailV3)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyPlanResponseV3) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyPlanResponseV3) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyPlanResponseV3) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyPlanResponseV3) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyPlanResponseV3) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyPlanResponseV3) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnergyPlanResponseV3) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


