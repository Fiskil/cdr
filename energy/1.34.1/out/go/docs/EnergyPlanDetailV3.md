# EnergyPlanDetailV3

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
**MeteringCharges** | Pointer to [**[]EnergyPlanDetailV3AllOfMeteringCharges**](EnergyPlanDetailV3AllOfMeteringCharges.md) | Charges for metering included in the plan. | [optional] 
**GasContract** | Pointer to [**EnergyPlanContractFullV3**](EnergyPlanContractFullV3.md) | The details of the terms for the supply of electricity under this plan. Is mandatory if _fuelType_ is set to &#x60;GAS&#x60; or &#x60;DUAL&#x60;. | [optional] 
**ElectricityContract** | Pointer to [**EnergyPlanContractFullV3**](EnergyPlanContractFullV3.md) | The details of the terms for the supply of electricity under this plan. Is mandatory if _fuelType_ is set to &#x60;ELECTRICITY&#x60; or &#x60;DUAL&#x60;. | [optional] 

## Methods

### NewEnergyPlanDetailV3

`func NewEnergyPlanDetailV3(planId string, lastUpdated string, type_ string, fuelType string, brand string, brandName string, ) *EnergyPlanDetailV3`

NewEnergyPlanDetailV3 instantiates a new EnergyPlanDetailV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanDetailV3WithDefaults

`func NewEnergyPlanDetailV3WithDefaults() *EnergyPlanDetailV3`

NewEnergyPlanDetailV3WithDefaults instantiates a new EnergyPlanDetailV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *EnergyPlanDetailV3) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *EnergyPlanDetailV3) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *EnergyPlanDetailV3) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetEffectiveFrom

`func (o *EnergyPlanDetailV3) GetEffectiveFrom() string`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *EnergyPlanDetailV3) GetEffectiveFromOk() (*string, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *EnergyPlanDetailV3) SetEffectiveFrom(v string)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *EnergyPlanDetailV3) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *EnergyPlanDetailV3) GetEffectiveTo() string`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *EnergyPlanDetailV3) GetEffectiveToOk() (*string, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *EnergyPlanDetailV3) SetEffectiveTo(v string)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *EnergyPlanDetailV3) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EnergyPlanDetailV3) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EnergyPlanDetailV3) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EnergyPlanDetailV3) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.


### GetDisplayName

`func (o *EnergyPlanDetailV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanDetailV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanDetailV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnergyPlanDetailV3) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *EnergyPlanDetailV3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanDetailV3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanDetailV3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanDetailV3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *EnergyPlanDetailV3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanDetailV3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanDetailV3) SetType(v string)`

SetType sets Type field to given value.


### GetFuelType

`func (o *EnergyPlanDetailV3) GetFuelType() string`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *EnergyPlanDetailV3) GetFuelTypeOk() (*string, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *EnergyPlanDetailV3) SetFuelType(v string)`

SetFuelType sets FuelType field to given value.


### GetBrand

`func (o *EnergyPlanDetailV3) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *EnergyPlanDetailV3) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *EnergyPlanDetailV3) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetBrandName

`func (o *EnergyPlanDetailV3) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *EnergyPlanDetailV3) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *EnergyPlanDetailV3) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.


### GetApplicationUri

`func (o *EnergyPlanDetailV3) GetApplicationUri() string`

GetApplicationUri returns the ApplicationUri field if non-nil, zero value otherwise.

### GetApplicationUriOk

`func (o *EnergyPlanDetailV3) GetApplicationUriOk() (*string, bool)`

GetApplicationUriOk returns a tuple with the ApplicationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUri

`func (o *EnergyPlanDetailV3) SetApplicationUri(v string)`

SetApplicationUri sets ApplicationUri field to given value.

### HasApplicationUri

`func (o *EnergyPlanDetailV3) HasApplicationUri() bool`

HasApplicationUri returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *EnergyPlanDetailV3) GetAdditionalInformation() EnergyPlanAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *EnergyPlanDetailV3) GetAdditionalInformationOk() (*EnergyPlanAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *EnergyPlanDetailV3) SetAdditionalInformation(v EnergyPlanAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *EnergyPlanDetailV3) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetCustomerType

`func (o *EnergyPlanDetailV3) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *EnergyPlanDetailV3) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *EnergyPlanDetailV3) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *EnergyPlanDetailV3) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetGeography

`func (o *EnergyPlanDetailV3) GetGeography() EnergyPlanGeography`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *EnergyPlanDetailV3) GetGeographyOk() (*EnergyPlanGeography, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *EnergyPlanDetailV3) SetGeography(v EnergyPlanGeography)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *EnergyPlanDetailV3) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### GetMeteringCharges

`func (o *EnergyPlanDetailV3) GetMeteringCharges() []EnergyPlanDetailV3AllOfMeteringCharges`

GetMeteringCharges returns the MeteringCharges field if non-nil, zero value otherwise.

### GetMeteringChargesOk

`func (o *EnergyPlanDetailV3) GetMeteringChargesOk() (*[]EnergyPlanDetailV3AllOfMeteringCharges, bool)`

GetMeteringChargesOk returns a tuple with the MeteringCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringCharges

`func (o *EnergyPlanDetailV3) SetMeteringCharges(v []EnergyPlanDetailV3AllOfMeteringCharges)`

SetMeteringCharges sets MeteringCharges field to given value.

### HasMeteringCharges

`func (o *EnergyPlanDetailV3) HasMeteringCharges() bool`

HasMeteringCharges returns a boolean if a field has been set.

### GetGasContract

`func (o *EnergyPlanDetailV3) GetGasContract() EnergyPlanContractFullV3`

GetGasContract returns the GasContract field if non-nil, zero value otherwise.

### GetGasContractOk

`func (o *EnergyPlanDetailV3) GetGasContractOk() (*EnergyPlanContractFullV3, bool)`

GetGasContractOk returns a tuple with the GasContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasContract

`func (o *EnergyPlanDetailV3) SetGasContract(v EnergyPlanContractFullV3)`

SetGasContract sets GasContract field to given value.

### HasGasContract

`func (o *EnergyPlanDetailV3) HasGasContract() bool`

HasGasContract returns a boolean if a field has been set.

### GetElectricityContract

`func (o *EnergyPlanDetailV3) GetElectricityContract() EnergyPlanContractFullV3`

GetElectricityContract returns the ElectricityContract field if non-nil, zero value otherwise.

### GetElectricityContractOk

`func (o *EnergyPlanDetailV3) GetElectricityContractOk() (*EnergyPlanContractFullV3, bool)`

GetElectricityContractOk returns a tuple with the ElectricityContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityContract

`func (o *EnergyPlanDetailV3) SetElectricityContract(v EnergyPlanContractFullV3)`

SetElectricityContract sets ElectricityContract field to given value.

### HasElectricityContract

`func (o *EnergyPlanDetailV3) HasElectricityContract() bool`

HasElectricityContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


