package UseCaseCommonContract

type IMessageBroker interface {
	Subscribe(queue string)
	Publish(event interface{})
}
