package reports

import (
	"github.com/yudapc/go-rupiah"
	"strconv"
)

func (r *Report) CashReport() {
	if r.document == nil {
		return
	}

	const (
		colCount = 4
		rowCount = 3
		margin1  = 32.0
		fontHt   = 10.0 // point
	)

	r.document.SetFillColor(255, 255, 255)
	r.document.SetTextColor(0, 0, 0)

	// title
	r.document.SetFont("Arial", "", (float64)(24))
	r.document.CellFormat((float64)(len("Detail Posisi Keuangan")*7), (float64)(24/2), "Detail Posisi Keuangan", "0", 0, "L", true, 0, "")
	r.document.CellFormat((float64)(len("")*10), (float64)(50/2), "", "0", 0, "L", true, 0, "")
	r.document.Ln(-1)

	cellHt := r.document.PointToUnitConvert(fontHt) + 20.0
	innerCellHt := r.document.PointToUnitConvert(fontHt) + 10.0
	r.document.SetFont("Arial", "B", fontHt)

	// header-style
	r.document.SetFillColor(247, 249, 255)
	r.document.SetFontSize(11)

	r.document.CellFormat((float64)(len("Jenis Donasi")*4), (float64)(cellHt/2), "Jenis Donasi", "0", 0, "L", true, 0, "")
	r.document.CellFormat((float64)(len("Jumlah Dana Terkumpul")*4), (float64)(cellHt/2), "Jumlah Dana Terkumpul", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len("% Total")*3), (float64)(cellHt/2), "% Total", "0", 0, "C", true, 0, "")
	r.document.Ln(-1)

	data := r.data.ReportCash
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
		r.document.CellFormat((float64)(len("Jenis Donasi")*4), (float64)(innerCellHt/2), typeDonation, "0", 0, "L", true, 0, "")
		r.document.CellFormat((float64)(len("Jumlah Dana Terkumpul")*4), (float64)(innerCellHt/2), totalString, "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("% Total")*3), (float64)(innerCellHt/2), percentString+"%", "0", 0, "C", true, 0, "")
		r.document.Ln(-1)
	}
}
