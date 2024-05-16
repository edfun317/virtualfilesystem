package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	register struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (r *register) parseCommand(s []string) {

	length := len(s)
	if length >= 2 {
		r.User = s[1]
	}
}

// Execute performs the user registration operation
func (r *register) execute(users *domain.Users) {

	if err := event.Register(users, r.User); err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Add '%s' successfully\n", r.User)
}
