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
Clean teeth
```

--task

Add a task to the ToDo list

```
$ ./todo --task "Buy milk"
$ ./todo -list
Clean teeth
Buy milk
```

--complete

Marks the list at that position in the list as complete, removing it from the list

```
$ ./todo -list
Clean teeth
Buy milk
$ ./todo --complete 1
$ ./todo -list
Buy milk
```
