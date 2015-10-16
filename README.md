# Executor [![GoDoc](https://godoc.org/github.com/EMSSConsulting/Executor?status.png)](https://godoc.org/github.com/EMSSConsulting/Executor) [![Build Status](https://travis-ci.org/EMSSConsulting/Executor.svg)](https://travis-ci.org/EMSSConsulting/Executor)
**Go based task runner for cross platform script execution**

Executor is designed to simplify the task of running user configurable scripts
across a wide range of platforms. This use case is common when writing a test
runner or automated deployment tool like [Depro](https://github.com/EMSSConsulting/Depro).

## Features

 - **Simple API** allows you to quickly get started and minimize boilerplate
 - **Pluggable Shells** allow you to pick and chose the shell plugins you wish to use,
   or even create your own.
 - **Powerful Configuration** of each task's arguments and environment, as well
   as the environment of your executor instance - allowing you to quickly build
   hierarchical environments.

## Usage

```go
package main

import (
    "os"

    executor "github.com/EMSSConsulting/Executor"
)

func main() {
    exec := executor.NewExecutor("powershell")
    exec.Environment["RUNNER"] = "Executor"

    task := executor.NewTask([]string{
        "Write-Host 'Hello World!'",
    }, nil, nil)

    err := exec.Run(task)
    if err != nil {
        os.Exit(1)
    }
}
```
