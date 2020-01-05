package impl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/light-seeker-service/proto"
	"strings"
)

const (
	userCollection = "users"
)

type LightSeekerRepository struct {
	client *mongo.Client
	config common.ServiceDbConfigurations
}

func NewLightSeekerRepository(ctx context.Context, config common.ServiceConfigurations) LightSeekerRepository {
	logger := middleware.GetLogger(ctx)
	logger.Infoln("Connecting to DB.")
	client, err := config.DbConfigs.ConnectToMongo()
	if err != nil {
		logger.WithField("Error", err.Error()).Panicln("Could not connect to the database.")
	}
	logger.Infoln("Connected to Mongo.")
	return LightSeekerRepository{client, config.DbConfigs}
}

func (l LightSeekerRepository) Close() {
	defer l.client.Disconnect(context.Background())
}

func (l LightSeekerRepository) userCollection() *mongo.Collection {
	return l.client.Database(l.config.DbName).Collection(userCollection)
}

func (l LightSeekerRepository) CreateLightSeeker(ctx context.Context, lightSeeker *proto.LightSeeker) error {
	lightSeeker.User.Email = strings.ToLower(lightSeeker.User.Email)
	_, err := l.userCollection().InsertOne(ctx, lightSeeker)
	if err != nil {
		return errors.Wrap(err, "Could not insert light seeker.")
	} else {
		return nil
	}
}

func (l LightSeekerRepository) UpdateLightSeeker(ctx context.Context, lightSeeker *proto.LightSeeker) error {
	lightSeeker.User.Email = strings.ToLower(lightSeeker.User.Email)
	if _, err := l.userCollection().UpdateOne(ctx, bson.M{"id": lightSeeker.Id}, bson.D{{"$set", lightSeeker}}); err != nil {
		return errors.Wrap(err, "Could not update the light seeker.")
	} else {
		return nil
	}
}

func (l LightSeekerRepository) GetLightSeekerByEmail(ctx context.Context, email string) (*proto.LightSeeker, error) {
	result := proto.LightSeeker{}
	if err := l.userCollection().FindOne(ctx, bson.M{"email": email}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the light seeker by email: %s", email))
	} else {
		return &result, nil
	}
}

func (l LightSeekerRepository) GetLightSeekerById(ctx context.Context, id string) (*proto.LightSeeker, error) {
	result := proto.LightSeeker{}
	if err := l.userCollection().FindOne(ctx, bson.M{"id": id}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the light seeker by id:s %s", id))
	} else {
		return &result, nil
	}
}

func (l LightSeekerRepository) GetAllLightSeekers(ctx context.Context) ([]*proto.LightSeeker, error) {
	var result []*proto.LightSeeker
	cur, err := l.userCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "Could not get the light seeker.")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		client := &proto.LightSeeker{}
		err = cur.Decode(client)
		if err != nil {
			return nil, errors.Wrap(err, "Could not decode the light seeker.")
		}
		result = append(result, client)
	}
	return result, nil
}

