package usecases

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	DateTimeFormat = "%d-%02d-%02d %02d:%02d:%02d"
)

type UsecaseHandlerResponse struct {
	Ok   bool
	Errs []string
	Data interface{}
}

func TimeNowToDateTime() string {
	now := time.Now()
	return fmt.Sprintf(DateTimeFormat, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

type EventDrivenUsecase interface {
	Handle(in proto.Message) (out proto.Message)
}

type EventDrivenDownstreamUsecase interface {
	Handle(in proto.Message) *UsecaseHandlerResponse
}

type ResultSchemer interface {
	SetResultSources(i ...interface{}) ResultSchemer
	ParseToResult() []byte
}
