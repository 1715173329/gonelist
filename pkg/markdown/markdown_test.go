package markdown

import (
	"fmt"
	"testing"
)

const example = "../../example/"

func TestMarkdownToHTML(t *testing.T) {
	output, err := MarkdownToHTML(example + "README.md")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}
