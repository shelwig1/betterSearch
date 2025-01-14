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
