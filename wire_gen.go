// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/jinzhu/gorm"
	"ziswaf-backend/application/use_case/authentication/login"
	"ziswaf-backend/application/use_case/category/create_category"
	"ziswaf-backend/application/use_case/category/create_statement_category"
	"ziswaf-backend/application/use_case/category/delete_statement_category"
	"ziswaf-backend/application/use_case/category/list_category"
	"ziswaf-backend/application/use_case/category/list_statement_category"
	"ziswaf-backend/application/use_case/category/show_category"
	"ziswaf-backend/application/use_case/category/update_category"
	"ziswaf-backend/application/use_case/category/update_statement_category"
	"ziswaf-backend/application/use_case/district/list_district"
	"ziswaf-backend/application/use_case/division/create_division"
	"ziswaf-backend/application/use_case/division/list_division"
	"ziswaf-backend/application/use_case/division/show_division"
	"ziswaf-backend/application/use_case/division/update_division"
	"ziswaf-backend/application/use_case/donor/create_donor"
	"ziswaf-backend/application/use_case/donor/delete_donor"
	"ziswaf-backend/application/use_case/donor/list_donor"
	"ziswaf-backend/application/use_case/donor/show_donor"
	"ziswaf-backend/application/use_case/donor/update_donor"
	"ziswaf-backend/application/use_case/employee/create_employee"
	"ziswaf-backend/application/use_case/employee/delete_employee"
	"ziswaf-backend/application/use_case/employee/list_employee"
	"ziswaf-backend/application/use_case/employee/show_employee"
	"ziswaf-backend/application/use_case/employee/update_employee"
	"ziswaf-backend/application/use_case/prognosis/create_prognosis"
	"ziswaf-backend/application/use_case/prognosis/list_prognosis"
	"ziswaf-backend/application/use_case/prognosis/show_prognosis"
	"ziswaf-backend/application/use_case/prognosis/update_prognosis"
	"ziswaf-backend/application/use_case/province/list_province"
	"ziswaf-backend/application/use_case/province/show_province"
	"ziswaf-backend/application/use_case/regency/list_regency"
	"ziswaf-backend/application/use_case/school/create_school"
	"ziswaf-backend/application/use_case/school/delete_school"
	"ziswaf-backend/application/use_case/school/list_school"
	"ziswaf-backend/application/use_case/school/record_school"
	"ziswaf-backend/application/use_case/school/show_school"
	"ziswaf-backend/application/use_case/school/update_school"
	"ziswaf-backend/application/use_case/student/create_student"
	"ziswaf-backend/application/use_case/student/delete_student"
	"ziswaf-backend/application/use_case/student/list_student"
	"ziswaf-backend/application/use_case/student/show_student"
	"ziswaf-backend/application/use_case/student/update_student"
	"ziswaf-backend/application/use_case/transaction/create_transaction"
	"ziswaf-backend/application/use_case/transaction/export_kwitansi"
	"ziswaf-backend/application/use_case/transaction/export_pdf"
	"ziswaf-backend/application/use_case/transaction/list_report"
	"ziswaf-backend/application/use_case/transaction/list_report_donation"
	"ziswaf-backend/application/use_case/transaction/list_report_operator"
	"ziswaf-backend/application/use_case/transaction/list_transaction"
	"ziswaf-backend/application/use_case/transaction/show_transaction"
	"ziswaf-backend/application/use_case/user/change_password"
	"ziswaf-backend/application/use_case/user/create_user"
	"ziswaf-backend/application/use_case/user/delete_user"
	"ziswaf-backend/application/use_case/user/list_users"
	"ziswaf-backend/application/use_case/user/show_user"
	"ziswaf-backend/application/use_case/user/update_status"
	"ziswaf-backend/application/use_case/user/update_user"
	"ziswaf-backend/application/use_case/village/list_village"
	"ziswaf-backend/infrastructure/persistence/repository/db"
	"ziswaf-backend/infrastructure/transport/http"
)

// Injectors from container.go:

func LoginHandler(db2 *gorm.DB) login.LoginHandler {
	loginRepository := db.NewLoginRepository(db2)
	userRepository := db.NewUserRepository(db2)
	loginHandler := login.NewLoginHandler(loginRepository, userRepository)
	return loginHandler
}

func CreateUserHandler(db2 *gorm.DB) create_user.CreateUserHandler {
	infrastructureRequest := request.NewRequest()
	userRepository := db.NewUserRepository(db2)
	createUserHandler := create_user.NewCreateUserHandler(infrastructureRequest, userRepository)
	return createUserHandler
}

func ListUsersHandler(db2 *gorm.DB) list_users.ListUsersHandler {
	infrastructureRequest := request.NewRequest()
	userRepository := db.NewUserRepository(db2)
	employeeRepository := db.NewEmployeeRepository(db2)
	listUsersHandler := list_users.NewListUsersHandler(infrastructureRequest, userRepository, employeeRepository)
	return listUsersHandler
}

func ShowUserHandler(db2 *gorm.DB) show_user.ShowUserHandler {
	infrastructureRequest := request.NewRequest()
	userRepository := db.NewUserRepository(db2)
	showUserHandler := show_user.NewShowUserHandler(infrastructureRequest, userRepository)
	return showUserHandler
}

func UpdateUserHandler(db2 *gorm.DB) update_user.UpdateUserHandler {
	infrastructureRequest := request.NewRequest()
	userRepository := db.NewUserRepository(db2)
	updateUserHandler := update_user.NewUpdateUserHandler(infrastructureRequest, userRepository)
	return updateUserHandler
}

func DeleteUserHandler(db2 *gorm.DB) delete_user.DeleteUserHandler {
	infrastructureRequest := request.NewRequest()
	userRepository := db.NewUserRepository(db2)
	deleteUserHandler := delete_user.NewDeleteUserHandler(infrastructureRequest, userRepository)
	return deleteUserHandler
}

func ChangePasswordHandler(db2 *gorm.DB) change_password.ChangePasswordHandler {
	userRepository := db.NewUserRepository(db2)
	changePasswordHandler := change_password.NewChangePasswordHandler(userRepository)
	return changePasswordHandler
}

func UpdateStatusHandler(db2 *gorm.DB) update_status.UpdateStatusHandler {
	infrastructureRequest := request.NewRequest()
	userRepository := db.NewUserRepository(db2)
	updateStatusHandler := update_status.NewUpdateStatusHandler(infrastructureRequest, userRepository)
	return updateStatusHandler
}

func CreateCategoryHandler(db2 *gorm.DB) create_category.CreateCategoryHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	createCategoryHandler := create_category.NewCreateCategoryHandler(infrastructureRequest, categoryRepository)
	return createCategoryHandler
}

func ListCategoriesHandler(db2 *gorm.DB) list_category.ListCategoriesHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	listCategoriesHandler := list_category.NewListCategoriesHandler(infrastructureRequest, categoryRepository)
	return listCategoriesHandler
}

func ShowCategoryHandler(db2 *gorm.DB) show_category.ShowCategoryHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	showCategoryHandler := show_category.NewShowCategoryHandler(infrastructureRequest, categoryRepository)
	return showCategoryHandler
}

func UpdateCategoryHandler(db2 *gorm.DB) update_category.UpdateCategoryHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	updateCategoryHandler := update_category.NewUpdateCategoryHandler(infrastructureRequest, categoryRepository)
	return updateCategoryHandler
}

func CreateStatementCategoryHandler(db2 *gorm.DB) create_statement_category.CreateStatementCategoryHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	createStatementCategoryHandler := create_statement_category.NewCreateStatementCategoryHandler(infrastructureRequest, categoryRepository)
	return createStatementCategoryHandler
}

func ListStatementCategoriesHandler(db2 *gorm.DB) list_statement_category.ListStatementCategoriesHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	listStatementCategoriesHandler := list_statement_category.NewListStatementCategoriesHandler(infrastructureRequest, categoryRepository)
	return listStatementCategoriesHandler
}

func UpdateStatementCategoryHandler(db2 *gorm.DB) update_statement_category.UpdateStatementCategoryHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	updateStatementCategoryHandler := update_statement_category.NewUpdateStatementCategoryHandler(infrastructureRequest, categoryRepository)
	return updateStatementCategoryHandler
}

func DeleteStatementCategoryHandler(db2 *gorm.DB) delete_statement_category.DeleteStatementCategoryHandler {
	infrastructureRequest := request.NewRequest()
	categoryRepository := db.NewCategoryRepository(db2)
	deleteStatementCategoryHandler := delete_statement_category.NewDeleteStatementCategoryHandler(infrastructureRequest, categoryRepository)
	return deleteStatementCategoryHandler
}

func CreateDivisionHandler(db2 *gorm.DB) create_division.CreateDivisionHandler {
	infrastructureRequest := request.NewRequest()
	divisionRepository := db.NewDivisionRepository(db2)
	createDivisionHandler := create_division.NewCreateDivisionHandler(infrastructureRequest, divisionRepository)
	return createDivisionHandler
}

func ListDivisionsHandler(db2 *gorm.DB) list_division.ListDivisionsHandler {
	infrastructureRequest := request.NewRequest()
	divisionRepository := db.NewDivisionRepository(db2)
	listDivisionsHandler := list_division.NewListDivisionsHandler(infrastructureRequest, divisionRepository)
	return listDivisionsHandler
}

func ShowDivisionHandler(db2 *gorm.DB) show_division.ShowDivisionHandler {
	infrastructureRequest := request.NewRequest()
	divisionRepository := db.NewDivisionRepository(db2)
	showDivisionHandler := show_division.NewShowDivisionHandler(infrastructureRequest, divisionRepository)
	return showDivisionHandler
}

func UpdateDivisionHandler(db2 *gorm.DB) update_division.UpdateDivisionHandler {
	infrastructureRequest := request.NewRequest()
	divisionRepository := db.NewDivisionRepository(db2)
	updateDivisionHandler := update_division.NewUpdateDivisionHandler(infrastructureRequest, divisionRepository)
	return updateDivisionHandler
}

func CreateDonorHandler(db2 *gorm.DB) create_donor.CreateDonorHandler {
	infrastructureRequest := request.NewRequest()
	donorRepository := db.NewDonorRepository(db2)
	createDonorHandler := create_donor.NewCreateDonorHandler(infrastructureRequest, donorRepository)
	return createDonorHandler
}

func ListDonorsHandler(db2 *gorm.DB) list_donor.ListDonorsHandler {
	infrastructureRequest := request.NewRequest()
	donorRepository := db.NewDonorRepository(db2)
	regencyRepository := db.NewRegencyRepository(db2)
	listDonorsHandler := list_donor.NewListDonorsHandler(infrastructureRequest, donorRepository, regencyRepository)
	return listDonorsHandler
}

func ShowDonorHandler(db2 *gorm.DB) show_donor.ShowDonorHandler {
	infrastructureRequest := request.NewRequest()
	donorRepository := db.NewDonorRepository(db2)
	regencyRepository := db.NewRegencyRepository(db2)
	showDonorHandler := show_donor.NewShowDonorHandler(infrastructureRequest, donorRepository, regencyRepository)
	return showDonorHandler
}

func UpdateDonorHandler(db2 *gorm.DB) update_donor.UpdateDonorHandler {
	infrastructureRequest := request.NewRequest()
	donorRepository := db.NewDonorRepository(db2)
	updateDonorHandler := update_donor.NewUpdateDonorHandler(infrastructureRequest, donorRepository)
	return updateDonorHandler
}

func DeleteDonorHandler(db2 *gorm.DB) delete_donor.DeleteDonorHandler {
	infrastructureRequest := request.NewRequest()
	donorRepository := db.NewDonorRepository(db2)
	deleteDonorHandler := delete_donor.NewDeleteDonorHandler(infrastructureRequest, donorRepository)
	return deleteDonorHandler
}

func ListProvincesHandler(db2 *gorm.DB) list_province.ListProvincesHandler {
	infrastructureRequest := request.NewRequest()
	provinceRepository := db.NewProvinceRepository(db2)
	listProvincesHandler := list_province.NewListProvincesHandler(infrastructureRequest, provinceRepository)
	return listProvincesHandler
}

func ShowProvinceHandler(db2 *gorm.DB) show_province.ShowProvinceHandler {
	infrastructureRequest := request.NewRequest()
	provinceRepository := db.NewProvinceRepository(db2)
	showProvinceHandler := show_province.NewShowProvinceHandler(infrastructureRequest, provinceRepository)
	return showProvinceHandler
}

func CreateTransactionHandler(db2 *gorm.DB) create_transaction.CreateTransactionHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	employeeRepository := db.NewEmployeeRepository(db2)
	userRepository := db.NewUserRepository(db2)
	createTransactionHandler := create_transaction.NewCreateTransactionHandler(infrastructureRequest, transactionRepository, employeeRepository, userRepository)
	return createTransactionHandler
}

func ListTransactionHandler(db2 *gorm.DB) list_transaction.ListTransactionHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	donorRepository := db.NewDonorRepository(db2)
	listTransactionHandler := list_transaction.NewListTransactionHandler(infrastructureRequest, transactionRepository, donorRepository)
	return listTransactionHandler
}

func ShowTransactionHandler(db2 *gorm.DB) show_transaction.ShowTransactionHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	userRepository := db.NewUserRepository(db2)
	showTransactionHandler := show_transaction.NewShowTransactionHandler(infrastructureRequest, transactionRepository, userRepository)
	return showTransactionHandler
}

func ListReportHandler(db2 *gorm.DB) list_report.ListReportHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	listReportHandler := list_report.NewListReportHandler(infrastructureRequest, transactionRepository)
	return listReportHandler
}

func ListReportOperatorHandler(db2 *gorm.DB) list_report_operator.ListReportOperatorHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	listReportOperatorHandler := list_report_operator.NewListReportOperatorHandler(infrastructureRequest, transactionRepository)
	return listReportOperatorHandler
}

func ListReportDonationHandler(db2 *gorm.DB) list_report_donation.ListReportDonationHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	categoryRepository := db.NewCategoryRepository(db2)
	listReportDonationHandler := list_report_donation.NewListReportDonationHandler(infrastructureRequest, transactionRepository, categoryRepository)
	return listReportDonationHandler
}

func ExportPdfHandler(db2 *gorm.DB) export_pdf.ExportPdfHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	exportPdfHandler := export_pdf.NewExportPdfHandler(infrastructureRequest, transactionRepository)
	return exportPdfHandler
}

func ExportKwitansiHandler(db2 *gorm.DB) export_kwitansi.ExportKwitansiHandler {
	infrastructureRequest := request.NewRequest()
	transactionRepository := db.NewTransactionRepository(db2)
	userRepository := db.NewUserRepository(db2)
	exportKwitansiHandler := export_kwitansi.NewExportKwitansiHandler(infrastructureRequest, transactionRepository, userRepository)
	return exportKwitansiHandler
}

func CreatePrognosisHandler(db2 *gorm.DB) create_prognosis.CreatePrognosisHandler {
	infrastructureRequest := request.NewRequest()
	prognosisRepository := db.NewPrognosisRepository(db2)
	createPrognosisHandler := create_prognosis.NewCreatePrognosisHandler(infrastructureRequest, prognosisRepository)
	return createPrognosisHandler
}

func ListPrognosisHandler(db2 *gorm.DB) list_prognosis.ListPrognosisHandler {
	infrastructureRequest := request.NewRequest()
	prognosisRepository := db.NewPrognosisRepository(db2)
	divisionRepository := db.NewDivisionRepository(db2)
	listPrognosisHandler := list_prognosis.NewListPrognosisHandler(infrastructureRequest, prognosisRepository, divisionRepository)
	return listPrognosisHandler
}

func ShowPrognosisHandler(db2 *gorm.DB) show_prognosis.ShowPrognosisHandler {
	infrastructureRequest := request.NewRequest()
	prognosisRepository := db.NewPrognosisRepository(db2)
	divisionRepository := db.NewDivisionRepository(db2)
	showPrognosisHandler := show_prognosis.NewShowPrognosisHandler(infrastructureRequest, prognosisRepository, divisionRepository)
	return showPrognosisHandler
}

func UpdatePronosisHandler(db2 *gorm.DB) update_prognosis.UpdatePrognosisHandler {
	infrastructureRequest := request.NewRequest()
	prognosisRepository := db.NewPrognosisRepository(db2)
	updatePrognosisHandler := update_prognosis.NewUpdatePrognosisHandler(infrastructureRequest, prognosisRepository)
	return updatePrognosisHandler
}

func CreateSchoolHandler(db2 *gorm.DB) create_school.CreateSchoolHandler {
	infrastructureRequest := request.NewRequest()
	schoolRepository := db.NewSchoolRepository(db2)
	createSchoolHandler := create_school.NewCreateSchoolHandler(infrastructureRequest, schoolRepository)
	return createSchoolHandler
}

func ListSchoolHandler(db2 *gorm.DB) list_school.ListSchoolHandler {
	infrastructureRequest := request.NewRequest()
	schoolRepository := db.NewSchoolRepository(db2)
	listSchoolHandler := list_school.NewListSchoolHandler(infrastructureRequest, schoolRepository)
	return listSchoolHandler
}

func ShowSchoolHandler(db2 *gorm.DB) show_school.ShowSchoolHandler {
	infrastructureRequest := request.NewRequest()
	schoolRepository := db.NewSchoolRepository(db2)
	showSchoolHandler := show_school.NewShowSchoolHandler(infrastructureRequest, schoolRepository)
	return showSchoolHandler
}

func UpdateSchoolHandler(db2 *gorm.DB) update_school.UpdateSchoolHandler {
	infrastructureRequest := request.NewRequest()
	schoolRepository := db.NewSchoolRepository(db2)
	updateSchoolHandler := update_school.NewUpdateSchoolHandler(infrastructureRequest, schoolRepository)
	return updateSchoolHandler
}

func DeleteSchoolHandler(db2 *gorm.DB) delete_school.DeleteSchoolHandler {
	infrastructureRequest := request.NewRequest()
	schoolRepository := db.NewSchoolRepository(db2)
	deleteSchoolHandler := delete_school.NewDeleteSchoolHandler(infrastructureRequest, schoolRepository)
	return deleteSchoolHandler
}

func RecordSchoolHandler(db2 *gorm.DB) record_school.RecordSchoolHandler {
	infrastructureRequest := request.NewRequest()
	schoolRepository := db.NewSchoolRepository(db2)
	recordSchoolHandler := record_school.NewRecordSchoolHandler(infrastructureRequest, schoolRepository)
	return recordSchoolHandler
}

func CreateEmployeeHandler(db2 *gorm.DB) create_employee.CreateEmployeeHandler {
	infrastructureRequest := request.NewRequest()
	employeeRepository := db.NewEmployeeRepository(db2)
	createEmployeeHandler := create_employee.NewCreateEmployeeHandler(infrastructureRequest, employeeRepository)
	return createEmployeeHandler
}

func ListEmployeesHandler(db2 *gorm.DB) list_employee.ListEmployeesHandler {
	infrastructureRequest := request.NewRequest()
	employeeRepository := db.NewEmployeeRepository(db2)
	listEmployeesHandler := list_employee.NewListEmployeesHandler(infrastructureRequest, employeeRepository)
	return listEmployeesHandler
}

func ShowEmployeeHandler(db2 *gorm.DB) show_employee.ShowEmployeeHandler {
	infrastructureRequest := request.NewRequest()
	employeeRepository := db.NewEmployeeRepository(db2)
	showEmployeeHandler := show_employee.NewShowEmployeeHandler(infrastructureRequest, employeeRepository)
	return showEmployeeHandler
}

func UpdateEmployeeHandler(db2 *gorm.DB) update_employee.UpdateEmployeeHandler {
	infrastructureRequest := request.NewRequest()
	employeeRepository := db.NewEmployeeRepository(db2)
	updateEmployeeHandler := update_employee.NewUpdateEmployeeHandler(infrastructureRequest, employeeRepository)
	return updateEmployeeHandler
}

func DeleteEmployeeHandler(db2 *gorm.DB) delete_employee.DeleteEmployeeHandler {
	infrastructureRequest := request.NewRequest()
	employeeRepository := db.NewEmployeeRepository(db2)
	deleteEmployeeHandler := delete_employee.NewDeleteEmployeeHandler(infrastructureRequest, employeeRepository)
	return deleteEmployeeHandler
}

func CreateStudentsHandler(db2 *gorm.DB) create_student.CreateStudentsHandler {
	infrastructureRequest := request.NewRequest()
	studentRepository := db.NewStudentRepository(db2)
	employeeRepository := db.NewEmployeeRepository(db2)
	createStudentsHandler := create_student.NewCreateStudentsHandler(infrastructureRequest, studentRepository, employeeRepository)
	return createStudentsHandler
}

func ListStudentHandler(db2 *gorm.DB) list_student.ListStudentHandler {
	infrastructureRequest := request.NewRequest()
	studentRepository := db.NewStudentRepository(db2)
	listStudentHandler := list_student.NewListStudentHandler(infrastructureRequest, studentRepository)
	return listStudentHandler
}

func ShowStudentHandler(db2 *gorm.DB) show_student.ShowStudentHandler {
	infrastructureRequest := request.NewRequest()
	studentRepository := db.NewStudentRepository(db2)
	showStudentHandler := show_student.NewShowStudentHandler(infrastructureRequest, studentRepository)
	return showStudentHandler
}

func UpdateStudentHandler(db2 *gorm.DB) update_student.UpdateStudentHandler {
	infrastructureRequest := request.NewRequest()
	studentRepository := db.NewStudentRepository(db2)
	updateStudentHandler := update_student.NewUpdateStudentHandler(infrastructureRequest, studentRepository)
	return updateStudentHandler
}

func DeleteStudentHandler(db2 *gorm.DB) delete_student.DeleteStudentHandler {
	infrastructureRequest := request.NewRequest()
	studentRepository := db.NewStudentRepository(db2)
	deleteStudentHandler := delete_student.NewDeleteStudentHandler(infrastructureRequest, studentRepository)
	return deleteStudentHandler
}

func ListRegenciesHandler(db2 *gorm.DB) list_regency.ListRegenciesHandler {
	infrastructureRequest := request.NewRequest()
	regencyRepository := db.NewRegencyRepository(db2)
	listRegenciesHandler := list_regency.NewListRegenciesHandler(infrastructureRequest, regencyRepository)
	return listRegenciesHandler
}

func ListDistrictHandler(db2 *gorm.DB) list_district.ListDistrictHandler {
	infrastructureRequest := request.NewRequest()
	districtRepository := db.NewDistrictRepository(db2)
	listDistrictHandler := list_district.NewListDistrictHandler(infrastructureRequest, districtRepository)
	return listDistrictHandler
}

func ListVillageHandler(db2 *gorm.DB) list_village.ListVillageHandler {
	infrastructureRequest := request.NewRequest()
	villageRepository := db.NewVillageRepository(db2)
	listVillageHandler := list_village.NewListVillageHandler(infrastructureRequest, villageRepository)
	return listVillageHandler
}