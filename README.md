# CLI Task Tracker 📝

A powerful command-line task management tool built with Go! Keep track of your tasks, manage their status, and stay organized - all from your terminal! 🚀

## Features ✨

- ➕ Add new tasks
- 📋 List all tasks
- 🔄 Update existing tasks
- 🗑️ Delete tasks
- 📊 Track task status (todo, in-progress, done)
- ⏰ Automatic timestamp tracking
- 💾 Persistent storage using JSON

## Installation 🛠️

1. Make sure you have Go installed on your system
2. Clone this repository
3. Navigate to the `task-tracker` directory
4. Run the program using Go

## Usage 💻

Here are the available commands:

### Add a Task 📝
```bash
go run main.go add "Your task description"
```

### List Tasks 📋
```bash
# List all tasks
go run main.go list

# List tasks with specific status
go run main.go list todo
go run main.go list in-progress
go run main.go list done
```

### Update a Task ✏️
```bash
go run main.go update [task-id] "New task description"
```

### Delete a Task 🗑️
```bash
go run main.go delete [task-id]
```

### Update Task Status 🔄
```bash
# Mark task as in progress
go run main.go mark-in-progress [task-id]

# Mark task as done
go run main.go mark-done [task-id]
```

## Task Storage 💾

Tasks are stored in a `tasks.json` file in the following format:
```json
{
  "id": 1,
  "description": "Task description",
  "status": "todo",
  "createdAt": "2024-XX-XX...",
  "updatedAt": "2024-XX-XX..."
}
```

## Why Use CLI Task Tracker? 🤔

- 🚀 **Fast & Efficient**: Manage tasks directly from your terminal
- 🎯 **Focused**: No distracting UI elements
- 💪 **Powerful**: Complete task management functionality
- 🔄 **Status Tracking**: Keep track of task progress
- ⏱️ **Time Awareness**: Automatic creation and update timestamps
- 📁 **Persistent**: All tasks safely stored in JSON format

## Contributing 🤝

Feel free to fork this repository and submit pull requests! You can also open issues if you find any bugs or have feature suggestions! 🐛✨

## License 📄

This project is open source and available under the MIT License. Feel free to use it as you wish! 🆓