[![Go Report Card](https://goreportcard.com/badge/github.com/michebble/todo)](https://goreportcard.com/report/github.com/michebble/todo)

A ToDo list manager written in Go. From "Powerful Command-Line Applications in Go"

### Setup

Run all tests

```
$ go test ./...
```

Build program with

```
$ go build -o todo cmd/todo/main.go
```

### Usage

The program accepts the the following flags

--list

Displays the list of uncompleted tasks.

```
$ ./todo -list
  1: Make coffee
  2: Have breakfast
```

--task

Add a task to the ToDo list

```
$ ./todo --task "Clean teeth"
$ ./todo -list
  1: Make coffee
  2: Have breakfast
  3: Clean teeth
```

--complete

Marks the list at that position in the list as complete, removing it from the list

```
$ ./todo --complete 2
$ ./todo -list
  1: Make coffee
X 2: Have breakfast
  3: Clean teeth
```
