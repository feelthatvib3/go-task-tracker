# Task Tracker

A simple CLI application for managing tasks.

## Project URL
[Task Tracker Project](https://roadmap.sh/projects/task-tracker)

## Features
- Add tasks
- List tasks
- Update tasks
- Delete tasks
- Change task status

## Installation

To install and run the Task Tracker application, follow these steps:

### Prerequisites

- **Go**: Ensure you have Go installed on your machine. You can download and install Go from [golang.org](https://golang.org/dl/).

### Clone the Repository

1. Open your terminal or command prompt.
2. Clone the repository using the following command:

   ```bash
   git clone https://github.com/<your-username>/task-tracker.git
   ```

3. Navigate into the project directory:

   ```bash
   cd task-tracker
   ```

### Build the Application

To build the application, run the following command:

```bash
go build -o tt
```

This will create an executable file named `tt` in the project directory.

### Running the Application

You can run the application directly from the terminal after building:

```bash
./tt
```

Alternatively, if you're on Windows, you can run:

```bash
tt.exe
```

## Usage

The Task Tracker CLI allows you to manage your tasks through various commands. Below are the available commands and their descriptions.

### Add a Task

To add a new task, use the following command:

```bash
./tt add "Task description"
```

Example:

```bash
./tt add "Buy groceries"
```

### List Tasks

To list all tasks or filter by status, use the following command:

```bash
./tt list [status]
```

Examples:

- List all tasks:
  ```bash
  ./tt list
  ```

- List tasks with the status "done":
  ```bash
  ./tt list done
  ```

### Update a Task

To update the description of an existing task, use the following command:

```bash
./tt update <task-id> "<new-description>"
```

Example:

```bash
./tt update 1 "Buy groceries and cook dinner"
```

### Delete a Task

To delete a task, use the following command:

```bash
./tt delete <task-id>
```

Example:

```bash
./tt delete 1
```

### Change Task Status

To change the status of a task, you can use the following commands:

- Mark a task as "todo":
  ```bash
  ./tt mark-todo <task-id>
  ```

- Mark a task as "in-progress":
  ```bash
  ./tt mark-in-progress <task-id>
  ```

- Mark a task as "done":
  ```bash
  ./tt mark-done <task-id>
  ```

Example:

```bash
./tt mark-done 1
```

## License
This project is licensed under the MIT License.
