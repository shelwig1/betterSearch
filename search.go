package main

import "fmt"

func Search(directoryMap []string, target string) []Result {
	// ch := make(chan Result)
	length := len(directoryMap)
	ch := make(chan Result, 0)

	var results []Result

	for _, file := range directoryMap {
		go HandleFile(target, file, ch)
	}

	// for result := range ch {
	// 	results = append(results, result)
	// }

	for i := 0; i < length; i++ {
		value := <-ch
		fmt.Println("Search function received: ", value)
		// results = append(results, <-ch)
		results = append(results, value)

	}

	fmt.Println("Results :", results)
	return results

}
