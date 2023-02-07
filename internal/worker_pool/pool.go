package worker_pool

import (
	"fmt"
	"log"
	"sync"
)

type Task interface {
	Execute() error
	OnFailure(error)
}

type Pool interface {
	Start()
	Stop()
	AddWork(Task)
}

type SimplePool struct {
	numWorkers int
	tasks      chan Task
	start      sync.Once
	stop       sync.Once
	quit       chan struct{}
}

var ErrNoWorkers = fmt.Errorf("attempting to create worker pool with less than 1 worker")
var ErrNegativeChannelSize = fmt.Errorf("attempting to create worker pool with a negative channel size")

func NewSimplePool(numWorkers int, channelSize int) (Pool, error) {
	if numWorkers <= 0 {
		return nil, ErrNoWorkers
	}
	if channelSize < 0 {
		return nil, ErrNegativeChannelSize
	}

	tasks := make(chan Task, channelSize)

	return &SimplePool{
		numWorkers: numWorkers,
		tasks:      tasks,
		start:      sync.Once{},
		stop:       sync.Once{},
		quit:       make(chan struct{}),
	}, nil
}

func (p *SimplePool) Start() {
	p.start.Do(func() {
		log.Println("starting a simple worker pool")
		p.startWorkers()
	})
}

func (p *SimplePool) Stop() {
	p.stop.Do(func() {
		log.Print("stopping simple worker pool")
		close(p.quit)
	})
}

func (p *SimplePool) AddWork(t Task) {
	select {
	case p.tasks <- t:
	case <-p.quit:
	}
}

func (p *SimplePool) startWorkers() {
	for i := 0; i < p.numWorkers; i++ {
		go func(workerNum int) {
			log.Printf("starting worker %d", workerNum)

			for {
				select {
				case <-p.quit:
					log.Printf("stopping worker %d with quit channel\n", workerNum)
					return
				case task, ok := <-p.tasks:
					if !ok {
						log.Printf("stopping worker %d with closed tasks channel\n", workerNum)
						return
					}

					if err := task.Execute(); err != nil {
						task.OnFailure(err)
					}
				}
			}
		}(i)
	}
}
