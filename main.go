package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// title
	s := newSwagFile()
	s.GetTitle()
	s.GetPath()
	var d, _ = json.Marshal(s.Result)
	fmt.Println(string(d))
}
