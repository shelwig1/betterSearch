package main

import "fmt"

func Search(directoryMap []string, target string) []Result {
	// ch := make(chan Result)
	length := len(directoryMap)
	resultsChannel := make(chan Result, 0)

	var results []Result

	for _, file := range directoryMap {
		go HandleFile(target, file, resultsChannel)
	}

	for i := 0; i < length; i++ {
		value := <-resultsChannel

		if value.valid == true {
			results = append(results, value)
		}

	}

	fmt.Println("Results :", results)
	return results

}
