package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/pkg/errors"
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
	lightSeeker.Id = uuid.New().String()
	lightSeeker.Id = uuid.New().String()
	pwd, err := utils.HashPassword(lightSeeker.Password)
	if err != nil {
		return nil, errors.Wrap(err, "CreateLighter.")
	}
	lightSeeker.Password = pwd
	if err := s.repository.CreateLightSeeker(ctx, lightSeeker); err != nil {
		return nil, errors.Wrap(err, "service")
	} else {
		return lightSeeker, nil
	}
}

func (s LightSeekerService) UpdateLightSeeker(ctx context.Context, lightSeeker *proto.LightSeeker) (*proto.LightSeeker, error) {
	if err := s.repository.UpdateLightSeeker(ctx, lightSeeker); err != nil {
		return nil, errors.Wrap(err, "service")
	} else {
		return lightSeeker, nil
	}
}

func (s LightSeekerService) GetLightSeekerByEmail(ctx context.Context, email string) (*proto.LightSeeker, error) {
	lightSeeker, err := s.repository.GetLightSeekerByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "service")
	}
	return lightSeeker, nil
}

func (s LightSeekerService) GetLightSeekerById(ctx context.Context, id string) (*proto.LightSeeker, error) {
	lightSeeker, err := s.repository.GetLightSeekerById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service")
	}
	return lightSeeker, nil
}

func (s LightSeekerService) GetAllLightSeekers(ctx context.Context) ([]*proto.LightSeeker, error) {
	lightSeekers, err := s.GetAllLightSeekers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service")
	}
	return lightSeekers, nil
}
