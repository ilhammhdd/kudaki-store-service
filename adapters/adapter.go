package adapters

import (
	"github.com/golang/protobuf/proto"
)

type EventDrivenAdapter struct {
	Partition int32
	Offset    int64
	Key       string
	Message   []byte
	In        *proto.Message
	Out       *proto.Message
}

func (eda EventDrivenAdapter) UnmarshalRawInEvent() {

}

func (eda EventDrivenAdapter) CallUsecase() {

}

func (eda EventDrivenAdapter) MarshalOutEvent() (key string, msg []byte) {

	return "", nil
}
