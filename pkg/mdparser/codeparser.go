package mdparser

func CodeParser(code []string) string {
	result := `<pre style="display: block; padding: 9.5px; margin: 0 0 10px; font-size: 15px; line-height: 1.42857143; color: #333; word-break: break-all; word-wrap: break-word; background-color: #f5f5f5; border: 1px solid #ccc; border-radius: 4px;"><code>`
	for _, data := range code {
		result += data + "\n"
	}
	result += "</code></pre>"
	return result
}
