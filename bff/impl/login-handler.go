package impl

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/client"
	"github.com/urfave/negroni"
	"light-up-backend/authentication-service/proto"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
	"net/http"
)

type loginHandler struct {
	authenticationServiceClient proto.AuthenticationService
}

func RegisterLoginEndpoints(r *mux.Router, client client.Client) {
	handler := loginHandler{
		authenticationServiceClient: proto.CreateNewAuthenticationServiceClient(client),
	}

	loginBase := mux.NewRouter()
	r.PathPrefix("/api/Auth").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(utils.DelegatingHandler),
		negroni.Wrap(loginBase),
	))
	login := loginBase.PathPrefix("/api/Auth/Login").Subrouter()
	login.Path("/Lighter").HandlerFunc(handler.LighterLogin)
	login.Path("/LightSeeker").HandlerFunc(handler.LightSeekerLogin)
}

func (l loginHandler) LighterLogin(res http.ResponseWriter, req *http.Request) {
	request := &proto.LoginRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := l.authenticationServiceClient.LoginLighter(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(res).Encode(response)
}

func (l loginHandler) LightSeekerLogin(res http.ResponseWriter, req *http.Request) {
	request := &proto.LoginRequest{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := l.authenticationServiceClient.LoginLightSeeker(middleware.FromRequest(req), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(res).Encode(response)
}
