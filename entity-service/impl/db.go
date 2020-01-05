package impl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/proto"
)

const (
	institutesCollection = "institutes"
	occupationsCollection = "occupations"
	educationQualificationsCollection = "education-qualifications"
)

type EntityRepository struct {
	client *mongo.Client
	config common.ServiceDbConfigurations
}

func NewEntityRepository(ctx context.Context, config common.ServiceConfigurations) EntityRepository {
	logger := middleware.GetLogger(ctx)
	logger.Infoln("Connecting to DB.")
	client, err := config.DbConfigs.ConnectToMongo()
	if err != nil {
		logger.WithField("Error", err.Error()).Panicln("Could not connect to the database.")
	}
	logger.Infoln("Connected to Mongo.")
	return EntityRepository{client, config.DbConfigs}
}

func (e EntityRepository) Close() {
	defer e.client.Disconnect(context.Background())
}

func (e EntityRepository) instituteCollection() *mongo.Collection {
	return e.client.Database(e.config.DbName).Collection(institutesCollection)
}

func (e EntityRepository) occupationCollection() *mongo.Collection {
	return e.client.Database(e.config.DbName).Collection(institutesCollection)
}

func (e EntityRepository) educationQualificationsCollection() *mongo.Collection {
	return e.client.Database(e.config.DbName).Collection(institutesCollection)
}

func (e EntityRepository) CreateInstitute(ctx context.Context, institute *proto.Institute) error {
	_, err := e.instituteCollection().InsertOne(ctx, institute)
	if err != nil {
		return errors.Wrap(err, "Could not insert institute.")
	} else {
		return nil
	}
}

func (e EntityRepository) UpdateInstitute(ctx context.Context, institute *proto.Institute) error {
	if _, err := e.instituteCollection().UpdateOne(ctx, bson.M{"id": institute.Id}, bson.D{{"$set", institute}}); err != nil {
		return errors.Wrap(err, "Could not update the institute.")
	} else {
		return nil
	}
}

func (e EntityRepository) GetInstituteById(ctx context.Context, id string) (*proto.Institute, error) {
	result := proto.Institute{}
	if err := e.instituteCollection().FindOne(ctx, bson.M{"id": id}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the institute by id: %s", id))
	} else {
		return &result, nil
	}
}

func (e EntityRepository) GetAllInstitutes(ctx context.Context) ([]*proto.Institute, error) {
	var result []*proto.Institute
	cur, err := e.instituteCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "Could not get the institutes.")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		client := &proto.Institute{}
		err = cur.Decode(client)
		if err != nil {
			return nil, errors.Wrap(err, "Could not decode the institute.")
		}
		result = append(result, client)
	}
	return result, nil
}

func (e EntityRepository) CreateOccupation(ctx context.Context, occupation *proto.Occupation) error {
	_, err := e.occupationCollection().InsertOne(ctx, occupation)
	if err != nil {
		return errors.Wrap(err, "Could not insert occupation.")
	} else {
		return nil
	}
}

func (e EntityRepository) UpdateOccupation(ctx context.Context, occupation *proto.Occupation) error {
	if _, err := e.occupationCollection().UpdateOne(ctx, bson.M{"id": occupation.Id}, bson.D{{"$set", occupation}}); err != nil {
		return errors.Wrap(err, "Could not update the occupation.")
	} else {
		return nil
	}
}

func (e EntityRepository) GetOccupationById(ctx context.Context, id string) (*proto.Occupation, error) {
	result := proto.Occupation{}
	if err := e.occupationCollection().FindOne(ctx, bson.M{"id": id}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the occupation by id: %s", id))
	} else {
		return &result, nil
	}
}

func (e EntityRepository) GetAllOccupations(ctx context.Context) ([]*proto.Occupation, error) {
	var result []*proto.Occupation
	cur, err := e.occupationCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "Could not get the occupations.")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		client := &proto.Occupation{}
		err = cur.Decode(client)
		if err != nil {
			return nil, errors.Wrap(err, "Could not decode the occupation.")
		}
		result = append(result, client)
	}
	return result, nil
}

func (e EntityRepository) CreateEducationQualification(ctx context.Context, educationQualification *proto.EducationQualification) error {
	_, err := e.educationQualificationsCollection().InsertOne(ctx, educationQualification)
	if err != nil {
		return errors.Wrap(err, "Could not insert education qualification.")
	} else {
		return nil
	}
}

func (e EntityRepository) UpdateEducationQualification(ctx context.Context, educationQualification *proto.EducationQualification) error {
	if _, err := e.educationQualificationsCollection().UpdateOne(ctx, bson.M{"id": educationQualification.Id}, bson.D{{"$set", educationQualification}}); err != nil {
		return errors.Wrap(err, "Could not update the occupation.")
	} else {
		return nil
	}
}

func (e EntityRepository) GetEducationQualificationById(ctx context.Context, id string) (*proto.EducationQualification, error) {
	result := proto.EducationQualification{}
	if err := e.educationQualificationsCollection().FindOne(ctx, bson.M{"id": id}).Decode(&result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not get the education qualification by id: %s", id))
	} else {
		return &result, nil
	}
}

func (e EntityRepository) GetAllEducationQualifications(ctx context.Context) ([]*proto.EducationQualification, error) {
	var result []*proto.EducationQualification
	cur, err := e.educationQualificationsCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "Could not get the education qualifications.")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		client := &proto.EducationQualification{}
		err = cur.Decode(client)
		if err != nil {
			return nil, errors.Wrap(err, "Could not decode the education qualification.")
		}
		result = append(result, client)
	}
	return result, nil
}