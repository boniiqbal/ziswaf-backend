package show_prognosis

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowPrognosisHandler struct {
	request            infrastructure.Request
	repository         infrastructure.PrognosisRepository
	divisionRepository infrastructure.DivisionRepository
}

func NewShowPrognosisHandler(request infrastructure.Request, progRepo infrastructure.PrognosisRepository, divRepo infrastructure.DivisionRepository) ShowPrognosisHandler {
	return ShowPrognosisHandler{
		request:            request,
		repository:         progRepo,
		divisionRepository: divRepo,
	}
}

func (handler *ShowPrognosisHandler) ShowPrognosisHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	prognosis, err := handler.repository.GetPrognosisByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	division := handler.divisionRepository.GetListDivision(ctx)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(prognosis, division), "Show prognosis success", true))
}
