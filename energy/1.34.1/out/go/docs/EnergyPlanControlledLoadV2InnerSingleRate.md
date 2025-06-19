# EnergyPlanControlledLoadV2InnerSingleRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the controlled load rate. | 
**Description** | Pointer to **string** | Description of the controlled load rate. | [optional] 
**DailySupplyCharge** | Pointer to **string** | The daily supply charge (exclusive of GST) for this controlled load tier. | [optional] 
**Rates** | [**[]EnergyPlanControlledLoadV2InnerSingleRateRatesInner**](EnergyPlanControlledLoadV2InnerSingleRateRatesInner.md) | Array of controlled load rates in order of usage volume. | 
**Period** | Pointer to **string** | Usage period for which the block rate applies. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). Defaults to &#x60;P1Y&#x60; if absent. | [optional] [default to "P1Y"]

## Methods

### NewEnergyPlanControlledLoadV2InnerSingleRate

`func NewEnergyPlanControlledLoadV2InnerSingleRate(displayName string, rates []EnergyPlanControlledLoadV2InnerSingleRateRatesInner, ) *EnergyPlanControlledLoadV2InnerSingleRate`

NewEnergyPlanControlledLoadV2InnerSingleRate instantiates a new EnergyPlanControlledLoadV2InnerSingleRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanControlledLoadV2InnerSingleRateWithDefaults

`func NewEnergyPlanControlledLoadV2InnerSingleRateWithDefaults() *EnergyPlanControlledLoadV2InnerSingleRate`

NewEnergyPlanControlledLoadV2InnerSingleRateWithDefaults instantiates a new EnergyPlanControlledLoadV2InnerSingleRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDailySupplyCharge

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetDailySupplyCharge() string`

GetDailySupplyCharge returns the DailySupplyCharge field if non-nil, zero value otherwise.

### GetDailySupplyChargeOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetDailySupplyChargeOk() (*string, bool)`

GetDailySupplyChargeOk returns a tuple with the DailySupplyCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySupplyCharge

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) SetDailySupplyCharge(v string)`

SetDailySupplyCharge sets DailySupplyCharge field to given value.

### HasDailySupplyCharge

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) HasDailySupplyCharge() bool`

HasDailySupplyCharge returns a boolean if a field has been set.

### GetRates

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetRates() []EnergyPlanControlledLoadV2InnerSingleRateRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetRatesOk() (*[]EnergyPlanControlledLoadV2InnerSingleRateRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) SetRates(v []EnergyPlanControlledLoadV2InnerSingleRateRatesInner)`

SetRates sets Rates field to given value.


### GetPeriod

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EnergyPlanControlledLoadV2InnerSingleRate) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


