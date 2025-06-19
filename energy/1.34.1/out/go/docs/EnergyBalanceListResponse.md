# EnergyBalanceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyBalanceListResponseData**](EnergyBalanceListResponseData.md) |  | 
**Links** | [**LinksPaginated**](LinksPaginated.md) |  | 
**Meta** | [**MetaPaginated**](MetaPaginated.md) |  | 

## Methods

### NewEnergyBalanceListResponse

`func NewEnergyBalanceListResponse(data EnergyBalanceListResponseData, links LinksPaginated, meta MetaPaginated, ) *EnergyBalanceListResponse`

NewEnergyBalanceListResponse instantiates a new EnergyBalanceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBalanceListResponseWithDefaults

`func NewEnergyBalanceListResponseWithDefaults() *EnergyBalanceListResponse`

NewEnergyBalanceListResponseWithDefaults instantiates a new EnergyBalanceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyBalanceListResponse) GetData() EnergyBalanceListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyBalanceListResponse) GetDataOk() (*EnergyBalanceListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyBalanceListResponse) SetData(v EnergyBalanceListResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyBalanceListResponse) GetLinks() LinksPaginated`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyBalanceListResponse) GetLinksOk() (*LinksPaginated, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyBalanceListResponse) SetLinks(v LinksPaginated)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyBalanceListResponse) GetMeta() MetaPaginated`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyBalanceListResponse) GetMetaOk() (*MetaPaginated, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyBalanceListResponse) SetMeta(v MetaPaginated)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


