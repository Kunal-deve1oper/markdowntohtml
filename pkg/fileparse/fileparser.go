package fileparse

import (
	"os"
	"strings"
)

type FileContent struct {
	Content []string
}

func (f *FileContent) ParseMd(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	parsedData := string(data)
	i := 0
	letter := ""
	for i < len(parsedData) {
		ele := string(parsedData[i])
		if ele == "\n" {
			line := strings.TrimSpace(letter)
			if line != "" {
				f.Content = append(f.Content, line)
			}
			for ele == "\n" {
				i++
				if i < len(parsedData) {
					f.Content = append(f.Content, "\n")
					ele = string(parsedData[i])
				} else {
					break
				}
			}
			letter = ""
		} else {
			i++
			letter += ele
		}
	}
	return nil
}
