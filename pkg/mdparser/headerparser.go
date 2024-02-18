package mdparser

import "fmt"

// Used to parse `### Hello` format strings
func HeaderParser(line string) string {
	i := 0
	count := 0
	for i < len(line) {
		temp := string(line[i])
		if temp == "#" {
			count++
		}
		if count > 6 {
			return "<p>" + line + "</p>"
		}
		if temp != "#" {
			if temp == " " && count <= 6 {
				res := ""
				for temp == " " {
					i++
					temp = string(line[i])
				}
				for i < len(line) {
					res += string(line[i])
					i++
				}
				first := fmt.Sprintf(`<h%v id="%v">`, count, res)
				last := fmt.Sprintf("</h%v>", count)
				res = first + res + last
				return res
			} else {
				return "<p>" + line + "</p>\n"
			}
		}
		i++
	}
	return "<p>" + line + "</p>"
}
