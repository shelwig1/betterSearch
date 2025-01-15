package main

import "testing"

func TestReadPDF(t *testing.T) {
	path := projectDir + "test_files\\clone.pdf"

	readPDF(path)
}
