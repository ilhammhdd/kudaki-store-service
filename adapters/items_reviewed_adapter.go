package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type ItemsReviewed struct{}

func (ir *ItemsReviewed) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.ItemsReviewed
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
