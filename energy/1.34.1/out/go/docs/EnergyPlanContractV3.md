# EnergyPlanContractV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalFeeInformation** | Pointer to **string** | Free text field containing additional information of the fees for this contract. | [optional] 
**PricingModel** | **string** | The pricing model for the contract. Contracts for gas must use &#x60;SINGLE_RATE&#x60;. Note that the detail for the enumeration values are:&lt;ul&gt;&lt;li&gt;&#x60;SINGLE_RATE&#x60;: all energy usage is charged at a single unit rate no matter when it is consumed. Multiple unit rates may exist that correspond to varying volumes of usage i.e. a &#39;block&#39; or &#39;step&#39; tariff (first 50kWh @ X cents, next 50kWh at Y cents etc.)&lt;/li&gt;&lt;li&gt;&#x60;SINGLE_RATE_CONT_LOAD&#x60;: as above, but with an additional, separate unit rate charged for all energy usage from a controlled load i.e. separately metered appliance like hot water service, pool pump etc.&lt;/li&gt;&lt;li&gt;&#x60;TIME_OF_USE&#x60;: energy usage is charged at unit rates that vary dependent on time of day and day of week that the energy is consumed&lt;/li&gt;&lt;li&gt;&#x60;TIME_OF_USE_CONT_LOAD&#x60;: as above, but with an additional, separate unit rate charged for all energy usage from a controlled load i.e. separately metered appliance like hot water service, pool pump etc.&lt;/li&gt;&lt;li&gt;&#x60;FLEXIBLE&#x60;: energy usage is charged at unit rates that vary based on external factors&lt;/li&gt;&lt;li&gt;&#x60;FLEXIBLE_CONT_LOAD&#x60;: as above, but with an additional, separate unit rate charged for all energy usage from a controlled load i.e. separately metered appliance like hot water service, pool pump etc.&lt;/li&gt;&lt;li&gt;&#x60;QUOTA&#x60;: all energy usage is charged at a single fixed rate, up to a specified usage quota/allowance. All excess usage beyond the allowance is then charged at a single unit rate. i.e. $50/month for up to 150kWh included usage.&lt;/li&gt;&lt;/ul&gt; | 
**TimeZone** | Pointer to **string** | Required if _pricingModel_ is set to &#x60;TIME_OF_USE&#x60;. Defines the time zone to use for calculation of the time of use thresholds. Defaults to &#x60;AEST&#x60; if absent. | [optional] [default to "AEST"]
**IsFixed** | **bool** | Flag indicating whether prices are fixed or variable. | 
**Variation** | Pointer to **string** | Free text description of price variation policy and conditions for the contract. Mandatory if _isFixed_ is &#x60;false&#x60;. | [optional] 
**OnExpiryDescription** | Pointer to **string** | Free text field that describes what will occur on or prior to expiry of the fixed contract term or benefit period. | [optional] 
**PaymentOption** | **[]string** | Payment options for this contract. | 
**IntrinsicGreenPower** | Pointer to [**EnergyPlanContractV3IntrinsicGreenPower**](EnergyPlanContractV3IntrinsicGreenPower.md) |  | [optional] 
**ControlledLoad** | Pointer to [**[]EnergyPlanControlledLoadV2Inner**](EnergyPlanControlledLoadV2Inner.md) | Required if pricing model is &#x60;SINGLE_RATE_CONT_LOAD&#x60; or &#x60;TIME_OF_USE_CONT_LOAD&#x60; or &#x60;FLEXIBLE_CONT_LOAD&#x60;. | [optional] 
**Incentives** | Pointer to [**[]EnergyPlanIncentivesInner**](EnergyPlanIncentivesInner.md) | Optional list of incentives available for the contract. | [optional] 
**Discounts** | Pointer to [**[]EnergyPlanDiscountsInner**](EnergyPlanDiscountsInner.md) | Optional list of discounts available for the contract. | [optional] 
**GreenPowerCharges** | Pointer to [**[]EnergyPlanGreenPowerChargesInner**](EnergyPlanGreenPowerChargesInner.md) | Optional list of charges applicable to green power. | [optional] 
**Eligibility** | Pointer to [**[]EnergyPlanEligibilityInner**](EnergyPlanEligibilityInner.md) | Eligibility restrictions or requirements. | [optional] 
**Fees** | Pointer to [**[]EnergyPlanFeesInner**](EnergyPlanFeesInner.md) | An array of fees applicable to the plan. | [optional] 
**SolarFeedInTariff** | Pointer to [**[]EnergyPlanSolarFeedInTariffV3Inner**](EnergyPlanSolarFeedInTariffV3Inner.md) | Array of feed in tariffs for solar power. | [optional] 
**TariffPeriod** | [**[]EnergyPlanTariffPeriodV2Inner**](EnergyPlanTariffPeriodV2Inner.md) | Array of tariff periods. | 

## Methods

### NewEnergyPlanContractV3

`func NewEnergyPlanContractV3(pricingModel string, isFixed bool, paymentOption []string, tariffPeriod []EnergyPlanTariffPeriodV2Inner, ) *EnergyPlanContractV3`

NewEnergyPlanContractV3 instantiates a new EnergyPlanContractV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyPlanContractV3WithDefaults

`func NewEnergyPlanContractV3WithDefaults() *EnergyPlanContractV3`

NewEnergyPlanContractV3WithDefaults instantiates a new EnergyPlanContractV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalFeeInformation

`func (o *EnergyPlanContractV3) GetAdditionalFeeInformation() string`

GetAdditionalFeeInformation returns the AdditionalFeeInformation field if non-nil, zero value otherwise.

### GetAdditionalFeeInformationOk

`func (o *EnergyPlanContractV3) GetAdditionalFeeInformationOk() (*string, bool)`

GetAdditionalFeeInformationOk returns a tuple with the AdditionalFeeInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFeeInformation

`func (o *EnergyPlanContractV3) SetAdditionalFeeInformation(v string)`

SetAdditionalFeeInformation sets AdditionalFeeInformation field to given value.

### HasAdditionalFeeInformation

`func (o *EnergyPlanContractV3) HasAdditionalFeeInformation() bool`

HasAdditionalFeeInformation returns a boolean if a field has been set.

### GetPricingModel

`func (o *EnergyPlanContractV3) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *EnergyPlanContractV3) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *EnergyPlanContractV3) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetTimeZone

`func (o *EnergyPlanContractV3) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *EnergyPlanContractV3) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *EnergyPlanContractV3) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *EnergyPlanContractV3) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetIsFixed

`func (o *EnergyPlanContractV3) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *EnergyPlanContractV3) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *EnergyPlanContractV3) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.


### GetVariation

`func (o *EnergyPlanContractV3) GetVariation() string`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *EnergyPlanContractV3) GetVariationOk() (*string, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *EnergyPlanContractV3) SetVariation(v string)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *EnergyPlanContractV3) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetOnExpiryDescription

`func (o *EnergyPlanContractV3) GetOnExpiryDescription() string`

GetOnExpiryDescription returns the OnExpiryDescription field if non-nil, zero value otherwise.

### GetOnExpiryDescriptionOk

`func (o *EnergyPlanContractV3) GetOnExpiryDescriptionOk() (*string, bool)`

GetOnExpiryDescriptionOk returns a tuple with the OnExpiryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnExpiryDescription

`func (o *EnergyPlanContractV3) SetOnExpiryDescription(v string)`

SetOnExpiryDescription sets OnExpiryDescription field to given value.

### HasOnExpiryDescription

`func (o *EnergyPlanContractV3) HasOnExpiryDescription() bool`

HasOnExpiryDescription returns a boolean if a field has been set.

### GetPaymentOption

`func (o *EnergyPlanContractV3) GetPaymentOption() []string`

GetPaymentOption returns the PaymentOption field if non-nil, zero value otherwise.

### GetPaymentOptionOk

`func (o *EnergyPlanContractV3) GetPaymentOptionOk() (*[]string, bool)`

GetPaymentOptionOk returns a tuple with the PaymentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOption

`func (o *EnergyPlanContractV3) SetPaymentOption(v []string)`

SetPaymentOption sets PaymentOption field to given value.


### GetIntrinsicGreenPower

`func (o *EnergyPlanContractV3) GetIntrinsicGreenPower() EnergyPlanContractV3IntrinsicGreenPower`

GetIntrinsicGreenPower returns the IntrinsicGreenPower field if non-nil, zero value otherwise.

### GetIntrinsicGreenPowerOk

`func (o *EnergyPlanContractV3) GetIntrinsicGreenPowerOk() (*EnergyPlanContractV3IntrinsicGreenPower, bool)`

GetIntrinsicGreenPowerOk returns a tuple with the IntrinsicGreenPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrinsicGreenPower

`func (o *EnergyPlanContractV3) SetIntrinsicGreenPower(v EnergyPlanContractV3IntrinsicGreenPower)`

SetIntrinsicGreenPower sets IntrinsicGreenPower field to given value.

### HasIntrinsicGreenPower

`func (o *EnergyPlanContractV3) HasIntrinsicGreenPower() bool`

HasIntrinsicGreenPower returns a boolean if a field has been set.

### GetControlledLoad

`func (o *EnergyPlanContractV3) GetControlledLoad() []EnergyPlanControlledLoadV2Inner`

GetControlledLoad returns the ControlledLoad field if non-nil, zero value otherwise.

### GetControlledLoadOk

`func (o *EnergyPlanContractV3) GetControlledLoadOk() (*[]EnergyPlanControlledLoadV2Inner, bool)`

GetControlledLoadOk returns a tuple with the ControlledLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlledLoad

`func (o *EnergyPlanContractV3) SetControlledLoad(v []EnergyPlanControlledLoadV2Inner)`

SetControlledLoad sets ControlledLoad field to given value.

### HasControlledLoad

`func (o *EnergyPlanContractV3) HasControlledLoad() bool`

HasControlledLoad returns a boolean if a field has been set.

### GetIncentives

`func (o *EnergyPlanContractV3) GetIncentives() []EnergyPlanIncentivesInner`

GetIncentives returns the Incentives field if non-nil, zero value otherwise.

### GetIncentivesOk

`func (o *EnergyPlanContractV3) GetIncentivesOk() (*[]EnergyPlanIncentivesInner, bool)`

GetIncentivesOk returns a tuple with the Incentives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncentives

`func (o *EnergyPlanContractV3) SetIncentives(v []EnergyPlanIncentivesInner)`

SetIncentives sets Incentives field to given value.

### HasIncentives

`func (o *EnergyPlanContractV3) HasIncentives() bool`

HasIncentives returns a boolean if a field has been set.

### GetDiscounts

`func (o *EnergyPlanContractV3) GetDiscounts() []EnergyPlanDiscountsInner`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *EnergyPlanContractV3) GetDiscountsOk() (*[]EnergyPlanDiscountsInner, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *EnergyPlanContractV3) SetDiscounts(v []EnergyPlanDiscountsInner)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *EnergyPlanContractV3) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetGreenPowerCharges

`func (o *EnergyPlanContractV3) GetGreenPowerCharges() []EnergyPlanGreenPowerChargesInner`

GetGreenPowerCharges returns the GreenPowerCharges field if non-nil, zero value otherwise.

### GetGreenPowerChargesOk

`func (o *EnergyPlanContractV3) GetGreenPowerChargesOk() (*[]EnergyPlanGreenPowerChargesInner, bool)`

GetGreenPowerChargesOk returns a tuple with the GreenPowerCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreenPowerCharges

`func (o *EnergyPlanContractV3) SetGreenPowerCharges(v []EnergyPlanGreenPowerChargesInner)`

SetGreenPowerCharges sets GreenPowerCharges field to given value.

### HasGreenPowerCharges

`func (o *EnergyPlanContractV3) HasGreenPowerCharges() bool`

HasGreenPowerCharges returns a boolean if a field has been set.

### GetEligibility

`func (o *EnergyPlanContractV3) GetEligibility() []EnergyPlanEligibilityInner`

GetEligibility returns the Eligibility field if non-nil, zero value otherwise.

### GetEligibilityOk

`func (o *EnergyPlanContractV3) GetEligibilityOk() (*[]EnergyPlanEligibilityInner, bool)`

GetEligibilityOk returns a tuple with the Eligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibility

`func (o *EnergyPlanContractV3) SetEligibility(v []EnergyPlanEligibilityInner)`

SetEligibility sets Eligibility field to given value.

### HasEligibility

`func (o *EnergyPlanContractV3) HasEligibility() bool`

HasEligibility returns a boolean if a field has been set.

### GetFees

`func (o *EnergyPlanContractV3) GetFees() []EnergyPlanFeesInner`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *EnergyPlanContractV3) GetFeesOk() (*[]EnergyPlanFeesInner, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *EnergyPlanContractV3) SetFees(v []EnergyPlanFeesInner)`

SetFees sets Fees field to given value.

### HasFees

`func (o *EnergyPlanContractV3) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetSolarFeedInTariff

`func (o *EnergyPlanContractV3) GetSolarFeedInTariff() []EnergyPlanSolarFeedInTariffV3Inner`

GetSolarFeedInTariff returns the SolarFeedInTariff field if non-nil, zero value otherwise.

### GetSolarFeedInTariffOk

`func (o *EnergyPlanContractV3) GetSolarFeedInTariffOk() (*[]EnergyPlanSolarFeedInTariffV3Inner, bool)`

GetSolarFeedInTariffOk returns a tuple with the SolarFeedInTariff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarFeedInTariff

`func (o *EnergyPlanContractV3) SetSolarFeedInTariff(v []EnergyPlanSolarFeedInTariffV3Inner)`

SetSolarFeedInTariff sets SolarFeedInTariff field to given value.

### HasSolarFeedInTariff

`func (o *EnergyPlanContractV3) HasSolarFeedInTariff() bool`

HasSolarFeedInTariff returns a boolean if a field has been set.

### GetTariffPeriod

`func (o *EnergyPlanContractV3) GetTariffPeriod() []EnergyPlanTariffPeriodV2Inner`

GetTariffPeriod returns the TariffPeriod field if non-nil, zero value otherwise.

### GetTariffPeriodOk

`func (o *EnergyPlanContractV3) GetTariffPeriodOk() (*[]EnergyPlanTariffPeriodV2Inner, bool)`

GetTariffPeriodOk returns a tuple with the TariffPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffPeriod

`func (o *EnergyPlanContractV3) SetTariffPeriod(v []EnergyPlanTariffPeriodV2Inner)`

SetTariffPeriod sets TariffPeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


