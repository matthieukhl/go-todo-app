package task

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	tasks = []Task{}
	AddTask("Test Task")
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Description != "Test Task" {
		t.Errorf("Expected task description 'Test Task', got %s", tasks[0].Description)
	}
}

func TestCompleteTask(t *testing.T) {
	tasks = []Task{{ID: 1, Description: "Test Task", Completed: false}}
	CompleteTask(1)
	if !tasks[0].Completed {
		t.Errorf("Expected task to be completed, got %v", tasks[0].Completed)
	}
}

func TestDeleteTask(t *testing.T) {
	tasks = []Task{{ID: 1, Description: "Test Task", Completed: false}}
	DeleteTaskNoConfirm(1)
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}

func TestGetCompletedTasks(t *testing.T) {
	tasks = []Task{
		{ID: 1, Description: "Test Incomplete Task", Completed: false},
		{ID: 2, Description: "Test Completed Task", Completed: true},
	}

	completedTasks := GetCompletedTasks()

	if len(completedTasks) != 1 {
		t.Errorf("Expected 1 completed task, got %d", len(completedTasks))
	}
	if completedTasks[0].Description != "Test Completed Task" {
		t.Errorf("Expected 'Test Completed Task', got %s", completedTasks[0].Description)
	}
}
