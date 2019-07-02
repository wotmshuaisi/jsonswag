package main

import (
	"encoding/json"
	"flag"
	"os"
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
	var swag = newSwagFile(*filepath)
	if eof := swag.GetTitle(); eof {
		return
	}
	for {
		if eof := swag.GetPath(); eof {
			break
		}
	}
	var d []byte
	if *prettyprint {
		d, _ = json.MarshalIndent(swag.Result, "", "  ")
	} else {
		d, _ = json.Marshal(swag.Result)
	}
	var output, err = os.OpenFile(*outputpath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	output.Write(d)
	output.Close()
}
