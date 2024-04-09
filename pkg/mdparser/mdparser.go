package mdparser

import (
	"strings"
)

type MdToHtml struct {
	FinalHtml string
}

func (t *MdToHtml) HtmlParser(fileContent []string) (string, error) {
	i := 0
	for i < len(fileContent) {
		if strings.HasPrefix(fileContent[i], "#") {
			res := HeaderParser(fileContent[i])
			t.FinalHtml += res
			i++
		} else if strings.HasPrefix(fileContent[i], "- ") {
			temp := make([]string, 0)
			for i < len(fileContent) {
				if strings.HasPrefix(fileContent[i], "- ") {
					temp = append(temp, fileContent[i])
				} else {
					break
				}
				i++
			}
			res := ListParser(temp)
			t.FinalHtml += res
		} else if strings.HasPrefix(fileContent[i], "```") {
			temp := make([]string, 0)
			i++
			for i < len(fileContent) && !strings.HasPrefix(fileContent[i], "```") {
				if fileContent[i] != "" {
					temp = append(temp, fileContent[i])
				}
				i++
			}
			res := CodeParser(temp)
			t.FinalHtml += res
			i++
		} else if strings.HasPrefix(fileContent[i], ">") {
			temp := make([]string, 0)
			for i < len(fileContent) && strings.HasPrefix(fileContent[i], ">") && fileContent[i] != "" {
				temp = append(temp, fileContent[i])
				i++
			}
			res := QuoteParse(temp)
			t.FinalHtml += res
		} else {
			res := ParseSentence(fileContent[i])
			if res == "" {
				i++
				continue
			}
			t.FinalHtml += "<p>" + res + "</p>"
			i++
		}
	}
	return t.FinalHtml, nil
}
