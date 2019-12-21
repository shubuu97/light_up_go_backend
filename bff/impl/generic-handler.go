package impl

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/client"
	"github.com/urfave/negroni"
	authenticationService "light-up-backend/authentication-service/proto"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/proto"
	"light-up-backend/common/utils"
	lightSeekerService "light-up-backend/light-seeker-service/proto"
	lighterService "light-up-backend/lighter-service/proto"
	"net/http"
	"strings"
)

type genericHandler struct {
	LighterServiceClient        lighterService.LighterService
	LightSeekerServiceClient    lightSeekerService.LightSeekerService
	AuthenticationServiceClient authenticationService.AuthenticationService
}

func RegisterGenericEndpoints(router *mux.Router, client client.Client, appConfig common.ApplicationConfig) {
	handler := genericHandler{
		LighterServiceClient:        lighterService.CreateNewLighterServiceClient(client),
		LightSeekerServiceClient:    lightSeekerService.CreateNewLightSeekerServiceClient(client),
		AuthenticationServiceClient: authenticationService.CreateNewAuthenticationServiceClient(client),
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
	generic.Path("/Lighter/OnBoard").HandlerFunc(handler.onBoardLighter)
	generic.Path("/Lighter/getByEmail").HandlerFunc(handler.lighterOnly(handler.getLighterByEmail))
	generic.Path("/Lighter/getById").HandlerFunc(handler.lighterOnly(handler.getLighterById))
	generic.Path("/Lighter/getAll").HandlerFunc(handler.adminOnly(handler.getAllLighters))
	// LightSeeker
	generic.Path("/LightSeeker/OnBoard").HandlerFunc(handler.onBoardLightSeeker)
	generic.Path("/LightSeeker/getByEmail").HandlerFunc(handler.lightSeekerOnly(handler.getLightSeekerByEmail))
	generic.Path("/LightSeeker/getById").HandlerFunc(handler.lightSeekerOnly(handler.getLightSeekerById))
	generic.Path("/LightSeeker/getAll").HandlerFunc(handler.adminOnly(handler.getAllLightSeekers))

}

func (g genericHandler) onBoardLighter(res http.ResponseWriter, req *http.Request) {
	request := &lighterService.CreateLighterRequest{}
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

func (g genericHandler) onBoardLightSeeker(res http.ResponseWriter, req *http.Request) {
	request := &lightSeekerService.CreateLightSeekerRequest{}
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
