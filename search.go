package main

import "sync"

func Search(directoryMap []string, target string) {
	var wg sync.WaitGroup

	for _, file := range directoryMap {
		wg.Add(1)
		go HandleFile(target, file, &wg)
	}
}
