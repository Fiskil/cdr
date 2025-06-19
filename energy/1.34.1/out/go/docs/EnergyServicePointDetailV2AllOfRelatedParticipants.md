# EnergyServicePointDetailV2AllOfRelatedParticipants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Party** | **string** | The name of the party/organisation related to this service point. | 
**Role** | **string** | The role performed by this participant in relation to the service point. Note the details of enumeration values below: &lt;ul&gt;&lt;li&gt;&#x60;FRMP&#x60;: Financially Responsible Market Participant&lt;/li&gt;&lt;li&gt;&#x60;LNSP&#x60;: Local Network Service Provider or Embedded Network Manager for child connection points&lt;/li&gt;&lt;li&gt;&#x60;DRSP&#x60;: wholesale Demand Response and/or market ancillary Service Provider and note that where it is not relevant for a NMI it will not be included.&lt;/li&gt;&lt;/ul&gt; | 

## Methods

### NewEnergyServicePointDetailV2AllOfRelatedParticipants

`func NewEnergyServicePointDetailV2AllOfRelatedParticipants(party string, role string, ) *EnergyServicePointDetailV2AllOfRelatedParticipants`

NewEnergyServicePointDetailV2AllOfRelatedParticipants instantiates a new EnergyServicePointDetailV2AllOfRelatedParticipants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointDetailV2AllOfRelatedParticipantsWithDefaults

`func NewEnergyServicePointDetailV2AllOfRelatedParticipantsWithDefaults() *EnergyServicePointDetailV2AllOfRelatedParticipants`

NewEnergyServicePointDetailV2AllOfRelatedParticipantsWithDefaults instantiates a new EnergyServicePointDetailV2AllOfRelatedParticipants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParty

`func (o *EnergyServicePointDetailV2AllOfRelatedParticipants) GetParty() string`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *EnergyServicePointDetailV2AllOfRelatedParticipants) GetPartyOk() (*string, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *EnergyServicePointDetailV2AllOfRelatedParticipants) SetParty(v string)`

SetParty sets Party field to given value.


### GetRole

`func (o *EnergyServicePointDetailV2AllOfRelatedParticipants) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EnergyServicePointDetailV2AllOfRelatedParticipants) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EnergyServicePointDetailV2AllOfRelatedParticipants) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


