package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type OwnerOrderApproved struct{}

func (ooa *OwnerOrderApproved) Work() interface{} {
	usecase := usecases.OwnerOrderApproved{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: ooa,
		eventDrivenAdapter:  new(adapters.OwnerOrderApproved),
		eventDrivenUsecase:  &usecase,
		eventName:           events.OrderServiceEventTopic_OWNER_ORDER_APPROVED.String(),
		inTopics:            []string{events.OrderServiceEventTopic_OWNER_ORDER_APPROVED.String()}}

	edde.handle()
	return nil
}

func (ooa *OwnerOrderApproved) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	if !usecaseRes.Ok {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	items := usecaseRes.Data.([]*usecases.OwnerOrderApprovedUpdateItem)
	for i := 0; i < len(items); i++ {
		dbo.Command("INSERT INTO kudaki_store.storefronts(uuid, total_item) VALUES(?, ?) ON DUPLICATE KEY UPDATE total_item = ?;",
			items[i].StorefrontUuid, items[i].StorefrontTotalItem, items[i].StorefrontTotalItem)
	}

	dboItems := mysql.NewDBOperation(mysql.CommandDB)
	for j := 0; j < len(items); j++ {
		dboItems.Command("INSERT INTO kudaki_store.items(uuid, amount) VALUES(?, ?) ON DUPLICATE KEY UPDATE amount = ?;",
			items[j].ItemUuid, items[j].ItemAmount, items[j].ItemAmount)
	}
}
