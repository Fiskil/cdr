package energy

import "encoding/json"

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
