package mdparser

import (
	"strings"

	"github.com/Kunal-deve1oper/md-to-html/pkg/userRegexp"
)

// make word of format __hello__ or **hello** to bold
func Bold(word string) string {
	temp := "<strong>"
	temp += word[2 : len(word)-2]
	temp += "</strong>"
	return temp
}

// make word of format _hello_ or *hello* to italics
func Italic(word string) string {
	temp := "<em>"
	temp += word[1 : len(word)-1]
	temp += "</em>"
	return temp
}

// make word of format ***hello*** to bold and italics
func Bold_Italics(word string) string {
	temp := "<em><strong>"
	temp += word[3 : len(word)-3]
	temp += "</strong></em>"
	return temp
}

// Parse sentence to bold or italics or both
func ParseSentence(line string) string {
	// Check if the line starts and ends with "***", indicating bold and italics
	if strings.HasPrefix(line, "***") && strings.HasSuffix(line, "***") {
		return Bold_Italics(line)
	} else if (strings.HasPrefix(line, "**") && strings.HasSuffix(line, "**")) ||
		(strings.HasPrefix(line, "__") && strings.HasSuffix(line, "__")) {
		// Check if the line starts and ends with "**" or "__", indicating bold
		return Bold(line)
	} else if (strings.HasPrefix(line, "*") && strings.HasSuffix(line, "*")) ||
		(strings.HasPrefix(line, "_") && strings.HasSuffix(line, "_")) {
		// Check if the line starts and ends with "*" or "_", indicating italics
		return Italic(line)
	}
	// If the line doesn't match the above patterns, check individual words for formatting
	temp := strings.Split(line, " ")
	for i, word := range temp {
		if userRegexp.BoldAndItalics.MatchString(word) {
			// If the word matches the pattern for bold and italics, apply formatting
			temp[i] = Bold_Italics(word)
		} else if userRegexp.BoldWithUderscore.MatchString(word) || userRegexp.BoldWithStar.MatchString(word) {
			// If the word matches the pattern for bold, apply bold formatting
			temp[i] = Bold(word)
		} else if userRegexp.ItalicsWithStar.MatchString(word) || userRegexp.ItalicsWithUnderscore.MatchString(word) {
			// If the word matches the pattern for italics, apply italics formatting
			temp[i] = Italic(word)
		}
	}
	// Reconstruct the line with modified words and return
	line = strings.Join(temp, " ")
	return line
}
