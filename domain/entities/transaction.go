package entities

import "time"

type Transaction struct {
	ID                  uint64            `gorm:"primary_key" json:"id"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	DeletedAt           *time.Time        `json:"deleted_at"`
	DonorID             uint64            `gorm:"column:donor_id" json:"donor_id"`
	DivisionID          uint16            `gorm:"column:division_id" json:"division_id"`
	CategoryID          uint16            `gorm:"column:category_id" json:"category_id"`
	StatementCategoryID uint64            `gorm:"column:statement_category_id"`
	Description         string            `gorm:"column:description" json:"description"`
	ItemID              uint64            `gorm:"column:item_id" json:"item_id"`
	ItemType            string            `gorm:"column:item_type" json:"item_type"`
	SchoolID            uint64            `gorm:"column:school_id" json:"employee_id"`
	Status              int8              `gorm:"column:status" json:"status"`
	Total               uint64            `gorm:"column:total" json:"total"`
	Kwitansi            string            `gorm:"column:kwitansi" json:"kwitansi"`
	Donor               Donor             `json:"donor"`
	Division            Division          `json:"division"`
	School              School            `json:"school"`
	Category            Category          `json:"category"`
	StatementCategory   StatementCategory `json:"statement_category"`
	GoodsQuery          []Goods           `gorm:"foreignkey:ID;association_foreignkey:ItemID" json:"goods"`
	Cashes              []Cash            `gorm:"foreignkey:ID;association_foreignkey:ItemID" json:"cashes"`
	CreatedBy           string            `gorm:"column:created_by" json:"created_by"`
	UpdatedBy           string            `gorm:"column:updated_by" json:"updated_by"`
}

type TransactionFilter struct {
	Search            string `json:"search"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	StartTotal        string `json:"start_total"`
	EndTotal          string `json:"end_total"`
	Regency           string `json:"regency"`
	DonaturStatus     string `json:"status_donatur"`
	SchoolID          string `json:"school_id"`
	DivisionID        string `json:"division_id"`
	DonaturCategory   string `json:"donatur_category"`
	StatementCategory string `json:"statement_category"`
	CategoryID        string `json:"category_id"`
	CategoryType      string `json:"category_type"`
	DonationType      string `json:"donation_type"`
	Sort              string `json:"sort"`
	SortType          string `json:"sort_type"`
	Page              string `json:"page"`
	Limit             string `json:"limit"`
	Role              string `json:"role"`
}

type ReportResponse struct {
	Total                     uint64                   `json:"total"`
	TotalRowCount             int                      `json:"total_row_count"`
	TotalRowDonorCount        int                      `json:"total_row_donor_count"`
	TotalRowNewDonorCount     int                      `json:"total_row_new_donor_count"`
	TotalZakatMaal            uint64                   `json:"total_zakat_maal"`
	TotalWakaf                uint64                   `json:"total_wakaf"`
	TotalZakatFitrah          uint64                   `json:"total_zakat_fitrah"`
	TotalInfaq                uint64                   `json:"total_infaq"`
	TotalKurban               uint64                   `json:"total_kurban"`
	TotalOther                uint64                   `json:"total_other"`
	TotalRetail               uint64                   `json:"total_retail"`
	TotalRetailRowCount       int                      `json:"totalretail_row_count"`
	TotalCorporate            uint64                   `json:"total_corporate"`
	TotalCorporateRowCount    int                      `json:"total_corporate_row_count"`
	TotalUpz                  uint64                   `json:"total_upz"`
	TotalUpzRowCount          int                      `json:"total_upz_row_count"`
	TotalPrognosisRetail      uint64                   `json:"total_prognosis_retail"`
	TotalPrognosisCorporate   uint64                   `json:"total_prognosis_corporate"`
	TotalPrognosisUpz         uint64                   `json:"total_prognosis_upz"`
	TotalPercentRetail        uint64                   `json:"total_percent_retail"`
	TotalPercentCorporate     uint64                   `json:"total_percent_corporate"`
	TotalPercentUpz           uint64                   `json:"total_percent_upz"`
	TotalCompanyDonor         uint64                   `json:"total_company_donor"`
	TotalPersonDonor          uint64                   `json:"total_person_donor"`
	TotalCompanyDonorRowCount int                      `json:"total_company_donor_row_count"`
	TotalPersonDonorRowCount  int                      `json:"total_person_donor_row_count"`
	TotalPercentCompanyDonor  uint64                   `json:"total_percent_company_donor"`
	TotalPercentPersonDonor   uint64                   `json:"total_percent_person_donor"`
	TotalRitelZakatMaal       uint64                   `json:"total_ritel_zakat_maal"`
	TotalRitelWakaf           uint64                   `json:"total_ritel_wakaf"`
	TotalRitelZakatFitrah     uint64                   `json:"total_ritel_zakat_fitrah"`
	TotalRitelInfaq           uint64                   `json:"total_ritel_infaq"`
	TotalRitelKurban          uint64                   `json:"total_ritel_kurban"`
	TotalRitelOther           uint64                   `json:"total_ritel_other"`
	TotalCorporateZakatMaal   uint64                   `json:"total_corporate_zakat_maal"`
	TotalCorporateWakaf       uint64                   `json:"total_corporate_wakaf"`
	TotalCorporateZakatFitrah uint64                   `json:"total_corporate_zakat_fitrah"`
	TotalCorporateInfaq       uint64                   `json:"total_corporate_infaq"`
	TotalCorporateKurban      uint64                   `json:"total_corporate_kurban"`
	TotalCorporateOther       uint64                   `json:"total_corporate_other"`
	TotalUpzZakatMaal         uint64                   `json:"total_upz_zakat_maal"`
	TotalUpzWakaf             uint64                   `json:"total_upz_wakaf"`
	TotalUpzZakatFitrah       uint64                   `json:"total_upz_zakat_fitrah"`
	TotalUpzInfaq             uint64                   `json:"total_upz_infaq"`
	TotalUpzKurban            uint64                   `json:"total_upz_kurban"`
	TotalUpzOther             uint64                   `json:"total_upz_other"`
	TotalGood                 uint64                   `json:"total_good"`
	TotalGoodCount            uint64                   `json:"total_good_count"`
	TotalGoodCollect          int32                    `json:"total_good_collect"`
	TotalGoodMoveCount        int32                    `json:"total_good_move_count"`
	TotalGoodNotMoveCount     int32                    `json:"total_good_not_move_count"`
	TotalGoodFoodCount        int32                    `json:"total_good_food_count"`
	TotalGoodOtherCount       int32                    `json:"total_good_other_count"`
	TotalAllCash              uint64                   `json:"total_all_cash"`
	TotalCashCount            uint64                   `json:"total_cash_count"`
	TotalCash                 uint64                   `json:"total_cash"`
	TotalNonCashMuamalat      uint64                   `json:"total_non_cash_muamalat"`
	TotalNonCashMandiri       uint64                   `json:"total_non_cash_mandiri"`
	TotalNonCashBsm           uint64                   `json:"total_non_cash_bsm"`
	TotalNonCashBri           uint64                   `json:"total_non_cash_bri"`
	TotalNonCashBniLamp       uint64                   `json:"total_non_cash_bni_lamp"`
	TotalNonCashBniSy         uint64                   `json:"total_non_cash_bni_sy"`
	TotalNonCashBca           uint64                   `json:"total_non_cash_bca"`
	TotalDivisionPercent      TotalDivisionPercent     `json:"total_division_percent"`
	Prognosis                 []TransactionGoal        `json:"prognosis"`
	TotalUpzPerMonth          TotalUpzPerMonth         `json:"total_upz_per_month"`
	TotalRetailPerMonth       TotalRetailPerMonth      `json:"total_retail_per_month"`
	TotalCorporatePerMonth    TotalCorporatePerMonth   `json:"total_corporate_per_month"`
	TotalPercentZiswaf        TotalPercentZiswaf       `json:"total_ziswaf_percent"`
	TotalCashPercent          TotalCashPercent         `json:"total_cash_percent"`
	TotalZiswafPerMonth       TotalZiswafPerMonth      `json:"total_ziswaf_per_month"`
	TotalZakatMaalPerMonth    TotalZakatMaalPerMonth   `json:"total_zakat_maal_per_month"`
	TotalZakatFitrahPerMonth  TotalZakatFitrahPerMonth `json:"total_zakat_fitrah_per_month"`
	TotalInfaqPerMonth        TotalInfaqPerMonth       `json:"total_Infaq_per_month"`
	TotalWakafPerMonth        TotalWakafPerMonth       `json:"total_wakaf_per_month"`
	TotalQurbanPerMonth       TotalQurbanPerMonth      `json:"total_qurban_per_month"`
	TotalOtherPerMonth        TotalOtherPerMonth       `json:"total_other_per_month"`
	TotalPrognosisPerMonth    TotalPrognosisPerMonth   `json:"total_prognosis_per_month"`
	TotalTransactionPerMonth  TotalTransactionPerMonth `json:"total_transaction_per_month"`
	TotalTransactionLastYear  TotalTransactionLastYear `json:"total_transaction_last_year"`
}

type ReportOperatorResponse struct {
	Total                     uint64                   `json:"total"`
	TotalRowCount             int                      `json:"total_row_count"`
	TotalRowDonorCount        int                      `json:"total_row_donor_count"`
	TotalRowNewDonorCount     int                      `json:"total_row_new_donor_count"`
	TotalZakatMaal            uint64                   `json:"total_zakat_maal"`
	TotalWakaf                uint64                   `json:"total_wakaf"`
	TotalZakatFitrah          uint64                   `json:"total_zakat_fitrah"`
	TotalInfaq                uint64                   `json:"total_infaq"`
	TotalKurban               uint64                   `json:"total_kurban"`
	TotalOther                uint64                   `json:"total_other"`
	TotalCompanyDonor         uint64                   `json:"total_company_donor"`
	TotalPersonDonor          uint64                   `json:"total_person_donor"`
	TotalPercentCompanyDonor  uint64                   `json:"total_percent_company_donor"`
	TotalPercentPersonDonor   uint64                   `json:"total_percent_person_donor"`
	TotalCompanyDonorRowCount int                      `json:"total_company_donor_row_count"`
	TotalPersonDonorRowCount  int                      `json:"total_person_donor_row_count"`
	TotalRetail               uint64                   `json:"total_retail"`
	TotalCorporate            uint64                   `json:"total_corporate"`
	TotalUpz                  uint64                   `json:"total_upz"`
	TotalRitelZakatMaal       uint64                   `json:"total_ritel_zakat_maal"`
	TotalRitelWakaf           uint64                   `json:"total_ritel_wakaf"`
	TotalRitelZakatFitrah     uint64                   `json:"total_ritel_zakat_fitrah"`
	TotalRitelInfaq           uint64                   `json:"total_ritel_infaq"`
	TotalRitelKurban          uint64                   `json:"total_ritel_kurban"`
	TotalRitelOther           uint64                   `json:"total_ritel_other"`
	TotalCorporateZakatMaal   uint64                   `json:"total_corporate_zakat_maal"`
	TotalCorporateWakaf       uint64                   `json:"total_corporate_wakaf"`
	TotalCorporateZakatFitrah uint64                   `json:"total_corporate_zakat_fitrah"`
	TotalCorporateInfaq       uint64                   `json:"total_corporate_infaq"`
	TotalCorporateKurban      uint64                   `json:"total_corporate_kurban"`
	TotalCorporateOther       uint64                   `json:"total_corporate_other"`
	TotalUpzZakatMaal         uint64                   `json:"total_upz_zakat_maal"`
	TotalUpzWakaf             uint64                   `json:"total_upz_wakaf"`
	TotalUpzZakatFitrah       uint64                   `json:"total_upz_zakat_fitrah"`
	TotalUpzInfaq             uint64                   `json:"total_upz_infaq"`
	TotalUpzKurban            uint64                   `json:"total_upz_kurban"`
	TotalUpzOther             uint64                   `json:"total_upz_other"`
	TotalGood                 uint64                   `json:"total_good"`
	TotalGoodCount            uint64                   `json:"total_good_count"`
	TotalGoodCollect          int32                    `json:"total_good_collect"`
	TotalGoodMoveCount        int32                    `json:"total_good_move_count"`
	TotalGoodNotMoveCount     int32                    `json:"total_good_not_move_count"`
	TotalGoodFoodCount        int32                    `json:"total_good_food_count"`
	TotalGoodOtherCount       int32                    `json:"total_good_other_count"`
	TotalAllCash              uint64                   `json:"total_all_cash"`
	TotalCashCount            uint64                   `json:"total_cash_count"`
	TotalCash                 uint64                   `json:"total_cash"`
	TotalNonCashMuamalat      uint64                   `json:"total_non_cash_muamalat"`
	TotalNonCashMandiri       uint64                   `json:"total_non_cash_mandiri"`
	TotalNonCashBsm           uint64                   `json:"total_non_cash_bsm"`
	TotalNonCashBri           uint64                   `json:"total_non_cash_bri"`
	TotalNonCashBniLamp       uint64                   `json:"total_non_cash_bni_lamp"`
	TotalNonCashBniSy         uint64                   `json:"total_non_cash_bni_sy"`
	TotalNonCashBca           uint64                   `json:"total_non_cash_bca"`
	TotalCashPercent          TotalCashPercent         `json:"total_cash_percent"`
	TotalUpzPerDay            TotalUpzPerDay           `json:"total_upz_per_day"`
	TotalRetailPerDay         TotalRetailPerDay        `json:"total_retail_per_day"`
	TotalCorporatePerDay      TotalCorporatePerDay     `json:"total_corporate_per_day"`
	TotalDivisionPercent      TotalDivisionPercent     `json:"total_division_percent"`
	TotalPercentZiswaf        TotalPercentZiswaf       `json:"total_ziswaf_percent"`
	TotalZiswafPerDay         TotalZiswafPerDay        `json:"total_ziswaf_per_day"`
	TotalZakatMaalPerDay      TotalZakatMaalPerDay     `json:"total_zakat_maal_per_day"`
	TotalZakatFitrahPerDay    TotalZakatFitrahPerDay   `json:"total_zakat_fitrah_per_day"`
	TotalInfaqPerDay          TotalInfaqPerDay         `json:"total_infaq_per_day"`
	TotalWakafPerDay          TotalWakafPerDay         `json:"total_wakaf_per_day"`
	TotalQurbanPerDay         TotalQurbanPerDay        `json:"total_qurban_per_day"`
	TotalOtherPerDay          TotalOtherPerDay         `json:"total_other_per_day"`
	TotalTransactionLastMonth TotalTransactionLastYear `json:"total_transaction_last_month"`
}

type ReportStatementCategory struct {
	Data          []ReportStatementCategoryData `json:"data"`
	TotalRowCount int                           `json:"total_row_count"`
	TotalCash     uint64                        `json:"total_cash"`
}

type ReportStatementCategoryData struct {
	StatementCategoryID uint64 `json:"statement_category_id"`
	Name                string `json:"name"`
	Total               uint64 `json:"total"`
	TotalPercent        uint64 `json:"total_percent"`
}

type ExportPdf struct {
	Role              string
	SchoolName        string `json:"school_name"`
	RegencyName       string `json:"regency_name"`
	StartDate         string
	EndDate           string
	Data              []ExportPdfDetailData  `json:"data"`
	ReportPerCategory ExportPdfPerCategory   `json:"report_per_category"`
	ReportPerDivision []ExportPdfPerDivision `json:"report_per_division"`
	ReportCash        []ExportPdfCash        `json:"report_cash"`
	ReportGood        []ExportPdfGood        `json:"report_good"`
}

type ExportPdfData struct {
	TotalRowCount           uint64 `json:"total_row_count"`
	TotalCash               uint64 `json:"total_cash"`
	TotalGood               uint64 `json:"total_good"`
	TotalPrognosis          uint64 `json:"total_prognosis"`
	TotalPercent            uint64 `json:"total_percent"`
	TotalRowCountUpz        uint64 `json:"total_row_count_upz"`
	TotalCashUpz            uint64 `json:"total_cash_upz"`
	TotalGoodUpz            uint64 `json:"total_good_upz"`
	TotalPrognosisUpz       uint64 `json:"total_prognosis_upz"`
	TotalPercentUpz         uint64 `json:"total_percent_upz"`
	TotalRowCountRetail     uint64 `json:"total_row_count_retail"`
	TotalCashRetail         uint64 `json:"total_cash_retail"`
	TotalGoodRetail         uint64 `json:"total_good_retail"`
	TotalPrognosisRetail    uint64 `json:"total_prognosis_retail"`
	TotalPercentRetail      uint64 `json:"total_percent_retail"`
	TotalRowCountCorporate  uint64 `json:"total_row_count_corporate"`
	TotalCashCorporate      uint64 `json:"total_cash_corporate"`
	TotalGoodCorporate      uint64 `json:"total_good_corporate"`
	TotalPrognosisCorporate uint64 `json:"total_prognosis_corporate"`
	TotalPercentCorporate   uint64 `json:"total_percent_corporate"`
	TotalGap                uint64 `json:"total_gap"`
	TotalGapUpz             uint64 `json:"total_gap_upz"`
	TotalGapRetail          uint64 `json:"total_gap_retail"`
	TotalGapCorporate       uint64 `json:"total_gap_corporate"`
}

type ExportPdfPerCategoryDetail struct {
	Name    string `json:"name"`
	Total   uint64 `json:"total"`
	Percent uint64 `json:"percent"`
}

type ExportPdfPerCategory struct {
	TotalCorporate             uint64                       `json:"total_corporate"`
	TotalPersonal              uint64                       `json:"total_personal"`
	TotalRowCountCorporate     uint64                       `json:"total_row_count_corporate"`
	TotalRowCountPersonal      uint64                       `json:"total_row_count_personal"`
	TotalPercentCorporate      uint64                       `json:"total_percent_corporate"`
	TotalPercentPersonal       uint64                       `json:"total_percent_personal"`
	ExportPdfPerCategoryDetail []ExportPdfPerCategoryDetail `json:"export_pdf_per_category_detail"`
}
type ExportPdfPerDivision struct {
	Name  string `json:"name"`
	Total uint64 `json:"total"`
}
type ExportPdfCash struct {
	Name    string `json:"name"`
	Total   uint64 `json:"total"`
	Percent uint64 `json:"percent"`
}
type ExportPdfGood struct {
	Name                 string `json:"name"`
	TotalRowCount        int    `json:"total_row_count"`
	TotalCollectCount    int32  `json:"total_collect_count"`
	TotalCollect         uint64 `json:"total_collect"`
	TotalNotCollectCount int32  `json:"total_not_collect_count"`
	TotalNotCollect      uint64 `json:"total_not_collect"`
}

type ExportPdfDetailData struct {
	Name           string `json:"name"`
	Month          string `json:"month"`
	Total          uint64 `json:"total"`
	TotalRowCount  uint64 `json:"total_row_count"`
	TotalPrognosis uint64 `json:"total_prognosis"`
	TotalGap       uint64 `json:"total_gap"`
	TotalPercent   uint64 `json:"total_percent"`
}

type TotalUpzPerMonth struct {
	TotalUpzJanuary   uint64 `json:"total_upz_january"`
	TotalUpzFebruary  uint64 `json:"total_upz_february"`
	TotalUpzMarch     uint64 `json:"total_upz_march"`
	TotalUpzApril     uint64 `json:"total_upz_april"`
	TotalUpzMay       uint64 `json:"total_upz_may"`
	TotalUpzJune      uint64 `json:"total_upz_june"`
	TotalUpzJuly      uint64 `json:"total_upz_july"`
	TotalUpzAugust    uint64 `json:"total_upz_august"`
	TotalUpzSeptember uint64 `json:"total_upz_september"`
	TotalUpzOctober   uint64 `json:"total_upz_october"`
	TotalUpzNovember  uint64 `json:"total_upz_november"`
	TotalUpzDecember  uint64 `json:"total_upz_december"`
}

type TotalRetailPerMonth struct {
	TotalRetailJanuary   uint64 `json:"total_retail_january"`
	TotalRetailFebruary  uint64 `json:"total_retail_february"`
	TotalRetailMarch     uint64 `json:"total_retail_march"`
	TotalRetailApril     uint64 `json:"total_retail_april"`
	TotalRetailMay       uint64 `json:"total_retail_may"`
	TotalRetailJune      uint64 `json:"total_retail_june"`
	TotalRetailJuly      uint64 `json:"total_retail_july"`
	TotalRetailAugust    uint64 `json:"total_retail_august"`
	TotalRetailSeptember uint64 `json:"total_retail_september"`
	TotalRetailOctober   uint64 `json:"total_retail_october"`
	TotalRetailNovember  uint64 `json:"total_retail_november"`
	TotalRetailDecember  uint64 `json:"total_retail_december"`
}

type TotalCorporatePerMonth struct {
	TotalCorporateJanuary   uint64 `json:"total_corporate_january"`
	TotalCorporateFebruary  uint64 `json:"total_corporate_february"`
	TotalCorporateMarch     uint64 `json:"total_corporate_march"`
	TotalCorporateApril     uint64 `json:"total_corporate_april"`
	TotalCorporateMay       uint64 `json:"total_corporate_may"`
	TotalCorporateJune      uint64 `json:"total_corporate_june"`
	TotalCorporateJuly      uint64 `json:"total_corporate_july"`
	TotalCorporateAugust    uint64 `json:"total_corporate_august"`
	TotalCorporateSeptember uint64 `json:"total_corporate_september"`
	TotalCorporateOctober   uint64 `json:"total_corporate_october"`
	TotalCorporateNovember  uint64 `json:"total_corporate_november"`
	TotalCorporateDecember  uint64 `json:"total_corporate_december"`
}

type ZiswafDay struct {
	Total uint64 `json:"total"`
}
type ZakatMaalDay struct {
	Total uint64 `json:"total"`
}
type ZakatFitrahDay struct {
	Total uint64 `json:"total"`
}
type QurbanDay struct {
	Total uint64 `json:"total"`
}
type WaqafDay struct {
	Total uint64 `json:"total"`
}
type InfaqDay struct {
	Total uint64 `json:"total"`
}
type OtherDay struct {
	Total uint64 `json:"total"`
}

type TotalZiswafPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalZakatMaalPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalZakatFitrahPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalInfaqPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalWakafPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalQurbanPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalOtherPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalPercentZiswaf struct {
	ZakatMaalPercent   uint64 `json:"zakat_maal_percent"`
	ZakatFitrahPercent uint64 `json:"zakat_fitrah_percent"`
	InfaqPercent       uint64 `json:"infaq_percent"`
	WaqafPercent       uint64 `json:"waqaf_percent"`
	QurbanPercent      uint64 `json:"qurban_percent"`
	OtherPercent       uint64 `json:"other_percent"`
}

type TotalCashPercent struct {
	CashPercent            uint64 `json:"cash_percent"`
	NonCashMuamalatPercent uint64 `json:"non_cash_muamalat_percent"`
	NonCashMandiriPercent  uint64 `json:"non_cash_mandiri_percent"`
	NonCashBsmPercent      uint64 `json:"non_cash_bsm_percent"`
	NonCashBriPercent      uint64 `json:"non_cash_bri_percent"`
	NonCashBniLampPercent  uint64 `json:"non_cash_bni_lamp_percent"`
	NonCashBniSyPercent    uint64 `json:"non_cash_bni_sy_percent"`
	NonCashBcaPercent      uint64 `json:"non_cash_bca_percent"`
}

type ZiswafMonth struct {
	Total uint64 `json:"total"`
}
type ZakatMaalMonth struct {
	Total uint64 `json:"total"`
}
type ZakatFitrahMonth struct {
	Total uint64 `json:"total"`
}
type QurbanMonth struct {
	Total uint64 `json:"total"`
}
type WaqafMonth struct {
	Total uint64 `json:"total"`
}
type InfaqMonth struct {
	Total uint64 `json:"total"`
}
type OtherMonth struct {
	Total uint64 `json:"total"`
}

type TotalZiswafPerMonth struct {
	Month []struct {
		Total uint64 `json:"total"`
	} `json:"month"`
}

type TotalZakatMaalPerMonth struct {
	Month []struct {
		Total uint64 `json:"total"`
	} `json:"month"`
}

type TotalZakatFitrahPerMonth struct {
	Month []struct {
		Total uint64 `json:"total"`
	} `json:"month"`
}

type TotalInfaqPerMonth struct {
	Month []struct {
		Total uint64 `json:"total"`
	} `json:"month"`
}

type TotalWakafPerMonth struct {
	Month []struct {
		Total uint64 `json:"total"`
	} `json:"month"`
}

type TotalQurbanPerMonth struct {
	Month []struct {
		Total uint64 `json:"total"`
	} `json:"month"`
}

type TotalOtherPerMonth struct {
	Month []struct {
		Total uint64 `json:"total"`
	} `json:"month"`
}

type TotalPrognosisPerMonth struct {
	January   uint64 `json:"january"`
	February  uint64 `json:"february"`
	March     uint64 `json:"march"`
	April     uint64 `json:"april"`
	May       uint64 `json:"may"`
	June      uint64 `json:"june"`
	July      uint64 `json:"july"`
	August    uint64 `json:"august"`
	September uint64 `json:"september"`
	October   uint64 `json:"october"`
	November  uint64 `json:"november"`
	December  uint64 `json:"december"`
}
type TotalTransactionPerMonth struct {
	January   uint64 `json:"january"`
	February  uint64 `json:"february"`
	March     uint64 `json:"march"`
	April     uint64 `json:"april"`
	May       uint64 `json:"may"`
	June      uint64 `json:"june"`
	July      uint64 `json:"july"`
	August    uint64 `json:"august"`
	September uint64 `json:"september"`
	October   uint64 `json:"october"`
	November  uint64 `json:"november"`
	December  uint64 `json:"december"`
}

type TotalTransactionLastYear struct {
	TotalUp      uint64 `json:"total_up"`
	TotalDown    uint64 `json:"total_down"`
	CountUp      int    `json:"count_up"`
	CountDown    int    `json:"count_down"`
	DonorUp      int    `json:"donor_up"`
	DonorDown    int    `json:"donor_down"`
	NewDonorUp   int    `json:"new_donor_up"`
	NewDonorDown int    `json:"new_donor_down"`
}

type TotalUpzPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalRetailPerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type TotalCorporatePerDay struct {
	Day []struct {
		Total uint64 `json:"total"`
	} `json:"day"`
}

type UpzDay struct {
	Total uint64 `json:"total"`
}

type RetailDay struct {
	Total uint64 `json:"total"`
}

type CorporateDay struct {
	Total uint64 `json:"total"`
}

type TotalDivisionPercent struct {
	TotalDivisionUpzPercent       uint64 `json:"total_division_upz_percent"`
	TotalDivisionRetailPercent    uint64 `json:"total_division_retail_percent"`
	TotalDivisionCorporatePercent uint64 `json:"total_division_corporate_percent"`
}

type Kwitansi struct {
	ID               uint64    `json:"id"`
	Description      string    `json:"description"`
	DonorName        string    `json:"donor_name"`
	DonorAddress     string    `json:"donor_address,omitempty"`
	DonorPhone       string    `json:"donor_phone"`
	DonorEmail       string    `json:"donor_email"`
	DonorNPWP        int64     `json:"donor_npwp"`
	DivisionName     string    `json:"division_name"`
	Unit             string    `json:"unit"`
	City             string    `json:"city"`
	SchoolName       string    `json:"school_name"`
	Kwitansi         string    `json:"kwitansi"`
	Category         string    `json:"category"`
	StatmentCategory string    `json:"statement_category,omitempty"`
	Total            uint64    `json:"total"`
	ItemType         string    `json:"item_type"`
	ItemCategory     string    `json:"item_category"`
	RefNumber        string    `json:"ref_number"`
	Quantity         int32     `json:"quantity"`
	Status           int8      `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	CashDescription  string    `json:"cash_description"`
	GoodDescription  string    `json:"good_description"`
	DonorCategory    bool      `json:"donor_category"`
	GoodStatus       int8      `json:"good_status"`
	CreatedBy        string    `json:"created_by"`
}
