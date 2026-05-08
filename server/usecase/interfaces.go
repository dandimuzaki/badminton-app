// usecase/interfaces.go
package usecase

import (
	"context"

	"github.com/dandimuzaki/badminton-server/infra"
)

type TxManager interface {
	WithinTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type TokenService interface {
	GenerateToken(userID uint, role string) (string, error)
	ValidateToken(tokenString string) (*infra.Claims, error)
}