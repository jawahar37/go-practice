# print-directory

`print-directory` is a simple Go module to print the structure of a given directory in ASCII format. It can display files and subdirectories with proper indentation, making it easier to visualize the hierarchy of a directory.

## Features

- Prints the structure of a directory in ASCII format
- Handles nested directories
- Includes test cases to verify functionality

## Installation

Clone the repository:

To run the program and print the structure of a directory, use:
```sh
go run main.go /path/to/directory
```
Replace /path/to/directory with the path to the directory you want to display.

## Example
Given the following directory structure:

```
example-dir
├── file1.txt
├── subdir1
│   └── file2.txt
└── subdir2
```

Running go run main.go example-dir will output:

```
example-dir
  file1.txt
  subdir1
    file2.txt
  subdir2
```
## Tests
The module includes test cases to cover scenarios, such as:
- Printing an empty directory
- Printing a directory with one file
- Printing a directory with only subdirectories
- Printing a directory with nested subdirectories

To run the tests, use:
```sh
go test -v
```