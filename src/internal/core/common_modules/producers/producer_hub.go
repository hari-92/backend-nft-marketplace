package common_producers

import (
	"context"
	"gitlab.com/hari-92/nft-market-server/internal/core/adapter/kafka"
	"sync"

	"github.com/golibs-starter/golib/pubsub"
	"github.com/golibs-starter/golib/web/event"
	"gitlab.com/hari-92/nft-market-server/internal/core/base"
	commonProducersEventGateway "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers/events/gateway"
	commonProducersEvent "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers/events/token"
)

var instanceProducerHub IProducerHub
var once sync.Once

type IProducerHub interface {
	TestPublishEvent(event *commonProducersEventGateway.TestPublishEvent)
	TrackUserLoginEvent(event *commonProducersEventGateway.TrackUserLoginEvent)
	UserCreatedEvent(event *commonProducersEventGateway.UserCreatedEvent)

	// Token event
	InitTokenEvent(event *commonProducersEvent.InitTokenEvent)
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
	ev := kafka.Event[T]{
		AbstractEvent: event.NewAbstractEvent(ctx, msg.Name()),
		PayloadData:   msg,
	}
	pubsub.Publish(ev)
}

func (p *ProducerHub) TestPublishEvent(event *commonProducersEventGateway.TestPublishEvent) {
	ctx := context.Background()
	publishEvent(ctx, event)
}

func (p *ProducerHub) TrackUserLoginEvent(event *commonProducersEventGateway.TrackUserLoginEvent) {
	ctx := context.Background()
	publishEvent(ctx, event)
}

func (p *ProducerHub) UserCreatedEvent(event *commonProducersEventGateway.UserCreatedEvent) {
	ctx := context.Background()
	publishEvent(ctx, event)
}

func (p *ProducerHub) InitTokenEvent(event *commonProducersEvent.InitTokenEvent) {
	publishEvent(context.Background(), event)
}
