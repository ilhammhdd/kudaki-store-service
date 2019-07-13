package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type UpdateStorefrontItem struct{}

func (usi *UpdateStorefrontItem) Work() interface{} {
	usecase := &usecases.UpdateStorefrontItem{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: usi,
		eventDrivenAdapter:  new(adapters.UpdateStorefrontItem),
		eventDrivenUsecase:  usecase,
		eventName:           events.StorefrontServiceCommandTopic_UPDATE_STOREFRONT_ITEM.String(),
		inTopics:            []string{events.StorefrontServiceCommandTopic_UPDATE_STOREFRONT_ITEM.String()},
		outTopic:            events.StorefrontServiceEventTopic_STOREFRONT_ITEMS_UPDATED.String()}

	ede.handle()
	return nil
}

func (usi *UpdateStorefrontItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.StorefrontItemsUpdated)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	usi.updateStorefront(out.ItemsAfter[0].Storefront)
	usi.updateItem(out.ItemsAfter[0])
}

func (usi *UpdateStorefrontItem) updateStorefront(updatedStorefront *store.Storefront) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?;", updatedStorefront.TotalItem, updatedStorefront.Uuid)
	errorkit.ErrorHandled(err)
}

func (usi *UpdateStorefrontItem) updateItem(updatedItem *store.Item) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE items SET name=?,amount=?,unit=?,price=?,description=?,photo=? WHERE uuid=?;",
		updatedItem.Name,
		updatedItem.Amount,
		updatedItem.Unit,
		updatedItem.Price,
		updatedItem.Description,
		updatedItem.Photo,
		updatedItem.Uuid)
	errorkit.ErrorHandled(err)
}
