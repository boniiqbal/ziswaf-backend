package list_report_operator

import (
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListReportOperatorResponse struct {
		base.BaseResponse
		Data ListReportOperatorResponseData `json:"data"`
	}

	ListReportOperatorResponseData struct {
		DashboardOperator      dashboardOperatorResponse     `json:"dashboard_operator"`
		CommonReport           commonReport                  `json:"common_report"`
		DivisionReport         divisionReport                `json:"division_report"`
		NominalReport          nominalReport                 `json:"nominal_report"`
		TotalZiswafPerDay      domain.TotalZiswafPerDay      `json:"total_ziswaf_per_day"`
		TotalZakatMaalPerDay   domain.TotalZakatMaalPerDay   `json:"total_zakat_maal_per_day"`
		TotalZakatFitrahPerDay domain.TotalZakatFitrahPerDay `json:"total_zakat_fitrah_per_day"`
		TotalInfaqPerDay       domain.TotalInfaqPerDay       `json:"total_infaq_per_day"`
		TotalWakafPerDay       domain.TotalWakafPerDay       `json:"total_wakaf_per_day"`
		TotalQurbanPerDay      domain.TotalQurbanPerDay      `json:"total_qurban_per_day"`
		TotalOtherPerDay       domain.TotalOtherPerDay       `json:"total_other_per_day"`
	}

	dashboardOperatorResponse struct {
		Total                 uint64                          `json:"total"`
		TotalRowCount         int                             `json:"total_row_count"`
		TotalRowDonorCount    int                             `json:"total_row_donor_count"`
		TotalRowNewDonorCount int                             `json:"total_row_new_donor_count"`
		TotalZakatMaal        uint64                          `json:"total_zakat_maal"`
		TotalWakaf            uint64                          `json:"total_wakaf"`
		TotalZakatFitrah      uint64                          `json:"total_zakat_fitrah"`
		TotalInfaq            uint64                          `json:"total_infaq"`
		TotalKurban           uint64                          `json:"total_kurban"`
		TotalOther            uint64                          `json:"total_other"`
		TotalPercentZiswaf    domain.TotalPercentZiswaf       `json:"total_percent_ziswaf"`
		TotalLastMonth        domain.TotalTransactionLastYear `json:"total_transaction_last_month"`
	}

	commonReport struct {
		Total                     uint64               `json:"total"`
		TotalRowCount             int                  `json:"total_row_count"`
		TotalRowDonorCount        int                  `json:"total_row_donor_count"`
		TotalZakatMaal            uint64               `json:"total_zakat_maal"`
		TotalWakaf                uint64               `json:"total_wakaf"`
		TotalZakatFitrah          uint64               `json:"total_zakat_fitrah"`
		TotalInfaq                uint64               `json:"total_infaq"`
		TotalKurban               uint64               `json:"total_kurban"`
		TotalOther                uint64               `json:"total_other"`
		TotalCompanyDonor         uint64               `json:"total_company_donor"`
		TotalPersonDonor          uint64               `json:"total_person_donor"`
		TotalCompanyDonorRowCount int                  `json:"total_company_donor_row_count"`
		TotalPersonDonorRowCount  int                  `json:"total_person_donor_row_count"`
		TotalPercentCompanyDonor  uint64               `json:"total_percent_company_donor"`
		TotalPercentPersonDonor   uint64               `json:"total_percent_person_donor"`
		TotalRetail               uint64               `json:"total_retail"`
		TotalCorporate            uint64               `json:"total_corporate"`
		TotalUpz                  uint64               `json:"total_upz"`
		TotalPercentZiswaf        TotalPercentZiswaf   `json:"total_ziswaf_percent"`
		TotalDivisionPercent      TotalDivisionPercent `json:"total_division_percent"`
	}
	divisionReport struct {
		Total                     uint64                      `json:"total"`
		TotalPercentRetail        uint64                      `json:"total_percent_retail"`
		TotalPercentCorporate     uint64                      `json:"total_percent_corporate"`
		TotalPercentUpz           uint64                      `json:"total_percent_upz"`
		TotalRitelZakatMaal       uint64                      `json:"total_ritel_zakat_maal"`
		TotalRitelWakaf           uint64                      `json:"total_ritel_wakaf"`
		TotalRitelZakatFitrah     uint64                      `json:"total_ritel_zakat_fitrah"`
		TotalRitelInfaq           uint64                      `json:"total_ritel_infaq"`
		TotalRitelKurban          uint64                      `json:"total_ritel_kurban"`
		TotalRitelOther           uint64                      `json:"total_ritel_other"`
		TotalCorporateZakatMaal   uint64                      `json:"total_corporate_zakat_maal"`
		TotalCorporateWakaf       uint64                      `json:"total_corporate_wakaf"`
		TotalCorporateZakatFitrah uint64                      `json:"total_corporate_zakat_fitrah"`
		TotalCorporateInfaq       uint64                      `json:"total_corporate_infaq"`
		TotalCorporateKurban      uint64                      `json:"total_corporate_kurban"`
		TotalCorporateOther       uint64                      `json:"total_corporate_other"`
		TotalUpzZakatMaal         uint64                      `json:"total_upz_zakat_maal"`
		TotalUpzWakaf             uint64                      `json:"total_upz_wakaf"`
		TotalUpzZakatFitrah       uint64                      `json:"total_upz_zakat_fitrah"`
		TotalUpzInfaq             uint64                      `json:"total_upz_infaq"`
		TotalUpzKurban            uint64                      `json:"total_upz_kurban"`
		TotalUpzOther             uint64                      `json:"total_upz_other"`
		TotalUpzPerDay            domain.TotalUpzPerDay       `json:"total_upz_per_day"`
		TotalRetailPerDay         domain.TotalRetailPerDay    `json:"total_retail_per_day"`
		TotalCorporatePerDay      domain.TotalCorporatePerDay `json:"total_corporate_per_day"`
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

	TotalDivisionPercent struct {
		TotalDivisionUpzPercent       uint64 `json:"total_division_upz_percent"`
		TotalDivisionRetailPercent    uint64 `json:"total_division_retail_percent"`
		TotalDivisionCorporatePercent uint64 `json:"total_division_corporate_percent"`
	}
)

func SetResponse(domain ListReportOperatorResponseData, message string, success bool) ListReportOperatorResponse {
	return ListReportOperatorResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.ReportOperatorResponse) ListReportOperatorResponseData {
	dashOperator := dashboardOperatorResponse{}
	comReport := commonReport{}
	divReport := divisionReport{}
	nomReport := nominalReport{}

	dashOperator.Total = domain.Total
	dashOperator.TotalRowCount = domain.TotalRowCount
	dashOperator.TotalRowDonorCount = domain.TotalRowDonorCount
	dashOperator.TotalRowNewDonorCount = domain.TotalRowNewDonorCount
	dashOperator.TotalZakatMaal = domain.TotalZakatMaal
	dashOperator.TotalWakaf = domain.TotalWakaf
	dashOperator.TotalZakatFitrah = domain.TotalZakatFitrah
	dashOperator.TotalInfaq = domain.TotalInfaq
	dashOperator.TotalKurban = domain.TotalKurban
	dashOperator.TotalOther = domain.TotalOther
	dashOperator.TotalLastMonth = domain.TotalTransactionLastMonth
	dashOperator.TotalPercentZiswaf = domain.TotalPercentZiswaf

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
	comReport.TotalPercentZiswaf = TotalPercentZiswaf{
		ZakatMaalPercent:   domain.TotalPercentZiswaf.ZakatMaalPercent,
		ZakatFitrahPercent: domain.TotalPercentZiswaf.ZakatFitrahPercent,
		InfaqPercent:       domain.TotalPercentZiswaf.InfaqPercent,
		WaqafPercent:       domain.TotalPercentZiswaf.WaqafPercent,
		QurbanPercent:      domain.TotalPercentZiswaf.QurbanPercent,
		OtherPercent:       domain.TotalPercentZiswaf.OtherPercent,
	}
	comReport.TotalDivisionPercent = TotalDivisionPercent{
		TotalDivisionUpzPercent:       domain.TotalDivisionPercent.TotalDivisionUpzPercent,
		TotalDivisionRetailPercent:    domain.TotalDivisionPercent.TotalDivisionRetailPercent,
		TotalDivisionCorporatePercent: domain.TotalDivisionPercent.TotalDivisionCorporatePercent,
	}

	divReport.Total = domain.Total
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
	divReport.TotalUpzPerDay = domain.TotalUpzPerDay
	divReport.TotalRetailPerDay = domain.TotalRetailPerDay
	divReport.TotalCorporatePerDay = domain.TotalCorporatePerDay

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

	return ListReportOperatorResponseData{
		DashboardOperator:      dashOperator,
		CommonReport:           comReport,
		DivisionReport:         divReport,
		NominalReport:          nomReport,
		TotalZiswafPerDay:      domain.TotalZiswafPerDay,
		TotalZakatMaalPerDay:   domain.TotalZakatMaalPerDay,
		TotalZakatFitrahPerDay: domain.TotalZakatFitrahPerDay,
		TotalInfaqPerDay:       domain.TotalInfaqPerDay,
		TotalWakafPerDay:       domain.TotalWakafPerDay,
		TotalQurbanPerDay:      domain.TotalQurbanPerDay,
		TotalOtherPerDay:       domain.TotalOtherPerDay,
	}
}
