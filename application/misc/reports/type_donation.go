package reports

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/yudapc/go-rupiah"
	"strconv"
)

func (r *Report) TypeDonationReport() {
	if r.document == nil {
		return
	}

	/**
		Header
	 */

	r.document.SetFillColor(255, 255, 255)
	r.document.SetTextColor(0, 0, 0)

	// title
	r.document.SetFont("Arial", "", (float64)(24))
	r.document.CellFormat((float64)(len("Laporan Donasi Per Jenis Donasi")*7), (float64)(24/2), "Laporan Donasi Per jenis Donasi", "0", 0, "L", true, 0, "")
	r.document.CellFormat((float64)(len("")*10), (float64)(50/2), "", "0", 0, "L", true, 0, "")
	r.document.Ln(-1)

	/**
		End of header
	 */

	const (
		colCount = 4
		rowCount = 3
		margin1  = 32.0
		fontHt   = 10.0 // point
	)
	cellHt := r.document.PointToUnitConvert(fontHt) + 20.0
	innerCellHt := r.document.PointToUnitConvert(fontHt) + 10.0
	r.document.SetFont("Arial", "B", fontHt)

	// header-style
	r.document.SetFillColor(247, 249, 255)
	r.document.SetFontSize(11)

	r.document.CellFormat((float64)(len("Jenis Donasi")*3), (float64)(cellHt/2), "Jenis Donasi", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len("Jumlah Dana Terkumpul")*4), (float64)(cellHt/2), "Jumlah Dana Terkumpul", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len("% Total")*3), (float64)(cellHt/2), "% Total", "0", 0, "C", true, 0, "")
	r.document.Ln(-1)

	data := r.data.ReportPerCategory.ExportPdfPerCategoryDetail
	for i := 0; i < len(data); i++ {
		// table coloring
		if i%2 == 1 {
			r.document.SetFillColor(247, 249, 250)
		} else {
			r.document.SetFillColor(255, 255, 255)
		}

		if data[i].Name == "Total" {
			r.document.SetFillColor(227, 225, 232)
			r.document.SetFont("Arial", "B", fontHt)
		} else {
			r.document.SetFont("Arial", "", fontHt)
		}

		totalString := rupiah.FormatRupiah(float64(data[i].Total))
		percentString := strconv.FormatUint(data[i].Percent, 10)
		typeDonation := data[i].Name
		r.document.CellFormat((float64)(len("Jenis Donasi")*3), (float64)(innerCellHt/2), typeDonation, "0", 0, "L", true, 0, "")
		r.document.CellFormat((float64)(len("Jumlah Dana Terkumpul")*4), (float64)(innerCellHt/2), totalString, "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("% Total")*3), (float64)(innerCellHt/2), percentString+"%", "0", 0, "C", true, 0, "")
		r.document.Ln(-1)
	}

	totalpercentCorporate := strconv.FormatUint(r.data.ReportPerCategory.TotalPercentCorporate, 10)
	totalCorporate := rupiah.FormatRupiah(float64(r.data.ReportPerCategory.TotalCorporate))
	totalDonationCorporate := strconv.FormatUint(r.data.ReportPerCategory.TotalRowCountCorporate, 10)

	totalpercentPersonal := strconv.FormatUint(r.data.ReportPerCategory.TotalPercentPersonal, 10)
	totalPersonal := rupiah.FormatRupiah(float64(r.data.ReportPerCategory.TotalPersonal))
	totalDonationPersonal := strconv.FormatUint(r.data.ReportPerCategory.TotalRowCountPersonal, 10)

	r.document.SetFillColor(255, 255, 255)
	r.document.SetTextColor(0, 0, 0)
	r.document.SetFont("Arial", "", fontHt)
	LineBreak(r.document, 20)

	r.document.SetTextColor(42, 69, 191)
	r.document.SetFontSize(16)
	r.document.CellFormat(120, (float64)(cellHt/2), totalpercentCorporate+"%", "0", 0, "L", true, 0, "")

	r.document.SetTextColor(222, 13, 86)
	r.document.CellFormat(120, (float64)(cellHt/2), totalpercentPersonal+"%", "0", 0, "L", true, 0, "")
	r.document.Ln(-1)

	r.document.SetTextColor(0, 0, 0)
	r.document.SetFontSize(8)
	r.document.CellFormat(120, 2, "Terkumpul dari Corp", "0", 0, "L", true, 0, "")
	r.document.CellFormat(120, 2, "Terkumpul dari Perorangan", "0", 0, "L", true, 0, "")
	r.document.Ln(-1)

	r.document.SetTextColor(0, 0, 0)
	r.document.SetFont("Arial", "B", fontHt)
	r.document.SetFontSize(16)
	r.document.CellFormat(120, (float64)(20/2), totalCorporate, "0", 0, "L", true, 0, "")
	r.document.CellFormat(120, (float64)(20/2), totalPersonal, "0", 0, "L", true, 0, "")
	r.document.Ln(-1)

	r.document.SetTextColor(0, 0, 0)
	r.document.SetFont("Arial", "", fontHt)
	r.document.SetFontSize(8)
	r.document.CellFormat(120, (float64)(15/2), totalDonationCorporate+" Donatur", "0", 0, "L", true, 0, "")
	r.document.CellFormat(120, (float64)(15/2), totalDonationPersonal+" Donatur", "0", 0, "L", true, 0, "")

	r.document.AddPage()
}

func LineBreak(pdf *gofpdf.Fpdf, size int) {
	pdf.CellFormat((float64)(len("")*10), (float64)(size/2), "", "0", 0, "L", true, 0, "")
	pdf.Ln(-1)
}
