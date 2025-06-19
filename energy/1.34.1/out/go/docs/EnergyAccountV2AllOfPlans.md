# EnergyAccountV2AllOfPlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nickname** | Pointer to **string** | Optional display name for the plan provided by the customer to help differentiate multiple plans. | [optional] 
**ServicePointIds** | **[]string** | An array of _servicePointId_ values, representing NMIs, that this plan is linked to. If there are no service points allocated to this plan then an empty array would be expected. | 
**PlanOverview** | Pointer to [**EnergyAccountV2AllOfPlanOverview**](EnergyAccountV2AllOfPlanOverview.md) |  | [optional] 

## Methods

### NewEnergyAccountV2AllOfPlans

`func NewEnergyAccountV2AllOfPlans(servicePointIds []string, ) *EnergyAccountV2AllOfPlans`

NewEnergyAccountV2AllOfPlans instantiates a new EnergyAccountV2AllOfPlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyAccountV2AllOfPlansWithDefaults

`func NewEnergyAccountV2AllOfPlansWithDefaults() *EnergyAccountV2AllOfPlans`

NewEnergyAccountV2AllOfPlansWithDefaults instantiates a new EnergyAccountV2AllOfPlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNickname

`func (o *EnergyAccountV2AllOfPlans) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *EnergyAccountV2AllOfPlans) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *EnergyAccountV2AllOfPlans) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *EnergyAccountV2AllOfPlans) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetServicePointIds

`func (o *EnergyAccountV2AllOfPlans) GetServicePointIds() []string`

GetServicePointIds returns the ServicePointIds field if non-nil, zero value otherwise.

### GetServicePointIdsOk

`func (o *EnergyAccountV2AllOfPlans) GetServicePointIdsOk() (*[]string, bool)`

GetServicePointIdsOk returns a tuple with the ServicePointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointIds

`func (o *EnergyAccountV2AllOfPlans) SetServicePointIds(v []string)`

SetServicePointIds sets ServicePointIds field to given value.


### GetPlanOverview

`func (o *EnergyAccountV2AllOfPlans) GetPlanOverview() EnergyAccountV2AllOfPlanOverview`

GetPlanOverview returns the PlanOverview field if non-nil, zero value otherwise.

### GetPlanOverviewOk

`func (o *EnergyAccountV2AllOfPlans) GetPlanOverviewOk() (*EnergyAccountV2AllOfPlanOverview, bool)`

GetPlanOverviewOk returns a tuple with the PlanOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanOverview

`func (o *EnergyAccountV2AllOfPlans) SetPlanOverview(v EnergyAccountV2AllOfPlanOverview)`

SetPlanOverview sets PlanOverview field to given value.

### HasPlanOverview

`func (o *EnergyAccountV2AllOfPlans) HasPlanOverview() bool`

HasPlanOverview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


