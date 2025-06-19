# EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the charging time period. If absent applies to all periods. | [optional] 
**DisplayName** | **string** | Display name of the tariff. | 
**Rates** | Pointer to [**[]EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner**](EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner.md) | Array of feed in rates. | [optional] 
**Period** | Pointer to **string** | Usage period for which the block rate applies. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). Defaults to &#x60;P1Y&#x60; if absent. | [optional] [default to "P1Y"]
**TimeVariations** | [**[]EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner**](EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner.md) | Array of time periods for which this tariff is applicable. | 

## Methods

### NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner

`func NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner(displayName string, timeVariations []EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner, ) *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner`

NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner instantiates a new EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerWithDefaults

`func NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerWithDefaults() *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner`

NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerWithDefaults instantiates a new EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRates

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetRates() []EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetRatesOk() (*[]EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) SetRates(v []EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetPeriod

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTimeVariations

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetTimeVariations() []EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner`

GetTimeVariations returns the TimeVariations field if non-nil, zero value otherwise.

### GetTimeVariationsOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) GetTimeVariationsOk() (*[]EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner, bool)`

GetTimeVariationsOk returns a tuple with the TimeVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeVariations

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner) SetTimeVariations(v []EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner)`

SetTimeVariations sets TimeVariations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


