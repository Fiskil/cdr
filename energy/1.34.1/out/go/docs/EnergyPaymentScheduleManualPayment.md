# EnergyPaymentScheduleManualPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillFrequency** | **string** | The frequency with which a bill will be issued. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | 

## Methods

### NewEnergyPaymentScheduleManualPayment

`func NewEnergyPaymentScheduleManualPayment(billFrequency string, ) *EnergyPaymentScheduleManualPayment`

NewEnergyPaymentScheduleManualPayment instantiates a new EnergyPaymentScheduleManualPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPaymentScheduleManualPaymentWithDefaults

`func NewEnergyPaymentScheduleManualPaymentWithDefaults() *EnergyPaymentScheduleManualPayment`

NewEnergyPaymentScheduleManualPaymentWithDefaults instantiates a new EnergyPaymentScheduleManualPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillFrequency

`func (o *EnergyPaymentScheduleManualPayment) GetBillFrequency() string`

GetBillFrequency returns the BillFrequency field if non-nil, zero value otherwise.

### GetBillFrequencyOk

`func (o *EnergyPaymentScheduleManualPayment) GetBillFrequencyOk() (*string, bool)`

GetBillFrequencyOk returns a tuple with the BillFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillFrequency

`func (o *EnergyPaymentScheduleManualPayment) SetBillFrequency(v string)`

SetBillFrequency sets BillFrequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


