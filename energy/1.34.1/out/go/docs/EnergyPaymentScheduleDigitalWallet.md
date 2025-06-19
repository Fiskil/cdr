# EnergyPaymentScheduleDigitalWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name of the wallet as given by the customer, else a default value defined by the data holder. | 
**Identifier** | **string** | The identifier of the digital wallet (dependent on type). | 
**Type** | **string** | The type of the digital wallet identifier. | 
**Provider** | **string** | The provider of the digital wallet. | 
**PaymentFrequency** | **string** | The frequency that payments will occur. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | 
**CalculationType** | **string** | The mechanism by which the payment amount is calculated. Explanation of values are as follows:&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&#x60;STATIC&#x60;: Indicates a consistent, static amount, per payment&lt;/li&gt;&lt;li&gt;&#x60;BALANCE&#x60;: Indicates that the outstanding balance for the account is paid per period&lt;/li&gt;&lt;li&gt;&#x60;CALCULATED&#x60;: Indicates that the payment amount is variable and calculated using a pre-defined algorithm.&lt;/li&gt;&lt;/ul&gt; | 

## Methods

### NewEnergyPaymentScheduleDigitalWallet

`func NewEnergyPaymentScheduleDigitalWallet(name string, identifier string, type_ string, provider string, paymentFrequency string, calculationType string, ) *EnergyPaymentScheduleDigitalWallet`

NewEnergyPaymentScheduleDigitalWallet instantiates a new EnergyPaymentScheduleDigitalWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPaymentScheduleDigitalWalletWithDefaults

`func NewEnergyPaymentScheduleDigitalWalletWithDefaults() *EnergyPaymentScheduleDigitalWallet`

NewEnergyPaymentScheduleDigitalWalletWithDefaults instantiates a new EnergyPaymentScheduleDigitalWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnergyPaymentScheduleDigitalWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnergyPaymentScheduleDigitalWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnergyPaymentScheduleDigitalWallet) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *EnergyPaymentScheduleDigitalWallet) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EnergyPaymentScheduleDigitalWallet) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EnergyPaymentScheduleDigitalWallet) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetType

`func (o *EnergyPaymentScheduleDigitalWallet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPaymentScheduleDigitalWallet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPaymentScheduleDigitalWallet) SetType(v string)`

SetType sets Type field to given value.


### GetProvider

`func (o *EnergyPaymentScheduleDigitalWallet) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EnergyPaymentScheduleDigitalWallet) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EnergyPaymentScheduleDigitalWallet) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPaymentFrequency

`func (o *EnergyPaymentScheduleDigitalWallet) GetPaymentFrequency() string`

GetPaymentFrequency returns the PaymentFrequency field if non-nil, zero value otherwise.

### GetPaymentFrequencyOk

`func (o *EnergyPaymentScheduleDigitalWallet) GetPaymentFrequencyOk() (*string, bool)`

GetPaymentFrequencyOk returns a tuple with the PaymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrequency

`func (o *EnergyPaymentScheduleDigitalWallet) SetPaymentFrequency(v string)`

SetPaymentFrequency sets PaymentFrequency field to given value.


### GetCalculationType

`func (o *EnergyPaymentScheduleDigitalWallet) GetCalculationType() string`

GetCalculationType returns the CalculationType field if non-nil, zero value otherwise.

### GetCalculationTypeOk

`func (o *EnergyPaymentScheduleDigitalWallet) GetCalculationTypeOk() (*string, bool)`

GetCalculationTypeOk returns a tuple with the CalculationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationType

`func (o *EnergyPaymentScheduleDigitalWallet) SetCalculationType(v string)`

SetCalculationType sets CalculationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


