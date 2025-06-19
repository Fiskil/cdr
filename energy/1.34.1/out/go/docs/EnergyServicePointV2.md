# EnergyServicePointV2

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

## Methods

### NewEnergyServicePointV2

`func NewEnergyServicePointV2(servicePointId string, nationalMeteringId string, servicePointClassification string, servicePointStatus string, jurisdictionCode string, validFromDate string, lastUpdateDateTime string, ) *EnergyServicePointV2`

NewEnergyServicePointV2 instantiates a new EnergyServicePointV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointV2WithDefaults

`func NewEnergyServicePointV2WithDefaults() *EnergyServicePointV2`

NewEnergyServicePointV2WithDefaults instantiates a new EnergyServicePointV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyServicePointV2) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyServicePointV2) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyServicePointV2) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.


### GetNationalMeteringId

`func (o *EnergyServicePointV2) GetNationalMeteringId() string`

GetNationalMeteringId returns the NationalMeteringId field if non-nil, zero value otherwise.

### GetNationalMeteringIdOk

`func (o *EnergyServicePointV2) GetNationalMeteringIdOk() (*string, bool)`

GetNationalMeteringIdOk returns a tuple with the NationalMeteringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalMeteringId

`func (o *EnergyServicePointV2) SetNationalMeteringId(v string)`

SetNationalMeteringId sets NationalMeteringId field to given value.


### GetServicePointClassification

`func (o *EnergyServicePointV2) GetServicePointClassification() string`

GetServicePointClassification returns the ServicePointClassification field if non-nil, zero value otherwise.

### GetServicePointClassificationOk

`func (o *EnergyServicePointV2) GetServicePointClassificationOk() (*string, bool)`

GetServicePointClassificationOk returns a tuple with the ServicePointClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointClassification

`func (o *EnergyServicePointV2) SetServicePointClassification(v string)`

SetServicePointClassification sets ServicePointClassification field to given value.


### GetServicePointStatus

`func (o *EnergyServicePointV2) GetServicePointStatus() string`

GetServicePointStatus returns the ServicePointStatus field if non-nil, zero value otherwise.

### GetServicePointStatusOk

`func (o *EnergyServicePointV2) GetServicePointStatusOk() (*string, bool)`

GetServicePointStatusOk returns a tuple with the ServicePointStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointStatus

`func (o *EnergyServicePointV2) SetServicePointStatus(v string)`

SetServicePointStatus sets ServicePointStatus field to given value.


### GetJurisdictionCode

`func (o *EnergyServicePointV2) GetJurisdictionCode() string`

GetJurisdictionCode returns the JurisdictionCode field if non-nil, zero value otherwise.

### GetJurisdictionCodeOk

`func (o *EnergyServicePointV2) GetJurisdictionCodeOk() (*string, bool)`

GetJurisdictionCodeOk returns a tuple with the JurisdictionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionCode

`func (o *EnergyServicePointV2) SetJurisdictionCode(v string)`

SetJurisdictionCode sets JurisdictionCode field to given value.


### GetIsGenerator

`func (o *EnergyServicePointV2) GetIsGenerator() bool`

GetIsGenerator returns the IsGenerator field if non-nil, zero value otherwise.

### GetIsGeneratorOk

`func (o *EnergyServicePointV2) GetIsGeneratorOk() (*bool, bool)`

GetIsGeneratorOk returns a tuple with the IsGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenerator

`func (o *EnergyServicePointV2) SetIsGenerator(v bool)`

SetIsGenerator sets IsGenerator field to given value.

### HasIsGenerator

`func (o *EnergyServicePointV2) HasIsGenerator() bool`

HasIsGenerator returns a boolean if a field has been set.

### GetValidFromDate

`func (o *EnergyServicePointV2) GetValidFromDate() string`

GetValidFromDate returns the ValidFromDate field if non-nil, zero value otherwise.

### GetValidFromDateOk

`func (o *EnergyServicePointV2) GetValidFromDateOk() (*string, bool)`

GetValidFromDateOk returns a tuple with the ValidFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromDate

`func (o *EnergyServicePointV2) SetValidFromDate(v string)`

SetValidFromDate sets ValidFromDate field to given value.


### GetLastUpdateDateTime

`func (o *EnergyServicePointV2) GetLastUpdateDateTime() string`

GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.

### GetLastUpdateDateTimeOk

`func (o *EnergyServicePointV2) GetLastUpdateDateTimeOk() (*string, bool)`

GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateDateTime

`func (o *EnergyServicePointV2) SetLastUpdateDateTime(v string)`

SetLastUpdateDateTime sets LastUpdateDateTime field to given value.


### GetLastConsumerChangeDate

`func (o *EnergyServicePointV2) GetLastConsumerChangeDate() string`

GetLastConsumerChangeDate returns the LastConsumerChangeDate field if non-nil, zero value otherwise.

### GetLastConsumerChangeDateOk

`func (o *EnergyServicePointV2) GetLastConsumerChangeDateOk() (*string, bool)`

GetLastConsumerChangeDateOk returns a tuple with the LastConsumerChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumerChangeDate

`func (o *EnergyServicePointV2) SetLastConsumerChangeDate(v string)`

SetLastConsumerChangeDate sets LastConsumerChangeDate field to given value.

### HasLastConsumerChangeDate

`func (o *EnergyServicePointV2) HasLastConsumerChangeDate() bool`

HasLastConsumerChangeDate returns a boolean if a field has been set.

### GetConsumerProfile

`func (o *EnergyServicePointV2) GetConsumerProfile() EnergyServicePointV2ConsumerProfile`

GetConsumerProfile returns the ConsumerProfile field if non-nil, zero value otherwise.

### GetConsumerProfileOk

`func (o *EnergyServicePointV2) GetConsumerProfileOk() (*EnergyServicePointV2ConsumerProfile, bool)`

GetConsumerProfileOk returns a tuple with the ConsumerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerProfile

`func (o *EnergyServicePointV2) SetConsumerProfile(v EnergyServicePointV2ConsumerProfile)`

SetConsumerProfile sets ConsumerProfile field to given value.

### HasConsumerProfile

`func (o *EnergyServicePointV2) HasConsumerProfile() bool`

HasConsumerProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


