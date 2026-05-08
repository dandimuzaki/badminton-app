package usecase

import (
	"context"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/model"
	"github.com/dandimuzaki/badminton-server/repository"
	"go.uber.org/zap"
)

type CourtUsecase interface {
	GetAvailableCourts(ctx context.Context, req dto.AvailableCourtRequest) ([]model.Court, error)
}

type courtUsecase struct {
	Repo *repository.Repository
	Log *zap.Logger
}

func NewCourtUsecase(repo *repository.Repository, log *zap.Logger) CourtUsecase {
	return &courtUsecase{
		Repo: repo,
		Log: log,
	}
}

func (u *courtUsecase) GetAvailableCourts(ctx context.Context, req dto.AvailableCourtRequest) ([]model.Court, error) {
	return u.Repo.CourtRepo.GetAvailableCourts(ctx, req)
}