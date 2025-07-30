package presentation

import (
	"hands_on_go/internal/statuserr"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidatorImpl_ValidateCreateUser_MissingLastName_ReturnError(t *testing.T) {
	// Init
	testCases := []struct {
		name          string
		body          string
		expectedError string
	}{
		{
			name: "missing last name",
			body: `
			{
				"firstName": "Joe",
				"age": 25,
				"phone": "Joe",
				"phoneVerified": true
			}
			`,
			expectedError: "required field 'LastName' is null",
		},
		{
			name: "missing first name",
			body: `
			{
				"LastName": "Joe",
				"age": 25,
				"phone": "Joe",
				"phoneVerified": true
			}
			`,
			expectedError: "required field 'FirstName' is null",
		},
	}

	for _, params := range testCases {
		t.Run(params.name, func(t *testing.T) {
			jsonBody := params.body
			reader := io.NopCloser(strings.NewReader(jsonBody))
			request := http.Request{Body: reader}

			validator := NewUserValidatorImpl()

			// Execute
			_, err := validator.ValidateCreateUser(&request)

			// Asserts
			assert.NotNil(t, err)
			assert.Equal(t, statuserr.KindInvalidRequest, statuserr.GetKind(err))

			msg, ok := statuserr.GetMessage(err)
			assert.True(t, ok)
			assert.Equal(t, params.expectedError, msg)
		})
	}

}
