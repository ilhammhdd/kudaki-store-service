package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
)

type CartItemUpdated struct{}

func (ciu *CartItemUpdated) Work() interface{} {
	// usecase := &usecases.CartItemUpdated{DBO: mysql.NewDBOperation()}
	// adapter := &adapters.CartItemUpdated{Sanitizer: new(kudakiredisearch.RedisearchText)}

	// edde := EventDrivenDownstreamExternal{
	// 	PostUsecaseExecutor: ciu,
	// 	eventDrivenAdapter:  adapter,
	// 	eventDrivenUsecase:  usecase,
	// 	eventName:           events.RentalTopic_CART_ITEM_UPDATED.String(),
	// 	inTopics:            []string{events.RentalTopic_CART_ITEM_UPDATED.String()}}

	// edde.handle()
	return nil
}

func (ciu *CartItemUpdated) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	// if !usecaseRes.Ok {
	// 	return
	// }

	// in := inEvent.(*events.CartItemUpdated)

	// ciu.updateItem(in.UpdatedCartItem.Item)
	// ciu.updateStorefront(in.UpdatedCartItem.Item.Storefront)
}

func (ciu *CartItemUpdated) updateItem(item *store.Item) {
	// dbo := mysql.NewDBOperation()
	// _, err := dbo.Command("UPDATE items SET amount=? WHERE uuid=?;", item.Amount, item.Uuid)
	// errorkit.ErrorHandled(err)

}

func (ciu *CartItemUpdated) updateStorefront(storefront *store.Storefront) {
	// dbo := mysql.NewDBOperation()
	// _, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?;", storefront.TotalItem, storefront.Uuid)
	// errorkit.ErrorHandled(err)

}
