package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_processJSON(t *testing.T) {
	var data = []byte(`{"str1":"string","num1":0,"num2":0,"str2":"string", "bool1": true}`)
	d := processJSON(data)
	r, _ := json.Marshal(d)
	fmt.Println(string(r))
	data = []byte(`["1", "2"]`)
	d = processJSON(data)
	r, _ = json.Marshal(d)
	fmt.Println(string(r))
	t.Fatal()
}
