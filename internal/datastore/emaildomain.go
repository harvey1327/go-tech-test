package datastore

import (
	"sort"
	"sync"
)

type item struct {
	key   string
	value int
}

func (itm item) GetKey() string {
	return itm.key
}

func (itm item) GetValue() int {
	return itm.value
}

type EmailDomain struct {
	itemIndexMap map[string]int
	orderdItems  []item
	errors       []error
	mu           *sync.RWMutex
}

// EmailDomainStore holds the domains with their count within orderdItems
// itemIndexMap holds the domain against it's current index within the slice
// errors holds any errors due to processing
func EmailDomainStore() *EmailDomain {
	return &EmailDomain{
		itemIndexMap: make(map[string]int),
		orderdItems:  make([]item, 0),
		errors:       make([]error, 0),
		mu:           &sync.RWMutex{},
	}
}

func (eds *EmailDomain) IncrementOrAddKey(key string) {
	eds.mu.Lock()
	if itemIndex, ok := eds.itemIndexMap[key]; ok {
		eds.orderdItems[itemIndex] = item{key: key, value: eds.orderdItems[itemIndex].value + 1}
	} else {
		if len(eds.orderdItems) == 0 {
			eds.orderdItems = append(eds.orderdItems, item{key: key, value: 1})
		} else {
			index := sort.Search(len(eds.orderdItems), func(i int) bool { return eds.orderdItems[i].key >= key })
			eds.orderdItems = append(eds.orderdItems, item{})
			copy(eds.orderdItems[index+1:], eds.orderdItems[index:])
			eds.orderdItems[index] = item{key: key, value: 1}
		}
		//update indexMap
		for idx, itm := range eds.orderdItems {
			eds.itemIndexMap[itm.key] = idx
		}
	}
	eds.mu.Unlock()
}

func (eds *EmailDomain) AddError(err error) {
	eds.mu.Lock()
	eds.errors = append(eds.errors, err)
	eds.mu.Unlock()
}

func (eds *EmailDomain) GetErrors() []error {
	return eds.errors
}

func (eds *EmailDomain) GetDomainsWithCounts() []item {
	return eds.orderdItems
}
