# EnergyBalanceListResponseDataBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**Balance** | **string** | The current balance of the account. A positive value indicates that amount is owing to be paid. A negative value indicates that the account is in credit. | 

## Methods

### NewEnergyBalanceListResponseDataBalances

`func NewEnergyBalanceListResponseDataBalances(accountId string, balance string, ) *EnergyBalanceListResponseDataBalances`

NewEnergyBalanceListResponseDataBalances instantiates a new EnergyBalanceListResponseDataBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBalanceListResponseDataBalancesWithDefaults

`func NewEnergyBalanceListResponseDataBalancesWithDefaults() *EnergyBalanceListResponseDataBalances`

NewEnergyBalanceListResponseDataBalancesWithDefaults instantiates a new EnergyBalanceListResponseDataBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EnergyBalanceListResponseDataBalances) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EnergyBalanceListResponseDataBalances) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EnergyBalanceListResponseDataBalances) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBalance

`func (o *EnergyBalanceListResponseDataBalances) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EnergyBalanceListResponseDataBalances) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EnergyBalanceListResponseDataBalances) SetBalance(v string)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


