package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type UserVerificationEmailSent struct{}

func (su *UserVerificationEmailSent) Work() interface{} {
	usecase := usecases.UserVerificationEmailSent{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: su,
		eventDrivenAdapter:  new(adapters.UserVerificationEmailSent),
		eventDrivenUsecase:  &usecase,
		eventName:           events.UserAuthServiceEventTopic_USER_VERIFICATION_EMAIL_SENT.String(),
		inTopics:            []string{events.UserAuthServiceEventTopic_USER_VERIFICATION_EMAIL_SENT.String()}}

	edde.handle()
	return nil
}

func (su *UserVerificationEmailSent) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	initStorefront := usecaseRes.Data.(*store.Storefront)

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_store.storefronts(uuid,user_uuid,total_item,rating,created_at) VALUES(?,?,?,?,UNIX_TIMESTAMP());",
		initStorefront.Uuid, initStorefront.UserUuid, initStorefront.TotalItem, initStorefront.Rating)
	errorkit.ErrorHandled(err)
}
