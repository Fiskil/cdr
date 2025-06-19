# EnergyPlanSolarFeedInTariffV3InnerSingleTariff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rates** | [**[]EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner**](EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner.md) | Array of feed in rates. | 
**Period** | Pointer to **string** | Usage period for which the block rate applies. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). Defaults to &#x60;P1Y&#x60; if absent. | [optional] [default to "P1Y"]

## Methods

### NewEnergyPlanSolarFeedInTariffV3InnerSingleTariff

`func NewEnergyPlanSolarFeedInTariffV3InnerSingleTariff(rates []EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner, ) *EnergyPlanSolarFeedInTariffV3InnerSingleTariff`

NewEnergyPlanSolarFeedInTariffV3InnerSingleTariff instantiates a new EnergyPlanSolarFeedInTariffV3InnerSingleTariff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffWithDefaults

`func NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffWithDefaults() *EnergyPlanSolarFeedInTariffV3InnerSingleTariff`

NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffWithDefaults instantiates a new EnergyPlanSolarFeedInTariffV3InnerSingleTariff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRates

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariff) GetRates() []EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariff) GetRatesOk() (*[]EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariff) SetRates(v []EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner)`

SetRates sets Rates field to given value.


### GetPeriod

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariff) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariff) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariff) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariff) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


