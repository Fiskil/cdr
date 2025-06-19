# EnergyInvoiceAccountCharges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCharges** | **string** | The aggregate total of account level charges for the period covered by the invoice. | 
**TotalDiscounts** | **string** | The aggregate total of account level discounts or credits for the period covered by the invoice. | 
**TotalGst** | Pointer to **string** | The total GST for all account level charges. If absent then zero is assumed. | [optional] [default to "0.00"]

## Methods

### NewEnergyInvoiceAccountCharges

`func NewEnergyInvoiceAccountCharges(totalCharges string, totalDiscounts string, ) *EnergyInvoiceAccountCharges`

NewEnergyInvoiceAccountCharges instantiates a new EnergyInvoiceAccountCharges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyInvoiceAccountChargesWithDefaults

`func NewEnergyInvoiceAccountChargesWithDefaults() *EnergyInvoiceAccountCharges`

NewEnergyInvoiceAccountChargesWithDefaults instantiates a new EnergyInvoiceAccountCharges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCharges

`func (o *EnergyInvoiceAccountCharges) GetTotalCharges() string`

GetTotalCharges returns the TotalCharges field if non-nil, zero value otherwise.

### GetTotalChargesOk

`func (o *EnergyInvoiceAccountCharges) GetTotalChargesOk() (*string, bool)`

GetTotalChargesOk returns a tuple with the TotalCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCharges

`func (o *EnergyInvoiceAccountCharges) SetTotalCharges(v string)`

SetTotalCharges sets TotalCharges field to given value.


### GetTotalDiscounts

`func (o *EnergyInvoiceAccountCharges) GetTotalDiscounts() string`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *EnergyInvoiceAccountCharges) GetTotalDiscountsOk() (*string, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *EnergyInvoiceAccountCharges) SetTotalDiscounts(v string)`

SetTotalDiscounts sets TotalDiscounts field to given value.


### GetTotalGst

`func (o *EnergyInvoiceAccountCharges) GetTotalGst() string`

GetTotalGst returns the TotalGst field if non-nil, zero value otherwise.

### GetTotalGstOk

`func (o *EnergyInvoiceAccountCharges) GetTotalGstOk() (*string, bool)`

GetTotalGstOk returns a tuple with the TotalGst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGst

`func (o *EnergyInvoiceAccountCharges) SetTotalGst(v string)`

SetTotalGst sets TotalGst field to given value.

### HasTotalGst

`func (o *EnergyInvoiceAccountCharges) HasTotalGst() bool`

HasTotalGst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


