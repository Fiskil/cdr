# EnergyPlanContractFullV3AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermType** | Pointer to **string** | The term for the contract. If absent assumes no specified term. | [optional] 
**BenefitPeriod** | Pointer to **string** | Description of the benefit period. Should only be present if _termType_ has the value &#x60;ONGOING&#x60;. | [optional] 
**Terms** | Pointer to **string** | Free text description of the terms for the contract. | [optional] 
**MeterTypes** | Pointer to **[]string** | An array of the meter types that this contract is available for. | [optional] 
**CoolingOffDays** | Pointer to **int32** | Number of days in the cooling off period for the contract. Mandatory for plans with type of &#x60;MARKET&#x60;. | [optional] 
**BillFrequency** | **[]string** | An array of the available billing schedules for this contract. Formatted according to [ISO 8601 Durations](https://en.wikipedia.org/wiki/ISO_8601#Durations) (excludes recurrence syntax). | 

## Methods

### NewEnergyPlanContractFullV3AllOf

`func NewEnergyPlanContractFullV3AllOf(billFrequency []string, ) *EnergyPlanContractFullV3AllOf`

NewEnergyPlanContractFullV3AllOf instantiates a new EnergyPlanContractFullV3AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanContractFullV3AllOfWithDefaults

`func NewEnergyPlanContractFullV3AllOfWithDefaults() *EnergyPlanContractFullV3AllOf`

NewEnergyPlanContractFullV3AllOfWithDefaults instantiates a new EnergyPlanContractFullV3AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermType

`func (o *EnergyPlanContractFullV3AllOf) GetTermType() string`

GetTermType returns the TermType field if non-nil, zero value otherwise.

### GetTermTypeOk

`func (o *EnergyPlanContractFullV3AllOf) GetTermTypeOk() (*string, bool)`

GetTermTypeOk returns a tuple with the TermType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermType

`func (o *EnergyPlanContractFullV3AllOf) SetTermType(v string)`

SetTermType sets TermType field to given value.

### HasTermType

`func (o *EnergyPlanContractFullV3AllOf) HasTermType() bool`

HasTermType returns a boolean if a field has been set.

### GetBenefitPeriod

`func (o *EnergyPlanContractFullV3AllOf) GetBenefitPeriod() string`

GetBenefitPeriod returns the BenefitPeriod field if non-nil, zero value otherwise.

### GetBenefitPeriodOk

`func (o *EnergyPlanContractFullV3AllOf) GetBenefitPeriodOk() (*string, bool)`

GetBenefitPeriodOk returns a tuple with the BenefitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitPeriod

`func (o *EnergyPlanContractFullV3AllOf) SetBenefitPeriod(v string)`

SetBenefitPeriod sets BenefitPeriod field to given value.

### HasBenefitPeriod

`func (o *EnergyPlanContractFullV3AllOf) HasBenefitPeriod() bool`

HasBenefitPeriod returns a boolean if a field has been set.

### GetTerms

`func (o *EnergyPlanContractFullV3AllOf) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *EnergyPlanContractFullV3AllOf) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *EnergyPlanContractFullV3AllOf) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *EnergyPlanContractFullV3AllOf) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetMeterTypes

`func (o *EnergyPlanContractFullV3AllOf) GetMeterTypes() []string`

GetMeterTypes returns the MeterTypes field if non-nil, zero value otherwise.

### GetMeterTypesOk

`func (o *EnergyPlanContractFullV3AllOf) GetMeterTypesOk() (*[]string, bool)`

GetMeterTypesOk returns a tuple with the MeterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterTypes

`func (o *EnergyPlanContractFullV3AllOf) SetMeterTypes(v []string)`

SetMeterTypes sets MeterTypes field to given value.

### HasMeterTypes

`func (o *EnergyPlanContractFullV3AllOf) HasMeterTypes() bool`

HasMeterTypes returns a boolean if a field has been set.

### GetCoolingOffDays

`func (o *EnergyPlanContractFullV3AllOf) GetCoolingOffDays() int32`

GetCoolingOffDays returns the CoolingOffDays field if non-nil, zero value otherwise.

### GetCoolingOffDaysOk

`func (o *EnergyPlanContractFullV3AllOf) GetCoolingOffDaysOk() (*int32, bool)`

GetCoolingOffDaysOk returns a tuple with the CoolingOffDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingOffDays

`func (o *EnergyPlanContractFullV3AllOf) SetCoolingOffDays(v int32)`

SetCoolingOffDays sets CoolingOffDays field to given value.

### HasCoolingOffDays

`func (o *EnergyPlanContractFullV3AllOf) HasCoolingOffDays() bool`

HasCoolingOffDays returns a boolean if a field has been set.

### GetBillFrequency

`func (o *EnergyPlanContractFullV3AllOf) GetBillFrequency() []string`

GetBillFrequency returns the BillFrequency field if non-nil, zero value otherwise.

### GetBillFrequencyOk

`func (o *EnergyPlanContractFullV3AllOf) GetBillFrequencyOk() (*[]string, bool)`

GetBillFrequencyOk returns a tuple with the BillFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillFrequency

`func (o *EnergyPlanContractFullV3AllOf) SetBillFrequency(v []string)`

SetBillFrequency sets BillFrequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


