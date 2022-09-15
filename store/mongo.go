package store

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"okki.hu/garric/ppnext/model"
)

const MongoDatabase = "ppnext"  // mongodb database name
const MongoCollection = "rooms" // mongodb collection name

// MongoRepository implements Repository using mongodb
type MongoRepository struct {
	Client *mongo.Client
}

// NewMongoRepository creates a new MongoRepository containing
// a connected client. The connection string used for the client
// is taken from the MONGODB_URI environment variable.
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

// Load the model.Room with the given name.
// If the collection does not contain a room with the given name,
// a new room is created and saved into the db.
func (m *MongoRepository) Load(name string) (*model.Room, error) {
	var r model.Room
	filter := bson.D{{
		Key:   "name",
		Value: name,
	}}
	err := m.rooms().FindOne(context.Background(), filter).Decode(&r)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r = *model.NewRoom(name)
			err = m.Save(&r)
			return &r, err
		}
	}
	return &r, err
}

// Save the given model.Room.
// This operation replaces the document in the collection.
// If no document with the give name exists, it will be inserted instead.
func (m *MongoRepository) Save(room *model.Room) error {
	filter := bson.D{{
		Key:   "name",
		Value: room.Name,
	}}
	_, err := m.rooms().ReplaceOne(context.Background(), filter, room, options.Replace().SetUpsert(true))
	return err
}

// Delete removes the room with the given name.
// If no room exists with the given name, the operation is still successful.
func (m *MongoRepository) Delete(name string) error {
	filter := bson.D{{
		Key:   "name",
		Value: name,
	}}
	_, err := m.rooms().DeleteOne(context.Background(), filter)
	return err
}

// Exists returns true if the given user exists in any
// room in the repository, or false otherwise.
func (m *MongoRepository) Exists(user string) (bool, error) {
	filter := bson.D{{
		Key: "votes." + user,
		Value: bson.D{{
			Key:   "$exists",
			Value: 1,
		}},
	}}
	n, err := m.rooms().CountDocuments(context.Background(), filter)
	return n > 0, err
}

// Remove removes a user from all rooms.
// Returns an error if there is an underlying storage problem.
func (m *MongoRepository) Remove(user string) error {
	key := fmt.Sprintf("votes.%s", user)
	filter := bson.D{{}}
	update := bson.D{{
		Key: "$unset",
		Value: bson.D{{
			Key:   key,
			Value: 1,
		}},
	}}
	_, err := m.rooms().UpdateMany(context.Background(), filter, update)
	return err
}

func (m *MongoRepository) RoomCount() (int, error) {
	filter := bson.D{{}}
	c, err := m.rooms().CountDocuments(context.Background(), filter)
	return int(c), err
}

func (m *MongoRepository) UserCount() (int, error) {
	pipeline := []interface{}{
		bson.D{{"$project", bson.D{{"votes", bson.D{{"$objectToArray", "$votes"}}}}}},
		bson.D{{"$project", bson.D{{"user", "$votes.k"}}}},
		bson.D{{"$unwind", "$user"}},
		bson.D{{"$group", bson.D{{"_id", "$user"}}}},
		bson.D{{"$count", "count"}},
	}
	cur, err := m.rooms().Aggregate(context.Background(), pipeline)
	if err != nil {
		return 0, err
	}
	if !cur.Next(context.Background()) {
		return 0, nil
	}
	var res bson.M
	if err = cur.Decode(&res); err != nil {
		return 0, err
	}
	count, ok := res["count"].(int32)
	if !ok {
		return 0, errors.New("invalid result")
	}
	return int(count), nil
}

// Cleanup removes any document from the collection, where
// the resetTs for the room is older than maxAge.
func (m *MongoRepository) Cleanup(maxAge time.Duration) error {
	ts := time.Now().Add(-maxAge)
	filter := bson.D{{
		Key: "resetTs",
		Value: bson.D{{
			Key:   "$lt",
			Value: ts,
		}},
	}}
	res, err := m.rooms().DeleteMany(context.Background(), filter)
	if err == nil {
		log.Printf("cleanup removed %d rooms\n", res.DeletedCount)
	}
	return err
}
