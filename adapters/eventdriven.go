package adapters

import (
	"log"

	"github.com/golang/protobuf/proto"
)

type Sanitizer interface {
	Set(string)
	Sanitize() string
	UnSanitize() string
}

type ConsumerLog struct {
	EventName string
}

func (cl *ConsumerLog) Log(partition int32, offset int64, key string) {
	log.Printf("consumed %s : partition = %d, offset = %d, key = %s", cl.EventName, partition, offset, key)
}

type EventDrivenAdapter interface {
	ParseIn(msg []byte) (proto.Message, bool)
	ParseOut(out proto.Message) (key string, message []byte)
}

type EventDrivenDownstreamAdapter interface {
	ParseIn(msg []byte) (proto.Message, bool)
}
