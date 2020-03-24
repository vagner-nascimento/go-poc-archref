package integration

type subscription interface {
	getTopic() string
	getConsumer() string
	getHandler() func([]byte)
}
