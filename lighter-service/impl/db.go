package impl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/lighter-service/proto"
	"strings"
)

const (
	userCollection = "users"
)

type LighterRepository struct {
	client *mongo.Client
	config common.ServiceDbConfigurations
}

func NewLighterRepository(ctx context.Context, config common.ServiceConfigurations) LighterRepository {
	logger := middleware.GetLogger(ctx)
	logger.Infoln("Connecting to DB.")
	client, err := config.DbConfigs.ConnectToMongo()
	if err != nil {
		logger.WithField("Error", err.Error()).Panicln("Could not connect to the database.")
	}
	logger.Infoln("Connected to Mongo.")
	return LighterRepository{client, config.DbConfigs}
}

func (l LighterRepository) Close() {
	defer l.client.Disconnect(context.Background())
}

func (l LighterRepository) userCollection() *mongo.Collection {
	return l.client.Database(l.config.DbName).Collection(userCollection)
}

func (l LighterRepository) CreateLighter(ctx context.Context, lighter *proto.Lighter) error {
	lighter.User.Email = strings.ToLower(lighter.User.Email)
	_, err := l.userCollection().InsertOne(ctx, lighter)
	if err != nil {
		return errors.Wrap(err, "Could not insert lighter.")
	} else {
		return nil
	}
}

func (l LighterRepository) UpdateLighter(ctx context.Context, lighter *proto.Lighter) error {
	lighter.User.Email = strings.ToLower(lighter.User.Email)
	if _, err := l.userCollection().UpdateOne(ctx, bson.M{"id": lighter.Id}, bson.D{{"$set", lighter}}); err != nil {
		return errors.Wrap(err, "Could not update the lighter.")
	} else {
		return nil
	}
}

func (l LighterRepository) GetLighterByEmail(ctx context.Context, email string) (*proto.Lighter, error) {
	result := proto.Lighter{}
	if err := l.userCollection().FindOne(ctx, bson.M{"email": email}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the lighter by email: %s", email))
	} else {
		return &result, nil
	}
}

func (l LighterRepository) GetLighterById(ctx context.Context, id string) (*proto.Lighter, error) {
	result := proto.Lighter{}
	if err := l.userCollection().FindOne(ctx, bson.M{"id": id}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the lighter by id: %s", id))
	} else {
		return &result, nil
	}
}

func (l LighterRepository) GetAllLighters(ctx context.Context) ([]*proto.Lighter, error) {
	var result []*proto.Lighter
	cur, err := l.userCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "Could not get the lighters.")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		client := &proto.Lighter{}
		err = cur.Decode(client)
		if err != nil {
			return nil, errors.Wrap(err, "Could not decode the lighter.")
		}
		result = append(result, client)
	}
	return result, nil
}
