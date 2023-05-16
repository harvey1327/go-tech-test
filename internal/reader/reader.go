package reader

import (
	"bufio"
	"os"
	"sync"
)

// ReadFile takes in a path to the csv file to read, skipLine is used to skip the first N lines of the file
func ReadFile(path string, skipLine int) (<-chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	outputChan := make(chan string)

	for i := 0; i < skipLine; i++ {
		fileScanner.Scan()
	}

	var wg sync.WaitGroup
	for fileScanner.Scan() {
		wg.Add(1)
		go sendToChannel(&wg, outputChan, fileScanner.Text())
	}
	go cleanUp(&wg, outputChan)
	return outputChan, nil
}

func sendToChannel(wg *sync.WaitGroup, output chan string, message string) {
	output <- message
	wg.Done()
}

func cleanUp(wg *sync.WaitGroup, output chan string) {
	wg.Wait()
	close(output)
}
