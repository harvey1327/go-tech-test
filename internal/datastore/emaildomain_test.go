package datastore_test

import (
	"errors"
	"sort"
	"testing"

	"github.com/harvey1327/emaildomainstats/internal/datastore"

	"github.com/stretchr/testify/assert"
)

func TestIncrementKey(t *testing.T) {

	tests := []struct {
		name          string
		key           string
		expectedValue int
	}{
		{"IncrementOrAddKey increments the key by 1 for a non existing key", "key1", 1},
		{"IncrementOrAddKey increments the key by 1 for an existing key", "key2", 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			eds := datastore.EmailDomainStore()

			for i := 0; i < test.expectedValue; i++ {
				eds.IncrementOrAddKey(test.key)
			}

			for _, v := range eds.GetDomainsWithCounts() {
				assert.Equal(t, test.expectedValue, v.GetValue())
			}
		})
	}
}

func TestSortedOrder(t *testing.T) {
	t.Run("IncrementOrAddKey adds a new key to an already populated datastore", func(t *testing.T) {
		eds := datastore.EmailDomainStore()
		eds.IncrementOrAddKey("d")
		eds.IncrementOrAddKey("e")
		eds.IncrementOrAddKey("b")
		eds.IncrementOrAddKey("a")
		eds.IncrementOrAddKey("c")

		items := eds.GetDomainsWithCounts()

		assert.Equal(t, 5, len(items))
		assert.True(t, sort.SliceIsSorted(items, func(i, j int) bool { return len(items[i].GetKey()) > len(items[j].GetKey()) }))
	})
}

func TestAddError(t *testing.T) {
	t.Run("AddError adds an error", func(t *testing.T) {
		eds := datastore.EmailDomainStore()
		eds.AddError(errors.New("some error"))
		assert.Equal(t, 1, len(eds.GetErrors()))
	})
}
