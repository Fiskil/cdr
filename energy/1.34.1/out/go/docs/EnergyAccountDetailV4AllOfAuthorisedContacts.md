# EnergyAccountDetailV4AllOfAuthorisedContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | For people with single names this field need not be present. The single name should be in the _lastName_ field. | [optional] 
**LastName** | **string** | For people with single names the single name should be in this field. | 
**MiddleNames** | Pointer to **[]string** | Field is mandatory but array may be empty. | [optional] 
**Prefix** | Pointer to **string** | Also known as title or salutation. The prefix to the name (e.g., Mr, Mrs, Ms, Miss, Sir, etc.) | [optional] 
**Suffix** | Pointer to **string** | Used for a trailing suffix to the name (e.g., Jr.) | [optional] 

## Methods

### NewEnergyAccountDetailV4AllOfAuthorisedContacts

`func NewEnergyAccountDetailV4AllOfAuthorisedContacts(lastName string, ) *EnergyAccountDetailV4AllOfAuthorisedContacts`

NewEnergyAccountDetailV4AllOfAuthorisedContacts instantiates a new EnergyAccountDetailV4AllOfAuthorisedContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyAccountDetailV4AllOfAuthorisedContactsWithDefaults

`func NewEnergyAccountDetailV4AllOfAuthorisedContactsWithDefaults() *EnergyAccountDetailV4AllOfAuthorisedContacts`

NewEnergyAccountDetailV4AllOfAuthorisedContactsWithDefaults instantiates a new EnergyAccountDetailV4AllOfAuthorisedContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMiddleNames

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetMiddleNames() []string`

GetMiddleNames returns the MiddleNames field if non-nil, zero value otherwise.

### GetMiddleNamesOk

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetMiddleNamesOk() (*[]string, bool)`

GetMiddleNamesOk returns a tuple with the MiddleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleNames

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) SetMiddleNames(v []string)`

SetMiddleNames sets MiddleNames field to given value.

### HasMiddleNames

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) HasMiddleNames() bool`

HasMiddleNames returns a boolean if a field has been set.

### GetPrefix

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSuffix

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *EnergyAccountDetailV4AllOfAuthorisedContacts) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


