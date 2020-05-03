package reports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/yudapc/go-rupiah"
	"log"
	"os"
	"strconv"
	"ziswaf-backend/domain/entities"
)

type KwitansiReport struct {
	htmlContent gofpdf.HTMLBasicType
	document    *gofpdf.Fpdf
	data        entities.Kwitansi
}

func NewKwitansiPdf(data entities.Kwitansi) KwitansiReport {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	html := pdf.HTMLBasicNew()

	report := KwitansiReport{
		htmlContent: html,
		data:        data,
	}

	report.document = pdf
	return report
}

func (r *KwitansiReport) Header() {
	r.document.SetFillColor(255, 255, 255)
	r.document.SetTextColor(0, 0, 0)

	// title
	r.document.SetFont("Arial", "", (float64)(24))
	r.document.CellFormat((float64)(len("DETAIL ZISWAF")*7), (float64)(24/2), "DETAIL ZISWAF", "0", 0, "L", true, 0, "")
	r.document.Ln(-10)
	r.document.CellFormat(float64((10*9)*3), 0, "", "1", 0, "L", true, 0, "")

	r.document.Ln(-1)
}

func (r *KwitansiReport) Body() {
	if r.document == nil {
		return
	}

	const (
		colCount    = 4
		rowCount    = 3
		margin1     = 32.0
		align       = "L"
		fontHt      = 10.0 // point
		fontHeader  = 12
		fontSection = 10
		width       = float64(10 * 9)
		height      = float64(18 / 2)
		timeLayout  = "02/01/2006 03PM"
	)

	r.document.SetFont("Arial", "", fontHt)
	data := r.data

	var (
		title       string
		description string
		totalItem   string
		refNumber   string
	)

	if data.ItemType == "Uang" {
		title = "DESKIPRSI UANG"

		description = data.CashDescription

		if data.CashDescription == "" {
			description = "-"
		}

		totalItem = rupiah.FormatRupiah(float64(data.Total))
		refNumber = data.RefNumber
	}

	if data.ItemType == "Barang" {
		title = "DESKRIPSI BARANG"

		if data.GoodDescription == "" {
			description = "-"
		}

		description = data.GoodDescription
		totalItem = strconv.Itoa(int(data.Quantity))
		refNumber = ""
	}

	r.document.Ln(10)
	r.document.SetFont("Arial", "B", fontHeader)
	r.document.CellFormat(width, height, "KATEGORI ZISWAF", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "NAMA DONATUR", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "NO HP", "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "", fontSection)
	r.document.CellFormat(width, height, data.DivisionName, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, data.DonorName, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, data.DonorPhone, "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "B", fontHeader)
	r.document.CellFormat(width, height, "TANGGAL DAN WAKTU", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "ALAMAT DONATUR", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "SUREL", "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "", fontSection)
	date := data.CreatedAt.Format(timeLayout)
	r.document.CellFormat(width, height, date, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, data.DonorAddress, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, data.DonorEmail, "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "B", fontHeader)
	r.document.CellFormat(width, height, "MA'HAD", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "NO NPWP", "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "", fontSection)
	npwp := strconv.Itoa(int(data.DonorNPWP))
	r.document.CellFormat(width, height, data.SchoolName, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, npwp, "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "B", fontHeader)
	r.document.CellFormat(width, height, "NOMOR KWITANSI", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "DESKRIPSI DONASI", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "BENTUK DONASI", "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "", fontSection)
	r.document.CellFormat(width, height, data.Kwitansi, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, data.Description, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, data.ItemType, "0", 0, align, true, 0, "")

	if data.ItemType == "Barang" {
		r.document.Ln(-1)
		r.document.SetFont("Arial", "B", fontHeader)
		r.document.CellFormat(width, height, "NAMA PETUGAS", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, title, "0", 0, align, true, 0, "")

		r.document.Ln(-1)
		r.document.SetFont("Arial", "", fontSection)
		r.document.CellFormat(width, height, data.CreatedBy, "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, description, "0", 0, align, true, 0, "")
	}

	r.document.Ln(-1)

	if data.ItemType == "Uang" {
		r.document.Ln(-1)
		r.document.SetFont("Arial", "B", fontHeader)
		r.document.CellFormat(width, height, "NAMA PETUGAS", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "TUNAI / NON TUNAI", "0", 0, align, true, 0, "")

		r.document.Ln(-1)
		r.document.SetFont("Arial", "", fontSection)
		r.document.CellFormat(width, height, data.CreatedBy, "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, data.ItemCategory, "0", 0, align, true, 0, "")

		r.document.Ln(-1)
		r.document.SetFont("Arial", "B", fontHeader)
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "REF NUMBER", "0", 0, align, true, 0, "")

		r.document.Ln(-1)
		r.document.SetFont("Arial", "", fontSection)
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
		r.document.CellFormat(width, height, refNumber, "0", 0, align, true, 0, "")
	}

	r.document.Ln(-1)
	r.document.SetFont("Arial", "B", fontHeader)
	r.document.CellFormat(width, height, "JENIS DONASI", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "JUMLAH", "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "", fontSection)
	r.document.CellFormat(width, height, data.Category, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, totalItem, "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "B", fontHeader)
	r.document.CellFormat(width, height, "KETERANGAN DONASI", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "STATUS", "0", 0, align, true, 0, "")

	r.document.Ln(-1)
	r.document.SetFont("Arial", "", fontSection)
	r.document.CellFormat(width, height, data.StatmentCategory, "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "", "0", 0, align, true, 0, "")
	r.document.CellFormat(width, height, "Diterima Petugas", "0", 0, align, true, 0, "")
}

func (r *KwitansiReport) Output(filename string, c *gin.Context) {
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
