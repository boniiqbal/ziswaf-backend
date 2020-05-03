package update_prognosis

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdatePrognosisHandler struct {
	request    infrastructure.Request
	repository infrastructure.PrognosisRepository
}

func NewUpdatePrognosisHandler(request infrastructure.Request, progRepo infrastructure.PrognosisRepository) UpdatePrognosisHandler {
	return UpdatePrognosisHandler{
		request:    request,
		repository: progRepo,
	}
}

func (handler *UpdatePrognosisHandler) UpdatePrognosisHandler(c *gin.Context) {
	request := UpdatePronosisRequest{}
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err := handler.repository.GetPrognosisByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
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

	donData, err := handler.repository.UpdatePrognosis(ctx, RequestMapper(request), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(donData), "Update prognosis berhasil", true))
}
