package utils

import (
	"testing"
)

func TestValidateBRSRJSON(t *testing.T) {
	err := ValidateBRSRJSON("./schemas/test_input.json", "./schemas/brsr_schema.json")
	if err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}
