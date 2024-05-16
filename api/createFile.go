package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	createFile struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (c *createFile) parseCommand(s []string) {

	length := len(s)

	if length >= 2 {
		c.User = s[1]
	}
	if length >= 3 {
		c.Folder = s[2]
	}
	if length >= 4 {
		c.File = s[3]
	}
	if length >= 5 {
		c.Desc = s[4]
	}
}

// Execute performs the create file operation
func (c *createFile) execute(users *domain.Users) {

	if err := event.CreateFile(users, c.User, c.Folder, c.File, c.Desc); err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Create '%s' in '%s/%s' successfully\n", c.File, c.User, c.Folder)

}
