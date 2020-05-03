package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type TransactionRepository interface {
	CreateTransaction(context.Context, domain.Transaction) (domain.Transaction, error)
	CreateCash(context.Context, domain.Cash) (domain.Cash, error)
	CreateGoods(context.Context, domain.Goods) (domain.Goods, error)
	GetListTransaction(context.Context, domain.TransactionFilter, []domain.Donor) ([]domain.Transaction, int)
	GetListReport(context.Context, domain.TransactionFilter) (domain.ReportResponse, error)
	ShowTransaction(context.Context, string) (domain.Transaction, error)
	GetListReportOperator(context.Context, domain.TransactionFilter) (domain.ReportOperatorResponse, error)
	ExportPdf(context.Context, domain.TransactionFilter) (domain.ExportPdf, error)
	GetListDonationReport(context.Context, domain.TransactionFilter, []domain.StatementCategory) (domain.ReportStatementCategory, error)
}
