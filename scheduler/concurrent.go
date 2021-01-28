package scheduler

import "github.com/liang24/go-crawler/engine"

type ConcurrentScheduler struct {
	workerChan chan engine.Request
}

func (s *ConcurrentScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

func (s *ConcurrentScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *ConcurrentScheduler) WorkerReady(chan engine.Request) {

}

func (s *ConcurrentScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
