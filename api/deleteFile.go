package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	deleteFile struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (d *deleteFile) parseCommand(s []string) {
	length := len(s)
	if length >= 2 {
		d.User = s[1]
	}
	if length >= 3 {
		d.Folder = s[2]
	}
	if length >= 4 {
		d.File = s[3]
	}
}

// Execute performs the delete file operation
func (d *deleteFile) execute(users *domain.Users) {

	if err := event.DeleteFile(users, d.User, d.Folder, d.File); err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Delete '%s' successfully\n", d.User)

}
