package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
)

type CartItemDeleted struct{}

func (cid *CartItemDeleted) Work() interface{} {
	// adapter := &adapters.CartItemDeleted{
	// 	Sanitizer: new(kudakiredisearch.RedisearchText)}
	// usecase := &usecases.CartItemDeleted{
	// 	DBO: mysql.NewDBOperation()}

	// edde := EventDrivenDownstreamExternal{
	// 	PostUsecaseExecutor: cid,
	// 	eventDrivenAdapter:  adapter,
	// 	eventDrivenUsecase:  usecase,
	// 	eventName:           events.RentalTopic_CART_ITEM_DELETED.String(),
	// 	inTopics:            []string{events.RentalTopic_CART_ITEM_DELETED.String()}}

	// edde.handle()
	return nil
}

func (cid *CartItemDeleted) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	// if !usecaseRes.Ok {
	// 	return
	// }

	// in := inEvent.(*events.CartItemDeleted)

	// cid.updateItem(in.CartItem.Item)
	// cid.updateStorefront(in.CartItem.Item.Storefront)
}

func (cid *CartItemDeleted) updateItem(item *store.Item) {
	// dbo := mysql.NewDBOperation()
	// _, err := dbo.Command("UPDATE items SET amount=? WHERE uuid=?;", item.Amount, item.Uuid)
	// errorkit.ErrorHandled(err)

}

func (cid *CartItemDeleted) updateStorefront(storefront *store.Storefront) {
	// dbo := mysql.NewDBOperation()
	// _, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?;", storefront.TotalItem, storefront.Uuid)
	// errorkit.ErrorHandled(err)

}
