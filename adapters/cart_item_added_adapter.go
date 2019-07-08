package adapters

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type CartItemAdded struct{}

func (cia *CartItemAdded) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.CartItemAdded

	if proto.Unmarshal(msg, &inEvent) == nil {
		if inEvent.EventStatus.HttpCode == http.StatusOK {
			return &inEvent, true
		}
	}
	return nil, false
}
