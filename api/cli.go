package api

import (
	"bufio"
	"fmt"
	"iscoollab/filesystem/domain"
	"os"
	"strings"
)

const (
	//event of the user domain
	Register Event = "register" //register [username]

	//event of the folder domain
	CreateFolder Event = "create-folder" //create-folder [username] [foldername] [description]?
	DeleteFolder Event = "delete-folder" //delete-folder [username] [foldername]
	ListFolders  Event = "list-folders"  //list-folders [username] [--sort-name|--sort-created] [asc|desc]
	RenameFolder Event = "rename-folder" //rename-folder [username] [foldername] [new-folder-name]

	//event of the file domain
	CreateFile Event = "create-file" //create-file [username] [foldername] [filename] [description]?
	DeleteFile Event = "delete-file" //delete-file [username] [foldername] [filename]
	ListFiles  Event = "list-files"  //list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
)

type (

	// Command structures the necessary components of a command, organizing user and file operations.
	Command struct {
		Event     string
		User      string
		Folder    string
		NewFolder string
		File      string
		Desc      string
		SortBy    string
		SortOrder string
	}

	// Event defines a type for command identifiers.
	Event string

	// Handler defines the interface for command parsing and execution.
	Handler interface {
		parseCommand([]string)
		execute(*domain.Users)
	}
)

// CoolRun initiates a command-processing loop that reads from standard input.
func CoolRun() error {

	scanner := bufio.NewScanner(os.Stdin)
	users := domain.NewUsers()
	for {

		fmt.Print("\n enter command\n >> ")
		scanner.Scan()
		command := scanner.Text()

		if command == "exit" {
			fmt.Println("Exiting...")
			break
		}
		processCommand(command, users)
	}
	return nil
}

// processCommand interprets and dispatches commands based on input strings.
func processCommand(command string, users *domain.Users) {

	parts := strings.Fields(command)
	for _, v := range parts {
		fmt.Println("parts", v)
	}
	if len(parts) > 0 {

		//fetch the event from parts[0]
		e := parts[0]
		switch e {
		case string(Register), string(CreateFolder), string(DeleteFolder), string(ListFolders), string(RenameFolder), string(CreateFile), string(DeleteFile), string(ListFiles):

			s := eventFactory(e) // Create specific handler based on event type
			eventHandle(s, parts, users)
		default:
			fmt.Fprintln(os.Stderr, "Error: Unrecognized command")
		}
		return
	}

}

// eventHandle handles the execution of parsed commands.
func eventHandle(h Handler, s []string, users *domain.Users) {

	h.parseCommand(s)
	h.execute(users)
}

// eventFactory returns a Handler based on the event type.
func eventFactory(event string) Handler {

	// Factory pattern to instantiate handlers based on the event type.
	switch event {
	case string(Register):
		return &register{}
	case string(CreateFolder):
		return &createFolder{}
	case string(DeleteFolder):
		return &deleteFolder{}
	case string(ListFolders):
		return &listFolders{}
	case string(RenameFolder):
		return &renameFolder{}
	case string(CreateFile):
		return &createFile{}
	case string(DeleteFile):
		return &deleteFile{}
	case string(ListFiles):
		return &listFiles{}
	}
	return nil
}
