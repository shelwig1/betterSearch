package main

import (
	"testing"
)

// const searchTestDir = "C:\\Users\\seanh\\Documents\\Projects\\betterSearch\\test_files"
const searchTestDir = projectDir + "test_files"

func TestSearch(t *testing.T) {
	dirMap := getDirectoryMap(searchTestDir)
	target := "Pineapple"

	Search(dirMap, target)
}
