package adaptor

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/usecase"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TimeslotHandler struct {
	Usecase usecase.TimeslotUsecase
	Log *zap.Logger
}

func NewTimeslotHandler(u usecase.TimeslotUsecase, log *zap.Logger) TimeslotHandler {
	return TimeslotHandler{
		Usecase: u,
		Log: log,
	}
}

func (h *TimeslotHandler) GetAvailableTimeslots(c *gin.Context) {
	var query dto.AvailableTimeslotQuery
	if err := c.BindQuery(&query); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	req, err := dto.ToTimeslotRequest(&query)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	timeslots, err := h.Usecase.GetAvailableTimeslots(c, *req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to get available timeslots", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "success to get available timeslots", timeslots)
}