package main

import (
	"os"
	"regexp"
	"sync"
)

func HandleFile(target string, file string, wg *sync.WaitGroup) bool {
	defer wg.Done()

	text := extractData(file)
	return initialCheck(target, text)
}

func extractData(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	text := string(data)

	return text
}

func initialCheck(target string, text string) bool {
	re := regexp.MustCompile(target)

	return re.MatchString(text)
}
