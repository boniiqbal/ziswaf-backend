package db

import (
	"context"
	"time"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type loginRepository struct {
	DB *gorm.DB
}

func NewLoginRepository(DB *gorm.DB) infrastructure.LoginRepository {
	return &loginRepository{
		DB: DB,
	}
}

func (p *loginRepository) Login(ctx context.Context, accessToken domain.AccessToken, userID uint64) (domain.AccessToken, error) {
	accToken := domain.AccessToken{}
	user := domain.User{}
	userData := domain.User{}
	user.LastLogin = time.Now()

	errUpUser := p.DB.Model(&userData).Where("id = ?", userID).Update(&user).Error
	if errUpUser != nil {
		return accessToken, errUpUser
	}

	checkToken := p.DB.Where("user_id = ?", userID).First(&accToken).RecordNotFound()
	if checkToken {
		err := p.DB.Create(&accessToken).Error
		if err != nil {
			return accessToken, err
		}
		return accessToken, nil

	} else {
		if err := p.DB.Model(&accToken).Where("user_id = ?", userID).Update(&accessToken).Error; err != nil {
			return accessToken, err
		}
		return accessToken, nil
	}
}

func (p *loginRepository) GetAccessTokenByUserID(ctx context.Context, userID uint64) bool {
	accessToken := domain.AccessToken{}

	userData := p.DB.Where("user_id = ?", userID).First(&accessToken).RecordNotFound()
	if userData {
		return true
	}
	return false
}
