package backend

type Result struct {
	Valid  bool   `json:"valid"`
	Target string `json:"target"`
	File   string `json:"file"`
}

// Do I even need the target here?
// I'm just returning files that match it
