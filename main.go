package main

import (
	"os"
	utils "pdfreader/utils"

	"github.com/ledongthuc/pdf"
)

func main() {
	pdf.DebugOn = true
	content, err := utils.ReadPdf("test_files/0478_s16_pm_21.pdf")
	if err != nil {
		panic(err)
	}
	os.WriteFile("output.txt", []byte(content), 0644)
}
