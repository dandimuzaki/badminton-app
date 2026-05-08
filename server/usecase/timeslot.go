package usecase

import (
	"context"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/model"
	"github.com/dandimuzaki/badminton-server/repository"
	"go.uber.org/zap"
)

type TimeslotUsecase interface {
	GetAvailableTimeslots(ctx context.Context, req dto.AvailableTimeslotRequest) ([]model.Timeslot, error)
}

type timeslotUsecase struct {
	Repo *repository.Repository
	Log *zap.Logger
}

func NewTimeslotUsecase(repo *repository.Repository, log *zap.Logger) TimeslotUsecase {
	return &timeslotUsecase{
		Repo: repo,
		Log: log,
	}
}

func (u *timeslotUsecase) GetAvailableTimeslots(ctx context.Context, req dto.AvailableTimeslotRequest) ([]model.Timeslot, error) {
	return u.Repo.TimeslotRepo.GetAvailableTimeslots(ctx, req)
}