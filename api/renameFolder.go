package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	renameFolder struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (r *renameFolder) parseCommand(s []string) {

	length := len(s)
	if length >= 2 {
		r.User = s[1]
	}
	if length >= 3 {
		r.Folder = s[2]
	}
	if length >= 4 {
		r.NewFolder = s[3]
	}
}

// Execute performs the rename folder operation
func (r *renameFolder) execute(users *domain.Users) {

	if err := event.RenameFolder(users, r.User, r.Folder, r.NewFolder); err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Rename '%s' successfully\n", r.User)

}
