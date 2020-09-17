package impl

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/client"
	"github.com/urfave/negroni"
	adminService "light-up-backend/admin-service/proto"
	authenticationService "light-up-backend/authentication-service/proto"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/proto"
	"light-up-backend/common/utils"
	entityService "light-up-backend/entity-service/proto"
	lightSeekerService "light-up-backend/light-seeker-service/proto"
	lighterService "light-up-backend/lighter-service/proto"
	"net/http"
	"strings"
)

type genericHandler struct {
	LighterServiceClient        lighterService.LighterService
	LightSeekerServiceClient    lightSeekerService.LightSeekerService
	AuthenticationServiceClient authenticationService.AuthenticationService
	AdminServiceClient          adminService.AdminService
	EntityServiceClient         entityService.EntityService
}

func RegisterGenericEndpoints(router *mux.Router, client client.Client, appConfig common.ApplicationConfig) {
	handler := genericHandler{
		LighterServiceClient:        lighterService.CreateNewLighterServiceClient(client),
		LightSeekerServiceClient:    lightSeekerService.CreateNewLightSeekerServiceClient(client),
		AuthenticationServiceClient: authenticationService.CreateNewAuthenticationServiceClient(client),
		AdminServiceClient:          adminService.CreateNewAdminServiceClient(client),
		EntityServiceClient:         entityService.CreateNewLightSeekerServiceClient(client),
	}

	genericBase := mux.NewRouter()
	router.PathPrefix("/api").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(utils.DelegatingHandler),
		negroni.Wrap(genericBase),
	))

	generic := genericBase.PathPrefix("/api").Subrouter()

	// Lighter
	generic.Path("/Lighter/onBoard").HandlerFunc(handler.onBoardLighter)
	generic.Path("/Lighter/getByEmail").HandlerFunc(handler.lighterOnly(handler.getLighterByEmail))
	generic.Path("/Lighter/getById").HandlerFunc(handler.lighterOnly(handler.getLighterById))
	generic.Path("/Lighter/getAll").HandlerFunc(handler.adminOnly(handler.getAllLighters))
	// LightSeeker
	generic.Path("/LightSeeker/onBoard").HandlerFunc(handler.onBoardLightSeeker)
	generic.Path("/LightSeeker/getByEmail").HandlerFunc(handler.lightSeekerOnly(handler.getLightSeekerByEmail))
	generic.Path("/LightSeeker/getById").HandlerFunc(handler.lightSeekerOnly(handler.getLightSeekerById))
	generic.Path("/LightSeeker/getAll").HandlerFunc(handler.adminOnly(handler.getAllLightSeekers))
	// Admin
	generic.Path("/Admin/onBoard").HandlerFunc(handler.onBoardAdmin)
	// Entity
	// Institute
	generic.Path("/Entity/Institute/create").HandlerFunc(handler.adminOnly(handler.createInstitute))
	generic.Path("/Entity/Institute/getById").HandlerFunc(handler.getInstituteById)
	generic.Path("/Entity/Institute/getAll").HandlerFunc(handler.getAllInstitutes)
	// Occupation
	generic.Path("/Entity/Occupation/create").HandlerFunc(handler.adminOnly(handler.createOccupation))
	generic.Path("/Entity/Occupation/getById").HandlerFunc(handler.getOccupationById)
	generic.Path("/Entity/Occupation/getAll").HandlerFunc(handler.getAllOccupations)
	// Educational Qualifications
	generic.Path("/Entity/EducationalQualification/create").HandlerFunc(handler.adminOnly(handler.createEducationalQualifications))
	generic.Path("/Entity/EducationalQualification/getById").HandlerFunc(handler.getEducationalQualificationsById)
	generic.Path("/Entity/EducationalQualification/getAll").HandlerFunc(handler.getAllEducationalQualifications)
}

// Lighter
func (g genericHandler) onBoardLighter(res http.ResponseWriter, req *http.Request) {
	request := &lighterService.Lighter{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LighterServiceClient.CreateLighter(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getLighterByEmail(res http.ResponseWriter, req *http.Request) {
	request := &proto.EmailRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LighterServiceClient.GetLighterByEmail(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getLighterById(res http.ResponseWriter, req *http.Request) {
	request := &proto.IdRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LighterServiceClient.GetLighterById(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getAllLighters(res http.ResponseWriter, req *http.Request) {
	request := &proto.Empty{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LighterServiceClient.GetLighters(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

// Light Seeker
func (g genericHandler) onBoardLightSeeker(res http.ResponseWriter, req *http.Request) {
	request := &lightSeekerService.LightSeeker{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LightSeekerServiceClient.CreateLightSeeker(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getLightSeekerByEmail(res http.ResponseWriter, req *http.Request) {
	request := &proto.EmailRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LightSeekerServiceClient.GetLightSeekerByEmail(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getLightSeekerById(res http.ResponseWriter, req *http.Request) {
	request := &proto.IdRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LightSeekerServiceClient.GetLightSeekerById(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getAllLightSeekers(res http.ResponseWriter, req *http.Request) {
	request := &proto.Empty{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.LightSeekerServiceClient.GetLightSeekers(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

// Admin
func (g genericHandler) onBoardAdmin(res http.ResponseWriter, req *http.Request) {
	request := &adminService.AdminRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.AdminServiceClient.CreateAdmin(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

// Entity
// Institute
func (g genericHandler) createInstitute(res http.ResponseWriter, req *http.Request) {
	request := &entityService.InstituteRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.AddInstitute(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getInstituteById(res http.ResponseWriter, req *http.Request) {
	request := &proto.IdRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.GetInstituteById(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getAllInstitutes(res http.ResponseWriter, req *http.Request) {
	request := &proto.Empty{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.GetAllInstitutes(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

// Occupation
func (g genericHandler) createOccupation(res http.ResponseWriter, req *http.Request) {
	request := &entityService.OccupationRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.AddOccupation(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getOccupationById(res http.ResponseWriter, req *http.Request) {
	request := &proto.IdRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.GetOccupationById(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getAllOccupations(res http.ResponseWriter, req *http.Request) {
	request := &proto.Empty{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.GetAllOccupations(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

// Educational Qualifications
func (g genericHandler) createEducationalQualifications(res http.ResponseWriter, req *http.Request) {
	request := &entityService.EducationQualificationRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.AddEducationQualification(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getEducationalQualificationsById(res http.ResponseWriter, req *http.Request) {
	request := &proto.IdRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.GetEducationQualificationById(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

func (g genericHandler) getAllEducationalQualifications(res http.ResponseWriter, req *http.Request) {
	request := &proto.Empty{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := g.EntityServiceClient.GetAllEducationQualifications(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		_ = json.NewEncoder(res).Encode(response)
	}
}

// Authentication
func (g genericHandler) lighterOnly(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if token, err := extractToken(req); err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		} else {
			if response, err := g.AuthenticationServiceClient.ValidateToken(
				middleware.FromRequest(req),
				&authenticationService.TokenValidationRequest{Token: token, UserType: []proto.UserTypes{
					proto.UserTypes_ADMIN,
					proto.UserTypes_LIGHTER,
				}},
			); err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			} else {
				if !response.Result.Valid {
					res.WriteHeader(http.StatusUnauthorized)
					return
				} else {
					handlerFunc(res, req)
				}
			}
		}
	}
}

func (g genericHandler) lightSeekerOnly(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if token, err := extractToken(req); err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		} else {
			if response, err := g.AuthenticationServiceClient.ValidateToken(
				middleware.FromRequest(req),
				&authenticationService.TokenValidationRequest{Token: token, UserType: []proto.UserTypes{
					proto.UserTypes_ADMIN,
					proto.UserTypes_LIGHT_SEEKER,
				}},
			); err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			} else {
				if !response.Result.Valid {
					res.WriteHeader(http.StatusUnauthorized)
					return
				} else {
					handlerFunc(res, req)
				}
			}
		}
	}
}

func (g genericHandler) adminOnly(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if token, err := extractToken(req); err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		} else {
			if response, err := g.AuthenticationServiceClient.ValidateToken(
				middleware.FromRequest(req),
				&authenticationService.TokenValidationRequest{Token: token, UserType: []proto.UserTypes{
					proto.UserTypes_ADMIN,
				}},
			); err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			} else {
				if !response.Result.Valid {
					res.WriteHeader(http.StatusUnauthorized)
					return
				} else {
					handlerFunc(res, req)
				}
			}
		}
	}
}

func extractToken(req *http.Request) (string, error) {
	authorizationHeader := req.Header.Get("authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			return bearerToken[1], nil
		} else {
			return "", errors.New("bearer token not found")
		}
	} else {
		return "", errors.New("authorization token not found in header")
	}
}
