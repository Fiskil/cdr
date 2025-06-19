# EnergyDerRecordAcConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionIdentifier** | **float32** | AC Connection ID as defined in the DER register. Does not align with CDR ID permanence standards. | 
**Count** | **int32** | Number of AC Connections in the group. For the suite of AC Connections to be considered as a group, all of the AC Connections included must have the same attributes. | 
**EquipmentType** | Pointer to **string** | Indicates whether the DER device is connected via an inverter (and what category of inverter it is) or not (e.g., rotating machine). If absent, assume equipment type to be &#x60;OTHER&#x60;. | [optional] [default to "OTHER"]
**ManufacturerName** | Pointer to **string** | The name of the inverter manufacturer. Mandatory if _equipmentType_ is &#x60;INVERTER&#x60;. | [optional] 
**InverterSeries** | Pointer to **string** | The inverter series. Mandatory if _equipmentType_ is &#x60;INVERTER&#x60;. | [optional] 
**InverterModelNumber** | Pointer to **string** | The inverter model number. Mandatory if _equipmentType_ is &#x60;INVERTER&#x60;. | [optional] 
**CommissioningDate** | **string** | The date that the DER installation is commissioned. | 
**Status** | **string** | Code used to indicate the status of the Inverter. This will be used to identify if an inverter is active or inactive or decommissioned. | 
**InverterDeviceCapacity** | Pointer to **float32** | The rated AC output power that is listed in the product specified by the manufacturer. Mandatory if _equipmentType_ is &#x60;INVERTER&#x60;. Default is &#x60;0&#x60; if value not known. | [optional] 
**DerDevices** | [**[]EnergyDerRecordDerDevices**](EnergyDerRecordDerDevices.md) |  | 

## Methods

### NewEnergyDerRecordAcConnections

`func NewEnergyDerRecordAcConnections(connectionIdentifier float32, count int32, commissioningDate string, status string, derDevices []EnergyDerRecordDerDevices, ) *EnergyDerRecordAcConnections`

NewEnergyDerRecordAcConnections instantiates a new EnergyDerRecordAcConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyDerRecordAcConnectionsWithDefaults

`func NewEnergyDerRecordAcConnectionsWithDefaults() *EnergyDerRecordAcConnections`

NewEnergyDerRecordAcConnectionsWithDefaults instantiates a new EnergyDerRecordAcConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionIdentifier

`func (o *EnergyDerRecordAcConnections) GetConnectionIdentifier() float32`

GetConnectionIdentifier returns the ConnectionIdentifier field if non-nil, zero value otherwise.

### GetConnectionIdentifierOk

`func (o *EnergyDerRecordAcConnections) GetConnectionIdentifierOk() (*float32, bool)`

GetConnectionIdentifierOk returns a tuple with the ConnectionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionIdentifier

`func (o *EnergyDerRecordAcConnections) SetConnectionIdentifier(v float32)`

SetConnectionIdentifier sets ConnectionIdentifier field to given value.


### GetCount

`func (o *EnergyDerRecordAcConnections) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EnergyDerRecordAcConnections) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EnergyDerRecordAcConnections) SetCount(v int32)`

SetCount sets Count field to given value.


### GetEquipmentType

`func (o *EnergyDerRecordAcConnections) GetEquipmentType() string`

GetEquipmentType returns the EquipmentType field if non-nil, zero value otherwise.

### GetEquipmentTypeOk

`func (o *EnergyDerRecordAcConnections) GetEquipmentTypeOk() (*string, bool)`

GetEquipmentTypeOk returns a tuple with the EquipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentType

`func (o *EnergyDerRecordAcConnections) SetEquipmentType(v string)`

SetEquipmentType sets EquipmentType field to given value.

### HasEquipmentType

`func (o *EnergyDerRecordAcConnections) HasEquipmentType() bool`

HasEquipmentType returns a boolean if a field has been set.

### GetManufacturerName

`func (o *EnergyDerRecordAcConnections) GetManufacturerName() string`

GetManufacturerName returns the ManufacturerName field if non-nil, zero value otherwise.

### GetManufacturerNameOk

`func (o *EnergyDerRecordAcConnections) GetManufacturerNameOk() (*string, bool)`

GetManufacturerNameOk returns a tuple with the ManufacturerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerName

`func (o *EnergyDerRecordAcConnections) SetManufacturerName(v string)`

SetManufacturerName sets ManufacturerName field to given value.

### HasManufacturerName

`func (o *EnergyDerRecordAcConnections) HasManufacturerName() bool`

HasManufacturerName returns a boolean if a field has been set.

### GetInverterSeries

`func (o *EnergyDerRecordAcConnections) GetInverterSeries() string`

GetInverterSeries returns the InverterSeries field if non-nil, zero value otherwise.

### GetInverterSeriesOk

`func (o *EnergyDerRecordAcConnections) GetInverterSeriesOk() (*string, bool)`

GetInverterSeriesOk returns a tuple with the InverterSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverterSeries

`func (o *EnergyDerRecordAcConnections) SetInverterSeries(v string)`

SetInverterSeries sets InverterSeries field to given value.

### HasInverterSeries

`func (o *EnergyDerRecordAcConnections) HasInverterSeries() bool`

HasInverterSeries returns a boolean if a field has been set.

### GetInverterModelNumber

`func (o *EnergyDerRecordAcConnections) GetInverterModelNumber() string`

GetInverterModelNumber returns the InverterModelNumber field if non-nil, zero value otherwise.

### GetInverterModelNumberOk

`func (o *EnergyDerRecordAcConnections) GetInverterModelNumberOk() (*string, bool)`

GetInverterModelNumberOk returns a tuple with the InverterModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverterModelNumber

`func (o *EnergyDerRecordAcConnections) SetInverterModelNumber(v string)`

SetInverterModelNumber sets InverterModelNumber field to given value.

### HasInverterModelNumber

`func (o *EnergyDerRecordAcConnections) HasInverterModelNumber() bool`

HasInverterModelNumber returns a boolean if a field has been set.

### GetCommissioningDate

`func (o *EnergyDerRecordAcConnections) GetCommissioningDate() string`

GetCommissioningDate returns the CommissioningDate field if non-nil, zero value otherwise.

### GetCommissioningDateOk

`func (o *EnergyDerRecordAcConnections) GetCommissioningDateOk() (*string, bool)`

GetCommissioningDateOk returns a tuple with the CommissioningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissioningDate

`func (o *EnergyDerRecordAcConnections) SetCommissioningDate(v string)`

SetCommissioningDate sets CommissioningDate field to given value.


### GetStatus

`func (o *EnergyDerRecordAcConnections) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnergyDerRecordAcConnections) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnergyDerRecordAcConnections) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInverterDeviceCapacity

`func (o *EnergyDerRecordAcConnections) GetInverterDeviceCapacity() float32`

GetInverterDeviceCapacity returns the InverterDeviceCapacity field if non-nil, zero value otherwise.

### GetInverterDeviceCapacityOk

`func (o *EnergyDerRecordAcConnections) GetInverterDeviceCapacityOk() (*float32, bool)`

GetInverterDeviceCapacityOk returns a tuple with the InverterDeviceCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverterDeviceCapacity

`func (o *EnergyDerRecordAcConnections) SetInverterDeviceCapacity(v float32)`

SetInverterDeviceCapacity sets InverterDeviceCapacity field to given value.

### HasInverterDeviceCapacity

`func (o *EnergyDerRecordAcConnections) HasInverterDeviceCapacity() bool`

HasInverterDeviceCapacity returns a boolean if a field has been set.

### GetDerDevices

`func (o *EnergyDerRecordAcConnections) GetDerDevices() []EnergyDerRecordDerDevices`

GetDerDevices returns the DerDevices field if non-nil, zero value otherwise.

### GetDerDevicesOk

`func (o *EnergyDerRecordAcConnections) GetDerDevicesOk() (*[]EnergyDerRecordDerDevices, bool)`

GetDerDevicesOk returns a tuple with the DerDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerDevices

`func (o *EnergyDerRecordAcConnections) SetDerDevices(v []EnergyDerRecordDerDevices)`

SetDerDevices sets DerDevices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


