package markdowntohtml

import (
	"github.com/Kunal-deve1oper/markdowntohtml/pkg/fileparse"
	"github.com/Kunal-deve1oper/markdowntohtml/pkg/mdparser"
)

// Function takes path to the markdownfile as argument
func Convert(filename string) error {
	v := &fileparse.FileContent{}
	err := v.ParseMd(filename)
	if err != nil {
		return err
	}
	t := &mdparser.MdToHtml{}
	err = t.HtmlParser(v.Content)
	if err != nil {
		return err
	}
	return nil
}
