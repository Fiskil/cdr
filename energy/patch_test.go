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
