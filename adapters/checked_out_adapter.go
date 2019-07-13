package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type CheckedOut struct{}

func (co *CheckedOut) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.CheckedOut
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (co *CheckedOut) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.StorefrontItemsUpdated)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
