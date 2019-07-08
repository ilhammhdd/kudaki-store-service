package externals

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
)

type DeleteStorefrontItem struct{}

func (dsi *DeleteStorefrontItem) Work() interface{} {
	usecase := &usecases.DeleteStorefrontItem{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: dsi,
		eventDrivenAdapter:  new(adapters.DeleteStorefrontItem),
		eventDrivenUsecase:  usecase,
		eventName:           events.StorefrontServiceCommandTopic_DELETE_STOREFRONT_ITEM.String(),
		inTopics:            []string{events.StorefrontServiceCommandTopic_DELETE_STOREFRONT_ITEM.String()},
		outTopic:            events.StorefrontServiceEventTopic_STOREFRONT_ITEM_DELETED.String()}

	ede.handle()
	return nil
}

func (dsi *DeleteStorefrontItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.StorefrontItemDeleted)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dsi.deleteItemFromDB(out.Item)
	dsi.updateStorefront(out.Item)
}

func (dsi *DeleteStorefrontItem) deleteItemFromDB(item *store.Item) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("DELETE FROM items WHERE uuid=? AND storefront_uuid=?;", item.Uuid, item.Storefront.Uuid)
	errorkit.ErrorHandled(err)
}

func (dsi *DeleteStorefrontItem) updateStorefront(item *store.Item) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?;", item.Storefront.TotalItem, item.Storefront.Uuid)
	errorkit.ErrorHandled(err)
}
