package list_users

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	UserFilterRequest struct {
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

func RequestMapper(req UserFilterRequest) domain.UserFilter {
	return domain.UserFilter{
		Search:       req.Search,
		SchoolID:     req.SchoolID,
		Role:         req.Role,
		CreatedStart: req.CreatedStart,
		CreatedEnd:   req.CreatedEnd,
		LoginStart:   req.LoginStart,
		LoginEnd:     req.LoginEnd,
		Sort:         req.Sort,
		Page:         req.Page,
		Limit:        req.Limit,
	}
}
