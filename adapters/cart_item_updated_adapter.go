package adapters

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type CartItemUpdated struct {
	Sanitizer Sanitizer
}

func (ciu *CartItemUpdated) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.CartItemUpdated

	if proto.Unmarshal(msg, &inEvent) == nil {
		if inEvent.EventStatus.HttpCode == http.StatusOK {
			ciu.Sanitizer.Set(inEvent.InitialCartItem.ItemUuid)
			return &inEvent, true
		}
	}
	return nil, false
}
