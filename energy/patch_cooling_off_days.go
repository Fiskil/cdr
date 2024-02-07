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
