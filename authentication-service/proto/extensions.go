package proto

import (
	"github.com/micro/go-micro/client"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
)

func CreateNewAuthenticationServiceClient(client client.Client) AuthenticationService {
	return NewAuthenticationService(common.AuthenticationServiceName, middleware.NewClientWrapper(client))
}
