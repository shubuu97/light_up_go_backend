package impl

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/pkg/errors"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
	"light-up-backend/light-seeker-service/proto"
)

type LightSeekerService struct {
	repository LightSeekerRepository
}

func NewLightSeekerService(repository LightSeekerRepository, client client.Client) *LightSeekerService {
	return &LightSeekerService{
		repository: repository,
	}
}

func (s LightSeekerService) CreateLightSeeker(ctx context.Context, lightSeeker *proto.LightSeeker) (*proto.LightSeeker, error) {
	logger := middleware.GetLogger(ctx)
	lightSeeker.Id = uuid.New().String()
	pwd, err := utils.HashPassword(lightSeeker.User.Password)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not hash password.")
		return nil, errors.Wrap(err, "CreateLighter.")
	}
	lightSeeker.User.Password = pwd
	lightSeeker.User.IsValid = true
	lightSeeker.CreatedOn = ptypes.TimestampNow()
	lightSeeker.ModifiedOn = ptypes.TimestampNow()
	if err := s.repository.CreateLightSeeker(ctx, lightSeeker); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not create light seeker.")
		return nil, errors.Wrap(err, "service")
	} else {
		return lightSeeker, nil
	}
}

func (s LightSeekerService) UpdateLightSeeker(ctx context.Context, lightSeeker *proto.LightSeeker) (*proto.LightSeeker, error) {
	logger := middleware.GetLogger(ctx)
	lightSeeker.ModifiedOn = ptypes.TimestampNow()
	if err := s.repository.UpdateLightSeeker(ctx, lightSeeker); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not update light seeker.")
		return nil, errors.Wrap(err, "service")
	} else {
		return lightSeeker, nil
	}
}

func (s LightSeekerService) GetLightSeekerByEmail(ctx context.Context, email string) (*proto.LightSeeker, error) {
	logger := middleware.GetLogger(ctx)
	lightSeeker, err := s.repository.GetLightSeekerByEmail(ctx, email)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could get light seeker by email.")
		return nil, errors.Wrap(err, "service")
	}
	return lightSeeker, nil
}

func (s LightSeekerService) GetLightSeekerById(ctx context.Context, id string) (*proto.LightSeeker, error) {
	logger := middleware.GetLogger(ctx)
	lightSeeker, err := s.repository.GetLightSeekerById(ctx, id)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get light seeker by id.")
		return nil, errors.Wrap(err, "service")
	}
	return lightSeeker, nil
}

func (s LightSeekerService) GetAllLightSeekers(ctx context.Context) ([]*proto.LightSeeker, error) {
	logger := middleware.GetLogger(ctx)
	lightSeekers, err := s.repository.GetAllLightSeekers(ctx)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get all light seekers.")
		return nil, errors.Wrap(err, "service")
	}
	return lightSeekers, nil
}

func (s LightSeekerService) ValidateLightSeekerUser(ctx context.Context, id string) (*proto.LightSeeker, error) {
	logger := middleware.GetLogger(ctx)
	lightSeeker, err := s.repository.GetLightSeekerById(ctx, id)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get light seeker by id.")
		return nil, errors.Wrap(err, "service")
	}
	lightSeeker.User.IsValid = true
	lightSeeker.ModifiedOn = ptypes.TimestampNow()
	err = s.repository.UpdateLightSeeker(ctx, lightSeeker)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Failed to mark it as a valid user.")
		return nil, errors.Wrap(err, "service")
	}
	return lightSeeker, nil
}

func (s LightSeekerService) InValidateLightSeekerUser(ctx context.Context, id string) (*proto.LightSeeker, error) {
	logger := middleware.GetLogger(ctx)
	lightSeeker, err := s.repository.GetLightSeekerById(ctx, id)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get light seeker by id.")
		return nil, errors.Wrap(err, "service")
	}
	lightSeeker.User.IsValid = false
	lightSeeker.ModifiedOn = ptypes.TimestampNow()
	err = s.repository.UpdateLightSeeker(ctx, lightSeeker)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Failed to mark it as a valid user.")
		return nil, errors.Wrap(err, "service")
	}
	return lightSeeker, nil
}