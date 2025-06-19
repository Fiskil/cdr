# EnergyInvoicePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | The start date of the period covered by this invoice. | 
**EndDate** | **string** | The end date of the period covered by this invoice. | 

## Methods

### NewEnergyInvoicePeriod

`func NewEnergyInvoicePeriod(startDate string, endDate string, ) *EnergyInvoicePeriod`

NewEnergyInvoicePeriod instantiates a new EnergyInvoicePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyInvoicePeriodWithDefaults

`func NewEnergyInvoicePeriodWithDefaults() *EnergyInvoicePeriod`

NewEnergyInvoicePeriodWithDefaults instantiates a new EnergyInvoicePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *EnergyInvoicePeriod) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyInvoicePeriod) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyInvoicePeriod) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *EnergyInvoicePeriod) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyInvoicePeriod) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyInvoicePeriod) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


