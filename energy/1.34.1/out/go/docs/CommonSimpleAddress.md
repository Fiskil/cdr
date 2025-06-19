# CommonSimpleAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailingName** | Pointer to **string** | Name of the individual or business formatted for inclusion in an address used for physical mail. | [optional] 
**AddressLine1** | **string** | First line of the standard address object. | 
**AddressLine2** | Pointer to **string** | Second line of the standard address object. | [optional] 
**AddressLine3** | Pointer to **string** | Third line of the standard address object. | [optional] 
**Postcode** | Pointer to **string** | Mandatory for Australian addresses. | [optional] 
**City** | **string** | Name of the city or locality. | 
**State** | **string** | Free text if the country is not Australia. If country is Australia then must be one of the values defined by the [State Type Abbreviation](https://auspost.com.au/content/dam/auspost_corp/media/documents/australia-post-data-guide.pdf) in the PAF file format. &#x60;NSW&#x60;, &#x60;QLD&#x60;, &#x60;VIC&#x60;, &#x60;NT&#x60;, &#x60;WA&#x60;, &#x60;SA&#x60;, &#x60;TAS&#x60;, &#x60;ACT&#x60;, &#x60;AAT&#x60;. | 
**Country** | Pointer to **string** | A valid [ISO 3166 Alpha-3](https://www.iso.org/iso-3166-country-codes.html) country code. Australia (&#x60;AUS&#x60;) is assumed if country is not present. | [optional] [default to "AUS"]

## Methods

### NewCommonSimpleAddress

`func NewCommonSimpleAddress(addressLine1 string, city string, state string, ) *CommonSimpleAddress`

NewCommonSimpleAddress instantiates a new CommonSimpleAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSimpleAddressWithDefaults

`func NewCommonSimpleAddressWithDefaults() *CommonSimpleAddress`

NewCommonSimpleAddressWithDefaults instantiates a new CommonSimpleAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailingName

`func (o *CommonSimpleAddress) GetMailingName() string`

GetMailingName returns the MailingName field if non-nil, zero value otherwise.

### GetMailingNameOk

`func (o *CommonSimpleAddress) GetMailingNameOk() (*string, bool)`

GetMailingNameOk returns a tuple with the MailingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingName

`func (o *CommonSimpleAddress) SetMailingName(v string)`

SetMailingName sets MailingName field to given value.

### HasMailingName

`func (o *CommonSimpleAddress) HasMailingName() bool`

HasMailingName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CommonSimpleAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CommonSimpleAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CommonSimpleAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *CommonSimpleAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CommonSimpleAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CommonSimpleAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CommonSimpleAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *CommonSimpleAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *CommonSimpleAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *CommonSimpleAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *CommonSimpleAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetPostcode

`func (o *CommonSimpleAddress) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *CommonSimpleAddress) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *CommonSimpleAddress) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *CommonSimpleAddress) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetCity

`func (o *CommonSimpleAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CommonSimpleAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CommonSimpleAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *CommonSimpleAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonSimpleAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonSimpleAddress) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *CommonSimpleAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CommonSimpleAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CommonSimpleAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CommonSimpleAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


