package test

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"testing"
	"timertracker/internel/types"
	"timertracker/internel/utils"
)

func TestPasportValidation(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("passport", utils.ValidatePassport)

	testCases := []struct {
		name          string
		passport      string
		expectedValid bool
	}{
		{
			name:          "Valid passport number",
			passport:      "1234 567890",
			expectedValid: true,
		},
		{
			name:          "Invalid passport number",
			passport:      "1234567890",
			expectedValid: false,
		},
		{
			name:          "Invalid passport number (short)",
			passport:      "123 567890",
			expectedValid: false,
		},
		{
			name:          "Invalid passport number (long)",
			passport:      "12345 567890",
			expectedValid: false,
		},
		{
			name:          "Invalid passport number (symbol)",
			passport:      "1234 A67890",
			expectedValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := types.User{
				PassportNumber: tc.passport,
			}

			err := validate.Struct(user)
			if tc.expectedValid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}