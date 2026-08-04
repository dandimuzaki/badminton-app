package adaptor

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/usecase"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReservationHandler struct {
	Usecase usecase.ReservationUsecase
	Log *zap.Logger
}

func NewReservationHandler(u usecase.ReservationUsecase, log *zap.Logger) ReservationHandler {
	return ReservationHandler{
		Usecase: u,
		Log: log,
	}
}

func (h *ReservationHandler) GetReservationHistory(c *gin.Context) {
	reservations, err := h.Usecase.GetReservationHistory(c)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to get reservation history", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "success to get available reservation history", reservations)
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var query dto.ReservationQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	req, err := dto.ToReservationRequest(&query)
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	if err := h.Usecase.CreateReservation(c, *req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "failed to create reservation", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "success to create reservation", nil)
}

func (h *ReservationHandler) CancelReservation(c *gin.Context) {
	id, err := utils.GetUintParam(c, "id")
	if err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid reservation id", err.Error())
		return
	}

	if err := h.Usecase.CancelReservation(c, id); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "failed to create reservation", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "success to create reservation", nil)
}