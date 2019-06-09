package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// title
	s := newSwagFile()
	s.getTitle()
	var d, _ = json.Marshal(s.Result)
	fmt.Println(string(d))
}
