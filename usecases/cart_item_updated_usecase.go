package usecases

import (
	"database/sql"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
)

type CartItemUpdated struct {
	DBO DBOperator
}

func (ciu *CartItemUpdated) Handle(in proto.Message) *UsecaseHandlerResponse {
	// inEvent := in.(*events.CartItemUpdated)

	// existedItem := ciu.itemExists(inEvent.InitialCartItem.Item.Uuid)
	// if existedItem == nil {
	// 	return &UsecaseHandlerResponse{
	// 		Errs: []string{"item with the given uuid not found"},
	// 		Ok:   false}
	// }

	// existedStorefront := ciu.storefrontExists(existedItem.Storefront.Uuid)
	// if existedStorefront == nil {
	// 	return &UsecaseHandlerResponse{
	// 		Errs: []string{"storefront not found"},
	// 		Ok:   false}
	// }

	// inEvent.UpdatedCartItem.Item = existedItem
	// inEvent.UpdatedCartItem.Item.Storefront = existedStorefront
	// if inEvent.UpdatedCartItem.TotalAmount > inEvent.InitialCartItem.TotalAmount {
	// 	inEvent.UpdatedCartItem.Item.Amount -= int32(inEvent.UpdatedCartItem.TotalAmount - inEvent.InitialCartItem.TotalAmount)
	// 	inEvent.UpdatedCartItem.Item.Storefront.TotalItem -= int32(inEvent.UpdatedCartItem.TotalAmount - inEvent.InitialCartItem.TotalAmount)
	// } else if inEvent.UpdatedCartItem.TotalAmount < inEvent.InitialCartItem.TotalAmount {
	// 	inEvent.UpdatedCartItem.Item.Amount += int32(inEvent.InitialCartItem.TotalAmount - inEvent.UpdatedCartItem.TotalAmount)
	// 	inEvent.UpdatedCartItem.Item.Storefront.TotalItem += int32(inEvent.InitialCartItem.TotalAmount - inEvent.UpdatedCartItem.TotalAmount)
	// }

	// return &UsecaseHandlerResponse{Ok: true}
	return nil
}

func (ciu *CartItemUpdated) itemExists(itemUUID string) *store.Item {
	row, err := ciu.DBO.QueryRow("SELECT storefront_uuid,name,amount,unit,price,description,photo,rating FROM kudaki_store.items WHERE uuid=?;", itemUUID)
	errorkit.ErrorHandled(err)

	var item store.Item
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
		return nil
	}

	item.Uuid = itemUUID
	return &item
}

func (ciu *CartItemUpdated) storefrontExists(storefrontUUID string) *store.Storefront {
	// row, err := ciu.DBO.QueryRow("SELECT user_uuid,total_item,rating FROM storefronts WHERE uuid=?;", storefrontUUID)
	// errorkit.ErrorHandled(err)

	// var storefront store.Storefront
	// storefront.User = new(user.User)
	// if row.Scan(
	// 	&storefront.User.Uuid,
	// 	&storefront.TotalItem,
	// 	&storefront.Rating) == sql.ErrNoRows {
	// 	return nil
	// }
	// storefront.Uuid = storefrontUUID
	// return &storefront

	return nil
}
