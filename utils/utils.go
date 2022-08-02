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

		for _, text := range texts {

			if format.IsSameSentence(text, TextContent) {
				TextContent.S = TextContent.S + text.S
			} else {
				TextContent = text
				buf.WriteString(text.S)
			}
		}
	}
	// b, err := r.GetPlainText()
	// if err != nil {
	// 	return "", err
	// }
	// buf.ReadFrom(b)
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
			// println(">>>> row: ", row.Position)
			for _, word := range row.Content {
				fmt.Println(word.S)
			}
		}
	}
	return "", nil
}


func ReadPdfTextWithFormattingWithRow(path string) (string, error) {
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
		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			// println(">>>> row: ", row.Position)
			for _, word := range row.Content {
				fmt.Println(word.S)
				buf.WriteString(word.S)
			}
		}
	}
	return buf.String(), nil
}


func ReadPdfTextWithFormattingWithRowAndPage(path string) (string,