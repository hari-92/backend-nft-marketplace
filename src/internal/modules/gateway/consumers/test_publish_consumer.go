package gateway_consumers

import (
	"github.com/golibs-starter/golib-message-bus/kafka/core"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/event"
	jsoniter "github.com/json-iterator/go"
)

type TestPublish struct {
	UserID  uint32 `json:"user_id"`
	Message string `json:"message"`
}

type TestPublishEvent struct {
	*event.AbstractEvent
	PayloadData *TestPublish `json:"payload"`
}

func (t TestPublishEvent) Payload() interface{} {
	return t.Payload
}

func (t TestPublishEvent) String() string {
	return t.ToString(t)
}

type TestPublishConsumer struct {
}

func (c TestPublishConsumer) HandlerFunc(message *core.ConsumerMessage) {
	log.Infof("[TestPublishConsumer] HandleFunc called with message: %s", message.String())
	var eventData TestPublishEvent
	if err := jsoniter.Unmarshal(message.Value, &eventData); err != nil {
		log.Errorf("[TestPublishConsumer] json.Unmarshal error: %s", err)
		return
	}
	log.Infof("[TestPublishConsumer] EventData: %s", eventData.String())
}

func (c TestPublishConsumer) Close() {}

func NewTestPublishConsumer() core.ConsumerHandler {
	log.Info("[TestPublishConsumer] NewTestPublishConsumer called")
	return &TestPublishConsumer{}
}
