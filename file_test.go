package main

import (
	"testing"
)

func Test_processJSON(t *testing.T) {
	var data = []byte("[1,2,3]")
	processJSON(data)
}
