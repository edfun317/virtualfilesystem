# Virtual File System

This project implements a virtual file system with user and file management capabilities using Go 1.20.14.

## Supported Go Version

- Go 1.20.14 and above (1.20.x)

## Running the Program

To run the program, follow these steps:

1. Ensure you have Go 1.20.14 or a higher 1.20.x version installed on your machine.
2. Extract the project files from the provided archive.
3. Open a terminal and navigate to the project directory.
4. Run the following command to start the program:
    ```sh
    go run main.go
    ```

## Validation Rules

### User, Folder, and File Name Rules

- Length: 3 to 20 characters
- Case insensitive
- Only alphanumeric characters (letters and numbers)
- No spaces allowed (it will be treated as different parameters)

## Command Instructions

Below are the available commands and their respective responses.

### User Registration

#### Command

```sh
register [username]
```

#### Response

- **Success**: 
  - `Add '[username]' successfully.`
- **Error**:
  - `The '[username]' has already existed.`
  - `The '[username]' contains invalid chars.`

### Folder Management

#### Create Folder

##### Command

```sh
create-folder [username] [foldername] [description]?
```

##### Response

- **Success**: 
  - `Create '[foldername]' successfully.`
- **Error**:
  - `The '[username]' doesn't exist.`
  - `The '[foldername]' contains invalid chars.`
  - `The '[foldername]' has already existed.`

#### Delete Folder

##### Command

```sh
delete-folder [username] [foldername]
```

##### Response

- **Success**: 
  - `Delete '[foldername]' successfully.`
- **Error**:
  - `The '[username]' doesn't exist.`
  - `The '[foldername]' doesn't exist.`

#### List Folders

##### Command
```sh
list-folders [username] [--sort-name|--sort-created] [asc|desc]
```

##### Response

- **List of folders**:
  - `[foldername] [description] [created at] [username]`
- **Warning**:
  - `The '[username]' doesn't have any folders.`
- **Error**:
  - `The '[username]' doesn't exist.`

#### Rename Folder

##### Command
```sh
rename-folder [username] [foldername] [new-folder-name]
```

##### Response

- **Success**: 
  - `Rename '[foldername]' to '[new-folder-name]' successfully.`
- **Error**:
  - `The '[username]' doesn't exist.`
  - `The '[foldername]' doesn't exist.`

### File Management

#### Create File

##### Command

```sh
create-file [username] [foldername] [filename] [description]?
```

##### Response

- **Success**: 
  - `Create '[filename]' in '[username]/[foldername]' successfully.`
- **Error**:
  - `The '[username]' doesn't exist.`
  - `The '[foldername]' doesn't exist.`
  - `The '[filename]' contains invalid chars.`
  - `The '[filename]' has already existed.`

#### Delete File

##### Command

```sh
delete-file [username] [foldername] [filename]
```
##### Response

- **Success**: 
  - `Delete '[filename]' from '[username]/[foldername]' successfully.`
- **Error**:
  - `The '[username]' doesn't exist.`
  - `The '[foldername]' doesn't exist.`
  - `The '[filename]' doesn't exist.`

#### List Files

##### Command
```sh
list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
```

##### Response

- **List of files**:
  - `[filename] [description] [created at] [foldername] [username]`
- **Warning**:
  - `The folder is empty.`
- **Error**:
  - `The '[username]' doesn't exist.`
  - `The '[foldername]' doesn't exist.`

## Project Design and Testing

I designed the architecture based on the concepts of Domain-Driven Design (DDD), striving to maintain a balance between design and complexity to avoid over-engineering. Unit tests, integration tests, and end-to-end tests are conducted separately to ensure comprehensive test coverage.
