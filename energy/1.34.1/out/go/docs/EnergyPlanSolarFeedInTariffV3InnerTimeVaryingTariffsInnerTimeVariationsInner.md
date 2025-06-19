# EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | [**[]EnergyDaysEnum**](EnergyDaysEnum.md) | The days that the tariff applies to. At least one entry required. | 
**StartTime** | Pointer to **string** | The beginning of the time period per day for which the tariff applies. If absent assumes start of day (i.e. midnight).  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | [optional] 
**EndTime** | Pointer to **string** | The end of the time period per day for which the tariff applies. If absent assumes end of day (i.e. one second before midnight).  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | [optional] 

## Methods

### NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner

`func NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner(days []EnergyDaysEnum, ) *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner`

NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner instantiates a new EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInnerWithDefaults

`func NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInnerWithDefaults() *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner`

NewEnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInnerWithDefaults instantiates a new EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) GetDays() []EnergyDaysEnum`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) GetDaysOk() (*[]EnergyDaysEnum, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) SetDays(v []EnergyDaysEnum)`

SetDays sets Days field to given value.


### GetStartTime

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInnerTimeVariationsInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


