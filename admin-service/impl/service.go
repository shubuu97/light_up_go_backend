package impl

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/pkg/errors"
	"light-up-backend/admin-service/proto"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
)

type AdminService struct {
	repository AdminRepository
}

func NewAdminService(repository AdminRepository, client client.Client) *AdminService {
	return &AdminService{
		repository:repository,
	}
}

func (s AdminService) CreateAdmin(ctx context.Context, admin *proto.Admin) (*proto.Admin, error) {
	logger := middleware.GetLogger(ctx)
	admin.Id = uuid.New().String()
	pwd, err := utils.HashPassword(admin.User.Password)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not hash password.")
		return nil, errors.Wrap(err, "CreateAdmin.")
	}
	admin.User.Password = pwd
	admin.User.IsValid = true
	admin.CreatedOn = ptypes.TimestampNow()
	admin.ModifiedOn = ptypes.TimestampNow()
	if err := s.repository.CreateAdmin(ctx, admin); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not create admin.")
		return nil, errors.Wrap(err, "service")
	} else {
		return admin, nil
	}
}

func (s AdminService) UpdateAdmin(ctx context.Context, admin *proto.Admin) (*proto.Admin, error) {
	logger := middleware.GetLogger(ctx)
	pwd, err := utils.HashPassword(admin.User.Password)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not hash password.")
		return nil, errors.Wrap(err, "UpdateAdmin.")
	}
	admin.User.Password = pwd
	admin.User.IsValid = true
	admin.ModifiedOn = ptypes.TimestampNow()
	if err := s.repository.UpdateAdmin(ctx, admin); err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not update admin.")
		return nil, errors.Wrap(err, "service")
	} else {
		return admin, nil
	}
}

func (s AdminService) GetAdminById(ctx context.Context, id string) (*proto.Admin, error) {
	logger := middleware.GetLogger(ctx)
	admin, err := s.repository.GetAdminById(ctx, id)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get admin by id.")
		return nil, errors.Wrap(err, "service")
	}
	if admin.User.IsValid != true {
		logger.Warningln("Could not get invalid admin by Id.")
		return nil, errors.New("Invalid user.")
	}
	return admin, nil
}

func (s AdminService) GetAdminByEmail(ctx context.Context, email string) (*proto.Admin, error) {
	logger := middleware.GetLogger(ctx)
	admin, err := s.repository.GetAdminByEmail(ctx, email)
	if err != nil {
		logger.WithField("Error", err.Error()).Warningln("Could not get admin by email.")
		return nil, errors.Wrap(err, "service")
	}
	if admin.User.IsValid != true {
		logger.Warningln("Could not get invalid admin by email.")
		return nil, errors.New("Invalid user.")
	}
	return admin, nil
}