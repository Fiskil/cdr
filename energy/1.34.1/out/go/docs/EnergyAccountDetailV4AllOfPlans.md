# EnergyAccountDetailV4AllOfPlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nickname** | Pointer to **string** | Optional display name for the plan provided by the customer to help differentiate multiple plans. | [optional] 
**ServicePointIds** | **[]string** | An array of _servicePointId_ values, representing NMIs, that this account is linked to. | 
**PlanOverview** | Pointer to [**EnergyAccountV2AllOfPlanOverview**](EnergyAccountV2AllOfPlanOverview.md) |  | [optional] 
**PlanDetail** | Pointer to [**EnergyAccountDetailV4AllOfPlanDetail**](EnergyAccountDetailV4AllOfPlanDetail.md) |  | [optional] 
**AuthorisedContacts** | Pointer to [**[]EnergyAccountDetailV4AllOfAuthorisedContacts**](EnergyAccountDetailV4AllOfAuthorisedContacts.md) | An array of additional contacts that are authorised to act on this account. | [optional] 

## Methods

### NewEnergyAccountDetailV4AllOfPlans

`func NewEnergyAccountDetailV4AllOfPlans(servicePointIds []string, ) *EnergyAccountDetailV4AllOfPlans`

NewEnergyAccountDetailV4AllOfPlans instantiates a new EnergyAccountDetailV4AllOfPlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyAccountDetailV4AllOfPlansWithDefaults

`func NewEnergyAccountDetailV4AllOfPlansWithDefaults() *EnergyAccountDetailV4AllOfPlans`

NewEnergyAccountDetailV4AllOfPlansWithDefaults instantiates a new EnergyAccountDetailV4AllOfPlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNickname

`func (o *EnergyAccountDetailV4AllOfPlans) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *EnergyAccountDetailV4AllOfPlans) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *EnergyAccountDetailV4AllOfPlans) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *EnergyAccountDetailV4AllOfPlans) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetServicePointIds

`func (o *EnergyAccountDetailV4AllOfPlans) GetServicePointIds() []string`

GetServicePointIds returns the ServicePointIds field if non-nil, zero value otherwise.

### GetServicePointIdsOk

`func (o *EnergyAccountDetailV4AllOfPlans) GetServicePointIdsOk() (*[]string, bool)`

GetServicePointIdsOk returns a tuple with the ServicePointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointIds

`func (o *EnergyAccountDetailV4AllOfPlans) SetServicePointIds(v []string)`

SetServicePointIds sets ServicePointIds field to given value.


### GetPlanOverview

`func (o *EnergyAccountDetailV4AllOfPlans) GetPlanOverview() EnergyAccountV2AllOfPlanOverview`

GetPlanOverview returns the PlanOverview field if non-nil, zero value otherwise.

### GetPlanOverviewOk

`func (o *EnergyAccountDetailV4AllOfPlans) GetPlanOverviewOk() (*EnergyAccountV2AllOfPlanOverview, bool)`

GetPlanOverviewOk returns a tuple with the PlanOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanOverview

`func (o *EnergyAccountDetailV4AllOfPlans) SetPlanOverview(v EnergyAccountV2AllOfPlanOverview)`

SetPlanOverview sets PlanOverview field to given value.

### HasPlanOverview

`func (o *EnergyAccountDetailV4AllOfPlans) HasPlanOverview() bool`

HasPlanOverview returns a boolean if a field has been set.

### GetPlanDetail

`func (o *EnergyAccountDetailV4AllOfPlans) GetPlanDetail() EnergyAccountDetailV4AllOfPlanDetail`

GetPlanDetail returns the PlanDetail field if non-nil, zero value otherwise.

### GetPlanDetailOk

`func (o *EnergyAccountDetailV4AllOfPlans) GetPlanDetailOk() (*EnergyAccountDetailV4AllOfPlanDetail, bool)`

GetPlanDetailOk returns a tuple with the PlanDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDetail

`func (o *EnergyAccountDetailV4AllOfPlans) SetPlanDetail(v EnergyAccountDetailV4AllOfPlanDetail)`

SetPlanDetail sets PlanDetail field to given value.

### HasPlanDetail

`func (o *EnergyAccountDetailV4AllOfPlans) HasPlanDetail() bool`

HasPlanDetail returns a boolean if a field has been set.

### GetAuthorisedContacts

`func (o *EnergyAccountDetailV4AllOfPlans) GetAuthorisedContacts() []EnergyAccountDetailV4AllOfAuthorisedContacts`

GetAuthorisedContacts returns the AuthorisedContacts field if non-nil, zero value otherwise.

### GetAuthorisedContactsOk

`func (o *EnergyAccountDetailV4AllOfPlans) GetAuthorisedContactsOk() (*[]EnergyAccountDetailV4AllOfAuthorisedContacts, bool)`

GetAuthorisedContactsOk returns a tuple with the AuthorisedContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisedContacts

`func (o *EnergyAccountDetailV4AllOfPlans) SetAuthorisedContacts(v []EnergyAccountDetailV4AllOfAuthorisedContacts)`

SetAuthorisedContacts sets AuthorisedContacts field to given value.

### HasAuthorisedContacts

`func (o *EnergyAccountDetailV4AllOfPlans) HasAuthorisedContacts() bool`

HasAuthorisedContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


