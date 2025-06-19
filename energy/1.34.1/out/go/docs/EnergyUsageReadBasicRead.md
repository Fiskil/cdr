# EnergyUsageReadBasicRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quality** | Pointer to **string** | The quality of the read taken. If absent then assumed to be &#x60;ACTUAL&#x60;. | [optional] [default to "ACTUAL"]
**Value** | **float32** | Meter read value. If positive then it means consumption, if negative it means export. | 

## Methods

### NewEnergyUsageReadBasicRead

`func NewEnergyUsageReadBasicRead(value float32, ) *EnergyUsageReadBasicRead`

NewEnergyUsageReadBasicRead instantiates a new EnergyUsageReadBasicRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyUsageReadBasicReadWithDefaults

`func NewEnergyUsageReadBasicReadWithDefaults() *EnergyUsageReadBasicRead`

NewEnergyUsageReadBasicReadWithDefaults instantiates a new EnergyUsageReadBasicRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuality

`func (o *EnergyUsageReadBasicRead) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *EnergyUsageReadBasicRead) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *EnergyUsageReadBasicRead) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *EnergyUsageReadBasicRead) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetValue

`func (o *EnergyUsageReadBasicRead) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnergyUsageReadBasicRead) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnergyUsageReadBasicRead) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


