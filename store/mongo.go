package store

import (
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"okki.hu/garric/ppnext/model"
)

type MongoRepository struct {
	Client *mongo.Client
}

func NewMongoRepository() *MongoRepository {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &MongoRepository{
		Client: client,
	}
}

func (*MongoRepository) Load(name string) (*model.Room, error) {
	return nil, nil
}

func (*MongoRepository) Save(room *model.Room) error {
	return nil
}

func (*MongoRepository) Delete(name string) error {
	return nil
}

func (*MongoRepository) Exists(user string) (bool, error) {
	return false, nil
}

func (*MongoRepository) Cleanup(maxAge time.Duration) error {
	return nil
}
