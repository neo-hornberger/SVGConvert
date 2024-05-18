package main

import (
	"fmt"
	"os"

	"github.com/JoshVarga/svgparser"
)

var ATTRIBUTES []string = []string{
	"x",
	"y",
	"width",
	"height",
}

func handleSVG(elem *svgparser.Element) {
	for i, child := range elem.Children {
		switch child.Name {
		case "use":
			if href, ok := child.Attributes["href"]; ok {
				if href[0] == '#' {
					break
				}

				extern_elem := ParseSVG(href)

				Pushd(RelDir(href))
				handleSVG(extern_elem)
				Popd()

				for _, attr := range ATTRIBUTES {
					if val, ok := child.Attributes[attr]; ok {
						extern_elem.Attributes[attr] = val
					}
				}
				extern_elem.Attributes["data-clone-of"] = href

				elem.Children[i] = extern_elem
			}
		}

		if len(child.Children) > 0 {
			handleSVG(child)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a file path")
		os.Exit(1)
	}

	file := os.Args[1]

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting current working directory")
		os.Exit(1)
	}

	elem := ParseSVG(file)

	Pushd(cwd)
	Pushd(RelDir(file))
	handleSVG(elem)
	Popd()

	if DIR_STACK.Len() > 1 && Popd() != cwd {
		fmt.Fprintln(os.Stderr, "Error: Directory stack corrupted")
		os.Exit(1)
	}

	WriteSVG(os.Stdout, elem)
}
