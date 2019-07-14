package usecases

import (
	"database/sql"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
)

type CartItemAdded struct {
	DBO DBOperator
}

func (cia *CartItemAdded) Handle(in proto.Message) *UsecaseHandlerResponse {
	// inEvent := in.(*events.CartItemAdded)

	// stat := new(UsecaseHandlerResponse)

	// if ok := cia.itemExists(inEvent.CartItem.Item); !ok {
	// 	stat.Errs = []string{"item not found"}
	// 	stat.Ok = false
	// 	return stat
	// }

	// if ok := cia.storefrontExists(inEvent.CartItem.Item.Storefront); !ok {
	// 	stat.Errs = []string{"storefront not found"}
	// 	stat.Ok = false
	// 	return stat
	// }
	// (*inEvent.CartItem.Item).Amount -= int32(inEvent.AddCartItem.ItemAmount)
	// (*inEvent.CartItem.Item.Storefront).TotalItem -= int32(inEvent.AddCartItem.ItemAmount)

	// stat.Ok = true
	// return stat

	return nil
}

func (cia *CartItemAdded) itemExists(item *store.Item) bool {
	row, err := cia.DBO.QueryRow("SELECT id FROM kudaki_store.items WHERE uuid=?;", item.Uuid)
	errorkit.ErrorHandled(err)

	var itemID uint64
	if row.Scan(&itemID) != sql.ErrNoRows {
		return true
	}
	return false
}

func (cia *CartItemAdded) storefrontExists(storefront *store.Storefront) bool {
	row, err := cia.DBO.QueryRow("SELECT id FROM kudaki_store.storefronts WHERE uuid=?;", storefront.Uuid)
	errorkit.ErrorHandled(err)

	var storefrontID uint64
	if row.Scan(&storefrontID) != sql.ErrNoRows {
		return true
	}
	return false
}
