package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

// LoginData struct models, collection of record access token
type (
	AccessToken struct {
		base.Model
		UserID    uint64 `json:"user_id"`
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expiredAt"`
	}
)
