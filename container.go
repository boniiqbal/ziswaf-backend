//+build wireinject

package main

import (
	"github.com/google/wire"

	// user routes
	"ziswaf-backend/application/use_case/user/change_password"
	"ziswaf-backend/application/use_case/user/create_user"
	"ziswaf-backend/application/use_case/user/delete_user"
	"ziswaf-backend/application/use_case/user/list_users"
	"ziswaf-backend/application/use_case/user/show_user"
	"ziswaf-backend/application/use_case/user/update_status"
	"ziswaf-backend/application/use_case/user/update_user"

	// authentication routes
	"ziswaf-backend/application/use_case/authentication/login"

	// category routes
	"ziswaf-backend/application/use_case/category/create_category"
	"ziswaf-backend/application/use_case/category/create_statement_category"
	"ziswaf-backend/application/use_case/category/delete_statement_category"
	"ziswaf-backend/application/use_case/category/list_category"
	"ziswaf-backend/application/use_case/category/list_statement_category"
	"ziswaf-backend/application/use_case/category/show_category"
	"ziswaf-backend/application/use_case/category/update_category"
	"ziswaf-backend/application/use_case/category/update_statement_category"

	// division routes
	"ziswaf-backend/application/use_case/division/create_division"
	"ziswaf-backend/application/use_case/division/list_division"
	"ziswaf-backend/application/use_case/division/show_division"
	"ziswaf-backend/application/use_case/division/update_division"

	// donor routes
	"ziswaf-backend/application/use_case/donor/create_donor"
	"ziswaf-backend/application/use_case/donor/delete_donor"
	"ziswaf-backend/application/use_case/donor/list_donor"
	"ziswaf-backend/application/use_case/donor/show_donor"
	"ziswaf-backend/application/use_case/donor/update_donor"

	// province routes
	"ziswaf-backend/application/use_case/province/list_province"
	"ziswaf-backend/application/use_case/province/show_province"

	// transaction routes
	"ziswaf-backend/application/use_case/transaction/create_transaction"
	"ziswaf-backend/application/use_case/transaction/export_kwitansi"
	"ziswaf-backend/application/use_case/transaction/export_pdf"
	"ziswaf-backend/application/use_case/transaction/list_report"
	"ziswaf-backend/application/use_case/transaction/list_report_donation"
	"ziswaf-backend/application/use_case/transaction/list_report_operator"
	"ziswaf-backend/application/use_case/transaction/list_transaction"
	"ziswaf-backend/application/use_case/transaction/show_transaction"

	// prognosis routes
	"ziswaf-backend/application/use_case/prognosis/create_prognosis"
	"ziswaf-backend/application/use_case/prognosis/list_prognosis"
	"ziswaf-backend/application/use_case/prognosis/show_prognosis"
	"ziswaf-backend/application/use_case/prognosis/update_prognosis"

	// school routes
	"ziswaf-backend/application/use_case/school/create_school"
	"ziswaf-backend/application/use_case/school/delete_school"
	"ziswaf-backend/application/use_case/school/list_school"
	"ziswaf-backend/application/use_case/school/record_school"
	"ziswaf-backend/application/use_case/school/show_school"
	"ziswaf-backend/application/use_case/school/update_school"

	// employee routes
	"ziswaf-backend/application/use_case/employee/create_employee"
	"ziswaf-backend/application/use_case/employee/delete_employee"
	"ziswaf-backend/application/use_case/employee/list_employee"
	"ziswaf-backend/application/use_case/employee/show_employee"
	"ziswaf-backend/application/use_case/employee/update_employee"

	// student routes
	"ziswaf-backend/application/use_case/student/create_student"
	"ziswaf-backend/application/use_case/student/delete_student"
	"ziswaf-backend/application/use_case/student/list_student"
	"ziswaf-backend/application/use_case/student/show_student"
	"ziswaf-backend/application/use_case/student/update_student"

	// regency routes
	"ziswaf-backend/application/use_case/regency/list_regency"

	// district routes
	"ziswaf-backend/application/use_case/district/list_district"

	// village routes
	"ziswaf-backend/application/use_case/village/list_village"

	repo "ziswaf-backend/infrastructure/persistence/repository/db"
	request "ziswaf-backend/infrastructure/transport/http"

	"github.com/jinzhu/gorm"
)

func LoginHandler(db *gorm.DB) login.LoginHandler {
	wire.Build(repo.NewUserRepository, repo.NewLoginRepository, login.NewLoginHandler)
	return login.LoginHandler{}
}

func CreateUserHandler(db *gorm.DB) create_user.CreateUserHandler {
	wire.Build(repo.NewUserRepository, request.NewRequest, create_user.NewCreateUserHandler)
	return create_user.CreateUserHandler{}
}

func ListUsersHandler(db *gorm.DB) list_users.ListUsersHandler {
	wire.Build(request.NewRequest, repo.NewUserRepository, repo.NewEmployeeRepository, list_users.NewListUsersHandler)
	return list_users.ListUsersHandler{}
}

func ShowUserHandler(db *gorm.DB) show_user.ShowUserHandler {
	wire.Build(request.NewRequest, repo.NewUserRepository, show_user.NewShowUserHandler)
	return show_user.ShowUserHandler{}
}

func UpdateUserHandler(db *gorm.DB) update_user.UpdateUserHandler {
	wire.Build(request.NewRequest, repo.NewUserRepository, update_user.NewUpdateUserHandler)
	return update_user.UpdateUserHandler{}
}

func DeleteUserHandler(db *gorm.DB) delete_user.DeleteUserHandler {
	wire.Build(request.NewRequest, repo.NewUserRepository, delete_user.NewDeleteUserHandler)
	return delete_user.DeleteUserHandler{}
}

func ChangePasswordHandler(db *gorm.DB) change_password.ChangePasswordHandler {
	wire.Build(repo.NewUserRepository, change_password.NewChangePasswordHandler)
	return change_password.ChangePasswordHandler{}
}

func UpdateStatusHandler(db *gorm.DB) update_status.UpdateStatusHandler {
	wire.Build(request.NewRequest, repo.NewUserRepository, update_status.NewUpdateStatusHandler)
	return update_status.UpdateStatusHandler{}
}

func CreateCategoryHandler(db *gorm.DB) create_category.CreateCategoryHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, create_category.NewCreateCategoryHandler)
	return create_category.CreateCategoryHandler{}
}

func ListCategoriesHandler(db *gorm.DB) list_category.ListCategoriesHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, list_category.NewListCategoriesHandler)
	return list_category.ListCategoriesHandler{}
}

func ShowCategoryHandler(db *gorm.DB) show_category.ShowCategoryHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, show_category.NewShowCategoryHandler)
	return show_category.ShowCategoryHandler{}
}

func UpdateCategoryHandler(db *gorm.DB) update_category.UpdateCategoryHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, update_category.NewUpdateCategoryHandler)
	return update_category.UpdateCategoryHandler{}
}

func CreateStatementCategoryHandler(db *gorm.DB) create_statement_category.CreateStatementCategoryHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, create_statement_category.NewCreateStatementCategoryHandler)
	return create_statement_category.CreateStatementCategoryHandler{}
}

func ListStatementCategoriesHandler(db *gorm.DB) list_statement_category.ListStatementCategoriesHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, list_statement_category.NewListStatementCategoriesHandler)
	return list_statement_category.ListStatementCategoriesHandler{}
}

func UpdateStatementCategoryHandler(db *gorm.DB) update_statement_category.UpdateStatementCategoryHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, update_statement_category.NewUpdateStatementCategoryHandler)
	return update_statement_category.UpdateStatementCategoryHandler{}
}

func DeleteStatementCategoryHandler(db *gorm.DB) delete_statement_category.DeleteStatementCategoryHandler {
	wire.Build(request.NewRequest, repo.NewCategoryRepository, delete_statement_category.NewDeleteStatementCategoryHandler)
	return delete_statement_category.DeleteStatementCategoryHandler{}
}

func CreateDivisionHandler(db *gorm.DB) create_division.CreateDivisionHandler {
	wire.Build(request.NewRequest, repo.NewDivisionRepository, create_division.NewCreateDivisionHandler)
	return create_division.CreateDivisionHandler{}
}

func ListDivisionsHandler(db *gorm.DB) list_division.ListDivisionsHandler {
	wire.Build(request.NewRequest, repo.NewDivisionRepository, list_division.NewListDivisionsHandler)
	return list_division.ListDivisionsHandler{}
}

func ShowDivisionHandler(db *gorm.DB) show_division.ShowDivisionHandler {
	wire.Build(request.NewRequest, repo.NewDivisionRepository, show_division.NewShowDivisionHandler)
	return show_division.ShowDivisionHandler{}
}

func UpdateDivisionHandler(db *gorm.DB) update_division.UpdateDivisionHandler {
	wire.Build(request.NewRequest, repo.NewDivisionRepository, update_division.NewUpdateDivisionHandler)
	return update_division.UpdateDivisionHandler{}
}

func CreateDonorHandler(db *gorm.DB) create_donor.CreateDonorHandler {
	wire.Build(request.NewRequest, repo.NewDonorRepository, create_donor.NewCreateDonorHandler)
	return create_donor.CreateDonorHandler{}
}

func ListDonorsHandler(db *gorm.DB) list_donor.ListDonorsHandler {
	wire.Build(request.NewRequest, repo.NewDonorRepository, repo.NewRegencyRepository, list_donor.NewListDonorsHandler)
	return list_donor.ListDonorsHandler{}
}

func ShowDonorHandler(db *gorm.DB) show_donor.ShowDonorHandler {
	wire.Build(request.NewRequest, repo.NewDonorRepository, repo.NewRegencyRepository, show_donor.NewShowDonorHandler)
	return show_donor.ShowDonorHandler{}
}

func UpdateDonorHandler(db *gorm.DB) update_donor.UpdateDonorHandler {
	wire.Build(request.NewRequest, repo.NewDonorRepository, update_donor.NewUpdateDonorHandler)
	return update_donor.UpdateDonorHandler{}
}

func DeleteDonorHandler(db *gorm.DB) delete_donor.DeleteDonorHandler {
	wire.Build(request.NewRequest, repo.NewDonorRepository, delete_donor.NewDeleteDonorHandler)
	return delete_donor.DeleteDonorHandler{}
}

func ListProvincesHandler(db *gorm.DB) list_province.ListProvincesHandler {
	wire.Build(request.NewRequest, repo.NewProvinceRepository, list_province.NewListProvincesHandler)
	return list_province.ListProvincesHandler{}
}

func ShowProvinceHandler(db *gorm.DB) show_province.ShowProvinceHandler {
	wire.Build(request.NewRequest, repo.NewProvinceRepository, show_province.NewShowProvinceHandler)
	return show_province.ShowProvinceHandler{}
}

func CreateTransactionHandler(db *gorm.DB) create_transaction.CreateTransactionHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, repo.NewUserRepository, repo.NewEmployeeRepository, create_transaction.NewCreateTransactionHandler)
	return create_transaction.CreateTransactionHandler{}
}

func ListTransactionHandler(db *gorm.DB) list_transaction.ListTransactionHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, repo.NewDonorRepository, list_transaction.NewListTransactionHandler)
	return list_transaction.ListTransactionHandler{}
}

func ShowTransactionHandler(db *gorm.DB) show_transaction.ShowTransactionHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, repo.NewUserRepository, show_transaction.NewShowTransactionHandler)
	return show_transaction.ShowTransactionHandler{}
}

func ListReportHandler(db *gorm.DB) list_report.ListReportHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, list_report.NewListReportHandler)
	return list_report.ListReportHandler{}
}

func ListReportOperatorHandler(db *gorm.DB) list_report_operator.ListReportOperatorHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, list_report_operator.NewListReportOperatorHandler)
	return list_report_operator.ListReportOperatorHandler{}
}

func ListReportDonationHandler(db *gorm.DB) list_report_donation.ListReportDonationHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, repo.NewCategoryRepository, list_report_donation.NewListReportDonationHandler)
	return list_report_donation.ListReportDonationHandler{}
}

func ExportPdfHandler(db *gorm.DB) export_pdf.ExportPdfHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, export_pdf.NewExportPdfHandler)
	return export_pdf.ExportPdfHandler{}
}

func ExportKwitansiHandler(db *gorm.DB) export_kwitansi.ExportKwitansiHandler {
	wire.Build(request.NewRequest, repo.NewTransactionRepository, repo.NewUserRepository, export_kwitansi.NewExportKwitansiHandler)
	return export_kwitansi.ExportKwitansiHandler{}
}

func CreatePrognosisHandler(db *gorm.DB) create_prognosis.CreatePrognosisHandler {
	wire.Build(request.NewRequest, repo.NewPrognosisRepository, create_prognosis.NewCreatePrognosisHandler)
	return create_prognosis.CreatePrognosisHandler{}
}

func ListPrognosisHandler(db *gorm.DB) list_prognosis.ListPrognosisHandler {
	wire.Build(request.NewRequest, repo.NewPrognosisRepository, repo.NewDivisionRepository, list_prognosis.NewListPrognosisHandler)
	return list_prognosis.ListPrognosisHandler{}
}

func ShowPrognosisHandler(db *gorm.DB) show_prognosis.ShowPrognosisHandler {
	wire.Build(request.NewRequest, repo.NewPrognosisRepository, repo.NewDivisionRepository, show_prognosis.NewShowPrognosisHandler)
	return show_prognosis.ShowPrognosisHandler{}
}

func UpdatePronosisHandler(db *gorm.DB) update_prognosis.UpdatePrognosisHandler {
	wire.Build(request.NewRequest, repo.NewPrognosisRepository, update_prognosis.NewUpdatePrognosisHandler)
	return update_prognosis.UpdatePrognosisHandler{}
}

func CreateSchoolHandler(db *gorm.DB) create_school.CreateSchoolHandler {
	wire.Build(request.NewRequest, repo.NewSchoolRepository, create_school.NewCreateSchoolHandler)
	return create_school.CreateSchoolHandler{}
}

func ListSchoolHandler(db *gorm.DB) list_school.ListSchoolHandler {
	wire.Build(request.NewRequest, repo.NewSchoolRepository, list_school.NewListSchoolHandler)
	return list_school.ListSchoolHandler{}
}

func ShowSchoolHandler(db *gorm.DB) show_school.ShowSchoolHandler {
	wire.Build(request.NewRequest, repo.NewSchoolRepository, show_school.NewShowSchoolHandler)
	return show_school.ShowSchoolHandler{}
}

func UpdateSchoolHandler(db *gorm.DB) update_school.UpdateSchoolHandler {
	wire.Build(request.NewRequest, repo.NewSchoolRepository, update_school.NewUpdateSchoolHandler)
	return update_school.UpdateSchoolHandler{}
}

func DeleteSchoolHandler(db *gorm.DB) delete_school.DeleteSchoolHandler {
	wire.Build(request.NewRequest, repo.NewSchoolRepository, delete_school.NewDeleteSchoolHandler)
	return delete_school.DeleteSchoolHandler{}
}

func RecordSchoolHandler(db *gorm.DB) record_school.RecordSchoolHandler {
	wire.Build(request.NewRequest, repo.NewSchoolRepository, record_school.NewRecordSchoolHandler)
	return record_school.RecordSchoolHandler{}
}

func CreateEmployeeHandler(db *gorm.DB) create_employee.CreateEmployeeHandler {
	wire.Build(request.NewRequest, repo.NewEmployeeRepository, create_employee.NewCreateEmployeeHandler)
	return create_employee.CreateEmployeeHandler{}
}

func ListEmployeesHandler(db *gorm.DB) list_employee.ListEmployeesHandler {
	wire.Build(request.NewRequest, repo.NewEmployeeRepository, list_employee.NewListEmployeesHandler)
	return list_employee.ListEmployeesHandler{}
}

func ShowEmployeeHandler(db *gorm.DB) show_employee.ShowEmployeeHandler {
	wire.Build(request.NewRequest, repo.NewEmployeeRepository, show_employee.NewShowEmployeeHandler)
	return show_employee.ShowEmployeeHandler{}
}

func UpdateEmployeeHandler(db *gorm.DB) update_employee.UpdateEmployeeHandler {
	wire.Build(request.NewRequest, repo.NewEmployeeRepository, update_employee.NewUpdateEmployeeHandler)
	return update_employee.UpdateEmployeeHandler{}
}

func DeleteEmployeeHandler(db *gorm.DB) delete_employee.DeleteEmployeeHandler {
	wire.Build(request.NewRequest, repo.NewEmployeeRepository, delete_employee.NewDeleteEmployeeHandler)
	return delete_employee.DeleteEmployeeHandler{}
}

func CreateStudentsHandler(db *gorm.DB) create_student.CreateStudentsHandler {
	wire.Build(request.NewRequest, repo.NewStudentRepository, repo.NewEmployeeRepository, create_student.NewCreateStudentsHandler)
	return create_student.CreateStudentsHandler{}
}

func ListStudentHandler(db *gorm.DB) list_student.ListStudentHandler {
	wire.Build(request.NewRequest, repo.NewStudentRepository, list_student.NewListStudentHandler)
	return list_student.ListStudentHandler{}
}

func ShowStudentHandler(db *gorm.DB) show_student.ShowStudentHandler {
	wire.Build(request.NewRequest, repo.NewStudentRepository, show_student.NewShowStudentHandler)
	return show_student.ShowStudentHandler{}
}

func UpdateStudentHandler(db *gorm.DB) update_student.UpdateStudentHandler {
	wire.Build(request.NewRequest, repo.NewStudentRepository, update_student.NewUpdateStudentHandler)
	return update_student.UpdateStudentHandler{}
}

func DeleteStudentHandler(db *gorm.DB) delete_student.DeleteStudentHandler {
	wire.Build(request.NewRequest, repo.NewStudentRepository, delete_student.NewDeleteStudentHandler)
	return delete_student.DeleteStudentHandler{}
}

func ListRegenciesHandler(db *gorm.DB) list_regency.ListRegenciesHandler {
	wire.Build(request.NewRequest, repo.NewRegencyRepository, list_regency.NewListRegenciesHandler)
	return list_regency.ListRegenciesHandler{}
}

func ListDistrictHandler(db *gorm.DB) list_district.ListDistrictHandler {
	wire.Build(request.NewRequest, repo.NewDistrictRepository, list_district.NewListDistrictHandler)
	return list_district.ListDistrictHandler{}
}

func ListVillageHandler(db *gorm.DB) list_village.ListVillageHandler {
	wire.Build(request.NewRequest, repo.NewVillageRepository, list_village.NewListVillageHandler)
	return list_village.ListVillageHandler{}
}
