package emaildomainstats_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/harvey1327/emaildomainstats"
	"github.com/stretchr/testify/assert"
)

func TestGenerateOuputReturnsDataStructure(t *testing.T) {
	tests := []struct {
		name               string
		fileContent        []string
		expectedDomains    []string
		expectedCounts     []int
		expectedErrorCount int
	}{
		{"GenerateOuput returns data structure with no errors", []string{"first_name,last_name,email,gender,ip_address", "Rowland,Aldins,raldins0@google.com,Female,46.148.249.225"}, []string{"google.com"}, []int{1}, 0},
		{"GenerateOuput returns data structure with errors", []string{"first_name,last_name,email,gender,ip_address", "Rowland,Aldins,raldins0@google.com,Female,46.148.249.225", "some incorrect data"}, []string{"google.com"}, []int{1}, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			file := "test.csv"
			writeFile(file, test.fileContent)
			ds, err := emaildomainstats.GenerateOuput(file)

			domains := ds.GetDomainsWithCounts()
			errors := ds.GetErrors()

			assert.Nil(t, err)
			assert.Equal(t, len(test.expectedDomains), len(domains))
			assert.Equal(t, len(test.expectedCounts), len(domains))
			assert.Equal(t, test.expectedErrorCount, len(errors))

			for idx, expectedDomain := range test.expectedDomains {
				assert.Equal(t, expectedDomain, domains[idx].GetKey())
				assert.Equal(t, test.expectedCounts[idx], domains[idx].GetValue())
			}

			deleteFile(file)
		})
	}
}

func TestGenerateOuputReturnsError(t *testing.T) {
	t.Run("GenerateOuput returns error", func(t *testing.T) {
		_, err := emaildomainstats.GenerateOuput("non-existing-file.csv")
		assert.NotNil(t, err)
	})
}

func writeFile(path string, content []string) {
	err := os.WriteFile(path, []byte(strings.Join(content, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}
