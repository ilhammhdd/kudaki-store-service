package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type UpdateStorefrontItem struct{}

func (usi *UpdateStorefrontItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.UpdateStorefrontItem

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
