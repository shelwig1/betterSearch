package main

import (
	"testing"
)

const searchTestDir = "C:\\Users\\seanh\\Documents\\Projects\\betterSearch\\test_files"

func TestSearch(t *testing.T) {
	dirMap := GetDirectoryMap(searchTestDir)
	target := "Pineapple"

	Search(dirMap, target)
}
