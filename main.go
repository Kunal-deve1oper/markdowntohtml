package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Kunal-deve1oper/md-to-html/pkg/fileparse"
	"github.com/Kunal-deve1oper/md-to-html/pkg/mdparser"
)

func main() {
	filepath := flag.String("f", "", "path to the markdown file")
	flag.Parse()
	if *filepath == "" {
		fmt.Println("Error: Please provide the path to the markdown file using -f flag.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	v := &fileparse.FileContent{}
	err := v.ParseMd(*filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t := &mdparser.MdToHtml{}
	err = t.HtmlParser(v.Content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("HTML conversion successful!")
}
