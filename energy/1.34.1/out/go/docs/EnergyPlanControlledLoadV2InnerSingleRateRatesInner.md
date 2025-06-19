# EnergyPlanControlledLoadV2InnerSingleRateRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | **string** | Unit price of usage per measure unit (exclusive of GST). | 
**MeasureUnit** | Pointer to [**MeasureUnitEnum**](MeasureUnitEnum.md) | The measurement unit of rate. Assumed to be &#x60;KWH&#x60; if absent. | [optional] [default to KWH]
**Volume** | Pointer to **float32** | Volume in kWh that this rate applies to. Only applicable for &#39;stepped&#39; rates where different rates apply for different volumes of usage in a period. | [optional] 

## Methods

### NewEnergyPlanControlledLoadV2InnerSingleRateRatesInner

`func NewEnergyPlanControlledLoadV2InnerSingleRateRatesInner(unitPrice string, ) *EnergyPlanControlledLoadV2InnerSingleRateRatesInner`

NewEnergyPlanControlledLoadV2InnerSingleRateRatesInner instantiates a new EnergyPlanControlledLoadV2InnerSingleRateRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanControlledLoadV2InnerSingleRateRatesInnerWithDefaults

`func NewEnergyPlanControlledLoadV2InnerSingleRateRatesInnerWithDefaults() *EnergyPlanControlledLoadV2InnerSingleRateRatesInner`

NewEnergyPlanControlledLoadV2InnerSingleRateRatesInnerWithDefaults instantiates a new EnergyPlanControlledLoadV2InnerSingleRateRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.


### GetMeasureUnit

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) GetMeasureUnit() MeasureUnitEnum`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) GetMeasureUnitOk() (*MeasureUnitEnum, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) SetMeasureUnit(v MeasureUnitEnum)`

SetMeasureUnit sets MeasureUnit field to given value.

### HasMeasureUnit

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) HasMeasureUnit() bool`

HasMeasureUnit returns a boolean if a field has been set.

### GetVolume

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *EnergyPlanControlledLoadV2InnerSingleRateRatesInner) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


