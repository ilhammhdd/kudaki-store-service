package usecases

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	DateTimeFormat = "%d-%02d-%02d %02d:%02d:%02d"
)

func TimeNowToDateTime() string {
	now := time.Now()
	return fmt.Sprintf(DateTimeFormat, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

type EventDrivenUsecase interface {
	Handle(in proto.Message) (out proto.Message)
}
