# MetaPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRecords** | **int32** | The total number of records in the full set. See [pagination](#pagination). | 
**TotalPages** | **int32** | The total number of pages in the full set. See [pagination](#pagination). | 

## Methods

### NewMetaPaginated

`func NewMetaPaginated(totalRecords int32, totalPages int32, ) *MetaPaginated`

NewMetaPaginated instantiates a new MetaPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPaginatedWithDefaults

`func NewMetaPaginatedWithDefaults() *MetaPaginated`

NewMetaPaginatedWithDefaults instantiates a new MetaPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRecords

`func (o *MetaPaginated) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *MetaPaginated) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *MetaPaginated) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.


### GetTotalPages

`func (o *MetaPaginated) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *MetaPaginated) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *MetaPaginated) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


