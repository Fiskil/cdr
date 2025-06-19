# EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to [**[]EnergyDaysEnum**](EnergyDaysEnum.md) | The days that the rate applies to. | [optional] 
**StartTime** | Pointer to **string** | The beginning of the time period per day for which the controlled load rate applies. Required if _endTime_ provided.  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | [optional] 
**EndTime** | Pointer to **string** | The end of the time period per day for which the controlled load rate applies. Required if _startTime_ provided.  Formatted according to [ISO 8601 Times](https://en.wikipedia.org/wiki/ISO_8601#Times). If the time is provided without a UTC offset, the time zone will be determined by the value of EnergyPlanContract.timeZone. | [optional] 
**AdditionalInfo** | Pointer to **string** | Display text providing more information on the controlled load, for e.g., controlled load availability if specific day/time is not known. Required if _startTime_ and _endTime_ absent or if _additionalInfoUri_ provided. | [optional] 
**AdditionalInfoUri** | Pointer to **string** | Optional link to additional information regarding the controlled load. | [optional] 

## Methods

### NewEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner

`func NewEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner() *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner`

NewEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner instantiates a new EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInnerWithDefaults

`func NewEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInnerWithDefaults() *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner`

NewEnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInnerWithDefaults instantiates a new EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetDays() []EnergyDaysEnum`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetDaysOk() (*[]EnergyDaysEnum, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetDays(v []EnergyDaysEnum)`

SetDays sets Days field to given value.

### HasDays

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStartTime

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAdditionalInfoUri

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetAdditionalInfoUri() string`

GetAdditionalInfoUri returns the AdditionalInfoUri field if non-nil, zero value otherwise.

### GetAdditionalInfoUriOk

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) GetAdditionalInfoUriOk() (*string, bool)`

GetAdditionalInfoUriOk returns a tuple with the AdditionalInfoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfoUri

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) SetAdditionalInfoUri(v string)`

SetAdditionalInfoUri sets AdditionalInfoUri field to given value.

### HasAdditionalInfoUri

`func (o *EnergyPlanControlledLoadV2InnerTimeOfUseRatesInnerTimeOfUseInner) HasAdditionalInfoUri() bool`

HasAdditionalInfoUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


