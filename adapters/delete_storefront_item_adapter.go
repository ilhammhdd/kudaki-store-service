package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type DeleteStorefrontItem struct{}

func (dsi *DeleteStorefrontItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.DeleteStorefrontItem

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
