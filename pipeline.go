package hw06pipelineexecution

import (
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	out := make(Bi)
	go func() {
		for i := 0; i < 6; i++ {
			out <- in
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return out
}
