# EnergyPaymentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Optional payment amount indicating that a constant payment amount is scheduled to be paid (used in bill smoothing scenarios). | [optional] 
**PaymentScheduleUType** | **string** | The type of object present in this response. | 
**CardDebit** | Pointer to [**EnergyPaymentScheduleCardDebit**](EnergyPaymentScheduleCardDebit.md) |  | [optional] 
**DirectDebit** | Pointer to [**EnergyPaymentScheduleDirectDebit**](EnergyPaymentScheduleDirectDebit.md) |  | [optional] 
**DigitalWallet** | Pointer to [**EnergyPaymentScheduleDigitalWallet**](EnergyPaymentScheduleDigitalWallet.md) |  | [optional] 
**ManualPayment** | Pointer to [**EnergyPaymentScheduleManualPayment**](EnergyPaymentScheduleManualPayment.md) |  | [optional] 

## Methods

### NewEnergyPaymentSchedule

`func NewEnergyPaymentSchedule(paymentScheduleUType string, ) *EnergyPaymentSchedule`

NewEnergyPaymentSchedule instantiates a new EnergyPaymentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPaymentScheduleWithDefaults

`func NewEnergyPaymentScheduleWithDefaults() *EnergyPaymentSchedule`

NewEnergyPaymentScheduleWithDefaults instantiates a new EnergyPaymentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EnergyPaymentSchedule) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyPaymentSchedule) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyPaymentSchedule) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EnergyPaymentSchedule) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPaymentScheduleUType

`func (o *EnergyPaymentSchedule) GetPaymentScheduleUType() string`

GetPaymentScheduleUType returns the PaymentScheduleUType field if non-nil, zero value otherwise.

### GetPaymentScheduleUTypeOk

`func (o *EnergyPaymentSchedule) GetPaymentScheduleUTypeOk() (*string, bool)`

GetPaymentScheduleUTypeOk returns a tuple with the PaymentScheduleUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheduleUType

`func (o *EnergyPaymentSchedule) SetPaymentScheduleUType(v string)`

SetPaymentScheduleUType sets PaymentScheduleUType field to given value.


### GetCardDebit

`func (o *EnergyPaymentSchedule) GetCardDebit() EnergyPaymentScheduleCardDebit`

GetCardDebit returns the CardDebit field if non-nil, zero value otherwise.

### GetCardDebitOk

`func (o *EnergyPaymentSchedule) GetCardDebitOk() (*EnergyPaymentScheduleCardDebit, bool)`

GetCardDebitOk returns a tuple with the CardDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDebit

`func (o *EnergyPaymentSchedule) SetCardDebit(v EnergyPaymentScheduleCardDebit)`

SetCardDebit sets CardDebit field to given value.

### HasCardDebit

`func (o *EnergyPaymentSchedule) HasCardDebit() bool`

HasCardDebit returns a boolean if a field has been set.

### GetDirectDebit

`func (o *EnergyPaymentSchedule) GetDirectDebit() EnergyPaymentScheduleDirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *EnergyPaymentSchedule) GetDirectDebitOk() (*EnergyPaymentScheduleDirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *EnergyPaymentSchedule) SetDirectDebit(v EnergyPaymentScheduleDirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *EnergyPaymentSchedule) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### GetDigitalWallet

`func (o *EnergyPaymentSchedule) GetDigitalWallet() EnergyPaymentScheduleDigitalWallet`

GetDigitalWallet returns the DigitalWallet field if non-nil, zero value otherwise.

### GetDigitalWalletOk

`func (o *EnergyPaymentSchedule) GetDigitalWalletOk() (*EnergyPaymentScheduleDigitalWallet, bool)`

GetDigitalWalletOk returns a tuple with the DigitalWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWallet

`func (o *EnergyPaymentSchedule) SetDigitalWallet(v EnergyPaymentScheduleDigitalWallet)`

SetDigitalWallet sets DigitalWallet field to given value.

### HasDigitalWallet

`func (o *EnergyPaymentSchedule) HasDigitalWallet() bool`

HasDigitalWallet returns a boolean if a field has been set.

### GetManualPayment

`func (o *EnergyPaymentSchedule) GetManualPayment() EnergyPaymentScheduleManualPayment`

GetManualPayment returns the ManualPayment field if non-nil, zero value otherwise.

### GetManualPaymentOk

`func (o *EnergyPaymentSchedule) GetManualPaymentOk() (*EnergyPaymentScheduleManualPayment, bool)`

GetManualPaymentOk returns a tuple with the ManualPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualPayment

`func (o *EnergyPaymentSchedule) SetManualPayment(v EnergyPaymentScheduleManualPayment)`

SetManualPayment sets ManualPayment field to given value.

### HasManualPayment

`func (o *EnergyPaymentSchedule) HasManualPayment() bool`

HasManualPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


