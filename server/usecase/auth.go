package usecase

import (
	"context"

	"errors"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/infra"
	"github.com/dandimuzaki/badminton-server/model"
	"github.com/dandimuzaki/badminton-server/repository"
	"github.com/dandimuzaki/badminton-server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, *string, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, *string, error)
}

type authUsecase struct {
	Repo     *repository.Repository
	TokenService infra.TokenService
	Log          *zap.Logger
}

func NewAuthUsecase(
	repo *repository.Repository,
	tokenService infra.TokenService,
	log *zap.Logger,
) AuthUsecase {
	return &authUsecase{
		Repo:     repo,
		TokenService: tokenService,
		Log:          log,
	}
}

func (u *authUsecase) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, *string, error) {
	// Check if email exists
	_, err := u.Repo.UserRepo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, nil, errors.New("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	}

	hashedPassword := utils.HashPassword(req.Password)

	var user *model.User
	user = &model.User{
		Name: req.Name,
		Email:        req.Email,
		Password: hashedPassword,
		Role:         model.RoleCustomer,
	}

	user, err = u.Repo.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	token, err := u.TokenService.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, nil, err
	}

	authResponse := dto.AuthResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return &authResponse, &token, nil
}

func (u *authUsecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, *string, error) {
	user, err := u.Repo.UserRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("invalid email or password")
		}
		return nil, nil, err
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, nil, errors.New("invalid email or password")
	}

	token, err := u.TokenService.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, nil, err
	}

	authResponse := dto.AuthResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return &authResponse, &token, nil
}

