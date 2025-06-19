# EnergyServicePointDetailV2AllOfRegisters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterId** | **string** | Unique identifier of the register within this service point. Is not globally unique. | 
**RegisterSuffix** | Pointer to **string** | Register suffix of the meter register where the meter reads are obtained. | [optional] 
**AveragedDailyLoad** | Pointer to **float32** | The energy delivered through a connection point or metering point over an extended period normalised to a &#39;per day&#39; basis (kWh). This value is calculated annually. | [optional] 
**RegisterConsumptionType** | **string** | Indicates the consumption type of register. | 
**NetworkTariffCode** | Pointer to **string** | The Network Tariff Code is a free text field containing a code supplied and published by the local network service provider. | [optional] 
**UnitOfMeasure** | Pointer to **string** | The unit of measure for data held in this register. | [optional] 
**TimeOfDay** | Pointer to **string** | Code to identify the time validity of register contents. | [optional] 
**Multiplier** | Pointer to **float32** | Multiplier required to take a register value and turn it into a value representing billable energy. | [optional] 
**ControlledLoad** | Pointer to **bool** | Indicates whether the energy recorded by this register is created under a Controlled Load regime. | [optional] 
**ConsumptionType** | Pointer to **string** | Actual/Subtractive Indicator. Note the details of enumeration values below: &lt;ul&gt;&lt;li&gt;&#x60;ACTUAL&#x60;: implies volume of energy actually metered between two dates&lt;/li&gt;&lt;li&gt;&#x60;CUMULATIVE&#x60;: indicates a meter reading for a specific date. A second Meter Reading is required to determine the consumption between those two Meter Reading dates.&lt;/li&gt;&lt;/ul&gt; | [optional] 

## Methods

### NewEnergyServicePointDetailV2AllOfRegisters

`func NewEnergyServicePointDetailV2AllOfRegisters(registerId string, registerConsumptionType string, ) *EnergyServicePointDetailV2AllOfRegisters`

NewEnergyServicePointDetailV2AllOfRegisters instantiates a new EnergyServicePointDetailV2AllOfRegisters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointDetailV2AllOfRegistersWithDefaults

`func NewEnergyServicePointDetailV2AllOfRegistersWithDefaults() *EnergyServicePointDetailV2AllOfRegisters`

NewEnergyServicePointDetailV2AllOfRegistersWithDefaults instantiates a new EnergyServicePointDetailV2AllOfRegisters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterId

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetRegisterId() string`

GetRegisterId returns the RegisterId field if non-nil, zero value otherwise.

### GetRegisterIdOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetRegisterIdOk() (*string, bool)`

GetRegisterIdOk returns a tuple with the RegisterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterId

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetRegisterId(v string)`

SetRegisterId sets RegisterId field to given value.


### GetRegisterSuffix

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetRegisterSuffix() string`

GetRegisterSuffix returns the RegisterSuffix field if non-nil, zero value otherwise.

### GetRegisterSuffixOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetRegisterSuffixOk() (*string, bool)`

GetRegisterSuffixOk returns a tuple with the RegisterSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSuffix

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetRegisterSuffix(v string)`

SetRegisterSuffix sets RegisterSuffix field to given value.

### HasRegisterSuffix

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasRegisterSuffix() bool`

HasRegisterSuffix returns a boolean if a field has been set.

### GetAveragedDailyLoad

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetAveragedDailyLoad() float32`

GetAveragedDailyLoad returns the AveragedDailyLoad field if non-nil, zero value otherwise.

### GetAveragedDailyLoadOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetAveragedDailyLoadOk() (*float32, bool)`

GetAveragedDailyLoadOk returns a tuple with the AveragedDailyLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragedDailyLoad

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetAveragedDailyLoad(v float32)`

SetAveragedDailyLoad sets AveragedDailyLoad field to given value.

### HasAveragedDailyLoad

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasAveragedDailyLoad() bool`

HasAveragedDailyLoad returns a boolean if a field has been set.

### GetRegisterConsumptionType

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetRegisterConsumptionType() string`

GetRegisterConsumptionType returns the RegisterConsumptionType field if non-nil, zero value otherwise.

### GetRegisterConsumptionTypeOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetRegisterConsumptionTypeOk() (*string, bool)`

GetRegisterConsumptionTypeOk returns a tuple with the RegisterConsumptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterConsumptionType

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetRegisterConsumptionType(v string)`

SetRegisterConsumptionType sets RegisterConsumptionType field to given value.


### GetNetworkTariffCode

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetNetworkTariffCode() string`

GetNetworkTariffCode returns the NetworkTariffCode field if non-nil, zero value otherwise.

### GetNetworkTariffCodeOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetNetworkTariffCodeOk() (*string, bool)`

GetNetworkTariffCodeOk returns a tuple with the NetworkTariffCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTariffCode

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetNetworkTariffCode(v string)`

SetNetworkTariffCode sets NetworkTariffCode field to given value.

### HasNetworkTariffCode

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasNetworkTariffCode() bool`

HasNetworkTariffCode returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetTimeOfDay() string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetTimeOfDayOk() (*string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetTimeOfDay(v string)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetMultiplier

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetMultiplier() float32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetMultiplierOk() (*float32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetMultiplier(v float32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetControlledLoad

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetControlledLoad() bool`

GetControlledLoad returns the ControlledLoad field if non-nil, zero value otherwise.

### GetControlledLoadOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetControlledLoadOk() (*bool, bool)`

GetControlledLoadOk returns a tuple with the ControlledLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlledLoad

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetControlledLoad(v bool)`

SetControlledLoad sets ControlledLoad field to given value.

### HasControlledLoad

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasControlledLoad() bool`

HasControlledLoad returns a boolean if a field has been set.

### GetConsumptionType

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetConsumptionType() string`

GetConsumptionType returns the ConsumptionType field if non-nil, zero value otherwise.

### GetConsumptionTypeOk

`func (o *EnergyServicePointDetailV2AllOfRegisters) GetConsumptionTypeOk() (*string, bool)`

GetConsumptionTypeOk returns a tuple with the ConsumptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionType

`func (o *EnergyServicePointDetailV2AllOfRegisters) SetConsumptionType(v string)`

SetConsumptionType sets ConsumptionType field to given value.

### HasConsumptionType

`func (o *EnergyServicePointDetailV2AllOfRegisters) HasConsumptionType() bool`

HasConsumptionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


