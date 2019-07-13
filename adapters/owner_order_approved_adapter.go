package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type OwnerOrderApproved struct{}

func (ooa *OwnerOrderApproved) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.OwnerOrderApproved
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
