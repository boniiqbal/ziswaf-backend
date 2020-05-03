package login

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/dgrijalva/jwt-go"

	"github.com/refactory-id/go-core-package/response"
)

type LoginHandler struct {
	repository    infrastructure.LoginRepository
	usrRepository infrastructure.UserRepository
}

func NewLoginHandler(prdRepo infrastructure.LoginRepository, usrRepo infrastructure.UserRepository) LoginHandler {
	return LoginHandler{
		repository:    prdRepo,
		usrRepository: usrRepo,
	}
}

func (handler *LoginHandler) LoginHandler(c *gin.Context) {
	request := LoginRequest{}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		errRequest := misc.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	// Get user by username
	user, err := handler.usrRepository.GetUserByUsername(ctx, request.Data.Username)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	if user.Status != 1 {
		c.JSON(misc.GetErrorStatusCode(401), response.SetMessage("Akun ini tidak aktif", false))
		return
	}

	byteDBPass := []byte(user.Password)
	byteReqPass := []byte(request.Data.Password)

	// Compare password user
	if error := bcrypt.CompareHashAndPassword(byteDBPass, byteReqPass); error != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Password Anda Salah", false))
		return
	}

	// Create the Claims
	claims := misc.CreateClaims(user.ID, user.Name, user.Role, user.EmployeeID, misc.EXPIRED_AT)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims) //Create Token
	signed, err := token.SignedString([]byte("secret"))

	// Login Handler
	_, errCreateAccessToken := handler.repository.Login(ctx, RequestMapper(user.ID, signed), user.ID)
	if errCreateAccessToken != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(errCreateAccessToken.Error(), false))
		return
	}

	userID := strconv.FormatUint(user.ID, 10)

	// Get User By ID
	userData, errGetUser := handler.usrRepository.GetUserById(ctx, userID)
	if errGetUser != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(errGetUser.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(userData), "Login Berhasil", true))
}
