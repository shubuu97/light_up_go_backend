package proto

import (
	"github.com/micro/go-micro/client"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
)

func CreateNewAdminServiceClient(client client.Client) AdminService {
	return NewAdminService(common.AdminServiceName, middleware.NewClientWrapper(client))
}
