package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/parallellink/srg"
)

func main() {
	expression := flag.String("e", "", "range expression. \"-e=a,b1~b3,c\"")
	delimiter := flag.String("d", " ", "delimiter character. \"-d='\\n'\"")
	surround := flag.String("s", "", "surround character. \"-s=' '\" ")
	flag.Parse()

	if *expression == "" {
		fmt.Println("range expression is empty")
		os.Exit(1)
	}

	if *delimiter == "\\n" {
		*delimiter = "\n"
	}

	hosts, err := srg.ParseRange(*expression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buffer bytes.Buffer
	first := true
	for _, h := range hosts {
		if first {
			first = false
		} else {
			buffer.WriteString(*delimiter)
		}
		buffer.WriteString(*surround)
		buffer.WriteString(h)
		buffer.WriteString(*surround)
	}

	fmt.Println(buffer.String())
}
