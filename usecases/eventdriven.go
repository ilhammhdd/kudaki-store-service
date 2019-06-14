package usecases

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	DateTimeFormat = "%d-%02d-%02d %02d:%02d:%02d"
)

type UsecaseHandlerStatus struct {
	Ok   bool
	Errs []string
}

func TimeNowToDateTime() string {
	now := time.Now()
	return fmt.Sprintf(DateTimeFormat, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

type EventDrivenUsecase interface {
	Handle(in proto.Message) (out proto.Message)
}

type EventDrivenDownstreamUsecase interface {
	Handle(in proto.Message) *UsecaseHandlerStatus
}
