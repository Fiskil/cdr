# EnergyBillingUsageTransactionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePointId** | Pointer to **string** | Unique identifier for the service point. | [optional] 
**InvoiceNumber** | Pointer to **string** | The number of the invoice in which this transaction is included if it has been issued. | [optional] 
**TimeOfUseType** | **string** | The time of use type that the transaction applies to. | 
**Description** | Pointer to **string** | Optional description of the transaction that can be used for display purposes. | [optional] 
**IsEstimate** | Pointer to **bool** | Flag indicating if the usage is estimated or actual. &#x60;true&#x60; indicates estimate. &#x60;false&#x60; or absent indicates actual. | [optional] [default to false]
**StartDate** | **string** | Date and time when the usage period starts. | 
**EndDate** | **string** | Date and time when the usage period ends. | 
**MeasureUnit** | Pointer to [**MeasureUnitEnum**](MeasureUnitEnum.md) | The measurement unit of rate. Assumed to be &#x60;KWH&#x60; if absent. | [optional] [default to KWH]
**Usage** | **float32** | The usage for the period in measure unit. A negative value indicates power generated. | 
**Amount** | **string** | The amount charged or credited for this transaction prior to any adjustments being applied. A negative value indicates a credit. | 
**CalculationFactors** | Pointer to [**[]EnergyBillingUsageTransactionV2CalculationFactors**](EnergyBillingUsageTransactionV2CalculationFactors.md) | Additional calculation factors that inform the transaction. | [optional] 
**Adjustments** | Pointer to [**[]EnergyBillingUsageTransactionV2Adjustments**](EnergyBillingUsageTransactionV2Adjustments.md) | Optional array of adjustments arising for this transaction. | [optional] 

## Methods

### NewEnergyBillingUsageTransactionV2

`func NewEnergyBillingUsageTransactionV2(timeOfUseType string, startDate string, endDate string, usage float32, amount string, ) *EnergyBillingUsageTransactionV2`

NewEnergyBillingUsageTransactionV2 instantiates a new EnergyBillingUsageTransactionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBillingUsageTransactionV2WithDefaults

`func NewEnergyBillingUsageTransactionV2WithDefaults() *EnergyBillingUsageTransactionV2`

NewEnergyBillingUsageTransactionV2WithDefaults instantiates a new EnergyBillingUsageTransactionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyBillingUsageTransactionV2) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyBillingUsageTransactionV2) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyBillingUsageTransactionV2) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.

### HasServicePointId

`func (o *EnergyBillingUsageTransactionV2) HasServicePointId() bool`

HasServicePointId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *EnergyBillingUsageTransactionV2) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *EnergyBillingUsageTransactionV2) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *EnergyBillingUsageTransactionV2) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *EnergyBillingUsageTransactionV2) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetTimeOfUseType

`func (o *EnergyBillingUsageTransactionV2) GetTimeOfUseType() string`

GetTimeOfUseType returns the TimeOfUseType field if non-nil, zero value otherwise.

### GetTimeOfUseTypeOk

`func (o *EnergyBillingUsageTransactionV2) GetTimeOfUseTypeOk() (*string, bool)`

GetTimeOfUseTypeOk returns a tuple with the TimeOfUseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfUseType

`func (o *EnergyBillingUsageTransactionV2) SetTimeOfUseType(v string)`

SetTimeOfUseType sets TimeOfUseType field to given value.


### GetDescription

`func (o *EnergyBillingUsageTransactionV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyBillingUsageTransactionV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyBillingUsageTransactionV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyBillingUsageTransactionV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEstimate

`func (o *EnergyBillingUsageTransactionV2) GetIsEstimate() bool`

GetIsEstimate returns the IsEstimate field if non-nil, zero value otherwise.

### GetIsEstimateOk

`func (o *EnergyBillingUsageTransactionV2) GetIsEstimateOk() (*bool, bool)`

GetIsEstimateOk returns a tuple with the IsEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEstimate

`func (o *EnergyBillingUsageTransactionV2) SetIsEstimate(v bool)`

SetIsEstimate sets IsEstimate field to given value.

### HasIsEstimate

`func (o *EnergyBillingUsageTransactionV2) HasIsEstimate() bool`

HasIsEstimate returns a boolean if a field has been set.

### GetStartDate

`func (o *EnergyBillingUsageTransactionV2) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyBillingUsageTransactionV2) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyBillingUsageTransactionV2) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *EnergyBillingUsageTransactionV2) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyBillingUsageTransactionV2) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyBillingUsageTransactionV2) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetMeasureUnit

`func (o *EnergyBillingUsageTransactionV2) GetMeasureUnit() MeasureUnitEnum`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *EnergyBillingUsageTransactionV2) GetMeasureUnitOk() (*MeasureUnitEnum, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *EnergyBillingUsageTransactionV2) SetMeasureUnit(v MeasureUnitEnum)`

SetMeasureUnit sets MeasureUnit field to given value.

### HasMeasureUnit

`func (o *EnergyBillingUsageTransactionV2) HasMeasureUnit() bool`

HasMeasureUnit returns a boolean if a field has been set.

### GetUsage

`func (o *EnergyBillingUsageTransactionV2) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EnergyBillingUsageTransactionV2) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EnergyBillingUsageTransactionV2) SetUsage(v float32)`

SetUsage sets Usage field to given value.


### GetAmount

`func (o *EnergyBillingUsageTransactionV2) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyBillingUsageTransactionV2) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyBillingUsageTransactionV2) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCalculationFactors

`func (o *EnergyBillingUsageTransactionV2) GetCalculationFactors() []EnergyBillingUsageTransactionV2CalculationFactors`

GetCalculationFactors returns the CalculationFactors field if non-nil, zero value otherwise.

### GetCalculationFactorsOk

`func (o *EnergyBillingUsageTransactionV2) GetCalculationFactorsOk() (*[]EnergyBillingUsageTransactionV2CalculationFactors, bool)`

GetCalculationFactorsOk returns a tuple with the CalculationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFactors

`func (o *EnergyBillingUsageTransactionV2) SetCalculationFactors(v []EnergyBillingUsageTransactionV2CalculationFactors)`

SetCalculationFactors sets CalculationFactors field to given value.

### HasCalculationFactors

`func (o *EnergyBillingUsageTransactionV2) HasCalculationFactors() bool`

HasCalculationFactors returns a boolean if a field has been set.

### GetAdjustments

`func (o *EnergyBillingUsageTransactionV2) GetAdjustments() []EnergyBillingUsageTransactionV2Adjustments`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *EnergyBillingUsageTransactionV2) GetAdjustmentsOk() (*[]EnergyBillingUsageTransactionV2Adjustments, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *EnergyBillingUsageTransactionV2) SetAdjustments(v []EnergyBillingUsageTransactionV2Adjustments)`

SetAdjustments sets Adjustments field to given value.

### HasAdjustments

`func (o *EnergyBillingUsageTransactionV2) HasAdjustments() bool`

HasAdjustments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


