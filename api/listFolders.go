package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	listFolders struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (l *listFolders) parseCommand(s []string) {
	length := len(s)

	if length >= 2 {
		l.User = s[1]
	}

	if length >= 4 {
		l.SortBy = s[2]
		l.SortOrder = s[3]
	}
}

// Execute performs the list folders query
func (l *listFolders) execute(users *domain.Users) {

	list, err := event.ListFolders(users, l.User, l.SortBy, l.SortOrder)
	if err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, v := range list {

		fmt.Println(v)
	}
}
