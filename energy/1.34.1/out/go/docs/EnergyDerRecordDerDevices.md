# EnergyDerRecordDerDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIdentifier** | **float32** | Unique identifier for a single DER device or a group of DER devices with the same attributes. Does not align with CDR ID permanence standards. | 
**Count** | **int32** | Number of devices in the group of DER devices. | 
**Manufacturer** | Pointer to **string** | The name of the device manufacturer. If absent then assumed to be \&quot;unknown\&quot;. | [optional] 
**ModelNumber** | Pointer to **string** | The model number of the device. If absent then assumed to be \&quot;unknown\&quot;. | [optional] 
**Status** | Pointer to **string** | Code used to indicate the status of the device. This will be used to identify if an inverter is active or inactive or decommissioned. | [optional] 
**Type** | **string** | Used to indicate the primary technology used in the DER device. | 
**Subtype** | Pointer to **string** | Used to indicate the primary technology used in the DER device. This field is also used to record for example the battery chemistry, or the type of PV panel. It is also used to record if a battery is contained in an electric vehicle connected in a vehicle-to-grid arrangement. If absent then assumed to be \&quot;other\&quot;. | [optional] 
**NominalRatedCapacity** | **float32** | Maximum output in kVA that is listed in the product specification by the manufacturer. This refers to the capacity of each unit within the device group. Default is &#x60;0&#x60; if value not known. | 
**NominalStorageCapacity** | Pointer to **float32** | Maximum storage capacity in kVAh. This refers to the capacity of each storage module within the device group. Mandatory if type is equal to &#x60;STORAGE&#x60;. Default is &#x60;0&#x60; if value not known. | [optional] 

## Methods

### NewEnergyDerRecordDerDevices

`func NewEnergyDerRecordDerDevices(deviceIdentifier float32, count int32, type_ string, nominalRatedCapacity float32, ) *EnergyDerRecordDerDevices`

NewEnergyDerRecordDerDevices instantiates a new EnergyDerRecordDerDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyDerRecordDerDevicesWithDefaults

`func NewEnergyDerRecordDerDevicesWithDefaults() *EnergyDerRecordDerDevices`

NewEnergyDerRecordDerDevicesWithDefaults instantiates a new EnergyDerRecordDerDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIdentifier

`func (o *EnergyDerRecordDerDevices) GetDeviceIdentifier() float32`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *EnergyDerRecordDerDevices) GetDeviceIdentifierOk() (*float32, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *EnergyDerRecordDerDevices) SetDeviceIdentifier(v float32)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.


### GetCount

`func (o *EnergyDerRecordDerDevices) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EnergyDerRecordDerDevices) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EnergyDerRecordDerDevices) SetCount(v int32)`

SetCount sets Count field to given value.


### GetManufacturer

`func (o *EnergyDerRecordDerDevices) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *EnergyDerRecordDerDevices) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *EnergyDerRecordDerDevices) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *EnergyDerRecordDerDevices) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModelNumber

`func (o *EnergyDerRecordDerDevices) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *EnergyDerRecordDerDevices) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *EnergyDerRecordDerDevices) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *EnergyDerRecordDerDevices) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetStatus

`func (o *EnergyDerRecordDerDevices) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnergyDerRecordDerDevices) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnergyDerRecordDerDevices) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnergyDerRecordDerDevices) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EnergyDerRecordDerDevices) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyDerRecordDerDevices) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyDerRecordDerDevices) SetType(v string)`

SetType sets Type field to given value.


### GetSubtype

`func (o *EnergyDerRecordDerDevices) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EnergyDerRecordDerDevices) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EnergyDerRecordDerDevices) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *EnergyDerRecordDerDevices) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetNominalRatedCapacity

`func (o *EnergyDerRecordDerDevices) GetNominalRatedCapacity() float32`

GetNominalRatedCapacity returns the NominalRatedCapacity field if non-nil, zero value otherwise.

### GetNominalRatedCapacityOk

`func (o *EnergyDerRecordDerDevices) GetNominalRatedCapacityOk() (*float32, bool)`

GetNominalRatedCapacityOk returns a tuple with the NominalRatedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominalRatedCapacity

`func (o *EnergyDerRecordDerDevices) SetNominalRatedCapacity(v float32)`

SetNominalRatedCapacity sets NominalRatedCapacity field to given value.


### GetNominalStorageCapacity

`func (o *EnergyDerRecordDerDevices) GetNominalStorageCapacity() float32`

GetNominalStorageCapacity returns the NominalStorageCapacity field if non-nil, zero value otherwise.

### GetNominalStorageCapacityOk

`func (o *EnergyDerRecordDerDevices) GetNominalStorageCapacityOk() (*float32, bool)`

GetNominalStorageCapacityOk returns a tuple with the NominalStorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominalStorageCapacity

`func (o *EnergyDerRecordDerDevices) SetNominalStorageCapacity(v float32)`

SetNominalStorageCapacity sets NominalStorageCapacity field to given value.

### HasNominalStorageCapacity

`func (o *EnergyDerRecordDerDevices) HasNominalStorageCapacity() bool`

HasNominalStorageCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


