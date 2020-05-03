package reports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"log"
	"os"
	"ziswaf-backend/domain/entities"
)

type Report struct {
	htmlContent gofpdf.HTMLBasicType
	document    *gofpdf.Fpdf
	data        entities.ExportPdf
}

func NewPdf(data entities.ExportPdf) Report {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)
	pdf.AddPage()

	html := pdf.HTMLBasicNew()

	report := Report{
		htmlContent: html,
		data:        data,
	}

	report.document = pdf
	return report
}

func (r *Report) Footer() {
	r.document.SetFooterFunc(func() {
		r.document.SetY(-10)
		r.document.SetFont("Arial", "I", 8)
		r.document.CellFormat(0, 2, fmt.Sprintf("Page %d/{nb}", r.document.PageNo()),
			"", 0, "R", false, 0, "")
	})
	r.document.AliasNbPages("")
}

func (r *Report) Output(filename string, c *gin.Context) {
	folderPath := "./attachments/pdf/"
	_ = os.MkdirAll(folderPath, 0755)

	folderPathAndFilename := folderPath + filename
	fileURI := folderPathAndFilename + ".pdf"

	err := r.document.OutputFileAndClose(fmt.Sprintf(fileURI))
	if err != nil {
		log.Println("ERROR", err.Error())
	}
	c.File(fileURI)
}
