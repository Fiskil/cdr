# EnergyDerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyDerListResponseData**](EnergyDerListResponseData.md) |  | 
**Links** | [**LinksPaginated**](LinksPaginated.md) |  | 
**Meta** | [**MetaPaginated**](MetaPaginated.md) |  | 

## Methods

### NewEnergyDerListResponse

`func NewEnergyDerListResponse(data EnergyDerListResponseData, links LinksPaginated, meta MetaPaginated, ) *EnergyDerListResponse`

NewEnergyDerListResponse instantiates a new EnergyDerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyDerListResponseWithDefaults

`func NewEnergyDerListResponseWithDefaults() *EnergyDerListResponse`

NewEnergyDerListResponseWithDefaults instantiates a new EnergyDerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyDerListResponse) GetData() EnergyDerListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyDerListResponse) GetDataOk() (*EnergyDerListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyDerListResponse) SetData(v EnergyDerListResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyDerListResponse) GetLinks() LinksPaginated`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyDerListResponse) GetLinksOk() (*LinksPaginated, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyDerListResponse) SetLinks(v LinksPaginated)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyDerListResponse) GetMeta() MetaPaginated`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyDerListResponse) GetMetaOk() (*MetaPaginated, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyDerListResponse) SetMeta(v MetaPaginated)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


