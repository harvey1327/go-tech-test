package reader_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/harvey1327/emaildomainstats/internal/reader"

	"github.com/stretchr/testify/assert"
)

func TestReadFileReturnsError(t *testing.T) {
	t.Run("readFile returns an error if path is incorrect", func(t *testing.T) {
		_, err := reader.ReadFile("non-existantFile", 0)
		assert.NotNil(t, err)
	})
}

func TestReadFileReturnsNonEmptyChannel(t *testing.T) {

	tests := []struct {
		name          string
		content       []string
		skipLines     int
		contentLength int
	}{
		{"readFile returns a populated channel of items", []string{"one,two,three", "four,five,six"}, 0, 2},
		{"readFile returns a populated channel of items", []string{"ignore,this,line", "one,two,three"}, 1, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fileName := "testFile.csv"
			writeFile(fileName, test.content)

			actualResult, err := reader.ReadFile(fileName, test.skipLines)
			assert.Nil(t, err)
			count := 0
			for data := range actualResult {
				assert.NotNil(t, data)
				count++
			}
			assert.Equal(t, test.contentLength, count)

			deleteFile(fileName)
		})
	}
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
