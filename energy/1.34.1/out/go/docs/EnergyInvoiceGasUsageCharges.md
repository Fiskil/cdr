# EnergyInvoiceGasUsageCharges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalUsageCharges** | **string** | The aggregate total of usage charges for the period covered by the invoice (exclusive of GST). | 
**TotalGenerationCredits** | **string** | The aggregate total of generation credits for the period covered by the invoice (exclusive of GST). | 
**TotalOnceOffCharges** | **string** | The aggregate total of any once off charges arising from gas usage for the period covered by the invoice (exclusive of GST). | 
**TotalOnceOffDiscounts** | **string** | The aggregate total of any once off discounts or credits arising from gas usage for the period covered by the invoice (exclusive of GST). | 
**OtherCharges** | Pointer to [**[]EnergyInvoiceGasUsageChargesOtherCharges**](EnergyInvoiceGasUsageChargesOtherCharges.md) | Optional array of charges that may be part of the invoice (for e.g., environmental charges for C&amp;I consumers) (exclusive of GST). | [optional] 
**TotalGst** | Pointer to **string** | The total GST for all gas usage charges. If absent then zero is assumed. | [optional] [default to "0.00"]

## Methods

### NewEnergyInvoiceGasUsageCharges

`func NewEnergyInvoiceGasUsageCharges(totalUsageCharges string, totalGenerationCredits string, totalOnceOffCharges string, totalOnceOffDiscounts string, ) *EnergyInvoiceGasUsageCharges`

NewEnergyInvoiceGasUsageCharges instantiates a new EnergyInvoiceGasUsageCharges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyInvoiceGasUsageChargesWithDefaults

`func NewEnergyInvoiceGasUsageChargesWithDefaults() *EnergyInvoiceGasUsageCharges`

NewEnergyInvoiceGasUsageChargesWithDefaults instantiates a new EnergyInvoiceGasUsageCharges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalUsageCharges

`func (o *EnergyInvoiceGasUsageCharges) GetTotalUsageCharges() string`

GetTotalUsageCharges returns the TotalUsageCharges field if non-nil, zero value otherwise.

### GetTotalUsageChargesOk

`func (o *EnergyInvoiceGasUsageCharges) GetTotalUsageChargesOk() (*string, bool)`

GetTotalUsageChargesOk returns a tuple with the TotalUsageCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCharges

`func (o *EnergyInvoiceGasUsageCharges) SetTotalUsageCharges(v string)`

SetTotalUsageCharges sets TotalUsageCharges field to given value.


### GetTotalGenerationCredits

`func (o *EnergyInvoiceGasUsageCharges) GetTotalGenerationCredits() string`

GetTotalGenerationCredits returns the TotalGenerationCredits field if non-nil, zero value otherwise.

### GetTotalGenerationCreditsOk

`func (o *EnergyInvoiceGasUsageCharges) GetTotalGenerationCreditsOk() (*string, bool)`

GetTotalGenerationCreditsOk returns a tuple with the TotalGenerationCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGenerationCredits

`func (o *EnergyInvoiceGasUsageCharges) SetTotalGenerationCredits(v string)`

SetTotalGenerationCredits sets TotalGenerationCredits field to given value.


### GetTotalOnceOffCharges

`func (o *EnergyInvoiceGasUsageCharges) GetTotalOnceOffCharges() string`

GetTotalOnceOffCharges returns the TotalOnceOffCharges field if non-nil, zero value otherwise.

### GetTotalOnceOffChargesOk

`func (o *EnergyInvoiceGasUsageCharges) GetTotalOnceOffChargesOk() (*string, bool)`

GetTotalOnceOffChargesOk returns a tuple with the TotalOnceOffCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnceOffCharges

`func (o *EnergyInvoiceGasUsageCharges) SetTotalOnceOffCharges(v string)`

SetTotalOnceOffCharges sets TotalOnceOffCharges field to given value.


### GetTotalOnceOffDiscounts

`func (o *EnergyInvoiceGasUsageCharges) GetTotalOnceOffDiscounts() string`

GetTotalOnceOffDiscounts returns the TotalOnceOffDiscounts field if non-nil, zero value otherwise.

### GetTotalOnceOffDiscountsOk

`func (o *EnergyInvoiceGasUsageCharges) GetTotalOnceOffDiscountsOk() (*string, bool)`

GetTotalOnceOffDiscountsOk returns a tuple with the TotalOnceOffDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnceOffDiscounts

`func (o *EnergyInvoiceGasUsageCharges) SetTotalOnceOffDiscounts(v string)`

SetTotalOnceOffDiscounts sets TotalOnceOffDiscounts field to given value.


### GetOtherCharges

`func (o *EnergyInvoiceGasUsageCharges) GetOtherCharges() []EnergyInvoiceGasUsageChargesOtherCharges`

GetOtherCharges returns the OtherCharges field if non-nil, zero value otherwise.

### GetOtherChargesOk

`func (o *EnergyInvoiceGasUsageCharges) GetOtherChargesOk() (*[]EnergyInvoiceGasUsageChargesOtherCharges, bool)`

GetOtherChargesOk returns a tuple with the OtherCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCharges

`func (o *EnergyInvoiceGasUsageCharges) SetOtherCharges(v []EnergyInvoiceGasUsageChargesOtherCharges)`

SetOtherCharges sets OtherCharges field to given value.

### HasOtherCharges

`func (o *EnergyInvoiceGasUsageCharges) HasOtherCharges() bool`

HasOtherCharges returns a boolean if a field has been set.

### GetTotalGst

`func (o *EnergyInvoiceGasUsageCharges) GetTotalGst() string`

GetTotalGst returns the TotalGst field if non-nil, zero value otherwise.

### GetTotalGstOk

`func (o *EnergyInvoiceGasUsageCharges) GetTotalGstOk() (*string, bool)`

GetTotalGstOk returns a tuple with the TotalGst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGst

`func (o *EnergyInvoiceGasUsageCharges) SetTotalGst(v string)`

SetTotalGst sets TotalGst field to given value.

### HasTotalGst

`func (o *EnergyInvoiceGasUsageCharges) HasTotalGst() bool`

HasTotalGst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


