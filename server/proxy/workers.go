package proxy

import (
	"net"
)

type Job struct {
	s *Server
	conn net.Conn
}

var JobQueue chan Job

type Worker struct {
	WorkerPool  chan chan Job
	JobChannel  chan Job
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job)}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				NewSession(job.s, job.conn).Serve()

			}
		}
	}()
}
