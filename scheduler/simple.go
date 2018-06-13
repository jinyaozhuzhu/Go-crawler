package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	//这里是的workerChan就是被配置成了in
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

//此处submit成功的前提是必须有空闲的worker在等待（在消耗WorkerChan）
//有空闲的worker在等待的前提是worker必须把上一件事情做完
//worker必须把上一件事情做完的前提是engine必须提交request成功
//以上形成循环等待

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkerChan <- r
	}()
}
