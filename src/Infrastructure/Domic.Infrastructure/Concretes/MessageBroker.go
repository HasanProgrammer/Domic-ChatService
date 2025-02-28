package InfrastructureConcrete

import (
	"Domic.Domain/Commons/Contracts"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type MessageBroker struct {
	serializer DomainCommonContract.ISerializer
	connection *amqp.Connection
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (broker *MessageBroker) Subscribe(queue string, closure func(body []byte) error) {

	ch, err := broker.connection.Channel()

	if err != nil {
		ch.Close()
		failOnError(err, "Failed to open a channel")
	}

	err = ch.ExchangeDeclare(
		queue+"-e", // name
		"fanout",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)

	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		queue+"-q", // name
		true,       // durable
		false,      // delete when unused
		true,       // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,     // queue name
		"",         // routing key
		queue+"-e", // exchange
		false,
		nil,
	)

	failOnError(err, "Failed to bind a queue")

	//long runing goroutin
	go func() {

		messageChannel, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)

		failOnError(err, "Failed to register a consumer")

		waiter := &sync.WaitGroup{}

		var counter int = 0

		for {

			//throttle policy
			if counter <= 10000 {

				counter++

				waiter.Add(counter)

				newMessage := <-messageChannel

				//concurrent processing current message ( event )
				go func() {

					defer waiter.Done()

					result := closure(newMessage.Body)

					if result == nil {
						newMessage.Acknowledger.Ack(newMessage.DeliveryTag, false)
					} else {
						newMessage.Acknowledger.Reject(newMessage.DeliveryTag, false)
					}

				}()

			} else {

				waiter.Wait()

				counter = 0

			}

		}

	}()

}

func (broker *MessageBroker) Publish(event interface{}, exchange string) {
	ch, err := broker.connection.Channel()

	defer ch.Close()

	if err != nil {
		ch.Close()
		failOnError(err, "Failed to publish a message")
	}

	message, stringifyError := broker.serializer.Serialize(&event)

	if stringifyError != nil {

	}

	//todo: retry policy
	err = ch.Publish(exchange+"-e", "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

	failOnError(err, "Failed to publish a message")
}

func NewMessageBroker(connection string) *MessageBroker {
	conn, err := amqp.Dial(connection)
	failOnError(err, "Failed to connect to [RabbitMQ]")
	return &MessageBroker{connection: conn}
}
