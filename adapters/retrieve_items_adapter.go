package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type RetrieveItems struct{}

func (ri RetrieveItems) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveItems

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ri RetrieveItems) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ItemsRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
