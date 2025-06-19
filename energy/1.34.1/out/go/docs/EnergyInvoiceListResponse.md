# EnergyInvoiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnergyInvoiceListResponseData**](EnergyInvoiceListResponseData.md) |  | 
**Links** | [**LinksPaginated**](LinksPaginated.md) |  | 
**Meta** | [**MetaPaginated**](MetaPaginated.md) |  | 

## Methods

### NewEnergyInvoiceListResponse

`func NewEnergyInvoiceListResponse(data EnergyInvoiceListResponseData, links LinksPaginated, meta MetaPaginated, ) *EnergyInvoiceListResponse`

NewEnergyInvoiceListResponse instantiates a new EnergyInvoiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyInvoiceListResponseWithDefaults

`func NewEnergyInvoiceListResponseWithDefaults() *EnergyInvoiceListResponse`

NewEnergyInvoiceListResponseWithDefaults instantiates a new EnergyInvoiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnergyInvoiceListResponse) GetData() EnergyInvoiceListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnergyInvoiceListResponse) GetDataOk() (*EnergyInvoiceListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnergyInvoiceListResponse) SetData(v EnergyInvoiceListResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *EnergyInvoiceListResponse) GetLinks() LinksPaginated`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnergyInvoiceListResponse) GetLinksOk() (*LinksPaginated, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnergyInvoiceListResponse) SetLinks(v LinksPaginated)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EnergyInvoiceListResponse) GetMeta() MetaPaginated`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnergyInvoiceListResponse) GetMetaOk() (*MetaPaginated, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnergyInvoiceListResponse) SetMeta(v MetaPaginated)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


