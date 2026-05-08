package infra

import (
	"github.com/dandimuzaki/badminton-server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	TxManager    GormTxManager
	EmailService EmailService
	TokenService TokenService
}

func NewContainer(
	db *gorm.DB,
	config utils.Configuration,
	log *zap.Logger,

) (*Container, error) {
	return &Container{
		TxManager: *NewGormTxManager(db),
		EmailService: NewEmailService(config.SMTP, log),
		TokenService: NewTokenService(config.JWTSecret,config.Issuer),
	}, nil
}