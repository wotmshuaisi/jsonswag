package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var (
	hashTagRegex = regexp.MustCompile("^#+ (.*)")
)

type swagFile struct {
	*bufio.Reader
	*bufio.Writer
	Result *Swag
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
		panic("invalid format")
	}
	return d[0][1]
}

func (s *swagFile) getTitle() {
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
