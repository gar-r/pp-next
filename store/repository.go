package store

import (
	"time"

	"okki.hu/garric/ppnext/model"
)

// Repository represents a storage interface for model.Room objects.
type Repository interface {

	// Load the model.Room with the given name.
	// Returns a new model.Room if a Room with the given name does not exist yet.
	// Returns an error if there is an underlying storage problem.
	Load(name string) (*model.Room, error)

	// Save the given model.Room.
	// Any existing model.Room with the same name gets overwritten.
	// Returns an error if there is an underlying storage problem.
	Save(room *model.Room) error

	// Delete the model.Room with the given name.
	// If the model.Room does not exist, nothing happens.
	// Returns an error if there is an underlying storage problem.
	Delete(name string) error

	// Exists returns true if the given user exists in any
	// room in the repository, or false otherwise.
	// Returns an error if there is an underlying storage problem.
	Exists(user string) (bool, error)

	// Remove removes a user from all rooms.
	// Returns an error if there is an underlying storage problem.
	Remove(user string) error

	// Cleanup removes obsolete model.Room files from the repository.
	// A model.Room is considered obsolete, when a certain amount
	// of time has elapsed since it was last updated.
	Cleanup(maxAge time.Duration) error
}
