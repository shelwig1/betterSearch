package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Recursive reading of directories
func GetDirectoryMap(dir string) []string {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking filepath: ", err)
	}

	return files
}

/*
func TraverseDirectoryMap(directoryMap []string) {
	var wg sync.WaitGroup

	for _, file := range directoryMap {
		wg.Add(1)
		go handleFile(file, &wg)
	}

	wg.Wait()
}

// Figure out channels I guess

// We can do the queuing thing for how much active memory we want to use so we don't bomb the fucking computer
// Channels may be the way to handle this one

// Channels -> send updated memory values between goroutines
// Mutexes -> protect shared memory when goroutines update the avaiable memory
func handleFile(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	ReadFile(path)
	//fmt.Println("File name: ", path)
}

// What does reading from a text file return?
// Need to check if we can load the whole fucker into memory
// How to clear it out?
// May want to stream this fucker but I'm worried about missing a match along a chunk
func ReadFile(dir string) {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	data = nil
}
*/
