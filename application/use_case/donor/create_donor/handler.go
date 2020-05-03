package create_donor

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type CreateDonorHandler struct {
	request    infrastructure.Request
	repository infrastructure.DonorRepository
}

func NewCreateDonorHandler(request infrastructure.Request, donorRepo infrastructure.DonorRepository) CreateDonorHandler {
	return CreateDonorHandler{
		request:    request,
		repository: donorRepo,
	}
}

func (handler *CreateDonorHandler) CreateDonorHandler(c *gin.Context) {
	request := CreateDonorRequest{}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	donorType := c.Query("type")
	var donorID uint64

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		errRequest := misc.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	donorData, err := handler.repository.GetDonorByNamePhone(ctx, RequestMapper(request))
	if err == nil {
		donorID = donorData.ID
	} else {
		donorID = 0
	}

	if donorType == "" && donorID != 0 {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Nama dan Nomor Handphone yang sama ditemukan dalam database", false))
		return
	}

	if donorType == "update" && donorID == 0 {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Donatur tidak ditemukan", false))
		return
	}

	donor, err := handler.repository.CreateDonor(ctx, RequestMapper(request), donorID)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(donor), "Donatur berhasil dibuat", true))
}
