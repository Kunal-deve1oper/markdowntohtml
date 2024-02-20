package mdparser

// used to parse "- my list 1" type of strings
func ListParser(content []string) string {
	res := ""
	for _, data := range content {
		i := 2
		for string(data[i]) == " " {
			i++
		}
		pasresdData := data[i:]
		pasresdData = ParseSentence(pasresdData)
		res += `<li style="color:green; font-size: 20px;">` + pasresdData + "</li>\n"
	}
	ans := ""
	ans += "<ul>\n" + res + "</ul>\n"
	return ans
}
