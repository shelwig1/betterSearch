package main

type Result struct {
	valid  bool
	target string
	file   string
}

// Do I even need the target here?
// I'm just returning files that match it
