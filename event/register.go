package event

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"os"
)

// Register add a new user with the name to the user collection.
// It calls the AddUser method from the user domain to handle the actual addition of the user.
// If AddUser returns an error (e.g., if the username already exists), the error is printed to standard error.
// On successful user registration, it prints a success message to standard output.
func Register(users *domain.Users, name string) {

	if err := users.AddUser(name); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Add '%s' successfully.\n", name)
}
