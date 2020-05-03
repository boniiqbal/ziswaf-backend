package create_user

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

type CreateUserHandler struct {
	request    infrastructure.Request
	repository infrastructure.UserRepository
}

func NewCreateUserHandler(request infrastructure.Request, prdRepo infrastructure.UserRepository) CreateUserHandler {
	return CreateUserHandler{
		request:    request,
		repository: prdRepo,
	}
}

func (handler *CreateUserHandler) CreateUserHandler(c *gin.Context) {
	request := CreateUserRequest{}
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatUint(acc.(uint64), 10)

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
	user, _ := handler.repository.GetUserByUsername(ctx, strings.ToLower(request.Data.Username))
	if strings.ToLower(request.Data.Username) == user.Username {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Username sudah tersedia", false))
		return
	}

	// Hashing password user
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(request.Data.Password), bcrypt.DefaultCost)
	if errHash != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(errHash.Error(), false))
		return
	}

	prd, err := handler.repository.CreateUser(ctx, RequestMapper(request, string(hashedPassword), strings.ToLower(request.Data.Username), accountID))

	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(prd), "User berhasil dibuat", true))
}
