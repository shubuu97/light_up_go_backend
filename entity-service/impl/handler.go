package impl

import (
	"context"
	common "light-up-backend/common/proto"
	"light-up-backend/entity-service/proto"
)

type Handler struct {
	service *EntityService
}

func (h Handler) AddEducationQualification(ctx context.Context, req *proto.EducationQualificationRequest, res *proto.EducationQualificationResponse) error {
	if educationQualification, err := h.service.CreateEducationQualification(ctx, req.EducationQualification); err != nil {
		return err
	} else {
		res.EducationQualification = educationQualification
		return nil
	}
}

func (h Handler) GetEducationQualificationById(ctx context.Context, req *common.IdRequest, res *proto.EducationQualificationResponse) error {
	if educationQualification, err := h.service.GetEducationQualificationById(ctx, req.Id); err != nil {
		return err
	} else {
		res.EducationQualification = educationQualification
		return nil
	}
}

func (h Handler) GetAllEducationQualifications(ctx context.Context, req *common.Empty, res *proto.EducationQualificationResponse) error {
	if educationQualifications, err := h.service.GetAllEducationQualifications(ctx); err != nil {
		return err
	} else {
		res.EducationQualifications = educationQualifications
		return nil
	}
}

func (h Handler) AddOccupation(ctx context.Context, req *proto.OccupationRequest, res *proto.OccupationResponse) error {
	if occupation, err := h.service.CreateOccupation(ctx, req.Occupation); err != nil {
		return err
	} else {
		res.Occupation = occupation
		return nil
	}
}

func (h Handler) GetOccupationById(ctx context.Context, req *common.IdRequest, res *proto.OccupationResponse) error {
	if occupation, err := h.service.GetOccupationById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Occupation = occupation
		return nil
	}
}

func (h Handler) GetAllOccupations(ctx context.Context, req *common.Empty, res *proto.OccupationResponse) error {
	if occupations, err := h.service.GetAllOccupations(ctx); err != nil {
		return err
	} else {
		res.Occupations = occupations
		return nil
	}
}

func (h Handler) AddInstitute(ctx context.Context, req *proto.InstituteRequest, res *proto.InstituteResponse) error {
	if institute, err := h.service.CreateInstitute(ctx, req.Institute); err != nil {
		return err
	} else {
		res.Institute = institute
		return nil
	}
}

func (h Handler) GetInstituteById(ctx context.Context, req *common.IdRequest, res *proto.InstituteResponse) error {
	if institute, err := h.service.GetInstituteById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Institute = institute
		return nil
	}
}

func (h Handler) GetAllInstitutes(ctx context.Context, req *common.Empty, res *proto.InstituteResponse) error {
	if institutes, err := h.service.GetAllInstitutes(ctx); err != nil {
		return err
	} else {
		res.Institutes = institutes
		return nil
	}
}

func NewHandler(service *EntityService) Handler {
	return Handler{
		service: service,
	}
}
