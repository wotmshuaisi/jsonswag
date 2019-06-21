package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// title
	s := newSwagFile("./plain.txt")
	s.GetTitle()
	s.GetPath()
	s.GetPath()
	s.GetPath()
	s.GetPath()
	var d, _ = json.Marshal(s.Result)
	// var d, _ = json.MarshalIndent(s.Result, "", "  ")
	fmt.Println(string(d))
}
