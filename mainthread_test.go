package mainthread_test

import (
	"errors"
	"testing"

	"github.com/faiface/mainthread"
)

func TestRunE(t *testing.T) {
	ok := false
	fun := func() error {
		mainthread.Call(func() {
			ok = true
		})
		return errors.New("this is a test")
	}

	err := mainthread.RunE(fun)
	if err == nil || err.Error() != "this is a test" {
		t.Errorf("Did not receive expected error but %+v", err)
	}

	if ok != true {
		t.Error("Function should have made a mainthread call but it did not")
	}
}

func BenchmarkCall(b *testing.B) {
	run := func() {
		f := func() {}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mainthread.Call(f)
		}
	}
	mainthread.Run(run)
}

func BenchmarkCallErr(b *testing.B) {
	run := func() {
		f := func() error {
			return errors.New("foo")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mainthread.CallErr(f)
		}
	}
	mainthread.Run(run)
}

func BenchmarkCallVal(b *testing.B) {
	run := func() {
		f := func() interface{} {
			return 42
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mainthread.CallVal(f)
		}
	}
	mainthread.Run(run)
}
