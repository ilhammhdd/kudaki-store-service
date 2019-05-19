package safekit

import (
	"os"
	"os/signal"
	"sync"
)

type Job interface {
	Work() interface{}
	Handle(interface{})
}

type WorkerPool struct {
	Job    chan Job
	Work   chan func()
	Done   chan bool
	PoolWG sync.WaitGroup
	wg     sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	wp := &WorkerPool{}

	wp.Work = make(chan func())
	wp.Job = make(chan Job)
	wp.Done = make(chan bool)

	wp.PoolWG.Add(2)

	Do(func() {
		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt)

	JobLoop:
		for {
			select {
			case job := <-wp.Job:
				wp.wg.Add(1)
				Do(func() {
					job.Handle(job.Work())
					wp.wg.Done()
				})
			case <-wp.Done:
				break JobLoop
			case <-signals:
				break JobLoop
			}
		}

		close(wp.Job)
		wp.PoolWG.Done()
	})

	Do(func() {
		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt)

	WorkLoop:
		for {
			select {
			case work := <-wp.Work:
				wp.wg.Add(1)
				Do(func() {
					work()
				})
				wp.wg.Done()
			case <-wp.Done:
				break WorkLoop
			case <-signals:
				break WorkLoop
			}
		}

		close(wp.Work)
		wp.PoolWG.Done()
	})

	return wp
}
