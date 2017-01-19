package mainthread_test

import (
	"fmt"

	"github.com/faiface/mainthread"
)

func run() {
	mainthread.Call(func() {
		// this function will be called on the main thread
		fmt.Println("Hello, world!")
	})
}

func ExampleRun() {
	mainthread.Run(run)
	// Output: Hello, world!
}
