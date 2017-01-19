package mainthread_test

import (
	"fmt"

	"github.com/faiface/mainthread"
)

func run() {
	mainthread.Call(func() {
		fmt.Println("Hello, world!")
	})
}

func ExampleRun() {
	mainthread.Run(run)
}
