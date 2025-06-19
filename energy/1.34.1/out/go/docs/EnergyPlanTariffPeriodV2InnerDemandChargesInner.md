# EnergyPlanTariffPeriodV2InnerDemandChargesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the charge. | 
**Description** | Pointer to **string** | Description of the charge. | [optional] 
**Amount** | **string** | The charge amount per measure unit exclusive of GST. | 
**MeasureUnit** | Pointer to [**MeasureUnitEnum**](MeasureUnitEnum.md) | The measurement unit of charge amount. Assumed to be &#x60;KWH&#x60; if absent. | [optional] [default to KWH]
**StartTime** | **string** | Start of the period.  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | 
**EndTime** | **string** | End of the period.  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | 
**Days** | Pointer to [**[]EnergyDaysEnum**](EnergyDaysEnum.md) | The days that the demand tariff applies to. | [optional] 
**MinDemand** | Pointer to **string** | Minimum demand for this demand tariff in kW. If absent then &#x60;0.00&#x60; is assumed. | [optional] [default to "0.00"]
**MaxDemand** | Pointer to **string** | Maximum demand for this demand tariff in kW. If present, must be higher than the value of the _minDemand_ field. | [optional] 
**MeasurementPeriod** | **string** | Application period for the demand tariff. | 
**ChargePeriod** | **string** | Charge period for the demand tariff. | 

## Methods

### NewEnergyPlanTariffPeriodV2InnerDemandChargesInner

`func NewEnergyPlanTariffPeriodV2InnerDemandChargesInner(displayName string, amount string, startTime string, endTime string, measurementPeriod string, chargePeriod string, ) *EnergyPlanTariffPeriodV2InnerDemandChargesInner`

NewEnergyPlanTariffPeriodV2InnerDemandChargesInner instantiates a new EnergyPlanTariffPeriodV2InnerDemandChargesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanTariffPeriodV2InnerDemandChargesInnerWithDefaults

`func NewEnergyPlanTariffPeriodV2InnerDemandChargesInnerWithDefaults() *EnergyPlanTariffPeriodV2InnerDemandChargesInner`

NewEnergyPlanTariffPeriodV2InnerDemandChargesInnerWithDefaults instantiates a new EnergyPlanTariffPeriodV2InnerDemandChargesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMeasureUnit

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMeasureUnit() MeasureUnitEnum`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMeasureUnitOk() (*MeasureUnitEnum, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetMeasureUnit(v MeasureUnitEnum)`

SetMeasureUnit sets MeasureUnit field to given value.

### HasMeasureUnit

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) HasMeasureUnit() bool`

HasMeasureUnit returns a boolean if a field has been set.

### GetStartTime

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetDays

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetDays() []EnergyDaysEnum`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetDaysOk() (*[]EnergyDaysEnum, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetDays(v []EnergyDaysEnum)`

SetDays sets Days field to given value.

### HasDays

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetMinDemand

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMinDemand() string`

GetMinDemand returns the MinDemand field if non-nil, zero value otherwise.

### GetMinDemandOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMinDemandOk() (*string, bool)`

GetMinDemandOk returns a tuple with the MinDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDemand

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetMinDemand(v string)`

SetMinDemand sets MinDemand field to given value.

### HasMinDemand

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) HasMinDemand() bool`

HasMinDemand returns a boolean if a field has been set.

### GetMaxDemand

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMaxDemand() string`

GetMaxDemand returns the MaxDemand field if non-nil, zero value otherwise.

### GetMaxDemandOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMaxDemandOk() (*string, bool)`

GetMaxDemandOk returns a tuple with the MaxDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDemand

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetMaxDemand(v string)`

SetMaxDemand sets MaxDemand field to given value.

### HasMaxDemand

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) HasMaxDemand() bool`

HasMaxDemand returns a boolean if a field has been set.

### GetMeasurementPeriod

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMeasurementPeriod() string`

GetMeasurementPeriod returns the MeasurementPeriod field if non-nil, zero value otherwise.

### GetMeasurementPeriodOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetMeasurementPeriodOk() (*string, bool)`

GetMeasurementPeriodOk returns a tuple with the MeasurementPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriod

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetMeasurementPeriod(v string)`

SetMeasurementPeriod sets MeasurementPeriod field to given value.


### GetChargePeriod

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetChargePeriod() string`

GetChargePeriod returns the ChargePeriod field if non-nil, zero value otherwise.

### GetChargePeriodOk

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) GetChargePeriodOk() (*string, bool)`

GetChargePeriodOk returns a tuple with the ChargePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargePeriod

`func (o *EnergyPlanTariffPeriodV2InnerDemandChargesInner) SetChargePeriod(v string)`

SetChargePeriod sets ChargePeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


