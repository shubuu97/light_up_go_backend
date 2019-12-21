package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/pkg/errors"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
	"light-up-backend/lighter-service/proto"
)

type LighterService struct {
	repository LighterRepository
}

func NewLighterService(repository LighterRepository, client client.Client) *LighterService {
	return &LighterService{
		repository: repository,
	}
}

func (s LighterService) CreateLighter(ctx context.Context, lighter *proto.Lighter) (*proto.Lighter, error) {
	logger := middleware.GetLogger(ctx)
	lighter.Id = uuid.New().String()
	pwd, err := utils.HashPassword(lighter.Password)
	if err != nil {
		return nil, errors.Wrap(err, "CreateLighter.")
	}
	lighter.Password = pwd
	if err := s.repository.CreateLighter(ctx, lighter); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not create lighter.")
		return nil, errors.Wrap(err, "service")
	} else {
		return lighter, nil
	}
}

func (s LighterService) UpdateLighter(ctx context.Context, lighter *proto.Lighter) (*proto.Lighter, error) {
	logger := middleware.GetLogger(ctx)
	if err := s.repository.UpdateLighter(ctx, lighter); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not update lighter.")
		return nil, errors.Wrap(err, "service")
	} else {
		return lighter, nil
	}
}

func (s LighterService) GetLighterByEmail(ctx context.Context, email string) (*proto.Lighter, error) {
	logger := middleware.GetLogger(ctx)
	lighter, err := s.repository.GetLighterByEmail(ctx, email)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get lighter by email.")
		return nil, errors.Wrap(err, "service")
	}
	return lighter, nil
}

func (s LighterService) GetLighterById(ctx context.Context, id string) (*proto.Lighter, error) {
	logger := middleware.GetLogger(ctx)
	lighter, err := s.repository.GetLighterById(ctx, id)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get lighter by Id.")
		return nil, errors.Wrap(err, "service")
	}
	return lighter, nil
}

func (s LighterService) GetAllLighters(ctx context.Context) ([]*proto.Lighter, error) {
	logger := middleware.GetLogger(ctx)
	lighters, err := s.repository.GetAllLighters(ctx)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get all lighters.")
		return nil, errors.Wrap(err, "service")
	}
	return lighters, nil
}
