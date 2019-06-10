package main

import (
	"encoding/json"
	"testing"
)

func Test_typeDetection(t *testing.T) {
	var testVar interface{}
	testVar = 1
	if typeDetection(testVar) != "number" {
		t.Fatal()
	}
	testVar = 1.1
	if typeDetection(testVar) != "number" {
		t.Fatal()
	}
	testVar = "1"
	if typeDetection(testVar) != "string" {
		t.Fatal()
	}
	testVar = true
	if typeDetection(testVar) != "boolean" {
		t.Fatal()
	}

	var j = []byte(`{"test": {"test1": 1}}`)
	var dd = map[string]interface{}{}
	json.Unmarshal(j, &dd)
	testVar = dd["test"]
	if typeDetection(testVar) != "object" {
		t.Fatal()
	}

	j = []byte(`{"test": ["test1", "1"]}`)
	dd = map[string]interface{}{}
	json.Unmarshal(j, &dd)
	testVar = dd["test"]
	if typeDetection(testVar) != "array" {
		t.Fatal()
	}
}
