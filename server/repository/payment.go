package repository

import (
	"context"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/infra"
	"github.com/dandimuzaki/badminton-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *model.Payment) error
	PaymentCallback(ctx context.Context, payload dto.PaymentCallback) error
}

type paymentRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewPaymentRepository(db *gorm.DB, log *zap.Logger) PaymentRepository {
	return &paymentRepository{
		DB:  db,
		Log: log,
	}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment *model.Payment) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.Create(payment).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}

func (r *paymentRepository) FindPaymentByID(ctx context.Context, id uint) (*model.Payment, error) {
	db := infra.GetDB(ctx, r.DB)
	var payment model.Payment
	if err := db.WithContext(ctx).First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) PaymentCallback(ctx context.Context, payload dto.PaymentCallback) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.Model(model.Payment{}).Where("transaction_id = ?", payload.OrderID).
		Updates(map[string]interface{}{
      "status": "expire",
    }).Error; err != nil {
		r.Log.Error(err.Error())
		return err
	}
	return nil
}