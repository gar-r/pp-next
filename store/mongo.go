package store

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"okki.hu/garric/ppnext/model"
)

const MongoDatabase = "ppnext"
const MongoCollection = "rooms"

type MongoRepository struct {
	Client *mongo.Client
}

func NewMongoRepository() *MongoRepository {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &MongoRepository{
		Client: client,
	}
}

func (m *MongoRepository) rooms() *mongo.Collection {
	return m.Client.Database(MongoDatabase).Collection(MongoCollection)
}

func (m *MongoRepository) Load(name string) (*model.Room, error) {
	var r model.Room
	filter := bson.D{{
		Key:   "name",
		Value: name,
	}}
	err := m.rooms().FindOne(context.Background(), filter).Decode(&r)
	return &r, err
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
