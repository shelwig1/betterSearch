package main

import (
	"testing"
)

const testDir = projectDir + "test_files\\dirTest\\testFile.txt"
const higherDir = projectDir + "test_files"

func TestExtractData(t *testing.T) {
	answer := "Here is an example of my text. Pineapple."
	result, _ := extractData(testDir)

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

// func TestHandleFile(t *testing.T) {
// 	var wg sync.WaitGroup

// 	result := HandleFile("Pineapple", testDir, &wg)

//		if !result {
//			t.Errorf("HandleFile failed")
//		}
//	}
func TestHandleFile(t *testing.T) {
	ch := make(chan Result)

	go HandleFile("Pineapple", testDir, ch)

	result := <-ch

	if result.file != testDir || result.target != "Pineapple" {
		t.Errorf("Expected result with:\ntarget: \"Pineapple\"\nfile : file ending with testFile.txt")
	}

}
