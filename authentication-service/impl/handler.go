package impl

import (
	"context"
	"github.com/pkg/errors"
	"light-up-backend/authentication-service/proto"
	common "light-up-backend/common/proto"
)

type Handler struct {
	Service *Service
}

func (h Handler) LoginLightSeeker(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	token, err := h.Service.LoginLightSeeker(ctx, req.Password, req.Email)
	if err != nil {
		return errors.Wrap(err, "handler")
	}
	res.Token = token
	res.Message = "LightSeeker logged in successfully."
	return nil
}

func (h Handler) LoginLighter(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	token, err := h.Service.LoginLighter(ctx, req.Password, req.Email)
	if err != nil {
		return errors.Wrap(err, "handler")
	}
	res.Token = token
	res.Message = "Lighter logged in successfully."
	return nil

}

func (h Handler) ValidateToken(ctx context.Context, req *proto.TokenValidationRequest, res *proto.TokenResponse) error {
	isValid := false

	for _, userType := range req.UserType {
		if userType == common.UserTypes_LIGHT_SEEKER {
			_, err := h.Service.ValidateLightSeeker(ctx, req.Token)
			if err == nil {
				isValid = true
				break
			}
		} else if userType == common.UserTypes_LIGHTER {
			_, err := h.Service.ValidateLighter(ctx, req.Token)
			if err == nil {
				isValid = true
				break
			}
		} else if userType == common.UserTypes_ADMIN {
			_, err := h.Service.ValidateAdmin(ctx, req.Token)
			if err == nil {
				isValid = true
				break
			}
		}
	}

	if isValid == false {
		return errors.New("Token is not valid")
	} else {
		res.Result = &proto.Token{Token: req.Token, Valid: true}
		return nil
	}
}
