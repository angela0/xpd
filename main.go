package main

import (
	"fmt"
	"os"

	"github.com/angela0/xpd/query"
)

func main() {
	var file *os.File
	var err error

	switch len(os.Args) {
	case 2:
		file = os.Stdin
	case 3:
		file, err = os.Open(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	default:
		fmt.Fprintln(os.Stderr, "Usage: xpd <selector> [filename]")
		return
	}

	doc, err := query.Parse(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	nodes, err := query.QueryAll(doc, os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, node := range nodes {
		fmt.Println(query.OutputHTML(node, true))
	}
}
