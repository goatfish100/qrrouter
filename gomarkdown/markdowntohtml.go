package main

import (
	"fmt"
	"io/ioutil"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	// read the whole file at once
	b, err := ioutil.ReadFile("markdown.md")
	if err != nil {
		panic(err)
	}

	html := bluemonday.UGCPolicy().SanitizeBytes(b)

	output := blackfriday.MarkdownCommon(html)
	fmt.Print(string(output))

}
