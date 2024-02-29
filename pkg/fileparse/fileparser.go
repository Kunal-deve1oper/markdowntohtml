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
			f.Content = append(f.Content, line)
			letter = ""
		} else {
			letter += ele
		}
		i++
	}
	if letter != "" {
		f.Content = append(f.Content, letter)
	}
	return nil
}
