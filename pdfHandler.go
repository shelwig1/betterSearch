package main

import (
	"fmt"

	"github.com/dslipak/pdf"
)

// Would be beneficial to send the result with a page number, no? -> at that point you can ctrl+f it

// New concern: loading a whole pdf into memory might suck. Plain text is cheap though, so it may not be a problem
// Let's go until memory becomes an issue, then figure it out
func readPDF(path string) (string, error) {
	reader, err := pdf.Open(path)
	if err != nil {
		fmt.Println("Error!!!: ", err)
		return "", err
	}

	totalPage := reader.NumPage()

	var result string

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := reader.Page(pageIndex)
		if p.V.IsNull() {
			fmt.Println("Page was null")
			continue
		}

		// var lastTextStyle pdf.Text
		texts := p.Content().Text
		for _, text := range texts {
			result += text.S
		}

	}
	fmt.Println(result)

	return "", nil
}
