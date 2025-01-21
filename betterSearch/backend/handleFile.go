package backend

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Current implementation means we can't just send an empty struct because it's not of type Result
// So we just set the fucker with a flag that immediately lets it be discarded as needed
func HandleFile(target string, file string, ch chan Result) {
	fmt.Println("File started processing: ", file)

	var result Result
	result.Target = target
	result.File = file

	text, err := extractData(file)
	if err != nil {
		fmt.Println("Error extracting text from ", file, " : ", err)

		result.Valid = false
		ch <- result

		return
	}

	if initialCheck(target, text) {
		result.Valid = true
	} else {
		result.Valid = false
	}

	ch <- result

}

// We can offload any issues here -> return an error if we find something gay
// Next step is making it so we handle different file types
func extractData(file string) (string, error) {
	// Check the file type
	ext := filepath.Ext(file)
	if ext != ".txt" {
		return "", errors.New("file is not an accepted format")
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
