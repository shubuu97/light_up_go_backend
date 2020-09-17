package impl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"light-up-backend/admin-service/proto"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
)

const (
	userCollection = "users"
)

type AdminRepository struct {
	client *mongo.Client
	config common.ServiceDbConfigurations
}

func NewAdminRepository(ctx context.Context, config common.ServiceConfigurations) AdminRepository {
	logger := middleware.GetLogger(ctx)
	logger.Infoln("Connecting to DB.")
	client, err := config.DbConfigs.ConnectToMongo()
	if err != nil {
		logger.WithField("Error", err.Error()).Panicln("Could not connect to the database.")
	}
	logger.Infoln("Connected to Mongo.")
	return AdminRepository{ client, config.DbConfigs}
}

func (a AdminRepository) Close() {
	defer a.client.Disconnect(context.Background())
}

func (a AdminRepository) userCollection() *mongo.Collection {
	return a.client.Database(a.config.DbName).Collection(userCollection)
}

func (a AdminRepository) CreateAdmin(ctx context.Context, admin *proto.Admin) error {
	_, err := a.userCollection().InsertOne(ctx, admin)
	if err != nil {
		return errors.Wrap(err, "Could not insert admin.")
	} else {
		return nil
	}
}

func (a AdminRepository) UpdateAdmin(ctx context.Context, admin *proto.Admin) error {
	if _, err := a.userCollection().UpdateOne(ctx, bson.M{"id": admin.Id}, bson.D{{"$set", admin}}); err != nil {
		return errors.Wrap(err, "Could not update the admin.")
	} else {
		return nil
	}
}

func (a AdminRepository) GetAdminById(ctx context.Context, id string) (*proto.Admin, error) {
	result := proto.Admin{}
	if err := a.userCollection().FindOne(ctx, bson.M{"id": id}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the admin by id: %s", id))
	} else {
		return &result, nil
	}
}

func (a AdminRepository) GetAdminByEmail(ctx context.Context, email string) (*proto.Admin, error) {
	result := proto.Admin{}
	if err := a.userCollection().FindOne(ctx, bson.M{"email": email}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the admin by email: %s", email))
	} else {
		return &result, nil
	}
}

func (a AdminRepository) GetAllAdmins(ctx context.Context) ([]*proto.Admin, error) {
	var result []*proto.Admin
	cur, err := a.userCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "Could not get the admins.")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		client := &proto.Admin{}
		err = cur.Decode(client)
		if err != nil {
			return nil, errors.Wrap(err, "Could not decode the admin.")
		}
		result = append(result, client)
	}
	return result, nil
}