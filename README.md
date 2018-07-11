# Loki

A simple way to create fakes in Go

## Basic Usage

```go
package example

import "github.com/rdumont/loki"

type TodoAdder interface {
    // Add creates a new todo and returns its position in the list
    Add(description string) int
}

type FakeTodoAdder interface {
    AddCalls    loki.Method
}

func (ta *FakeTodoAdder) Add(description string) int {
    return ta.AddCalls.Receive(description).Get(0)
}
```

```go
package example

func DoSomething() {
    adder := new(FakeTodoAdder)

    adder.AddCalls.On("some todo").Return(5)

    pos := adder.Add("some todo")
    fmt.Println("Position is", pos)
    // pos == 5
}
```