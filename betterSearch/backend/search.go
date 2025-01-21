package backend

import "fmt"

type SearchStruct struct{}

func (s *SearchStruct) Search(target string) []Result {
	// ch := make(chan Result)
	currentDirectoryMap = GetCurrentDirectoryMap()
	length := len(currentDirectoryMap)
	// Consider killing the 0 at the end - don't know if it would still functions
	//	resultsChannel := make(chan Result, 0)
	resultsChannel := make(chan Result)

	var results []Result

	for _, file := range currentDirectoryMap {
		go HandleFile(target, file, resultsChannel)
	}

	for i := 0; i < length; i++ {
		value := <-resultsChannel

		// If we get any issues this is where it will come from
		if value.Valid {
			results = append(results, value)
		}

	}

	fmt.Println("Results :", results)
	return results

}
