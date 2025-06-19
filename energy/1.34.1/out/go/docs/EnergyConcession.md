# EnergyConcession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Indicator of the method of concession calculation. | 
**DisplayName** | **string** | The display name of the concession. | 
**AdditionalInfo** | Pointer to **string** | Display text providing more information on the concession. Mandatory if _type_ is &#x60;VARIABLE&#x60;. | [optional] 
**AdditionalInfoUri** | Pointer to **string** | Optional link to additional information regarding the concession. | [optional] 
**StartDate** | Pointer to **string** | Optional start date for the application of the concession. | [optional] 
**EndDate** | Pointer to **string** | Optional end date for the application of the concession. | [optional] 
**DiscountFrequency** | Pointer to **string** | Conditional attribute for frequency at which a concession is applied. Required if type is &#x60;FIXED_AMOUNT&#x60; or &#x60;FIXED_PERCENTAGE&#x60;. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | [optional] 
**Amount** | Pointer to **string** | Conditional attribute for the amount of discount for the concession- required if type is &#x60;FIXED_AMOUNT&#x60;. | [optional] 
**Percentage** | Pointer to **string** | Conditional attribute for the percentage of discount of concession - required if type is &#x60;FIXED_PERCENTAGE&#x60;. | [optional] 
**AppliedTo** | Pointer to **[]string** | Array of ENUMs to specify what the concession applies to. Multiple ENUM values can be provided. If absent, &#x60;USAGE&#x60; is assumed. | [optional] 

## Methods

### NewEnergyConcession

`func NewEnergyConcession(type_ string, displayName string, ) *EnergyConcession`

NewEnergyConcession instantiates a new EnergyConcession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyConcessionWithDefaults

`func NewEnergyConcessionWithDefaults() *EnergyConcession`

NewEnergyConcessionWithDefaults instantiates a new EnergyConcession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnergyConcession) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyConcession) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyConcession) SetType(v string)`

SetType sets Type field to given value.


### GetDisplayName

`func (o *EnergyConcession) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyConcession) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyConcession) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetAdditionalInfo

`func (o *EnergyConcession) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *EnergyConcession) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *EnergyConcession) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *EnergyConcession) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAdditionalInfoUri

`func (o *EnergyConcession) GetAdditionalInfoUri() string`

GetAdditionalInfoUri returns the AdditionalInfoUri field if non-nil, zero value otherwise.

### GetAdditionalInfoUriOk

`func (o *EnergyConcession) GetAdditionalInfoUriOk() (*string, bool)`

GetAdditionalInfoUriOk returns a tuple with the AdditionalInfoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfoUri

`func (o *EnergyConcession) SetAdditionalInfoUri(v string)`

SetAdditionalInfoUri sets AdditionalInfoUri field to given value.

### HasAdditionalInfoUri

`func (o *EnergyConcession) HasAdditionalInfoUri() bool`

HasAdditionalInfoUri returns a boolean if a field has been set.

### GetStartDate

`func (o *EnergyConcession) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyConcession) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyConcession) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EnergyConcession) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EnergyConcession) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyConcession) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyConcession) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EnergyConcession) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDiscountFrequency

`func (o *EnergyConcession) GetDiscountFrequency() string`

GetDiscountFrequency returns the DiscountFrequency field if non-nil, zero value otherwise.

### GetDiscountFrequencyOk

`func (o *EnergyConcession) GetDiscountFrequencyOk() (*string, bool)`

GetDiscountFrequencyOk returns a tuple with the DiscountFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountFrequency

`func (o *EnergyConcession) SetDiscountFrequency(v string)`

SetDiscountFrequency sets DiscountFrequency field to given value.

### HasDiscountFrequency

`func (o *EnergyConcession) HasDiscountFrequency() bool`

HasDiscountFrequency returns a boolean if a field has been set.

### GetAmount

`func (o *EnergyConcession) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyConcession) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyConcession) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EnergyConcession) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPercentage

`func (o *EnergyConcession) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *EnergyConcession) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *EnergyConcession) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *EnergyConcession) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetAppliedTo

`func (o *EnergyConcession) GetAppliedTo() []string`

GetAppliedTo returns the AppliedTo field if non-nil, zero value otherwise.

### GetAppliedToOk

`func (o *EnergyConcession) GetAppliedToOk() (*[]string, bool)`

GetAppliedToOk returns a tuple with the AppliedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTo

`func (o *EnergyConcession) SetAppliedTo(v []string)`

SetAppliedTo sets AppliedTo field to given value.

### HasAppliedTo

`func (o *EnergyConcession) HasAppliedTo() bool`

HasAppliedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


