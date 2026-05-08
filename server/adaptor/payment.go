package adaptor

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/usecase"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentHandler struct {
	Usecase usecase.PaymentUsecase
	Log *zap.Logger
}

func NewPaymentHandler(u usecase.PaymentUsecase, log *zap.Logger) PaymentHandler {
	return PaymentHandler{
		Usecase: u,
		Log: log,
	}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req dto.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	res, err := h.Usecase.CreatePayment(c, req)
	if err == utils.ErrUnauthorized {
		utils.ResponseFailed(c, http.StatusUnauthorized, "failed to create payment", err.Error())
		return
	}
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to create payment", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "success to create payment", res)
}

func (h *PaymentHandler) PaymentCallback(c *gin.Context) {
	var req dto.PaymentCallback
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	err := h.Usecase.PaymentCallback(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to update payment", err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "success to update payment", nil)
}