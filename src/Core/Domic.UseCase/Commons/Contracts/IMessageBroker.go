package UseCaseCommonContract

type IMessageBroker interface {
	Subscribe(queue string, closure func(body []byte) error)
	Publish(event interface{}, exchange string)
}
