package main

import (
	"fmt"
	"os"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {

	}

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(os.Stderr, "findlinks:: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc){
		fmt.Println(link)
	}
}