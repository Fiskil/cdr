# EnergyBalanceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** | The current balance of the account. A positive value indicates that amount is owing to be paid. A negative value indicates that the account is in credit. | 

## Methods

### NewEnergyBalanceResponseData

`func NewEnergyBalanceResponseData(balance string, ) *EnergyBalanceResponseData`

NewEnergyBalanceResponseData instantiates a new EnergyBalanceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyBalanceResponseDataWithDefaults

`func NewEnergyBalanceResponseDataWithDefaults() *EnergyBalanceResponseData`

NewEnergyBalanceResponseDataWithDefaults instantiates a new EnergyBalanceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *EnergyBalanceResponseData) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EnergyBalanceResponseData) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EnergyBalanceResponseData) SetBalance(v string)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


