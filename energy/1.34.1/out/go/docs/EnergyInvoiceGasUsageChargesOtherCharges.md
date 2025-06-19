# EnergyInvoiceGasUsageChargesOtherCharges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of charge. Assumed to be &#x60;OTHER&#x60; if absent. | [optional] [default to "OTHER"]
**Amount** | **string** | The aggregate total of charges for this item (exclusive of GST). | 
**Description** | **string** | A free text description of the type of charge. | 

## Methods

### NewEnergyInvoiceGasUsageChargesOtherCharges

`func NewEnergyInvoiceGasUsageChargesOtherCharges(amount string, description string, ) *EnergyInvoiceGasUsageChargesOtherCharges`

NewEnergyInvoiceGasUsageChargesOtherCharges instantiates a new EnergyInvoiceGasUsageChargesOtherCharges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyInvoiceGasUsageChargesOtherChargesWithDefaults

`func NewEnergyInvoiceGasUsageChargesOtherChargesWithDefaults() *EnergyInvoiceGasUsageChargesOtherCharges`

NewEnergyInvoiceGasUsageChargesOtherChargesWithDefaults instantiates a new EnergyInvoiceGasUsageChargesOtherCharges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyInvoiceGasUsageChargesOtherCharges) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


