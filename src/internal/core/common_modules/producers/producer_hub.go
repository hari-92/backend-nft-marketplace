package common_producers

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib/pubsub"
	"github.com/golibs-starter/golib/web/event"
	kafka "gitlab.com/hari-92/nft-market-server/internal/core/adapter"
	"gitlab.com/hari-92/nft-market-server/internal/core/base"
	"sync"
)

var instanceProducerHub IProducerHub
var once sync.Once

type TestPublishEvent struct {
	UserID   uint32 `json:"user_id"`
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

func (t TestPublishEvent) Name() string {
	return "TestPublishEvent"
}

func (t TestPublishEvent) Data() interface{} {
	return &t
}

type IProducerHub interface {
	TestPublishEvent(event *TestPublishEvent)
}

func NewCommonProducerHub() {
	once.Do(func() {
		instanceProducerHub = &ProducerHub{}
	})
}

type ProducerHub struct {
}

func GetProducerHubInstance() IProducerHub {
	if instanceProducerHub == nil {
		panic("producer hub instance is not initialized")
	}
	return instanceProducerHub
}

func publishEvent[T base.IKafkaEvent](ctx context.Context, msg T) {
	pubsub.Publish(kafka.Event[T]{
		AbstractEvent: event.NewAbstractEvent(ctx, msg.Name()),
		PayloadData:   msg,
	})
}

func (p *ProducerHub) TestPublishEvent(event *TestPublishEvent) {
	fmt.Println("Publishing TestPublishEvent event:", event)
	publishEvent(context.Background(), event)
}
