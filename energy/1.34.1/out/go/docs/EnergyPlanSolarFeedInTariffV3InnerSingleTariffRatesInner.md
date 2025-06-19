# EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | **string** | Unit price of usage per measure unit (exclusive of GST). | 
**MeasureUnit** | Pointer to [**MeasureUnitEnum**](MeasureUnitEnum.md) | The measurement unit of rate. Assumed to be &#x60;KWH&#x60; if absent. | [optional] [default to KWH]
**Volume** | Pointer to **float32** | Volume that this rate applies to. Only applicable for &#39;stepped&#39; rates where different rates apply for different volumes of usage in a period. | [optional] 

## Methods

### NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner

`func NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner(unitPrice string, ) *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner`

NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner instantiates a new EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInnerWithDefaults

`func NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInnerWithDefaults() *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner`

NewEnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInnerWithDefaults instantiates a new EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.


### GetMeasureUnit

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) GetMeasureUnit() MeasureUnitEnum`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) GetMeasureUnitOk() (*MeasureUnitEnum, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) SetMeasureUnit(v MeasureUnitEnum)`

SetMeasureUnit sets MeasureUnit field to given value.

### HasMeasureUnit

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) HasMeasureUnit() bool`

HasMeasureUnit returns a boolean if a field has been set.

### GetVolume

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *EnergyPlanSolarFeedInTariffV3InnerSingleTariffRatesInner) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


