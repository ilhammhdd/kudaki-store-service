package adapters

import (
	"github.com/golang/protobuf/proto"
)

type CartItemUpdated struct {
	Sanitizer Sanitizer
}

func (ciu *CartItemUpdated) ParseIn(msg []byte) (proto.Message, bool) {

	return nil, false
}
