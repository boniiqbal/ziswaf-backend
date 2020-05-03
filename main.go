package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// Auth Service
	_auth "ziswaf-backend/middleware"

	// configs and interfaces
	_db "ziswaf-backend/infrastructure/persistence/repository/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "*"},
		ExposeHeaders:    []string{"Accept", "Content-Length", "Content-Type", "Authorization", "Accept:Encoding"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("/api/v1") // initial route
	db := _db.DBInit()            // initial db configuration

	// login route
	LoginRoute(v1, db)

	// Image Route
	ImgRoute(v1, db)

	// Pdf Route
	GenerateReportRoute(v1, db)

	// middleware
	v1.Use(_auth.AuthenticationRequired())

	// register
	UserRoute(v1, db)
	CategoryRoute(v1, db)
	DivisionRoute(v1, db)
	DonorRoute(v1, db)
	ProvinceRoute(v1, db)
	TransactionRoute(v1, db)
	PrognosisRoute(v1, db)
	SchoolRoute(v1, db)
	EmployeeRoute(v1, db)
	StudentRoute(v1, db)
	RegencyRoute(v1, db)
	DistrictRoute(v1, db)
	VillageRoute(v1, db)

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal(fmt.Sprintf("PORT must be set [%s]", port))
	}

	router.Run(":" + port)
}

// Route User
func UserRoute(route *gin.RouterGroup, db *gorm.DB) {
	cpHandler := CreateUserHandler(db)
	lpHandler := ListUsersHandler(db)
	spHandler := ShowUserHandler(db)
	usHandler := UpdateUserHandler(db)
	chpHandler := ChangePasswordHandler(db)
	usdHandler := DeleteUserHandler(db)
	usuHandler := UpdateStatusHandler(db)

	v1 := route.Group("/manager")
	{
		v1.POST("/users", cpHandler.CreateUserHandler)
		v1.GET("/users", lpHandler.ListUsersHandler)
		v1.GET("/users/:id", spHandler.ShowUserHandler)
		v1.PUT("/users/:id", usHandler.UpdateUserHandler)
		v1.PUT("/users/:id/status", usuHandler.UpdateStatusHandler)
		v1.DELETE("/users/:id", usdHandler.DeleteUserHandler)
	}
	v2 := route.Group("/users")
	{
		v2.PUT("/change-password/:id", chpHandler.ChangePasswordHandler)
		v2.GET("/:id", spHandler.ShowUserHandler)
		v2.PUT("/update/:id", usHandler.UpdateUserHandler)
	}
}

// Login Route
func LoginRoute(route *gin.RouterGroup, db *gorm.DB) {
	loHandler := LoginHandler(db)

	v1 := route.Group("/login")
	{
		v1.POST("", loHandler.LoginHandler)
	}
}

// Category Route
func CategoryRoute(route *gin.RouterGroup, db *gorm.DB) {
	caHandler := CreateCategoryHandler(db)
	clHandler := ListCategoriesHandler(db)
	scHandler := ShowCategoryHandler(db)
	cuHandler := UpdateCategoryHandler(db)
	crtStCtrgHandler := CreateStatementCategoryHandler(db)
	lStCtrgHandler := ListStatementCategoriesHandler(db)
	upCtrgHandler := UpdateStatementCategoryHandler(db)
	deCtrgHandler := DeleteStatementCategoryHandler(db)
	rpTrxHandler := ListReportHandler(db)

	v1 := route.Group("/manager")
	{
		v1.POST("/category", caHandler.CreateCategoryHandler)
		v1.GET("/category", clHandler.ListCategoriesHandler)
		v1.GET("/statement/category", lStCtrgHandler.ListStatementCategoriesHandler)
		v1.GET("/category/:id", scHandler.ShowCategoryHandler)
		v1.PUT("/category/:id", cuHandler.UpdateCategoryHandler)
		v1.POST("/category/jenis", crtStCtrgHandler.CreateStatementCategoryHandler)
		v1.PUT("/statement/category/:id", upCtrgHandler.UpdateStatementCategoryHandler)
		v1.DELETE("/statement/category/:id", deCtrgHandler.DeleteStatementCategoryHandler)
		v1.GET("/transaction/report", rpTrxHandler.ListReportHandler)
	}

	v2 := route.Group("/category")
	{
		v2.GET("", clHandler.ListCategoriesHandler)
		v2.GET("/:id", scHandler.ShowCategoryHandler)
	}
}

// Division Route
func DivisionRoute(route *gin.RouterGroup, db *gorm.DB) {
	divcHandler := CreateDivisionHandler(db)
	divlHandler := ListDivisionsHandler(db)
	divsHandler := ShowDivisionHandler(db)
	divuHandler := UpdateDivisionHandler(db)

	v1 := route.Group("/manager")
	{
		v1.POST("/division", divcHandler.CreateDivisionHandler)
		v1.GET("/division", divlHandler.ListDivisionsHandler)
		v1.GET("/division/:id", divsHandler.ShowDivisionHandler)
		v1.PUT("/division/:id", divuHandler.UpdateDivisionHandler)
	}

	v2 := route.Group("/division")
	{
		v2.GET("", divlHandler.ListDivisionsHandler)
		v2.GET("/:id", divsHandler.ShowDivisionHandler)
	}
}

// Donor Route
func DonorRoute(route *gin.RouterGroup, db *gorm.DB) {
	doncHandler := CreateDonorHandler(db)
	donlHandler := ListDonorsHandler(db)
	donsHandler := ShowDonorHandler(db)
	donuHandler := UpdateDonorHandler(db)
	dondHandler := DeleteDonorHandler(db)

	v1 := route.Group("/donor")
	{
		v1.POST("", doncHandler.CreateDonorHandler)
		v1.GET("", donlHandler.ListDonorsHandler)
		v1.GET("/:id", donsHandler.ShowDonorHandler)
		v1.PUT("/:id", donuHandler.UpdateDonorHandler)
		v1.DELETE("/:id", dondHandler.DeleteDonorHandler)
	}
}

// Province Route
func ProvinceRoute(route *gin.RouterGroup, db *gorm.DB) {
	provincelHandler := ListProvincesHandler(db)
	provincesHandler := ShowProvinceHandler(db)

	v1 := route.Group("/province")
	{
		v1.GET("", provincelHandler.ListProvincesHandler)
		v1.GET("/:id", provincesHandler.ShowProvinceHandler)
	}
}

// Transaction Route
func TransactionRoute(route *gin.RouterGroup, db *gorm.DB) {
	crTrxHandler := CreateTransactionHandler(db)
	lsTrxHandler := ListTransactionHandler(db)
	shTrxHandler := ShowTransactionHandler(db)
	roTrxHandler := ListReportOperatorHandler(db)
	rdTrxHandler := ListReportDonationHandler(db)

	v1 := route.Group("/transaction")
	{
		v1.POST("", crTrxHandler.CreateTransactionHandler)
		v1.GET("/detail/:id", shTrxHandler.ShowTransactionHandler)
		v1.GET("", lsTrxHandler.ListTransactionHandler)
		v1.GET("/report-operator", roTrxHandler.ListReportOperatorHandler)
		v1.GET("/report-donation", rdTrxHandler.ListReportDonationHandler)
	}
}

// Prognosis Route
func PrognosisRoute(route *gin.RouterGroup, db *gorm.DB) {
	proglHandler := ListPrognosisHandler(db)
	progsHandler := ShowPrognosisHandler(db)
	progcHandler := CreatePrognosisHandler(db)
	progupHandler := UpdatePronosisHandler(db)

	v1 := route.Group("/prognosis")
	{
		v1.GET("", proglHandler.ListPrognosisHandler)
		v1.GET("/:id", progsHandler.ShowPrognosisHandler)
		v1.POST("", progcHandler.CreatePrognosisHandler)
		v1.PUT("/:id", progupHandler.UpdatePrognosisHandler)
	}
}

// School Route
func SchoolRoute(route *gin.RouterGroup, db *gorm.DB) {
	schlHandler := ListSchoolHandler(db)
	schsHandler := ShowSchoolHandler(db)
	schcHandler := CreateSchoolHandler(db)
	schupHandler := UpdateSchoolHandler(db)
	schdpHandler := DeleteSchoolHandler(db)
	schrdHandler := RecordSchoolHandler(db)

	v1 := route.Group("/school")
	{
		v1.GET("", schlHandler.ListSchoolHandler)
		v1.GET("/:id", schsHandler.ShowSchoolHandler)
		v1.POST("", schcHandler.CreateSchoolHandler)
		v1.PUT("/:id", schupHandler.UpdateSchoolHandler)
		v1.DELETE("/:id", schdpHandler.DeleteSchoolHandler)
	}

	route.GET("/record/school/:id", schrdHandler.RecordSchoolHandler)
}

// Employee Route
func EmployeeRoute(route *gin.RouterGroup, db *gorm.DB) {
	emplHandler := ListEmployeesHandler(db)
	empsHandler := ShowEmployeeHandler(db)
	empcHandler := CreateEmployeeHandler(db)
	empuHandler := UpdateEmployeeHandler(db)
	empdHandler := DeleteEmployeeHandler(db)

	v1 := route.Group("/employee")
	{
		v1.GET("", emplHandler.ListEmployeesHandler)
		v1.GET("/:id", empsHandler.ShowEmployeeHandler)
		v1.POST("", empcHandler.CreateEmployeeHandler)
		v1.PUT("/:id", empuHandler.UpdateEmployeeHandler)
		v1.DELETE("/:id", empdHandler.DeleteEmployeeHandler)
	}
}

// Image Route
func ImgRoute(route *gin.RouterGroup, db *gorm.DB) {
	route.GET("/images/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.File("./attachments/images/default-image.png")
		} else {
			c.File("./attachments/images/" + id)
		}
	})
}

// Student Route
func StudentRoute(route *gin.RouterGroup, db *gorm.DB) {
	stdlHandler := ListStudentHandler(db)
	stdcHandler := CreateStudentsHandler(db)
	stduHandler := UpdateStudentHandler(db)
	stdsHandler := ShowStudentHandler(db)
	stddHandler := DeleteStudentHandler(db)

	v1 := route.Group("/student")
	{
		v1.GET("", stdlHandler.ListStudentHandler)
		v1.GET("/:id", stdsHandler.ShowStudentHandler)
		v1.POST("", stdcHandler.CreateStudentsHandler)
		v1.PUT("/:id", stduHandler.UpdateStudentHandler)
		v1.DELETE("/:id", stddHandler.DeleteStudentHandler)
	}
}

// Regency Route
func RegencyRoute(route *gin.RouterGroup, db *gorm.DB) {
	rgcHandler := ListRegenciesHandler(db)

	v1 := route.Group("/regency")
	{
		v1.GET("", rgcHandler.ListRegenciesHandler)
	}
}

// District Route
func DistrictRoute(route *gin.RouterGroup, db *gorm.DB) {
	drcHandler := ListDistrictHandler(db)

	v1 := route.Group("/district")
	{
		v1.GET("", drcHandler.ListDistrictHandler)
	}
}

// Village Route
func VillageRoute(route *gin.RouterGroup, db *gorm.DB) {
	villHandler := ListVillageHandler(db)

	v1 := route.Group("/village")
	{
		v1.GET("", villHandler.ListVillageHandler)
	}
}

func GenerateReportRoute(route *gin.RouterGroup, db *gorm.DB) {
	xpTrxHandler := ExportPdfHandler(db)
	xkTrxHandler := ExportKwitansiHandler(db)

	v1 := route.Group("/transaction")
	{
		v1.GET("/export-pdf", xpTrxHandler.ExportPdfHandler)
		v1.GET("/export-kwitansi/:id", xkTrxHandler.ExportKwitansiHandler)
	}
}
