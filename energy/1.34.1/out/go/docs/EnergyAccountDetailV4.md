# EnergyAccountDetailV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**AccountNumber** | Pointer to **string** | Optional identifier of the account as defined by the data holder. This must be the value presented on physical statements (if it exists) and must not be used for the value of _accountId_. | [optional] 
**DisplayName** | Pointer to **string** | An optional display name for the account if one exists or can be derived. The content of this field is at the discretion of the data holder. | [optional] 
**OpenStatus** | Pointer to **string** | Open or closed status for the account. If not present then &#x60;OPEN&#x60; is assumed. | [optional] [default to "OPEN"]
**CreationDate** | Pointer to **string** | The date that the account was created or opened. Mandatory if _openStatus_ is &#x60;OPEN&#x60;. | [optional] 
**Plans** | [**[]EnergyAccountDetailV4AllOfPlans**](EnergyAccountDetailV4AllOfPlans.md) | The array of plans containing service points and associated plan details. | 

## Methods

### NewEnergyAccountDetailV4

`func NewEnergyAccountDetailV4(accountId string, plans []EnergyAccountDetailV4AllOfPlans, ) *EnergyAccountDetailV4`

NewEnergyAccountDetailV4 instantiates a new EnergyAccountDetailV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyAccountDetailV4WithDefaults

`func NewEnergyAccountDetailV4WithDefaults() *EnergyAccountDetailV4`

NewEnergyAccountDetailV4WithDefaults instantiates a new EnergyAccountDetailV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EnergyAccountDetailV4) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EnergyAccountDetailV4) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EnergyAccountDetailV4) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountNumber

`func (o *EnergyAccountDetailV4) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *EnergyAccountDetailV4) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *EnergyAccountDetailV4) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *EnergyAccountDetailV4) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnergyAccountDetailV4) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyAccountDetailV4) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyAccountDetailV4) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnergyAccountDetailV4) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOpenStatus

`func (o *EnergyAccountDetailV4) GetOpenStatus() string`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *EnergyAccountDetailV4) GetOpenStatusOk() (*string, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *EnergyAccountDetailV4) SetOpenStatus(v string)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *EnergyAccountDetailV4) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetCreationDate

`func (o *EnergyAccountDetailV4) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *EnergyAccountDetailV4) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *EnergyAccountDetailV4) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *EnergyAccountDetailV4) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetPlans

`func (o *EnergyAccountDetailV4) GetPlans() []EnergyAccountDetailV4AllOfPlans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *EnergyAccountDetailV4) GetPlansOk() (*[]EnergyAccountDetailV4AllOfPlans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *EnergyAccountDetailV4) SetPlans(v []EnergyAccountDetailV4AllOfPlans)`

SetPlans sets Plans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


