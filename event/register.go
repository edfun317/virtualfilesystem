package event

import (
	"errors"
	"iscoollab/filesystem/domain"
)

// Register add a new user with the name to the user collection.
// It calls the AddUser method from the user domain to handle the actual addition of the user.
func Register(users *domain.Users, name string) error {

	if name == "" {
		err := errors.New("Usage: register [username]")
		return err
	}
	if err := users.AddUser(name); err != nil {

		return err
	}

	return nil
}
