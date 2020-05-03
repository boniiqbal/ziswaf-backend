package create_prognosis

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type CreatePrognosisHandler struct {
	request    infrastructure.Request
	repository infrastructure.PrognosisRepository
}

func NewCreatePrognosisHandler(request infrastructure.Request, prognosisRepo infrastructure.PrognosisRepository) CreatePrognosisHandler {
	return CreatePrognosisHandler{
		request:    request,
		repository: prognosisRepo,
	}
}

func (handler *CreatePrognosisHandler) CreatePrognosisHandler(c *gin.Context) {
	request := CreatePrognosisRequest{}

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

	prognosis, err := handler.repository.CreatePrognosis(ctx, RequestMapper(request))
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(prognosis), "Prognosis berhasil dibuat", true))
}
