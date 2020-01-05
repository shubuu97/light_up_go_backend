package impl

import (
	"context"
	"light-up-backend/admin-service/proto"
)

type Handler struct {
	service *AdminService
}

func (h Handler) CreateAdmin(ctx context.Context, req *proto.AdminRequest, res *proto.AdminResponse) error {
	if admin, err := h.service.CreateAdmin(ctx, req.Admin); err != nil {
		return err
	} else {
		res.Admin = admin
		return nil
	}
}

func NewHandler(service *AdminService) Handler {
	return Handler{
		service: service,
	}
}
