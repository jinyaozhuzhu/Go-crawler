package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	panic("implement me")

}

func (s *SimpleScheduler) Submit(r engine.Request) {
	s.WorkerChan <- r
}
