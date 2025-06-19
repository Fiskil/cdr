# EnergyPlanSolarFeedInTariffV3Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The name of the tariff. | 
**Description** | Pointer to **string** | A description of the tariff. | [optional] 
**StartDate** | Pointer to **string** | The start date of the application of the feed in tariff. | [optional] 
**EndDate** | Pointer to **string** | The end date of the application of the feed in tariff. | [optional] 
**Scheme** | **string** | The applicable scheme. | 
**PayerType** | **string** | The type of the payer. | 
**TariffUType** | **string** | Reference to the applicable tariff structure. | 
**SingleTariff** | Pointer to [**EnergyPlanSolarFeedInTariffV3InnerSingleTariff**](EnergyPlanSolarFeedInTariffV3InnerSingleTariff.md) |  | [optional] 
**TimeVaryingTariffs** | Pointer to [**[]EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner**](EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner.md) | Represents a tariff based on time. Mandatory if _tariffUType_ is set to &#x60;timeVaryingTariffs&#x60;. | [optional] 

## Methods

### NewEnergyPlanSolarFeedInTariffV3Inner

`func NewEnergyPlanSolarFeedInTariffV3Inner(displayName string, scheme string, payerType string, tariffUType string, ) *EnergyPlanSolarFeedInTariffV3Inner`

NewEnergyPlanSolarFeedInTariffV3Inner instantiates a new EnergyPlanSolarFeedInTariffV3Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanSolarFeedInTariffV3InnerWithDefaults

`func NewEnergyPlanSolarFeedInTariffV3InnerWithDefaults() *EnergyPlanSolarFeedInTariffV3Inner`

NewEnergyPlanSolarFeedInTariffV3InnerWithDefaults instantiates a new EnergyPlanSolarFeedInTariffV3Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnergyPlanSolarFeedInTariffV3Inner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EnergyPlanSolarFeedInTariffV3Inner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EnergyPlanSolarFeedInTariffV3Inner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetScheme

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetPayerType

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetPayerType() string`

GetPayerType returns the PayerType field if non-nil, zero value otherwise.

### GetPayerTypeOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetPayerTypeOk() (*string, bool)`

GetPayerTypeOk returns a tuple with the PayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerType

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetPayerType(v string)`

SetPayerType sets PayerType field to given value.


### GetTariffUType

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetTariffUType() string`

GetTariffUType returns the TariffUType field if non-nil, zero value otherwise.

### GetTariffUTypeOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetTariffUTypeOk() (*string, bool)`

GetTariffUTypeOk returns a tuple with the TariffUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffUType

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetTariffUType(v string)`

SetTariffUType sets TariffUType field to given value.


### GetSingleTariff

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetSingleTariff() EnergyPlanSolarFeedInTariffV3InnerSingleTariff`

GetSingleTariff returns the SingleTariff field if non-nil, zero value otherwise.

### GetSingleTariffOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetSingleTariffOk() (*EnergyPlanSolarFeedInTariffV3InnerSingleTariff, bool)`

GetSingleTariffOk returns a tuple with the SingleTariff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTariff

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetSingleTariff(v EnergyPlanSolarFeedInTariffV3InnerSingleTariff)`

SetSingleTariff sets SingleTariff field to given value.

### HasSingleTariff

`func (o *EnergyPlanSolarFeedInTariffV3Inner) HasSingleTariff() bool`

HasSingleTariff returns a boolean if a field has been set.

### GetTimeVaryingTariffs

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetTimeVaryingTariffs() []EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner`

GetTimeVaryingTariffs returns the TimeVaryingTariffs field if non-nil, zero value otherwise.

### GetTimeVaryingTariffsOk

`func (o *EnergyPlanSolarFeedInTariffV3Inner) GetTimeVaryingTariffsOk() (*[]EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner, bool)`

GetTimeVaryingTariffsOk returns a tuple with the TimeVaryingTariffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeVaryingTariffs

`func (o *EnergyPlanSolarFeedInTariffV3Inner) SetTimeVaryingTariffs(v []EnergyPlanSolarFeedInTariffV3InnerTimeVaryingTariffsInner)`

SetTimeVaryingTariffs sets TimeVaryingTariffs field to given value.

### HasTimeVaryingTariffs

`func (o *EnergyPlanSolarFeedInTariffV3Inner) HasTimeVaryingTariffs() bool`

HasTimeVaryingTariffs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


