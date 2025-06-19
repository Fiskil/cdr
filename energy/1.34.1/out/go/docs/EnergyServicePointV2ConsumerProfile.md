# EnergyServicePointV2ConsumerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | Pointer to **string** | A code that defines the consumer class as defined in the National Energy Retail Regulations, or in overriding Jurisdictional instruments. | [optional] 
**Threshold** | Pointer to **string** | A code that defines the consumption threshold as defined in the National Energy Retail Regulations, or in overriding Jurisdictional instruments. Note the details of enumeration values below: &lt;ul&gt;&lt;li&gt;&#x60;LOW&#x60;: Consumption is less than the &#39;lower consumption threshold&#39; as defined in the National Energy Retail Regulations&lt;/li&gt;&lt;li&gt;&#x60;MEDIUM&#x60;: Consumption is equal to or greater than the &#39;lower consumption threshold&#39;, but less than the &#39;upper consumption threshold&#39;, as defined in the National Energy Retail Regulations&lt;/li&gt;&lt;li&gt;&#x60;HIGH&#x60;: Consumption is equal to or greater than the &#39;upper consumption threshold&#39; as defined in the National Energy Retail Regulations.&lt;/li&gt;&lt;/ul&gt; | [optional] 

## Methods

### NewEnergyServicePointV2ConsumerProfile

`func NewEnergyServicePointV2ConsumerProfile() *EnergyServicePointV2ConsumerProfile`

NewEnergyServicePointV2ConsumerProfile instantiates a new EnergyServicePointV2ConsumerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyServicePointV2ConsumerProfileWithDefaults

`func NewEnergyServicePointV2ConsumerProfileWithDefaults() *EnergyServicePointV2ConsumerProfile`

NewEnergyServicePointV2ConsumerProfileWithDefaults instantiates a new EnergyServicePointV2ConsumerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *EnergyServicePointV2ConsumerProfile) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *EnergyServicePointV2ConsumerProfile) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *EnergyServicePointV2ConsumerProfile) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *EnergyServicePointV2ConsumerProfile) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetThreshold

`func (o *EnergyServicePointV2ConsumerProfile) GetThreshold() string`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *EnergyServicePointV2ConsumerProfile) GetThresholdOk() (*string, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *EnergyServicePointV2ConsumerProfile) SetThreshold(v string)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *EnergyServicePointV2ConsumerProfile) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


