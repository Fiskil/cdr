# EnergyConcessionsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concessions** | [**[]EnergyConcession**](EnergyConcession.md) | Array may be empty if no concessions exist. | 

## Methods

### NewEnergyConcessionsResponseData

`func NewEnergyConcessionsResponseData(concessions []EnergyConcession, ) *EnergyConcessionsResponseData`

NewEnergyConcessionsResponseData instantiates a new EnergyConcessionsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyConcessionsResponseDataWithDefaults

`func NewEnergyConcessionsResponseDataWithDefaults() *EnergyConcessionsResponseData`

NewEnergyConcessionsResponseDataWithDefaults instantiates a new EnergyConcessionsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcessions

`func (o *EnergyConcessionsResponseData) GetConcessions() []EnergyConcession`

GetConcessions returns the Concessions field if non-nil, zero value otherwise.

### GetConcessionsOk

`func (o *EnergyConcessionsResponseData) GetConcessionsOk() (*[]EnergyConcession, bool)`

GetConcessionsOk returns a tuple with the Concessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcessions

`func (o *EnergyConcessionsResponseData) SetConcessions(v []EnergyConcession)`

SetConcessions sets Concessions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


