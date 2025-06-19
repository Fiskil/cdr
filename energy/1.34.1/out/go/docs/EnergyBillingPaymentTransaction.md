# EnergyBillingPaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The amount paid. | 
**Method** | **string** | The method of payment. | 

## Methods

### NewEnergyBillingPaymentTransaction

`func NewEnergyBillingPaymentTransaction(amount string, method string, ) *EnergyBillingPaymentTransaction`

NewEnergyBillingPaymentTransaction instantiates a new EnergyBillingPaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBillingPaymentTransactionWithDefaults

`func NewEnergyBillingPaymentTransactionWithDefaults() *EnergyBillingPaymentTransaction`

NewEnergyBillingPaymentTransactionWithDefaults instantiates a new EnergyBillingPaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EnergyBillingPaymentTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnergyBillingPaymentTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnergyBillingPaymentTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMethod

`func (o *EnergyBillingPaymentTransaction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EnergyBillingPaymentTransaction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EnergyBillingPaymentTransaction) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


