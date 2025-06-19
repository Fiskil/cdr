# EnergyBillingOnceOffTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePointId** | Pointer to **string** | Unique identifier for the service point. | [optional] 
**InvoiceNumber** | Pointer to **string** | The number of the invoice in which this transaction is included if it has been issued. | [optional] 
**Amount** | **string** | The amount of the charge or credit. A positive value indicates a charge and a negative value indicates a credit. | 
**Description** | **string** | A free text description of the item. | 

## Methods

### NewEnergyBillingOnceOffTransaction

`func NewEnergyBillingOnceOffTransaction(amount string, description string, ) *EnergyBillingOnceOffTransaction`

NewEnergyBillingOnceOffTransaction instantiates a new EnergyBillingOnceOffTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBillingOnceOffTransactionWithDefaults

`func NewEnergyBillingOnceOffTransactionWithDefaults() *EnergyBillingOnceOffTransaction`

NewEnergyBillingOnceOffTransactionWithDefaults instantiates a new EnergyBillingOnceOffTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyBillingOnceOffTransaction) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyBillingOnceOffTransaction) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyBillingOnceOffTransaction) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.

### HasServicePointId

`func (o *EnergyBillingOnceOffTransaction) HasServicePointId() bool`

HasServicePointId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *EnergyBillingOnceOffTransaction) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *EnergyBillingOnceOffTransaction) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *EnergyBillingOnceOffTransaction) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *EnergyBillingOnceOffTransaction) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetAmount

`func (o *EnergyBillingOnceOffTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyBillingOnceOffTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyBillingOnceOffTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *EnergyBillingOnceOffTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyBillingOnceOffTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyBillingOnceOffTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


