package mainthread_test

import (
	"fmt"

	"github.com/faiface/mainthread"
)

func run() {
	mainthread.Call(func() {
		fmt.Println("i'm printing from the main thread")
	})
}

func ExampleRun() {
	mainthread.Run(run)
	// Output: i'm printing from the main thread
}
