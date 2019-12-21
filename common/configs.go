package common

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"light-up-backend/common/middleware"
	"os"
	"time"
)

const (
	AuthenticationServiceName = "light-up-authentication-service"
	LightSeekerServiceName    = "light-up-light-seeker-service"
	LighterServiceName        = "light-up-lighter-service"
	BFFServiceName            = "light-up-bff-service"
	LighterServiceDB          = "light-up-lighters"
	LightSeekerServiceDB      = "light-up-light-seekers"
)

type ApplicationDbCredentials struct {
	Username string
	Password string
	Hostname string
	Port     string
}

type ServiceDbConfigurations struct {
	ApplicationDbCredentials
	DbName string
}

type ServiceConfigurations struct {
	ServiceName string
	DbConfigs   ServiceDbConfigurations
}

func (s ServiceDbConfigurations) ConnectToMongo() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb://" + s.Hostname
	if s.Username != "" && s.Password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@"+s.Hostname, s.Username, s.Password)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		client.Disconnect(context.Background())
		return nil, err
	}
	return client, nil
}

type ApplicationConfig struct {
	Environment   string
	DbCredentials ApplicationDbCredentials
}

func (a ApplicationConfig) IsLocalEnv() bool {
	return a.Environment == "local"
}

func (a ApplicationConfig) RequiresSetup() bool {
	return a.Environment == "local" || a.Environment == "dev"
}

type BffConfig struct {
	ServiceName string
	Port        string
}

func (a ApplicationConfig) BffConfig() BffConfig {
	return BffConfig{
		ServiceName: BFFServiceName,
		Port:        ":20021",
	}
}

func (a ApplicationConfig) LighterServiceConfiguration() ServiceConfigurations {
	return a.buildServiceConfig(
		LighterServiceName,
		LighterServiceDB,
	)
}

func (a ApplicationConfig) LightSeekerServiceConfiguration() ServiceConfigurations {
	return a.buildServiceConfig(
		LightSeekerServiceName,
		LightSeekerServiceDB,
	)
}

func (a ApplicationConfig) AuthenticationServiceConfig() ServiceConfigurations {
	return a.buildServiceConfig(
		AuthenticationServiceName,
		"",
	)
}

func (a ApplicationConfig) buildServiceConfig(serviceName, serviceDbName string) ServiceConfigurations {
	return ServiceConfigurations{
		ServiceName: serviceName,
		DbConfigs: ServiceDbConfigurations{
			ApplicationDbCredentials: a.DbCredentials,
			DbName:                   serviceDbName,
		},
	}
}

func buildApplicationConfig(env string, DbCredentials ApplicationDbCredentials) ApplicationConfig {
	return ApplicationConfig{
		Environment:   env,
		DbCredentials: DbCredentials,
	}
}

func buildDbCredentials(hostname, port, username, password string) ApplicationDbCredentials {
	return ApplicationDbCredentials{
		Hostname: hostname,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func LocalApplicationConfig() ApplicationConfig {
	return buildApplicationConfig(
		"local",
		buildDbCredentials("localhost", "3306", "", ""),
	)
}

func DevApplicationConfig() ApplicationConfig {
	return buildApplicationConfig(
		"dev",
		buildDbCredentials("", "", "root", "12345678"),
	)
}

func GetApplicationConfig() ApplicationConfig {
	env := os.Getenv("MICRO_ENV")
	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	logger.WithField("MICRO_ENV", env).Infoln("Loading configs.")
	switch env {
	case "local":
		return LocalApplicationConfig()
	case "dev":
		return DevApplicationConfig()
	default:
		panic(fmt.Sprintf("Could not source the application config for Environment = %s", env))
	}
}
