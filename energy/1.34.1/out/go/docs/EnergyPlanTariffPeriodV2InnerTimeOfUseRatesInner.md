# EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the rate. | 
**Description** | Pointer to **string** | Description of the rate. | [optional] 
**Rates** | [**[]EnergyPlanControlledLoadV2InnerSingleRateRatesInner**](EnergyPlanControlledLoadV2InnerSingleRateRatesInner.md) | Array of controlled load rates in order of usage volume. | 
**Period** | Pointer to **string** | Usage period for which the block rate applies. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). Defaults to &#x60;P1Y&#x60; if absent. | [optional] [default to "P1Y"]
**TimeOfUse** | [**[]EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner**](EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner.md) | Array of times of use. | 
**Type** | **string** | The type of usage that the rate applies to. | 

## Methods

### NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner

`func NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner(displayName string, rates []EnergyPlanControlledLoadV2InnerSingleRateRatesInner, timeOfUse []EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner, type_ string, ) *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner`

NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner instantiates a new EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerWithDefaults

`func NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerWithDefaults() *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner`

NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerWithDefaults instantiates a new EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRates

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetRates() []EnergyPlanControlledLoadV2InnerSingleRateRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetRatesOk() (*[]EnergyPlanControlledLoadV2InnerSingleRateRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) SetRates(v []EnergyPlanControlledLoadV2InnerSingleRateRatesInner)`

SetRates sets Rates field to given value.


### GetPeriod

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTimeOfUse

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetTimeOfUse() []EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner`

GetTimeOfUse returns the TimeOfUse field if non-nil, zero value otherwise.

### GetTimeOfUseOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetTimeOfUseOk() (*[]EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner, bool)`

GetTimeOfUseOk returns a tuple with the TimeOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfUse

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) SetTimeOfUse(v []EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner)`

SetTimeOfUse sets TimeOfUse field to given value.


### GetType

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


