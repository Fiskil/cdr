# EnergyPlanControlledLoadV2Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A display name for the controlled load. | 
**RateBlockUType** | **string** | Specifies the type of controlled load rate. | 
**StartDate** | Pointer to **string** | Optional start date of the application of the controlled load rate. | [optional] 
**EndDate** | Pointer to **string** | Optional end date of the application of the controlled load rate. | [optional] 
**SingleRate** | Pointer to [**EnergyPlanControlledLoadV2InnerSingleRate**](EnergyPlanControlledLoadV2InnerSingleRate.md) |  | [optional] 
**TimeOfUseRates** | Pointer to [**[]EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner**](EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner.md) | Array of objects representing time of use rates. Required if _rateBlockUType_ is &#x60;timeOfUseRates&#x60;. | [optional] 

## Methods

### NewEnergyPlanControlledLoadV2Inner

`func NewEnergyPlanControlledLoadV2Inner(displayName string, rateBlockUType string, ) *EnergyPlanControlledLoadV2Inner`

NewEnergyPlanControlledLoadV2Inner instantiates a new EnergyPlanControlledLoadV2Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanControlledLoadV2InnerWithDefaults

`func NewEnergyPlanControlledLoadV2InnerWithDefaults() *EnergyPlanControlledLoadV2Inner`

NewEnergyPlanControlledLoadV2InnerWithDefaults instantiates a new EnergyPlanControlledLoadV2Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnergyPlanControlledLoadV2Inner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanControlledLoadV2Inner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanControlledLoadV2Inner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRateBlockUType

`func (o *EnergyPlanControlledLoadV2Inner) GetRateBlockUType() string`

GetRateBlockUType returns the RateBlockUType field if non-nil, zero value otherwise.

### GetRateBlockUTypeOk

`func (o *EnergyPlanControlledLoadV2Inner) GetRateBlockUTypeOk() (*string, bool)`

GetRateBlockUTypeOk returns a tuple with the RateBlockUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBlockUType

`func (o *EnergyPlanControlledLoadV2Inner) SetRateBlockUType(v string)`

SetRateBlockUType sets RateBlockUType field to given value.


### GetStartDate

`func (o *EnergyPlanControlledLoadV2Inner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyPlanControlledLoadV2Inner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyPlanControlledLoadV2Inner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EnergyPlanControlledLoadV2Inner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EnergyPlanControlledLoadV2Inner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyPlanControlledLoadV2Inner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyPlanControlledLoadV2Inner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EnergyPlanControlledLoadV2Inner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetSingleRate

`func (o *EnergyPlanControlledLoadV2Inner) GetSingleRate() EnergyPlanControlledLoadV2InnerSingleRate`

GetSingleRate returns the SingleRate field if non-nil, zero value otherwise.

### GetSingleRateOk

`func (o *EnergyPlanControlledLoadV2Inner) GetSingleRateOk() (*EnergyPlanControlledLoadV2InnerSingleRate, bool)`

GetSingleRateOk returns a tuple with the SingleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleRate

`func (o *EnergyPlanControlledLoadV2Inner) SetSingleRate(v EnergyPlanControlledLoadV2InnerSingleRate)`

SetSingleRate sets SingleRate field to given value.

### HasSingleRate

`func (o *EnergyPlanControlledLoadV2Inner) HasSingleRate() bool`

HasSingleRate returns a boolean if a field has been set.

### GetTimeOfUseRates

`func (o *EnergyPlanControlledLoadV2Inner) GetTimeOfUseRates() []EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner`

GetTimeOfUseRates returns the TimeOfUseRates field if non-nil, zero value otherwise.

### GetTimeOfUseRatesOk

`func (o *EnergyPlanControlledLoadV2Inner) GetTimeOfUseRatesOk() (*[]EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner, bool)`

GetTimeOfUseRatesOk returns a tuple with the TimeOfUseRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfUseRates

`func (o *EnergyPlanControlledLoadV2Inner) SetTimeOfUseRates(v []EnergyPlanControlledLoadV2InnerTimeOfUseRatesInner)`

SetTimeOfUseRates sets TimeOfUseRates field to given value.

### HasTimeOfUseRates

`func (o *EnergyPlanControlledLoadV2Inner) HasTimeOfUseRates() bool`

HasTimeOfUseRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


