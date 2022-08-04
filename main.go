package main

import (
	"os"
	// "pdfreader/format"
	"pdfreader/format"
	utils "pdfreader/utils"
)

func main() {

	content, err := utils.ReadPdfText("test_files/LECTURE-1.pdf")
	if err != nil {
		panic(err)
	}

	os.WriteFile("output.txt", []byte(content), 0644)

	format.UpdateOutput("output.txt")

	format.FormatLines("output.txt")
}
