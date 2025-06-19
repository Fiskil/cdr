# EnergyInvoicePayOnTimeDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountAmount** | **string** | The amount that will be discounted if the invoice is paid by the date specified. | 
**GstAmount** | Pointer to **string** | The GST amount that will be discounted if the invoice is paid by the date specified. If absent then zero is assumed. | [optional] [default to "0.00"]
**Date** | **string** | The date by which the invoice must be paid to receive the pay on time discount. | 

## Methods

### NewEnergyInvoicePayOnTimeDiscount

`func NewEnergyInvoicePayOnTimeDiscount(discountAmount string, date string, ) *EnergyInvoicePayOnTimeDiscount`

NewEnergyInvoicePayOnTimeDiscount instantiates a new EnergyInvoicePayOnTimeDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyInvoicePayOnTimeDiscountWithDefaults

`func NewEnergyInvoicePayOnTimeDiscountWithDefaults() *EnergyInvoicePayOnTimeDiscount`

NewEnergyInvoicePayOnTimeDiscountWithDefaults instantiates a new EnergyInvoicePayOnTimeDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountAmount

`func (o *EnergyInvoicePayOnTimeDiscount) GetDiscountAmount() string`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *EnergyInvoicePayOnTimeDiscount) GetDiscountAmountOk() (*string, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *EnergyInvoicePayOnTimeDiscount) SetDiscountAmount(v string)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetGstAmount

`func (o *EnergyInvoicePayOnTimeDiscount) GetGstAmount() string`

GetGstAmount returns the GstAmount field if non-nil, zero value otherwise.

### GetGstAmountOk

`func (o *EnergyInvoicePayOnTimeDiscount) GetGstAmountOk() (*string, bool)`

GetGstAmountOk returns a tuple with the GstAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGstAmount

`func (o *EnergyInvoicePayOnTimeDiscount) SetGstAmount(v string)`

SetGstAmount sets GstAmount field to given value.

### HasGstAmount

`func (o *EnergyInvoicePayOnTimeDiscount) HasGstAmount() bool`

HasGstAmount returns a boolean if a field has been set.

### GetDate

`func (o *EnergyInvoicePayOnTimeDiscount) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EnergyInvoicePayOnTimeDiscount) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EnergyInvoicePayOnTimeDiscount) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


