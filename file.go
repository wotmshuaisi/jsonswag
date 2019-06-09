package main

import (
	"bufio"
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

func (s *swagFile) ReadNext() string {
	data, _, err := s.ReadLine()
	if err != nil {
		panic(err)
	}
	// regex
	var d = hashTagRegex.FindAllStringSubmatch(string(data), 1)
	if len(d[0]) == 1 {
		panic("invalid format - line" + strconv.Itoa(s.Seek))
	}
	s.Seek += 1
	return d[0][1]
}

func (s *swagFile) GetTitle() {
	if s.bool {
		return
	}
	var data = s.ReadNext()
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
	var data = strings.Split(s.ReadNext(), " | ")
	if s.Result.Paths[data[1]] == nil {
		s.Result.Paths[data[1]] = map[string]Path{}
	}
	s.Result.Paths[data[1]][data[0]] = Path{
		Summary:    data[2],
		Parameters: []Parameter{},
		Responses:  map[string]Response{},
	}
	// definitions
	// Request
	// Response
	return ""
}
