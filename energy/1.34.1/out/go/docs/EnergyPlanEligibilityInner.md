# EnergyPlanEligibilityInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the eligibility restriction.&lt;br/&gt;The &#x60;CONTINGENT_PLAN&#x60; value indicates that the plan is contingent on the customer taking up an alternate fuel plan from the same retailer (for instance, if the _fuelType_ is &#x60;ELECTRICITY&#x60; then a &#x60;GAS&#x60; plan from the same retailer must be taken up). | 
**Information** | **string** | Information of the eligibility restriction specific to the type of the restriction. | 
**Description** | Pointer to **string** | A description of the eligibility restriction. | [optional] 

## Methods

### NewEnergyPlanEligibilityInner

`func NewEnergyPlanEligibilityInner(type_ string, information string, ) *EnergyPlanEligibilityInner`

NewEnergyPlanEligibilityInner instantiates a new EnergyPlanEligibilityInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanEligibilityInnerWithDefaults

`func NewEnergyPlanEligibilityInnerWithDefaults() *EnergyPlanEligibilityInner`

NewEnergyPlanEligibilityInnerWithDefaults instantiates a new EnergyPlanEligibilityInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnergyPlanEligibilityInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanEligibilityInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanEligibilityInner) SetType(v string)`

SetType sets Type field to given value.


### GetInformation

`func (o *EnergyPlanEligibilityInner) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *EnergyPlanEligibilityInner) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *EnergyPlanEligibilityInner) SetInformation(v string)`

SetInformation sets Information field to given value.


### GetDescription

`func (o *EnergyPlanEligibilityInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanEligibilityInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanEligibilityInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanEligibilityInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


