# EnergyServicePointDetailV2AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistributionLossFactor** | [**EnergyServicePointDetailV2AllOfDistributionLossFactor**](EnergyServicePointDetailV2AllOfDistributionLossFactor.md) |  | 
**RelatedParticipants** | [**[]EnergyServicePointDetailV2AllOfRelatedParticipants**](EnergyServicePointDetailV2AllOfRelatedParticipants.md) |  | 
**Location** | [**CommonPhysicalAddress**](CommonPhysicalAddress.md) | Location of the servicepoint. | 
**Meters** | Pointer to [**[]EnergyServicePointDetailV2AllOfMeters**](EnergyServicePointDetailV2AllOfMeters.md) | The meters associated with the service point. This may be empty where there are no meters physically installed at the service point. | [optional] 

## Methods

### NewEnergyServicePointDetailV2AllOf

`func NewEnergyServicePointDetailV2AllOf(distributionLossFactor EnergyServicePointDetailV2AllOfDistributionLossFactor, relatedParticipants []EnergyServicePointDetailV2AllOfRelatedParticipants, location CommonPhysicalAddress, ) *EnergyServicePointDetailV2AllOf`

NewEnergyServicePointDetailV2AllOf instantiates a new EnergyServicePointDetailV2AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointDetailV2AllOfWithDefaults

`func NewEnergyServicePointDetailV2AllOfWithDefaults() *EnergyServicePointDetailV2AllOf`

NewEnergyServicePointDetailV2AllOfWithDefaults instantiates a new EnergyServicePointDetailV2AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistributionLossFactor

`func (o *EnergyServicePointDetailV2AllOf) GetDistributionLossFactor() EnergyServicePointDetailV2AllOfDistributionLossFactor`

GetDistributionLossFactor returns the DistributionLossFactor field if non-nil, zero value otherwise.

### GetDistributionLossFactorOk

`func (o *EnergyServicePointDetailV2AllOf) GetDistributionLossFactorOk() (*EnergyServicePointDetailV2AllOfDistributionLossFactor, bool)`

GetDistributionLossFactorOk returns a tuple with the DistributionLossFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionLossFactor

`func (o *EnergyServicePointDetailV2AllOf) SetDistributionLossFactor(v EnergyServicePointDetailV2AllOfDistributionLossFactor)`

SetDistributionLossFactor sets DistributionLossFactor field to given value.


### GetRelatedParticipants

`func (o *EnergyServicePointDetailV2AllOf) GetRelatedParticipants() []EnergyServicePointDetailV2AllOfRelatedParticipants`

GetRelatedParticipants returns the RelatedParticipants field if non-nil, zero value otherwise.

### GetRelatedParticipantsOk

`func (o *EnergyServicePointDetailV2AllOf) GetRelatedParticipantsOk() (*[]EnergyServicePointDetailV2AllOfRelatedParticipants, bool)`

GetRelatedParticipantsOk returns a tuple with the RelatedParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParticipants

`func (o *EnergyServicePointDetailV2AllOf) SetRelatedParticipants(v []EnergyServicePointDetailV2AllOfRelatedParticipants)`

SetRelatedParticipants sets RelatedParticipants field to given value.


### GetLocation

`func (o *EnergyServicePointDetailV2AllOf) GetLocation() CommonPhysicalAddress`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EnergyServicePointDetailV2AllOf) GetLocationOk() (*CommonPhysicalAddress, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EnergyServicePointDetailV2AllOf) SetLocation(v CommonPhysicalAddress)`

SetLocation sets Location field to given value.


### GetMeters

`func (o *EnergyServicePointDetailV2AllOf) GetMeters() []EnergyServicePointDetailV2AllOfMeters`

GetMeters returns the Meters field if non-nil, zero value otherwise.

### GetMetersOk

`func (o *EnergyServicePointDetailV2AllOf) GetMetersOk() (*[]EnergyServicePointDetailV2AllOfMeters, bool)`

GetMetersOk returns a tuple with the Meters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeters

`func (o *EnergyServicePointDetailV2AllOf) SetMeters(v []EnergyServicePointDetailV2AllOfMeters)`

SetMeters sets Meters field to given value.

### HasMeters

`func (o *EnergyServicePointDetailV2AllOf) HasMeters() bool`

HasMeters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


