# EnergyPlanTariffPeriodV2Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of charge. Assumed to be &#x60;OTHER&#x60; if absent. | [optional] [default to "OTHER"]
**DisplayName** | **string** | The name of the tariff period. | 
**StartDate** | **string** | The start date of the tariff period in a calendar year. Formatted in mm-dd format. | 
**EndDate** | **string** | The end date of the tariff period in a calendar year. Formatted in mm-dd format. | 
**DailySupplyChargeType** | Pointer to **string** | Specifies if daily supply charge is single or banded. | [optional] 
**DailySupplyCharge** | Pointer to **string** | The amount of access charge for the tariff period, in dollars per day exclusive of GST. Mandatory if _dailySupplyChargeType_ is &#x60;SINGLE&#x60;. | [optional] 
**BandedDailySupplyCharges** | Pointer to [**[]EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner**](EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner.md) | Array representing banded daily supply charge rates. Mandatory if _dailySupplyChargeType_ is &#x60;BAND&#x60;. | [optional] 
**TimeZone** | Pointer to **string** | Specifies the charge specific time zone for calculation of the time of use thresholds. If absent, timezone value in EnergyPlanContract is assumed. | [optional] 
**RateBlockUType** | **string** | Specifies the type of rate applicable to this tariff period. | 
**SingleRate** | Pointer to [**EnergyPlanTariffPeriodV2InnerSingleRate**](EnergyPlanTariffPeriodV2InnerSingleRate.md) |  | [optional] 
**TimeOfUseRates** | Pointer to [**[]EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner**](EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner.md) | Array of objects representing time of use rates. Required if _rateBlockUType_ is &#x60;timeOfUseRates&#x60;. | [optional] 
**DemandCharges** | Pointer to [**[]EnergyPlanTariffPeriodV2InnerDemandChargesInner**](EnergyPlanTariffPeriodV2InnerDemandChargesInner.md) | Array of demand charges. Required if _rateBlockUType_ is &#x60;demandCharges&#x60;. | [optional] 

## Methods

### NewEnergyPlanTariffPeriodV2Inner

`func NewEnergyPlanTariffPeriodV2Inner(displayName string, startDate string, endDate string, rateBlockUType string, ) *EnergyPlanTariffPeriodV2Inner`

NewEnergyPlanTariffPeriodV2Inner instantiates a new EnergyPlanTariffPeriodV2Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanTariffPeriodV2InnerWithDefaults

`func NewEnergyPlanTariffPeriodV2InnerWithDefaults() *EnergyPlanTariffPeriodV2Inner`

NewEnergyPlanTariffPeriodV2InnerWithDefaults instantiates a new EnergyPlanTariffPeriodV2Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnergyPlanTariffPeriodV2Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnergyPlanTariffPeriodV2Inner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnergyPlanTariffPeriodV2Inner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnergyPlanTariffPeriodV2Inner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnergyPlanTariffPeriodV2Inner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetStartDate

`func (o *EnergyPlanTariffPeriodV2Inner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EnergyPlanTariffPeriodV2Inner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *EnergyPlanTariffPeriodV2Inner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EnergyPlanTariffPeriodV2Inner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetDailySupplyChargeType

`func (o *EnergyPlanTariffPeriodV2Inner) GetDailySupplyChargeType() string`

GetDailySupplyChargeType returns the DailySupplyChargeType field if non-nil, zero value otherwise.

### GetDailySupplyChargeTypeOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetDailySupplyChargeTypeOk() (*string, bool)`

GetDailySupplyChargeTypeOk returns a tuple with the DailySupplyChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySupplyChargeType

`func (o *EnergyPlanTariffPeriodV2Inner) SetDailySupplyChargeType(v string)`

SetDailySupplyChargeType sets DailySupplyChargeType field to given value.

### HasDailySupplyChargeType

`func (o *EnergyPlanTariffPeriodV2Inner) HasDailySupplyChargeType() bool`

HasDailySupplyChargeType returns a boolean if a field has been set.

### GetDailySupplyCharge

`func (o *EnergyPlanTariffPeriodV2Inner) GetDailySupplyCharge() string`

GetDailySupplyCharge returns the DailySupplyCharge field if non-nil, zero value otherwise.

### GetDailySupplyChargeOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetDailySupplyChargeOk() (*string, bool)`

GetDailySupplyChargeOk returns a tuple with the DailySupplyCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySupplyCharge

`func (o *EnergyPlanTariffPeriodV2Inner) SetDailySupplyCharge(v string)`

SetDailySupplyCharge sets DailySupplyCharge field to given value.

### HasDailySupplyCharge

`func (o *EnergyPlanTariffPeriodV2Inner) HasDailySupplyCharge() bool`

HasDailySupplyCharge returns a boolean if a field has been set.

### GetBandedDailySupplyCharges

`func (o *EnergyPlanTariffPeriodV2Inner) GetBandedDailySupplyCharges() []EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner`

GetBandedDailySupplyCharges returns the BandedDailySupplyCharges field if non-nil, zero value otherwise.

### GetBandedDailySupplyChargesOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetBandedDailySupplyChargesOk() (*[]EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner, bool)`

GetBandedDailySupplyChargesOk returns a tuple with the BandedDailySupplyCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandedDailySupplyCharges

`func (o *EnergyPlanTariffPeriodV2Inner) SetBandedDailySupplyCharges(v []EnergyPlanTariffPeriodV2InnerBandedDailySupplyChargesInner)`

SetBandedDailySupplyCharges sets BandedDailySupplyCharges field to given value.

### HasBandedDailySupplyCharges

`func (o *EnergyPlanTariffPeriodV2Inner) HasBandedDailySupplyCharges() bool`

HasBandedDailySupplyCharges returns a boolean if a field has been set.

### GetTimeZone

`func (o *EnergyPlanTariffPeriodV2Inner) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *EnergyPlanTariffPeriodV2Inner) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *EnergyPlanTariffPeriodV2Inner) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetRateBlockUType

`func (o *EnergyPlanTariffPeriodV2Inner) GetRateBlockUType() string`

GetRateBlockUType returns the RateBlockUType field if non-nil, zero value otherwise.

### GetRateBlockUTypeOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetRateBlockUTypeOk() (*string, bool)`

GetRateBlockUTypeOk returns a tuple with the RateBlockUType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBlockUType

`func (o *EnergyPlanTariffPeriodV2Inner) SetRateBlockUType(v string)`

SetRateBlockUType sets RateBlockUType field to given value.


### GetSingleRate

`func (o *EnergyPlanTariffPeriodV2Inner) GetSingleRate() EnergyPlanTariffPeriodV2InnerSingleRate`

GetSingleRate returns the SingleRate field if non-nil, zero value otherwise.

### GetSingleRateOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetSingleRateOk() (*EnergyPlanTariffPeriodV2InnerSingleRate, bool)`

GetSingleRateOk returns a tuple with the SingleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleRate

`func (o *EnergyPlanTariffPeriodV2Inner) SetSingleRate(v EnergyPlanTariffPeriodV2InnerSingleRate)`

SetSingleRate sets SingleRate field to given value.

### HasSingleRate

`func (o *EnergyPlanTariffPeriodV2Inner) HasSingleRate() bool`

HasSingleRate returns a boolean if a field has been set.

### GetTimeOfUseRates

`func (o *EnergyPlanTariffPeriodV2Inner) GetTimeOfUseRates() []EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner`

GetTimeOfUseRates returns the TimeOfUseRates field if non-nil, zero value otherwise.

### GetTimeOfUseRatesOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetTimeOfUseRatesOk() (*[]EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner, bool)`

GetTimeOfUseRatesOk returns a tuple with the TimeOfUseRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfUseRates

`func (o *EnergyPlanTariffPeriodV2Inner) SetTimeOfUseRates(v []EnergyPlanTariffPeriodV2InnerTimeOfUseRatesInner)`

SetTimeOfUseRates sets TimeOfUseRates field to given value.

### HasTimeOfUseRates

`func (o *EnergyPlanTariffPeriodV2Inner) HasTimeOfUseRates() bool`

HasTimeOfUseRates returns a boolean if a field has been set.

### GetDemandCharges

`func (o *EnergyPlanTariffPeriodV2Inner) GetDemandCharges() []EnergyPlanTariffPeriodV2InnerDemandChargesInner`

GetDemandCharges returns the DemandCharges field if non-nil, zero value otherwise.

### GetDemandChargesOk

`func (o *EnergyPlanTariffPeriodV2Inner) GetDemandChargesOk() (*[]EnergyPlanTariffPeriodV2InnerDemandChargesInner, bool)`

GetDemandChargesOk returns a tuple with the DemandCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandCharges

`func (o *EnergyPlanTariffPeriodV2Inner) SetDemandCharges(v []EnergyPlanTariffPeriodV2InnerDemandChargesInner)`

SetDemandCharges sets DemandCharges field to given value.

### HasDemandCharges

`func (o *EnergyPlanTariffPeriodV2Inner) HasDemandCharges() bool`

HasDemandCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


