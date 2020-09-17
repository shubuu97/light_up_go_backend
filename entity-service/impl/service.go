package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/pkg/errors"
	"light-up-backend/common/middleware"
	"light-up-backend/common/proto"
)

type EntityService struct {
	repository EntityRepository
}

func NewEntityService(repository EntityRepository, client client.Client) *EntityService {
	return &EntityService{
		repository: repository,
	}
}

func (s EntityService) CreateInstitute(ctx context.Context, institute *proto.Institute) (*proto.Institute, error) {
	logger := middleware.GetLogger(ctx)
	institute.Id = uuid.New().String()
	if err := s.repository.CreateInstitute(ctx, institute); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not create institute.")
		return nil, errors.Wrap(err, "service")
	} else {
		return institute, nil
	}
}

func (s EntityService) UpdateInstitute(ctx context.Context, institute *proto.Institute) (*proto.Institute, error) {
	logger := middleware.GetLogger(ctx)
	if err := s.repository.UpdateInstitute(ctx, institute); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not update institute.")
		return nil, errors.Wrap(err, "service")
	} else {
		return institute, nil
	}
}

func (s EntityService) GetInstituteById(ctx context.Context, id string) (*proto.Institute, error) {
	logger := middleware.GetLogger(ctx)
	if institute, err := s.repository.GetInstituteById(ctx, id); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get institute by Id.")
		return nil, errors.Wrap(err, "service")
	} else {
		return institute, nil
	}
}

func (s EntityService) GetAllInstitutes(ctx context.Context) ([]*proto.Institute, error) {
	logger := middleware.GetLogger(ctx)
	if institutes, err := s.repository.GetAllInstitutes(ctx); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not all institutes.")
		return nil, errors.Wrap(err, "service")
	} else {
		return institutes, nil
	}
}

func (s EntityService) CreateOccupation(ctx context.Context, occupation *proto.Occupation) (*proto.Occupation, error) {
	logger := middleware.GetLogger(ctx)
	occupation.Id = uuid.New().String()
	if err := s.repository.CreateOccupation(ctx, occupation); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not create occupation.")
		return nil, errors.Wrap(err, "service")
	} else {
		return occupation, nil
	}
}

func (s EntityService) UpdateOccupation(ctx context.Context, occupation *proto.Occupation) (*proto.Occupation, error) {
	logger := middleware.GetLogger(ctx)
	if err := s.repository.UpdateOccupation(ctx, occupation); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not update occupation.")
		return nil, errors.Wrap(err, "service")
	} else {
		return occupation, nil
	}
}

func (s EntityService) GetOccupationById(ctx context.Context, id string) (*proto.Occupation, error) {
	logger := middleware.GetLogger(ctx)
	if occupation, err := s.repository.GetOccupationById(ctx, id); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get occupation by Id.")
		return nil, errors.Wrap(err, "service")
	} else {
		return occupation, nil
	}
}

func (s EntityService) GetAllOccupations(ctx context.Context) ([]*proto.Occupation, error) {
	logger := middleware.GetLogger(ctx)
	if occupations, err := s.repository.GetAllOccupations(ctx); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not all occupations.")
		return nil, errors.Wrap(err, "service")
	} else {
		return occupations, nil
	}
}

func (s EntityService) CreateEducationQualification(ctx context.Context, educationQualification *proto.EducationQualification) (*proto.EducationQualification, error) {
	logger := middleware.GetLogger(ctx)
	educationQualification.Id = uuid.New().String()
	if err := s.repository.CreateEducationQualification(ctx, educationQualification); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not create education qualification.")
		return nil, errors.Wrap(err, "service")
	} else {
		return educationQualification, nil
	}
}

func (s EntityService) UpdateEducationQualification(ctx context.Context, educationQualification *proto.EducationQualification) (*proto.EducationQualification, error) {
	logger := middleware.GetLogger(ctx)
	if err := s.repository.UpdateEducationQualification(ctx, educationQualification); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not update education qualification.")
		return nil, errors.Wrap(err, "service")
	} else {
		return educationQualification, nil
	}
}

func (s EntityService) GetEducationQualificationById(ctx context.Context, id string) (*proto.EducationQualification, error) {
	logger := middleware.GetLogger(ctx)
	if educationQualification, err := s.repository.GetEducationQualificationById(ctx, id); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get education qualification by Id.")
		return nil, errors.Wrap(err, "service")
	} else {
		return educationQualification, nil
	}
}

func (s EntityService) GetAllEducationQualifications(ctx context.Context) ([]*proto.EducationQualification, error) {
	logger := middleware.GetLogger(ctx)
	if educationQualifications, err := s.repository.GetAllEducationQualifications(ctx); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not all education qualifications.")
		return nil, errors.Wrap(err, "service")
	} else {
		return educationQualifications, nil
	}
}