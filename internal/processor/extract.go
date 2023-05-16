package processor

import (
	"fmt"
	"strings"
)

// ExtractEmailDomain takes in a single csv line and extracts out the domian from the email
func ExtractEmailDomain(input string) (string, error) {
	fields := strings.Split(input, ",")
	if len(fields) != 5 {
		return "", fmt.Errorf("length of data line does not conform to standard requirement of 5 fields: %s", input)
	}
	email := fields[2]
	if !strings.Contains(email, "@") {
		return "", fmt.Errorf("email is malformed: %s", input)
	}
	return strings.Split(email, "@")[1], nil
}
