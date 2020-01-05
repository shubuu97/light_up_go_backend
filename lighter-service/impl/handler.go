package impl

import (
	"context"
	common "light-up-backend/common/proto"
	"light-up-backend/lighter-service/proto"
)

type Handler struct {
	service *LighterService
}

func NewHandler(service *LighterService) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) CreateLighter(ctx context.Context, req *proto.LighterRequest, res *proto.LighterResponse) error {
	if lighter, err := h.service.CreateLighter(ctx, req.Lighter); err != nil {
		return err
	} else {
		res.Lighter = lighter
		return nil
	}
}

func (h Handler) GetLighterById(ctx context.Context, req *common.IdRequest, res *proto.LighterResponse) error {
	if lighter, err := h.service.GetLighterById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Lighter = lighter
		return nil
	}
}

func (h Handler) GetLighterByEmail(ctx context.Context, req *common.EmailRequest, res *proto.LighterResponse) error {
	if lighter, err := h.service.GetLighterByEmail(ctx, req.Email); err != nil {
		return err
	} else {
		res.Lighter = lighter
		return nil
	}
}

func (h Handler) GetLighters(ctx context.Context, req *common.Empty, res *proto.LighterResponse) error {
	if lighters, err := h.service.GetAllLighters(ctx); err != nil {
		return err
	} else {
		res.Lighters = lighters
		return nil
	}
}

func (h Handler) VerifyLighterUser(ctx context.Context, req *common.IdRequest, res *proto.LighterResponse) error {
	if lighter, err := h.service.VerifyLighterUser(ctx, req.Id); err != nil {
		return err
	} else {
		res.Lighter = lighter
		return nil
	}
}

func (h Handler) ValidateLighterUser(ctx context.Context, req *common.IdRequest, res *proto.LighterResponse) error {
	if lighter, err := h.service.ValidateLighterUser(ctx, req.Id); err != nil {
		return err
	} else {
		res.Lighter = lighter
		return nil
	}
}

func (h Handler) InValidateLighterUser(ctx context.Context, req *common.IdRequest, res *proto.LighterResponse) error {
	if lighter, err := h.service.InValidateLighterUser(ctx, req.Id); err != nil {
		return err
	} else {
		res.Lighter = lighter
		return nil
	}
}