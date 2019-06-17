package main

import (
	"bufio"
	"encoding/json"
	"io"
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

func newSwagFile(path string) *swagFile {
	var input, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic("error occured while opening")
	}
	if (fileRows(input)-1)%3 != 0 {
		panic("invalid file(lines).")
	}
	input.Seek(0, 0) // seek to file top
	// output file
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
func (s *swagFile) GetPath() {
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
		Summary: summary,
		Parameters: []Parameter{Parameter{
			Name:     "body",
			In:       "body",
			Required: true,
			Schema:   map[string]string{"$ref": "#/definitions/"},
		}},
		Responses: map[string]Response{"200": Response{
			Description: "OK",
			Schema:      map[string]string{"$ref": "#/definitions/"},
		}},
	}
	// Request
	var reqID = strings.ToUpper(method) + strings.ReplaceAll(uri, "/", "_") + "__Request"
	s.Result.Paths[uri][method].Parameters[0].Schema["$ref"] = "#/definitions/" + reqID // set definition
	var _, request = s.ReadNext(true)
	var reqDef = processJSON(request)
	s.Result.Definitions[reqID] = reqDef
	// Response
	var resID = strings.ToUpper(method) + strings.ReplaceAll(uri, "/", "_") + "__Response"
	s.Result.Paths[uri][method].Responses["200"].Schema["$ref"] = "#/definitions/" + resID // set definition
	var _, response = s.ReadNext(true)
	var resDef = processJSON(response)
	s.Result.Definitions[resID] = resDef
	// definitions
}

// public functions

func processJSON(j []byte) *Definition {
	// return  definition
	var result = &Definition{}
	var d = map[string]interface{}{}
	var s = []interface{}{}
	var object = true
	if len(j) == 0 {
		return nil
	}
	if err := json.Unmarshal(j, &d); err != nil {
		if err := json.Unmarshal(j, &s); err != nil {
			panic(err)
		}
		object = false
	}
	// object
	if object {
		if len(d) == 0 {
			return &Definition{}
		}
		result.Type = Object
		var properties = map[string]*Definition{}
		for k, v := range d {
			var childData, _ = json.Marshal(v)
			properties[k] = &Definition{
				Type: typeDetection(v),
			}
			switch properties[k].Type {
			case Object:
				var childMap = map[string]interface{}{}
				json.Unmarshal(childData, &childMap)
				properties[k].Properties = map[string]*Definition{}
				for kk, vv := range childMap {
					var t = typeDetection(vv)
					if t != Object {
						properties[k].Properties[kk] = &Definition{
							Type: t,
						}
						continue
					}
					var mapData, _ = json.Marshal(vv)
					properties[k].Properties[kk] = processJSON(mapData)
				}
			case Array:
				var childArray = []interface{}{}
				json.Unmarshal(childData, &childArray)
				var arrayData, _ = json.Marshal(childArray[0])
				switch typeDetection(childArray[0]) {
				case Object, Array:
					properties[k].Items = processJSON(arrayData)
				default:
					properties[k].Items = &Definition{
						Type: typeDetection(childArray[0]),
					}
				}
			}
		}
		result.Properties = properties
		return result
	}
	// array
	if len(s) == 0 {
		return &Definition{}
	}
	result.Type = "array"
	result.Items = &Definition{
		Type:       typeDetection(s[0]),
		Properties: map[string]*Definition{},
	}
	var childData, _ = json.Marshal(s[0]) // get first element
	var childMap = map[string]interface{}{}
	json.Unmarshal(childData, &childMap)
	for kk, vv := range childMap {
		var t = typeDetection(vv)
		if t != Object {
			result.Items.Properties[kk] = &Definition{
				Type: t,
			}
			continue
		}
		var mapData, _ = json.Marshal(vv)
		result.Items.Properties[kk] = processJSON(mapData)
	}
	return result
}

func fileRows(r io.Reader) int {
	var rb = bufio.NewReader(r)
	var count = 0
	for {
		_, _, err := rb.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		count++
	}
	return count
}
