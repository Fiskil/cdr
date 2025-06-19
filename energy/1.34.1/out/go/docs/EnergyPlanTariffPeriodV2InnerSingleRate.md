# EnergyPlanTariffPeriodV2InnerSingleRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the rate. | 
**Description** | Pointer to **string** | Description of the rate. | [optional] 
**GeneralUnitPrice** | Pointer to **string** | The block rate (unit price) for any usage above the included fixed usage, in dollars per kWh inclusive of GST. Only required if _pricingModel_ field is &#x60;QUOTA&#x60;. | [optional] 
**Rates** | [**[]EnergyPlanControlledLoadV2InnerSingleRateRatesInner**](EnergyPlanControlledLoadV2InnerSingleRateRatesInner.md) | Array of controlled load rates in order of usage volume. | 
**Period** | Pointer to **string** | Usage period for which the block rate applies. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | [optional] 

## Methods

### NewEnergyPlanTariffPeriodV2InnerSingleRate

`func NewEnergyPlanTariffPeriodV2InnerSingleRate(displayName string, rates []EnergyPlanControlledLoadV2InnerSingleRateRatesInner, ) *EnergyPlanTariffPeriodV2InnerSingleRate`

NewEnergyPlanTariffPeriodV2InnerSingleRate instantiates a new EnergyPlanTariffPeriodV2InnerSingleRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanTariffPeriodV2InnerSingleRateWithDefaults

`func NewEnergyPlanTariffPeriodV2InnerSingleRateWithDefaults() *EnergyPlanTariffPeriodV2InnerSingleRate`

NewEnergyPlanTariffPeriodV2InnerSingleRateWithDefaults instantiates a new EnergyPlanTariffPeriodV2InnerSingleRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGeneralUnitPrice

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetGeneralUnitPrice() string`

GetGeneralUnitPrice returns the GeneralUnitPrice field if non-nil, zero value otherwise.

### GetGeneralUnitPriceOk

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetGeneralUnitPriceOk() (*string, bool)`

GetGeneralUnitPriceOk returns a tuple with the GeneralUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralUnitPrice

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) SetGeneralUnitPrice(v string)`

SetGeneralUnitPrice sets GeneralUnitPrice field to given value.

### HasGeneralUnitPrice

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) HasGeneralUnitPrice() bool`

HasGeneralUnitPrice returns a boolean if a field has been set.

### GetRates

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetRates() []EnergyPlanControlledLoadV2InnerSingleRateRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetRatesOk() (*[]EnergyPlanControlledLoadV2InnerSingleRateRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) SetRates(v []EnergyPlanControlledLoadV2InnerSingleRateRatesInner)`

SetRates sets Rates field to given value.


### GetPeriod

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EnergyPlanTariffPeriodV2InnerSingleRate) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


