# EnergyBillingTransactionV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**ExecutionDateTime** | **string** | The date and time that the transaction occurred. | 
**Gst** | Pointer to **string** | The GST incurred in the transaction. Should not be included for credits or payments. If absent then zero is assumed. | [optional] [default to "0.00"]
**TransactionUType** | **string** | Indicator of the type of transaction object present in this record. | 
**Usage** | Pointer to [**EnergyBillingUsageTransactionV2**](EnergyBillingUsageTransactionV2.md) | Represents a usage charge or generation credit. Mandatory if _transactionUType_ is equal to &#x60;usage&#x60;. | [optional] 
**Demand** | Pointer to [**EnergyBillingDemandTransactionV3**](EnergyBillingDemandTransactionV3.md) | Represents a demand charge or generation credit. Mandatory if _transactionUType_ is equal to &#x60;demand&#x60;. | [optional] 
**OnceOff** | Pointer to [**EnergyBillingOnceOffTransaction**](EnergyBillingOnceOffTransaction.md) | Represents a once off charge or credit. Mandatory if _transactionUType_ is equal to &#x60;onceOff&#x60;. | [optional] 
**OtherCharges** | Pointer to [**EnergyBillingOtherTransaction**](EnergyBillingOtherTransaction.md) | Represents charge other than usage and once off. Mandatory if _transactionUType_ is equal to &#x60;otherCharges&#x60;. | [optional] 
**Payment** | Pointer to [**EnergyBillingPaymentTransaction**](EnergyBillingPaymentTransaction.md) | Represents a payment to the account. Mandatory if _transactionUType_ is equal to &#x60;payment&#x60;. | [optional] 

## Methods

### NewEnergyBillingTransactionV3

`func NewEnergyBillingTransactionV3(accountId string, executionDateTime string, transactionUType string, ) *EnergyBillingTransactionV3`

NewEnergyBillingTransactionV3 instantiates a new EnergyBillingTransactionV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBillingTransactionV3WithDefaults

`func NewEnergyBillingTransactionV3WithDefaults() *EnergyBillingTransactionV3`

NewEnergyBillingTransactionV3WithDefaults instantiates a new EnergyBillingTransactionV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EnergyBillingTransactionV3) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EnergyBillingTransactionV3) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EnergyBillingTransactionV3) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetExecutionDateTime

`func (o *EnergyBillingTransactionV3) GetExecutionDateTime() string`

GetExecutionDateTime returns the ExecutionDateTime field if non-nil, zero value otherwise.

### GetExecutionDateTimeOk

`func (o *EnergyBillingTransactionV3) GetExecutionDateTimeOk() (*string, bool)`

GetExecutionDateTimeOk returns a tuple with the ExecutionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDateTime

`func (o *EnergyBillingTransactionV3) SetExecutionDateTime(v string)`

SetExecutionDateTime sets ExecutionDateTime field to given value.


### GetGst

`func (o *EnergyBillingTransactionV3) GetGst() string`

GetGst returns the Gst field if non-nil, zero value otherwise.

### GetGstOk

`func (o *EnergyBillingTransactionV3) GetGstOk() (*string, bool)`

GetGstOk returns a tuple with the Gst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGst

`func (o *EnergyBillingTransactionV3) SetGst(v string)`

SetGst sets Gst field to given value.

### HasGst

`func (o *EnergyBillingTransactionV3) HasGst() bool`

HasGst returns a boolean if a field has been set.

### GetTransactionUType

`func (o *EnergyBillingTransactionV3) GetTransactionUType() string`

GetTransactionUType returns the TransactionUType field if non-nil, zero value otherwise.

### GetTransactionUTypeOk

`func (o *EnergyBillingTransactionV3) GetTransactionUTypeOk() (*string, bool)`

GetTransactionUTypeOk returns a tuple with the TransactionUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUType

`func (o *EnergyBillingTransactionV3) SetTransactionUType(v string)`

SetTransactionUType sets TransactionUType field to given value.


### GetUsage

`func (o *EnergyBillingTransactionV3) GetUsage() EnergyBillingUsageTransactionV2`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EnergyBillingTransactionV3) GetUsageOk() (*EnergyBillingUsageTransactionV2, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EnergyBillingTransactionV3) SetUsage(v EnergyBillingUsageTransactionV2)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *EnergyBillingTransactionV3) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetDemand

`func (o *EnergyBillingTransactionV3) GetDemand() EnergyBillingDemandTransactionV3`

GetDemand returns the Demand field if non-nil, zero value otherwise.

### GetDemandOk

`func (o *EnergyBillingTransactionV3) GetDemandOk() (*EnergyBillingDemandTransactionV3, bool)`

GetDemandOk returns a tuple with the Demand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemand

`func (o *EnergyBillingTransactionV3) SetDemand(v EnergyBillingDemandTransactionV3)`

SetDemand sets Demand field to given value.

### HasDemand

`func (o *EnergyBillingTransactionV3) HasDemand() bool`

HasDemand returns a boolean if a field has been set.

### GetOnceOff

`func (o *EnergyBillingTransactionV3) GetOnceOff() EnergyBillingOnceOffTransaction`

GetOnceOff returns the OnceOff field if non-nil, zero value otherwise.

### GetOnceOffOk

`func (o *EnergyBillingTransactionV3) GetOnceOffOk() (*EnergyBillingOnceOffTransaction, bool)`

GetOnceOffOk returns a tuple with the OnceOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnceOff

`func (o *EnergyBillingTransactionV3) SetOnceOff(v EnergyBillingOnceOffTransaction)`

SetOnceOff sets OnceOff field to given value.

### HasOnceOff

`func (o *EnergyBillingTransactionV3) HasOnceOff() bool`

HasOnceOff returns a boolean if a field has been set.

### GetOtherCharges

`func (o *EnergyBillingTransactionV3) GetOtherCharges() EnergyBillingOtherTransaction`

GetOtherCharges returns the OtherCharges field if non-nil, zero value otherwise.

### GetOtherChargesOk

`func (o *EnergyBillingTransactionV3) GetOtherChargesOk() (*EnergyBillingOtherTransaction, bool)`

GetOtherChargesOk returns a tuple with the OtherCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCharges

`func (o *EnergyBillingTransactionV3) SetOtherCharges(v EnergyBillingOtherTransaction)`

SetOtherCharges sets OtherCharges field to given value.

### HasOtherCharges

`func (o *EnergyBillingTransactionV3) HasOtherCharges() bool`

HasOtherCharges returns a boolean if a field has been set.

### GetPayment

`func (o *EnergyBillingTransactionV3) GetPayment() EnergyBillingPaymentTransaction`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *EnergyBillingTransactionV3) GetPaymentOk() (*EnergyBillingPaymentTransaction, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *EnergyBillingTransactionV3) SetPayment(v EnergyBillingPaymentTransaction)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *EnergyBillingTransactionV3) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


