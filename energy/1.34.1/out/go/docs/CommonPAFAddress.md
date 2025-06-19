# CommonPAFAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dpid** | Pointer to **string** | Unique identifier for an address as defined by Australia Post. Also known as Delivery Point Identifier. | [optional] 
**ThoroughfareNumber1** | Pointer to **int32** | Thoroughfare number for a property (first number in a property ranged address). | [optional] 
**ThoroughfareNumber1Suffix** | Pointer to **string** | Suffix for the thoroughfare number. Only relevant if _thoroughfareNumber1_ is populated. | [optional] 
**ThoroughfareNumber2** | Pointer to **int32** | Second thoroughfare number (only used if the property has a ranged address e.g., 23-25). | [optional] 
**ThoroughfareNumber2Suffix** | Pointer to **string** | Suffix for the second thoroughfare number. Only relevant if _thoroughfareNumber2_ is populated. | [optional] 
**FlatUnitType** | Pointer to **string** | Type of flat or unit for the address. | [optional] 
**FlatUnitNumber** | Pointer to **string** | Unit number (including suffix, if applicable). | [optional] 
**FloorLevelType** | Pointer to **string** | Type of floor or level for the address. | [optional] 
**FloorLevelNumber** | Pointer to **string** | Floor or level number (including alpha characters). | [optional] 
**LotNumber** | Pointer to **string** | Allotment number for the address. | [optional] 
**BuildingName1** | Pointer to **string** | Building/Property name 1. | [optional] 
**BuildingName2** | Pointer to **string** | Building/Property name 2. | [optional] 
**StreetName** | Pointer to **string** | The name of the street. | [optional] 
**StreetType** | Pointer to **string** | The street type. Valid enumeration defined by Australia Post PAF code file. | [optional] 
**StreetSuffix** | Pointer to **string** | The street type suffix. Valid enumeration defined by Australia Post PAF code file. | [optional] 
**PostalDeliveryType** | Pointer to **string** | Postal delivery type. (e.g., PO BOX). Valid enumeration defined by Australia Post PAF code file. | [optional] 
**PostalDeliveryNumber** | Pointer to **int32** | Postal delivery number if the address is a postal delivery type. | [optional] 
**PostalDeliveryNumberPrefix** | Pointer to **string** | Postal delivery number prefix related to the postal delivery number. | [optional] 
**PostalDeliveryNumberSuffix** | Pointer to **string** | Postal delivery number suffix related to the postal delivery number. | [optional] 
**LocalityName** | **string** | Full name of locality. | 
**Postcode** | **string** | Postcode for the locality. | 
**State** | **string** | State in which the address belongs. Valid enumeration defined by Australia Post PAF code file [State Type Abbreviation](https://auspost.com.au/content/dam/auspost_corp/media/documents/australia-post-data-guide.pdf). &#x60;NSW&#x60;, &#x60;QLD&#x60;, &#x60;VIC&#x60;, &#x60;NT&#x60;, &#x60;WA&#x60;, &#x60;SA&#x60;, &#x60;TAS&#x60;, &#x60;ACT&#x60;, &#x60;AAT&#x60;. | 

## Methods

### NewCommonPAFAddress

`func NewCommonPAFAddress(localityName string, postcode string, state string, ) *CommonPAFAddress`

NewCommonPAFAddress instantiates a new CommonPAFAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPAFAddressWithDefaults

`func NewCommonPAFAddressWithDefaults() *CommonPAFAddress`

NewCommonPAFAddressWithDefaults instantiates a new CommonPAFAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpid

`func (o *CommonPAFAddress) GetDpid() string`

GetDpid returns the Dpid field if non-nil, zero value otherwise.

### GetDpidOk

`func (o *CommonPAFAddress) GetDpidOk() (*string, bool)`

GetDpidOk returns a tuple with the Dpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpid

`func (o *CommonPAFAddress) SetDpid(v string)`

SetDpid sets Dpid field to given value.

### HasDpid

`func (o *CommonPAFAddress) HasDpid() bool`

HasDpid returns a boolean if a field has been set.

### GetThoroughfareNumber1

`func (o *CommonPAFAddress) GetThoroughfareNumber1() int32`

GetThoroughfareNumber1 returns the ThoroughfareNumber1 field if non-nil, zero value otherwise.

### GetThoroughfareNumber1Ok

`func (o *CommonPAFAddress) GetThoroughfareNumber1Ok() (*int32, bool)`

GetThoroughfareNumber1Ok returns a tuple with the ThoroughfareNumber1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThoroughfareNumber1

`func (o *CommonPAFAddress) SetThoroughfareNumber1(v int32)`

SetThoroughfareNumber1 sets ThoroughfareNumber1 field to given value.

### HasThoroughfareNumber1

`func (o *CommonPAFAddress) HasThoroughfareNumber1() bool`

HasThoroughfareNumber1 returns a boolean if a field has been set.

### GetThoroughfareNumber1Suffix

`func (o *CommonPAFAddress) GetThoroughfareNumber1Suffix() string`

GetThoroughfareNumber1Suffix returns the ThoroughfareNumber1Suffix field if non-nil, zero value otherwise.

### GetThoroughfareNumber1SuffixOk

`func (o *CommonPAFAddress) GetThoroughfareNumber1SuffixOk() (*string, bool)`

GetThoroughfareNumber1SuffixOk returns a tuple with the ThoroughfareNumber1Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThoroughfareNumber1Suffix

`func (o *CommonPAFAddress) SetThoroughfareNumber1Suffix(v string)`

SetThoroughfareNumber1Suffix sets ThoroughfareNumber1Suffix field to given value.

### HasThoroughfareNumber1Suffix

`func (o *CommonPAFAddress) HasThoroughfareNumber1Suffix() bool`

HasThoroughfareNumber1Suffix returns a boolean if a field has been set.

### GetThoroughfareNumber2

`func (o *CommonPAFAddress) GetThoroughfareNumber2() int32`

GetThoroughfareNumber2 returns the ThoroughfareNumber2 field if non-nil, zero value otherwise.

### GetThoroughfareNumber2Ok

`func (o *CommonPAFAddress) GetThoroughfareNumber2Ok() (*int32, bool)`

GetThoroughfareNumber2Ok returns a tuple with the ThoroughfareNumber2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThoroughfareNumber2

`func (o *CommonPAFAddress) SetThoroughfareNumber2(v int32)`

SetThoroughfareNumber2 sets ThoroughfareNumber2 field to given value.

### HasThoroughfareNumber2

`func (o *CommonPAFAddress) HasThoroughfareNumber2() bool`

HasThoroughfareNumber2 returns a boolean if a field has been set.

### GetThoroughfareNumber2Suffix

`func (o *CommonPAFAddress) GetThoroughfareNumber2Suffix() string`

GetThoroughfareNumber2Suffix returns the ThoroughfareNumber2Suffix field if non-nil, zero value otherwise.

### GetThoroughfareNumber2SuffixOk

`func (o *CommonPAFAddress) GetThoroughfareNumber2SuffixOk() (*string, bool)`

GetThoroughfareNumber2SuffixOk returns a tuple with the ThoroughfareNumber2Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThoroughfareNumber2Suffix

`func (o *CommonPAFAddress) SetThoroughfareNumber2Suffix(v string)`

SetThoroughfareNumber2Suffix sets ThoroughfareNumber2Suffix field to given value.

### HasThoroughfareNumber2Suffix

`func (o *CommonPAFAddress) HasThoroughfareNumber2Suffix() bool`

HasThoroughfareNumber2Suffix returns a boolean if a field has been set.

### GetFlatUnitType

`func (o *CommonPAFAddress) GetFlatUnitType() string`

GetFlatUnitType returns the FlatUnitType field if non-nil, zero value otherwise.

### GetFlatUnitTypeOk

`func (o *CommonPAFAddress) GetFlatUnitTypeOk() (*string, bool)`

GetFlatUnitTypeOk returns a tuple with the FlatUnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatUnitType

`func (o *CommonPAFAddress) SetFlatUnitType(v string)`

SetFlatUnitType sets FlatUnitType field to given value.

### HasFlatUnitType

`func (o *CommonPAFAddress) HasFlatUnitType() bool`

HasFlatUnitType returns a boolean if a field has been set.

### GetFlatUnitNumber

`func (o *CommonPAFAddress) GetFlatUnitNumber() string`

GetFlatUnitNumber returns the FlatUnitNumber field if non-nil, zero value otherwise.

### GetFlatUnitNumberOk

`func (o *CommonPAFAddress) GetFlatUnitNumberOk() (*string, bool)`

GetFlatUnitNumberOk returns a tuple with the FlatUnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatUnitNumber

`func (o *CommonPAFAddress) SetFlatUnitNumber(v string)`

SetFlatUnitNumber sets FlatUnitNumber field to given value.

### HasFlatUnitNumber

`func (o *CommonPAFAddress) HasFlatUnitNumber() bool`

HasFlatUnitNumber returns a boolean if a field has been set.

### GetFloorLevelType

`func (o *CommonPAFAddress) GetFloorLevelType() string`

GetFloorLevelType returns the FloorLevelType field if non-nil, zero value otherwise.

### GetFloorLevelTypeOk

`func (o *CommonPAFAddress) GetFloorLevelTypeOk() (*string, bool)`

GetFloorLevelTypeOk returns a tuple with the FloorLevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorLevelType

`func (o *CommonPAFAddress) SetFloorLevelType(v string)`

SetFloorLevelType sets FloorLevelType field to given value.

### HasFloorLevelType

`func (o *CommonPAFAddress) HasFloorLevelType() bool`

HasFloorLevelType returns a boolean if a field has been set.

### GetFloorLevelNumber

`func (o *CommonPAFAddress) GetFloorLevelNumber() string`

GetFloorLevelNumber returns the FloorLevelNumber field if non-nil, zero value otherwise.

### GetFloorLevelNumberOk

`func (o *CommonPAFAddress) GetFloorLevelNumberOk() (*string, bool)`

GetFloorLevelNumberOk returns a tuple with the FloorLevelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorLevelNumber

`func (o *CommonPAFAddress) SetFloorLevelNumber(v string)`

SetFloorLevelNumber sets FloorLevelNumber field to given value.

### HasFloorLevelNumber

`func (o *CommonPAFAddress) HasFloorLevelNumber() bool`

HasFloorLevelNumber returns a boolean if a field has been set.

### GetLotNumber

`func (o *CommonPAFAddress) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *CommonPAFAddress) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *CommonPAFAddress) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *CommonPAFAddress) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetBuildingName1

`func (o *CommonPAFAddress) GetBuildingName1() string`

GetBuildingName1 returns the BuildingName1 field if non-nil, zero value otherwise.

### GetBuildingName1Ok

`func (o *CommonPAFAddress) GetBuildingName1Ok() (*string, bool)`

GetBuildingName1Ok returns a tuple with the BuildingName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingName1

`func (o *CommonPAFAddress) SetBuildingName1(v string)`

SetBuildingName1 sets BuildingName1 field to given value.

### HasBuildingName1

`func (o *CommonPAFAddress) HasBuildingName1() bool`

HasBuildingName1 returns a boolean if a field has been set.

### GetBuildingName2

`func (o *CommonPAFAddress) GetBuildingName2() string`

GetBuildingName2 returns the BuildingName2 field if non-nil, zero value otherwise.

### GetBuildingName2Ok

`func (o *CommonPAFAddress) GetBuildingName2Ok() (*string, bool)`

GetBuildingName2Ok returns a tuple with the BuildingName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingName2

`func (o *CommonPAFAddress) SetBuildingName2(v string)`

SetBuildingName2 sets BuildingName2 field to given value.

### HasBuildingName2

`func (o *CommonPAFAddress) HasBuildingName2() bool`

HasBuildingName2 returns a boolean if a field has been set.

### GetStreetName

`func (o *CommonPAFAddress) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *CommonPAFAddress) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *CommonPAFAddress) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *CommonPAFAddress) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetStreetType

`func (o *CommonPAFAddress) GetStreetType() string`

GetStreetType returns the StreetType field if non-nil, zero value otherwise.

### GetStreetTypeOk

`func (o *CommonPAFAddress) GetStreetTypeOk() (*string, bool)`

GetStreetTypeOk returns a tuple with the StreetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetType

`func (o *CommonPAFAddress) SetStreetType(v string)`

SetStreetType sets StreetType field to given value.

### HasStreetType

`func (o *CommonPAFAddress) HasStreetType() bool`

HasStreetType returns a boolean if a field has been set.

### GetStreetSuffix

`func (o *CommonPAFAddress) GetStreetSuffix() string`

GetStreetSuffix returns the StreetSuffix field if non-nil, zero value otherwise.

### GetStreetSuffixOk

`func (o *CommonPAFAddress) GetStreetSuffixOk() (*string, bool)`

GetStreetSuffixOk returns a tuple with the StreetSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetSuffix

`func (o *CommonPAFAddress) SetStreetSuffix(v string)`

SetStreetSuffix sets StreetSuffix field to given value.

### HasStreetSuffix

`func (o *CommonPAFAddress) HasStreetSuffix() bool`

HasStreetSuffix returns a boolean if a field has been set.

### GetPostalDeliveryType

`func (o *CommonPAFAddress) GetPostalDeliveryType() string`

GetPostalDeliveryType returns the PostalDeliveryType field if non-nil, zero value otherwise.

### GetPostalDeliveryTypeOk

`func (o *CommonPAFAddress) GetPostalDeliveryTypeOk() (*string, bool)`

GetPostalDeliveryTypeOk returns a tuple with the PostalDeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalDeliveryType

`func (o *CommonPAFAddress) SetPostalDeliveryType(v string)`

SetPostalDeliveryType sets PostalDeliveryType field to given value.

### HasPostalDeliveryType

`func (o *CommonPAFAddress) HasPostalDeliveryType() bool`

HasPostalDeliveryType returns a boolean if a field has been set.

### GetPostalDeliveryNumber

`func (o *CommonPAFAddress) GetPostalDeliveryNumber() int32`

GetPostalDeliveryNumber returns the PostalDeliveryNumber field if non-nil, zero value otherwise.

### GetPostalDeliveryNumberOk

`func (o *CommonPAFAddress) GetPostalDeliveryNumberOk() (*int32, bool)`

GetPostalDeliveryNumberOk returns a tuple with the PostalDeliveryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalDeliveryNumber

`func (o *CommonPAFAddress) SetPostalDeliveryNumber(v int32)`

SetPostalDeliveryNumber sets PostalDeliveryNumber field to given value.

### HasPostalDeliveryNumber

`func (o *CommonPAFAddress) HasPostalDeliveryNumber() bool`

HasPostalDeliveryNumber returns a boolean if a field has been set.

### GetPostalDeliveryNumberPrefix

`func (o *CommonPAFAddress) GetPostalDeliveryNumberPrefix() string`

GetPostalDeliveryNumberPrefix returns the PostalDeliveryNumberPrefix field if non-nil, zero value otherwise.

### GetPostalDeliveryNumberPrefixOk

`func (o *CommonPAFAddress) GetPostalDeliveryNumberPrefixOk() (*string, bool)`

GetPostalDeliveryNumberPrefixOk returns a tuple with the PostalDeliveryNumberPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalDeliveryNumberPrefix

`func (o *CommonPAFAddress) SetPostalDeliveryNumberPrefix(v string)`

SetPostalDeliveryNumberPrefix sets PostalDeliveryNumberPrefix field to given value.

### HasPostalDeliveryNumberPrefix

`func (o *CommonPAFAddress) HasPostalDeliveryNumberPrefix() bool`

HasPostalDeliveryNumberPrefix returns a boolean if a field has been set.

### GetPostalDeliveryNumberSuffix

`func (o *CommonPAFAddress) GetPostalDeliveryNumberSuffix() string`

GetPostalDeliveryNumberSuffix returns the PostalDeliveryNumberSuffix field if non-nil, zero value otherwise.

### GetPostalDeliveryNumberSuffixOk

`func (o *CommonPAFAddress) GetPostalDeliveryNumberSuffixOk() (*string, bool)`

GetPostalDeliveryNumberSuffixOk returns a tuple with the PostalDeliveryNumberSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalDeliveryNumberSuffix

`func (o *CommonPAFAddress) SetPostalDeliveryNumberSuffix(v string)`

SetPostalDeliveryNumberSuffix sets PostalDeliveryNumberSuffix field to given value.

### HasPostalDeliveryNumberSuffix

`func (o *CommonPAFAddress) HasPostalDeliveryNumberSuffix() bool`

HasPostalDeliveryNumberSuffix returns a boolean if a field has been set.

### GetLocalityName

`func (o *CommonPAFAddress) GetLocalityName() string`

GetLocalityName returns the LocalityName field if non-nil, zero value otherwise.

### GetLocalityNameOk

`func (o *CommonPAFAddress) GetLocalityNameOk() (*string, bool)`

GetLocalityNameOk returns a tuple with the LocalityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityName

`func (o *CommonPAFAddress) SetLocalityName(v string)`

SetLocalityName sets LocalityName field to given value.


### GetPostcode

`func (o *CommonPAFAddress) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *CommonPAFAddress) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *CommonPAFAddress) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetState

`func (o *CommonPAFAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonPAFAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonPAFAddress) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


