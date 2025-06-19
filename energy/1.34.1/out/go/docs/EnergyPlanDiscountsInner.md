# EnergyPlanDiscountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the discount. | 
**Description** | Pointer to **string** | The description of the discount. | [optional] 
**Type** | **string** | The type of the discount. | 
**Category** | Pointer to **string** | The type of the discount. Mandatory if the discount _type_ is &#x60;CONDITIONAL&#x60;. | [optional] 
**EndDate** | Pointer to **string** | Optional end date for the discount after which the discount is no longer available. | [optional] 
**MethodUType** | **string** | The method of calculation of the discount. | 
**PercentOfBill** | Pointer to [**EnergyPlanDiscountsInnerPercentOfBill**](EnergyPlanDiscountsInnerPercentOfBill.md) |  | [optional] 
**PercentOfUse** | Pointer to [**EnergyPlanDiscountsInnerPercentOfUse**](EnergyPlanDiscountsInnerPercentOfUse.md) |  | [optional] 
**FixedAmount** | Pointer to [**EnergyPlanDiscountsInnerFixedAmount**](EnergyPlanDiscountsInnerFixedAmount.md) |  | [optional] 
**PercentOverThreshold** | Pointer to [**EnergyPlanDiscountsInnerPercentOverThreshold**](EnergyPlanDiscountsInnerPercentOverThreshold.md) |  | [optional] 

## Methods

### NewEnergyPlanDiscountsInner

`func NewEnergyPlanDiscountsInner(displayName string, type_ string, methodUType string, ) *EnergyPlanDiscountsInner`

NewEnergyPlanDiscountsInner instantiates a new EnergyPlanDiscountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanDiscountsInnerWithDefaults

`func NewEnergyPlanDiscountsInnerWithDefaults() *EnergyPlanDiscountsInner`

NewEnergyPlanDiscountsInnerWithDefaults instantiates a new EnergyPlanDiscountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanDiscountsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanDiscountsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanDiscountsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanDiscountsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanDiscountsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanDiscountsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanDiscountsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *EnergyPlanDiscountsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanDiscountsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanDiscountsInner) SetType(v string)`

SetType sets Type field to given value.


### GetCategory

`func (o *EnergyPlanDiscountsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EnergyPlanDiscountsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EnergyPlanDiscountsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EnergyPlanDiscountsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEndDate

`func (o *EnergyPlanDiscountsInner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyPlanDiscountsInner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyPlanDiscountsInner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EnergyPlanDiscountsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMethodUType

`func (o *EnergyPlanDiscountsInner) GetMethodUType() string`

GetMethodUType returns the MethodUType field if non-nil, zero value otherwise.

### GetMethodUTypeOk

`func (o *EnergyPlanDiscountsInner) GetMethodUTypeOk() (*string, bool)`

GetMethodUTypeOk returns a tuple with the MethodUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodUType

`func (o *EnergyPlanDiscountsInner) SetMethodUType(v string)`

SetMethodUType sets MethodUType field to given value.


### GetPercentOfBill

`func (o *EnergyPlanDiscountsInner) GetPercentOfBill() EnergyPlanDiscountsInnerPercentOfBill`

GetPercentOfBill returns the PercentOfBill field if non-nil, zero value otherwise.

### GetPercentOfBillOk

`func (o *EnergyPlanDiscountsInner) GetPercentOfBillOk() (*EnergyPlanDiscountsInnerPercentOfBill, bool)`

GetPercentOfBillOk returns a tuple with the PercentOfBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfBill

`func (o *EnergyPlanDiscountsInner) SetPercentOfBill(v EnergyPlanDiscountsInnerPercentOfBill)`

SetPercentOfBill sets PercentOfBill field to given value.

### HasPercentOfBill

`func (o *EnergyPlanDiscountsInner) HasPercentOfBill() bool`

HasPercentOfBill returns a boolean if a field has been set.

### GetPercentOfUse

`func (o *EnergyPlanDiscountsInner) GetPercentOfUse() EnergyPlanDiscountsInnerPercentOfUse`

GetPercentOfUse returns the PercentOfUse field if non-nil, zero value otherwise.

### GetPercentOfUseOk

`func (o *EnergyPlanDiscountsInner) GetPercentOfUseOk() (*EnergyPlanDiscountsInnerPercentOfUse, bool)`

GetPercentOfUseOk returns a tuple with the PercentOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfUse

`func (o *EnergyPlanDiscountsInner) SetPercentOfUse(v EnergyPlanDiscountsInnerPercentOfUse)`

SetPercentOfUse sets PercentOfUse field to given value.

### HasPercentOfUse

`func (o *EnergyPlanDiscountsInner) HasPercentOfUse() bool`

HasPercentOfUse returns a boolean if a field has been set.

### GetFixedAmount

`func (o *EnergyPlanDiscountsInner) GetFixedAmount() EnergyPlanDiscountsInnerFixedAmount`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *EnergyPlanDiscountsInner) GetFixedAmountOk() (*EnergyPlanDiscountsInnerFixedAmount, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *EnergyPlanDiscountsInner) SetFixedAmount(v EnergyPlanDiscountsInnerFixedAmount)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *EnergyPlanDiscountsInner) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetPercentOverThreshold

`func (o *EnergyPlanDiscountsInner) GetPercentOverThreshold() EnergyPlanDiscountsInnerPercentOverThreshold`

GetPercentOverThreshold returns the PercentOverThreshold field if non-nil, zero value otherwise.

### GetPercentOverThresholdOk

`func (o *EnergyPlanDiscountsInner) GetPercentOverThresholdOk() (*EnergyPlanDiscountsInnerPercentOverThreshold, bool)`

GetPercentOverThresholdOk returns a tuple with the PercentOverThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOverThreshold

`func (o *EnergyPlanDiscountsInner) SetPercentOverThreshold(v EnergyPlanDiscountsInnerPercentOverThreshold)`

SetPercentOverThreshold sets PercentOverThreshold field to given value.

### HasPercentOverThreshold

`func (o *EnergyPlanDiscountsInner) HasPercentOverThreshold() bool`

HasPercentOverThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


