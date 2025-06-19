# EnergyUsageListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reads** | [**[]EnergyUsageRead**](EnergyUsageRead.md) | Array of meter reads sorted by NMI in ascending order followed by _readStartDate_ in descending order. | 

## Methods

### NewEnergyUsageListResponseData

`func NewEnergyUsageListResponseData(reads []EnergyUsageRead, ) *EnergyUsageListResponseData`

NewEnergyUsageListResponseData instantiates a new EnergyUsageListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyUsageListResponseDataWithDefaults

`func NewEnergyUsageListResponseDataWithDefaults() *EnergyUsageListResponseData`

NewEnergyUsageListResponseDataWithDefaults instantiates a new EnergyUsageListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReads

`func (o *EnergyUsageListResponseData) GetReads() []EnergyUsageRead`

GetReads returns the Reads field if non-nil, zero value otherwise.

### GetReadsOk

`func (o *EnergyUsageListResponseData) GetReadsOk() (*[]EnergyUsageRead, bool)`

GetReadsOk returns a tuple with the Reads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReads

`func (o *EnergyUsageListResponseData) SetReads(v []EnergyUsageRead)`

SetReads sets Reads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


