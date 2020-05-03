package change_password

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type ChangePasswordHandler struct {
	usrRepository infrastructure.UserRepository
}

func NewChangePasswordHandler(usRepo infrastructure.UserRepository) ChangePasswordHandler {
	return ChangePasswordHandler{
		usrRepository: usRepo,
	}
}

func (handler *ChangePasswordHandler) ChangePasswordHandler(c *gin.Context) {
	var (
		request    ChangeUserRequest
		byteDBPass []byte
	)

	accountID, _ := c.Get("UserId")
	accID := strconv.FormatUint(accountID.(uint64), 10)

	id := c.Param("id")
	changeType := c.Query("type")

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

	if changeType == "reset" {
		// Get User Account
		accountData, err := handler.usrRepository.GetUserById(ctx, accID)
		if err != nil {
			c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
			return
		}
		byteDBPass = []byte(accountData.Password)
	} else if changeType == "change" {
		userById, err := handler.usrRepository.GetUserById(ctx, id)
		if err != nil {
			c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
			return
		}
		byteDBPass = []byte(userById.Password)
	}

	byteReqPass := []byte(request.Data.Password)

	// Compare password user
	if error := bcrypt.CompareHashAndPassword(byteDBPass, byteReqPass); error != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Password salah", false))
		return
	}

	// Confirm New Password
	if request.Data.ConfirmPassword != request.Data.NewPassword {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Password konfirmasi salah", false))
		return
	}

	// Hashing new password user
	hashedNewPassword, errHash := bcrypt.GenerateFromPassword([]byte(request.Data.NewPassword), bcrypt.DefaultCost)
	if errHash != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(errHash.Error(), false))
		return
	}

	// Update User Data
	_, err := handler.usrRepository.UpdateUser(ctx, RequestMapper(string(hashedNewPassword), accID), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	userData, err := handler.usrRepository.GetUserById(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(userData), "Password berhasil diubah", true))
}
