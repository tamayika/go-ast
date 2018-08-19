package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tamayika/go-ast/pkg/ast"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	n, err := ast.Parse(readSrc())
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(b))
}

func readSrc() string {
	if flag.NArg() == 0 {
		src, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		return string(src)
	}
	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		switch _, err := os.Stat(path); {
		case err != nil:
			log.Fatal(err)
		default:
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			return string(buf)
		}
	}
	return ""
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: go-ast [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}
