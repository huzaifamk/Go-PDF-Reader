package format

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/ledongthuc/pdf"
)

func IsSameSentence(text pdf.Text, lastTextStyle pdf.Text) bool {
	return (text.Font == lastTextStyle.Font) && (text.FontSize == lastTextStyle.FontSize) && (text.X == lastTextStyle.X)
}

func UpdateOutput(filename string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "") {
			lines[i] = strings.Replace(line, "", "", -1)
		}
	}

	for i, line := range lines {
		if strings.Contains(line, ":") {
			lines[i] = strings.Replace(line, ":", ":\n", -1)
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func FormatLines(filename string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		lines[1] = strings.Split(line, "\n")[0]
	}
	for _, line := range lines {
		lines[2] = strings.Split(line, "\n")[0]
	}
	for _, line := range lines {
		lines[3] = strings.Split(line, "\n")[0]
	}
	for _, line := range lines {
		lines[4] = strings.Split(line, "\n")[0]
	}
	for _, line := range lines {
		lines[5] = strings.Split(line, "\n")[0]
	}
	for _, line := range lines {
		lines[6] = strings.Split(line, "\n")[0]
	}
	for _, line := range lines {
		lines[7] = strings.Split(line, "\n")[0]
	}

	for i, line := range lines {
		if strings.Contains(line, `.`) {
			lines[i] = strings.Replace(line, `.`, "", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `(a)`) {
			lines[i] = strings.Replace(line, `(a)`, "\n\n(a)", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `(b)`) {
			lines[i] = strings.Replace(line, `(b)`, "\n\n(b)", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `(c)`) {
			lines[i] = strings.Replace(line, `(c)`, "\n\n(c)", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `(d)`) {
			lines[i] = strings.Replace(line, `(d)`, "\n\n(d)", -1)
		}
	}
	for i, line := range lines {
		if strings.Contains(line, `(f)`) {
			lines[i] = strings.Replace(line, `(f)`, "\n\n(f)", -1)
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
