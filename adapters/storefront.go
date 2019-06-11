package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type AddStorefrontItem struct{}

func (asi *AddStorefrontItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.AddStorefrontItemRequested

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (asi *AddStorefrontItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.StorefrontItemAdded)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type UpdateStorefrontItem struct{}

func (usi *UpdateStorefrontItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.UpdateStorefrontItemRequested

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (usi *UpdateStorefrontItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.StorefrontItemUpdated)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type DeleteStorefrontItem struct{}

func (dsi *DeleteStorefrontItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.DeleteStorefrontItemRequested

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (dsi *DeleteStorefrontItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.StorefrontItemDeleted)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
