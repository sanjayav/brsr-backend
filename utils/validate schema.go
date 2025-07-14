package utils

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

// ValidateBRSRJSON validates a BRSR input JSON against the brsr_full_schema.json
func ValidateBRSRJSON(jsonPath string, schemaPath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewReferenceLoader("file://" + jsonPath)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	if !result.Valid() {
		errMsg := "JSON validation failed:\n"
		for _, desc := range result.Errors() {
			errMsg += fmt.Sprintf("- %s\n", desc)
		}
		return fmt.Errorf("%s", errMsg)
	}

	return nil
}
