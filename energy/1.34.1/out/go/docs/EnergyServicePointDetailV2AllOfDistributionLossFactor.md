# EnergyServicePointDetailV2AllOfDistributionLossFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code used to identify data loss factor for the service point values. Refer to AEMO distribution loss factor documents for each financial year to interpret. | 
**Description** | **string** | Description of the data loss factor code and value. | 
**LossValue** | **string** | The value associated with the loss factor code. | 

## Methods

### NewEnergyServicePointDetailV2AllOfDistributionLossFactor

`func NewEnergyServicePointDetailV2AllOfDistributionLossFactor(code string, description string, lossValue string, ) *EnergyServicePointDetailV2AllOfDistributionLossFactor`

NewEnergyServicePointDetailV2AllOfDistributionLossFactor instantiates a new EnergyServicePointDetailV2AllOfDistributionLossFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointDetailV2AllOfDistributionLossFactorWithDefaults

`func NewEnergyServicePointDetailV2AllOfDistributionLossFactorWithDefaults() *EnergyServicePointDetailV2AllOfDistributionLossFactor`

NewEnergyServicePointDetailV2AllOfDistributionLossFactorWithDefaults instantiates a new EnergyServicePointDetailV2AllOfDistributionLossFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLossValue

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) GetLossValue() string`

GetLossValue returns the LossValue field if non-nil, zero value otherwise.

### GetLossValueOk

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) GetLossValueOk() (*string, bool)`

GetLossValueOk returns a tuple with the LossValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossValue

`func (o *EnergyServicePointDetailV2AllOfDistributionLossFactor) SetLossValue(v string)`

SetLossValue sets LossValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


