package usecases

import (
	"database/sql"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
)

type CartItemDeleted struct {
	DBO DBOperator
}

func (cid *CartItemDeleted) Handle(in proto.Message) *UsecaseHandlerResponse {
	// inEvent := in.(*events.CartItemDeleted)

	// if !cid.itemExists(inEvent.CartItem.Item) {
	// 	return &UsecaseHandlerResponse{
	// 		Errs: []string{"item with the given uuid not found"},
	// 		Ok:   false}
	// }

	// if !cid.storefrontExists(inEvent.CartItem.Item.Storefront) {
	// 	return &UsecaseHandlerResponse{
	// 		Errs: []string{"storefront with the given uuid not found"},
	// 		Ok:   false}
	// }

	// inEvent.CartItem.Item.Amount += int32(inEvent.CartItem.TotalAmount)
	// inEvent.CartItem.Item.Storefront.TotalItem += int32(inEvent.CartItem.TotalAmount)

	// return &UsecaseHandlerResponse{Ok: true}
	return nil
}

func (cid *CartItemDeleted) itemExists(item *store.Item) bool {
	row, err := cid.DBO.QueryRow("SELECT storefront_uuid,name,amount,unit,price,description,photo,rating FROM kudaki_store.items WHERE uuid=?;", item.Uuid)
	errorkit.ErrorHandled(err)

	item.Storefront = new(store.Storefront)
	if row.Scan(
		&item.Storefront.Uuid,
		&item.Name,
		&item.Amount,
		&item.Unit,
		&item.Price,
		&item.Description,
		&item.Photo,
		&item.Rating) == sql.ErrNoRows {
		return false
	}

	return true
}

func (cid *CartItemDeleted) storefrontExists(storefront *store.Storefront) bool {
	// row, err := cid.DBO.QueryRow("SELECT user_uuid,total_item,rating FROM storefronts WHERE uuid=?;", storefront.Uuid)
	// errorkit.ErrorHandled(err)

	// storefront.User = new(user.User)
	// if row.Scan(
	// 	&storefront.User.Uuid,
	// 	&storefront.TotalItem,
	// 	&storefront.Rating) == sql.ErrNoRows {
	// 	return false
	// }
	// return true

	return false
}
