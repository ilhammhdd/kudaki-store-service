package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type AddStorefrontItem struct {
	DBO DBOperator
}

func (asi *AddStorefrontItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := asi.initInOutEvent(in)

	usr := GetUserFromKudakiToken(inEvent.KudakiToken)
	newItem := asi.initItem(inEvent)
	if storefront, ok := StorefrontExists(usr, asi.DBO); ok {
		storefront.TotalItem = storefront.TotalItem + newItem.Amount
		outEvent.Storefront = storefront
	}

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Item = newItem
	outEvent.Requester = usr
	return outEvent
}

func (asi *AddStorefrontItem) initInOutEvent(in proto.Message) (inEvent *events.AddStorefrontItem, outEvent *events.StorefrontItemAdded) {
	inEvent = in.(*events.AddStorefrontItem)

	outEvent = new(events.StorefrontItemAdded)
	outEvent.AddStorefrontItem = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}

func (asi *AddStorefrontItem) initItem(inEvent *events.AddStorefrontItem) *store.Item {
	itemDimension := new(store.ItemDimension)
	itemDimension.Height = inEvent.Height
	itemDimension.Length = inEvent.Length
	itemDimension.UnitOfMeasurement = inEvent.UnitOfMeasurement
	itemDimension.Width = inEvent.Width

	item := new(store.Item)
	item.Amount = inEvent.Amount
	item.Color = inEvent.Color
	item.Description = inEvent.Description
	item.Name = inEvent.Name
	item.Photo = inEvent.Photo
	item.Price = inEvent.Price
	item.PriceDuration = inEvent.PriceDuration
	item.Rating = 0.0
	item.Unit = inEvent.Unit
	item.Uuid = uuid.New().String()

	item.ItemDimension = itemDimension

	return item
}
