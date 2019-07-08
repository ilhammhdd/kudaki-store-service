package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type AddStorefrontItem struct{}

func (asi *AddStorefrontItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.AddStorefrontItem

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
