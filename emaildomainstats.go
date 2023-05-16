/*
This package is required to provide functionality to process a csv file and return a sorted (by email domain) data
structure of your choice containing the email domains along with the number of customers for each domain. The customer_data.csv
file provides an example csv file to work with. Any errors should be logged (or handled) or returned to the consumer of
this package. Performance matters, the sample file may only contain 1K lines but the package may be expected to be used on
files with 10 million lines or run on a small machine.

Write this package as you normally would for any production grade code that would be deployed to a live system.

Please stick to using the standard library.
*/

package emaildomainstats

import (
	"sync"

	"github.com/harvey1327/emaildomainstats/internal/datastore"
	"github.com/harvey1327/emaildomainstats/internal/processor"
	"github.com/harvey1327/emaildomainstats/internal/reader"
)

// GenerateOuput will egenrate the required ouput which contains the required data structure and any processing errors
// will return an error if there are issues reading the file
func GenerateOuput(filePath string) (*datastore.EmailDomain, error) {
	lines, err := reader.ReadFile(filePath, 1)
	if err != nil {
		return nil, err
	}

	emailDomains := datastore.EmailDomainStore()

	var wg sync.WaitGroup
	for {
		line, ok := <-lines
		if !ok {
			break
		}
		wg.Add(1)
		go updateDatastore(&wg, line, emailDomains)
	}
	wg.Wait()
	return emailDomains, nil
}

func updateDatastore(wg *sync.WaitGroup, line string, emailDomains *datastore.EmailDomain) {
	domain, err := processor.ExtractEmailDomain(line)
	if err != nil {
		emailDomains.AddError(err)
	} else {
		emailDomains.IncrementOrAddKey(domain)
	}
	wg.Done()
}
