package misc

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	UserID     uint64
	Name       string
	Role       int
	EmployeeID uint64
	jwt.StandardClaims
}

// GetExpiryTime for jwt
func GetExpiryTime(exp time.Duration) int64 {
	return time.Now().Add(time.Hour * exp).Unix()
}

// CreateClaims .
func CreateClaims(userID uint64, name string, role int, employeeID uint64, exp time.Duration) MyCustomClaims {
	return MyCustomClaims{
		userID,
		name,
		role,
		employeeID,
		jwt.StandardClaims{
			ExpiresAt: GetExpiryTime(exp),
			IssuedAt:  time.Now().Unix(),
		},
	}
}
