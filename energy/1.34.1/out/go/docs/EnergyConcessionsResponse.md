# EnergyConcessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyConcessionsResponseData**](EnergyConcessionsResponseData.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEnergyConcessionsResponse

`func NewEnergyConcessionsResponse(data EnergyConcessionsResponseData, links Links, ) *EnergyConcessionsResponse`

NewEnergyConcessionsResponse instantiates a new EnergyConcessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyConcessionsResponseWithDefaults

`func NewEnergyConcessionsResponseWithDefaults() *EnergyConcessionsResponse`

NewEnergyConcessionsResponseWithDefaults instantiates a new EnergyConcessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyConcessionsResponse) GetData() EnergyConcessionsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyConcessionsResponse) GetDataOk() (*EnergyConcessionsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyConcessionsResponse) SetData(v EnergyConcessionsResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyConcessionsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyConcessionsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyConcessionsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyConcessionsResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyConcessionsResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyConcessionsResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnergyConcessionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


