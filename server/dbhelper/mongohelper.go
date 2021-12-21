package dbhelper

import (
	"context"
	"datarepository/server/helper"
	"datarepository/server/helper/loggermdl"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoHost -MongoHost
type MongoHost struct {
	HostName        string        `json:"hostName"`
	Server          string        `json:"server"`
	Port            int           `json:"port"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	Database        string        `json:"database"`
	IsDefault       bool          `json:"isDefault"`
	MaxIdleConns    int           `json:"maxIdleConns" `
	MaxOpenConns    int           `json:"maxOpenConns"`
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" `
}

// TomlConfig - TomlConfig
type TomlConfig struct {
	MongoHosts map[string]MongoHost
}

var mutex sync.Mutex
var once sync.Once
var hostDetails MongoHost
var DefaultCtx context.Context

// ClientInstance -
var ClientInstance *mongo.Client
var DbInstance *mongo.Database

// Init initializes Mongo Connections for give toml file
func Init(tomlFilepath, hostName string) error {
	var sessionError error
	once.Do(func() {
		defer mutex.Unlock()
		mutex.Lock()
		_, err := helper.InitConfig(tomlFilepath, &hostDetails)
		if err != nil {
			loggermdl.LogError(err)
			sessionError = err
			return
		}

		var authCredentials *options.Credential
		// loggermdl.LogDebug(hostDetails.Password)
		if len(hostDetails.Password) > 0 {
			authCredentials = &options.Credential{
				AuthMechanism: "SCRAM-SHA-1",
				AuthSource:    hostDetails.Database,
				Username:      hostDetails.Username,
				Password:      hostDetails.Password,
				PasswordSet:   false,
			}
		}
		ctx := context.Background()
		client, err := mongo.Connect(ctx, &options.ClientOptions{
			Auth:  authCredentials,
			Hosts: []string{hostDetails.Server},
		})
		if err != nil {
			sessionError = err
			return
		}
		DefaultCtx = context.Background()
		ClientInstance = client
		DbInstance = client.Database(hostDetails.Database)
		// loggermdl.LogDebug(client.Database("datarepo").ListCollectionNames(DefaultCtx, bson.M{}))
	})
	return sessionError
}
