package adaptor

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/usecase"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CourtHandler struct {
	Usecase usecase.CourtUsecase
	Log *zap.Logger
}

func NewCourtHandler(u usecase.CourtUsecase, log *zap.Logger) CourtHandler {
	return CourtHandler{
		Usecase: u,
		Log: log,
	}
}

func (h *CourtHandler) GetAvailableCourts(c *gin.Context) {
	var query dto.AvailableCourtQuery
	if err := c.BindQuery(&query); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	req, err := dto.ToCourtRequest(&query)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	courts, err := h.Usecase.GetAvailableCourts(c, *req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to get available courts", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "success to get available courts", courts)
}