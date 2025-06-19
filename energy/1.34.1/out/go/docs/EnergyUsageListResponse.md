# EnergyUsageListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyUsageListResponseData**](EnergyUsageListResponseData.md) |  | 
**Links** | [**LinksPaginated**](LinksPaginated.md) |  | 
**Meta** | [**MetaPaginated**](MetaPaginated.md) |  | 

## Methods

### NewEnergyUsageListResponse

`func NewEnergyUsageListResponse(data EnergyUsageListResponseData, links LinksPaginated, meta MetaPaginated, ) *EnergyUsageListResponse`

NewEnergyUsageListResponse instantiates a new EnergyUsageListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyUsageListResponseWithDefaults

`func NewEnergyUsageListResponseWithDefaults() *EnergyUsageListResponse`

NewEnergyUsageListResponseWithDefaults instantiates a new EnergyUsageListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyUsageListResponse) GetData() EnergyUsageListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyUsageListResponse) GetDataOk() (*EnergyUsageListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyUsageListResponse) SetData(v EnergyUsageListResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyUsageListResponse) GetLinks() LinksPaginated`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyUsageListResponse) GetLinksOk() (*LinksPaginated, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyUsageListResponse) SetLinks(v LinksPaginated)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyUsageListResponse) GetMeta() MetaPaginated`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyUsageListResponse) GetMetaOk() (*MetaPaginated, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyUsageListResponse) SetMeta(v MetaPaginated)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


