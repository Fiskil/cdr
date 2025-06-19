# EnergyAccountBaseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**AccountNumber** | Pointer to **string** | Optional identifier of the account as defined by the data holder. This must be the value presented on physical statements (if it exists) and must not be used for the value of _accountId_. | [optional] 
**DisplayName** | Pointer to **string** | An optional display name for the account if one exists or can be derived. The content of this field is at the discretion of the data holder. | [optional] 
**OpenStatus** | Pointer to **string** | Open or closed status for the account. If not present then &#x60;OPEN&#x60; is assumed. | [optional] [default to "OPEN"]
**CreationDate** | Pointer to **string** | The date that the account was created or opened. Mandatory if _openStatus_ is &#x60;OPEN&#x60;. | [optional] 

## Methods

### NewEnergyAccountBaseV2

`func NewEnergyAccountBaseV2(accountId string, ) *EnergyAccountBaseV2`

NewEnergyAccountBaseV2 instantiates a new EnergyAccountBaseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyAccountBaseV2WithDefaults

`func NewEnergyAccountBaseV2WithDefaults() *EnergyAccountBaseV2`

NewEnergyAccountBaseV2WithDefaults instantiates a new EnergyAccountBaseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EnergyAccountBaseV2) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EnergyAccountBaseV2) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EnergyAccountBaseV2) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountNumber

`func (o *EnergyAccountBaseV2) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *EnergyAccountBaseV2) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *EnergyAccountBaseV2) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *EnergyAccountBaseV2) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnergyAccountBaseV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyAccountBaseV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyAccountBaseV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnergyAccountBaseV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOpenStatus

`func (o *EnergyAccountBaseV2) GetOpenStatus() string`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *EnergyAccountBaseV2) GetOpenStatusOk() (*string, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *EnergyAccountBaseV2) SetOpenStatus(v string)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *EnergyAccountBaseV2) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetCreationDate

`func (o *EnergyAccountBaseV2) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *EnergyAccountBaseV2) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *EnergyAccountBaseV2) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *EnergyAccountBaseV2) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


