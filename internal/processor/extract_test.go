package processor_test

import (
	"errors"
	"testing"

	"github.com/harvey1327/emaildomainstats/internal/processor"

	"github.com/stretchr/testify/assert"
)

func TestExtractEmailDomain(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{"ExtractEmailDomain returns a domain", "first_name,last_name,email@domain.com,gender,ip_address", "domain.com", nil},
		{"ExtractEmailDomain returns error with incomplete input", "last_name,email@domain.com,gender,ip_address", "", errors.New("length of data line does not conform to standard requirement of 5 fields: last_name,email@domain.com,gender,ip_address")},
		{"ExtractEmailDomain returns error with malformed email", "first_name,last_name,email_domain.com,gender,ip_address", "", errors.New("email is malformed: first_name,last_name,email_domain.com,gender,ip_address")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := processor.ExtractEmailDomain(test.input)
			if test.err == nil {
				assert.Equal(t, test.expected, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, test.err, err)
			}
		})
	}
}
