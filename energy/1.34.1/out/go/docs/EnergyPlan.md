# EnergyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | The unique identifier for the Energy plan. | 
**EffectiveFrom** | Pointer to **string** | The date and time from which this plan is effective (i.e. is available for origination). Used to enable the articulation of products to the regime before they are available for customers to originate. | [optional] 
**EffectiveTo** | Pointer to **string** | The date and time at which this plan will be retired and will no longer be offered. Used to enable the managed deprecation of plans. | [optional] 
**LastUpdated** | **string** | The last date and time that the information for this plan was changed (or the creation date for the plan if it has never been altered). | 
**DisplayName** | Pointer to **string** | The display name of the plan. | [optional] 
**Description** | Pointer to **string** | A description of the plan. | [optional] 
**Type** | **string** | The type of the plan. | 
**FuelType** | **string** | The fuel types covered by the plan. | 
**Brand** | **string** | The ID of the brand under which this plan is offered. | 
**BrandName** | **string** | The display name of the brand under which this plan is offered. | 
**ApplicationUri** | Pointer to **string** | A link to an application web page where this plan can be applied for. | [optional] 
**AdditionalInformation** | Pointer to [**EnergyPlanAdditionalInformation**](EnergyPlanAdditionalInformation.md) |  | [optional] 
**CustomerType** | Pointer to **string** | The type of customer that the plan is offered to. If absent then the plan is available to all customers. | [optional] 
**Geography** | Pointer to [**EnergyPlanGeography**](EnergyPlanGeography.md) |  | [optional] 

## Methods

### NewEnergyPlan

`func NewEnergyPlan(planId string, lastUpdated string, type_ string, fuelType string, brand string, brandName string, ) *EnergyPlan`

NewEnergyPlan instantiates a new EnergyPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanWithDefaults

`func NewEnergyPlanWithDefaults() *EnergyPlan`

NewEnergyPlanWithDefaults instantiates a new EnergyPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *EnergyPlan) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *EnergyPlan) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *EnergyPlan) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetEffectiveFrom

`func (o *EnergyPlan) GetEffectiveFrom() string`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *EnergyPlan) GetEffectiveFromOk() (*string, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *EnergyPlan) SetEffectiveFrom(v string)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *EnergyPlan) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *EnergyPlan) GetEffectiveTo() string`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *EnergyPlan) GetEffectiveToOk() (*string, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *EnergyPlan) SetEffectiveTo(v string)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *EnergyPlan) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EnergyPlan) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EnergyPlan) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EnergyPlan) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.


### GetDisplayName

`func (o *EnergyPlan) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlan) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlan) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnergyPlan) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *EnergyPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *EnergyPlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlan) SetType(v string)`

SetType sets Type field to given value.


### GetFuelType

`func (o *EnergyPlan) GetFuelType() string`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *EnergyPlan) GetFuelTypeOk() (*string, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *EnergyPlan) SetFuelType(v string)`

SetFuelType sets FuelType field to given value.


### GetBrand

`func (o *EnergyPlan) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *EnergyPlan) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *EnergyPlan) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetBrandName

`func (o *EnergyPlan) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *EnergyPlan) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *EnergyPlan) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.


### GetApplicationUri

`func (o *EnergyPlan) GetApplicationUri() string`

GetApplicationUri returns the ApplicationUri field if non-nil, zero value otherwise.

### GetApplicationUriOk

`func (o *EnergyPlan) GetApplicationUriOk() (*string, bool)`

GetApplicationUriOk returns a tuple with the ApplicationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUri

`func (o *EnergyPlan) SetApplicationUri(v string)`

SetApplicationUri sets ApplicationUri field to given value.

### HasApplicationUri

`func (o *EnergyPlan) HasApplicationUri() bool`

HasApplicationUri returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *EnergyPlan) GetAdditionalInformation() EnergyPlanAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *EnergyPlan) GetAdditionalInformationOk() (*EnergyPlanAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *EnergyPlan) SetAdditionalInformation(v EnergyPlanAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *EnergyPlan) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetCustomerType

`func (o *EnergyPlan) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *EnergyPlan) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *EnergyPlan) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *EnergyPlan) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetGeography

`func (o *EnergyPlan) GetGeography() EnergyPlanGeography`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *EnergyPlan) GetGeographyOk() (*EnergyPlanGeography, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *EnergyPlan) SetGeography(v EnergyPlanGeography)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *EnergyPlan) HasGeography() bool`

HasGeography returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


