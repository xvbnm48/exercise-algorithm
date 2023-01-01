package main

import (
	"encoding/json"
	"fmt"
)

// Operation represents a single OT operation
type Operation struct {
	Type string `json:"type"`
	Pos  int    `json:"pos"`
	Val  string `json:"val"`
}

// ValidateOT takes in the stale file contents, the latest file contents, and a slice of OT operations as inputs
// and returns a boolean indicating whether the sequence of operations, when applied to the stale contents,
// produces the latest contents
func ValidateOT(stale, latest string, ops []Operation) bool {
	current := stale
	for _, op := range ops {
		if op.Type == "insert" {
			current = current[:op.Pos] + op.Val + current[op.Pos:]
		} else if op.Type == "delete" {
			current = current[:op.Pos] + current[op.Pos+len(op.Val):]
		} else {
			// invalid operation type
			return false
		}
	}

	return current == latest
}

func main() {
	stale := "The quick brown fox jumps over the lazy dog"
	latest := "The quick brown cat jumps over the lazy dog"
	opsJSON := `[{"type":"delete","pos":16,"val":"fox"},{"type":"insert","pos":16,"val":"cat"}]`

	var ops []Operation
	err := json.Unmarshal([]byte(opsJSON), &ops)
	if err != nil {
		fmt.Println(err)
		return
	}

	valid := ValidateOT(stale, latest, ops)
	fmt.Println(valid) // prints "true"
}
