# EnergyPaymentScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyPaymentScheduleResponseData**](EnergyPaymentScheduleResponseData.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEnergyPaymentScheduleResponse

`func NewEnergyPaymentScheduleResponse(data EnergyPaymentScheduleResponseData, links Links, ) *EnergyPaymentScheduleResponse`

NewEnergyPaymentScheduleResponse instantiates a new EnergyPaymentScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPaymentScheduleResponseWithDefaults

`func NewEnergyPaymentScheduleResponseWithDefaults() *EnergyPaymentScheduleResponse`

NewEnergyPaymentScheduleResponseWithDefaults instantiates a new EnergyPaymentScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyPaymentScheduleResponse) GetData() EnergyPaymentScheduleResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyPaymentScheduleResponse) GetDataOk() (*EnergyPaymentScheduleResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyPaymentScheduleResponse) SetData(v EnergyPaymentScheduleResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyPaymentScheduleResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyPaymentScheduleResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyPaymentScheduleResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyPaymentScheduleResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyPaymentScheduleResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyPaymentScheduleResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnergyPaymentScheduleResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


