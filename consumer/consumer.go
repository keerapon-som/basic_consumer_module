package consumer

type Consumer struct {
	Id      string
	Message string
	Jobtype string
}

type ConsumerHandler func(ConsumerClient Consumer)

type ConsumerBuildStep1 interface {
	ConsumerType(consumerType string) ConsumerBuildStep2
}

type ConsumerBuildStep2 interface {
	Handler(ConsumerHandler) ConsumerBuildStep3
}

type ConsumerBuildStep3 interface {
	Open() ConsumerService
}

type ConsumerService interface {
	Close()
}
