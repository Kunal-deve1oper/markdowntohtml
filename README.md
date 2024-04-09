# Markdown to HTML Converter

## Overview

The `markdowntohtml` package provides functionalities to convert Markdown files to HTML format. It includes a `ConvertToFile` function that takes the path to a Markdown file as input and generates an `index.html` file containing the HTML representation of the Markdown content.

## Installation

To install the `markdowntohtml` package, you can use the following `go get` command:

```go
go get github.com/Kunal-deve1oper/markdowntohtml
```

## Usage

### Importing the Package

Import the package in your Go code:

```go
import (
	"github.com/Kunal-deve1oper/markdowntohtml"
)
```

# Converting Markdown to HTML

Use the `ConvertToFile` function to convert a Markdown file to HTML:

```go
err := markdowntohtml.ConvertToFile("path/to/markdownfile.md")
if err != nil {
    // Handle error
}
```

This function reads the Markdown file located at the specified path, converts its contents to HTML, and saves the HTML output to an index.html file in the current directory.

# Example

```go
package main

import (
	"github.com/Kunal-deve1oper/markdowntohtml"
)

func main() {
	err := markdowntohtml.ConvertToFile("example.md")
	if err != nil {
		// Handle error
	}
}
```

Use the `ConvertToString` function to convert a Markdown file to HTML:

```go
res, err := markdowntohtml.ConvertToString("path to markdownfile.md")
if err != nil {
    // Handle error
}
fmt.Println(res)
```

Function takes path to the markdownfile as argument and give output of the HTML code as a string

# Example

```go
package main

import (
	"fmt"

	"github.com/Kunal-deve1oper/markdowntohtml"
)

func main() {
	res, err := markdowntohtml.ConvertToString("path to markdownfile.md")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```

Use the `ConvertInputToString` function to convert a input slice of sting to HTML:

```go
res, err := markdowntohtml.ConvertInputToString("takes input a string slice")
if err != nil {
    // Handle error
}
fmt.Println(res)
```

ConvertInputToString takes a slice of strings where each element represents a line of Markdown content. The function will process each line of this input slice and convert it into HTML code.

```go
package main

import (
	"fmt"

	"github.com/Kunal-deve1oper/markdowntohtml"
)

func main() {
	temp := make([]string, 0)
	temp = append(temp, "```")
	temp = append(temp, "npm install multer")
	temp = append(temp, "```")
	temp = append(temp, "___Hello___")

	res, err := markdowntohtml.ConvertInputToString(temp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

```

# API Reference

### ConvertToFile(filename string) error

- The ConvertToFile function takes the path to a Markdown file as its argument and generates an index.html file containing the HTML representation of the Markdown content.
- filename: Path to the Markdown file to be converted.
- Returns an error if any operation fails during the conversion process.

### ConvertToString(filename string) (string, error)

- The ConvertToFile function takes the path to a Markdown file as its argument and generates an index.html file containing the HTML representation of the Markdown content.
- filename: Path to the Markdown file to be converted.
- Returns an error if any operation fails during the conversion process.

### ConvertInputToString(input []string) (string, error)
- The ConvertInputToString function takes a slice of strings representing Markdown content and converts it into HTML format.
- input: A slice of strings where each element represents a line of Markdown content.
- Returns the resulting HTML string and an error if any operation fails during the conversion process.

# NOTE

has all the necessary features and functions.
Currently working on parsing table.
