# Go Todo App

This is a simple CLI-based Todo application built in Go. It allows you to manage tasks by adding, marking as complete, viewing, and deleting them. Tasks are saved to a JSON file so they persist between sessions.

## Features

- Add tasks with descriptions
- Mark tasks as complete
- Delete tasks with confirmation
- View all tasks or individual tasks by ID
- Persistent storage using a JSON file
- Modular code structure with a `task` package
- Unit tests for core functionality

## Getting Started

### Prerequisites

- Go installed on your machine (version 1.18+ recommended)

### Installation

1. Clone the repository to your local machine.
2. Initialize the module if not already done:

   ```bash
   go mod init todo-app
   ```

3. Usage

   ```bash
   go run main.go
   ```

## Usage
Upon running the application, a menu will be displayed with options:
1. List all tasks - View all tasks with their completion status.
2. Add a task - Create a new task by providing a description.
3. Mark a task as complete - Set the completion status of a task by entering its ID.
4. View a specific task - Display details of a task by entering its ID.
5. Delete a task - Remove a task by ID, with confirmation.
6. Exit - Quit the application.

## File Structure
- `main.go` - Contains the main program and manu handling
- `task/` - Contains the `Task` struct, functions for managing tasks, and data persistence.
- `task/task_test.go` - Contains unit tests for task-related functions.

## Testing
To run the unit tests for this project:

   ```bash
   go test ./task
   ```
This will execute tests for the core functionality in the `task` package, ensuring that adding, completing and deleting tasks works as expected.

## Future Enhancements
- **Filtering** - Add options to filter completed and incomplete tasks.
- **API Conversion** - Convert the application to a REST API using `net/http` and `gorilla/mux`.
- **Enhanced Validation** - Add checks to prevent adding empty tasks.

## License
This project is open-source and available under the MIT License.