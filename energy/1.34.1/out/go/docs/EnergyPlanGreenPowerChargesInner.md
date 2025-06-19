# EnergyPlanGreenPowerChargesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the charge. | 
**Description** | Pointer to **string** | The description of the charge. | [optional] 
**Scheme** | **string** | The applicable green power scheme. | 
**Type** | **string** | The type of charge. | 
**Tiers** | [**[]EnergyPlanGreenPowerChargesInnerTiersInner**](EnergyPlanGreenPowerChargesInnerTiersInner.md) | Array of charge tiers based on the percentage of green power used for the period implied by the type. Array is in order of increasing percentage of green power. | 

## Methods

### NewEnergyPlanGreenPowerChargesInner

`func NewEnergyPlanGreenPowerChargesInner(displayName string, scheme string, type_ string, tiers []EnergyPlanGreenPowerChargesInnerTiersInner, ) *EnergyPlanGreenPowerChargesInner`

NewEnergyPlanGreenPowerChargesInner instantiates a new EnergyPlanGreenPowerChargesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanGreenPowerChargesInnerWithDefaults

`func NewEnergyPlanGreenPowerChargesInnerWithDefaults() *EnergyPlanGreenPowerChargesInner`

NewEnergyPlanGreenPowerChargesInnerWithDefaults instantiates a new EnergyPlanGreenPowerChargesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanGreenPowerChargesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanGreenPowerChargesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanGreenPowerChargesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanGreenPowerChargesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanGreenPowerChargesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanGreenPowerChargesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanGreenPowerChargesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheme

`func (o *EnergyPlanGreenPowerChargesInner) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *EnergyPlanGreenPowerChargesInner) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *EnergyPlanGreenPowerChargesInner) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetType

`func (o *EnergyPlanGreenPowerChargesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanGreenPowerChargesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanGreenPowerChargesInner) SetType(v string)`

SetType sets Type field to given value.


### GetTiers

`func (o *EnergyPlanGreenPowerChargesInner) GetTiers() []EnergyPlanGreenPowerChargesInnerTiersInner`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *EnergyPlanGreenPowerChargesInner) GetTiersOk() (*[]EnergyPlanGreenPowerChargesInnerTiersInner, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *EnergyPlanGreenPowerChargesInner) SetTiers(v []EnergyPlanGreenPowerChargesInnerTiersInner)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


