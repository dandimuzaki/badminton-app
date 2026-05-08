package adaptor

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/usecase"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Usecase usecase.AuthUsecase
	Log *zap.Logger
	Config utils.Configuration
}

func NewAuthHandler(u usecase.AuthUsecase, log *zap.Logger) AuthHandler {
	return AuthHandler{
		Usecase: u,
		Log: log,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	resp, token, err := h.Usecase.Register(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to register", err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
    Name:     "token",
    Value:    *token,
    Path:     "/",
    MaxAge:   3600 * 24,
    HttpOnly: true,
    Secure:   true,
    SameSite: http.SameSiteNoneMode,
    Domain:   h.Config.ServerHost,
	})

	utils.ResponseSuccess(c, http.StatusOK, "success to register", resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseFailed(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	resp, token, err := h.Usecase.Login(c, req)
	if err != nil {
		utils.ResponseFailed(c, http.StatusInternalServerError, "failed to register", err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
    Name:     "token",
    Value:    *token,
    Path:     "/",
    MaxAge:   3600 * 24,
    HttpOnly: true,
    Secure:   true,
    SameSite: http.SameSiteNoneMode,
    Domain:   h.Config.ServerHost,
	})

	utils.ResponseSuccess(c, http.StatusOK, "success to register", resp)
}