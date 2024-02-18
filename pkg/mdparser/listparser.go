package mdparser

import (
	"strings"

	"github.com/Kunal-deve1oper/md-to-html/pkg/userRegexp"
)

// used to parse "- my list 1" type of strings
func ListParser(content []string) string {
	res := ""
	for _, data := range content {
		i := 2
		for string(data[i]) == " " {
			i++
		}
		pasresdData := data[i:]
		temp := strings.Split(pasresdData, " ")
		for i, word := range temp {
			if userRegexp.BoldAndItalics.MatchString(word) {
				temp[i] = Bold_Italics(word)
			} else if userRegexp.BoldWithUderscore.MatchString(word) || userRegexp.BoldWithStar.MatchString(word) {
				temp[i] = Bold(word)
			} else if userRegexp.ItalicsWithStar.MatchString(word) || userRegexp.ItalicsWithUnderscore.MatchString(word) {
				temp[i] = Italic(word)
			}
		}
		pasresdData = strings.Join(temp, " ")
		res += `<li style="color:green; font-size: 20px;">` + pasresdData + "</li>\n"
	}
	ans := ""
	ans += "<ul>\n" + res + "</ul>\n"
	return ans
}
