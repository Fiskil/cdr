# EnergyBillingListResponseV3Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]EnergyBillingTransactionV3**](EnergyBillingTransactionV3.md) | Array of transactions sorted by date and time in descending order. | 

## Methods

### NewEnergyBillingListResponseV3Data

`func NewEnergyBillingListResponseV3Data(transactions []EnergyBillingTransactionV3, ) *EnergyBillingListResponseV3Data`

NewEnergyBillingListResponseV3Data instantiates a new EnergyBillingListResponseV3Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBillingListResponseV3DataWithDefaults

`func NewEnergyBillingListResponseV3DataWithDefaults() *EnergyBillingListResponseV3Data`

NewEnergyBillingListResponseV3DataWithDefaults instantiates a new EnergyBillingListResponseV3Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *EnergyBillingListResponseV3Data) GetTransactions() []EnergyBillingTransactionV3`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *EnergyBillingListResponseV3Data) GetTransactionsOk() (*[]EnergyBillingTransactionV3, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *EnergyBillingListResponseV3Data) SetTransactions(v []EnergyBillingTransactionV3)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


