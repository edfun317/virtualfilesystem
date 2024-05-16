package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	createFolder struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (c *createFolder) parseCommand(s []string) {
	length := len(s)

	if length >= 2 {
		c.User = s[1]
	}

	if length >= 3 {
		c.Folder = s[2]
	}

	if length >= 4 {
		c.Desc = s[3]
	}
}

// Execute performs the create folder operation
func (c *createFolder) execute(users *domain.Users) {

	if err := event.CreateFolder(users, c.User, c.Folder, c.Desc); err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Create '%s' successfully\n", c.User)

}
