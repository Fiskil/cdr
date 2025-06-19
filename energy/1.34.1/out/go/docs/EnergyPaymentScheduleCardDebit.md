# EnergyPaymentScheduleCardDebit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardScheme** | **string** | The type of credit card held on file. | 
**PaymentFrequency** | **string** | The frequency that payments will occur. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | 
**CalculationType** | **string** | The mechanism by which the payment amount is calculated. Explanation of values are as follows:&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&#x60;STATIC&#x60;: Indicates a consistent, static amount, per payment&lt;/li&gt;&lt;li&gt;&#x60;BALANCE&#x60;: Indicates that the outstanding balance for the account is paid per period&lt;/li&gt;&lt;li&gt;&#x60;CALCULATED&#x60;: Indicates that the payment amount is variable and calculated using a pre-defined algorithm.&lt;/li&gt;&lt;/ul&gt; | 

## Methods

### NewEnergyPaymentScheduleCardDebit

`func NewEnergyPaymentScheduleCardDebit(cardScheme string, paymentFrequency string, calculationType string, ) *EnergyPaymentScheduleCardDebit`

NewEnergyPaymentScheduleCardDebit instantiates a new EnergyPaymentScheduleCardDebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPaymentScheduleCardDebitWithDefaults

`func NewEnergyPaymentScheduleCardDebitWithDefaults() *EnergyPaymentScheduleCardDebit`

NewEnergyPaymentScheduleCardDebitWithDefaults instantiates a new EnergyPaymentScheduleCardDebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardScheme

`func (o *EnergyPaymentScheduleCardDebit) GetCardScheme() string`

GetCardScheme returns the CardScheme field if non-nil, zero value otherwise.

### GetCardSchemeOk

`func (o *EnergyPaymentScheduleCardDebit) GetCardSchemeOk() (*string, bool)`

GetCardSchemeOk returns a tuple with the CardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardScheme

`func (o *EnergyPaymentScheduleCardDebit) SetCardScheme(v string)`

SetCardScheme sets CardScheme field to given value.


### GetPaymentFrequency

`func (o *EnergyPaymentScheduleCardDebit) GetPaymentFrequency() string`

GetPaymentFrequency returns the PaymentFrequency field if non-nil, zero value otherwise.

### GetPaymentFrequencyOk

`func (o *EnergyPaymentScheduleCardDebit) GetPaymentFrequencyOk() (*string, bool)`

GetPaymentFrequencyOk returns a tuple with the PaymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrequency

`func (o *EnergyPaymentScheduleCardDebit) SetPaymentFrequency(v string)`

SetPaymentFrequency sets PaymentFrequency field to given value.


### GetCalculationType

`func (o *EnergyPaymentScheduleCardDebit) GetCalculationType() string`

GetCalculationType returns the CalculationType field if non-nil, zero value otherwise.

### GetCalculationTypeOk

`func (o *EnergyPaymentScheduleCardDebit) GetCalculationTypeOk() (*string, bool)`

GetCalculationTypeOk returns a tuple with the CalculationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationType

`func (o *EnergyPaymentScheduleCardDebit) SetCalculationType(v string)`

SetCalculationType sets CalculationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


