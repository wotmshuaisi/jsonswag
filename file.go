package main

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	hashTagRegex = regexp.MustCompile("^#+ (.*)")
)

type swagFile struct {
	*bufio.Reader
	*bufio.Writer
	Result *Swag
	Seek   int
	bool
}

// Methods

func newSwagFile() *swagFile {
	var input, err = os.OpenFile("./plain.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic("error occured while opening")
	}
	output, err := os.OpenFile("./result.json", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic("error occured while opening")
	}
	var f = &swagFile{
		bufio.NewReader(input),
		bufio.NewWriter(output),
		&Swag{
			Paths:       map[string]map[string]Path{},
			Definitions: map[string]interface{}{},
		},
		0,
		false,
	}
	f.Result.Version = "2.0"
	return f
}

func (s *swagFile) ReadNext(b bool) (string, []byte) {
	data, _, err := s.ReadLine()
	if err != nil {
		panic(err)
	}
	// regex
	if !b {
		var d = hashTagRegex.FindAllStringSubmatch(string(data), 1)
		if len(d[0]) == 1 {
			panic("invalid format - line" + strconv.Itoa(s.Seek))
		}
		s.Seek++
		return d[0][1], nil
	}
	var d = hashTagRegex.FindAllSubmatch(data, 1)
	if len(d[0]) == 1 {
		panic("invalid format - line" + strconv.Itoa(s.Seek))
	}
	s.Seek++
	return "", d[0][1]
}

func (s *swagFile) GetTitle() {
	if s.bool {
		return
	}
	var data, _ = s.ReadNext(false)
	// title processing
	var title = strings.Split(data, " - ")
	s.Result.Info = map[string]string{}
	switch len(title) {
	case 3:
		s.Result.Info["description"] = title[1]
		fallthrough
	case 2:
		s.Result.Info["version"] = title[len(title)-1]
		fallthrough
	case 1:
		s.Result.Info["title"] = title[0]
	}
	s.bool = true
}

// GetPath write a path into Swag return path
func (s *swagFile) GetPath() string {
	if s.bool == false {
		panic("you forgot get title")
	}
	// Method URL
	var data, _ = s.ReadNext(false)
	var d = strings.Split(data, " | ")
	var method, uri, summary = d[0], d[1], d[2]
	if s.Result.Paths[uri] == nil {
		s.Result.Paths[uri] = map[string]Path{}
	}
	s.Result.Paths[uri][method] = Path{
		Summary:    summary,
		Parameters: []Parameter{},
		Responses:  map[string]Response{},
	}
	// Request
	var _, request = s.ReadNext(true)
	processJSON(request)
	// Response
	// definitions
	return ""
}

// public functions

func processJSON(j []byte) Definition {
	// return  definition
	var result = Definition{}
	var d = map[string]interface{}{}
	var s = []interface{}{}
	var object = true
	if err := json.Unmarshal(j, &d); err != nil {
		if err := json.Unmarshal(j, &s); err != nil {
			panic(err)
		}
		object = false
	}
	// object
	if object {
		if len(d) == 0 {
			return Definition{}
		}
		result.Type = Object
		var properties = map[string]Definition{}
		for k, v := range d {
			properties[k] = Definition{
				Type: typeDetection(v),
			}
		}
		result.Properties = properties
		return result
	}
	// array
	if len(s) == 0 {
		return Definition{}
	}
	result.Type = "array"
	result.Items = Definition{
		Type: typeDetection(s[0]),
	}
	return result
}
