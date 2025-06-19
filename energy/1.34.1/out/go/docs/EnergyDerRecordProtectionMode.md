# EnergyDerRecordProtectionMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportLimitKva** | Pointer to **float32** | Maximum amount of power (kVA) that may be exported from a connection point to the grid, as monitored by a control/relay function. An absent value indicates no limit. | [optional] 
**UnderFrequencyProtection** | Pointer to **float32** | Protective function limit in Hz. | [optional] 
**UnderFrequencyProtectionDelay** | Pointer to **float32** | Trip delay time in seconds. | [optional] 
**OverFrequencyProtection** | Pointer to **float32** | Protective function limit in Hz. | [optional] 
**OverFrequencyProtectionDelay** | Pointer to **float32** | Trip delay time in seconds. | [optional] 
**UnderVoltageProtection** | Pointer to **float32** | Protective function limit in V. | [optional] 
**UnderVoltageProtectionDelay** | Pointer to **float32** | Trip delay time in seconds. | [optional] 
**OverVoltageProtection** | Pointer to **float32** | Protective function limit in V. | [optional] 
**OverVoltageProtectionDelay** | Pointer to **float32** | Trip delay time in seconds. | [optional] 
**SustainedOverVoltage** | Pointer to **float32** | Sustained over voltage. | [optional] 
**SustainedOverVoltageDelay** | Pointer to **float32** | Sustained Over voltage protection delay in seconds. | [optional] 
**FrequencyRateOfChange** | Pointer to **float32** | Rate of change of frequency trip point (Hz/s). | [optional] 
**VoltageVectorShift** | Pointer to **float32** | Trip angle in degrees. | [optional] 
**InterTripScheme** | Pointer to **string** | Description of the form of inter-trip (e.g., &#39;from local substation&#39;). | [optional] 
**NeutralVoltageDisplacement** | Pointer to **float32** | Trip voltage. | [optional] 

## Methods

### NewEnergyDerRecordProtectionMode

`func NewEnergyDerRecordProtectionMode() *EnergyDerRecordProtectionMode`

NewEnergyDerRecordProtectionMode instantiates a new EnergyDerRecordProtectionMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyDerRecordProtectionModeWithDefaults

`func NewEnergyDerRecordProtectionModeWithDefaults() *EnergyDerRecordProtectionMode`

NewEnergyDerRecordProtectionModeWithDefaults instantiates a new EnergyDerRecordProtectionMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportLimitKva

`func (o *EnergyDerRecordProtectionMode) GetExportLimitKva() float32`

GetExportLimitKva returns the ExportLimitKva field if non-nil, zero value otherwise.

### GetExportLimitKvaOk

`func (o *EnergyDerRecordProtectionMode) GetExportLimitKvaOk() (*float32, bool)`

GetExportLimitKvaOk returns a tuple with the ExportLimitKva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportLimitKva

`func (o *EnergyDerRecordProtectionMode) SetExportLimitKva(v float32)`

SetExportLimitKva sets ExportLimitKva field to given value.

### HasExportLimitKva

`func (o *EnergyDerRecordProtectionMode) HasExportLimitKva() bool`

HasExportLimitKva returns a boolean if a field has been set.

### GetUnderFrequencyProtection

`func (o *EnergyDerRecordProtectionMode) GetUnderFrequencyProtection() float32`

GetUnderFrequencyProtection returns the UnderFrequencyProtection field if non-nil, zero value otherwise.

### GetUnderFrequencyProtectionOk

`func (o *EnergyDerRecordProtectionMode) GetUnderFrequencyProtectionOk() (*float32, bool)`

GetUnderFrequencyProtectionOk returns a tuple with the UnderFrequencyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderFrequencyProtection

`func (o *EnergyDerRecordProtectionMode) SetUnderFrequencyProtection(v float32)`

SetUnderFrequencyProtection sets UnderFrequencyProtection field to given value.

### HasUnderFrequencyProtection

`func (o *EnergyDerRecordProtectionMode) HasUnderFrequencyProtection() bool`

HasUnderFrequencyProtection returns a boolean if a field has been set.

### GetUnderFrequencyProtectionDelay

`func (o *EnergyDerRecordProtectionMode) GetUnderFrequencyProtectionDelay() float32`

GetUnderFrequencyProtectionDelay returns the UnderFrequencyProtectionDelay field if non-nil, zero value otherwise.

### GetUnderFrequencyProtectionDelayOk

`func (o *EnergyDerRecordProtectionMode) GetUnderFrequencyProtectionDelayOk() (*float32, bool)`

GetUnderFrequencyProtectionDelayOk returns a tuple with the UnderFrequencyProtectionDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderFrequencyProtectionDelay

`func (o *EnergyDerRecordProtectionMode) SetUnderFrequencyProtectionDelay(v float32)`

SetUnderFrequencyProtectionDelay sets UnderFrequencyProtectionDelay field to given value.

### HasUnderFrequencyProtectionDelay

`func (o *EnergyDerRecordProtectionMode) HasUnderFrequencyProtectionDelay() bool`

HasUnderFrequencyProtectionDelay returns a boolean if a field has been set.

### GetOverFrequencyProtection

`func (o *EnergyDerRecordProtectionMode) GetOverFrequencyProtection() float32`

GetOverFrequencyProtection returns the OverFrequencyProtection field if non-nil, zero value otherwise.

### GetOverFrequencyProtectionOk

`func (o *EnergyDerRecordProtectionMode) GetOverFrequencyProtectionOk() (*float32, bool)`

GetOverFrequencyProtectionOk returns a tuple with the OverFrequencyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverFrequencyProtection

`func (o *EnergyDerRecordProtectionMode) SetOverFrequencyProtection(v float32)`

SetOverFrequencyProtection sets OverFrequencyProtection field to given value.

### HasOverFrequencyProtection

`func (o *EnergyDerRecordProtectionMode) HasOverFrequencyProtection() bool`

HasOverFrequencyProtection returns a boolean if a field has been set.

### GetOverFrequencyProtectionDelay

`func (o *EnergyDerRecordProtectionMode) GetOverFrequencyProtectionDelay() float32`

GetOverFrequencyProtectionDelay returns the OverFrequencyProtectionDelay field if non-nil, zero value otherwise.

### GetOverFrequencyProtectionDelayOk

`func (o *EnergyDerRecordProtectionMode) GetOverFrequencyProtectionDelayOk() (*float32, bool)`

GetOverFrequencyProtectionDelayOk returns a tuple with the OverFrequencyProtectionDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverFrequencyProtectionDelay

`func (o *EnergyDerRecordProtectionMode) SetOverFrequencyProtectionDelay(v float32)`

SetOverFrequencyProtectionDelay sets OverFrequencyProtectionDelay field to given value.

### HasOverFrequencyProtectionDelay

`func (o *EnergyDerRecordProtectionMode) HasOverFrequencyProtectionDelay() bool`

HasOverFrequencyProtectionDelay returns a boolean if a field has been set.

### GetUnderVoltageProtection

`func (o *EnergyDerRecordProtectionMode) GetUnderVoltageProtection() float32`

GetUnderVoltageProtection returns the UnderVoltageProtection field if non-nil, zero value otherwise.

### GetUnderVoltageProtectionOk

`func (o *EnergyDerRecordProtectionMode) GetUnderVoltageProtectionOk() (*float32, bool)`

GetUnderVoltageProtectionOk returns a tuple with the UnderVoltageProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderVoltageProtection

`func (o *EnergyDerRecordProtectionMode) SetUnderVoltageProtection(v float32)`

SetUnderVoltageProtection sets UnderVoltageProtection field to given value.

### HasUnderVoltageProtection

`func (o *EnergyDerRecordProtectionMode) HasUnderVoltageProtection() bool`

HasUnderVoltageProtection returns a boolean if a field has been set.

### GetUnderVoltageProtectionDelay

`func (o *EnergyDerRecordProtectionMode) GetUnderVoltageProtectionDelay() float32`

GetUnderVoltageProtectionDelay returns the UnderVoltageProtectionDelay field if non-nil, zero value otherwise.

### GetUnderVoltageProtectionDelayOk

`func (o *EnergyDerRecordProtectionMode) GetUnderVoltageProtectionDelayOk() (*float32, bool)`

GetUnderVoltageProtectionDelayOk returns a tuple with the UnderVoltageProtectionDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderVoltageProtectionDelay

`func (o *EnergyDerRecordProtectionMode) SetUnderVoltageProtectionDelay(v float32)`

SetUnderVoltageProtectionDelay sets UnderVoltageProtectionDelay field to given value.

### HasUnderVoltageProtectionDelay

`func (o *EnergyDerRecordProtectionMode) HasUnderVoltageProtectionDelay() bool`

HasUnderVoltageProtectionDelay returns a boolean if a field has been set.

### GetOverVoltageProtection

`func (o *EnergyDerRecordProtectionMode) GetOverVoltageProtection() float32`

GetOverVoltageProtection returns the OverVoltageProtection field if non-nil, zero value otherwise.

### GetOverVoltageProtectionOk

`func (o *EnergyDerRecordProtectionMode) GetOverVoltageProtectionOk() (*float32, bool)`

GetOverVoltageProtectionOk returns a tuple with the OverVoltageProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverVoltageProtection

`func (o *EnergyDerRecordProtectionMode) SetOverVoltageProtection(v float32)`

SetOverVoltageProtection sets OverVoltageProtection field to given value.

### HasOverVoltageProtection

`func (o *EnergyDerRecordProtectionMode) HasOverVoltageProtection() bool`

HasOverVoltageProtection returns a boolean if a field has been set.

### GetOverVoltageProtectionDelay

`func (o *EnergyDerRecordProtectionMode) GetOverVoltageProtectionDelay() float32`

GetOverVoltageProtectionDelay returns the OverVoltageProtectionDelay field if non-nil, zero value otherwise.

### GetOverVoltageProtectionDelayOk

`func (o *EnergyDerRecordProtectionMode) GetOverVoltageProtectionDelayOk() (*float32, bool)`

GetOverVoltageProtectionDelayOk returns a tuple with the OverVoltageProtectionDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverVoltageProtectionDelay

`func (o *EnergyDerRecordProtectionMode) SetOverVoltageProtectionDelay(v float32)`

SetOverVoltageProtectionDelay sets OverVoltageProtectionDelay field to given value.

### HasOverVoltageProtectionDelay

`func (o *EnergyDerRecordProtectionMode) HasOverVoltageProtectionDelay() bool`

HasOverVoltageProtectionDelay returns a boolean if a field has been set.

### GetSustainedOverVoltage

`func (o *EnergyDerRecordProtectionMode) GetSustainedOverVoltage() float32`

GetSustainedOverVoltage returns the SustainedOverVoltage field if non-nil, zero value otherwise.

### GetSustainedOverVoltageOk

`func (o *EnergyDerRecordProtectionMode) GetSustainedOverVoltageOk() (*float32, bool)`

GetSustainedOverVoltageOk returns a tuple with the SustainedOverVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainedOverVoltage

`func (o *EnergyDerRecordProtectionMode) SetSustainedOverVoltage(v float32)`

SetSustainedOverVoltage sets SustainedOverVoltage field to given value.

### HasSustainedOverVoltage

`func (o *EnergyDerRecordProtectionMode) HasSustainedOverVoltage() bool`

HasSustainedOverVoltage returns a boolean if a field has been set.

### GetSustainedOverVoltageDelay

`func (o *EnergyDerRecordProtectionMode) GetSustainedOverVoltageDelay() float32`

GetSustainedOverVoltageDelay returns the SustainedOverVoltageDelay field if non-nil, zero value otherwise.

### GetSustainedOverVoltageDelayOk

`func (o *EnergyDerRecordProtectionMode) GetSustainedOverVoltageDelayOk() (*float32, bool)`

GetSustainedOverVoltageDelayOk returns a tuple with the SustainedOverVoltageDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainedOverVoltageDelay

`func (o *EnergyDerRecordProtectionMode) SetSustainedOverVoltageDelay(v float32)`

SetSustainedOverVoltageDelay sets SustainedOverVoltageDelay field to given value.

### HasSustainedOverVoltageDelay

`func (o *EnergyDerRecordProtectionMode) HasSustainedOverVoltageDelay() bool`

HasSustainedOverVoltageDelay returns a boolean if a field has been set.

### GetFrequencyRateOfChange

`func (o *EnergyDerRecordProtectionMode) GetFrequencyRateOfChange() float32`

GetFrequencyRateOfChange returns the FrequencyRateOfChange field if non-nil, zero value otherwise.

### GetFrequencyRateOfChangeOk

`func (o *EnergyDerRecordProtectionMode) GetFrequencyRateOfChangeOk() (*float32, bool)`

GetFrequencyRateOfChangeOk returns a tuple with the FrequencyRateOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyRateOfChange

`func (o *EnergyDerRecordProtectionMode) SetFrequencyRateOfChange(v float32)`

SetFrequencyRateOfChange sets FrequencyRateOfChange field to given value.

### HasFrequencyRateOfChange

`func (o *EnergyDerRecordProtectionMode) HasFrequencyRateOfChange() bool`

HasFrequencyRateOfChange returns a boolean if a field has been set.

### GetVoltageVectorShift

`func (o *EnergyDerRecordProtectionMode) GetVoltageVectorShift() float32`

GetVoltageVectorShift returns the VoltageVectorShift field if non-nil, zero value otherwise.

### GetVoltageVectorShiftOk

`func (o *EnergyDerRecordProtectionMode) GetVoltageVectorShiftOk() (*float32, bool)`

GetVoltageVectorShiftOk returns a tuple with the VoltageVectorShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltageVectorShift

`func (o *EnergyDerRecordProtectionMode) SetVoltageVectorShift(v float32)`

SetVoltageVectorShift sets VoltageVectorShift field to given value.

### HasVoltageVectorShift

`func (o *EnergyDerRecordProtectionMode) HasVoltageVectorShift() bool`

HasVoltageVectorShift returns a boolean if a field has been set.

### GetInterTripScheme

`func (o *EnergyDerRecordProtectionMode) GetInterTripScheme() string`

GetInterTripScheme returns the InterTripScheme field if non-nil, zero value otherwise.

### GetInterTripSchemeOk

`func (o *EnergyDerRecordProtectionMode) GetInterTripSchemeOk() (*string, bool)`

GetInterTripSchemeOk returns a tuple with the InterTripScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterTripScheme

`func (o *EnergyDerRecordProtectionMode) SetInterTripScheme(v string)`

SetInterTripScheme sets InterTripScheme field to given value.

### HasInterTripScheme

`func (o *EnergyDerRecordProtectionMode) HasInterTripScheme() bool`

HasInterTripScheme returns a boolean if a field has been set.

### GetNeutralVoltageDisplacement

`func (o *EnergyDerRecordProtectionMode) GetNeutralVoltageDisplacement() float32`

GetNeutralVoltageDisplacement returns the NeutralVoltageDisplacement field if non-nil, zero value otherwise.

### GetNeutralVoltageDisplacementOk

`func (o *EnergyDerRecordProtectionMode) GetNeutralVoltageDisplacementOk() (*float32, bool)`

GetNeutralVoltageDisplacementOk returns a tuple with the NeutralVoltageDisplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralVoltageDisplacement

`func (o *EnergyDerRecordProtectionMode) SetNeutralVoltageDisplacement(v float32)`

SetNeutralVoltageDisplacement sets NeutralVoltageDisplacement field to given value.

### HasNeutralVoltageDisplacement

`func (o *EnergyDerRecordProtectionMode) HasNeutralVoltageDisplacement() bool`

HasNeutralVoltageDisplacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


