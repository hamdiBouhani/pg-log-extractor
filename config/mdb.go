package config

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConn(mongoAddr, rs string, hosts []string) *types.MongoConn {
	log := logrus.New()
	var Config = types.MongoConn{}
	if mongoAddr == "" {
		log.Fatalf("Failed to create mongo client: %v", errors.New("mongo Addr should not be empty"))
	}

	Config.MongoAddr = mongoAddr
	opt := options.Client().ApplyURI(mongoAddr)

	// if len(rs) > 0 && len(hosts) > 0 {
	// 	opt.SetHosts(hosts)
	// 	opt.SetReadPreference(readpref.Primary())
	// 	opt.SetServerSelectionTimeout(time.Duration(2 * time.Second))
	// 	opt.SetReplicaSet(rs)
	// }

	MongoClient, err := mongo.NewClient(opt)
	if err != nil {
		log.Fatalf("Failed to create mongo client: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	log.Info("connect to Mongodb ")

	err = MongoClient.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to create mongo client: %v", errors.Wrap(err, "unable to connect to mongo db"))
	}

	Config.MongoClt = MongoClient
	return &Config
}
