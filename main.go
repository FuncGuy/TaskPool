package main

import (
	"TaskPool/worker"
	"time"
)

func main() {
defer worker.Wait()
for i :=0; i < 10000000; i++ {
	job := worker.Job{
		Action: PrintPayload,
		Payload: map[string]string{
			"time": time.Now().String(),
		},
	}
	job.Fire()
}

time.Sleep(time.Second * 3)
}
