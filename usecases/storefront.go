package usecases

import (
	"database/sql"
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

	return nil
}

func (usi *UpdateStorefrontItem) initInOutEvent(in proto.Message) (inEvent *events.UpdateStorefrontItemRequested, outEvent *events.StorefrontItemUpdated) {

	return nil, nil
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
