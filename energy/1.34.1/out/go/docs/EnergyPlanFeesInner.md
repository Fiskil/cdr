# EnergyPlanFeesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the fee. | 
**Term** | **string** | The term of the fee. | 
**Amount** | Pointer to **string** | The fee amount. Required if _term_ is not &#x60;PERCENT_OF_BILL&#x60;. | [optional] 
**Rate** | Pointer to **string** | The fee rate. Required if _term_ is &#x60;PERCENT_OF_BILL&#x60;. | [optional] 
**Description** | Pointer to **string** | A description of the fee. | [optional] 

## Methods

### NewEnergyPlanFeesInner

`func NewEnergyPlanFeesInner(type_ string, term string, ) *EnergyPlanFeesInner`

NewEnergyPlanFeesInner instantiates a new EnergyPlanFeesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanFeesInnerWithDefaults

`func NewEnergyPlanFeesInnerWithDefaults() *EnergyPlanFeesInner`

NewEnergyPlanFeesInnerWithDefaults instantiates a new EnergyPlanFeesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnergyPlanFeesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanFeesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanFeesInner) SetType(v string)`

SetType sets Type field to given value.


### GetTerm

`func (o *EnergyPlanFeesInner) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *EnergyPlanFeesInner) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *EnergyPlanFeesInner) SetTerm(v string)`

SetTerm sets Term field to given value.


### GetAmount

`func (o *EnergyPlanFeesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyPlanFeesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyPlanFeesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EnergyPlanFeesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRate

`func (o *EnergyPlanFeesInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *EnergyPlanFeesInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *EnergyPlanFeesInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *EnergyPlanFeesInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetDescription

`func (o *EnergyPlanFeesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanFeesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanFeesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanFeesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


