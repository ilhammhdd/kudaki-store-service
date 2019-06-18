package adapters

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-entities/events"
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

type CartItemDeleted struct {
	Sanitizer Sanitizer
}

func (cid *CartItemDeleted) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.CartItemDeleted

	if proto.Unmarshal(msg, &inEvent) == nil {
		if inEvent.EventStatus.HttpCode == http.StatusOK {
			return &inEvent, true
		}
	}
	return nil, false
}

type CartItemUpdated struct {
	Sanitizer Sanitizer
}

func (ciu *CartItemUpdated) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.CartItemUpdated

	if proto.Unmarshal(msg, &inEvent) == nil {
		if inEvent.EventStatus.HttpCode == http.StatusOK {
			ciu.Sanitizer.Set(inEvent.InitialCartItem.Item.Uuid)
			inEvent.InitialCartItem.Item.Uuid = ciu.Sanitizer.UnSanitize()
			return &inEvent, true
		}
	}
	return nil, false
}
