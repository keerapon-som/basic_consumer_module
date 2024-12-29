package consumer

import (
	"fmt"
	"sync"
	"time"
)

type jobPoller struct {
	closeSignal chan struct{}
}

func (poller *jobPoller) RunJob(builder *consumerBuild, closeWait *sync.WaitGroup) {
	defer closeWait.Done()
	fmt.Println("Running job")

	builder.ConsumerHandler(Consumer{
		Id:      "123",
		Message: "Hello, World!",
		Jobtype: builder.consumerType,
	})

	for {
		select {
		case <-poller.closeSignal:
			fmt.Println("Close signal received")
			return
		default:
			fmt.Println("Close signal is nil")
			time.Sleep(1 * time.Second)
		}
	}

}
