# EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | **string** | The amount of daily supply charge for the band, in dollars per day exclusive of GST. | 
**MeasureUnit** | Pointer to [**MeasureUnitEnum**](MeasureUnitEnum.md) | The measurement unit of rate. Assumed to be &#x60;DAYS&#x60; if absent. | [optional] [default to DAYS]
**Volume** | Pointer to **float32** | Volume the charge applies to. | [optional] 

## Methods

### NewEnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner

`func NewEnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner(unitPrice string, ) *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner`

NewEnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner instantiates a new EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInnerWithDefaults

`func NewEnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInnerWithDefaults() *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner`

NewEnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInnerWithDefaults instantiates a new EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.


### GetMeasureUnit

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) GetMeasureUnit() MeasureUnitEnum`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) GetMeasureUnitOk() (*MeasureUnitEnum, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) SetMeasureUnit(v MeasureUnitEnum)`

SetMeasureUnit sets MeasureUnit field to given value.

### HasMeasureUnit

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) HasMeasureUnit() bool`

HasMeasureUnit returns a boolean if a field has been set.

### GetVolume

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


