package reports

import (
	"github.com/yudapc/go-rupiah"
	"strconv"
	"time"
	"ziswaf-backend/application/misc"
)

func (r *Report) MonthlyReport() {
	var (
		schoolName      string
		regencyName     string
		startDate       string
		endDate         string
		divisiWidth     int
		jmlhDonasiWidth int
		JmlnDanaWidth   int
	)

	const (
		colCount = 4
		rowCount = 3
		margin1  = 32.0
		fontHt   = 10.0 // point

	)

	if r.data.SchoolName == "" {
		schoolName = "Semua"
	} else {
		schoolName = r.data.SchoolName
	}

	if r.data.RegencyName == "" {
		regencyName = "Kota : Semua"
	} else {
		regencyName = r.data.RegencyName
	}

	if r.data.StartDate == "" {
		startDate = "-"
	} else {
		strDate, _ := time.Parse(time.RFC3339, r.data.StartDate)
		startDate = strDate.Format("January 2006")
	}

	if r.data.EndDate == "" {
		endDate = "-"
	} else {
		edDate, _ := time.Parse(time.RFC3339, r.data.EndDate)
		endDate = edDate.Format("January 2006")
	}

	role := r.data.Role

	if role == misc.ADMIN {
		divisiWidth = 3
		jmlhDonasiWidth = 2
		JmlnDanaWidth = 3
	} else {
		divisiWidth = 5
		jmlhDonasiWidth = 5
		JmlnDanaWidth = 5
	}

	/**
	Header Part
	*/
	r.document.SetFillColor(255, 255, 255)
	r.document.SetTextColor(0, 0, 0)

	// title
	r.document.SetFont("Arial", "", (float64)(24))
	r.document.CellFormat((float64)(len("Laporan Donasi Bulanan")*7), (float64)(24/2), "Laporan Donasi Bulanan", "0", 0, "L", true, 0, "")

	// date
	r.document.SetFont("Arial", "", (float64)(12))
	r.document.CellFormat((float64)(len("from "+startDate)*5), (float64)(12/2), "from "+startDate, "0", 0, "R", true, 0, "")
	r.document.CellFormat((float64)(len("to "+endDate)*4), (float64)(12/2), "to "+endDate, "0", 0, "L", true, 0, "")
	r.document.CellFormat((float64)(len("")*10), (float64)(25/2), "", "0", 0, "L", true, 0, "")

	// school name
	r.document.Ln(-1)

	//r.document.CellFormat((float64)(len("")*2), (float64)(10/2), "", "1", 0, "L", true, 0, "")
	r.document.CellFormat(165, (float64)(20/2), "Unit Ma`had : ", "0", 0, "R", true, 0, "")
	r.document.CellFormat(55, (float64)(20/2), schoolName, "0", 0, "L", true, 0, "")
	r.document.CellFormat(65, (float64)(20/2), regencyName, "0", 0, "L", true, 0, "")
	r.document.CellFormat((float64)(len("")*10), (float64)(25/2), "", "0", 0, "L", true, 0, "")

	r.document.Ln(-1)

	/**
	End of header
	*/

	cellHt := r.document.PointToUnitConvert(fontHt) + 20.0
	innerCellHt := r.document.PointToUnitConvert(fontHt) + 10.0
	r.document.SetFont("Arial", "B", fontHt)
	// r.document.SetXY(margin1, margin1)

	r.document.SetFillColor(247, 249, 255)
	r.document.SetFontSize(11)
	r.document.CellFormat((float64)(len("Bulan")*5), (float64)(cellHt/2), "Bulan", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len(" Divisi ")*divisiWidth), (float64)(cellHt/2), "Divisi", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len(" Jumlah Donasi ")*jmlhDonasiWidth), (float64)(cellHt/2), "Jumlah Donasi", "0", 0, "C", true, 0, "")
	r.document.CellFormat((float64)(len(" Jumlah Dana Terkumpul ")*JmlnDanaWidth), (float64)(cellHt/2), "Jumlah Dana Terkumpul", "0", 0, "C", true, 0, "")

	if role == misc.ADMIN {
		r.document.CellFormat((float64)(len("Prognosis")*6), (float64)(cellHt/2), "Prognosis", "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("Gap")*20), (float64)(cellHt/2), "Gap", "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len("% Capai")*3), (float64)(cellHt/2), "% Capai", "0", 0, "C", true, 0, "")
	}

	r.document.Ln(-1)

	for i := 0; i < len(r.data.Data); i++ {
		// table coloring
		if i%2 == 1 {
			r.document.SetFillColor(247, 249, 250)
		} else {
			r.document.SetFillColor(255, 255, 255)
		}

		if r.data.Data[i].Name == "Total" {
			r.document.SetFillColor(227, 225, 232)
			r.document.SetFont("Arial", "B", fontHt)
		} else if r.data.Data[i].Month == "Total" {
			r.document.SetFillColor(198, 196, 204)
			r.document.SetFont("Arial", "B", fontHt)
		} else {
			r.document.SetFont("Arial", "", fontHt)
		}

		rowString := strconv.FormatUint(r.data.Data[i].TotalRowCount, 10)
		totString := rupiah.FormatRupiah(float64(r.data.Data[i].Total))
		proString := rupiah.FormatRupiah(float64(r.data.Data[i].TotalPrognosis))
		gapString := rupiah.FormatRupiah(float64(r.data.Data[i].TotalGap))
		capString := strconv.FormatUint(r.data.Data[i].TotalPercent, 10)
		r.document.CellFormat((float64)(len("Bulan")*5), (float64)(innerCellHt/2), r.data.Data[i].Month, "0", 0, "L", true, 0, "")
		r.document.CellFormat((float64)(len(" Divisi ")*divisiWidth), (float64)(innerCellHt/2), r.data.Data[i].Name, "0", 0, "L", true, 0, "")
		r.document.CellFormat((float64)(len(" Jumlah Donasi ")*jmlhDonasiWidth), (float64)(innerCellHt/2), rowString, "0", 0, "C", true, 0, "")
		r.document.CellFormat((float64)(len(" Jumlah Dana Terkumpul ")*JmlnDanaWidth), (float64)(innerCellHt/2), totString, "0", 0, "C", true, 0, "")

		if role == misc.ADMIN {
			r.document.CellFormat((float64)(len("Prognosis")*6), (float64)(innerCellHt/2), proString, "0", 0, "C", true, 0, "")
			r.document.CellFormat((float64)(len("Gap")*20), (float64)(innerCellHt/2), gapString, "0", 0, "C", true, 0, "")
			r.document.CellFormat((float64)(len("% Capai")*3), (float64)(innerCellHt/2), capString+" %", "0", 0, "C", true, 0, "")
		}

		r.document.Ln(-1)
	}

	r.document.AddPage()
}
