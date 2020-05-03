package list_report

import (
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListReportResponse struct {
		base.BaseResponse
		Data ListReportResponseData `json:"data"`
	}

	ListReportResponseData struct {
		DashboardAdmin     dashboardAdminResponse        `json:"dashboard_admin"`
		CommonReport       commonReport                  `json:"common_report"`
		DivisionReport     divisionReport                `json:"division_report"`
		NominalReport      nominalReport                 `json:"nominal_report"`
		PrognosisUpz       PrognosisUpz                  `json:"prognosis_upz"`
		PrognosisRetail    PrognosisRetail               `json:"prognosis_retail"`
		PrognosisCorporate PrognosisCorporate            `json:"prognosis_corporate"`
		UpzPerMonth        domain.TotalUpzPerMonth       `json:"upz_per_month"`
		RetailPerMonth     domain.TotalRetailPerMonth    `json:"retail_per_month"`
		CorporatePerMonth  domain.TotalCorporatePerMonth `json:"corporate_per_month"`
	}

	PrognosisUpz struct {
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

	PrognosisRetail struct {
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

	PrognosisCorporate struct {
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
	dashboardAdminResponse struct {
		Total                    uint64                          `json:"total"`
		TotalRowCount            int                             `json:"total_row_count"`
		TotalRowDonorCount       int                             `json:"total_row_donor_count"`
		TotalRowNewDonorCount    int                             `json:"total_row_new_donor_count"`
		TotalZakatMaal           uint64                          `json:"total_zakat_maal"`
		TotalWakaf               uint64                          `json:"total_wakaf"`
		TotalZakatFitrah         uint64                          `json:"total_zakat_fitrah"`
		TotalInfaq               uint64                          `json:"total_infaq"`
		TotalKurban              uint64                          `json:"total_kurban"`
		TotalOther               uint64                          `json:"total_other"`
		TotalRetail              uint64                          `json:"total_retail"`
		TotalRetailRowCount      int                             `json:"total_retail_row_count"`
		TotalCorporate           uint64                          `json:"total_corporate"`
		TotalCorporateRowCount   int                             `json:"total_corporate_row_count"`
		TotalUpz                 uint64                          `json:"total_upz"`
		TotalUpzRowCount         int                             `json:"total_upz_row_count"`
		TotalPrognosisRetail     uint64                          `json:"total_prognosis_retail"`
		TotalPrognosisCorporate  uint64                          `json:"total_prognosis_corporate"`
		TotalPrognosisUpz        uint64                          `json:"total_prognosis_upz"`
		TotalPercentRetail       uint64                          `json:"total_percent_retail"`
		TotalPercentCorporate    uint64                          `json:"total_percent_corporate"`
		TotalPercentUpz          uint64                          `json:"total_percent_upz"`
		TotalLastYear            domain.TotalTransactionLastYear `json:"total_transaction_last_year"`
		TotalZiswafPerMonth      domain.TotalZiswafPerMonth      `json:"total_ziswaf_per_month"`
		TotalZakatMaalPerMonth   domain.TotalZakatMaalPerMonth   `json:"total_zakat_maal_per_month"`
		TotalZakatFitrahPerMonth domain.TotalZakatFitrahPerMonth `json:"total_zakat_fitrah_per_month"`
		TotalInfaqPerMonth       domain.TotalInfaqPerMonth       `json:"total_Infaq_per_month"`
		TotalWakafPerMonth       domain.TotalWakafPerMonth       `json:"total_wakaf_per_month"`
		TotalQurbanPerMonth      domain.TotalQurbanPerMonth      `json:"total_qurban_per_month"`
		TotalOtherPerMonth       domain.TotalOtherPerMonth       `json:"total_other_per_month"`
	}

	commonReport struct {
		Total                     uint64                   `json:"total"`
		TotalRowCount             int                      `json:"total_row_count"`
		TotalRowDonorCount        int                      `json:"total_row_donor_count"`
		TotalZakatMaal            uint64                   `json:"total_zakat_maal"`
		TotalWakaf                uint64                   `json:"total_wakaf"`
		TotalZakatFitrah          uint64                   `json:"total_zakat_fitrah"`
		TotalInfaq                uint64                   `json:"total_infaq"`
		TotalKurban               uint64                   `json:"total_kurban"`
		TotalOther                uint64                   `json:"total_other"`
		TotalCompanyDonor         uint64                   `json:"total_company_donor"`
		TotalPersonDonor          uint64                   `json:"total_person_donor"`
		TotalCompanyDonorRowCount int                      `json:"total_company_donor_row_count"`
		TotalPersonDonorRowCount  int                      `json:"total_person_donor_row_count"`
		TotalPercentCompanyDonor  uint64                   `json:"total_percent_company_donor"`
		TotalPercentPersonDonor   uint64                   `json:"total_percent_person_donor"`
		TotalRetail               uint64                   `json:"total_retail"`
		TotalCorporate            uint64                   `json:"total_corporate"`
		TotalUpz                  uint64                   `json:"total_upz"`
		TotalDivisionPercent      TotalDivisionPercent     `json:"total_division_percent"`
		TotalPercentZiswaf        TotalPercentZiswaf       `json:"total_ziswaf_percent"`
		TotalPrognosisPerMonth    TotalPrognosisPerMonth   `json:"total_prognosis_per_month"`
		TotalTransactionPerMonth  TotalTransactionPerMonth `json:"total_transaction_per_month"`
	}
	divisionReport struct {
		Total                     uint64 `json:"total"`
		TotalPrognosisRetail      uint64 `json:"total_prognosis_retail"`
		TotalPrognosisCorporate   uint64 `json:"total_prognosis_corporate"`
		TotalPrognosisUpz         uint64 `json:"total_prognosis_upz"`
		TotalPercentRetail        uint64 `json:"total_percent_retail"`
		TotalPercentCorporate     uint64 `json:"total_percent_corporate"`
		TotalPercentUpz           uint64 `json:"total_percent_upz"`
		TotalRitelZakatMaal       uint64 `json:"total_ritel_zakat_maal"`
		TotalRitelWakaf           uint64 `json:"total_ritel_wakaf"`
		TotalRitelZakatFitrah     uint64 `json:"total_ritel_zakat_fitrah"`
		TotalRitelInfaq           uint64 `json:"total_ritel_infaq"`
		TotalRitelKurban          uint64 `json:"total_ritel_kurban"`
		TotalRitelOther           uint64 `json:"total_ritel_other"`
		TotalCorporateZakatMaal   uint64 `json:"total_corporate_zakat_maal"`
		TotalCorporateWakaf       uint64 `json:"total_corporate_wakaf"`
		TotalCorporateZakatFitrah uint64 `json:"total_corporate_zakat_fitrah"`
		TotalCorporateInfaq       uint64 `json:"total_corporate_infaq"`
		TotalCorporateKurban      uint64 `json:"total_corporate_kurban"`
		TotalCorporateOther       uint64 `json:"total_corporate_other"`
		TotalUpzZakatMaal         uint64 `json:"total_upz_zakat_maal"`
		TotalUpzWakaf             uint64 `json:"total_upz_wakaf"`
		TotalUpzZakatFitrah       uint64 `json:"total_upz_zakat_fitrah"`
		TotalUpzInfaq             uint64 `json:"total_upz_infaq"`
		TotalUpzKurban            uint64 `json:"total_upz_kurban"`
		TotalUpzOther             uint64 `json:"total_upz_other"`
	}
	nominalReport struct {
		TotalGood             uint64           `json:"total_good"`
		TotalGoodCount        uint64           `json:"total_good_count"`
		TotalGoodCollect      int32            `json:"total_good_collect"`
		TotalGoodMoveCount    int32            `json:"total_good_move_count"`
		TotalGoodNotMoveCount int32            `json:"total_good_not_move_count"`
		TotalGoodFoodCount    int32            `json:"total_good_food_count"`
		TotalGoodOtherCount   int32            `json:"total_good_other_count"`
		TotalAllCash          uint64           `json:"total_all_cash"`
		TotalCashCount        uint64           `json:"total_cash_count"`
		TotalCash             uint64           `json:"total_cash"`
		TotalNonCashMuamalat  uint64           `json:"total_non_cash_muamalat"`
		TotalNonCashMandiri   uint64           `json:"total_non_cash_mandiri"`
		TotalNonCashBsm       uint64           `json:"total_non_cash_bsm"`
		TotalNonCashBri       uint64           `json:"total_non_cash_bri"`
		TotalNonCashBniLamp   uint64           `json:"total_non_cash_bni_lamp"`
		TotalNonCashBniSy     uint64           `json:"total_non_cash_bni_sy"`
		TotalNonCashBca       uint64           `json:"total_non_cash_bca"`
		TotalCashPercent      TotalCashPercent `json:"total_cash_percent"`
	}

	TotalPercentZiswaf struct {
		ZakatMaalPercent   uint64 `json:"zakat_maal_percent"`
		ZakatFitrahPercent uint64 `json:"zakat_fitrah_percent"`
		InfaqPercent       uint64 `json:"infaq_percent"`
		WaqafPercent       uint64 `json:"waqaf_percent"`
		QurbanPercent      uint64 `json:"qurban_percent"`
		OtherPercent       uint64 `json:"other_percent"`
	}

	TotalCashPercent struct {
		CashPercent            uint64 `json:"cash_percent"`
		NonCashMuamalatPercent uint64 `json:"non_cash_muamalat_percent"`
		NonCashMandiriPercent  uint64 `json:"non_cash_mandiri_percent"`
		NonCashBsmPercent      uint64 `json:"non_cash_bsm_percent"`
		NonCashBriPercent      uint64 `json:"non_cash_bri_percent"`
		NonCashBniLampPercent  uint64 `json:"non_cash_bni_lamp_percent"`
		NonCashBniSyPercent    uint64 `json:"non_cash_bni_sy_percent"`
		NonCashBcaPercent      uint64 `json:"non_cash_bca_percent"`
	}

	TotalPrognosisPerMonth struct {
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
	TotalTransactionPerMonth struct {
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

	TotalDivisionPercent struct {
		TotalDivisionUpzPercent       uint64 `json:"total_division_upz_percent"`
		TotalDivisionRetailPercent    uint64 `json:"total_division_retail_percent"`
		TotalDivisionCorporatePercent uint64 `json:"total_division_corporate_percent"`
	}
)

func SetResponse(domain ListReportResponseData, message string, success bool) ListReportResponse {
	return ListReportResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.ReportResponse) ListReportResponseData {
	dashAdmin := dashboardAdminResponse{}
	// dashOperator := dashboardOperatorResponse{}
	comReport := commonReport{}
	divReport := divisionReport{}
	nomReport := nominalReport{}
	progUpz := PrognosisUpz{}
	progRetail := PrognosisRetail{}
	progCorporate := PrognosisCorporate{}

	for _, v := range domain.Prognosis {
		if v.DivisionID == 1 {
			if v.Month == 1 {
				progUpz.January = v.Total
			} else if v.Month == 2 {
				progUpz.February = v.Total
			} else if v.Month == 3 {
				progUpz.March = v.Total
			} else if v.Month == 4 {
				progUpz.April = v.Total
			} else if v.Month == 5 {
				progUpz.May = v.Total
			} else if v.Month == 6 {
				progUpz.June = v.Total
			} else if v.Month == 7 {
				progUpz.July = v.Total
			} else if v.Month == 8 {
				progUpz.August = v.Total
			} else if v.Month == 9 {
				progUpz.September = v.Total
			} else if v.Month == 10 {
				progUpz.October = v.Total
			} else if v.Month == 11 {
				progUpz.November = v.Total
			} else if v.Month == 12 {
				progUpz.December = v.Total
			}
		} else if v.DivisionID == 2 {
			if v.Month == 1 {
				progRetail.January = v.Total
			} else if v.Month == 2 {
				progRetail.February = v.Total
			} else if v.Month == 3 {
				progRetail.March = v.Total
			} else if v.Month == 4 {
				progRetail.April = v.Total
			} else if v.Month == 5 {
				progRetail.May = v.Total
			} else if v.Month == 6 {
				progRetail.June = v.Total
			} else if v.Month == 7 {
				progRetail.July = v.Total
			} else if v.Month == 8 {
				progRetail.August = v.Total
			} else if v.Month == 9 {
				progRetail.September = v.Total
			} else if v.Month == 10 {
				progRetail.October = v.Total
			} else if v.Month == 11 {
				progRetail.November = v.Total
			} else if v.Month == 12 {
				progRetail.December = v.Total
			}
		} else if v.DivisionID == 3 {
			if v.Month == 1 {
				progCorporate.January = v.Total
			} else if v.Month == 2 {
				progCorporate.February = v.Total
			} else if v.Month == 3 {
				progCorporate.March = v.Total
			} else if v.Month == 4 {
				progCorporate.April = v.Total
			} else if v.Month == 5 {
				progCorporate.May = v.Total
			} else if v.Month == 6 {
				progCorporate.June = v.Total
			} else if v.Month == 7 {
				progCorporate.July = v.Total
			} else if v.Month == 8 {
				progCorporate.August = v.Total
			} else if v.Month == 9 {
				progCorporate.September = v.Total
			} else if v.Month == 10 {
				progCorporate.October = v.Total
			} else if v.Month == 11 {
				progCorporate.November = v.Total
			} else if v.Month == 12 {
				progCorporate.December = v.Total
			}
		}
	}

	dashAdmin.Total = domain.Total
	dashAdmin.TotalRowCount = domain.TotalRowCount
	dashAdmin.TotalRowDonorCount = domain.TotalRowDonorCount
	dashAdmin.TotalRowNewDonorCount = domain.TotalRowNewDonorCount
	dashAdmin.TotalZakatMaal = domain.TotalZakatMaal
	dashAdmin.TotalWakaf = domain.TotalWakaf
	dashAdmin.TotalZakatFitrah = domain.TotalZakatFitrah
	dashAdmin.TotalInfaq = domain.TotalInfaq
	dashAdmin.TotalKurban = domain.TotalKurban
	dashAdmin.TotalOther = domain.TotalOther
	dashAdmin.TotalRetail = domain.TotalRetail
	dashAdmin.TotalRetailRowCount = domain.TotalRetailRowCount
	dashAdmin.TotalCorporate = domain.TotalCorporate
	dashAdmin.TotalCorporateRowCount = domain.TotalCorporateRowCount
	dashAdmin.TotalUpz = domain.TotalUpz
	dashAdmin.TotalUpzRowCount = domain.TotalUpzRowCount
	dashAdmin.TotalPrognosisRetail = domain.TotalPrognosisRetail
	dashAdmin.TotalPrognosisCorporate = domain.TotalPrognosisCorporate
	dashAdmin.TotalPrognosisUpz = domain.TotalPrognosisUpz
	dashAdmin.TotalPercentRetail = domain.TotalPercentRetail
	dashAdmin.TotalPercentCorporate = domain.TotalPercentCorporate
	dashAdmin.TotalPercentUpz = domain.TotalPercentUpz
	dashAdmin.TotalZiswafPerMonth = domain.TotalZiswafPerMonth
	dashAdmin.TotalZakatMaalPerMonth = domain.TotalZakatMaalPerMonth
	dashAdmin.TotalZakatFitrahPerMonth = domain.TotalZakatFitrahPerMonth
	dashAdmin.TotalInfaqPerMonth = domain.TotalInfaqPerMonth
	dashAdmin.TotalWakafPerMonth = domain.TotalWakafPerMonth
	dashAdmin.TotalQurbanPerMonth = domain.TotalQurbanPerMonth
	dashAdmin.TotalOtherPerMonth = domain.TotalOtherPerMonth
	dashAdmin.TotalLastYear = domain.TotalTransactionLastYear

	comReport.Total = domain.Total
	comReport.TotalRowCount = domain.TotalRowCount
	comReport.TotalRowDonorCount = domain.TotalRowDonorCount
	comReport.TotalZakatMaal = domain.TotalZakatMaal
	comReport.TotalWakaf = domain.TotalWakaf
	comReport.TotalZakatFitrah = domain.TotalZakatFitrah
	comReport.TotalInfaq = domain.TotalInfaq
	comReport.TotalKurban = domain.TotalKurban
	comReport.TotalOther = domain.TotalOther
	comReport.TotalCompanyDonor = domain.TotalCompanyDonor
	comReport.TotalPersonDonor = domain.TotalPersonDonor
	comReport.TotalCompanyDonorRowCount = domain.TotalCompanyDonorRowCount
	comReport.TotalPersonDonorRowCount = domain.TotalPersonDonorRowCount
	comReport.TotalPercentCompanyDonor = domain.TotalPercentCompanyDonor
	comReport.TotalPercentPersonDonor = domain.TotalPercentPersonDonor
	comReport.TotalRetail = domain.TotalRetail
	comReport.TotalCorporate = domain.TotalCorporate
	comReport.TotalUpz = domain.TotalUpz
	comReport.TotalDivisionPercent = TotalDivisionPercent{
		TotalDivisionUpzPercent:       domain.TotalDivisionPercent.TotalDivisionUpzPercent,
		TotalDivisionRetailPercent:    domain.TotalDivisionPercent.TotalDivisionRetailPercent,
		TotalDivisionCorporatePercent: domain.TotalDivisionPercent.TotalDivisionCorporatePercent,
	}
	comReport.TotalPercentZiswaf = TotalPercentZiswaf{
		ZakatMaalPercent:   domain.TotalPercentZiswaf.ZakatMaalPercent,
		ZakatFitrahPercent: domain.TotalPercentZiswaf.ZakatFitrahPercent,
		InfaqPercent:       domain.TotalPercentZiswaf.InfaqPercent,
		WaqafPercent:       domain.TotalPercentZiswaf.WaqafPercent,
		QurbanPercent:      domain.TotalPercentZiswaf.QurbanPercent,
		OtherPercent:       domain.TotalPercentZiswaf.OtherPercent,
	}
	comReport.TotalPrognosisPerMonth = TotalPrognosisPerMonth{
		January:   domain.TotalPrognosisPerMonth.January,
		February:  domain.TotalPrognosisPerMonth.February,
		March:     domain.TotalPrognosisPerMonth.March,
		April:     domain.TotalPrognosisPerMonth.April,
		May:       domain.TotalPrognosisPerMonth.May,
		June:      domain.TotalPrognosisPerMonth.June,
		July:      domain.TotalPrognosisPerMonth.July,
		August:    domain.TotalPrognosisPerMonth.August,
		September: domain.TotalPrognosisPerMonth.September,
		October:   domain.TotalPrognosisPerMonth.October,
		November:  domain.TotalPrognosisPerMonth.November,
		December:  domain.TotalPrognosisPerMonth.December,
	}
	comReport.TotalTransactionPerMonth = TotalTransactionPerMonth{
		January:   domain.TotalTransactionPerMonth.January,
		February:  domain.TotalTransactionPerMonth.February,
		March:     domain.TotalTransactionPerMonth.March,
		April:     domain.TotalTransactionPerMonth.April,
		May:       domain.TotalTransactionPerMonth.May,
		June:      domain.TotalTransactionPerMonth.June,
		July:      domain.TotalTransactionPerMonth.July,
		August:    domain.TotalTransactionPerMonth.August,
		September: domain.TotalTransactionPerMonth.September,
		October:   domain.TotalTransactionPerMonth.October,
		November:  domain.TotalTransactionPerMonth.November,
		December:  domain.TotalTransactionPerMonth.December,
	}

	divReport.Total = domain.Total
	divReport.TotalPrognosisRetail = domain.TotalPrognosisRetail
	divReport.TotalPrognosisCorporate = domain.TotalPrognosisCorporate
	divReport.TotalPrognosisUpz = domain.TotalPrognosisUpz
	divReport.TotalPercentRetail = domain.TotalPercentRetail
	divReport.TotalPercentCorporate = domain.TotalPercentCorporate
	divReport.TotalPercentUpz = domain.TotalPercentUpz
	divReport.TotalRitelZakatMaal = domain.TotalRitelZakatMaal
	divReport.TotalRitelWakaf = domain.TotalRitelWakaf
	divReport.TotalRitelZakatFitrah = domain.TotalRitelZakatFitrah
	divReport.TotalRitelInfaq = domain.TotalRitelInfaq
	divReport.TotalRitelKurban = domain.TotalRitelKurban
	divReport.TotalRitelOther = domain.TotalRitelOther
	divReport.TotalCorporateZakatMaal = domain.TotalCorporateZakatMaal
	divReport.TotalCorporateWakaf = domain.TotalCorporateWakaf
	divReport.TotalCorporateZakatFitrah = domain.TotalCorporateZakatFitrah
	divReport.TotalCorporateInfaq = domain.TotalCorporateInfaq
	divReport.TotalCorporateKurban = domain.TotalCorporateKurban
	divReport.TotalCorporateOther = domain.TotalCorporateOther
	divReport.TotalUpzZakatMaal = domain.TotalUpzZakatMaal
	divReport.TotalUpzWakaf = domain.TotalUpzWakaf
	divReport.TotalUpzZakatFitrah = domain.TotalUpzZakatFitrah
	divReport.TotalUpzInfaq = domain.TotalUpzInfaq
	divReport.TotalUpzKurban = domain.TotalUpzKurban
	divReport.TotalUpzOther = domain.TotalUpzOther

	nomReport.TotalGood = domain.TotalGood
	nomReport.TotalGoodCount = domain.TotalGoodCount
	nomReport.TotalGoodCollect = domain.TotalGoodCollect
	nomReport.TotalGoodMoveCount = domain.TotalGoodMoveCount
	nomReport.TotalGoodNotMoveCount = domain.TotalGoodNotMoveCount
	nomReport.TotalGoodFoodCount = domain.TotalGoodFoodCount
	nomReport.TotalGoodOtherCount = domain.TotalGoodOtherCount
	nomReport.TotalAllCash = domain.TotalAllCash
	nomReport.TotalCashCount = domain.TotalCashCount
	nomReport.TotalCash = domain.TotalCash
	nomReport.TotalNonCashMuamalat = domain.TotalNonCashMuamalat
	nomReport.TotalNonCashMandiri = domain.TotalNonCashMandiri
	nomReport.TotalNonCashBsm = domain.TotalNonCashBsm
	nomReport.TotalNonCashBri = domain.TotalNonCashBri
	nomReport.TotalNonCashBniLamp = domain.TotalNonCashBniLamp
	nomReport.TotalNonCashBniSy = domain.TotalNonCashBniSy
	nomReport.TotalNonCashBca = domain.TotalNonCashBca
	nomReport.TotalCashPercent = TotalCashPercent{
		CashPercent:            domain.TotalCashPercent.CashPercent,
		NonCashMuamalatPercent: domain.TotalCashPercent.NonCashMuamalatPercent,
		NonCashMandiriPercent:  domain.TotalCashPercent.NonCashMandiriPercent,
		NonCashBsmPercent:      domain.TotalCashPercent.NonCashBsmPercent,
		NonCashBriPercent:      domain.TotalCashPercent.NonCashBriPercent,
		NonCashBniLampPercent:  domain.TotalCashPercent.NonCashBniLampPercent,
		NonCashBniSyPercent:    domain.TotalCashPercent.NonCashBniSyPercent,
		NonCashBcaPercent:      domain.TotalCashPercent.NonCashBcaPercent,
	}

	return ListReportResponseData{
		DashboardAdmin:     dashAdmin,
		CommonReport:       comReport,
		DivisionReport:     divReport,
		NominalReport:      nomReport,
		PrognosisUpz:       progUpz,
		PrognosisRetail:    progRetail,
		PrognosisCorporate: progCorporate,
		UpzPerMonth:        domain.TotalUpzPerMonth,
		RetailPerMonth:     domain.TotalRetailPerMonth,
		CorporatePerMonth:  domain.TotalCorporatePerMonth,
	}
}
