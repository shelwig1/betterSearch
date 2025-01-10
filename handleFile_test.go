package main

import (
	"sync"
	"testing"
)

const testDir = "C:\\Users\\seanh\\Documents\\Projects\\betterSearch\\test_files\\dirTest\\testFile.txt"

func TestExtractData(t *testing.T) {
	answer := "Here is an example of my text. Pineapple."
	result := extractData(testDir)

	if result != answer {
		t.Errorf("Expected answer, got %s", result)
	}
}

func TestInitialCheck(t *testing.T) {
	target := "Pineapple"
	text := "Here is an example of my text. Pineapple."

	result := initialCheck(target, text)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

func TestHandleFile(t *testing.T) {
	var wg sync.WaitGroup

	result := HandleFile("Pineapple", testDir, &wg)

	if !result {
		t.Errorf("HandleFile failed")
	}
}
