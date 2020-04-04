package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"copy/gocopy"
)

const usage = `
Hello, it's my program for copying files - gocopy. 
	This is flags for usage: 
		-from string - path to source. 
		-to string - path to destination. 
		-offset int - offset in file in bytes. 
		-limit int - copy limit of file.
`

var from, to string
var offset, limit int64

func main() {
	flag.StringVar(&from, "from", "", "from /path/to/source")
	flag.StringVar(&to, "to", "", "to /path/to/source")
	flag.Int64Var(&offset, "offset", 0, "offset 1024")
	flag.Int64Var(&limit, "limit", 0, "limit 2048")
	flag.Parse()


	if len(os.Args) == 1 {
		log.Fatal(usage)
	}

	bytes, err := gocopy.WriteFile(from, to, offset, limit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes sucessfully written.", bytes)
}
