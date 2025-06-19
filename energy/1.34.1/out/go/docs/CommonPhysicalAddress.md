# CommonPhysicalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressUType** | **string** | The type of address object present. | 
**Simple** | Pointer to [**CommonSimpleAddress**](CommonSimpleAddress.md) |  | [optional] 
**Paf** | Pointer to [**CommonPAFAddress**](CommonPAFAddress.md) |  | [optional] 

## Methods

### NewCommonPhysicalAddress

`func NewCommonPhysicalAddress(addressUType string, ) *CommonPhysicalAddress`

NewCommonPhysicalAddress instantiates a new CommonPhysicalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPhysicalAddressWithDefaults

`func NewCommonPhysicalAddressWithDefaults() *CommonPhysicalAddress`

NewCommonPhysicalAddressWithDefaults instantiates a new CommonPhysicalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressUType

`func (o *CommonPhysicalAddress) GetAddressUType() string`

GetAddressUType returns the AddressUType field if non-nil, zero value otherwise.

### GetAddressUTypeOk

`func (o *CommonPhysicalAddress) GetAddressUTypeOk() (*string, bool)`

GetAddressUTypeOk returns a tuple with the AddressUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressUType

`func (o *CommonPhysicalAddress) SetAddressUType(v string)`

SetAddressUType sets AddressUType field to given value.


### GetSimple

`func (o *CommonPhysicalAddress) GetSimple() CommonSimpleAddress`

GetSimple returns the Simple field if non-nil, zero value otherwise.

### GetSimpleOk

`func (o *CommonPhysicalAddress) GetSimpleOk() (*CommonSimpleAddress, bool)`

GetSimpleOk returns a tuple with the Simple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimple

`func (o *CommonPhysicalAddress) SetSimple(v CommonSimpleAddress)`

SetSimple sets Simple field to given value.

### HasSimple

`func (o *CommonPhysicalAddress) HasSimple() bool`

HasSimple returns a boolean if a field has been set.

### GetPaf

`func (o *CommonPhysicalAddress) GetPaf() CommonPAFAddress`

GetPaf returns the Paf field if non-nil, zero value otherwise.

### GetPafOk

`func (o *CommonPhysicalAddress) GetPafOk() (*CommonPAFAddress, bool)`

GetPafOk returns a tuple with the Paf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaf

`func (o *CommonPhysicalAddress) SetPaf(v CommonPAFAddress)`

SetPaf sets Paf field to given value.

### HasPaf

`func (o *CommonPhysicalAddress) HasPaf() bool`

HasPaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


