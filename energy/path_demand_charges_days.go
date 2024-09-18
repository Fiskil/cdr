package energy

import "encoding/json"

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
