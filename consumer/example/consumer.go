package exampleconsumer

import "org.idev.koala/backend/component/kafka"

type ExampleConsumer struct {
	consumerGroup *kafka.ConsumerGroup
}
