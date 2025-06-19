# EnergyAccountDetailV4AllOfPlanDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuelType** | **string** | The fuel types covered by the plan. | 
**IsContingentPlan** | Pointer to **bool** | Flag that indicates that the plan is contingent on the customer taking up an alternate fuel plan from the same retailer (for instance, if the _fuelType_ is &#x60;ELECTRICITY&#x60; then a &#x60;GAS&#x60; plan from the same retailer must be taken up). Has no meaning if the plan has a _fuelType_ of &#x60;DUAL&#x60;. If absent the value is assumed to be &#x60;false&#x60;. | [optional] [default to false]
**MeteringCharges** | Pointer to [**[]EnergyPlanDetailV3AllOfMeteringCharges**](EnergyPlanDetailV3AllOfMeteringCharges.md) | Charges for metering included in the plan. | [optional] 
**GasContract** | Pointer to [**EnergyPlanContractV3**](EnergyPlanContractV3.md) | The details of the terms for the supply of electricity under this plan. Is mandatory if _fuelType_ is set to &#x60;GAS&#x60; or &#x60;DUAL&#x60;. | [optional] 
**ElectricityContract** | Pointer to [**EnergyPlanContractV3**](EnergyPlanContractV3.md) | The details of the terms for the supply of electricity under this plan. Is mandatory if _fuelType_ is set to &#x60;ELECTRICITY&#x60; or &#x60;DUAL&#x60;. | [optional] 

## Methods

### NewEnergyAccountDetailV4AllOfPlanDetail

`func NewEnergyAccountDetailV4AllOfPlanDetail(fuelType string, ) *EnergyAccountDetailV4AllOfPlanDetail`

NewEnergyAccountDetailV4AllOfPlanDetail instantiates a new EnergyAccountDetailV4AllOfPlanDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyAccountDetailV4AllOfPlanDetailWithDefaults

`func NewEnergyAccountDetailV4AllOfPlanDetailWithDefaults() *EnergyAccountDetailV4AllOfPlanDetail`

NewEnergyAccountDetailV4AllOfPlanDetailWithDefaults instantiates a new EnergyAccountDetailV4AllOfPlanDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuelType

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetFuelType() string`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetFuelTypeOk() (*string, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *EnergyAccountDetailV4AllOfPlanDetail) SetFuelType(v string)`

SetFuelType sets FuelType field to given value.


### GetIsContingentPlan

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetIsContingentPlan() bool`

GetIsContingentPlan returns the IsContingentPlan field if non-nil, zero value otherwise.

### GetIsContingentPlanOk

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetIsContingentPlanOk() (*bool, bool)`

GetIsContingentPlanOk returns a tuple with the IsContingentPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContingentPlan

`func (o *EnergyAccountDetailV4AllOfPlanDetail) SetIsContingentPlan(v bool)`

SetIsContingentPlan sets IsContingentPlan field to given value.

### HasIsContingentPlan

`func (o *EnergyAccountDetailV4AllOfPlanDetail) HasIsContingentPlan() bool`

HasIsContingentPlan returns a boolean if a field has been set.

### GetMeteringCharges

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetMeteringCharges() []EnergyPlanDetailV3AllOfMeteringCharges`

GetMeteringCharges returns the MeteringCharges field if non-nil, zero value otherwise.

### GetMeteringChargesOk

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetMeteringChargesOk() (*[]EnergyPlanDetailV3AllOfMeteringCharges, bool)`

GetMeteringChargesOk returns a tuple with the MeteringCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringCharges

`func (o *EnergyAccountDetailV4AllOfPlanDetail) SetMeteringCharges(v []EnergyPlanDetailV3AllOfMeteringCharges)`

SetMeteringCharges sets MeteringCharges field to given value.

### HasMeteringCharges

`func (o *EnergyAccountDetailV4AllOfPlanDetail) HasMeteringCharges() bool`

HasMeteringCharges returns a boolean if a field has been set.

### GetGasContract

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetGasContract() EnergyPlanContractV3`

GetGasContract returns the GasContract field if non-nil, zero value otherwise.

### GetGasContractOk

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetGasContractOk() (*EnergyPlanContractV3, bool)`

GetGasContractOk returns a tuple with the GasContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasContract

`func (o *EnergyAccountDetailV4AllOfPlanDetail) SetGasContract(v EnergyPlanContractV3)`

SetGasContract sets GasContract field to given value.

### HasGasContract

`func (o *EnergyAccountDetailV4AllOfPlanDetail) HasGasContract() bool`

HasGasContract returns a boolean if a field has been set.

### GetElectricityContract

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetElectricityContract() EnergyPlanContractV3`

GetElectricityContract returns the ElectricityContract field if non-nil, zero value otherwise.

### GetElectricityContractOk

`func (o *EnergyAccountDetailV4AllOfPlanDetail) GetElectricityContractOk() (*EnergyPlanContractV3, bool)`

GetElectricityContractOk returns a tuple with the ElectricityContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityContract

`func (o *EnergyAccountDetailV4AllOfPlanDetail) SetElectricityContract(v EnergyPlanContractV3)`

SetElectricityContract sets ElectricityContract field to given value.

### HasElectricityContract

`func (o *EnergyAccountDetailV4AllOfPlanDetail) HasElectricityContract() bool`

HasElectricityContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


