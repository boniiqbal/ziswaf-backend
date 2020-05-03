package update_user

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateUserHandler struct {
	request    infrastructure.Request
	repository infrastructure.UserRepository
}

func NewUpdateUserHandler(request infrastructure.Request, prdRepo infrastructure.UserRepository) UpdateUserHandler {
	return UpdateUserHandler{
		request:    request,
		repository: prdRepo,
	}
}

func (handler *UpdateUserHandler) UpdateUserHandler(c *gin.Context) {
	request := UpdateUserRequest{}
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatUint(acc.(uint64), 10)
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// Get user by id
	userData, err := handler.repository.GetUserById(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	byteDBPass := []byte(userData.Password)

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		errRequest := misc.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	// Check username already exist
	user, _ := handler.repository.GetUserByUsername(ctx, strings.ToLower(request.Data.Username))
	if userData.Username == strings.ToLower(request.Data.Username) {
		//
	} else {
		if strings.ToLower(request.Data.Username) == user.Username {
			c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Username sudah tersedia", false))
			return
		}
	}

	// Hashing password user
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(request.Data.Password), bcrypt.DefaultCost)
	if errHash != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(errHash.Error(), false))
		return
	}

	// Compare password user
	if error := bcrypt.CompareHashAndPassword(byteDBPass, []byte(request.Data.Password)); error != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Password salah", false))
		return
	}

	_, err = handler.repository.UpdateUser(ctx, RequestMapper(request, string(hashedPassword), strings.ToLower(request.Data.Username), accountID), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	usData, err := handler.repository.GetUserById(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(usData), "Update user berhasil", true))
}
