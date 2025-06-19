# EnergyPaymentScheduleDirectDebit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTokenised** | Pointer to **bool** | Flag indicating that the account details are tokenised, or held in a closed system, and is not accessible through any other channels. &#x60;false&#x60; if absent. | [optional] [default to false]
**Bsb** | Pointer to **string** | The unmasked BSB for the account to be debited. Is expected to be formatted as digits only with leading zeros included and no punctuation or spaces. Is required if _isTokenised_ is absent or &#x60;false&#x60;. | [optional] 
**AccountNumber** | Pointer to **string** | The unmasked account number for the account to be debited. Is expected to be formatted as digits only with leading zeros included and no punctuation or spaces. Is required if _isTokenised_ is absent or &#x60;false&#x60;. | [optional] 
**PaymentFrequency** | **string** | The frequency that payments will occur. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | 
**CalculationType** | **string** | The mechanism by which the payment amount is calculated. Explanation of values are as follows:&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&#x60;STATIC&#x60;: Indicates a consistent, static amount, per payment&lt;/li&gt;&lt;li&gt;&#x60;BALANCE&#x60;: Indicates that the outstanding balance for the account is paid per period&lt;/li&gt;&lt;li&gt;&#x60;CALCULATED&#x60;: Indicates that the payment amount is variable and calculated using a pre-defined algorithm.&lt;/li&gt;&lt;/ul&gt; | 

## Methods

### NewEnergyPaymentScheduleDirectDebit

`func NewEnergyPaymentScheduleDirectDebit(paymentFrequency string, calculationType string, ) *EnergyPaymentScheduleDirectDebit`

NewEnergyPaymentScheduleDirectDebit instantiates a new EnergyPaymentScheduleDirectDebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPaymentScheduleDirectDebitWithDefaults

`func NewEnergyPaymentScheduleDirectDebitWithDefaults() *EnergyPaymentScheduleDirectDebit`

NewEnergyPaymentScheduleDirectDebitWithDefaults instantiates a new EnergyPaymentScheduleDirectDebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTokenised

`func (o *EnergyPaymentScheduleDirectDebit) GetIsTokenised() bool`

GetIsTokenised returns the IsTokenised field if non-nil, zero value otherwise.

### GetIsTokenisedOk

`func (o *EnergyPaymentScheduleDirectDebit) GetIsTokenisedOk() (*bool, bool)`

GetIsTokenisedOk returns a tuple with the IsTokenised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenised

`func (o *EnergyPaymentScheduleDirectDebit) SetIsTokenised(v bool)`

SetIsTokenised sets IsTokenised field to given value.

### HasIsTokenised

`func (o *EnergyPaymentScheduleDirectDebit) HasIsTokenised() bool`

HasIsTokenised returns a boolean if a field has been set.

### GetBsb

`func (o *EnergyPaymentScheduleDirectDebit) GetBsb() string`

GetBsb returns the Bsb field if non-nil, zero value otherwise.

### GetBsbOk

`func (o *EnergyPaymentScheduleDirectDebit) GetBsbOk() (*string, bool)`

GetBsbOk returns a tuple with the Bsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsb

`func (o *EnergyPaymentScheduleDirectDebit) SetBsb(v string)`

SetBsb sets Bsb field to given value.

### HasBsb

`func (o *EnergyPaymentScheduleDirectDebit) HasBsb() bool`

HasBsb returns a boolean if a field has been set.

### GetAccountNumber

`func (o *EnergyPaymentScheduleDirectDebit) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *EnergyPaymentScheduleDirectDebit) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *EnergyPaymentScheduleDirectDebit) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *EnergyPaymentScheduleDirectDebit) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetPaymentFrequency

`func (o *EnergyPaymentScheduleDirectDebit) GetPaymentFrequency() string`

GetPaymentFrequency returns the PaymentFrequency field if non-nil, zero value otherwise.

### GetPaymentFrequencyOk

`func (o *EnergyPaymentScheduleDirectDebit) GetPaymentFrequencyOk() (*string, bool)`

GetPaymentFrequencyOk returns a tuple with the PaymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrequency

`func (o *EnergyPaymentScheduleDirectDebit) SetPaymentFrequency(v string)`

SetPaymentFrequency sets PaymentFrequency field to given value.


### GetCalculationType

`func (o *EnergyPaymentScheduleDirectDebit) GetCalculationType() string`

GetCalculationType returns the CalculationType field if non-nil, zero value otherwise.

### GetCalculationTypeOk

`func (o *EnergyPaymentScheduleDirectDebit) GetCalculationTypeOk() (*string, bool)`

GetCalculationTypeOk returns a tuple with the CalculationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationType

`func (o *EnergyPaymentScheduleDirectDebit) SetCalculationType(v string)`

SetCalculationType sets CalculationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


