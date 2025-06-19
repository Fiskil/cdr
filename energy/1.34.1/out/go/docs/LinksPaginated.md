# LinksPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** | Fully qualified link that generated the current response document. | 
**First** | Pointer to **string** | URI to the first page of this set. Mandatory if this response is not the first page. | [optional] 
**Prev** | Pointer to **string** | URI to the previous page of this set. Mandatory if this response is not the first page. | [optional] 
**Next** | Pointer to **string** | URI to the next page of this set. Mandatory if this response is not the last page. | [optional] 
**Last** | Pointer to **string** | URI to the last page of this set. Mandatory if this response is not the last page. | [optional] 

## Methods

### NewLinksPaginated

`func NewLinksPaginated(self string, ) *LinksPaginated`

NewLinksPaginated instantiates a new LinksPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksPaginatedWithDefaults

`func NewLinksPaginatedWithDefaults() *LinksPaginated`

NewLinksPaginatedWithDefaults instantiates a new LinksPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *LinksPaginated) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LinksPaginated) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LinksPaginated) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetFirst

`func (o *LinksPaginated) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *LinksPaginated) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *LinksPaginated) SetFirst(v string)`

SetFirst sets First field to given value.

### HasFirst

`func (o *LinksPaginated) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetPrev

`func (o *LinksPaginated) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *LinksPaginated) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *LinksPaginated) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *LinksPaginated) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetNext

`func (o *LinksPaginated) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *LinksPaginated) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *LinksPaginated) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *LinksPaginated) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetLast

`func (o *LinksPaginated) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *LinksPaginated) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *LinksPaginated) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *LinksPaginated) HasLast() bool`

HasLast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


