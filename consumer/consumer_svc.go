package consumer

import (
	"sync"
)

type consumerBuild struct {
	consumerType    string
	ConsumerHandler ConsumerHandler
}

type consumerControll struct {
	closeSignal chan struct{}
}

func NewConsumer() ConsumerBuildStep1 {
	return &consumerBuild{}
}

func (builder *consumerBuild) ConsumerType(consumerType string) ConsumerBuildStep2 {
	builder.consumerType = consumerType
	return builder
}

// Define HandleFunction to accept a function parameter
func (builder *consumerBuild) Handler(f ConsumerHandler) ConsumerBuildStep3 {
	builder.ConsumerHandler = f
	return builder
}

func (builder *consumerBuild) Open() ConsumerService {

	var closeWait sync.WaitGroup
	closeWait.Add(1)

	p := jobPoller{
		closeSignal: make(chan struct{}),
	}

	go p.RunJob(builder, &closeWait)

	return &consumerControll{
		closeSignal: p.closeSignal,
	}
}

func (cc *consumerControll) Close() {
	close(cc.closeSignal)
}
