package entities

import (
	"time"
	base "ziswaf-backend/domain/infrastructure"
)

// UserRecord struct models, collection of record user management
type (
	User struct {
		base.ModelSoftDelete
		Name        string      `gorm:"column:name" json:"name"`
		Username    string      `gorm:"column:username" json:"username"`
		Password    string      `gorm:"column:password" json:"password"`
		EmployeeID  uint64      `gorm:"column:employee_id" json:"employee_id"`
		CreatedBy   string      `gorm:"column:created_by" json:"created_by"`
		UpdatedBy   string      `gorm:"column:updated_by" json:"updated_by"`
		Status      int         `gorm:"column:status" json:"status"`
		Role        int         `gorm:"column:role" json:"role"`
		LastLogin   time.Time   `json:"last_login"`
		AccessToken AccessToken `json:"access_token"`
		Employee    Employee    `json:"employee"`
	}

	UserFilter struct {
		Search       string `json:"search"`
		SchoolID     string `json:"school_id"`
		Role         string `json:"role"`
		CreatedStart string `json:"created_start"`
		CreatedEnd   string `json:"created_end"`
		LoginStart   string `json:"login_start"`
		LoginEnd     string `json:"login_end"`
		Sort         string `json:"Sort"`
		Page         string `json:"Page"`
		Limit        string `json:"Limit"`
	}
)
