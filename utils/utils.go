package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

func ReadPdfText(path string) (string, error) {
	
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
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
			if IsSameSentence(text, lastTextStyle) {
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

func IsSameSentence(text pdf.Text, lastTextStyle pdf.Text) bool {
	return (text.Font == lastTextStyle.Font) && (text.FontSize == lastTextStyle.FontSize) && (text.X == lastTextStyle.X)
}

func FormatLines(filename string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, `.`) {
			lines[i] = strings.Replace(line, `.`, "", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `(a)`) {
			lines[i] = strings.Replace(line, `(a)`, "\n\n(a)", -1)
		}

		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
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
			println(">>>> row: ", row.Position)
			for _, word := range row.Content {
				fmt.Println(word.S)
			}
		}
	}
	return "", nil
}
