package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type DeleteStorefrontItem struct {
	DBO DBOperator
}

func (dsi *DeleteStorefrontItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := dsi.initInOutEvent(in)

	usr := GetUserFromKudakiToken(inEvent.KudakiToken)

	existedStorefront, ok := StorefrontExists(usr, dsi.DBO)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"storefront for the given user not found"}
		return outEvent
	}

	existedItem, ok := IntendedItemExists(dsi.DBO, existedStorefront, inEvent.ItemUuid)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"item with the given uuid doesn't exists"}
		return outEvent
	}
	existedStorefront.TotalItem -= existedItem.Amount

	outEvent.Item = existedItem
	outEvent.Item.Storefront = existedStorefront
	outEvent.Requester = usr
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (dsi *DeleteStorefrontItem) initInOutEvent(in proto.Message) (inEvent *events.DeleteStorefrontItem, outEvent *events.StorefrontItemDeleted) {
	inEvent = in.(*events.DeleteStorefrontItem)

	outEvent = new(events.StorefrontItemDeleted)
	outEvent.DeleteStorefrontItem = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}
