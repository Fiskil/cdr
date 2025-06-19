# EnergyPlanGeography

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedPostcodes** | Pointer to **[]string** | Array of valid Australian postcodes that are specifically excluded from the plan. Each element is a single four digit postcode (e.g., &#x60;3000&#x60;) or a range of postcodes defined by two four digit postcodes and a hyphen (e.g., &#x60;3000-3999&#x60;). | [optional] 
**IncludedPostcodes** | Pointer to **[]string** | Array of valid Australian postcodes that are included from the plan. If absent defaults to all non-excluded postcodes. Each element is a single four digit postcode (e.g., &#x60;3000&#x60;) or a range of postcodes defined by two four digit postcodes and a hyphen (e.g., &#x60;3000-3999&#x60;). | [optional] 
**Distributors** | **[]string** | Array of distributors for the plan. Must have at least one entry. | 

## Methods

### NewEnergyPlanGeography

`func NewEnergyPlanGeography(distributors []string, ) *EnergyPlanGeography`

NewEnergyPlanGeography instantiates a new EnergyPlanGeography object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanGeographyWithDefaults

`func NewEnergyPlanGeographyWithDefaults() *EnergyPlanGeography`

NewEnergyPlanGeographyWithDefaults instantiates a new EnergyPlanGeography object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedPostcodes

`func (o *EnergyPlanGeography) GetExcludedPostcodes() []string`

GetExcludedPostcodes returns the ExcludedPostcodes field if non-nil, zero value otherwise.

### GetExcludedPostcodesOk

`func (o *EnergyPlanGeography) GetExcludedPostcodesOk() (*[]string, bool)`

GetExcludedPostcodesOk returns a tuple with the ExcludedPostcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPostcodes

`func (o *EnergyPlanGeography) SetExcludedPostcodes(v []string)`

SetExcludedPostcodes sets ExcludedPostcodes field to given value.

### HasExcludedPostcodes

`func (o *EnergyPlanGeography) HasExcludedPostcodes() bool`

HasExcludedPostcodes returns a boolean if a field has been set.

### GetIncludedPostcodes

`func (o *EnergyPlanGeography) GetIncludedPostcodes() []string`

GetIncludedPostcodes returns the IncludedPostcodes field if non-nil, zero value otherwise.

### GetIncludedPostcodesOk

`func (o *EnergyPlanGeography) GetIncludedPostcodesOk() (*[]string, bool)`

GetIncludedPostcodesOk returns a tuple with the IncludedPostcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedPostcodes

`func (o *EnergyPlanGeography) SetIncludedPostcodes(v []string)`

SetIncludedPostcodes sets IncludedPostcodes field to given value.

### HasIncludedPostcodes

`func (o *EnergyPlanGeography) HasIncludedPostcodes() bool`

HasIncludedPostcodes returns a boolean if a field has been set.

### GetDistributors

`func (o *EnergyPlanGeography) GetDistributors() []string`

GetDistributors returns the Distributors field if non-nil, zero value otherwise.

### GetDistributorsOk

`func (o *EnergyPlanGeography) GetDistributorsOk() (*[]string, bool)`

GetDistributorsOk returns a tuple with the Distributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributors

`func (o *EnergyPlanGeography) SetDistributors(v []string)`

SetDistributors sets Distributors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


