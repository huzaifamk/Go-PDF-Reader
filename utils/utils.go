package utils

import (
	"bytes"
	"fmt"
	"os"
	"pdfreader/format"

	"github.com/ledongthuc/pdf"
)

func ReadPdfText(path string) (string, error) {

	var buf bytes.Buffer
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {

		p := r.Page(pageIndex)
		var TextContent pdf.Text
		texts := p.Content().Text

		buf.WriteString(fmt.Sprintf("\n**Page %d**\n\n", pageIndex))

		for _, text := range texts {

			if format.IsSameSentence(text, TextContent) {
				TextContent.S = TextContent.S + text.S
			} else {
				TextContent = text
				buf.WriteString(text.S)
			}
		}
	}
	return buf.String(), nil
}

func ReadPdfTextWithFormatting(path string) (string, error) {
	var buf bytes.Buffer
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		var lastTextStyle pdf.Text
		texts := p.Content().Text
		for _, text := range texts {
			if format.IsSameSentence(text, lastTextStyle) {
				lastTextStyle.S = lastTextStyle.S + text.S
			} else {
				fmt.Print(lastTextStyle.S)
				os.WriteFile("output.txt", []byte(lastTextStyle.S), 0644)
				lastTextStyle = text
				buf.WriteString(text.S)
			}
		}
	}
	return buf.String(), nil
}

func ReadPdfRow(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			for _, word := range row.Content {
				fmt.Println(word.S)
			}
		}
	}
	return "", nil
}
