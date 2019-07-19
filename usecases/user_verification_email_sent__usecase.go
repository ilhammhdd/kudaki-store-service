package usecases

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/user"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type UserVerificationEmailSent struct {
	DBO DBOperator
}

func (su *UserVerificationEmailSent) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.UserVerificationEmailSent)

	return &UsecaseHandlerResponse{
		Data: su.initStorefront(inEvent.User),
		Ok:   true,
	}
}

func (su *UserVerificationEmailSent) initStorefront(usr *user.User) *store.Storefront {
	return &store.Storefront{
		CreatedAt:      ptypes.TimestampNow(),
		Rating:         0.0,
		TotalItem:      0,
		TotalRawRating: 0.0,
		UserUuid:       usr.Uuid,
		Uuid:           uuid.New().String()}
}
