package main

import (
	"fmt"
	"testing"
)

func TestGetDirectoryMap(t *testing.T) {
	result := GetDirectoryMap(projectDir + "test_files")

	for _, v := range result {
		fmt.Println(v)
	}

}

func TestGetDirectoryMap2(t *testing.T) {
	result := GetDirectoryMap("C:\\Program Files")

	for _, v := range result {
		fmt.Println(v)
	}

}

func TestOpenFileInOS(t *testing.T) {
	path := "C:\\Programming Projects\\betterSearch\\betterSearch\\test_files\\test.txt"

	OpenFileInOS(path)
}

func TestOpenFileExplorer(t *testing.T) {
	dir, err := ChooseDirectory()
	if err != nil {
		fmt.Println("Error retrieved: ", err)
	} else {
		fmt.Println("Directory: ", dir)
	}
}
