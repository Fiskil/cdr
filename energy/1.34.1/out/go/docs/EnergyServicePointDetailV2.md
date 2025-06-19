# EnergyServicePointDetailV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePointId** | **string** | Unique identifier for the service point. | 
**NationalMeteringId** | **string** | The independent ID of the service point, known in the industry as the NMI. | 
**ServicePointClassification** | **string** | The classification of the service point as defined in MSATS procedures. | 
**ServicePointStatus** | **string** | Code used to indicate the status of the service point. Note the details for the enumeration values below:&lt;ul&gt;&lt;li&gt;&#x60;ACTIVE&#x60;: An active, energised, service point&lt;/li&gt;&lt;li&gt;&#x60;DE_ENERGISED&#x60;: The service point exists but is deenergised&lt;/li&gt;&lt;li&gt;&#x60;EXTINCT&#x60;: The service point has been permanently decommissioned&lt;/li&gt;&lt;li&gt;&#x60;GREENFIELD&#x60;: Applies to a service point that has never been energised&lt;/li&gt;&lt;li&gt;&#x60;OFF_MARKET&#x60;: Applies when the service point is no longer settled in the NEM.&lt;/li&gt;&lt;/ul&gt; | 
**JurisdictionCode** | **string** | Jurisdiction code to which the service point belongs.This code defines the jurisdictional rules which apply to the service point. Note the details of enumeration values below:&lt;ul&gt;&lt;li&gt;&#x60;ALL&#x60;: All Jurisdictions&lt;/li&gt;&lt;li&gt;&#x60;ACT&#x60;: Australian Capital Territory&lt;/li&gt;&lt;li&gt;&#x60;NEM&#x60;: National Electricity Market&lt;/li&gt;&lt;li&gt;&#x60;NSW&#x60;: New South Wales&lt;/li&gt;&lt;li&gt;&#x60;QLD&#x60;: Queensland&lt;/li&gt;&lt;li&gt;&#x60;SA&#x60;: South Australia&lt;/li&gt;&lt;li&gt;&#x60;TAS&#x60;: Tasmania&lt;/li&gt;&lt;li&gt;&#x60;VIC&#x60;: Victoria.&lt;/li&gt;&lt;/ul&gt; | 
**IsGenerator** | Pointer to **bool** | This flag determines whether the energy at this connection point is to be treated as consumer load or as a generating unit (this may include generator auxiliary loads). If absent defaults to &#x60;false&#x60;. &lt;br&gt;**Note:** Only applicable for scheduled or semischeduled generators, does not indicate on site generation by consumer. | [optional] [default to false]
**ValidFromDate** | **string** | The latest start date from which the constituent data sets of this service point became valid. | 
**LastUpdateDateTime** | **string** | The date and time that the information for this service point was modified. | 
**LastConsumerChangeDate** | Pointer to **string** | The date the account holder changed for the NMI. | [optional] 
**ConsumerProfile** | Pointer to [**EnergyServicePointV2ConsumerProfile**](EnergyServicePointV2ConsumerProfile.md) |  | [optional] 
**DistributionLossFactor** | [**EnergyServicePointDetailV2AllOfDistributionLossFactor**](EnergyServicePointDetailV2AllOfDistributionLossFactor.md) |  | 
**RelatedParticipants** | [**[]EnergyServicePointDetailV2AllOfRelatedParticipants**](EnergyServicePointDetailV2AllOfRelatedParticipants.md) |  | 
**Location** | [**CommonPhysicalAddress**](CommonPhysicalAddress.md) | Location of the servicepoint. | 
**Meters** | Pointer to [**[]EnergyServicePointDetailV2AllOfMeters**](EnergyServicePointDetailV2AllOfMeters.md) | The meters associated with the service point. This may be empty where there are no meters physically installed at the service point. | [optional] 

## Methods

### NewEnergyServicePointDetailV2

`func NewEnergyServicePointDetailV2(servicePointId string, nationalMeteringId string, servicePointClassification string, servicePointStatus string, jurisdictionCode string, validFromDate string, lastUpdateDateTime string, distributionLossFactor EnergyServicePointDetailV2AllOfDistributionLossFactor, relatedParticipants []EnergyServicePointDetailV2AllOfRelatedParticipants, location CommonPhysicalAddress, ) *EnergyServicePointDetailV2`

NewEnergyServicePointDetailV2 instantiates a new EnergyServicePointDetailV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointDetailV2WithDefaults

`func NewEnergyServicePointDetailV2WithDefaults() *EnergyServicePointDetailV2`

NewEnergyServicePointDetailV2WithDefaults instantiates a new EnergyServicePointDetailV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyServicePointDetailV2) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyServicePointDetailV2) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyServicePointDetailV2) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.


### GetNationalMeteringId

`func (o *EnergyServicePointDetailV2) GetNationalMeteringId() string`

GetNationalMeteringId returns the NationalMeteringId field if non-nil, zero value otherwise.

### GetNationalMeteringIdOk

`func (o *EnergyServicePointDetailV2) GetNationalMeteringIdOk() (*string, bool)`

GetNationalMeteringIdOk returns a tuple with the NationalMeteringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalMeteringId

`func (o *EnergyServicePointDetailV2) SetNationalMeteringId(v string)`

SetNationalMeteringId sets NationalMeteringId field to given value.


### GetServicePointClassification

`func (o *EnergyServicePointDetailV2) GetServicePointClassification() string`

GetServicePointClassification returns the ServicePointClassification field if non-nil, zero value otherwise.

### GetServicePointClassificationOk

`func (o *EnergyServicePointDetailV2) GetServicePointClassificationOk() (*string, bool)`

GetServicePointClassificationOk returns a tuple with the ServicePointClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointClassification

`func (o *EnergyServicePointDetailV2) SetServicePointClassification(v string)`

SetServicePointClassification sets ServicePointClassification field to given value.


### GetServicePointStatus

`func (o *EnergyServicePointDetailV2) GetServicePointStatus() string`

GetServicePointStatus returns the ServicePointStatus field if non-nil, zero value otherwise.

### GetServicePointStatusOk

`func (o *EnergyServicePointDetailV2) GetServicePointStatusOk() (*string, bool)`

GetServicePointStatusOk returns a tuple with the ServicePointStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointStatus

`func (o *EnergyServicePointDetailV2) SetServicePointStatus(v string)`

SetServicePointStatus sets ServicePointStatus field to given value.


### GetJurisdictionCode

`func (o *EnergyServicePointDetailV2) GetJurisdictionCode() string`

GetJurisdictionCode returns the JurisdictionCode field if non-nil, zero value otherwise.

### GetJurisdictionCodeOk

`func (o *EnergyServicePointDetailV2) GetJurisdictionCodeOk() (*string, bool)`

GetJurisdictionCodeOk returns a tuple with the JurisdictionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionCode

`func (o *EnergyServicePointDetailV2) SetJurisdictionCode(v string)`

SetJurisdictionCode sets JurisdictionCode field to given value.


### GetIsGenerator

`func (o *EnergyServicePointDetailV2) GetIsGenerator() bool`

GetIsGenerator returns the IsGenerator field if non-nil, zero value otherwise.

### GetIsGeneratorOk

`func (o *EnergyServicePointDetailV2) GetIsGeneratorOk() (*bool, bool)`

GetIsGeneratorOk returns a tuple with the IsGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenerator

`func (o *EnergyServicePointDetailV2) SetIsGenerator(v bool)`

SetIsGenerator sets IsGenerator field to given value.

### HasIsGenerator

`func (o *EnergyServicePointDetailV2) HasIsGenerator() bool`

HasIsGenerator returns a boolean if a field has been set.

### GetValidFromDate

`func (o *EnergyServicePointDetailV2) GetValidFromDate() string`

GetValidFromDate returns the ValidFromDate field if non-nil, zero value otherwise.

### GetValidFromDateOk

`func (o *EnergyServicePointDetailV2) GetValidFromDateOk() (*string, bool)`

GetValidFromDateOk returns a tuple with the ValidFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromDate

`func (o *EnergyServicePointDetailV2) SetValidFromDate(v string)`

SetValidFromDate sets ValidFromDate field to given value.


### GetLastUpdateDateTime

`func (o *EnergyServicePointDetailV2) GetLastUpdateDateTime() string`

GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.

### GetLastUpdateDateTimeOk

`func (o *EnergyServicePointDetailV2) GetLastUpdateDateTimeOk() (*string, bool)`

GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateDateTime

`func (o *EnergyServicePointDetailV2) SetLastUpdateDateTime(v string)`

SetLastUpdateDateTime sets LastUpdateDateTime field to given value.


### GetLastConsumerChangeDate

`func (o *EnergyServicePointDetailV2) GetLastConsumerChangeDate() string`

GetLastConsumerChangeDate returns the LastConsumerChangeDate field if non-nil, zero value otherwise.

### GetLastConsumerChangeDateOk

`func (o *EnergyServicePointDetailV2) GetLastConsumerChangeDateOk() (*string, bool)`

GetLastConsumerChangeDateOk returns a tuple with the LastConsumerChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumerChangeDate

`func (o *EnergyServicePointDetailV2) SetLastConsumerChangeDate(v string)`

SetLastConsumerChangeDate sets LastConsumerChangeDate field to given value.

### HasLastConsumerChangeDate

`func (o *EnergyServicePointDetailV2) HasLastConsumerChangeDate() bool`

HasLastConsumerChangeDate returns a boolean if a field has been set.

### GetConsumerProfile

`func (o *EnergyServicePointDetailV2) GetConsumerProfile() EnergyServicePointV2ConsumerProfile`

GetConsumerProfile returns the ConsumerProfile field if non-nil, zero value otherwise.

### GetConsumerProfileOk

`func (o *EnergyServicePointDetailV2) GetConsumerProfileOk() (*EnergyServicePointV2ConsumerProfile, bool)`

GetConsumerProfileOk returns a tuple with the ConsumerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerProfile

`func (o *EnergyServicePointDetailV2) SetConsumerProfile(v EnergyServicePointV2ConsumerProfile)`

SetConsumerProfile sets ConsumerProfile field to given value.

### HasConsumerProfile

`func (o *EnergyServicePointDetailV2) HasConsumerProfile() bool`

HasConsumerProfile returns a boolean if a field has been set.

### GetDistributionLossFactor

`func (o *EnergyServicePointDetailV2) GetDistributionLossFactor() EnergyServicePointDetailV2AllOfDistributionLossFactor`

GetDistributionLossFactor returns the DistributionLossFactor field if non-nil, zero value otherwise.

### GetDistributionLossFactorOk

`func (o *EnergyServicePointDetailV2) GetDistributionLossFactorOk() (*EnergyServicePointDetailV2AllOfDistributionLossFactor, bool)`

GetDistributionLossFactorOk returns a tuple with the DistributionLossFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionLossFactor

`func (o *EnergyServicePointDetailV2) SetDistributionLossFactor(v EnergyServicePointDetailV2AllOfDistributionLossFactor)`

SetDistributionLossFactor sets DistributionLossFactor field to given value.


### GetRelatedParticipants

`func (o *EnergyServicePointDetailV2) GetRelatedParticipants() []EnergyServicePointDetailV2AllOfRelatedParticipants`

GetRelatedParticipants returns the RelatedParticipants field if non-nil, zero value otherwise.

### GetRelatedParticipantsOk

`func (o *EnergyServicePointDetailV2) GetRelatedParticipantsOk() (*[]EnergyServicePointDetailV2AllOfRelatedParticipants, bool)`

GetRelatedParticipantsOk returns a tuple with the RelatedParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParticipants

`func (o *EnergyServicePointDetailV2) SetRelatedParticipants(v []EnergyServicePointDetailV2AllOfRelatedParticipants)`

SetRelatedParticipants sets RelatedParticipants field to given value.


### GetLocation

`func (o *EnergyServicePointDetailV2) GetLocation() CommonPhysicalAddress`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EnergyServicePointDetailV2) GetLocationOk() (*CommonPhysicalAddress, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EnergyServicePointDetailV2) SetLocation(v CommonPhysicalAddress)`

SetLocation sets Location field to given value.


### GetMeters

`func (o *EnergyServicePointDetailV2) GetMeters() []EnergyServicePointDetailV2AllOfMeters`

GetMeters returns the Meters field if non-nil, zero value otherwise.

### GetMetersOk

`func (o *EnergyServicePointDetailV2) GetMetersOk() (*[]EnergyServicePointDetailV2AllOfMeters, bool)`

GetMetersOk returns a tuple with the Meters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeters

`func (o *EnergyServicePointDetailV2) SetMeters(v []EnergyServicePointDetailV2AllOfMeters)`

SetMeters sets Meters field to given value.

### HasMeters

`func (o *EnergyServicePointDetailV2) HasMeters() bool`

HasMeters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


