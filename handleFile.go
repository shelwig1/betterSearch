package main

import (
	"os"
	"regexp"
)

// wg *sync.WaitGroup
// Add this back into arguments if we go back to WaitGroups

// Testing this is going to be annoying - we'll figure it out
func HandleFile(target string, file string, ch chan Result) {
	// defer wg.Done()

	var result Result

	text := extractData(file)

	// return initialCheck(target, text)

	if initialCheck(target, text) {
		result.target = target
		result.file = file

		ch <- result
	} else {
		<-ch
	}

	// Figure out why the fuck we need to do this
	close(ch)
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
