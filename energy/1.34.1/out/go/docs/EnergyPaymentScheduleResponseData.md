# EnergyPaymentScheduleResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSchedules** | [**[]EnergyPaymentSchedule**](EnergyPaymentSchedule.md) | Array may be empty if no payment schedules exist. | 

## Methods

### NewEnergyPaymentScheduleResponseData

`func NewEnergyPaymentScheduleResponseData(paymentSchedules []EnergyPaymentSchedule, ) *EnergyPaymentScheduleResponseData`

NewEnergyPaymentScheduleResponseData instantiates a new EnergyPaymentScheduleResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPaymentScheduleResponseDataWithDefaults

`func NewEnergyPaymentScheduleResponseDataWithDefaults() *EnergyPaymentScheduleResponseData`

NewEnergyPaymentScheduleResponseDataWithDefaults instantiates a new EnergyPaymentScheduleResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSchedules

`func (o *EnergyPaymentScheduleResponseData) GetPaymentSchedules() []EnergyPaymentSchedule`

GetPaymentSchedules returns the PaymentSchedules field if non-nil, zero value otherwise.

### GetPaymentSchedulesOk

`func (o *EnergyPaymentScheduleResponseData) GetPaymentSchedulesOk() (*[]EnergyPaymentSchedule, bool)`

GetPaymentSchedulesOk returns a tuple with the PaymentSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedules

`func (o *EnergyPaymentScheduleResponseData) SetPaymentSchedules(v []EnergyPaymentSchedule)`

SetPaymentSchedules sets PaymentSchedules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


