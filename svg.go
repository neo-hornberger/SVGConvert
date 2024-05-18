package main

import (
	"fmt"
	"io"
	"os"

	"github.com/JoshVarga/svgparser"
)

func ParseSVG(file string) *svgparser.Element {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting current working directory")
			cwd = DIR_STACK.Peek().(string)
		}

		fmt.Fprintf(os.Stderr, "File %s does not exist in dir %s\n", file, cwd)
		os.Exit(1)
	}

	reader, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s\n", file)
		os.Exit(1)
	}

	elem, err := svgparser.Parse(reader, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing SVG")
		os.Exit(1)
	}

	return elem
}

func WriteSVG(w io.Writer, elem *svgparser.Element) {
	io.WriteString(w, "<"+elem.Name)
	for key, val := range elem.Attributes {
		io.WriteString(w, " "+key+"=\""+val+"\"")
	}
	io.WriteString(w, ">")

	if elem.Content != "" {
		io.WriteString(w, elem.Content)
	}
	for _, child := range elem.Children {
		WriteSVG(w, child)
	}

	io.WriteString(w, "</"+elem.Name+">")
}
