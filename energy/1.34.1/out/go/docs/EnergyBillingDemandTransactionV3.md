# EnergyBillingDemandTransactionV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePointId** | Pointer to **string** | Unique identifier for the service point. | [optional] 
**InvoiceNumber** | Pointer to **string** | The number of the invoice in which this transaction is included if it has been issued. | [optional] 
**TimeOfUseType** | **string** | The time of use type that the transaction applies to. | 
**Description** | Pointer to **string** | Optional description of the transaction that can be used for display purposes. | [optional] 
**IsEstimate** | Pointer to **bool** | Flag indicating if the usage is estimated or actual. &#x60;true&#x60; indicates estimate. &#x60;false&#x60; or absent indicates actual. | [optional] [default to false]
**StartDate** | **string** | Date and time when the demand period starts. | 
**EndDate** | **string** | Date and time when the demand period ends. | 
**MeasureUnit** | Pointer to [**MeasureUnitEnum**](MeasureUnitEnum.md) | The measurement unit of rate. Assumed to be &#x60;KVA&#x60; if absent. | [optional] [default to KVA]
**Rate** | **float32** | The rate for the demand charge in _measureUnit_. Assumed to be &#x60;KVA&#x60; if _measureUnit_ not provided. A negative value indicates power generated. | 
**Amount** | **string** | The amount charged or credited for this transaction prior to any adjustments being applied. A negative value indicates a credit. | 
**CalculationFactors** | Pointer to [**[]EnergyBillingUsageTransactionV2CalculationFactors**](EnergyBillingUsageTransactionV2CalculationFactors.md) | Additional calculation factors that inform the transaction. | [optional] 
**Adjustments** | Pointer to [**[]EnergyBillingUsageTransactionV2Adjustments**](EnergyBillingUsageTransactionV2Adjustments.md) | Optional array of adjustments arising for this transaction. | [optional] 

## Methods

### NewEnergyBillingDemandTransactionV3

`func NewEnergyBillingDemandTransactionV3(timeOfUseType string, startDate string, endDate string, rate float32, amount string, ) *EnergyBillingDemandTransactionV3`

NewEnergyBillingDemandTransactionV3 instantiates a new EnergyBillingDemandTransactionV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBillingDemandTransactionV3WithDefaults

`func NewEnergyBillingDemandTransactionV3WithDefaults() *EnergyBillingDemandTransactionV3`

NewEnergyBillingDemandTransactionV3WithDefaults instantiates a new EnergyBillingDemandTransactionV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePointId

`func (o *EnergyBillingDemandTransactionV3) GetServicePointId() string`

GetServicePointId returns the ServicePointId field if non-nil, zero value otherwise.

### GetServicePointIdOk

`func (o *EnergyBillingDemandTransactionV3) GetServicePointIdOk() (*string, bool)`

GetServicePointIdOk returns a tuple with the ServicePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePointId

`func (o *EnergyBillingDemandTransactionV3) SetServicePointId(v string)`

SetServicePointId sets ServicePointId field to given value.

### HasServicePointId

`func (o *EnergyBillingDemandTransactionV3) HasServicePointId() bool`

HasServicePointId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *EnergyBillingDemandTransactionV3) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *EnergyBillingDemandTransactionV3) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *EnergyBillingDemandTransactionV3) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *EnergyBillingDemandTransactionV3) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetTimeOfUseType

`func (o *EnergyBillingDemandTransactionV3) GetTimeOfUseType() string`

GetTimeOfUseType returns the TimeOfUseType field if non-nil, zero value otherwise.

### GetTimeOfUseTypeOk

`func (o *EnergyBillingDemandTransactionV3) GetTimeOfUseTypeOk() (*string, bool)`

GetTimeOfUseTypeOk returns a tuple with the TimeOfUseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfUseType

`func (o *EnergyBillingDemandTransactionV3) SetTimeOfUseType(v string)`

SetTimeOfUseType sets TimeOfUseType field to given value.


### GetDescription

`func (o *EnergyBillingDemandTransactionV3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyBillingDemandTransactionV3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyBillingDemandTransactionV3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyBillingDemandTransactionV3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEstimate

`func (o *EnergyBillingDemandTransactionV3) GetIsEstimate() bool`

GetIsEstimate returns the IsEstimate field if non-nil, zero value otherwise.

### GetIsEstimateOk

`func (o *EnergyBillingDemandTransactionV3) GetIsEstimateOk() (*bool, bool)`

GetIsEstimateOk returns a tuple with the IsEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEstimate

`func (o *EnergyBillingDemandTransactionV3) SetIsEstimate(v bool)`

SetIsEstimate sets IsEstimate field to given value.

### HasIsEstimate

`func (o *EnergyBillingDemandTransactionV3) HasIsEstimate() bool`

HasIsEstimate returns a boolean if a field has been set.

### GetStartDate

`func (o *EnergyBillingDemandTransactionV3) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyBillingDemandTransactionV3) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyBillingDemandTransactionV3) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *EnergyBillingDemandTransactionV3) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyBillingDemandTransactionV3) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyBillingDemandTransactionV3) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetMeasureUnit

`func (o *EnergyBillingDemandTransactionV3) GetMeasureUnit() MeasureUnitEnum`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *EnergyBillingDemandTransactionV3) GetMeasureUnitOk() (*MeasureUnitEnum, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *EnergyBillingDemandTransactionV3) SetMeasureUnit(v MeasureUnitEnum)`

SetMeasureUnit sets MeasureUnit field to given value.

### HasMeasureUnit

`func (o *EnergyBillingDemandTransactionV3) HasMeasureUnit() bool`

HasMeasureUnit returns a boolean if a field has been set.

### GetRate

`func (o *EnergyBillingDemandTransactionV3) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *EnergyBillingDemandTransactionV3) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *EnergyBillingDemandTransactionV3) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetAmount

`func (o *EnergyBillingDemandTransactionV3) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyBillingDemandTransactionV3) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyBillingDemandTransactionV3) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCalculationFactors

`func (o *EnergyBillingDemandTransactionV3) GetCalculationFactors() []EnergyBillingUsageTransactionV2CalculationFactors`

GetCalculationFactors returns the CalculationFactors field if non-nil, zero value otherwise.

### GetCalculationFactorsOk

`func (o *EnergyBillingDemandTransactionV3) GetCalculationFactorsOk() (*[]EnergyBillingUsageTransactionV2CalculationFactors, bool)`

GetCalculationFactorsOk returns a tuple with the CalculationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFactors

`func (o *EnergyBillingDemandTransactionV3) SetCalculationFactors(v []EnergyBillingUsageTransactionV2CalculationFactors)`

SetCalculationFactors sets CalculationFactors field to given value.

### HasCalculationFactors

`func (o *EnergyBillingDemandTransactionV3) HasCalculationFactors() bool`

HasCalculationFactors returns a boolean if a field has been set.

### GetAdjustments

`func (o *EnergyBillingDemandTransactionV3) GetAdjustments() []EnergyBillingUsageTransactionV2Adjustments`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *EnergyBillingDemandTransactionV3) GetAdjustmentsOk() (*[]EnergyBillingUsageTransactionV2Adjustments, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *EnergyBillingDemandTransactionV3) SetAdjustments(v []EnergyBillingUsageTransactionV2Adjustments)`

SetAdjustments sets Adjustments field to given value.

### HasAdjustments

`func (o *EnergyBillingDemandTransactionV3) HasAdjustments() bool`

HasAdjustments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


