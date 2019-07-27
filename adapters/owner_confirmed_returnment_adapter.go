package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type OwnerConfirmedReturnment struct{}

func (ocr *OwnerConfirmedReturnment) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.OwnerConfirmedReturnment
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
