# EnergyDerListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerRecords** | [**[]EnergyDerRecord**](EnergyDerRecord.md) | Array of meter reads. | 

## Methods

### NewEnergyDerListResponseData

`func NewEnergyDerListResponseData(derRecords []EnergyDerRecord, ) *EnergyDerListResponseData`

NewEnergyDerListResponseData instantiates a new EnergyDerListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyDerListResponseDataWithDefaults

`func NewEnergyDerListResponseDataWithDefaults() *EnergyDerListResponseData`

NewEnergyDerListResponseDataWithDefaults instantiates a new EnergyDerListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerRecords

`func (o *EnergyDerListResponseData) GetDerRecords() []EnergyDerRecord`

GetDerRecords returns the DerRecords field if non-nil, zero value otherwise.

### GetDerRecordsOk

`func (o *EnergyDerListResponseData) GetDerRecordsOk() (*[]EnergyDerRecord, bool)`

GetDerRecordsOk returns a tuple with the DerRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerRecords

`func (o *EnergyDerListResponseData) SetDerRecords(v []EnergyDerRecord)`

SetDerRecords sets DerRecords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


