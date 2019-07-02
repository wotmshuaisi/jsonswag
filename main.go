package main

import (
	"flag"
)

var (
	filepath    = flag.String("f", "", "*[required] file path")
	outputpath  = flag.String("o", "./swagger.json", "output file path")
	prettyprint = flag.Bool("p", false, "pretty print")
)

func main() {
	flag.Parse()
	if filepath == nil || *filepath == "" {
		flag.PrintDefaults()
		return
	}
	var swag = newSwagFile(*filepath, *outputpath)
	if eof := swag.GetTitle(); eof {
		return
	}
	for {
		if eof := swag.GetPath(); eof {
			break
		}
	}
	swag.SaveToPath(*prettyprint)
}
