package shells

import (
	"path"

	"github.com/SierraSoftworks/Executor"
)

// TestShell provides an example shell which allows you to provide runtime methods
// for its various features.
type TestShell struct {
	ShellAbstract

	TestCommand    string
	TestFilename   func(dir, id string) string
	TestArgs       func(executor *executor.Executor, task *executor.Task) []string
	TestOperations func(executor *executor.Executor, task *executor.Task) []string
}

// Name retrieves the name of this shell as it will be used when requesting a new
// executor instance.
func (s *TestShell) Name() string {
	return "test"
}

// Command retrieves the relative or full path to the executable which will be run
// for this shell.
// If a relative path is provided, the user's PATH environment variable will be
// searched in an attempt to locate the file to execute.
func (s *TestShell) Command() string {
	return s.TestCommand
}

// Filename is responsible for building a valid filename from the taskID and the
// temporary task directory.
func (s *TestShell) Filename(directory, taskID string) string {
	if s.TestFilename != nil {
		return s.TestFilename(directory, taskID)
	}

	return path.Join(directory, taskID)
}

// Args is responsible for building the full list of command line arguments
// which will be passed to the shell's executable.
// It should include the file to be run, which can be acquired by calling
// `task.ScriptFile(s)`.
func (s *TestShell) Args(executor *executor.Executor, task *executor.Task) []string {
	if s.TestArgs != nil {
		return s.TestArgs(executor, task)
	}

	return append([]string{task.ScriptFile(s)}, task.Args...)
}

// Operations is responsible for building the list of operations that will
// be included in the task script.
func (s *TestShell) Operations(executor *executor.Executor, task *executor.Task) []string {
	if s.TestOperations != nil {
		return s.TestOperations(executor, task)
	}

	return task.Operations
}
