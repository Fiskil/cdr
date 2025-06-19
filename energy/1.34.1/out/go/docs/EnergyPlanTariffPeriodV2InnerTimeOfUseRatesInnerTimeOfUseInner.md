# EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | [**[]EnergyDaysEnum**](EnergyDaysEnum.md) | The days that the rate applies to. | 
**StartTime** | **string** | Start of the period.  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | 
**EndTime** | **string** | End of the period.  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | 

## Methods

### NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner

`func NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner(days []EnergyDaysEnum, startTime string, endTime string, ) *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner`

NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner instantiates a new EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInnerWithDefaults

`func NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInnerWithDefaults() *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner`

NewEnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInnerWithDefaults instantiates a new EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetDays() []EnergyDaysEnum`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetDaysOk() (*[]EnergyDaysEnum, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetDays(v []EnergyDaysEnum)`

SetDays sets Days field to given value.


### GetStartTime

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


