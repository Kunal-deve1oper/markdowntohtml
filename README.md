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

Use the ConvertToFile function to convert a Markdown file to HTML:

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

# API Reference

### ConvertToFile(filename string) error

- The ConvertToFile function takes the path to a Markdown file as its argument and generates an index.html file containing the HTML representation of the Markdown content.
- filename: Path to the Markdown file to be converted.
- Returns an error if any operation fails during the conversion process.

# NOTE

has all the necessary features and functions.
Currently working on parsing table.
