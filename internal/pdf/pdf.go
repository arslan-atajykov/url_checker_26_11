package pdf

import (
	"bytes"
	"fmt"
	"url_checker/internal/model"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(tasks []model.Task) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.Ln(12)

	for _, task := range tasks {
		pdf.Cell(0, 8, fmt.Sprintf("Zadacha #%d", task.ID))
		pdf.Ln(8)

		for _, link := range task.Links {
			pdf.Cell(0, 6, fmt.Sprintf("%s  %s", link.URL, link.Lstatus))
			pdf.Ln(6)
		}

		pdf.Ln(4)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
