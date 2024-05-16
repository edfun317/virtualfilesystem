package api

import (
	"fmt"
	"iscoollab/filesystem/domain"
	"iscoollab/filesystem/event"
	"os"
)

type (
	listFiles struct {
		Command
	}
)

// parseComands processes an input slice of strings to extract relevant command parameters.
func (l *listFiles) parseCommand(s []string) {

	length := len(s)
	if length >= 2 {
		l.User = s[1]
	}
	if length >= 3 {
		l.Folder = s[2]
	}

	if length >= 5 {
		l.SortBy = s[3]
		l.SortOrder = s[4]
	}
}

// Execute performs the list files query
func (l *listFiles) execute(users *domain.Users) {

	list, err := event.ListFiles(users, l.User, l.Folder, l.SortBy, l.SortOrder)
	if err != nil {

		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, v := range list {

		fmt.Println(v)
	}
}
