package scheduler

import "go-reptile/engine"

/**
		持续将request放入channel，安排worker进行工作
 */
type QueuedScheduler struct {
	requestChan chan engine.Request  //创建一个engine.Request类型通道
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request{
	// 构建一个engine.Request类型通道
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request)  {
	//向通道中放入值
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request)  {
	//向通道中放入值 w
	s.workerChan <- w
}

//执行
func (s *QueuedScheduler) Run()  {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	//匿名函数并发执行(闭包)
	go func() {
		var requestQ  []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <- s.requestChan :  //获取s.requestChan 通道的值
				requestQ = append(requestQ,r)
			case w := <- s.workerChan :
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:  //将activeRequest 放入 activeWorker 通道中
				workerQ = workerQ[1:]
				requestQ = requestQ[1:] // 删除开头1个元素
			}
		}
	}()
}