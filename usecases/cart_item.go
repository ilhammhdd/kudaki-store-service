package usecases

import (
	"database/sql"

	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/store"
)

type CartItemAdded struct {
	DBO DBOperator
}

func (cia *CartItemAdded) Handle(in proto.Message) *UsecaseHandlerStatus {
	inEvent := in.(*events.CartItemAdded)

	stat := new(UsecaseHandlerStatus)

	if ok := cia.itemExists(inEvent.CartItem.Item); !ok {
		stat.Errs = []string{"item not found"}
		stat.Ok = false
		return stat
	}

	if ok := cia.storefrontExists(inEvent.CartItem.Item.Storefront); !ok {
		stat.Errs = []string{"storefront not found"}
		stat.Ok = false
		return stat
	}
	(*inEvent.CartItem.Item).Amount -= int32(inEvent.AddCartItemRequested.ItemAmount)
	(*inEvent.CartItem.Item.Storefront).TotalItem -= int32(inEvent.AddCartItemRequested.ItemAmount)

	stat.Ok = true
	return stat
}

func (cia *CartItemAdded) itemExists(item *store.Item) bool {
	row, err := cia.DBO.QueryRow("SELECT id FROM items WHERE uuid=?;", item.Uuid)
	errorkit.ErrorHandled(err)

	var itemID uint64
	if row.Scan(&itemID) != sql.ErrNoRows {
		return true
	}
	return false
}

func (cia *CartItemAdded) storefrontExists(storefront *store.Storefront) bool {
	row, err := cia.DBO.QueryRow("SELECT id FROM storefronts WHERE uuid=?;", storefront.Uuid)
	errorkit.ErrorHandled(err)

	var storefrontID uint64
	if row.Scan(&storefrontID) != sql.ErrNoRows {
		return true
	}
	return false
}

type CartItemDeleted struct {
	DBO DBOperator
}

func (cid *CartItemDeleted) Handle(in proto.Message) *UsecaseHandlerStatus {
	inEvent := in.(*events.CartItemDeleted)

	if !cid.itemExists(inEvent.CartItem.Item) {
		return &UsecaseHandlerStatus{
			Errs: []string{"item with the given uuid not found"},
			Ok:   false}
	}

	if !cid.storefrontExists(inEvent.CartItem.Item.Storefront) {
		return &UsecaseHandlerStatus{
			Errs: []string{"storefront with the given uuid not found"},
			Ok:   false}
	}

	inEvent.CartItem.Item.Amount += int32(inEvent.CartItem.TotalAmount)
	inEvent.CartItem.Item.Storefront.TotalItem += int32(inEvent.CartItem.TotalAmount)

	return &UsecaseHandlerStatus{Ok: true}
}

func (cid *CartItemDeleted) itemExists(item *store.Item) bool {
	row, err := cid.DBO.QueryRow("SELECT storefront_uuid,name,amount,unit,price,description,photo,rating FROM items WHERE uuid=?;", item.Uuid)
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
	row, err := cid.DBO.QueryRow("SELECT user_uuid,total_item,rating FROM storefronts WHERE uuid=?;", storefront.Uuid)
	errorkit.ErrorHandled(err)

	storefront.User = new(user.User)
	if row.Scan(
		&storefront.User.Uuid,
		&storefront.TotalItem,
		&storefront.Rating) == sql.ErrNoRows {
		return false
	}
	return true
}

type CartItemUpdated struct {
	DBO DBOperator
}

func (ciu *CartItemUpdated) Handle(in proto.Message) *UsecaseHandlerStatus {
	inEvent := in.(*events.CartItemUpdated)

	existedItem := ciu.itemExists(inEvent.InitialCartItem.Item.Uuid)
	if existedItem == nil {
		return &UsecaseHandlerStatus{
			Errs: []string{"item with the given uuid not found"},
			Ok:   false}
	}

	existedStorefront := ciu.storefrontExists(existedItem.Storefront.Uuid)
	if existedStorefront == nil {
		return &UsecaseHandlerStatus{
			Errs: []string{"storefront not found"},
			Ok:   false}
	}

	inEvent.UpdatedCartItem.Item = existedItem
	inEvent.UpdatedCartItem.Item.Storefront = existedStorefront
	if inEvent.UpdatedCartItem.TotalAmount > inEvent.InitialCartItem.TotalAmount {
		inEvent.UpdatedCartItem.Item.Amount -= int32(inEvent.UpdatedCartItem.TotalAmount - inEvent.InitialCartItem.TotalAmount)
		inEvent.UpdatedCartItem.Item.Storefront.TotalItem -= int32(inEvent.UpdatedCartItem.TotalAmount - inEvent.InitialCartItem.TotalAmount)
	} else if inEvent.UpdatedCartItem.TotalAmount < inEvent.InitialCartItem.TotalAmount {
		inEvent.UpdatedCartItem.Item.Amount += int32(inEvent.InitialCartItem.TotalAmount - inEvent.UpdatedCartItem.TotalAmount)
		inEvent.UpdatedCartItem.Item.Storefront.TotalItem += int32(inEvent.InitialCartItem.TotalAmount - inEvent.UpdatedCartItem.TotalAmount)
	}

	return &UsecaseHandlerStatus{Ok: true}
}

func (ciu *CartItemUpdated) itemExists(itemUUID string) *store.Item {
	row, err := ciu.DBO.QueryRow("SELECT storefront_uuid,name,amount,unit,price,description,photo,rating FROM items WHERE uuid=?;", itemUUID)
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
	row, err := ciu.DBO.QueryRow("SELECT user_uuid,total_item,rating FROM storefronts WHERE uuid=?;", storefrontUUID)
	errorkit.ErrorHandled(err)

	var storefront store.Storefront
	storefront.User = new(user.User)
	if row.Scan(
		&storefront.User.Uuid,
		&storefront.TotalItem,
		&storefront.Rating) == sql.ErrNoRows {
		return nil
	}
	storefront.Uuid = storefrontUUID
	return &storefront
}
