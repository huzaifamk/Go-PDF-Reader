package main

import (
	"os"
	utils "pdfreader/utils"
)

func main() {

	// content, err := utils.ReadPdfText("test_files/0478_s16_qp_11.pdf")
	// if err != nil {
	// 	panic(err)
	// }

	// content, err := utils.ReadPdfTextWithFormatting("test_files/0478_s16_qp_11.pdf")
	// if err != nil {
	// 	panic(err)
	// }

		content, err := utils.ReadPdfRow("test_files/0478_s16_qp_11.pdf")

	os.WriteFile("output.txt", []byte(content), 0644)

	utils.FormatLines("output.txt")
}
