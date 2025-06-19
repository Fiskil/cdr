# EnergyBillingOtherTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePointId** | Pointer to **string** | Unique identifier for the service point. | [optional] 
**InvoiceNumber** | Pointer to **string** | The number of the invoice in which this transaction is included if it has been issued. | [optional] 
**StartDate** | Pointer to **string** | Optional start date for the application of the charge. | [optional] 
**EndDate** | Pointer to **string** | Optional end date for the application of the charge. | [optional] 
**Type** | Pointer to **string** | Type of charge. Assumed to be &#x60;OTHER&#x60; if absent. | [optional] [default to "OTHER"]
**Amount** | **string** | The amount of the charge. | 
**Description** | **string** | A free text description of the item. | 
**CalculationFactors** | Pointer to [**[]EnergyBillingUsageTransactionV2CalculationFactors**](EnergyBillingUsageTransactionV2CalculationFactors.md) | Additional calculation factors that inform the transaction. | [optional] 
**Adjustments** | Pointer to [**[]EnergyBillingUsageTransactionV2Adjustments**](EnergyBillingUsageTransactionV2Adjustments.md) | Optional array of adjustments arising for this transaction. | [optional] 

## Methods

### NewEnergyBillingOtherTransaction

`func NewEnergyBillingOtherTransaction(amount string, description string, ) *EnergyBillingOtherTransaction`

NewEnergyBillingOtherTransaction instantiates a new EnergyBillingOtherTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBillingOtherTransactionWithDefaults

`func NewEnergyBillingOtherTransactionWithDefaults() *EnergyBillingOtherTransaction`

NewEnergyBillingOtherTransactionWithDefaults instantiates a new EnergyBillingOtherTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyBillingOtherTransaction) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyBillingOtherTransaction) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyBillingOtherTransaction) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.

### HasServicePointId

`func (o *EnergyBillingOtherTransaction) HasServicePointId() bool`

HasServicePointId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *EnergyBillingOtherTransaction) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *EnergyBillingOtherTransaction) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *EnergyBillingOtherTransaction) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *EnergyBillingOtherTransaction) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *EnergyBillingOtherTransaction) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyBillingOtherTransaction) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyBillingOtherTransaction) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EnergyBillingOtherTransaction) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EnergyBillingOtherTransaction) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyBillingOtherTransaction) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyBillingOtherTransaction) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EnergyBillingOtherTransaction) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetType

`func (o *EnergyBillingOtherTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyBillingOtherTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyBillingOtherTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnergyBillingOtherTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *EnergyBillingOtherTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyBillingOtherTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyBillingOtherTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *EnergyBillingOtherTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyBillingOtherTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyBillingOtherTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCalculationFactors

`func (o *EnergyBillingOtherTransaction) GetCalculationFactors() []EnergyBillingUsageTransactionV2CalculationFactors`

GetCalculationFactors returns the CalculationFactors field if non-nil, zero value otherwise.

### GetCalculationFactorsOk

`func (o *EnergyBillingOtherTransaction) GetCalculationFactorsOk() (*[]EnergyBillingUsageTransactionV2CalculationFactors, bool)`

GetCalculationFactorsOk returns a tuple with the CalculationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFactors

`func (o *EnergyBillingOtherTransaction) SetCalculationFactors(v []EnergyBillingUsageTransactionV2CalculationFactors)`

SetCalculationFactors sets CalculationFactors field to given value.

### HasCalculationFactors

`func (o *EnergyBillingOtherTransaction) HasCalculationFactors() bool`

HasCalculationFactors returns a boolean if a field has been set.

### GetAdjustments

`func (o *EnergyBillingOtherTransaction) GetAdjustments() []EnergyBillingUsageTransactionV2Adjustments`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *EnergyBillingOtherTransaction) GetAdjustmentsOk() (*[]EnergyBillingUsageTransactionV2Adjustments, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *EnergyBillingOtherTransaction) SetAdjustments(v []EnergyBillingUsageTransactionV2Adjustments)`

SetAdjustments sets Adjustments field to given value.

### HasAdjustments

`func (o *EnergyBillingOtherTransaction) HasAdjustments() bool`

HasAdjustments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


