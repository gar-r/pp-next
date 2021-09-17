package store

import "okki.hu/garric/ppnext/model"

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
}
