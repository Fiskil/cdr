package energy_test

import (
	"encoding/json"
	"testing"

	"github.com/fiskil/cdr/energy"
	"github.com/matryer/is"
)

func TestPatchedCoolingOffDays(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var i energy.PatchedCoolingOffDays

	// Unmarshal a number.
	err := json.Unmarshal([]byte("1"), &i)
	is.NoErr(err)
	is.Equal(i, energy.PatchedCoolingOffDays(1))

	// Unmarshal a string.
	err = json.Unmarshal([]byte(`"1"`), &i)
	is.NoErr(err)
	is.Equal(i, energy.PatchedCoolingOffDays(1))

	bytes, err := json.Marshal(i)
	is.NoErr(err)
	is.Equal(string(bytes), "1")
}

func TestPatchedDemandChargesDays(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var d energy.PatchedDemandChargesDays

	// Unmarshal from new enum format.
	err := json.Unmarshal([]byte(`["MON", "TUE", "WED"]`), &d)
	is.NoErr(err)
	is.Equal(d, energy.PatchedDemandChargesDays{"MON", "TUE", "WED"})

	// Unmarshal from old format.
	err = json.Unmarshal([]byte(`{"weekdays": true, "saturday": false, "sunday": false}`), &d)
	is.NoErr(err)
	is.Equal(d, energy.PatchedDemandChargesDays{"MON", "TUE", "WED", "THU", "FRI"})

	bytes, err := json.Marshal(d)
	is.NoErr(err)
	is.Equal(string(bytes), `["MON","TUE","WED","THU","FRI"]`)
}

func TestPatchedEnergyPlanSolarFeedInTariffSingleTariffRate(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var d energy.PatchedEnergyPlanSolarFeedInTariffSingleTariff

	// unmarshal v2 with rates
	err := json.Unmarshal([]byte(`{"rates": [{"measureUnit": "KVAR", "unitPrice": "3.0"}]}`), &d)
	is.NoErr(err)
	is.Equal(len(d.Rates), 1)
	is.Equal(*d.Rates[0].MeasureUnit, "KVAR")
	is.Equal(d.Rates[0].UnitPrice, "3.0")

	bytes, err := json.Marshal(d)
	is.NoErr(err)
	is.Equal(string(bytes), `{"rates":[{"measureUnit":"KVAR","unitPrice":"3.0"}]}`)

	// unmarshal v1 with amount
	err = json.Unmarshal([]byte(`{"amount": "4.0"}`), &d)
	is.NoErr(err)
	is.Equal(len(d.Rates), 1)
	is.Equal(*d.Rates[0].MeasureUnit, "KWH")
	is.Equal(d.Rates[0].UnitPrice, "4.0")

	bytes, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(string(bytes), `{"rates":[{"measureUnit":"KWH","unitPrice":"4.0"}]}`)
}

func TestPatchedEnergyPlanSolarFeedInTariffTimeVaryingTariffRate(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var d energy.PatchedEnergyPlanSolarFeedInTariffTimeVaryingTariff

	// unmarshal v2 with rates
	err := json.Unmarshal([]byte(`{"rates": [{"measureUnit": "KVAR", "unitPrice": "3.0"}]}`), &d)
	is.NoErr(err)
	is.Equal(len(*d.Rates), 1)
	is.Equal(*(*d.Rates)[0].MeasureUnit, "KVAR")
	is.Equal((*d.Rates)[0].UnitPrice, "3.0")

	bytes, err := json.Marshal(d)
	is.NoErr(err)
	is.Equal(string(bytes), `{"rates":[{"measureUnit":"KVAR","unitPrice":"3.0"}],"timeVariations":null}`)

	// unmarshal v1 with amount
	err = json.Unmarshal([]byte(`{"amount": "4.0"}`), &d)
	is.NoErr(err)
	is.Equal(len(*d.Rates), 1)
	is.Equal(*(*d.Rates)[0].MeasureUnit, "KWH")
	is.Equal((*d.Rates)[0].UnitPrice, "4.0")

	bytes, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(string(bytes), `{"rates":[{"measureUnit":"KWH","unitPrice":"4.0"}],"timeVariations":null}`)
}
