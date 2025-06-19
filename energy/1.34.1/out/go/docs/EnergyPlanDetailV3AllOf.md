# EnergyPlanDetailV3AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeteringCharges** | Pointer to [**[]EnergyPlanDetailV3AllOfMeteringCharges**](EnergyPlanDetailV3AllOfMeteringCharges.md) | Charges for metering included in the plan. | [optional] 
**GasContract** | Pointer to [**EnergyPlanContractFullV3**](EnergyPlanContractFullV3.md) | The details of the terms for the supply of electricity under this plan. Is mandatory if _fuelType_ is set to &#x60;GAS&#x60; or &#x60;DUAL&#x60;. | [optional] 
**ElectricityContract** | Pointer to [**EnergyPlanContractFullV3**](EnergyPlanContractFullV3.md) | The details of the terms for the supply of electricity under this plan. Is mandatory if _fuelType_ is set to &#x60;ELECTRICITY&#x60; or &#x60;DUAL&#x60;. | [optional] 

## Methods

### NewEnergyPlanDetailV3AllOf

`func NewEnergyPlanDetailV3AllOf() *EnergyPlanDetailV3AllOf`

NewEnergyPlanDetailV3AllOf instantiates a new EnergyPlanDetailV3AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanDetailV3AllOfWithDefaults

`func NewEnergyPlanDetailV3AllOfWithDefaults() *EnergyPlanDetailV3AllOf`

NewEnergyPlanDetailV3AllOfWithDefaults instantiates a new EnergyPlanDetailV3AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeteringCharges

`func (o *EnergyPlanDetailV3AllOf) GetMeteringCharges() []EnergyPlanDetailV3AllOfMeteringCharges`

GetMeteringCharges returns the MeteringCharges field if non-nil, zero value otherwise.

### GetMeteringChargesOk

`func (o *EnergyPlanDetailV3AllOf) GetMeteringChargesOk() (*[]EnergyPlanDetailV3AllOfMeteringCharges, bool)`

GetMeteringChargesOk returns a tuple with the MeteringCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringCharges

`func (o *EnergyPlanDetailV3AllOf) SetMeteringCharges(v []EnergyPlanDetailV3AllOfMeteringCharges)`

SetMeteringCharges sets MeteringCharges field to given value.

### HasMeteringCharges

`func (o *EnergyPlanDetailV3AllOf) HasMeteringCharges() bool`

HasMeteringCharges returns a boolean if a field has been set.

### GetGasContract

`func (o *EnergyPlanDetailV3AllOf) GetGasContract() EnergyPlanContractFullV3`

GetGasContract returns the GasContract field if non-nil, zero value otherwise.

### GetGasContractOk

`func (o *EnergyPlanDetailV3AllOf) GetGasContractOk() (*EnergyPlanContractFullV3, bool)`

GetGasContractOk returns a tuple with the GasContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasContract

`func (o *EnergyPlanDetailV3AllOf) SetGasContract(v EnergyPlanContractFullV3)`

SetGasContract sets GasContract field to given value.

### HasGasContract

`func (o *EnergyPlanDetailV3AllOf) HasGasContract() bool`

HasGasContract returns a boolean if a field has been set.

### GetElectricityContract

`func (o *EnergyPlanDetailV3AllOf) GetElectricityContract() EnergyPlanContractFullV3`

GetElectricityContract returns the ElectricityContract field if non-nil, zero value otherwise.

### GetElectricityContractOk

`func (o *EnergyPlanDetailV3AllOf) GetElectricityContractOk() (*EnergyPlanContractFullV3, bool)`

GetElectricityContractOk returns a tuple with the ElectricityContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityContract

`func (o *EnergyPlanDetailV3AllOf) SetElectricityContract(v EnergyPlanContractFullV3)`

SetElectricityContract sets ElectricityContract field to given value.

### HasElectricityContract

`func (o *EnergyPlanDetailV3AllOf) HasElectricityContract() bool`

HasElectricityContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


