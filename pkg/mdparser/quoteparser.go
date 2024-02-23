package mdparser

func QuoteParse(quotes []string) string {
	res := `<blockquote style="padding: 10px 20px; margin: 0 0 20px; font-size: 20.5px; border-left: 5px solid #eee;"><p>`
	for _, data := range quotes {
		if string(data[1]) == " " {
			res += data[2:] + "<br>"
		} else {
			res += data[1:] + "<br>"
		}
	}
	res += `</p></blockquote>`
	return res
}
