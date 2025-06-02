package bootstrap

import (
	golibmsg "github.com/golibs-starter/golib-message-bus"
	"go.uber.org/fx"
)

func WithKafkaOnlyConsumer() fx.Option {
	return fx.Options(
		// register message bus
		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaConsumerOpt(),
		golibmsg.OnStopConsumerOpt())
}

func WithKafkaOnlyProducer() fx.Option {
	return fx.Options(
		// register message bus
		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaProducerOpt(),
		golibmsg.OnStopProducerOpt())
}

func WithKafka() fx.Option {
	return fx.Options(
		// register message bus
		golibmsg.KafkaCommonOpt(),

		golibmsg.KafkaConsumerOpt(),
		golibmsg.OnStopConsumerOpt(),

		golibmsg.KafkaProducerOpt(),
		golibmsg.OnStopProducerOpt())
}
