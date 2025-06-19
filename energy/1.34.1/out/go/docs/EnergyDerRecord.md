# EnergyDerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePointId** | **string** | Unique identifier for the service point. | 
**ApprovedCapacity** | **float32** | Approved small generating unit capacity as agreed with NSP in the connection agreement, expressed in kVA. Value of &#x60;0&#x60; indicates no DER record exists for the given _servicePointId_. | 
**AvailablePhasesCount** | **int32** | The number of phases available for the installation of DER. Acceptable values are &#x60;0&#x60;, &#x60;1&#x60;, &#x60;2&#x60; or &#x60;3&#x60;. Value of &#x60;0&#x60; indicates no DER record exists for the given _servicePointId_. | 
**InstalledPhasesCount** | **int32** | The number of phases that DER is connected to. Acceptable values are &#x60;0&#x60;, &#x60;1&#x60;, &#x60;2&#x60; or &#x60;3&#x60;. Value of &#x60;0&#x60; indicates no DER record exists for the given _servicePointId_. | 
**IslandableInstallation** | **bool** | For identification of small generating units designed with the ability to operate in an islanded mode. | 
**HasCentralProtectionControl** | Pointer to **bool** | For DER installations where NSPs specify the need for additional forms of protection above those inbuilt in an inverter. If absent then assumed to be &#x60;false&#x60;. | [optional] [default to false]
**ProtectionMode** | Pointer to [**EnergyDerRecordProtectionMode**](EnergyDerRecordProtectionMode.md) |  | [optional] 
**AcConnections** | [**[]EnergyDerRecordAcConnections**](EnergyDerRecordAcConnections.md) |  | 

## Methods

### NewEnergyDerRecord

`func NewEnergyDerRecord(servicePointId string, approvedCapacity float32, availablePhasesCount int32, installedPhasesCount int32, islandableInstallation bool, acConnections []EnergyDerRecordAcConnections, ) *EnergyDerRecord`

NewEnergyDerRecord instantiates a new EnergyDerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyDerRecordWithDefaults

`func NewEnergyDerRecordWithDefaults() *EnergyDerRecord`

NewEnergyDerRecordWithDefaults instantiates a new EnergyDerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyDerRecord) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyDerRecord) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyDerRecord) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.


### GetApprovedCapacity

`func (o *EnergyDerRecord) GetApprovedCapacity() float32`

GetApprovedCapacity returns the ApprovedCapacity field if non-nil, zero value otherwise.

### GetApprovedCapacityOk

`func (o *EnergyDerRecord) GetApprovedCapacityOk() (*float32, bool)`

GetApprovedCapacityOk returns a tuple with the ApprovedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedCapacity

`func (o *EnergyDerRecord) SetApprovedCapacity(v float32)`

SetApprovedCapacity sets ApprovedCapacity field to given value.


### GetAvailablePhasesCount

`func (o *EnergyDerRecord) GetAvailablePhasesCount() int32`

GetAvailablePhasesCount returns the AvailablePhasesCount field if non-nil, zero value otherwise.

### GetAvailablePhasesCountOk

`func (o *EnergyDerRecord) GetAvailablePhasesCountOk() (*int32, bool)`

GetAvailablePhasesCountOk returns a tuple with the AvailablePhasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePhasesCount

`func (o *EnergyDerRecord) SetAvailablePhasesCount(v int32)`

SetAvailablePhasesCount sets AvailablePhasesCount field to given value.


### GetInstalledPhasesCount

`func (o *EnergyDerRecord) GetInstalledPhasesCount() int32`

GetInstalledPhasesCount returns the InstalledPhasesCount field if non-nil, zero value otherwise.

### GetInstalledPhasesCountOk

`func (o *EnergyDerRecord) GetInstalledPhasesCountOk() (*int32, bool)`

GetInstalledPhasesCountOk returns a tuple with the InstalledPhasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledPhasesCount

`func (o *EnergyDerRecord) SetInstalledPhasesCount(v int32)`

SetInstalledPhasesCount sets InstalledPhasesCount field to given value.


### GetIslandableInstallation

`func (o *EnergyDerRecord) GetIslandableInstallation() bool`

GetIslandableInstallation returns the IslandableInstallation field if non-nil, zero value otherwise.

### GetIslandableInstallationOk

`func (o *EnergyDerRecord) GetIslandableInstallationOk() (*bool, bool)`

GetIslandableInstallationOk returns a tuple with the IslandableInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIslandableInstallation

`func (o *EnergyDerRecord) SetIslandableInstallation(v bool)`

SetIslandableInstallation sets IslandableInstallation field to given value.


### GetHasCentralProtectionControl

`func (o *EnergyDerRecord) GetHasCentralProtectionControl() bool`

GetHasCentralProtectionControl returns the HasCentralProtectionControl field if non-nil, zero value otherwise.

### GetHasCentralProtectionControlOk

`func (o *EnergyDerRecord) GetHasCentralProtectionControlOk() (*bool, bool)`

GetHasCentralProtectionControlOk returns a tuple with the HasCentralProtectionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCentralProtectionControl

`func (o *EnergyDerRecord) SetHasCentralProtectionControl(v bool)`

SetHasCentralProtectionControl sets HasCentralProtectionControl field to given value.

### HasHasCentralProtectionControl

`func (o *EnergyDerRecord) HasHasCentralProtectionControl() bool`

HasHasCentralProtectionControl returns a boolean if a field has been set.

### GetProtectionMode

`func (o *EnergyDerRecord) GetProtectionMode() EnergyDerRecordProtectionMode`

GetProtectionMode returns the ProtectionMode field if non-nil, zero value otherwise.

### GetProtectionModeOk

`func (o *EnergyDerRecord) GetProtectionModeOk() (*EnergyDerRecordProtectionMode, bool)`

GetProtectionModeOk returns a tuple with the ProtectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionMode

`func (o *EnergyDerRecord) SetProtectionMode(v EnergyDerRecordProtectionMode)`

SetProtectionMode sets ProtectionMode field to given value.

### HasProtectionMode

`func (o *EnergyDerRecord) HasProtectionMode() bool`

HasProtectionMode returns a boolean if a field has been set.

### GetAcConnections

`func (o *EnergyDerRecord) GetAcConnections() []EnergyDerRecordAcConnections`

GetAcConnections returns the AcConnections field if non-nil, zero value otherwise.

### GetAcConnectionsOk

`func (o *EnergyDerRecord) GetAcConnectionsOk() (*[]EnergyDerRecordAcConnections, bool)`

GetAcConnectionsOk returns a tuple with the AcConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcConnections

`func (o *EnergyDerRecord) SetAcConnections(v []EnergyDerRecordAcConnections)`

SetAcConnections sets AcConnections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


