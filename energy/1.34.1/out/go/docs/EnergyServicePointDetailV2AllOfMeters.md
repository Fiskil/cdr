# EnergyServicePointDetailV2AllOfMeters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeterId** | **string** | The meter ID uniquely identifies a meter for a given service point. It is unique in the context of the service point. It is not globally unique. | 
**Specifications** | [**EnergyServicePointDetailV2AllOfSpecifications**](EnergyServicePointDetailV2AllOfSpecifications.md) |  | 
**Registers** | Pointer to [**[]EnergyServicePointDetailV2AllOfRegisters**](EnergyServicePointDetailV2AllOfRegisters.md) | Usage data registers available from the meter. This may be empty where there are no meters physically installed at the service point. | [optional] 

## Methods

### NewEnergyServicePointDetailV2AllOfMeters

`func NewEnergyServicePointDetailV2AllOfMeters(meterId string, specifications EnergyServicePointDetailV2AllOfSpecifications, ) *EnergyServicePointDetailV2AllOfMeters`

NewEnergyServicePointDetailV2AllOfMeters instantiates a new EnergyServicePointDetailV2AllOfMeters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointDetailV2AllOfMetersWithDefaults

`func NewEnergyServicePointDetailV2AllOfMetersWithDefaults() *EnergyServicePointDetailV2AllOfMeters`

NewEnergyServicePointDetailV2AllOfMetersWithDefaults instantiates a new EnergyServicePointDetailV2AllOfMeters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterId

`func (o *EnergyServicePointDetailV2AllOfMeters) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *EnergyServicePointDetailV2AllOfMeters) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *EnergyServicePointDetailV2AllOfMeters) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetSpecifications

`func (o *EnergyServicePointDetailV2AllOfMeters) GetSpecifications() EnergyServicePointDetailV2AllOfSpecifications`

GetSpecifications returns the Specifications field if non-nil, zero value otherwise.

### GetSpecificationsOk

`func (o *EnergyServicePointDetailV2AllOfMeters) GetSpecificationsOk() (*EnergyServicePointDetailV2AllOfSpecifications, bool)`

GetSpecificationsOk returns a tuple with the Specifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifications

`func (o *EnergyServicePointDetailV2AllOfMeters) SetSpecifications(v EnergyServicePointDetailV2AllOfSpecifications)`

SetSpecifications sets Specifications field to given value.


### GetRegisters

`func (o *EnergyServicePointDetailV2AllOfMeters) GetRegisters() []EnergyServicePointDetailV2AllOfRegisters`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *EnergyServicePointDetailV2AllOfMeters) GetRegistersOk() (*[]EnergyServicePointDetailV2AllOfRegisters, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *EnergyServicePointDetailV2AllOfMeters) SetRegisters(v []EnergyServicePointDetailV2AllOfRegisters)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *EnergyServicePointDetailV2AllOfMeters) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


