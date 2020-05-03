package update_donor

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateDonorHandler struct {
	request    infrastructure.Request
	repository infrastructure.DonorRepository
}

func NewUpdateDonorHandler(request infrastructure.Request, donorRepo infrastructure.DonorRepository) UpdateDonorHandler {
	return UpdateDonorHandler{
		request:    request,
		repository: donorRepo,
	}
}

func (handler *UpdateDonorHandler) UpdateDonorHandler(c *gin.Context) {
	request := UpdateDonorRequest{}
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	donorData, err := handler.repository.GetDonorByID(ctx, id)
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

	// Check donor by name and phone
	checkDonor, err := handler.repository.GetDonorByNamePhone(ctx, RequestMapper(request))
	if err == nil {
		if request.Data.Phone == donorData.Phone && request.Data.Name == donorData.Name {
		} else {
			if request.Data.Phone == checkDonor.Phone && request.Data.Name == checkDonor.Name {
				c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Nama dan No. Telepon sudah terdaftar", false))
				return
			}
		}
	}

	donData, err := handler.repository.UpdateDonor(ctx, RequestMapper(request), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(donData), "Update donatur berhasil", true))
}
