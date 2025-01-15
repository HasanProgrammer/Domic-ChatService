package InfrastructureConcrete

type MessageBroker struct {
}

func (broker *MessageBroker) Subscribe(queue string) {

}

func (broker *MessageBroker) Publish(event interface{}) {

}

func NewMessageBroker() *MessageBroker {
	return &MessageBroker{}
}
