package wallet_consumers

import (
	"github.com/golibs-starter/golib-message-bus/kafka/core"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/event"
	jsoniter "github.com/json-iterator/go"
)

type UserCreated struct {
	UserID uint32 `json:"user_id"`
}

type UserCreatedEvent struct {
	*event.AbstractEvent
	PayloadData *UserCreated `json:"payload"`
}

func (u UserCreatedEvent) Payload() interface{} {
	return u.PayloadData
}
func (u UserCreatedEvent) String() string {
	return u.ToString(u)
}

type UserCreatedConsumer struct {
}

func (u UserCreatedConsumer) HandlerFunc(message *core.ConsumerMessage) {
	var eventData UserCreatedEvent
	if err := jsoniter.Unmarshal(message.Value, &eventData); err != nil {
		log.Errorf("UserCreatedConsumer json unmarshal err: %s", err)
		return
	}

	log.Infof("UserCreatedConsumer event: %v", eventData)
}

func (u UserCreatedConsumer) Close() {}

func NewUserCreatedConsumer() core.ConsumerHandler {
	return &UserCreatedConsumer{}
}
