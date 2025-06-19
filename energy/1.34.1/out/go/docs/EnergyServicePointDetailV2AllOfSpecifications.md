# EnergyServicePointDetailV2AllOfSpecifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | A code to denote the status of the meter. Note the details of enumeration values below: &lt;ul&gt;&lt;li&gt;&#x60;CURRENT&#x60;: Applies when a meter is current and not disconnected&lt;/li&gt;&lt;li&gt;&#x60;DISCONNECTED&#x60;: Applies when a meter is present but has been remotely disconnected.&lt;/li&gt;&lt;/ul&gt; | 
**InstallationType** | **string** | The metering Installation type code indicates whether the metering installation has to be manually read. Note the details of enumeration values below: &lt;ul&gt;&lt;li&gt;&#x60;BASIC&#x60;: Accumulation Meter – Type 6&lt;/li&gt;&lt;li&gt;&#x60;COMMS1&#x60;: Interval Meter with communications – Type 1&lt;/li&gt;&lt;li&gt;&#x60;COMMS2&#x60;: Interval Meter with communications – Type 2&lt;/li&gt;&lt;li&gt;&#x60;COMMS3&#x60;: Interval Meter with communications – Type 3&lt;/li&gt;&lt;li&gt;&#x60;COMMS4&#x60;: Interval Meter with communications – Type 4&lt;/li&gt;&lt;li&gt;&#x60;COMMS4C&#x60;: CT connected metering installation that meets the minimum services specifications&lt;/li&gt;&lt;li&gt;&#x60;COMMS4D&#x60;: Whole current metering installation that meets the minimum services specifications&lt;/li&gt;&lt;li&gt;&#x60;MRAM&#x60;: Small customer metering installation – Type 4A&lt;/li&gt;&lt;li&gt;&#x60;MRIM&#x60;: Manually Read Interval Meter – Type 5&lt;/li&gt;&lt;li&gt;&#x60;UMCP&#x60;: Unmetered Supply – Type 7&lt;/li&gt;&lt;li&gt;&#x60;VICAMI&#x60;: A relevant metering installation as defined in clause 9.9C of the NER&lt;/li&gt;&lt;li&gt;&#x60;NCONUML&#x60;: Non-contestable unmeter load - Introduced as part of Global Settlement.&lt;/li&gt;&lt;/ul&gt; | 
**Manufacturer** | Pointer to **string** | Free text field to identify the manufacturer of the installed meter. | [optional] 
**Model** | Pointer to **string** | Free text field to identify the meter manufacturer&#39;s designation for the meter model. | [optional] 
**ReadType** | Pointer to **string** | Code to denote the method and frequency of Meter Reading. The value is formatted as follows: &lt;ul&gt;&lt;li&gt;First Character &#x3D; Remote (&#x60;R&#x60;) or Manual (&#x60;M&#x60;)&lt;/li&gt;&lt;li&gt;Second Character &#x3D; Mode: &#x60;T&#x60; &#x3D; telephone, &#x60;W&#x60; &#x3D; wireless, &#x60;P&#x60; &#x3D; powerline, &#x60;I&#x60; &#x3D; infra-red, &#x60;G&#x60; &#x3D; galvanic, &#x60;V&#x60; &#x3D; visual&lt;/li&gt;&lt;li&gt;Third Character &#x3D; Frequency of Scheduled Meter Readings: &#x60;1&#x60; &#x3D; Twelve times per year, &#x60;2&#x60; &#x3D; Six times per year, &#x60;3&#x60; &#x3D; Four times per year, &#x60;D&#x60; &#x3D; Daily or weekly&lt;/li&gt;&lt;li&gt;Optional Fourth Character &#x3D; to identify what interval length the meter is capable of reading. This includes five, 15 and 30 minute granularity as the following: &#x60;A&#x60; &#x3D; 5 minute, &#x60;B&#x60; &#x3D; 15 minute, &#x60;C&#x60; &#x3D; 30 minute, &#x60;D&#x60; &#x3D; Cannot convert to 5 minute (i.e. due to metering installation de-energised), &#x60;M&#x60; &#x3D; Manually Read Accumulation Meter.&lt;/li&gt;&lt;/ul&gt;For example, &lt;ul&gt;&lt;li&gt;&#x60;MV3&#x60; &#x3D; Manual, Visual, Quarterly&lt;/li&gt;&lt;li&gt;&#x60;MV3M&#x60; &#x3D; Manual, Visual, Quarterly, Manually Read Accumulation Meter&lt;/li&gt;&lt;li&gt;&#x60;RWDC&#x60; &#x3D; Remote, Wireless, Daily, 30 minutes interval.&lt;/li&gt;&lt;/ul&gt; | [optional] 
**NextScheduledReadDate** | Pointer to **string** | This date is the next scheduled meter read date (NSRD) if a manual Meter Reading is required. | [optional] 

## Methods

### NewEnergyServicePointDetailV2AllOfSpecifications

`func NewEnergyServicePointDetailV2AllOfSpecifications(status string, installationType string, ) *EnergyServicePointDetailV2AllOfSpecifications`

NewEnergyServicePointDetailV2AllOfSpecifications instantiates a new EnergyServicePointDetailV2AllOfSpecifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointDetailV2AllOfSpecificationsWithDefaults

`func NewEnergyServicePointDetailV2AllOfSpecificationsWithDefaults() *EnergyServicePointDetailV2AllOfSpecifications`

NewEnergyServicePointDetailV2AllOfSpecificationsWithDefaults instantiates a new EnergyServicePointDetailV2AllOfSpecifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnergyServicePointDetailV2AllOfSpecifications) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInstallationType

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetInstallationType() string`

GetInstallationType returns the InstallationType field if non-nil, zero value otherwise.

### GetInstallationTypeOk

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetInstallationTypeOk() (*string, bool)`

GetInstallationTypeOk returns a tuple with the InstallationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationType

`func (o *EnergyServicePointDetailV2AllOfSpecifications) SetInstallationType(v string)`

SetInstallationType sets InstallationType field to given value.


### GetManufacturer

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *EnergyServicePointDetailV2AllOfSpecifications) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *EnergyServicePointDetailV2AllOfSpecifications) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EnergyServicePointDetailV2AllOfSpecifications) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EnergyServicePointDetailV2AllOfSpecifications) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetReadType

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetReadType() string`

GetReadType returns the ReadType field if non-nil, zero value otherwise.

### GetReadTypeOk

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetReadTypeOk() (*string, bool)`

GetReadTypeOk returns a tuple with the ReadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadType

`func (o *EnergyServicePointDetailV2AllOfSpecifications) SetReadType(v string)`

SetReadType sets ReadType field to given value.

### HasReadType

`func (o *EnergyServicePointDetailV2AllOfSpecifications) HasReadType() bool`

HasReadType returns a boolean if a field has been set.

### GetNextScheduledReadDate

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetNextScheduledReadDate() string`

GetNextScheduledReadDate returns the NextScheduledReadDate field if non-nil, zero value otherwise.

### GetNextScheduledReadDateOk

`func (o *EnergyServicePointDetailV2AllOfSpecifications) GetNextScheduledReadDateOk() (*string, bool)`

GetNextScheduledReadDateOk returns a tuple with the NextScheduledReadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledReadDate

`func (o *EnergyServicePointDetailV2AllOfSpecifications) SetNextScheduledReadDate(v string)`

SetNextScheduledReadDate sets NextScheduledReadDate field to given value.

### HasNextScheduledReadDate

`func (o *EnergyServicePointDetailV2AllOfSpecifications) HasNextScheduledReadDate() bool`

HasNextScheduledReadDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


