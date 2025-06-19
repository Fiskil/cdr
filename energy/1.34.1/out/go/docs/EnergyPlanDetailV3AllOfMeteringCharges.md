# EnergyPlanDetailV3AllOfMeteringCharges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the charge. | 
**Description** | Pointer to **string** | Description of the charge. | [optional] 
**MinimumValue** | **string** | Minimum value of the charge if the charge is a range or the absolute value of the charge if no range is specified. | 
**MaximumValue** | Pointer to **string** | The upper limit of the charge if the charge could occur in a range. | [optional] 
**Period** | Pointer to **string** | The charges that occur on a schedule indicates the frequency. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | [optional] 

## Methods

### NewEnergyPlanDetailV3AllOfMeteringCharges

`func NewEnergyPlanDetailV3AllOfMeteringCharges(displayName string, minimumValue string, ) *EnergyPlanDetailV3AllOfMeteringCharges`

NewEnergyPlanDetailV3AllOfMeteringCharges instantiates a new EnergyPlanDetailV3AllOfMeteringCharges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanDetailV3AllOfMeteringChargesWithDefaults

`func NewEnergyPlanDetailV3AllOfMeteringChargesWithDefaults() *EnergyPlanDetailV3AllOfMeteringCharges`

NewEnergyPlanDetailV3AllOfMeteringChargesWithDefaults instantiates a new EnergyPlanDetailV3AllOfMeteringCharges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMinimumValue

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetMinimumValue() string`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetMinimumValueOk() (*string, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) SetMinimumValue(v string)`

SetMinimumValue sets MinimumValue field to given value.


### GetMaximumValue

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetMaximumValue() string`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetMaximumValueOk() (*string, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) SetMaximumValue(v string)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetPeriod

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EnergyPlanDetailV3AllOfMeteringCharges) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


