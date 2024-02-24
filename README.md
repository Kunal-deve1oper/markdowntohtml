# Markdown to HTML Converter

This Go program converts Markdown files to HTML.

## Usage

The program accepts a Markdown file as input and converts it to HTML.

### Prerequisites

Before using this tool, ensure that you have Go installed on your system.

### Installation

To install the tool, you need to have Go installed. Then, you can clone the repository and build the executable.

```bash
git clone https://github.com/Kunal-deve1oper/md-to-html.git
```

Navigate to the project directory:

```bash
cd md-to-html
```

Build the project:

```bash
go build
```

### Command Line Usage

```bash
./md-to-list -f <path_to_markdown_file>
```

### Flags

- `-f`: Specifies the path to the Markdown file to be converted.

## Dependencies

The program depends on the following packages:

- [fileparse](https://github.com/Kunal-deve1oper/md-to-html/pkg/fileparse): This package provides functionality to parse Markdown files.
- [mdparser](https://github.com/Kunal-deve1oper/md-to-html/pkg/mdparser): This package handles the conversion of Markdown to HTML.

## Example

```bash
./md-to-html -f ./kunal.md
```

## Output

The program will generate an index.html file.

## Errors

If there are any errors during parsing or conversion, the program will display an error message and exit with a non-zero status code.

## Note
This project is currently under development. Some features or functionalities may not be fully implemented yet.