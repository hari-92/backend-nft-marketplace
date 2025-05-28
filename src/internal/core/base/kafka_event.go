package base

type IKafkaEvent interface {
	Name() string
	Data() interface{}
}
