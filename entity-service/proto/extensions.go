package proto

import (
	"github.com/micro/go-micro/client"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
)

func CreateNewLightSeekerServiceClient(client client.Client) EntityService {
	return NewEntityService(common.EntityServiceName, middleware.NewClientWrapper(client))
}
