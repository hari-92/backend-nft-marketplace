package wallet_consumers

import (
	"github.com/golibs-starter/golib-message-bus/kafka/core"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/event"
	jsoniter "github.com/json-iterator/go"
	wallet_requests "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/requests"
	wallet_services "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/services"
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
	walletService wallet_services.IWalletService
}

func (u UserCreatedConsumer) HandlerFunc(message *core.ConsumerMessage) {
	var eventData UserCreatedEvent
	if err := jsoniter.Unmarshal(message.Value, &eventData); err != nil {
		log.Errorf("UserCreatedConsumer json unmarshal err: %s", err)
		return
	}
	res, err := u.walletService.CreateWallet(&wallet_requests.CreateWalletRequest{
		UserID: eventData.PayloadData.UserID,
	})
	if err != nil {
		log.Errorf("UserCreatedConsumer CreateWallet error: %s", err)
		return
	}

	log.Info("UserCreatedConsumer eventData and CreateWallet", eventData, res)
}

func (u UserCreatedConsumer) Close() {}

func NewUserCreatedConsumer(
	walletService wallet_services.IWalletService,
) core.ConsumerHandler {
	return &UserCreatedConsumer{
		walletService: walletService,
	}
}
