package executor_test

import (
	"os"
	"testing"

	"github.com/EMSSConsulting/Executor"
	"github.com/EMSSConsulting/Executor/shells" // Import the shells themselves
)

func TestTaskConstructor(t *testing.T) {
	t.Log("Testing operations argument")

	task, err := executor.NewTask([]string{"task1"}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(task.Operations) != 1 {
		t.Errorf("Expected task operations list to have length of 1, got %d", len(task.Operations))
	}

	if task.Operations[0] != "task1" {
		t.Errorf("Expected task operations to be [task1], got %#v", task.Operations)
	}

	if task.Args == nil {
		t.Error("Expected task Args to default to []string when nil is provided")
	}

	if task.Environment == nil {
		t.Error("Expected task Args to default to map[string]string when nil is provided")
	}

	t.Log("Testing args argument")

	task, err = executor.NewTask([]string{"task1"}, []string{"--arg"}, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(task.Args) != 1 {
		t.Errorf("Expected task args list to have length of 1, got %d", len(task.Args))
	}

	if task.Args[0] != "--arg" {
		t.Errorf("Expected task args to be [--arg], got %#v", task.Args)
	}

	t.Log("Testing environment argument")

	task, err = executor.NewTask([]string{"task1"}, nil, map[string]string{"test": "1"})
	if err != nil {
		t.Fatal(err)
	}

	if len(task.Environment) != 1 {
		t.Errorf("Expected task environment to have length of 1, got %d", len(task.Environment))
	}

	if task.Environment["test"] != "1" {
		t.Errorf("Expected task environemnt to be { test: 1 }, got %#v", task.Environment)
	}
}

func TestTaskUUID(t *testing.T) {
	task, err := executor.NewTask([]string{"task1"}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// A shell which just returns the raw UUID
	shell := &shells.TestShell{
		TestFilename: func(dir, id string) string {
			return id
		},
	}

	if task.ScriptFile(shell) == "" {
		t.Error("Expected tasks to have a non-nil UUID associated with them")
	}

	uuid := task.ScriptFile(shell)

	task, err = executor.NewTask([]string{"task1"}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if task.ScriptFile(shell) == uuid {
		t.Error("Expected each task to have a unique UUID irrespective of arguments")
	}
}

func TestTaskTempDirectory(t *testing.T) {
	task, err := executor.NewTask([]string{"task1"}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// A shell which just returns the raw temp directory
	shell := &shells.TestShell{
		TestFilename: func(dir, id string) string {
			return dir
		},
	}

	if task.ScriptFile(shell) != os.TempDir() {
		t.Error("Expected tasks use the operating system's temp directory")
	}
}
