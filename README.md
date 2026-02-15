# TODOlist-GO

Command-line TODO list application written in Go.

## Requirements

- Go 1.x or higher

## Installation

```bash
git clone https://github.com/ivaaneoski/TODOlist-GO.git
cd TODOlist-GO
```

## Run

```bash
go run main.go
```

## Build

```bash
go build -o todolist
./todolist
```

## Usage

The app runs in an interactive loop with the following commands:

**add** - Add a new task
```
> add
Enter the title: Buy groceries
```

**list** - Display all tasks with status
```
> list
Your Tasks:
1. Buy groceries ❌
2. Finish project ✅
```

**done** - Mark a task as complete
```
> done
Enter ID of the task which you want to mark done: 1
```

**delete** - Remove a task
```
> delete
Enter ID of the task which you want to delete: 2
```

**quit** - Exit the application

## Implementation Details

- Tasks are stored in memory (not persisted to disk)
- Each task has an ID, title, and completion status
- Tasks marked as done show ✅, incomplete tasks show ❌
- All data is lost when the program exits
