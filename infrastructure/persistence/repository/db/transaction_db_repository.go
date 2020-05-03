package db

import (
	"context"
	"errors"
	"math"

	"strconv"
	"time"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) infrastructure.TransactionRepository {
	return &transactionRepository{
		DB: DB,
	}
}

func (t *transactionRepository) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	tx := t.DB.Begin()

	err := tx.Create(&transaction).Error

	if err != nil {
		tx.Rollback()
		return transaction, err
	}

	tx.Commit()
	return transaction, nil
}

func (t *transactionRepository) CreateCash(ctx context.Context, cash domain.Cash) (domain.Cash, error) {
	var (
		donor    domain.Donor
		division domain.Division
		cat      domain.Category
		statCat  domain.StatementCategory
	)

	if t.DB.First(&donor, cash.Transaction.DonorID).RecordNotFound() {
		return domain.Cash{}, errors.New("Donatur tidak ditemukan")
	}
	if t.DB.First(&division, cash.Transaction.DivisionID).RecordNotFound() {
		return domain.Cash{}, errors.New("Sumber donasi tidak ditemukan")
	}
	if t.DB.First(&cat, cash.Transaction.CategoryID).RecordNotFound() {
		return domain.Cash{}, errors.New("Jenis donasi tidak ditemukan")
	}
	if t.DB.First(&statCat, cash.Transaction.StatementCategoryID).RecordNotFound() {
		return domain.Cash{}, errors.New("Keterangan donasi tidak ditemukan")
	}

	cashCategories := []int8{1, 2, 3, 4, 5, 6, 7, 8}
	_, found := Find(cashCategories, cash.CategoryID)
	if !found {
		return domain.Cash{}, errors.New("Jenis donasi uang tidak ditemukan")
	}

	tx := t.DB.Begin()

	err := tx.Save(&cash).Error
	if err != nil {
		return cash, err
	}

	if cash.Transaction.Kwitansi == "" {
		t.CreateKwitansi(tx, cash.Transaction.CreatedBy, cash.Transaction.DivisionID, cash.Transaction.ID)
	}

	if err != nil {
		tx.Rollback()
		return cash, err
	}

	tx.Commit()
	return cash, nil
}

func (t *transactionRepository) CreateGoods(ctx context.Context, goods domain.Goods) (domain.Goods, error) {
	var (
		donor    domain.Donor
		division domain.Division
		cat      domain.Category
		statCat  domain.StatementCategory
	)

	if t.DB.First(&donor, goods.Transaction.DonorID).RecordNotFound() {
		return domain.Goods{}, errors.New("Donatur tidak ditemukan")
	}
	if t.DB.First(&division, goods.Transaction.DivisionID).RecordNotFound() {
		return domain.Goods{}, errors.New("Sumber Donasi tidak ditemukan")
	}
	if t.DB.First(&cat, goods.Transaction.CategoryID).RecordNotFound() {
		return domain.Goods{}, errors.New("Jenis Donasi tidak ditemukan")
	}
	if t.DB.First(&statCat, goods.Transaction.StatementCategoryID).RecordNotFound() {
		return domain.Goods{}, errors.New("Keterangan donasi tidak ditemukan")
	}

	goodsCategories := []int8{1, 2, 3, 4, 5}
	_, found := Find(goodsCategories, goods.CategoryID)
	if !found {
		return domain.Goods{}, errors.New("Jenis donasi barang tidak ditemukan")
	}

	tx := t.DB.Begin()
	err := tx.Create(&goods).Error

	if goods.Transaction.Kwitansi == "" {
		t.CreateKwitansi(tx, goods.Transaction.CreatedBy, goods.Transaction.DivisionID, goods.Transaction.ID)
	}

	if err != nil {
		tx.Rollback()
		return goods, err
	}

	tx.Commit()
	return goods, nil
}

func (t *transactionRepository) CreateKwitansi(tx *gorm.DB, operatorID string, divisionID uint16, transactionID uint64) {
	user := domain.User{}
	trx := domain.Transaction{}

	tx.Preload("Employee").Preload("Employee.School").First(&user, "employee_id = ?", operatorID)

	schoolID := strconv.FormatUint(user.Employee.School.ModelSoftDelete.ID, 10)
	dvsID := strconv.FormatUint(uint64(divisionID), 10)

	trxID := strconv.FormatUint(transactionID, 10)
	year := strconv.Itoa(time.Now().Year())

	kwitansi := schoolID + year + dvsID + trxID

	err := tx.First(&trx, "id = ?", trxID).Update("kwitansi", kwitansi).Error
	if err != nil {
		tx.Rollback()
	}
}

func (p *transactionRepository) GetListTransaction(ctx context.Context, filter domain.TransactionFilter, filteredDonor []domain.Donor) ([]domain.Transaction, int) {
	var (
		donorArrayStaID []uint64
		donorArraySeaID []uint64
		donorArrayID    []uint64
		cashArrayID     []uint64
		schoolArrayID   []uint64
		count           int
	)

	page, _ := strconv.Atoi(filter.Page)
	limit, _ := strconv.Atoi(filter.Limit)
	offset := (page - 1) * limit

	transaction := []domain.Transaction{}
	donor := []domain.Donor{}
	cash := []domain.Cash{}
	school := []domain.School{}

	tx := p.DB.Model(&transaction)

	if filter.Search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if filter.Regency != "" {
		p.DB.Where("regency_id = ?", filter.Regency).Find(&school)
		for _, v := range school {
			schoolArrayID = append(schoolArrayID, v.ID)
		}
		tx = tx.Where("school_id IN (?)", schoolArrayID)
	}
	if filter.DonaturStatus != "" {
		for _, v := range filteredDonor {
			donorArrayStaID = append(donorArrayStaID, v.ID)
		}
		tx = tx.Where("donor_id IN (?)", donorArrayStaID)
	}
	if filter.StartTotal != "" {
		tx = tx.Where("total >= ?", filter.StartTotal)
	}
	if filter.EndTotal != "" {
		tx = tx.Where("total <= ?", filter.EndTotal)
	}
	if filter.StartDate != "" {
		tx = tx.Where("created_at >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		tx = tx.Where("created_at <= ?", filter.EndDate)
	}
	if filter.DonaturCategory != "" {
		p.DB.Where("is_company = ?", filter.DonaturCategory).Find(&donor)
		for _, v := range donor {
			donorArrayID = append(donorArrayID, v.ID)
		}
		tx = tx.Where("donor_id IN (?)", donorArrayID)
	}
	if filter.StatementCategory != "" {
		tx = tx.Where("statement_category_id = ?", filter.StatementCategory)
	}
	if filter.DonationType != "" {
		tx = tx.Where("item_type = ?", filter.DonationType)
	}
	if filter.CategoryType != "" {
		if filter.CategoryType == "1" {
			p.DB.Where("category_id IN (?)", 1).Find(&cash)
			for _, v := range cash {
				cashArrayID = append(cashArrayID, v.ID)
			}
		} else {
			p.DB.Where("category_id IN (?)", []uint64{2, 3, 4, 5, 6, 7, 8}).Find(&cash)
			for _, v := range cash {
				cashArrayID = append(cashArrayID, v.ID)
			}
		}
		tx = tx.Where("item_id IN (?) AND item_type = ?", cashArrayID, "cashs")
	}
	if filter.SchoolID != "" {
		tx = tx.Where("school_id = ?", filter.SchoolID)
	}
	if filter.CategoryID != "" {
		tx = tx.Where("category_id = ?", filter.CategoryID)
	}
	if filter.DivisionID != "" {
		tx = tx.Where("division_id = ?", filter.DivisionID)
	}

	if filter.Search != "" {
		for _, v := range filteredDonor {
			donorArraySeaID = append(donorArraySeaID, v.ID)
		}
		tx = tx.Where("donor_id IN (?)", donorArraySeaID)
	}

	tx.Count(&count)

	if filter.Page != "" || filter.Limit != "10" {
		tx.Preload("Division").
			Preload("Category").
			Preload("Donor").
			Preload("School").
			Preload("School.Regency").
			Preload("Cashes").
			Preload("Cashes.CashCategory").
			Preload("GoodsQuery").
			Preload("GoodsQuery.GoodsCategory").
			Preload("StatementCategory").
			Offset(offset).Limit(limit).Order(filter.Sort + "" + filter.SortType).Find(&transaction)
	} else {
		tx.Preload("Division").Preload("Category").Preload("Donor").Preload("School").Preload("School.Regency").Preload("Cashes").Preload("GoodsQuery").Order(filter.Sort + "" + filter.SortType).Find(&transaction)
	}

	return transaction, count
}

func (p *transactionRepository) GetListReport(ctx context.Context, filter domain.TransactionFilter) (domain.ReportResponse, error) {
	transaction := []domain.Transaction{}
	trxLastYear := []domain.Transaction{}
	trxJan := []domain.Transaction{}
	trxFeb := []domain.Transaction{}
	trxMar := []domain.Transaction{}
	trxApr := []domain.Transaction{}
	trxMay := []domain.Transaction{}
	trxJun := []domain.Transaction{}
	trxJul := []domain.Transaction{}
	trxAug := []domain.Transaction{}
	trxSep := []domain.Transaction{}
	trxOct := []domain.Transaction{}
	trxNov := []domain.Transaction{}
	trxDes := []domain.Transaction{}
	donor := []domain.Donor{}
	transGoal := []domain.TransactionGoal{}
	cash := []domain.Cash{}
	good := []domain.Goods{}
	percentZiswaf := domain.TotalPercentZiswaf{}
	percentCash := domain.TotalCashPercent{}
	school := []domain.School{}

	var (
		total                     uint64
		totalLast                 uint64
		totalPercentUp            uint64
		totalPercentDown          uint64
		totalRowCount             int
		totalRowCountLast         int
		totalRowCountUp           int
		totalRowCountDown         int
		totalRowDonorCount        int
		totalRowDonorCountLast    int
		totalRowDonorCountUp      int
		totalRowDonorCountDown    int
		totalRowNewDonorCount     int
		totalRowNewDonorCountLast int
		totalRowNewDonorCountUp   int
		totalRowNewDonorCountDown int
		totalZakatMaal            uint64
		totalWakaf                uint64
		totalZakatFitrah          uint64
		totalInfaq                uint64
		totalKurban               uint64
		totalOther                uint64
		totalRetail               uint64
		totalRetailRowCount       int
		totalCorporate            uint64
		totalCorporateRowCount    int
		totalUpz                  uint64
		totalUpzRowCount          int
		totalPrognosisRetail      uint64
		totalPrognosisCorporate   uint64
		totalPrognosisUpz         uint64
		totalPercentRetail        uint64
		totalPercentCorporate     uint64
		totalPercentUpz           uint64
		totalCompanyDonor         uint64
		totalPersonDonor          uint64
		totalCompanyDonorRowCount int
		totalPersonDonorRowCount  int
		totalPercentCompanyDonor  uint64
		totalPercentPersonDonor   uint64
		totalRitelZakatMaal       uint64
		totalRitelWakaf           uint64
		totalRitelZakatFitrah     uint64
		totalRitelInfaq           uint64
		totalRitelKurban          uint64
		totalRitelOther           uint64
		totalCorporateZakatFitrah uint64
		totalCorporateZakatMaal   uint64
		totalCorporateInfaq       uint64
		totalCorporateWakaf       uint64
		totalCorporateKurban      uint64
		totalCorporateOther       uint64
		totalUpzZakatMaal         uint64
		totalUpzWakaf             uint64
		totalUpzZakatFitrah       uint64
		totalUpzInfaq             uint64
		totalUpzKurban            uint64
		totalUpzOther             uint64
		totalGood                 uint64
		totalGoodCount            uint64
		totalGoodMoveCount        int32
		totalGoodNotMoveCount     int32
		totalGoodFoodCount        int32
		totalGoodOtherCount       int32
		totalAllCash              uint64
		totalCashCount            uint64
		totalCash                 uint64
		totalNonCashMuamalat      uint64
		totalNonCashMandiri       uint64
		totalNonCashBsm           uint64
		totalNonCashBri           uint64
		totalNonCashBniLamp       uint64
		totalNonCashBniSy         uint64
		totalNonCashBca           uint64
		totUpzJan                 uint64
		totUpzFeb                 uint64
		totUpzMar                 uint64
		totUpzApr                 uint64
		totUpzMay                 uint64
		totUpzJun                 uint64
		totUpzJul                 uint64
		totUpzAug                 uint64
		totUpzSep                 uint64
		totUpzOct                 uint64
		totUpzNov                 uint64
		totUpzDes                 uint64
		totRetJan                 uint64
		totRetFeb                 uint64
		totRetMar                 uint64
		totRetApr                 uint64
		totRetMay                 uint64
		totRetJun                 uint64
		totRetJul                 uint64
		totRetAug                 uint64
		totRetSep                 uint64
		totRetOct                 uint64
		totRetNov                 uint64
		totRetDes                 uint64
		totCorpJan                uint64
		totCorpFeb                uint64
		totCorpMar                uint64
		totCorpApr                uint64
		totCorpMay                uint64
		totCorpJun                uint64
		totCorpJul                uint64
		totCorpAug                uint64
		totCorpSep                uint64
		totCorpOct                uint64
		totCorpNov                uint64
		totCorpDes                uint64
		totalZiswaf               domain.TotalZiswafPerMonth
		totalZakMaal              domain.TotalZakatMaalPerMonth
		totalZakFit               domain.TotalZakatFitrahPerMonth
		totalInf                  domain.TotalInfaqPerMonth
		totalWaq                  domain.TotalWakafPerMonth
		totalQur                  domain.TotalQurbanPerMonth
		totalOth                  domain.TotalOtherPerMonth
		tot                       domain.ZiswafMonth
		totZaMaal                 domain.ZakatMaalMonth
		totZaFit                  domain.ZakatFitrahMonth
		totInfaq                  domain.InfaqMonth
		totKurban                 domain.QurbanMonth
		totWaqaf                  domain.WaqafMonth
		totOther                  domain.OtherMonth
		totTrxMonth               domain.TotalTransactionPerMonth
		totProgMonth              domain.TotalPrognosisPerMonth
		totLastYear               domain.TotalTransactionLastYear
		dataTrx                   []domain.Transaction
		totUpzMonth               domain.TotalUpzPerMonth
		totRetMonth               domain.TotalRetailPerMonth
		totCorMonth               domain.TotalCorporatePerMonth
		totDivPercent             domain.TotalDivisionPercent
		donorArrayID              []uint64
		donorLastArrayID          []uint64
		schoolArrayID             []uint64
	)

	if filter.StartDate == "" || filter.EndDate == "" {
		return domain.ReportResponse{}, errors.New("Date is empty")
	}

	// Get Year
	yearStart := misc.TimeHelper(filter.StartDate, "year")
	yearEnd := misc.TimeHelper(filter.EndDate, "year")
	month := misc.MonthHelper(yearEnd)

	if yearStart != yearEnd {
		return domain.ReportResponse{}, errors.New("Tahun pada filter harus pada tahun yang sama")
	}

	// Get Date last year
	dateLastYear, dateEndLastYear := misc.DateHelper(filter.StartDate, filter.EndDate, "year")

	tx := p.DB.Model(&transaction)
	ty := p.DB.Model(&trxLastYear)
	ta := p.DB.Model(&trxJan)
	tb := p.DB.Model(&trxFeb)
	tc := p.DB.Model(&trxMar)
	td := p.DB.Model(&trxApr)
	te := p.DB.Model(&trxMay)
	tf := p.DB.Model(&trxJun)
	tg := p.DB.Model(&trxJul)
	th := p.DB.Model(&trxAug)
	ti := p.DB.Model(&trxSep)
	tj := p.DB.Model(&trxOct)
	tk := p.DB.Model(&trxNov)
	tl := p.DB.Model(&trxDes)

	if filter.SchoolID == "" {
		tx = tx.Where("id > ?", 0)
		ty = ty.Where("id > ?", 0)
		ta = ta.Where("id > ?", 0)
		tb = tb.Where("id > ?", 0)
		tc = tc.Where("id > ?", 0)
		td = td.Where("id > ?", 0)
		te = te.Where("id > ?", 0)
		tf = tf.Where("id > ?", 0)
		tg = tg.Where("id > ?", 0)
		th = th.Where("id > ?", 0)
		ti = ti.Where("id > ?", 0)
		tj = tj.Where("id > ?", 0)
		tk = tk.Where("id > ?", 0)
		tl = tl.Where("id > ?", 0)
	}
	if filter.StartDate != "" {
		tx = tx.Where("created_at >= ?", filter.StartDate)
		ty = ty.Where("created_at >= ?", dateLastYear)
	}
	if filter.EndDate != "" {
		tx = tx.Where("created_at <= ?", filter.EndDate)
		ty = ty.Where("created_at <= ?", dateEndLastYear)
	}
	if filter.Regency != "" {
		p.DB.Where("regency_id = ?", filter.Regency).Find(&school)
		for _, v := range school {
			schoolArrayID = append(schoolArrayID, v.ID)
		}
		tx = tx.Where("school_id IN (?)", schoolArrayID)
	}
	if filter.SchoolID != "" {
		tx = tx.Where("school_id = ?", filter.SchoolID)
		ty = ty.Where("school_id = ?", filter.SchoolID)
		ta = ta.Where("school_id = ?", filter.SchoolID)
		tb = tb.Where("school_id = ?", filter.SchoolID)
		tc = tc.Where("school_id = ?", filter.SchoolID)
		td = td.Where("school_id = ?", filter.SchoolID)
		te = te.Where("school_id = ?", filter.SchoolID)
		tf = tf.Where("school_id = ?", filter.SchoolID)
		tg = tg.Where("school_id = ?", filter.SchoolID)
		th = th.Where("school_id = ?", filter.SchoolID)
		ti = ti.Where("school_id = ?", filter.SchoolID)
		tj = tj.Where("school_id = ?", filter.SchoolID)
		tk = tk.Where("school_id = ?", filter.SchoolID)
		tl = tl.Where("school_id = ?", filter.SchoolID)
	}

	tx.Preload("Donor").Find(&transaction).Count(&totalRowCount)
	ty.Find(&trxLastYear).Count(&totalRowCountLast)
	ta.Where("created_at >= ? AND created_at <= ?", month.JanuaryStart, month.JanuaryEnd).Find(&trxJan)
	tb.Where("created_at >= ? AND created_at <= ?", month.FebruaryStart, month.FebruaryEnd).Find(&trxFeb)
	tc.Where("created_at >= ? AND created_at <= ?", month.MarchStart, month.MarchEnd).Find(&trxMar)
	td.Where("created_at >= ? AND created_at <= ?", month.AprilStart, month.AprilEnd).Find(&trxApr)
	te.Where("created_at >= ? AND created_at <= ?", month.MayStart, month.MayEnd).Find(&trxMay)
	tf.Where("created_at >= ? AND created_at <= ?", month.JuneStart, month.JuneEnd).Find(&trxJun)
	tg.Where("created_at >= ? AND created_at <= ?", month.JulyStart, month.JulyEnd).Find(&trxJul)
	th.Where("created_at >= ? AND created_at <= ?", month.AugustStart, month.AugustEnd).Find(&trxAug)
	ti.Where("created_at >= ? AND created_at <= ?", month.SeptemberStart, month.SeptemberEnd).Find(&trxSep)
	tj.Where("created_at >= ? AND created_at <= ?", month.OctoberStart, month.OctoberEnd).Find(&trxOct)
	tk.Where("created_at >= ? AND created_at <= ?", month.NovemberStart, month.NovemberEnd).Find(&trxNov)
	tl.Where("created_at >= ? AND created_at <= ?", month.DecemberStart, month.DecemberEnd).Find(&trxDes)

	// Total Transaction
	for _, v := range transaction {
		total += v.Total
	}

	// Total Transaction Last Year
	for _, v := range trxLastYear {
		totalLast += v.Total
	}

	// Get different total
	if totalLast == 0 {
		totalPercentUp = 100
		totalPercentDown = 0
	} else {
		totPercent := percent(totalLast, total)
		if totPercent > 100 {
			totalPercentDown = totPercent - 100
			totalPercentUp = 0
		} else if totPercent < 100 {
			totalPercentUp = 100 - totPercent
			totalPercentDown = 0
		}
	}

	// Get different row count
	if totalRowCount > totalRowCountLast {
		totalRowCountUp = totalRowCount - totalRowCountLast
		totalRowCountDown = 0
	} else if totalRowCount < totalRowCountLast {
		totalRowCountDown = totalRowCountLast - totalRowCount
		totalRowCountUp = 0
	}

	if total == 0 {
		return domain.ReportResponse{}, errors.New("Data Not Found")
	}

	// Total Ziswaf January
	for _, v := range trxJan {
		if v.DivisionID == 1 {
			totUpzJan += v.Total
		} else if v.DivisionID == 2 {
			totRetJan += v.Total
		} else if v.DivisionID == 3 {
			totCorpJan += v.Total
		}
	}

	// Total Ziswaf February
	for _, v := range trxFeb {
		if v.DivisionID == 1 {
			totUpzFeb += v.Total
		} else if v.DivisionID == 2 {
			totRetFeb += v.Total
		} else if v.DivisionID == 3 {
			totCorpFeb += v.Total
		}
	}

	// Total Ziswaf March
	for _, v := range trxMar {
		if v.DivisionID == 1 {
			totUpzMar += v.Total
		} else if v.DivisionID == 2 {
			totRetMar += v.Total
		} else if v.DivisionID == 3 {
			totCorpMar += v.Total
		}
	}

	// Total Ziswaf April
	for _, v := range trxApr {
		if v.DivisionID == 1 {
			totUpzApr += v.Total
		} else if v.DivisionID == 2 {
			totRetApr += v.Total
		} else if v.DivisionID == 3 {
			totCorpApr += v.Total
		}
	}

	// Total Ziswaf May
	for _, v := range trxMay {
		if v.DivisionID == 1 {
			totUpzMay += v.Total
		} else if v.DivisionID == 2 {
			totRetMay += v.Total
		} else if v.DivisionID == 3 {
			totCorpMay += v.Total
		}
	}

	// Total Ziswaf Jun
	for _, v := range trxJun {
		if v.DivisionID == 1 {
			totUpzJun += v.Total
		} else if v.DivisionID == 2 {
			totRetJun += v.Total
		} else if v.DivisionID == 3 {
			totCorpJun += v.Total
		}
	}

	// Total Ziswaf Julye
	for _, v := range trxJul {
		if v.DivisionID == 1 {
			totUpzJul += v.Total
		} else if v.DivisionID == 2 {
			totRetJul += v.Total
		} else if v.DivisionID == 3 {
			totCorpJul += v.Total
		}
	}

	// Total Ziswaf August
	for _, v := range trxAug {
		if v.DivisionID == 1 {
			totUpzAug += v.Total
		} else if v.DivisionID == 2 {
			totRetAug += v.Total
		} else if v.DivisionID == 3 {
			totCorpAug += v.Total
		}
	}

	// Total Ziswaf September
	for _, v := range trxSep {
		if v.DivisionID == 1 {
			totUpzSep += v.Total
		} else if v.DivisionID == 2 {
			totRetSep += v.Total
		} else if v.DivisionID == 3 {
			totCorpSep += v.Total
		}
	}

	// Total Ziswaf October
	for _, v := range trxOct {
		if v.DivisionID == 1 {
			totUpzOct += v.Total
		} else if v.DivisionID == 2 {
			totRetOct += v.Total
		} else if v.DivisionID == 3 {
			totCorpOct += v.Total
		}
	}

	// Total Ziswaf November
	for _, v := range trxNov {
		if v.DivisionID == 1 {
			totUpzNov += v.Total
		} else if v.DivisionID == 2 {
			totRetNov += v.Total
		} else if v.DivisionID == 3 {
			totCorpNov += v.Total
		}
	}

	// Total Ziswaf Desember
	for _, v := range trxDes {
		if v.DivisionID == 1 {
			totUpzDes += v.Total
		} else if v.DivisionID == 2 {
			totRetDes += v.Total
		} else if v.DivisionID == 3 {
			totCorpDes += v.Total
		}
	}

	// Categories
	for _, v := range transaction {
		if v.CategoryID == 2 {
			totalZakatMaal += v.Total
		} else if v.CategoryID == 4 {
			totalWakaf += v.Total
		} else if v.CategoryID == 3 {
			totalInfaq += v.Total
		} else if v.CategoryID == 1 {
			totalZakatFitrah += v.Total
		} else if v.CategoryID == 5 {
			totalKurban += v.Total
		} else if v.CategoryID == 6 {
			totalOther += v.Total
		}
	}

	// Division
	for _, v := range transaction {
		if v.DivisionID == 2 {
			totalRetailRowCount += 1
			totalRetail += v.Total
		} else if v.DivisionID == 3 {
			totalCorporateRowCount += 1
			totalCorporate += v.Total
		} else if v.DivisionID == 1 {
			totalUpzRowCount += 1
			totalUpz += v.Total
		}
	}

	// Total Ziswaf UPZ Per Month
	totUpzMonth.TotalUpzJanuary = totUpzJan
	totUpzMonth.TotalUpzFebruary = totUpzFeb
	totUpzMonth.TotalUpzMarch = totUpzMar
	totUpzMonth.TotalUpzApril = totUpzApr
	totUpzMonth.TotalUpzMay = totUpzMay
	totUpzMonth.TotalUpzJune = totUpzJun
	totUpzMonth.TotalUpzJuly = totUpzJul
	totUpzMonth.TotalUpzAugust = totUpzAug
	totUpzMonth.TotalUpzSeptember = totUpzSep
	totUpzMonth.TotalUpzOctober = totUpzOct
	totUpzMonth.TotalUpzNovember = totUpzNov
	totUpzMonth.TotalUpzDecember = totUpzDes

	// Total Ziswaf Retail Per Month
	totRetMonth.TotalRetailJanuary = totRetJan
	totRetMonth.TotalRetailFebruary = totRetFeb
	totRetMonth.TotalRetailMarch = totRetMar
	totRetMonth.TotalRetailApril = totRetApr
	totRetMonth.TotalRetailMay = totRetMay
	totRetMonth.TotalRetailJune = totRetJun
	totRetMonth.TotalRetailJuly = totRetJul
	totRetMonth.TotalRetailAugust = totRetAug
	totRetMonth.TotalRetailSeptember = totRetSep
	totRetMonth.TotalRetailOctober = totRetOct
	totRetMonth.TotalRetailNovember = totRetNov
	totRetMonth.TotalRetailDecember = totRetDes

	// Total Ziswaf Corporate Per Month
	totCorMonth.TotalCorporateJanuary = totCorpJan
	totCorMonth.TotalCorporateFebruary = totCorpFeb
	totCorMonth.TotalCorporateMarch = totCorpMar
	totCorMonth.TotalCorporateApril = totCorpApr
	totCorMonth.TotalCorporateMay = totCorpMay
	totCorMonth.TotalCorporateJune = totCorpJun
	totCorMonth.TotalCorporateJuly = totCorpJul
	totCorMonth.TotalCorporateAugust = totCorpAug
	totCorMonth.TotalCorporateSeptember = totCorpSep
	totCorMonth.TotalCorporateOctober = totCorpOct
	totCorMonth.TotalCorporateNovember = totCorpNov
	totCorMonth.TotalCorporateDecember = totCorpDes

	// Detail Per Division
	for _, v := range transaction {
		if v.DivisionID == 1 && v.CategoryID == 1 {
			totalUpzZakatFitrah += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 2 {
			totalUpzZakatMaal += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 3 {
			totalUpzInfaq += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 4 {
			totalUpzWakaf += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 5 {
			totalUpzKurban += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 6 {
			totalUpzOther += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 1 {
			totalRitelZakatFitrah += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 2 {
			totalRitelZakatMaal += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 3 {
			totalRitelInfaq += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 4 {
			totalRitelWakaf += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 5 {
			totalRitelKurban += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 6 {
			totalRitelOther += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 1 {
			totalCorporateZakatFitrah += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 2 {
			totalCorporateZakatMaal += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 3 {
			totalCorporateInfaq += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 4 {
			totalCorporateWakaf += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 5 {
			totalCorporateKurban += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 6 {
			totalCorporateOther += v.Total
		}
	}

	p.DB.Find(&cash)
	p.DB.Find(&good)

	for _, v := range transaction {
		for _, j := range cash {
			if v.ItemType == "cashs" && j.ID == v.ItemID {
				if j.CategoryID == 1 {
					totalCash += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 2 {
					totalNonCashMuamalat += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 3 {
					totalNonCashMandiri += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 4 {
					totalNonCashBsm += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 5 {
					totalNonCashBri += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 6 {
					totalNonCashBniLamp += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 7 {
					totalNonCashBniSy += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 8 {
					totalNonCashBca += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				}
				continue
			}
		}
	}

	for _, v := range transaction {
		for _, j := range good {
			if v.ItemType == "goods" && j.ID == v.ItemID {
				if j.CategoryID == 1 {
					totalGoodNotMoveCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				} else if j.CategoryID == 2 {
					totalGoodMoveCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				} else if j.CategoryID == 3 {
					totalGoodFoodCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				} else if j.CategoryID == 4 {
					totalGoodOtherCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				}
				continue
			}
		}
	}

	time := time.Now()
	beginningOfMonth := misc.BeginningOfMonth(time)
	endOfMonth := misc.EndOfMonth(time)
	lastYearMonth, lastYearEndMonth := misc.DateHelper(beginningOfMonth.String(), endOfMonth.String(), "year")

	// Get Donor in transaction
	for _, v := range transaction {
		donorArrayID = append(donorArrayID, v.DonorID)
	}
	donorArrayID = unique(donorArrayID)

	for _, v := range trxLastYear {
		donorLastArrayID = append(donorLastArrayID, v.DonorID)
	}
	donorLastArrayID = unique(donorLastArrayID)

	// Get donor data
	p.DB.Where("id IN (?) AND created_at >= ? AND created_at <= ?", donorArrayID, filter.StartDate, filter.EndDate).Find(&donor).Count(&totalRowDonorCount)
	p.DB.Where("id IN (?) AND created_at >= ? and created_at <= ?", donorArrayID, beginningOfMonth, endOfMonth).Find(&donor).Count(&totalRowNewDonorCount)
	p.DB.Where("id IN (?) AND created_at >= ? AND created_at <= ?", donorLastArrayID, dateLastYear, dateEndLastYear).Find(&donor).Count(&totalRowDonorCountLast)
	p.DB.Where("id IN (?) AND created_at >= ? and created_at <= ?", donorLastArrayID, lastYearMonth, lastYearEndMonth).Find(&donor).Count(&totalRowNewDonorCountLast)

	// Get different total row count donor
	if totalRowDonorCount > totalRowDonorCountLast {
		totalRowDonorCountUp = totalRowDonorCount - totalRowDonorCountLast
		totalRowDonorCountDown = 0
	} else if totalRowDonorCount < totalRowDonorCountLast {
		totalRowDonorCountDown = totalRowDonorCountLast - totalRowDonorCount
		totalRowDonorCountUp = 0
	}

	// Get different total row count new donor
	if totalRowNewDonorCount > totalRowNewDonorCountLast {
		totalRowNewDonorCountUp = totalRowNewDonorCount - totalRowNewDonorCountLast
		totalRowNewDonorCountDown = 0
	} else if totalRowNewDonorCount < totalRowNewDonorCountLast {
		totalRowNewDonorCountDown = totalRowNewDonorCountLast - totalRowNewDonorCount
		totalRowNewDonorCountUp = 0
	}

	// Check Donor company
	for _, v := range transaction {
		if v.Donor.IsCompany == true {
			totalCompanyDonorRowCount += 1
			totalCompanyDonor += v.Total
		} else if v.Donor.IsCompany == false {
			totalPersonDonorRowCount += 1
			totalPersonDonor += v.Total
		}
	}

	// Get Prognosis
	p.DB.Where("year = ?", yearEnd).Find(&transGoal)

	// Prognosis
	for _, v := range transGoal {
		if v.DivisionID == 2 {
			totalPrognosisRetail += v.Total
		} else if v.DivisionID == 1 {
			totalPrognosisUpz += v.Total
		} else if v.DivisionID == 3 {
			totalPrognosisCorporate += v.Total
		}
	}

	// Total Ziswaf Per Month
	for i := 1; i < 13; i++ {
		switch i {
		case 1:
			dataTrx = trxJan
		case 2:
			dataTrx = trxFeb
		case 3:
			dataTrx = trxMar
		case 4:
			dataTrx = trxApr
		case 5:
			dataTrx = trxMay
		case 6:
			dataTrx = trxJun
		case 7:
			dataTrx = trxJul
		case 8:
			dataTrx = trxAug
		case 9:
			dataTrx = trxSep
		case 10:
			dataTrx = trxOct
		case 11:
			dataTrx = trxNov
		case 12:
			dataTrx = trxDes
		}
		// Total Ziswaf
		for _, v := range dataTrx {
			tot.Total += v.Total
			if v.CategoryID == 1 {
				totZaFit.Total += v.Total
			} else if v.CategoryID == 2 {
				totZaMaal.Total += v.Total
			} else if v.CategoryID == 3 {
				totInfaq.Total += v.Total
			} else if v.CategoryID == 4 {
				totWaqaf.Total += v.Total
			} else if v.CategoryID == 5 {
				totKurban.Total += v.Total
			} else if v.CategoryID == 6 {
				totOther.Total += v.Total
			}
		}
		totalZiswaf.Month = append(totalZiswaf.Month, tot)
		totalZakMaal.Month = append(totalZakMaal.Month, totZaMaal)
		totalZakFit.Month = append(totalZakFit.Month, totZaFit)
		totalInf.Month = append(totalInf.Month, totInfaq)
		totalWaq.Month = append(totalWaq.Month, totWaqaf)
		totalQur.Month = append(totalQur.Month, totKurban)
		totalOth.Month = append(totalOth.Month, totOther)
		tot.Total = 0
		totZaMaal.Total = 0
		totZaFit.Total = 0
		totInfaq.Total = 0
		totKurban.Total = 0
		totWaqaf.Total = 0
		totOther.Total = 0
	}

	// Get Total Prognosis Per Month
	for _, v := range transGoal {
		if v.Month == 1 {
			totProgMonth.January += v.Total
		} else if v.Month == 2 {
			totProgMonth.February += v.Total
		} else if v.Month == 3 {
			totProgMonth.March += v.Total
		} else if v.Month == 4 {
			totProgMonth.April += v.Total
		} else if v.Month == 5 {
			totProgMonth.May += v.Total
		} else if v.Month == 6 {
			totProgMonth.June += v.Total
		} else if v.Month == 7 {
			totProgMonth.July += v.Total
		} else if v.Month == 8 {
			totProgMonth.August += v.Total
		} else if v.Month == 9 {
			totProgMonth.September += v.Total
		} else if v.Month == 10 {
			totProgMonth.October += v.Total
		} else if v.Month == 11 {
			totProgMonth.November += v.Total
		} else if v.Month == 12 {
			totProgMonth.December += v.Total
		}
	}

	// Get Total Transaction Per month
	totTrxMonth.January = totalZiswaf.Month[0].Total
	totTrxMonth.February = totalZiswaf.Month[1].Total
	totTrxMonth.March = totalZiswaf.Month[2].Total
	totTrxMonth.April = totalZiswaf.Month[3].Total
	totTrxMonth.May = totalZiswaf.Month[4].Total
	totTrxMonth.June = totalZiswaf.Month[5].Total
	totTrxMonth.July = totalZiswaf.Month[6].Total
	totTrxMonth.August = totalZiswaf.Month[7].Total
	totTrxMonth.September = totalZiswaf.Month[8].Total
	totTrxMonth.October = totalZiswaf.Month[9].Total
	totTrxMonth.November = totalZiswaf.Month[10].Total
	totTrxMonth.December = totalZiswaf.Month[11].Total

	// Percent
	totalPercentCompanyDonor = percent(totalCompanyDonor, total)
	totalPercentPersonDonor = percent(totalPersonDonor, total)
	totalPercentRetail = percent(totalRetail, totalPrognosisRetail)
	totalPercentCorporate = percent(totalCorporate, totalPrognosisCorporate)
	totalPercentUpz = percent(totalUpz, totalPrognosisUpz)
	percentZiswaf.ZakatMaalPercent = percent(totalZakatMaal, total)
	percentZiswaf.ZakatFitrahPercent = percent(totalZakatFitrah, total)
	percentZiswaf.WaqafPercent = percent(totalWakaf, total)
	percentZiswaf.QurbanPercent = percent(totalKurban, total)
	percentZiswaf.InfaqPercent = percent(totalInfaq, total)
	percentZiswaf.OtherPercent = percent(totalOther, total)
	percentCash.CashPercent = percent(totalCash, totalAllCash)
	percentCash.NonCashBcaPercent = percent(totalNonCashBca, totalAllCash)
	percentCash.NonCashBniLampPercent = percent(totalNonCashBniLamp, totalAllCash)
	percentCash.NonCashBniSyPercent = percent(totalNonCashBniSy, totalAllCash)
	percentCash.NonCashBriPercent = percent(totalNonCashBri, totalAllCash)
	percentCash.NonCashBsmPercent = percent(totalNonCashBsm, totalAllCash)
	percentCash.NonCashMandiriPercent = percent(totalNonCashMandiri, totalAllCash)
	percentCash.NonCashMuamalatPercent = percent(totalNonCashMuamalat, totalAllCash)
	totDivPercent.TotalDivisionUpzPercent = percent(totalUpz, total)
	totDivPercent.TotalDivisionRetailPercent = percent(totalRetail, total)
	totDivPercent.TotalDivisionCorporatePercent = percent(totalCorporate, total)

	totLastYear.TotalUp = totalPercentUp
	totLastYear.TotalDown = totalPercentDown
	totLastYear.CountUp = totalRowCountUp
	totLastYear.CountDown = totalRowCountDown
	totLastYear.DonorUp = totalRowDonorCountUp
	totLastYear.DonorDown = totalRowDonorCountDown
	totLastYear.NewDonorUp = totalRowNewDonorCountUp
	totLastYear.NewDonorDown = totalRowNewDonorCountDown

	report := domain.ReportResponse{
		Total:                     total,
		TotalRowCount:             totalRowCount,
		TotalRowDonorCount:        totalRowDonorCount,
		TotalRowNewDonorCount:     totalRowNewDonorCount,
		TotalZakatMaal:            totalZakatMaal,
		TotalWakaf:                totalWakaf,
		TotalZakatFitrah:          totalZakatFitrah,
		TotalInfaq:                totalInfaq,
		TotalKurban:               totalKurban,
		TotalOther:                totalOther,
		TotalRetail:               totalRetail,
		TotalRetailRowCount:       totalRetailRowCount,
		TotalCorporate:            totalCorporate,
		TotalCorporateRowCount:    totalCorporateRowCount,
		TotalUpz:                  totalUpz,
		TotalUpzRowCount:          totalUpzRowCount,
		TotalPrognosisRetail:      totalPrognosisRetail,
		TotalPrognosisCorporate:   totalPrognosisCorporate,
		TotalPrognosisUpz:         totalPrognosisUpz,
		TotalPercentRetail:        totalPercentRetail,
		TotalPercentCorporate:     totalPercentCorporate,
		TotalPercentUpz:           totalPercentUpz,
		TotalCompanyDonor:         totalCompanyDonor,
		TotalPersonDonor:          totalPersonDonor,
		TotalCompanyDonorRowCount: totalCompanyDonorRowCount,
		TotalPersonDonorRowCount:  totalPersonDonorRowCount,
		TotalPercentCompanyDonor:  totalPercentCompanyDonor,
		TotalPercentPersonDonor:   totalPercentPersonDonor,
		TotalRitelZakatMaal:       totalRitelZakatMaal,
		TotalRitelWakaf:           totalRitelWakaf,
		TotalRitelZakatFitrah:     totalRitelZakatFitrah,
		TotalRitelInfaq:           totalRitelInfaq,
		TotalRitelKurban:          totalRitelKurban,
		TotalRitelOther:           totalRitelOther,
		TotalCorporateZakatFitrah: totalCorporateZakatFitrah,
		TotalCorporateZakatMaal:   totalCorporateZakatMaal,
		TotalCorporateInfaq:       totalCorporateInfaq,
		TotalCorporateWakaf:       totalCorporateWakaf,
		TotalCorporateKurban:      totalCorporateKurban,
		TotalCorporateOther:       totalCorporateOther,
		TotalUpzZakatMaal:         totalUpzZakatMaal,
		TotalUpzWakaf:             totalUpzWakaf,
		TotalUpzZakatFitrah:       totalUpzZakatFitrah,
		TotalUpzInfaq:             totalUpzInfaq,
		TotalUpzKurban:            totalUpzKurban,
		TotalUpzOther:             totalUpzOther,
		TotalGood:                 totalGood,
		TotalGoodCount:            totalGoodCount,
		TotalGoodCollect:          totalGoodMoveCount + totalGoodNotMoveCount + totalGoodFoodCount + totalGoodOtherCount,
		TotalGoodMoveCount:        totalGoodMoveCount,
		TotalGoodNotMoveCount:     totalGoodNotMoveCount,
		TotalGoodFoodCount:        totalGoodFoodCount,
		TotalGoodOtherCount:       totalGoodOtherCount,
		TotalAllCash:              totalAllCash,
		TotalCashCount:            totalCashCount,
		TotalCash:                 totalCash,
		TotalNonCashMuamalat:      totalNonCashMuamalat,
		TotalNonCashMandiri:       totalNonCashMandiri,
		TotalNonCashBsm:           totalNonCashBsm,
		TotalNonCashBri:           totalNonCashBri,
		TotalNonCashBniLamp:       totalNonCashBniLamp,
		TotalNonCashBniSy:         totalNonCashBniSy,
		TotalNonCashBca:           totalNonCashBca,
		Prognosis:                 transGoal,
		TotalUpzPerMonth:          totUpzMonth,
		TotalRetailPerMonth:       totRetMonth,
		TotalCorporatePerMonth:    totCorMonth,
		TotalCashPercent:          percentCash,
		TotalPercentZiswaf:        percentZiswaf,
		TotalZiswafPerMonth:       totalZiswaf,
		TotalZakatMaalPerMonth:    totalZakMaal,
		TotalZakatFitrahPerMonth:  totalZakFit,
		TotalInfaqPerMonth:        totalInf,
		TotalWakafPerMonth:        totalWaq,
		TotalQurbanPerMonth:       totalQur,
		TotalOtherPerMonth:        totalOth,
		TotalPrognosisPerMonth:    totProgMonth,
		TotalTransactionPerMonth:  totTrxMonth,
		TotalTransactionLastYear:  totLastYear,
		TotalDivisionPercent:      totDivPercent,
	}

	return report, nil
}

func (p *transactionRepository) ShowTransaction(ctx context.Context, id string) (domain.Transaction, error) {
	trx := domain.Transaction{}

	if p.DB.Preload("Donor").
		Preload("School").
		Preload("School.Regency").
		Preload("Division").
		Preload("Category").
		Preload("Cashes").
		Preload("Cashes.CashCategory").
		Preload("GoodsQuery").
		Preload("GoodsQuery.GoodsCategory").
		Preload("StatementCategory").
		First(&trx, "id = ?", id).RecordNotFound() {
		return trx, errors.New("Cannot find transaction with id + " + id)
	}

	return trx, nil
}

func (p *transactionRepository) GetListReportOperator(ctx context.Context, filter domain.TransactionFilter) (domain.ReportOperatorResponse, error) {
	transaction := []domain.Transaction{}
	trxLastMonth := []domain.Transaction{}
	donor := []domain.Donor{}
	cash := []domain.Cash{}
	good := []domain.Goods{}
	totalZiswaf := domain.TotalZiswafPerDay{}
	totalZakMaal := domain.TotalZakatMaalPerDay{}
	totalZakFit := domain.TotalZakatFitrahPerDay{}
	totalInf := domain.TotalInfaqPerDay{}
	totalWaq := domain.TotalWakafPerDay{}
	totalQur := domain.TotalQurbanPerDay{}
	totalOth := domain.TotalOtherPerDay{}
	trxDay1 := []domain.Transaction{}
	trxDay2 := []domain.Transaction{}
	trxDay3 := []domain.Transaction{}
	trxDay4 := []domain.Transaction{}
	trxDay5 := []domain.Transaction{}
	trxDay6 := []domain.Transaction{}
	trxDay7 := []domain.Transaction{}
	trxDay8 := []domain.Transaction{}
	trxDay9 := []domain.Transaction{}
	trxDay10 := []domain.Transaction{}
	trxDay11 := []domain.Transaction{}
	trxDay12 := []domain.Transaction{}
	trxDay13 := []domain.Transaction{}
	trxDay14 := []domain.Transaction{}
	trxDay15 := []domain.Transaction{}
	trxDay16 := []domain.Transaction{}
	trxDay17 := []domain.Transaction{}
	trxDay18 := []domain.Transaction{}
	trxDay19 := []domain.Transaction{}
	trxDay20 := []domain.Transaction{}
	trxDay21 := []domain.Transaction{}
	trxDay22 := []domain.Transaction{}
	trxDay23 := []domain.Transaction{}
	trxDay24 := []domain.Transaction{}
	trxDay25 := []domain.Transaction{}
	trxDay26 := []domain.Transaction{}
	trxDay27 := []domain.Transaction{}
	trxDay28 := []domain.Transaction{}
	trxDay29 := []domain.Transaction{}
	trxDay30 := []domain.Transaction{}
	trxDay31 := []domain.Transaction{}
	totalUpzDay := domain.TotalUpzPerDay{}
	totalRetailDay := domain.TotalRetailPerDay{}
	totalCorporateDay := domain.TotalCorporatePerDay{}
	percentCash := domain.TotalCashPercent{}

	var (
		total                     uint64
		totalRowCount             int
		totalRowDonorCount        int
		totalRowNewDonorCount     int
		totalLast                 uint64
		totalRowCountLast         int
		totalRowDonorCountLast    int
		totalRowNewDonorCountLast int
		totalPercentUp            uint64
		totalPercentDown          uint64
		totalRowCountUp           int
		totalRowCountDown         int
		totalRowDonorCountUp      int
		totalRowDonorCountDown    int
		totalRowNewDonorCountUp   int
		totalRowNewDonorCountDown int
		totalZakatMaal            uint64
		totalWakaf                uint64
		totalZakatFitrah          uint64
		totalInfaq                uint64
		totalKurban               uint64
		totalOther                uint64
		totalRetail               uint64
		totalCorporate            uint64
		totalUpz                  uint64
		totalUpzZakatFitrah       uint64
		totalUpzZakatMaal         uint64
		totalUpzInfaq             uint64
		totalUpzWakaf             uint64
		totalUpzKurban            uint64
		totalUpzOther             uint64
		totalRitelZakatFitrah     uint64
		totalRitelZakatMaal       uint64
		totalRitelInfaq           uint64
		totalRitelWakaf           uint64
		totalRitelKurban          uint64
		totalRitelOther           uint64
		totalCorporateZakatFitrah uint64
		totalCorporateZakatMaal   uint64
		totalCorporateInfaq       uint64
		totalCorporateWakaf       uint64
		totalCorporateKurban      uint64
		totalCorporateOther       uint64
		totalCompanyDonorRowCount int
		totalCompanyDonor         uint64
		totalPersonDonorRowCount  int
		totalPersonDonor          uint64
		totalPercentCompanyDonor  uint64
		totalPercentPersonDonor   uint64
		totalGood                 uint64
		totalGoodCount            uint64
		totalGoodMoveCount        int32
		totalGoodNotMoveCount     int32
		totalGoodFoodCount        int32
		totalGoodOtherCount       int32
		totalAllCash              uint64
		totalCashCount            uint64
		totalCash                 uint64
		totalNonCashMuamalat      uint64
		totalNonCashMandiri       uint64
		totalNonCashBsm           uint64
		totalNonCashBri           uint64
		totalNonCashBniLamp       uint64
		totalNonCashBniSy         uint64
		totalNonCashBca           uint64
		totUpz                    domain.UpzDay
		totRetail                 domain.RetailDay
		totCorporate              domain.CorporateDay
		tot                       domain.ZiswafDay
		totZaMaal                 domain.ZakatMaalDay
		totZaFit                  domain.ZakatFitrahDay
		totInfaq                  domain.InfaqDay
		totKurban                 domain.QurbanDay
		totWaqaf                  domain.WaqafDay
		totOther                  domain.OtherDay
		totLastMonth              domain.TotalTransactionLastYear
		percentZiswaf             domain.TotalPercentZiswaf
		percentDiv                domain.TotalDivisionPercent
		dataTrx                   []domain.Transaction
		donorArrayID              []uint64
		donorLastArrayID          []uint64
	)

	if filter.SchoolID == "" {
		return domain.ReportOperatorResponse{}, errors.New("School ID is empty")
	} else if filter.StartDate == "" || filter.EndDate == "" {
		return domain.ReportOperatorResponse{}, errors.New("Date is empty")
	}

	// Get Month and Year
	monthStart := misc.TimeHelper(filter.StartDate, "month")
	monthEnd := misc.TimeHelper(filter.EndDate, "month")
	yearStart := misc.TimeHelper(filter.StartDate, "year")
	yearEnd := misc.TimeHelper(filter.EndDate, "year")

	if monthStart != monthEnd {
		return domain.ReportOperatorResponse{}, errors.New("Bulan pada filter harus pada bulan yang sama")
	}

	if yearStart != yearEnd {
		return domain.ReportOperatorResponse{}, errors.New("Tahun pada filter harus pada tahun yang sama")
	}

	// Get Date last Month
	dateLastMonth, dateEndLastMonth := misc.DateHelper(filter.StartDate, filter.EndDate, "month")

	tx := p.DB.Model(&transaction)
	ty := p.DB.Model(&trxLastMonth)

	if filter.SchoolID == "" {
		tx = tx.Where("id > ?", 0)
		ty = ty.Where("id > ?", 0)
	}
	if filter.StartDate != "" {
		tx = tx.Where("created_at >= ?", filter.StartDate)
		ty = ty.Where("created_at >= ?", dateLastMonth)
	}
	if filter.EndDate != "" {
		tx = tx.Where("created_at <= ?", filter.EndDate)
		ty = ty.Where("created_at <= ?", dateEndLastMonth)
	}
	if filter.SchoolID != "" {
		tx = tx.Where("school_id = ?", filter.SchoolID)
		ty = ty.Where("school_id = ?", filter.SchoolID)
	}

	tx.Preload("Donor").Find(&transaction).Count(&totalRowCount)
	ty.Find(&trxLastMonth).Count(&totalRowCountLast)

	// Total Transaction
	for _, v := range transaction {
		total += v.Total
	}

	if total == 0 {
		return domain.ReportOperatorResponse{}, errors.New("Data Not Found")
	}

	// Total Transaction Last Year
	for _, v := range trxLastMonth {
		totalLast += v.Total
	}

	// Get different total
	if totalLast == 0 {
		totalPercentUp = 100
		totalPercentDown = 0
	} else {
		totPercent := percent(totalLast, total)
		if totPercent > 100 {
			totalPercentDown = totPercent - 100
			totalPercentUp = 0
		} else if totPercent < 100 {
			totalPercentUp = 100 - totPercent
			totalPercentDown = 0
		}
	}

	// Get different row count
	if totalRowCount > totalRowCountLast {
		totalRowCountUp = totalRowCount - totalRowCountLast
		totalRowCountDown = 0
	} else if totalRowCount < totalRowCountLast {
		totalRowCountDown = totalRowCountLast - totalRowCount
		totalRowCountUp = 0
	}

	day := misc.DayHelper(filter.StartDate)

	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day1, day.DayEnd1, filter.SchoolID).Find(&trxDay1)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day2, day.DayEnd2, filter.SchoolID).Find(&trxDay2)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day3, day.DayEnd3, filter.SchoolID).Find(&trxDay3)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day4, day.DayEnd4, filter.SchoolID).Find(&trxDay4)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day5, day.DayEnd5, filter.SchoolID).Find(&trxDay5)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day6, day.DayEnd6, filter.SchoolID).Find(&trxDay6)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day7, day.DayEnd7, filter.SchoolID).Find(&trxDay7)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day8, day.DayEnd8, filter.SchoolID).Find(&trxDay8)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day9, day.DayEnd9, filter.SchoolID).Find(&trxDay9)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day10, day.DayEnd10, filter.SchoolID).Find(&trxDay10)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day11, day.DayEnd11, filter.SchoolID).Find(&trxDay11)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day12, day.DayEnd12, filter.SchoolID).Find(&trxDay12)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day13, day.DayEnd13, filter.SchoolID).Find(&trxDay13)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day14, day.DayEnd14, filter.SchoolID).Find(&trxDay14)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day15, day.DayEnd15, filter.SchoolID).Find(&trxDay15)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day16, day.DayEnd16, filter.SchoolID).Find(&trxDay16)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day17, day.DayEnd17, filter.SchoolID).Find(&trxDay17)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day18, day.DayEnd18, filter.SchoolID).Find(&trxDay18)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day19, day.DayEnd19, filter.SchoolID).Find(&trxDay19)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day20, day.DayEnd20, filter.SchoolID).Find(&trxDay20)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day21, day.DayEnd21, filter.SchoolID).Find(&trxDay21)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day22, day.DayEnd22, filter.SchoolID).Find(&trxDay22)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day23, day.DayEnd23, filter.SchoolID).Find(&trxDay23)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day24, day.DayEnd24, filter.SchoolID).Find(&trxDay24)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day25, day.DayEnd25, filter.SchoolID).Find(&trxDay25)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day26, day.DayEnd26, filter.SchoolID).Find(&trxDay26)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day27, day.DayEnd27, filter.SchoolID).Find(&trxDay27)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day28, day.DayEnd28, filter.SchoolID).Find(&trxDay28)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day29, day.DayEnd29, filter.SchoolID).Find(&trxDay29)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day30, day.DayEnd30, filter.SchoolID).Find(&trxDay30)
	p.DB.Where("created_at >= ? AND created_at <= ? AND school_id = ?", day.Day31, day.DayEnd31, filter.SchoolID).Find(&trxDay31)

	for i := 1; i < 32; i++ {
		switch i {
		case 1:
			dataTrx = trxDay1
		case 2:
			dataTrx = trxDay2
		case 3:
			dataTrx = trxDay3
		case 4:
			dataTrx = trxDay4
		case 5:
			dataTrx = trxDay5
		case 6:
			dataTrx = trxDay6
		case 7:
			dataTrx = trxDay7
		case 8:
			dataTrx = trxDay8
		case 9:
			dataTrx = trxDay9
		case 10:
			dataTrx = trxDay10
		case 11:
			dataTrx = trxDay11
		case 12:
			dataTrx = trxDay12
		case 13:
			dataTrx = trxDay13
		case 14:
			dataTrx = trxDay14
		case 15:
			dataTrx = trxDay15
		case 16:
			dataTrx = trxDay16
		case 17:
			dataTrx = trxDay17
		case 18:
			dataTrx = trxDay18
		case 19:
			dataTrx = trxDay19
		case 20:
			dataTrx = trxDay20
		case 21:
			dataTrx = trxDay21
		case 22:
			dataTrx = trxDay22
		case 23:
			dataTrx = trxDay23
		case 24:
			dataTrx = trxDay24
		case 25:
			dataTrx = trxDay25
		case 26:
			dataTrx = trxDay26
		case 27:
			dataTrx = trxDay27
		case 28:
			dataTrx = trxDay28
		case 29:
			dataTrx = trxDay29
		case 30:
			dataTrx = trxDay30
		case 31:
			dataTrx = trxDay31
		}

		// Total Ziswaf
		for _, v := range dataTrx {
			tot.Total += v.Total
			if v.CategoryID == 1 {
				totZaFit.Total += v.Total
				if v.DivisionID == 1 {
					totUpz.Total += v.Total
				} else if v.DivisionID == 2 {
					totRetail.Total += v.Total
				} else {
					totCorporate.Total += v.Total
				}
			} else if v.CategoryID == 2 {
				totZaMaal.Total += v.Total
				if v.DivisionID == 1 {
					totUpz.Total += v.Total
				} else if v.DivisionID == 2 {
					totRetail.Total += v.Total
				} else {
					totCorporate.Total += v.Total
				}
			} else if v.CategoryID == 3 {
				totInfaq.Total += v.Total
				if v.DivisionID == 1 {
					totUpz.Total += v.Total
				} else if v.DivisionID == 2 {
					totRetail.Total += v.Total
				} else {
					totCorporate.Total += v.Total
				}
			} else if v.CategoryID == 4 {
				totWaqaf.Total += v.Total
				if v.DivisionID == 1 {
					totUpz.Total += v.Total
				} else if v.DivisionID == 2 {
					totRetail.Total += v.Total
				} else {
					totCorporate.Total += v.Total
				}
			} else if v.CategoryID == 5 {
				totKurban.Total += v.Total
				if v.DivisionID == 1 {
					totUpz.Total += v.Total
				} else if v.DivisionID == 2 {
					totRetail.Total += v.Total
				} else {
					totCorporate.Total += v.Total
				}
			} else if v.CategoryID == 6 {
				totOther.Total += v.Total
				if v.DivisionID == 1 {
					totUpz.Total += v.Total
				} else if v.DivisionID == 2 {
					totRetail.Total += v.Total
				} else {
					totCorporate.Total += v.Total
				}
			}
		}
		totalZiswaf.Day = append(totalZiswaf.Day, tot)
		totalZakMaal.Day = append(totalZakMaal.Day, totZaMaal)
		totalZakFit.Day = append(totalZakFit.Day, totZaFit)
		totalInf.Day = append(totalInf.Day, totInfaq)
		totalWaq.Day = append(totalWaq.Day, totWaqaf)
		totalQur.Day = append(totalQur.Day, totKurban)
		totalOth.Day = append(totalOth.Day, totOther)
		totalUpzDay.Day = append(totalUpzDay.Day, totUpz)
		totalRetailDay.Day = append(totalRetailDay.Day, totRetail)
		totalCorporateDay.Day = append(totalCorporateDay.Day, totCorporate)
		tot.Total = 0
		totZaMaal.Total = 0
		totZaFit.Total = 0
		totInfaq.Total = 0
		totKurban.Total = 0
		totWaqaf.Total = 0
		totOther.Total = 0
		totUpz.Total = 0
		totRetail.Total = 0
		totCorporate.Total = 0
	}

	p.DB.Find(&cash)
	p.DB.Find(&good)

	for _, v := range transaction {
		for _, j := range cash {
			if v.ItemType == "cashs" && j.ID == v.ItemID {
				if j.CategoryID == 1 {
					totalCash += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 2 {
					totalNonCashMuamalat += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 3 {
					totalNonCashMandiri += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 4 {
					totalNonCashBsm += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 5 {
					totalNonCashBri += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 6 {
					totalNonCashBniLamp += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 7 {
					totalNonCashBniSy += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				} else if j.CategoryID == 8 {
					totalNonCashBca += v.Total
					totalAllCash += v.Total
					totalCashCount += 1
				}
				continue
			}
		}
	}

	for _, v := range transaction {
		for _, j := range good {
			if v.ItemType == "goods" && j.ID == v.ItemID {
				if j.CategoryID == 1 {
					totalGoodNotMoveCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				} else if j.CategoryID == 2 {
					totalGoodMoveCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				} else if j.CategoryID == 3 {
					totalGoodFoodCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				} else if j.CategoryID == 4 {
					totalGoodOtherCount += j.Quantity
					totalGood += v.Total
					totalGoodCount += 1
				}
				continue
			}
		}
	}

	// Division
	for _, v := range transaction {
		if v.DivisionID == 2 {
			totalRetail += v.Total
		} else if v.DivisionID == 3 {
			totalCorporate += v.Total
		} else if v.DivisionID == 1 {
			totalUpz += v.Total
		}
	}

	// Detail Per Division
	for _, v := range transaction {
		if v.DivisionID == 1 && v.CategoryID == 1 {
			totalUpzZakatFitrah += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 2 {
			totalUpzZakatMaal += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 3 {
			totalUpzInfaq += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 4 {
			totalUpzWakaf += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 5 {
			totalUpzKurban += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 6 {
			totalUpzOther += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 1 {
			totalRitelZakatFitrah += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 2 {
			totalRitelZakatMaal += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 3 {
			totalRitelInfaq += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 4 {
			totalRitelWakaf += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 5 {
			totalRitelKurban += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 6 {
			totalRitelOther += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 1 {
			totalCorporateZakatFitrah += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 2 {
			totalCorporateZakatMaal += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 3 {
			totalCorporateInfaq += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 4 {
			totalCorporateWakaf += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 5 {
			totalCorporateKurban += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 6 {
			totalCorporateOther += v.Total
		}
	}

	// Categories
	for _, v := range transaction {
		if v.CategoryID == 2 {
			totalZakatMaal += v.Total
		} else if v.CategoryID == 4 {
			totalWakaf += v.Total
		} else if v.CategoryID == 3 {
			totalInfaq += v.Total
		} else if v.CategoryID == 1 {
			totalZakatFitrah += v.Total
		} else if v.CategoryID == 5 {
			totalKurban += v.Total
		} else if v.CategoryID == 6 {
			totalOther += v.Total
		}
	}

	// Check Donor company
	for _, v := range transaction {
		if v.Donor.IsCompany == true {
			totalCompanyDonorRowCount += 1
			totalCompanyDonor += v.Total
		} else if v.Donor.IsCompany == false {
			totalPersonDonorRowCount += 1
			totalPersonDonor += v.Total
		}
	}

	time := time.Now()
	beginningOfMonth := misc.BeginningOfMonth(time)
	endOfMonth := misc.EndOfMonth(time)
	lastMonth, lastEndMonth := misc.DateHelper(beginningOfMonth.String(), endOfMonth.String(), "month")

	// Ger donor in transaction
	for _, v := range transaction {
		donorArrayID = append(donorArrayID, v.DonorID)
	}
	donorArrayID = unique(donorArrayID)

	for _, v := range trxLastMonth {
		donorLastArrayID = append(donorLastArrayID, v.DonorID)
	}
	donorLastArrayID = unique(donorLastArrayID)

	p.DB.Where("id IN (?)", donorArrayID).Find(&donor).Count(&totalRowDonorCount)
	p.DB.Where("id IN (?) AND created_at >= ? and created_at <= ?", donorArrayID, beginningOfMonth, endOfMonth).Find(&donor).Count(&totalRowNewDonorCount)
	p.DB.Where("id IN (?) AND created_at >= ? AND created_at <= ?", donorLastArrayID, dateLastMonth, dateEndLastMonth).Find(&donor).Count(&totalRowDonorCountLast)
	p.DB.Where("id IN (?) AND created_at >= ? and created_at <= ?", donorLastArrayID, lastMonth, lastEndMonth).Find(&donor).Count(&totalRowNewDonorCountLast)

	// Get different total row count donor
	if totalRowDonorCount > totalRowDonorCountLast {
		totalRowDonorCountUp = totalRowDonorCount - totalRowDonorCountLast
		totalRowDonorCountDown = 0
	} else if totalRowDonorCount < totalRowDonorCountLast {
		totalRowDonorCountDown = totalRowDonorCountLast - totalRowDonorCount
		totalRowDonorCountUp = 0
	}

	// Get different total row count new donor
	if totalRowNewDonorCount > totalRowNewDonorCountLast {
		totalRowNewDonorCountUp = totalRowNewDonorCount - totalRowNewDonorCountLast
		totalRowNewDonorCountDown = 0
	} else if totalRowNewDonorCount < totalRowNewDonorCountLast {
		totalRowNewDonorCountDown = totalRowNewDonorCountLast - totalRowNewDonorCount
		totalRowNewDonorCountUp = 0
	}

	// Get percent ziswaf
	totalPercentCompanyDonor = percent(totalCompanyDonor, total)
	totalPercentPersonDonor = percent(totalPersonDonor, total)
	percentZiswaf.ZakatMaalPercent = percent(totalZakatMaal, total)
	percentZiswaf.ZakatFitrahPercent = percent(totalZakatFitrah, total)
	percentZiswaf.WaqafPercent = percent(totalWakaf, total)
	percentZiswaf.QurbanPercent = percent(totalKurban, total)
	percentZiswaf.InfaqPercent = percent(totalInfaq, total)
	percentZiswaf.OtherPercent = percent(totalOther, total)
	percentCash.CashPercent = percent(totalCash, totalAllCash)
	percentCash.NonCashBcaPercent = percent(totalNonCashBca, totalAllCash)
	percentCash.NonCashBniLampPercent = percent(totalNonCashBniLamp, totalAllCash)
	percentCash.NonCashBniSyPercent = percent(totalNonCashBniSy, totalAllCash)
	percentCash.NonCashBriPercent = percent(totalNonCashBri, totalAllCash)
	percentCash.NonCashBsmPercent = percent(totalNonCashBsm, totalAllCash)
	percentCash.NonCashMandiriPercent = percent(totalNonCashMandiri, totalAllCash)
	percentCash.NonCashMuamalatPercent = percent(totalNonCashMuamalat, totalAllCash)
	percentDiv.TotalDivisionCorporatePercent = percent(totalCorporate, total)
	percentDiv.TotalDivisionRetailPercent = percent(totalRetail, total)
	percentDiv.TotalDivisionUpzPercent = percent(totalUpz, total)

	totLastMonth.TotalUp = totalPercentUp
	totLastMonth.TotalDown = totalPercentDown
	totLastMonth.CountUp = totalRowCountUp
	totLastMonth.CountDown = totalRowCountDown
	totLastMonth.DonorUp = totalRowDonorCountUp
	totLastMonth.DonorDown = totalRowDonorCountDown
	totLastMonth.NewDonorUp = totalRowNewDonorCountUp
	totLastMonth.NewDonorDown = totalRowNewDonorCountDown

	report := domain.ReportOperatorResponse{
		Total:                     total,
		TotalRowCount:             totalRowCount,
		TotalRowDonorCount:        totalRowDonorCount,
		TotalRowNewDonorCount:     totalRowNewDonorCount,
		TotalZakatMaal:            totalZakatMaal,
		TotalWakaf:                totalWakaf,
		TotalZakatFitrah:          totalZakatFitrah,
		TotalInfaq:                totalInfaq,
		TotalKurban:               totalKurban,
		TotalOther:                totalOther,
		TotalZiswafPerDay:         totalZiswaf,
		TotalZakatMaalPerDay:      totalZakMaal,
		TotalZakatFitrahPerDay:    totalZakFit,
		TotalInfaqPerDay:          totalInf,
		TotalWakafPerDay:          totalWaq,
		TotalQurbanPerDay:         totalQur,
		TotalOtherPerDay:          totalOth,
		TotalTransactionLastMonth: totLastMonth,
		TotalDivisionPercent:      percentDiv,
		TotalPercentZiswaf:        percentZiswaf,
		TotalGood:                 totalGood,
		TotalGoodCount:            totalGoodCount,
		TotalGoodCollect:          totalGoodMoveCount + totalGoodNotMoveCount + totalGoodFoodCount + totalGoodOtherCount,
		TotalGoodMoveCount:        totalGoodMoveCount,
		TotalGoodNotMoveCount:     totalGoodNotMoveCount,
		TotalGoodFoodCount:        totalGoodFoodCount,
		TotalGoodOtherCount:       totalGoodOtherCount,
		TotalAllCash:              totalAllCash,
		TotalCashCount:            totalCashCount,
		TotalCash:                 totalCash,
		TotalNonCashMuamalat:      totalNonCashMuamalat,
		TotalNonCashMandiri:       totalNonCashMandiri,
		TotalNonCashBsm:           totalNonCashBsm,
		TotalNonCashBri:           totalNonCashBri,
		TotalNonCashBniLamp:       totalNonCashBniLamp,
		TotalNonCashBniSy:         totalNonCashBniSy,
		TotalNonCashBca:           totalNonCashBca,
		TotalCashPercent:          percentCash,
		TotalCompanyDonor:         totalCompanyDonor,
		TotalPersonDonor:          totalPersonDonor,
		TotalCompanyDonorRowCount: totalCompanyDonorRowCount,
		TotalPersonDonorRowCount:  totalPersonDonorRowCount,
		TotalPercentCompanyDonor:  totalPercentCompanyDonor,
		TotalPercentPersonDonor:   totalPercentPersonDonor,
		TotalUpzPerDay:            totalUpzDay,
		TotalRetailPerDay:         totalRetailDay,
		TotalCorporatePerDay:      totalCorporateDay,
		TotalRetail:               totalRetail,
		TotalCorporate:            totalCorporate,
		TotalUpz:                  totalUpz,
		TotalUpzZakatMaal:         totalUpzZakatMaal,
		TotalUpzWakaf:             totalUpzWakaf,
		TotalUpzZakatFitrah:       totalUpzZakatFitrah,
		TotalUpzInfaq:             totalUpzInfaq,
		TotalUpzKurban:            totalUpzKurban,
		TotalUpzOther:             totalUpzOther,
		TotalRitelZakatFitrah:     totalRitelZakatFitrah,
		TotalRitelZakatMaal:       totalRitelZakatMaal,
		TotalRitelInfaq:           totalRitelInfaq,
		TotalRitelWakaf:           totalRitelWakaf,
		TotalRitelKurban:          totalRitelKurban,
		TotalRitelOther:           totalRitelOther,
		TotalCorporateZakatFitrah: totalCorporateZakatFitrah,
		TotalCorporateZakatMaal:   totalCorporateZakatMaal,
		TotalCorporateInfaq:       totalCorporateInfaq,
		TotalCorporateWakaf:       totalCorporateWakaf,
		TotalCorporateKurban:      totalCorporateKurban,
		TotalCorporateOther:       totalCorporateOther,
	}

	return report, nil
}

func (p *transactionRepository) GetListDonationReport(ctx context.Context, filter domain.TransactionFilter, cat []domain.StatementCategory) (domain.ReportStatementCategory, error) {
	transaction := []domain.Transaction{}
	trx := []domain.Transaction{}
	donor := []domain.Donor{}
	school := []domain.School{}
	cash := []domain.Cash{}
	statCat := []domain.StatementCategory{}
	rpDonation := domain.ReportStatementCategory{}
	rpDonData := domain.ReportStatementCategoryData{}

	var (
		donorArrayID  []uint64
		schoolArrayID []uint64
		cashArrayID   []uint64
		staArray      []uint64
		totalRowCount int
		total         uint64
		totalData     uint64
		idCat         uint64
		nameCat       string
	)

	staArray = misc.QueryConvert(filter.StatementCategory)

	if filter.StatementCategory != "" {
		for _, v := range staArray {
			p.DB.Where("statement_category_id = ?", v).Find(&trx)
			if len(trx) == 0 {
				catID := strconv.FormatUint(v, 10)
				return domain.ReportStatementCategory{}, errors.New("Keterangan Donasi dengan id " + catID + " tidak ditemukan")
			}
		}
	}

	tx := p.DB.Model(&transaction)

	if filter.SchoolID == "" {
		tx = tx.Where("id > ?", 0)
	}
	if filter.StatementCategory == "" {
		staArray = []uint64{}
		for _, v := range cat {
			staArray = append(staArray, v.ID)
		}
		tx = tx.Where("statement_category_id IN (?)", staArray)
	}
	if filter.DivisionID != "" {
		divArrayID := misc.QueryConvert(filter.DivisionID)
		tx = tx.Where("division_id IN (?)", divArrayID)
	}
	if filter.CategoryID != "" {
		catArrayID := misc.QueryConvert(filter.CategoryID)
		tx = tx.Where("category_id IN (?)", catArrayID)
	}
	if filter.StatementCategory != "" {
		tx = tx.Where("statement_category_id IN (?)", staArray)
	}
	if filter.DonaturCategory != "" {
		donArrayID := misc.QueryConvert(filter.DonaturCategory)
		p.DB.Where("is_company IN (?)", donArrayID).Find(&donor)
		for _, v := range donor {
			donorArrayID = append(donorArrayID, v.ID)
		}
		tx = tx.Where("donor_id IN (?)", donorArrayID)
	}
	if filter.CategoryType != "" {
		catArrayID := misc.QueryConvert(filter.CategoryType)
		p.DB.Where("category_id IN (?)", catArrayID).Find(&cash)
		for _, v := range cash {
			cashArrayID = append(cashArrayID, v.ID)
		}
		tx = tx.Where("item_id IN (?) AND item_type = ?", cashArrayID, "cashs")
	}
	if filter.Regency != "" {
		regArrayID := misc.QueryConvert(filter.Regency)
		p.DB.Where("regency_id IN (?)", regArrayID).Find(&school)
		for _, v := range school {
			schoolArrayID = append(schoolArrayID, v.ID)
		}
		tx = tx.Where("school_id IN (?)", schoolArrayID)
	}
	if filter.StartDate != "" {
		tx = tx.Where("created_at >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		tx = tx.Where("created_at <= ?", filter.EndDate)
	}
	if filter.SchoolID != "" {
		schArrayID := misc.QueryConvert(filter.SchoolID)
		tx = tx.Where("school_id IN (?)", schArrayID)
	}

	tx.Order("statement_category_id ASC").Find(&transaction).Count(&totalRowCount)

	if len(transaction) == 0 {
		return domain.ReportStatementCategory{}, errors.New("Data Not Found")
	}

	p.DB.Find(&statCat)

	// Total Transaction
	for _, v := range transaction {
		total += v.Total
	}

	for i := 0; i < len(unique(staArray)); i++ {
		for _, v := range transaction {
			if staArray[i] == v.StatementCategoryID {
				idCat = v.StatementCategoryID
				totalData += v.Total
			} else {
				idCat = staArray[i]
			}
		}

		for _, j := range statCat {
			if idCat == j.ID {
				nameCat = j.Name
			}
		}

		rpDonData.TotalPercent = percent(totalData, total)
		rpDonData.Total = totalData
		rpDonData.StatementCategoryID = idCat
		rpDonData.Name = nameCat
		rpDonation.Data = append(rpDonation.Data, rpDonData)
		rpDonData.StatementCategoryID = 0
		rpDonData.TotalPercent = 0
		totalData = 0

	}

	rpDonation.TotalCash = total
	rpDonation.TotalRowCount = totalRowCount

	return rpDonation, nil
}

func (p *transactionRepository) ExportPdf(ctx context.Context, filter domain.TransactionFilter) (domain.ExportPdf, error) {
	transaction := []domain.Transaction{}
	transGoal := []domain.TransactionGoal{}
	dataTrx := []domain.Transaction{}
	trxJan := []domain.Transaction{}
	trxFeb := []domain.Transaction{}
	trxMar := []domain.Transaction{}
	trxApr := []domain.Transaction{}
	trxMay := []domain.Transaction{}
	trxJun := []domain.Transaction{}
	trxJul := []domain.Transaction{}
	trxAug := []domain.Transaction{}
	trxSep := []domain.Transaction{}
	trxOct := []domain.Transaction{}
	trxNov := []domain.Transaction{}
	trxDes := []domain.Transaction{}
	exPdf := domain.ExportPdf{}
	exPdfData := domain.ExportPdfData{}
	exPdfDetail := domain.ExportPdfDetailData{}
	exPdfCat := domain.ExportPdfPerCategory{}
	exPdfCatDet := domain.ExportPdfPerCategoryDetail{}
	exPdfDiv := domain.ExportPdfPerDivision{}
	school := []domain.School{}
	sch := domain.School{}
	regency := domain.Regency{}

	var (
		total           uint64
		totalRowCount   int
		schoolArrayID   []uint64
		totGap          uint64
		totGapUpz       uint64
		totGapRetail    uint64
		totGapCorporate uint64
	)

	if filter.StartDate == "" || filter.EndDate == "" {
		return domain.ExportPdf{}, errors.New("Date is empty")
	}

	// Get Year
	year := misc.TimeHelper(filter.StartDate, "year")
	month := misc.MonthHelper(year)

	tx := p.DB.Model(&transaction)
	ta := p.DB.Model(&trxJan)
	tb := p.DB.Model(&trxFeb)
	tc := p.DB.Model(&trxMar)
	td := p.DB.Model(&trxApr)
	te := p.DB.Model(&trxMay)
	tf := p.DB.Model(&trxJun)
	tg := p.DB.Model(&trxJul)
	th := p.DB.Model(&trxAug)
	ti := p.DB.Model(&trxSep)
	tj := p.DB.Model(&trxOct)
	tk := p.DB.Model(&trxNov)
	tl := p.DB.Model(&trxDes)

	if filter.SchoolID == "" {
		tx = tx.Where("id > ?", 0)
		ta = ta.Where("id > ?", 0)
		tb = tb.Where("id > ?", 0)
		tc = tc.Where("id > ?", 0)
		td = td.Where("id > ?", 0)
		te = te.Where("id > ?", 0)
		tf = tf.Where("id > ?", 0)
		tg = tg.Where("id > ?", 0)
		th = th.Where("id > ?", 0)
		ti = ti.Where("id > ?", 0)
		tj = tj.Where("id > ?", 0)
		tk = tk.Where("id > ?", 0)
		tl = tl.Where("id > ?", 0)
	}

	if filter.StartDate != "" {
		tx = tx.Where("created_at >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		tx = tx.Where("created_at <= ?", filter.EndDate)
	}
	if filter.Regency != "" {
		p.DB.Where("id = ?", filter.Regency).Find(&regency)
		p.DB.Where("regency_id = ?", filter.Regency).Find(&school)
		for _, v := range school {
			schoolArrayID = append(schoolArrayID, v.ID)
		}
		tx = tx.Where("school_id IN (?)", schoolArrayID)
		ta = ta.Where("school_id IN (?)", schoolArrayID)
		tb = tb.Where("school_id IN (?)", schoolArrayID)
		tc = tc.Where("school_id IN (?)", schoolArrayID)
		td = td.Where("school_id IN (?)", schoolArrayID)
		te = te.Where("school_id IN (?)", schoolArrayID)
		tf = tf.Where("school_id IN (?)", schoolArrayID)
		tg = tg.Where("school_id IN (?)", schoolArrayID)
		th = th.Where("school_id IN (?)", schoolArrayID)
		ti = ti.Where("school_id IN (?)", schoolArrayID)
		tj = tj.Where("school_id IN (?)", schoolArrayID)
		tk = tk.Where("school_id IN (?)", schoolArrayID)
		tl = tl.Where("school_id IN (?)", schoolArrayID)
	}
	if filter.SchoolID != "" {
		tx = tx.Where("school_id = ?", filter.SchoolID)
		ta = ta.Where("school_id = ?", filter.SchoolID)
		tb = tb.Where("school_id = ?", filter.SchoolID)
		tc = tc.Where("school_id = ?", filter.SchoolID)
		td = td.Where("school_id = ?", filter.SchoolID)
		te = te.Where("school_id = ?", filter.SchoolID)
		tf = tf.Where("school_id = ?", filter.SchoolID)
		tg = tg.Where("school_id = ?", filter.SchoolID)
		th = th.Where("school_id = ?", filter.SchoolID)
		ti = ti.Where("school_id = ?", filter.SchoolID)
		tj = tj.Where("school_id = ?", filter.SchoolID)
		tk = tk.Where("school_id = ?", filter.SchoolID)
		tl = tl.Where("school_id = ?", filter.SchoolID)

		p.DB.Where("id = ?", filter.SchoolID).Find(&sch)
	}

	tx.Preload("Donor").Find(&transaction).Count(&totalRowCount)
	ta.Where("created_at >= ? AND created_at <= ?", month.JanuaryStart, month.JanuaryEnd).Find(&trxJan)
	tb.Where("created_at >= ? AND created_at <= ?", month.FebruaryStart, month.FebruaryEnd).Find(&trxFeb)
	tc.Where("created_at >= ? AND created_at <= ?", month.MarchStart, month.MarchEnd).Find(&trxMar)
	td.Where("created_at >= ? AND created_at <= ?", month.AprilStart, month.AprilEnd).Find(&trxApr)
	te.Where("created_at >= ? AND created_at <= ?", month.MayStart, month.MayEnd).Find(&trxMay)
	tf.Where("created_at >= ? AND created_at <= ?", month.JuneStart, month.JuneEnd).Find(&trxJun)
	tg.Where("created_at >= ? AND created_at <= ?", month.JulyStart, month.JulyEnd).Find(&trxJul)
	th.Where("created_at >= ? AND created_at <= ?", month.AugustStart, month.AugustEnd).Find(&trxAug)
	ti.Where("created_at >= ? AND created_at <= ?", month.SeptemberStart, month.SeptemberEnd).Find(&trxSep)
	tj.Where("created_at >= ? AND created_at <= ?", month.OctoberStart, month.OctoberEnd).Find(&trxOct)
	tk.Where("created_at >= ? AND created_at <= ?", month.NovemberStart, month.NovemberEnd).Find(&trxNov)
	tl.Where("created_at >= ? AND created_at <= ?", month.DecemberStart, month.DecemberEnd).Find(&trxDes)

	// Total Transaction
	for _, v := range transaction {
		total += v.Total
	}

	// Total Ziswaf Per Month
	for i := 1; i < 14; i++ {
		switch i {
		case 1:
			dataTrx = trxJan
		case 2:
			dataTrx = trxFeb
		case 3:
			dataTrx = trxMar
		case 4:
			dataTrx = trxApr
		case 5:
			dataTrx = trxMay
		case 6:
			dataTrx = trxJun
		case 7:
			dataTrx = trxJul
		case 8:
			dataTrx = trxAug
		case 9:
			dataTrx = trxSep
		case 10:
			dataTrx = trxOct
		case 11:
			dataTrx = trxNov
		case 12:
			dataTrx = trxDes
		case 13:
			dataTrx = transaction
		}

		for _, v := range dataTrx {
			exPdfData.TotalRowCount += 1
			if v.DivisionID == 1 {
				exPdfData.TotalRowCountUpz += 1
				if v.ItemType == "cashs" {
					exPdfData.TotalCash += v.Total
					exPdfData.TotalCashUpz += v.Total
				} else if v.ItemType == "goods" {
					exPdfData.TotalGood += v.Total
					exPdfData.TotalGoodUpz += v.Total
				}
			} else if v.DivisionID == 2 {
				exPdfData.TotalRowCountRetail += 1
				if v.ItemType == "cashs" {
					exPdfData.TotalCash += v.Total
					exPdfData.TotalCashRetail += v.Total
				} else if v.ItemType == "goods" {
					exPdfData.TotalGood += v.Total
					exPdfData.TotalGoodRetail += v.Total
				}
			} else if v.DivisionID == 3 {
				exPdfData.TotalRowCountCorporate += 1
				if v.ItemType == "cashs" {
					exPdfData.TotalCash += v.Total
					exPdfData.TotalCashCorporate += v.Total
				} else if v.ItemType == "goods" {
					exPdfData.TotalGood += v.Total
					exPdfData.TotalGoodCorporate += v.Total
				}
			}
		}

		// Get Prognosis
		if i != 13 {
			p.DB.Where("year = ? AND month = ?", year, i).Find(&transGoal)
		} else {
			p.DB.Where("year = ?", year).Find(&transGoal)
		}

		for _, j := range transGoal {
			exPdfData.TotalPrognosis += j.Total
			if j.DivisionID == 1 {
				exPdfData.TotalPrognosisUpz = j.Total
			} else if j.DivisionID == 2 {
				exPdfData.TotalPrognosisRetail = j.Total
			} else if j.DivisionID == 3 {
				exPdfData.TotalPrognosisCorporate = j.Total
			}
		}

		exPdfData.TotalPercentUpz = percent((exPdfData.TotalCashUpz + exPdfData.TotalGoodUpz), exPdfData.TotalPrognosisUpz)
		exPdfData.TotalPercentRetail = percent((exPdfData.TotalCashRetail + exPdfData.TotalGoodRetail), exPdfData.TotalPrognosisRetail)
		exPdfData.TotalPercentCorporate = percent((exPdfData.TotalCashCorporate + exPdfData.TotalGoodCorporate), exPdfData.TotalPrognosisCorporate)
		exPdfData.TotalPercent = exPdfData.TotalPercentUpz + exPdfData.TotalPercentRetail + exPdfData.TotalPercentCorporate

		totProg := exPdfData.TotalPrognosis
		totProgUpz := exPdfData.TotalPrognosisUpz
		totProgRetail := exPdfData.TotalPrognosisRetail
		totProgCorp := exPdfData.TotalPrognosisCorporate

		if totProg > (exPdfData.TotalCash + exPdfData.TotalGood) {
			totGap = totProg - (exPdfData.TotalCash + exPdfData.TotalGood)
		} else {
			totGap = (exPdfData.TotalCash + exPdfData.TotalGood) - totProg
		}
		if totProgUpz > (exPdfData.TotalCashUpz + exPdfData.TotalGoodUpz) {
			totGapUpz = totProgUpz - (exPdfData.TotalCashUpz + exPdfData.TotalGoodUpz)
		} else {
			totGapUpz = (exPdfData.TotalCashUpz + exPdfData.TotalGoodUpz) - totProgUpz
		}
		if totProgRetail > (exPdfData.TotalCashRetail + exPdfData.TotalGoodRetail) {
			totGapRetail = totProgRetail - (exPdfData.TotalCashRetail + exPdfData.TotalGoodRetail)
		} else {
			totGapRetail = (exPdfData.TotalCashRetail + exPdfData.TotalGoodRetail) - totProgRetail
		}
		if totProgCorp > (exPdfData.TotalCashCorporate + exPdfData.TotalGoodCorporate) {
			totGapCorporate = totProgCorp - (exPdfData.TotalCashCorporate + exPdfData.TotalGoodCorporate)
		} else {
			totGapCorporate = (exPdfData.TotalCashCorporate + exPdfData.TotalGoodCorporate) - totProgCorp
		}

		exPdfData.TotalGap = totGap
		exPdfData.TotalGapUpz = totGapUpz
		exPdfData.TotalGapRetail = totGapRetail
		exPdfData.TotalGapCorporate = totGapCorporate

		var month string
		for j := 0; j < 5; j++ {
			if i == 1 {
				month = "Januari"
			} else if i == 2 {
				month = "Februari"
			} else if i == 3 {
				month = "Maret"
			} else if i == 4 {
				month = "April"
			} else if i == 5 {
				month = "Mei"
			} else if i == 6 {
				month = "Juni"
			} else if i == 7 {
				month = "Juli"
			} else if i == 8 {
				month = "Agustus"
			} else if i == 9 {
				month = "September"
			} else if i == 10 {
				month = "Oktober"
			} else if i == 11 {
				month = "November"
			} else if i == 12 {
				month = "Desember"
			} else if i == 13 {
				month = "Total"
			}
			if j == 0 {
				exPdfDetail.Month = month
				exPdfDetail.Name = "Upz"
				exPdfDetail.Total = exPdfData.TotalCashUpz + exPdfData.TotalGoodUpz
				exPdfDetail.TotalGap = exPdfData.TotalGapUpz
				exPdfDetail.TotalPercent = exPdfData.TotalPercentUpz
				exPdfDetail.TotalPrognosis = exPdfData.TotalPrognosisUpz
				exPdfDetail.TotalRowCount = exPdfData.TotalRowCountUpz

				exPdf.Data = append(exPdf.Data, exPdfDetail)
			} else if j == 1 {
				exPdfDetail.Month = month
				exPdfDetail.Name = "Retail"
				exPdfDetail.Total = exPdfData.TotalCashRetail + exPdfData.TotalGoodRetail
				exPdfDetail.TotalGap = exPdfData.TotalGapRetail
				exPdfDetail.TotalPercent = exPdfData.TotalPercentRetail
				exPdfDetail.TotalPrognosis = exPdfData.TotalPrognosisRetail
				exPdfDetail.TotalRowCount = exPdfData.TotalRowCountRetail

				exPdf.Data = append(exPdf.Data, exPdfDetail)
			} else if j == 2 {
				exPdfDetail.Month = month
				exPdfDetail.Name = "Corporate"
				exPdfDetail.Total = exPdfData.TotalCashCorporate + exPdfData.TotalGoodCorporate
				exPdfDetail.TotalGap = exPdfData.TotalGapCorporate
				exPdfDetail.TotalPercent = exPdfData.TotalPercentCorporate
				exPdfDetail.TotalPrognosis = exPdfData.TotalPrognosisCorporate
				exPdfDetail.TotalRowCount = exPdfData.TotalRowCountCorporate

				exPdf.Data = append(exPdf.Data, exPdfDetail)
			} else if j == 3 {
				exPdfDetail.Month = month
				exPdfDetail.Name = "Total"
				exPdfDetail.Total = exPdfData.TotalCash + exPdfData.TotalGood
				exPdfDetail.TotalGap = exPdfData.TotalGap
				exPdfDetail.TotalPercent = exPdfData.TotalPercent
				exPdfDetail.TotalPrognosis = exPdfData.TotalPrognosis
				exPdfDetail.TotalRowCount = exPdfData.TotalRowCount

				exPdf.Data = append(exPdf.Data, exPdfDetail)
			} else if j == 4 {
				exPdfDetail.Month = month
				exPdfDetail.Name = "Total"
				exPdfDetail.Total = exPdfData.TotalCash + exPdfData.TotalGood
				exPdfDetail.TotalGap = exPdfData.TotalGap
				exPdfDetail.TotalPercent = exPdfData.TotalPercent
				exPdfDetail.TotalPrognosis = exPdfData.TotalPrognosis
				exPdfDetail.TotalRowCount = exPdfData.TotalRowCount
			}
		}

		exPdfData.TotalCash = 0
		exPdfData.TotalCashCorporate = 0
		exPdfData.TotalCashRetail = 0
		exPdfData.TotalCashUpz = 0
		exPdfData.TotalGood = 0
		exPdfData.TotalGoodCorporate = 0
		exPdfData.TotalGoodRetail = 0
		exPdfData.TotalGoodUpz = 0
		exPdfData.TotalPercent = 0
		exPdfData.TotalPercentCorporate = 0
		exPdfData.TotalPercentRetail = 0
		exPdfData.TotalPercentUpz = 0
		exPdfData.TotalPrognosis = 0
		exPdfData.TotalPrognosisCorporate = 0
		exPdfData.TotalPrognosisRetail = 0
		exPdfData.TotalPrognosisUpz = 0
		exPdfData.TotalRowCount = 0
		exPdfData.TotalRowCountCorporate = 0
		exPdfData.TotalRowCountRetail = 0
		exPdfData.TotalRowCountUpz = 0
	}

	exPdf.Role = filter.Role
	exPdf.SchoolName = sch.Name
	exPdf.RegencyName = regency.Name
	exPdf.StartDate = filter.StartDate
	exPdf.EndDate = filter.EndDate

	var (
		totDon    uint64
		totMaal   uint64
		totFitrah uint64
		totWaqaf  uint64
		totQurban uint64
		totInfaq  uint64
		totOther  uint64
	)

	for _, v := range transaction {
		totDon += v.Total
		if v.CategoryID == 1 {
			totFitrah += v.Total
			if v.Donor.IsCompany == true {
				exPdfCat.TotalRowCountCorporate += 1
				exPdfCat.TotalCorporate += v.Total
			} else if v.Donor.IsCompany == false {
				exPdfCat.TotalPersonal += v.Total
				exPdfCat.TotalRowCountPersonal += 1
			}
		} else if v.CategoryID == 2 {
			totMaal += v.Total
			if v.Donor.IsCompany == true {
				exPdfCat.TotalRowCountCorporate += 1
				exPdfCat.TotalCorporate += v.Total
			} else if v.Donor.IsCompany == false {
				exPdfCat.TotalPersonal += v.Total
				exPdfCat.TotalRowCountPersonal += 1
			}
		} else if v.CategoryID == 3 {
			totInfaq += v.Total
			if v.Donor.IsCompany == true {
				exPdfCat.TotalRowCountCorporate += 1
				exPdfCat.TotalCorporate += v.Total
			} else if v.Donor.IsCompany == false {
				exPdfCat.TotalPersonal += v.Total
				exPdfCat.TotalRowCountPersonal += 1
			}
		} else if v.CategoryID == 4 {
			totWaqaf += v.Total
			if v.Donor.IsCompany == true {
				exPdfCat.TotalRowCountCorporate += 1
				exPdfCat.TotalCorporate += v.Total
			} else if v.Donor.IsCompany == false {
				exPdfCat.TotalPersonal += v.Total
				exPdfCat.TotalRowCountPersonal += 1
			}
		} else if v.CategoryID == 5 {
			totQurban += v.Total
			if v.Donor.IsCompany == true {
				exPdfCat.TotalRowCountCorporate += 1
				exPdfCat.TotalCorporate += v.Total
			} else if v.Donor.IsCompany == false {
				exPdfCat.TotalPersonal += v.Total
				exPdfCat.TotalRowCountPersonal += 1
			}
		} else if v.CategoryID == 6 {
			totOther += v.Total
			if v.Donor.IsCompany == true {
				exPdfCat.TotalRowCountCorporate += 1
				exPdfCat.TotalCorporate += v.Total
			} else if v.Donor.IsCompany == false {
				exPdfCat.TotalPersonal += v.Total
				exPdfCat.TotalRowCountPersonal += 1
			}
		}
	}

	for i := 0; i < 7; i++ {
		if i == 0 {
			exPdfCatDet.Name = "Zakat Fitrah"
			exPdfCatDet.Total = totFitrah
			exPdfCatDet.Percent = percent(totFitrah, total)
		} else if i == 1 {
			exPdfCatDet.Name = "Zakat Maal"
			exPdfCatDet.Total = totMaal
			exPdfCatDet.Percent = percent(totMaal, total)
		} else if i == 2 {
			exPdfCatDet.Name = "Infak/Sedekah"
			exPdfCatDet.Total = totInfaq
			exPdfCatDet.Percent = percent(totInfaq, total)
		} else if i == 3 {
			exPdfCatDet.Name = "Wakaf"
			exPdfCatDet.Total = totWaqaf
			exPdfCatDet.Percent = percent(totWaqaf, total)
		} else if i == 4 {
			exPdfCatDet.Name = "Qurban"
			exPdfCatDet.Total = totQurban
			exPdfCatDet.Percent = percent(totQurban, total)
		} else if i == 5 {
			exPdfCatDet.Name = "Penerimaan Lainnya"
			exPdfCatDet.Total = totOther
			exPdfCatDet.Percent = percent(totOther, total)
		} else if i == 6 {
			exPdfCatDet.Name = "Total"
			exPdfCatDet.Total = totDon
			exPdfCatDet.Percent = percent(totDon, total)
		}

		exPdfCat.TotalPercentCorporate = percent(exPdfCat.TotalCorporate, total)
		exPdfCat.TotalPercentPersonal = percent(exPdfCat.TotalPersonal, total)

		exPdfCat.ExportPdfPerCategoryDetail = append(exPdfCat.ExportPdfPerCategoryDetail, exPdfCatDet)
		exPdfCatDet.Percent = 0
		exPdfCatDet.Total = 0

		exPdf.ReportPerCategory = exPdfCat
	}

	var (
		dataPdf domain.ReportResponse
	)

	// Detail Per Division
	for _, v := range transaction {
		if v.DivisionID == 1 && v.CategoryID == 1 {
			dataPdf.TotalUpzZakatFitrah += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 2 {
			dataPdf.TotalUpzZakatMaal += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 3 {
			dataPdf.TotalUpzInfaq += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 4 {
			dataPdf.TotalUpzWakaf += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 5 {
			dataPdf.TotalUpzKurban += v.Total
		} else if v.DivisionID == 1 && v.CategoryID == 6 {
			dataPdf.TotalUpzOther += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 1 {
			dataPdf.TotalRitelZakatFitrah += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 2 {
			dataPdf.TotalRitelZakatMaal += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 3 {
			dataPdf.TotalRitelInfaq += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 4 {
			dataPdf.TotalRitelWakaf += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 5 {
			dataPdf.TotalRitelKurban += v.Total
		} else if v.DivisionID == 2 && v.CategoryID == 6 {
			dataPdf.TotalRitelOther += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 1 {
			dataPdf.TotalCorporateZakatFitrah += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 2 {
			dataPdf.TotalCorporateZakatMaal += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 3 {
			dataPdf.TotalCorporateInfaq += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 4 {
			dataPdf.TotalCorporateWakaf += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 5 {
			dataPdf.TotalCorporateKurban += v.Total
		} else if v.DivisionID == 3 && v.CategoryID == 6 {
			dataPdf.TotalCorporateOther += v.Total
		}
	}

	for i := 0; i < 19; i++ {
		if i == 0 {
			exPdfDiv.Name = "Ritel - Zakat Maal"
			exPdfDiv.Total = dataPdf.TotalRitelZakatMaal
		} else if i == 1 {
			exPdfDiv.Name = "Ritel - Zakat Fitrah"
			exPdfDiv.Total = dataPdf.TotalRitelZakatFitrah
		} else if i == 2 {
			exPdfDiv.Name = "Ritel - Infaq/Sedekah"
			exPdfDiv.Total = dataPdf.TotalRitelInfaq
		} else if i == 3 {
			exPdfDiv.Name = "Ritel - Waqaf"
			exPdfDiv.Total = dataPdf.TotalRitelWakaf
		} else if i == 4 {
			exPdfDiv.Name = "Ritel - Qurban"
			exPdfDiv.Total = dataPdf.TotalRitelKurban
		} else if i == 5 {
			exPdfDiv.Name = "Ritel - Penerimaan Lain "
			exPdfDiv.Total = dataPdf.TotalRitelOther
		} else if i == 6 {
			exPdfDiv.Name = "Corp - Zakat Maal"
			exPdfDiv.Total = dataPdf.TotalCorporateZakatMaal
		} else if i == 7 {
			exPdfDiv.Name = "Corp - Zakat Fitrah"
			exPdfDiv.Total = dataPdf.TotalCorporateZakatFitrah
		} else if i == 8 {
			exPdfDiv.Name = "Corp - Infaq/Sedekah"
			exPdfDiv.Total = dataPdf.TotalCorporateInfaq
		} else if i == 9 {
			exPdfDiv.Name = "Corp - Waqaf"
			exPdfDiv.Total = dataPdf.TotalCorporateWakaf
		} else if i == 10 {
			exPdfDiv.Name = "Corp - Qurban"
			exPdfDiv.Total = dataPdf.TotalCorporateKurban
		} else if i == 11 {
			exPdfDiv.Name = "Corp - Penerimaan Lain "
			exPdfDiv.Total = dataPdf.TotalCorporateOther
		} else if i == 12 {
			exPdfDiv.Name = "UPZ - Zakat Maal"
			exPdfDiv.Total = dataPdf.TotalUpzZakatMaal
		} else if i == 13 {
			exPdfDiv.Name = "UPZ - Zakat Fitrah"
			exPdfDiv.Total = dataPdf.TotalUpzZakatFitrah
		} else if i == 14 {
			exPdfDiv.Name = "UPZ - Infaq/Sedekah"
			exPdfDiv.Total = dataPdf.TotalUpzInfaq
		} else if i == 15 {
			exPdfDiv.Name = "UPZ - Waqaf"
			exPdfDiv.Total = dataPdf.TotalUpzWakaf
		} else if i == 16 {
			exPdfDiv.Name = "UPZ - Qurban"
			exPdfDiv.Total = dataPdf.TotalUpzKurban
		} else if i == 17 {
			exPdfDiv.Name = "UPZ - Penerimaan Lain "
			exPdfDiv.Total = dataPdf.TotalUpzOther
		} else if i == 18 {
			exPdfDiv.Name = "Total"
			exPdfDiv.Total = total
		}

		exPdf.ReportPerDivision = append(exPdf.ReportPerDivision, exPdfDiv)
	}

	var (
		goodConfirm       []domain.Goods
		goodNotConfirm    []domain.Goods
		exPdfGood         domain.ExportPdfGood
		qytConfNotMove    int32
		qytConfMove       int32
		qytConfFood       int32
		qytConfOther      int32
		qytConfAnimal     int32
		totConfNotMove    uint64
		totConfMove       uint64
		totConfFood       uint64
		totConfOther      uint64
		totConfAnimal     uint64
		qytNotConfNotMove int32
		qytNotConfMove    int32
		qytNotConfFood    int32
		qytNotConfOther   int32
		qytNotConfAnimal  int32
		totNotConfNotMove uint64
		totNotConfMove    uint64
		totNotConfFood    uint64
		totNotConfOther   uint64
		totNotConfAnimal  uint64
		totRowNotMove     int
		totRowMove        int
		totRowFood        int
		totRowOther       int
		totRowAnimal      int
	)

	p.DB.Where("status = ?", 1).Find(&goodConfirm)
	p.DB.Where("status IN (?)", []uint{2, 3, 4}).Find(&goodNotConfirm)

	// Good Confirm
	for _, v := range transaction {
		for _, j := range goodConfirm {
			if v.ItemType == "goods" && j.ID == v.ItemID {
				if j.CategoryID == 1 {
					qytConfNotMove += j.Quantity
					totConfNotMove += v.Total
					totRowNotMove += 1
				} else if j.CategoryID == 2 {
					qytConfMove += j.Quantity
					totConfMove += v.Total
					totRowMove += 1
				} else if j.CategoryID == 3 {
					qytConfFood += j.Quantity
					totConfFood += v.Total
					totRowFood += 1
				} else if j.CategoryID == 4 {
					qytConfOther += j.Quantity
					totConfOther += v.Total
					totRowOther += 1
				} else if j.CategoryID == 5 {
					qytConfAnimal += j.Quantity
					totConfAnimal += v.Total
					totRowAnimal += 1
				}
				continue
			}
		}
	}

	// Good Not Confirm
	for _, v := range transaction {
		for _, j := range goodNotConfirm {
			if v.ItemType == "goods" && j.ID == v.ItemID {
				if j.CategoryID == 1 {
					qytNotConfNotMove += j.Quantity
					totNotConfNotMove += v.Total
					totRowNotMove = +1
				} else if j.CategoryID == 2 {
					qytNotConfMove += j.Quantity
					totNotConfMove += v.Total
					totRowMove += 1
				} else if j.CategoryID == 3 {
					qytNotConfFood += j.Quantity
					totNotConfFood += v.Total
					totRowFood += 1
				} else if j.CategoryID == 4 {
					qytNotConfOther += j.Quantity
					totNotConfOther += v.Total
					totRowOther += 1
				} else if j.CategoryID == 5 {
					qytNotConfAnimal += j.Quantity
					totNotConfAnimal += v.Total
					totRowAnimal += 1
				}
				continue
			}
		}
	}

	for i := 0; i < 6; i++ {
		if i == 0 {
			exPdfGood.Name = "Aset Tidak Bergerak"
			exPdfGood.TotalRowCount = totRowNotMove
			exPdfGood.TotalCollectCount = qytConfNotMove
			exPdfGood.TotalNotCollectCount = qytNotConfNotMove
			exPdfGood.TotalCollect = totConfNotMove
			exPdfGood.TotalNotCollect = totNotConfNotMove
		} else if i == 1 {
			exPdfGood.Name = "Aset Bergerak"
			exPdfGood.TotalRowCount = totRowMove
			exPdfGood.TotalCollectCount = qytConfMove
			exPdfGood.TotalNotCollectCount = qytNotConfMove
			exPdfGood.TotalCollect = totConfMove
			exPdfGood.TotalNotCollect = totNotConfMove
		} else if i == 2 {
			exPdfGood.Name = "Makanan/Minuman"
			exPdfGood.TotalRowCount = totRowFood
			exPdfGood.TotalCollectCount = qytConfFood
			exPdfGood.TotalNotCollectCount = qytNotConfFood
			exPdfGood.TotalCollect = totConfFood
			exPdfGood.TotalNotCollect = totNotConfFood
		} else if i == 3 {
			exPdfGood.Name = "Benda Lainnya"
			exPdfGood.TotalRowCount = totRowOther
			exPdfGood.TotalCollectCount = qytConfOther
			exPdfGood.TotalNotCollectCount = qytNotConfOther
			exPdfGood.TotalCollect = totConfOther
			exPdfGood.TotalNotCollect = totNotConfOther
		} else if i == 4 {
			exPdfGood.Name = "Hewan Kurban"
			exPdfGood.TotalRowCount = totRowAnimal
			exPdfGood.TotalCollectCount = qytConfAnimal
			exPdfGood.TotalNotCollectCount = qytNotConfAnimal
			exPdfGood.TotalCollect = totConfAnimal
			exPdfGood.TotalNotCollect = totNotConfAnimal
		} else if i == 5 {
			exPdfGood.Name = "Total"
			exPdfGood.TotalRowCount = totRowNotMove + totRowMove + totRowFood + totRowOther + totRowAnimal
			exPdfGood.TotalCollectCount = qytConfNotMove + qytConfMove + qytConfFood + qytConfOther + qytConfAnimal
			exPdfGood.TotalNotCollectCount = qytNotConfNotMove + qytNotConfMove + qytNotConfFood + qytNotConfOther + qytNotConfAnimal
			exPdfGood.TotalCollect = totConfNotMove + totConfMove + totConfFood + totConfOther + totConfAnimal
			exPdfGood.TotalNotCollect = totNotConfNotMove + totNotConfMove + totNotConfFood + totNotConfOther
		}
		exPdf.ReportGood = append(exPdf.ReportGood, exPdfGood)
	}

	var (
		cash         []domain.Cash
		exPdfCash    domain.ExportPdfCash
		totalAllCash uint64
	)

	p.DB.Find(&cash)

	for _, v := range transaction {
		for _, j := range cash {
			if v.ItemType == "cashs" && j.ID == v.ItemID {
				if j.CategoryID == 1 {
					dataPdf.TotalCash += v.Total
					totalAllCash += v.Total
				} else if j.CategoryID == 2 {
					dataPdf.TotalNonCashMuamalat += v.Total
					totalAllCash += v.Total
				} else if j.CategoryID == 3 {
					dataPdf.TotalNonCashMandiri += v.Total
					totalAllCash += v.Total
				} else if j.CategoryID == 4 {
					dataPdf.TotalNonCashBsm += v.Total
					totalAllCash += v.Total
				} else if j.CategoryID == 5 {
					dataPdf.TotalNonCashBri += v.Total
					totalAllCash += v.Total
				} else if j.CategoryID == 6 {
					dataPdf.TotalNonCashBniLamp += v.Total
					totalAllCash += v.Total
				} else if j.CategoryID == 7 {
					dataPdf.TotalNonCashBniSy += v.Total
					totalAllCash += v.Total
				} else if j.CategoryID == 8 {
					dataPdf.TotalNonCashBca += v.Total
					totalAllCash += v.Total
				}
				continue
			}
		}
	}

	for i := 0; i < 9; i++ {
		if i == 0 {
			exPdfCash.Name = "Tunai"
			exPdfCash.Total = dataPdf.TotalCash
			exPdfCash.Percent = percent(dataPdf.TotalCash, totalAllCash)
		} else if i == 1 {
			exPdfCash.Name = "Non Tunai - Muamalat"
			exPdfCash.Total = dataPdf.TotalNonCashMuamalat
			exPdfCash.Percent = percent(dataPdf.TotalNonCashMuamalat, totalAllCash)
		} else if i == 2 {
			exPdfCash.Name = "Non Tunai - Mandiri"
			exPdfCash.Total = dataPdf.TotalNonCashMandiri
			exPdfCash.Percent = percent(dataPdf.TotalNonCashMandiri, totalAllCash)
		} else if i == 3 {
			exPdfCash.Name = "Non Tunai - BSM"
			exPdfCash.Total = dataPdf.TotalNonCashBsm
			exPdfCash.Percent = percent(dataPdf.TotalNonCashBsm, totalAllCash)
		} else if i == 4 {
			exPdfCash.Name = "Non Tunai - BRI Syariah"
			exPdfCash.Total = dataPdf.TotalNonCashBri
			exPdfCash.Percent = percent(dataPdf.TotalNonCashBri, totalAllCash)
		} else if i == 5 {
			exPdfCash.Name = "Non Tunai - BNI Sy Lamp"
			exPdfCash.Total = dataPdf.TotalNonCashBniLamp
			exPdfCash.Percent = percent(dataPdf.TotalNonCashBniLamp, totalAllCash)
		} else if i == 6 {
			exPdfCash.Name = "Non Tunai - BNI Syariah"
			exPdfCash.Total = dataPdf.TotalNonCashBniSy
			exPdfCash.Percent = percent(dataPdf.TotalNonCashBniSy, totalAllCash)
		} else if i == 7 {
			exPdfCash.Name = "Non Tunai - BCA"
			exPdfCash.Total = dataPdf.TotalNonCashBca
			exPdfCash.Percent = percent(dataPdf.TotalNonCashBca, totalAllCash)
		} else if i == 8 {
			exPdfCash.Name = "Total"
			exPdfCash.Total = totalAllCash
			exPdfCash.Percent = percent(totalAllCash, totalAllCash)
		}

		exPdf.ReportCash = append(exPdf.ReportCash, exPdfCash)
	}

	return exPdf, nil
}

func unique(intSlice []uint64) []uint64 {
	keys := make(map[uint64]bool)
	list := []uint64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func percent(a uint64, b uint64) uint64 {
	if b == 0 {
		return 0
	}
	return uint64(math.Round(float64((a * 100)) / float64(b)))
}

func Find(slice []int8, val int8) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
