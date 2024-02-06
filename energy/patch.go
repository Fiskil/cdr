package energy

import (
	"encoding/json"
	"strconv"
)

// PatchedCoolingOffDays is a type that can unmarshal an int from both a number and a string.
// issue: https://github.com/ConsumerDataStandardsAustralia/standards-maintenance/issues/582
type PatchedCoolingOffDays int

// UnmarshalJSON implements the json.Unmarshaller interface.
func (pc *PatchedCoolingOffDays) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(pc))
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	*pc = PatchedCoolingOffDays(i)
	return nil
}

// PatchedDemandChargesDays is a type that can unmarshal from new enum format and also from old format
// issue: https://github.com/ConsumerDataStandardsAustralia/standards-maintenance/issues/502
type PatchedDemandChargesDays []EnergyPlanTariffPeriodDemandChargesDays

type oldDemandChargesDays struct {
	Weekdays bool `json:"weekdays"`
	Saturday bool `json:"saturday"`
	Sunday   bool `json:"sunday"`
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (pd *PatchedDemandChargesDays) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		return json.Unmarshal(b, (*[]EnergyPlanTariffPeriodDemandChargesDays)(pd))
	}

	var oldDays oldDemandChargesDays
	if err := json.Unmarshal(b, &oldDays); err != nil {
		return err
	}

	var newDays []EnergyPlanTariffPeriodDemandChargesDays = make([]EnergyPlanTariffPeriodDemandChargesDays, 0)
	if oldDays.Weekdays {
		newDays = append(newDays,
			EnergyPlanTariffPeriodDemandChargesDaysMON,
			EnergyPlanTariffPeriodDemandChargesDaysTUE,
			EnergyPlanTariffPeriodDemandChargesDaysWED,
			EnergyPlanTariffPeriodDemandChargesDaysTHU,
			EnergyPlanTariffPeriodDemandChargesDaysFRI)
	}

	if oldDays.Saturday {
		newDays = append(newDays, EnergyPlanTariffPeriodDemandChargesDaysSAT)
	}

	if oldDays.Sunday {
		newDays = append(newDays, EnergyPlanTariffPeriodDemandChargesDaysSUN)
	}

	*pd = newDays
	return nil
}

type Rate struct {
	// MeasureUnit The measurement unit of rate. Assumed to be KWH if absent
	MeasureUnit *string `json:"measureUnit,omitempty"`
	// UnitPrice Unit price of usage per measure unit (exclusive of GST)
	UnitPrice string `json:"unitPrice"`
	// Volume Volume that this rate applies to. Only applicable for ‘stepped’ rates where different rates apply for different volumes of usage in a period
	Volume *float32 `json:"volume,omitempty"`
}

type PatchedEnergyPlanSolarFeedInTariffSingleTariff struct {
	// Amount is to support parsing EnergyPlanSolarFeedInTariff v1 response format
	Amount string `json:"amount,omitempty"`
	// Rates Array of feed in rates
	Rates []PatchedEnergyPlanSolarFeedInTariffSingleTariffRate `json:"rates"`
}

type PatchedEnergyPlanSolarFeedInTariffSingleTariffRate Rate

// UnmarshalJSON implements the json.Unmarshaller interface.
func (tariff *PatchedEnergyPlanSolarFeedInTariffSingleTariff) UnmarshalJSON(b []byte) error {
	type auxEnergyPlanSolarFeedInTariffSingleTariff PatchedEnergyPlanSolarFeedInTariffSingleTariff

	var auxTariff auxEnergyPlanSolarFeedInTariffSingleTariff
	if err := json.Unmarshal(b, &auxTariff); err != nil {
		return err
	}

	if auxTariff.Amount != "" && len(auxTariff.Rates) == 0 {
		unit := "KWH" // default to KWH
		auxTariff.Rates = []PatchedEnergyPlanSolarFeedInTariffSingleTariffRate{
			{
				UnitPrice:   auxTariff.Amount,
				MeasureUnit: &unit,
			},
		}
		auxTariff.Amount = ""
	}

	*tariff = PatchedEnergyPlanSolarFeedInTariffSingleTariff(auxTariff)
	return nil
}

type PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariff struct {
	// Amount is to support parsing EnergyPlanSolarFeedInTariff v1 response format
	Amount string `json:"amount,omitempty"`
	// Rates Array of feed in rates
	Rates *[]PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariffRate `json:"rates"`
	// TimeVariations Array of time periods for which this tariff is applicable
	TimeVariations []struct {
		// Days The days that the tariff applies to. At least one entry required
		Days []EnergyPlanSolarFeedInTariffV2TimeVaryingTariffsTimeVariationsDays `json:"days"`

		// EndTime The end of the time period per day for which the tariff applies.  If absent assumes end of day (ie. one second before midnight)
		EndTime *string `json:"endTime,omitempty"`

		// StartTime The beginning of the time period per day for which the tariff applies.  If absent assumes start of day (ie. midnight)
		StartTime *string `json:"startTime,omitempty"`
	} `json:"timeVariations"`

	// Type The type of the charging time period. If absent applies to all periods
	Type *EnergyPlanSolarFeedInTariffV2TimeVaryingTariffsType `json:"type,omitempty"`
}

type PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariffRate Rate

// UnmarshalJSON implements the json.Unmarshaller interface.
func (tariff *PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariff) UnmarshalJSON(b []byte) error {
	type auxEnergyPlanSolarFeedInTariffTimeVaryingTariff PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariff

	var auxTariff auxEnergyPlanSolarFeedInTariffTimeVaryingTariff
	if err := json.Unmarshal(b, &auxTariff); err != nil {
		return err
	}

	if auxTariff.Amount != "" && auxTariff.Rates == nil {
		unit := "KWH" // default to KWH
		auxTariff.Rates = &[]PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariffRate{
			{
				UnitPrice:   auxTariff.Amount,
				MeasureUnit: &unit,
			},
		}
		auxTariff.Amount = ""
	}

	*tariff = PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariff(auxTariff)
	return nil
}
