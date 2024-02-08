package energy

import (
	"encoding/json"
)

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
