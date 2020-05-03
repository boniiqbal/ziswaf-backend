package reports

import (
	"github.com/yudapc/go-rupiah"
	"strconv"
)

func (r *Report) GoodsReport() {
	if r.document == nil {
		return
	}

	const (
		colCount = 4
		rowCount = 3
		margin1  = 32.0
		fontHt   = 10.0 // point
	)

	/**
		Header
	 */
	r.document.SetFillColor(255, 255, 255)
	r.document.SetTextColor(0, 0, 0)

	// title
	r.document.SetFont("Arial", "", (float64)(24))
	r.document.CellFormat((float64)(len("Donasi Barang")*7), (float64)(24/2), "Donasi Barang", "0", 0, "L", true, 0, "")
	r.document.CellFormat((float64)(len("")*10), (float64)(50/2), "", "0", 0, "L", true, 0, "")
	r.document.Ln(-1)

	/**
		End of header
	 */

	cellHt := r.document.PointToUnitConvert(fontHt) + 20.0
	innerCellHt := r.document.PointToUnitConvert(fontHt) + 10.0
	r.document.SetFont("Arial", "B", fontHt)

	// header-style
	r.document.SetFillColor(247, 249, 255)
	r.document.SetFontSize(11)

	r.document.CellFormat((float64)(len("Jenis Barang")*4), (float64)(cellHt/2), "Jenis Barang", "0", 0, "L", true, 0, "")
	r.document.CellFormat((float64)(len("Jumlah Donasi")*4), (float64)(cellHt/2), "Jumlah Donasi", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len("Diterima")*3), (float64)(cellHt/2), "Diterima", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len("Estimasi Nilai")*3), (float64)(cellHt/2), "Estimasi Nilai", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len("Belum Diterima")*3), (float64)(cellHt/2), "Belum Diterima", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len("Estimasi Nilai")*3), (float64)(cellHt/2), "Estimasi Nilai", "0", 0, "C", true, 0, "")
	r.document.Ln(-1)

	data := r.data.ReportGood
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

		name := data[i].Name
		totalCount := strconv.Itoa(data[i].TotalRowCount)
		totalCollect := rupiah.FormatRupiah(float64(data[i].TotalCollect))
		countCollect := strconv.Itoa(int(data[i].TotalCollectCount))

		countNotCollect := strconv.Itoa(int(data[i].TotalNotCollectCount))
		totalNotCollect := rupiah.FormatRupiah(float64(data[i].TotalNotCollect))

		//estimate := strconv.FormatUint(data[i].)
		r.document.CellFormat((float64)(len("Jenis Barang")*4), (float64)(innerCellHt/2), name, "0", 0, "L", true, 0, "")
		r.document.CellFormat((float64)(len("Jumlah Donasi")*4), (float64)(innerCellHt/2), totalCount, "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("Diterima")*3), (float64)(innerCellHt/2), countCollect, "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("Estimasi Nilai")*3), (float64)(innerCellHt/2), totalCollect, "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("Belum Diterima")*3), (float64)(innerCellHt/2), countNotCollect, "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("Estimasi Nilai")*3), (float64)(innerCellHt/2), totalNotCollect, "0", 0, "C", true, 0, "")
		r.document.Ln(-1)
	}

	r.document.AddPage()
}