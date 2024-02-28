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
```

--add

Add a task to the ToDo list

```
$ ./todo --add Have breakfast
$ ./todo -list
  1: Make coffee
  2: Have breakfast
```

Tasks can also be accepted via standard in

```
$ echo "Clean teeth" | ./todo -add
$ ./todo -list
  1: Make coffee
  2: Have breakfast
  3: Clean teeth
```

--complete

Marks the task at that position in the list as complete

```
$ ./todo --complete 2
$ ./todo -list
  1: Make coffee
X 2: Have breakfast
  3: Clean teeth
```

--delete

Permanently removes the task at that position from the list

```
$ ./todo --delete 1
$ ./todo -list
X 1: Have breakfast
  2: Clean teeth
```

By default the ToDo list will be saved in `.todo.json`, but a different file can be specified by setting the `TODO_FILENAME` ENV.

```
export TODO_FILENAME=my-todo.json
```
