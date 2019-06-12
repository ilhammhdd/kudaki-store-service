package usecases

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ilhammhdd/kudaki-entities/store"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type AddStorefrontItem struct {
	DBO DBOperator
}

func (asi *AddStorefrontItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := asi.initInOutEvent(in)

	usr := asi.getUserFromKudakiToken(inEvent.KudakiToken)
	newItem := asi.initItem(inEvent)
	if storefront, ok := asi.storefrontExists(usr); ok {
		storefront.TotalItem = storefront.TotalItem + newItem.Amount
		outEvent.Storefront = storefront
		newItem.Storefront = storefront
	}

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Item = newItem
	outEvent.User = usr
	return outEvent
}

func (asi *AddStorefrontItem) initInOutEvent(in proto.Message) (inEvent *events.AddStorefrontItemRequested, outEvent *events.StorefrontItemAdded) {
	inEvent = in.(*events.AddStorefrontItemRequested)

	outEvent = new(events.StorefrontItemAdded)
	outEvent.AddStorefrontItemRequested = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}

func (asi *AddStorefrontItem) getUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.Role(user.Role_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}

func (asi *AddStorefrontItem) storefrontExists(usr *user.User) (*store.Storefront, bool) {
	row, err := asi.DBO.QueryRow("SELECT uuid,total_item,rating FROM storefronts WHERE user_uuid = ?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var storefront store.Storefront
	if row.Scan(&storefront.Uuid, &storefront.TotalItem, &storefront.Rating) == sql.ErrNoRows {
		return nil, false
	}
	storefront.User = usr
	return &storefront, true
}

func (asi *AddStorefrontItem) initItem(inEvent *events.AddStorefrontItemRequested) *store.Item {
	item := new(store.Item)
	item.Amount = inEvent.Amount
	item.Description = inEvent.Description
	item.Name = inEvent.Name
	item.Photo = inEvent.Photo
	item.Price = inEvent.Price
	item.Rating = 0.0
	item.Unit = inEvent.Unit
	item.Uuid = uuid.New().String()

	return item
}

type UpdateStorefrontItem struct {
	DBO DBOperator
}

func (usi *UpdateStorefrontItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := usi.initInOutEvent(in)

	usr := usi.getUserFromKudakiToken(inEvent.KudakiToken)
	existedStorefront, ok := usi.storefrontExists(usr)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"storefront not exists"}
		return outEvent
	}

	intendedItem, ok := usi.intendedItemExists(existedStorefront, inEvent.Uuid)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"the intended item for update doesn't exists"}
		return outEvent
	}

	outEvent.Item = usi.initUpdatedItem(inEvent, intendedItem, existedStorefront)
	outEvent.Storefront = existedStorefront
	outEvent.User = usr
	outEvent.EventStatus.HttpCode = http.StatusOK

	outEvent.Storefront.TotalItem = usi.addOrSubtractTotalItem(inEvent.Amount, intendedItem.Amount, outEvent.Storefront.TotalItem)
	log.Println("storefront state in usecase : ", outEvent.Storefront)

	return outEvent
}

func (usi *UpdateStorefrontItem) initInOutEvent(in proto.Message) (inEvent *events.UpdateStorefrontItemRequested, outEvent *events.StorefrontItemUpdated) {
	inEvent = in.(*events.UpdateStorefrontItemRequested)

	outEvent = new(events.StorefrontItemUpdated)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.UpdateStorefrontItemRequested = inEvent

	return
}

func (usi *UpdateStorefrontItem) getUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.Role(user.Role_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}

func (usi *UpdateStorefrontItem) storefrontExists(usr *user.User) (*store.Storefront, bool) {
	row, err := usi.DBO.QueryRow("SELECT uuid,total_item,rating FROM storefronts WHERE user_uuid=?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var storefront store.Storefront
	if row.Scan(&storefront.Uuid, &storefront.TotalItem, &storefront.Rating) == sql.ErrNoRows {
		return nil, false
	}
	storefront.User = usr

	return &storefront, true
}

func (usi *UpdateStorefrontItem) intendedItemExists(storefront *store.Storefront, itemUUID string) (*store.Item, bool) {
	row, err := usi.DBO.QueryRow("SELECT name,amount,unit,price,description,photo,rating FROM items WHERE storefront_uuid=? AND uuid=?;", storefront.Uuid, itemUUID)
	errorkit.ErrorHandled(err)

	var intendedItem store.Item
	if row.Scan(
		&intendedItem.Name,
		&intendedItem.Amount,
		&intendedItem.Unit,
		&intendedItem.Price,
		&intendedItem.Description,
		&intendedItem.Photo,
		&intendedItem.Rating) == sql.ErrNoRows {
		return nil, false
	}
	intendedItem.Storefront = storefront
	intendedItem.Uuid = itemUUID

	return &intendedItem, true
}

func (usi *UpdateStorefrontItem) initUpdatedItem(inEvent *events.UpdateStorefrontItemRequested, intendedItem *store.Item, existedStorefront *store.Storefront) *store.Item {
	updatedItem := new(store.Item)
	updatedItem.Amount = inEvent.Amount
	updatedItem.Description = inEvent.Description
	updatedItem.Name = inEvent.Name
	updatedItem.Photo = inEvent.Photo
	updatedItem.Price = inEvent.Price
	updatedItem.Rating = intendedItem.Rating
	updatedItem.Storefront = existedStorefront
	updatedItem.Unit = inEvent.Unit
	updatedItem.Uuid = intendedItem.Uuid
	return updatedItem
}

func (usi *UpdateStorefrontItem) addOrSubtractTotalItem(newAmount int32, oldAmount int32, initialTotalItem int32) (finalTotalItem int32) {
	log.Printf("newAmount = %d, oldAmount  = %d", newAmount, oldAmount)
	if newAmount > oldAmount {
		diff := newAmount - oldAmount
		return initialTotalItem + diff
	} else if newAmount < oldAmount {
		diff := oldAmount - newAmount
		return initialTotalItem - diff
	}
	return 0
}

type DeleteStorefrontItem struct {
	DBO DBOperator
}

func (dsi *DeleteStorefrontItem) Handle(in proto.Message) (out proto.Message) {

	return nil
}

func (dsi *DeleteStorefrontItem) initInOutEvent(in proto.Message) (inEvent *events.DeleteStorefrontItemRequested, outEvent *events.StorefrontItemDeleted) {

	return nil, nil
}
