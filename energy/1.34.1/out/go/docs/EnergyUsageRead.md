# EnergyUsageRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePointId** | **string** | Unique identifier for the service point. | 
**RegisterId** | Pointer to **string** | Register ID of the meter register where the meter reads are obtained. | [optional] 
**RegisterSuffix** | **string** | Register suffix of the meter register where the meter reads are obtained. | 
**MeterId** | Pointer to **string** | Meter id/serial number as it appears in customer&#39;s bill. ID permanence rules do not apply. | [optional] 
**ControlledLoad** | Pointer to **bool** | Indicates whether the energy recorded by this register is created under a Controlled Load regime. | [optional] 
**ReadStartDate** | **string** | Date when the meter reads start in AEST and assumed to start from 12:00am AEST. | 
**ReadEndDate** | Pointer to **string** | Date when the meter reads end in AEST. If absent then assumed to be equal to _readStartDate_. In this case the entry represents data for a single date specified by _readStartDate_. | [optional] 
**UnitOfMeasure** | Pointer to **string** | Unit of measure of the meter reads. Refer to Appendix B of &lt;a href&#x3D;&#39;https://www.aemo.com.au/-/media/files/stakeholder_consultation/consultations/nem-consultations/2019/5ms-metering-package-2/final-determination/mdff-specification-nem12-nem13-v21-final-determination-clean.pdf?la&#x3D;en&amp;hash&#x3D;03FCBA0D60E091DE00F2361AE76206EA&#39;&gt;MDFF Specification NEM12 NEM13 v2.1&lt;/a&gt; for a list of possible values. | [optional] 
**ReadUType** | **string** | Specify the type of the meter read data. | 
**BasicRead** | Pointer to [**EnergyUsageReadBasicRead**](EnergyUsageReadBasicRead.md) |  | [optional] 
**IntervalRead** | Pointer to [**EnergyUsageReadIntervalRead**](EnergyUsageReadIntervalRead.md) |  | [optional] 

## Methods

### NewEnergyUsageRead

`func NewEnergyUsageRead(servicePointId string, registerSuffix string, readStartDate string, readUType string, ) *EnergyUsageRead`

NewEnergyUsageRead instantiates a new EnergyUsageRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyUsageReadWithDefaults

`func NewEnergyUsageReadWithDefaults() *EnergyUsageRead`

NewEnergyUsageReadWithDefaults instantiates a new EnergyUsageRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyUsageRead) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyUsageRead) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyUsageRead) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.


### GetRegisterId

`func (o *EnergyUsageRead) GetRegisterId() string`

GetRegisterId returns the RegisterId field if non-nil, zero value otherwise.

### GetRegisterIdOk

`func (o *EnergyUsageRead) GetRegisterIdOk() (*string, bool)`

GetRegisterIdOk returns a tuple with the RegisterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterId

`func (o *EnergyUsageRead) SetRegisterId(v string)`

SetRegisterId sets RegisterId field to given value.

### HasRegisterId

`func (o *EnergyUsageRead) HasRegisterId() bool`

HasRegisterId returns a boolean if a field has been set.

### GetRegisterSuffix

`func (o *EnergyUsageRead) GetRegisterSuffix() string`

GetRegisterSuffix returns the RegisterSuffix field if non-nil, zero value otherwise.

### GetRegisterSuffixOk

`func (o *EnergyUsageRead) GetRegisterSuffixOk() (*string, bool)`

GetRegisterSuffixOk returns a tuple with the RegisterSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterSuffix

`func (o *EnergyUsageRead) SetRegisterSuffix(v string)`

SetRegisterSuffix sets RegisterSuffix field to given value.


### GetMeterId

`func (o *EnergyUsageRead) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *EnergyUsageRead) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *EnergyUsageRead) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *EnergyUsageRead) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetControlledLoad

`func (o *EnergyUsageRead) GetControlledLoad() bool`

GetControlledLoad returns the ControlledLoad field if non-nil, zero value otherwise.

### GetControlledLoadOk

`func (o *EnergyUsageRead) GetControlledLoadOk() (*bool, bool)`

GetControlledLoadOk returns a tuple with the ControlledLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlledLoad

`func (o *EnergyUsageRead) SetControlledLoad(v bool)`

SetControlledLoad sets ControlledLoad field to given value.

### HasControlledLoad

`func (o *EnergyUsageRead) HasControlledLoad() bool`

HasControlledLoad returns a boolean if a field has been set.

### GetReadStartDate

`func (o *EnergyUsageRead) GetReadStartDate() string`

GetReadStartDate returns the ReadStartDate field if non-nil, zero value otherwise.

### GetReadStartDateOk

`func (o *EnergyUsageRead) GetReadStartDateOk() (*string, bool)`

GetReadStartDateOk returns a tuple with the ReadStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadStartDate

`func (o *EnergyUsageRead) SetReadStartDate(v string)`

SetReadStartDate sets ReadStartDate field to given value.


### GetReadEndDate

`func (o *EnergyUsageRead) GetReadEndDate() string`

GetReadEndDate returns the ReadEndDate field if non-nil, zero value otherwise.

### GetReadEndDateOk

`func (o *EnergyUsageRead) GetReadEndDateOk() (*string, bool)`

GetReadEndDateOk returns a tuple with the ReadEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadEndDate

`func (o *EnergyUsageRead) SetReadEndDate(v string)`

SetReadEndDate sets ReadEndDate field to given value.

### HasReadEndDate

`func (o *EnergyUsageRead) HasReadEndDate() bool`

HasReadEndDate returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *EnergyUsageRead) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *EnergyUsageRead) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *EnergyUsageRead) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *EnergyUsageRead) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetReadUType

`func (o *EnergyUsageRead) GetReadUType() string`

GetReadUType returns the ReadUType field if non-nil, zero value otherwise.

### GetReadUTypeOk

`func (o *EnergyUsageRead) GetReadUTypeOk() (*string, bool)`

GetReadUTypeOk returns a tuple with the ReadUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadUType

`func (o *EnergyUsageRead) SetReadUType(v string)`

SetReadUType sets ReadUType field to given value.


### GetBasicRead

`func (o *EnergyUsageRead) GetBasicRead() EnergyUsageReadBasicRead`

GetBasicRead returns the BasicRead field if non-nil, zero value otherwise.

### GetBasicReadOk

`func (o *EnergyUsageRead) GetBasicReadOk() (*EnergyUsageReadBasicRead, bool)`

GetBasicReadOk returns a tuple with the BasicRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicRead

`func (o *EnergyUsageRead) SetBasicRead(v EnergyUsageReadBasicRead)`

SetBasicRead sets BasicRead field to given value.

### HasBasicRead

`func (o *EnergyUsageRead) HasBasicRead() bool`

HasBasicRead returns a boolean if a field has been set.

### GetIntervalRead

`func (o *EnergyUsageRead) GetIntervalRead() EnergyUsageReadIntervalRead`

GetIntervalRead returns the IntervalRead field if non-nil, zero value otherwise.

### GetIntervalReadOk

`func (o *EnergyUsageRead) GetIntervalReadOk() (*EnergyUsageReadIntervalRead, bool)`

GetIntervalReadOk returns a tuple with the IntervalRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalRead

`func (o *EnergyUsageRead) SetIntervalRead(v EnergyUsageReadIntervalRead)`

SetIntervalRead sets IntervalRead field to given value.

### HasIntervalRead

`func (o *EnergyUsageRead) HasIntervalRead() bool`

HasIntervalRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


