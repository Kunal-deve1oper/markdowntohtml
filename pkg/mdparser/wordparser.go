package mdparser

import (
	"github.com/Kunal-deve1oper/markdowntohtml/pkg/userRegexp"
)

// ParseSentence parses a sentence and converts Markdown to HTML.
func ParseSentence(line string) string {
	// Check if the sentence contains a Markdown image
	if userRegexp.Image.MatchString(line) {
		line = userRegexp.Image.ReplaceAllString(line, `<img src="$2" alt="$1">`)
	}

	// Check if the sentence contains a Markdown link
	if userRegexp.Link.MatchString(line) {
		line = userRegexp.Link.ReplaceAllString(line, `<a href="$2">$1</a>`)
	}

	// Check if the sentence contains bold and italic Markdown syntax
	if userRegexp.BoldAndItalics.MatchString(line) {
		line = userRegexp.BoldAndItalics.ReplaceAllString(line, `<em><strong>$1</strong></em>`)
	}

	// Check if the sentence contains bold Markdown syntax using asterisks
	if userRegexp.BoldWithStar.MatchString(line) {
		line = userRegexp.BoldWithStar.ReplaceAllString(line, `<strong>$1</strong>`)
	}

	// Check if the sentence contains bold Markdown syntax using underscores
	if userRegexp.BoldWithUderscore.MatchString(line) {
		line = userRegexp.BoldWithUderscore.ReplaceAllString(line, `<strong>$1</strong>`)
	}

	// Check if the sentence contains italic Markdown syntax using asterisks
	if userRegexp.ItalicsWithStar.MatchString(line) {
		line = userRegexp.ItalicsWithStar.ReplaceAllString(line, `<em>$1</em>`)
	}

	// Check if the sentence contains italic Markdown syntax using underscores
	if userRegexp.ItalicsWithUnderscore.MatchString(line) {
		line = userRegexp.ItalicsWithUnderscore.ReplaceAllString(line, `<em>$1</em>`)
	}

	return line
}
