package utils

import (
    "testing"
)

func TestValidateBRSRInput_Valid(t *testing.T) {
    validInput := map[string]interface{}{
        "company": "Tata Motors",
        "financialYear": "2023-24",
        "sectionA": map[string]interface{}{
            "A1": "Tata Motors",
            "A2": "L28920MH1945PLC004520",
        },
        "sectionB": map[string]interface{}{
            "B1": "Board has oversight",
        },
        "sectionC": map[string]interface{}{
            "C1_P1_Q1": "Yes",
            "C1_P2_Q1": "Yes",
        },
        "signatories": []interface{}{
            map[string]interface{}{
                "name": "John Doe",
                "role": "CFO",
                "date": "2024-04-20",
                "signed": true,
            },
        },
    }

    err := ValidateBRSRInput(validInput)
    if err != nil {
        t.Errorf("Expected valid input to pass, but got error: %v", err)
    }
}

func TestValidateBRSRInput_Invalid(t *testing.T) {
    invalidInput := map[string]interface{}{
        "company": "Tata Motors",
        // Missing required sectionB and sectionC
        "financialYear": "2023-24",
        "sectionA": map[string]interface{}{
            "A1": "Tata Motors",
        },
    }

    err := ValidateBRSRInput(invalidInput)
    if err == nil {
        t.Errorf("Expected error for invalid input but got nil")
    } else {
        t.Logf("Correctly failed validation: %v", err)
    }
}
