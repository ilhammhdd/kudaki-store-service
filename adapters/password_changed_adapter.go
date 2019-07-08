package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type PasswordChanged struct{}

func (pr *PasswordChanged) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.PasswordChanged
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
