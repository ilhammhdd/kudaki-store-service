package adapters

import (
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"

	"github.com/golang/protobuf/proto"
)

type PasswordReseted struct{}

func (pr *PasswordReseted) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.PasswordReseted
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
