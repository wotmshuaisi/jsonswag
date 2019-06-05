package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	hashTagRegex = regexp.MustCompile("#+ (.*)")
)

type swagFile struct {
	*bufio.Reader
}

func newSwagFile(path string) *swagFile {
	var file, err = os.Open("./plain.txt")
	if err != nil {
		panic("error occured while opening")
	}
	return &swagFile{bufio.NewReader(file)}
}

func (s *swagFile) getTitle() (string, bool) {
	var data, _, err = s.Reader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return "", false
		}
		panic("readline: " + err.Error())
	}

	return string(data), true
}

func main() {

	// title
	s := newSwagFile("./plain.txt")
	if data, ok := s.getTitle(); ok {
		fmt.Println(data)
	}
	// content
	// for {
	// 	var data, _, err = fileR.ReadLine()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		panic("readline: " + err.Error())
	// 	}
	// 	fmt.Println(string(data))
	// }
}
