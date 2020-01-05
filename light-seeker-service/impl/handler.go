package impl

import (
	"context"
	common "light-up-backend/common/proto"
	"light-up-backend/light-seeker-service/proto"
)

type Handler struct {
	service *LightSeekerService
}

func (h Handler) CreateLightSeeker(ctx context.Context, req *proto.LightSeekerRequest, res *proto.LightSeekerResponse) error {
	if lightSeeker, err := h.service.CreateLightSeeker(ctx, req.LightSeeker); err != nil {
		return err
	} else {
		res.LightSeeker = lightSeeker
		return nil
	}
}

func (h Handler) GetLightSeekerById(ctx context.Context, req *common.IdRequest, res *proto.LightSeekerResponse) error {
	if lightSeeker, err := h.service.GetLightSeekerById(ctx, req.Id); err != nil {
		return err
	} else {
		res.LightSeeker = lightSeeker
		return nil
	}
}

func (h Handler) GetLightSeekerByEmail(ctx context.Context, req *common.EmailRequest, res *proto.LightSeekerResponse) error {
	if lightSeeker, err := h.service.GetLightSeekerByEmail(ctx, req.Email); err != nil {
		return err
	} else {
		res.LightSeeker = lightSeeker
		return nil
	}
}

func (h Handler) GetLightSeekers(ctx context.Context, req *common.Empty, res *proto.LightSeekerResponse) error {
	if lightSeekers, err := h.service.GetAllLightSeekers(ctx); err != nil {
		return err
	} else {
		res.LightSeekers = lightSeekers
		return nil
	}
}

func (h Handler) ValidateLightSeekerUser(ctx context.Context, req *common.IdRequest, res *proto.LightSeekerResponse) error {
	if lightSeeker, err := h.service.ValidateLightSeekerUser(ctx, req.Id); err != nil {
		return err
	} else {
		res.LightSeeker = lightSeeker
		return nil
	}
}

func (h Handler) InValidateLightSeekerUser(ctx context.Context, req *common.IdRequest, res *proto.LightSeekerResponse) error {
	if lightSeeker, err := h.service.InValidateLightSeekerUser(ctx, req.Id); err != nil {
		return err
	} else {
		res.LightSeeker = lightSeeker
		return nil
	}
}

func NewHandler(service *LightSeekerService) Handler {
	return Handler{
		service: service,
	}
}


