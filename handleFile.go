package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// wg *sync.WaitGroup
// Add this back into arguments if we go back to WaitGroups

// Testing this is going to be annoying - we'll figure it out
func HandleFile(target string, file string, ch chan Result) {
	// defer wg.Done()

	// Not returning because we're pulling files we can't pull data from
	// Check the type of the file, if it's bad news return an error
	fmt.Println("File started processing: ", file)

	var result Result

	text, err := extractData(file)
	if err != nil {
		fmt.Println("Error extracting text from ", file, " : ", err)
		<-ch
		return
	}

	if initialCheck(target, text) {
		result.target = target
		result.file = file

		ch <- result
	} else {
		<-ch
	}

}

// We can offload any issues here -> return an error if we find something gay
func extractData(file string) (string, error) {
	// Check the file type
	ext := filepath.Ext(file)
	if ext != ".txt" {
		return "", errors.New("File is not an accepted format")
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	text := string(data)

	return text, nil
}

func initialCheck(target string, text string) bool {
	re := regexp.MustCompile(target)

	return re.MatchString(text)
}
