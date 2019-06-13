package main

import (
	"encoding/json"
	"testing"
)

func Test_processJSON(t *testing.T) {
	var data = []byte(`{"str1":"string","num1":0,"num2":0,"str2":"string", "bool1": true, "object1": {"a": {"a": 1}, "b": 2}}`)
	d := processJSON(data)
	r, err := json.Marshal(d)
	if err != nil {
		t.Fatal(r)
	}
	data = []byte(`["1", "2"]`)
	d = processJSON(data)
	r, _ = json.Marshal(d)
	if err != nil {
		t.Fatal(r)
	}
}
