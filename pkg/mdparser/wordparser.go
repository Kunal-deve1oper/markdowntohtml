package mdparser

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

func Bold_Italics(word string) string {
	temp := "<em><strong>"
	temp += word[3 : len(word)-3]
	temp += "</strong></em>"
	return temp
}
