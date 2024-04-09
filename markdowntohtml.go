package markdowntohtml

import (
	"os"

	"github.com/Kunal-deve1oper/markdowntohtml/pkg/fileparse"
	"github.com/Kunal-deve1oper/markdowntohtml/pkg/mdparser"
)

/*
Function takes path to the markdownfile as argument
and generate an index.html file
*/
func ConvertToFile(filename string) error {
	v := &fileparse.FileContent{}
	err := v.ParseMd(filename)
	if err != nil {
		return err
	}
	t := &mdparser.MdToHtml{}
	temp := `<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>HTML 5 Boilerplate</title>
		<link rel="stylesheet" href="style.css">
	  </head>
	  <body>`
	res, err := t.HtmlParser(v.Content)
	if err != nil {
		return err
	}
	res = temp + res + `</body></html>`
	err = os.WriteFile("./index.html", []byte(res), 0644)
	if err != nil {
		return err
	}
	return nil
}

/*
Function takes path to the markdownfile as argument
and give output of the HTML code as a string
*/
func ConvertToString(filename string) (string, error) {
	v := &fileparse.FileContent{}
	err := v.ParseMd(filename)
	if err != nil {
		return "", err
	}
	t := &mdparser.MdToHtml{}
	temp := `<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>HTML 5 Boilerplate</title>
		<link rel="stylesheet" href="style.css">
	  </head>
	  <body>`
	res, err := t.HtmlParser(v.Content)
	if err != nil {
		return "", err
	}
	res = temp + res + `</body></html>`
	return res, nil
}

/*
ConvertInputToString takes a slice of strings where each element represents a line of Markdown content.
The function will process each line of this input slice and convert it into HTML code.
*/
func ConvertInputToString(input []string) (string, error) {
	t := &mdparser.MdToHtml{}
	res, err := t.HtmlParser(input)
	if err != nil {
		return "", nil
	}
	return res, nil
}
