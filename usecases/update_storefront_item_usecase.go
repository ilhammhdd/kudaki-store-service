package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type UpdateStorefrontItem struct {
	DBO DBOperator
}

func (usi *UpdateStorefrontItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := usi.initInOutEvent(in)

	usr := GetUserFromKudakiToken(inEvent.KudakiToken)
	existedStorefront, ok := StorefrontExists(usr, usi.DBO)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"storefront not exists"}
		return outEvent
	}

	intendedItem, ok := IntendedItemExists(usi.DBO, existedStorefront, inEvent.ItemUuid)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"the intended item for update doesn't exists"}
		return outEvent
	}

	outEvent.ItemsBefore = append(outEvent.ItemsBefore, intendedItem)
	outEvent.ItemsBefore[0].Storefront = existedStorefront

	outEvent.ItemsAfter = append(outEvent.ItemsAfter, usi.initUpdatedItem(inEvent, intendedItem, existedStorefront))
	outEvent.ItemsAfter[0].Storefront = existedStorefront
	outEvent.ItemsAfter[0].Storefront = usi.addOrSubtractTotalItem(inEvent.Amount, intendedItem.Amount, outEvent.ItemsBefore[0].Storefront)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Requester = usr

	return outEvent

	return nil
}

func (usi *UpdateStorefrontItem) initInOutEvent(in proto.Message) (inEvent *events.UpdateStorefrontItem, outEvent *events.StorefrontItemsUpdated) {
	inEvent = in.(*events.UpdateStorefrontItem)

	outEvent = new(events.StorefrontItemsUpdated)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.UpdateStorefrontItem = inEvent

	return
}

func (usi *UpdateStorefrontItem) initUpdatedItem(inEvent *events.UpdateStorefrontItem, intendedItem *store.Item, existedStorefront *store.Storefront) *store.Item {
	updatedItemDimension := new(store.ItemDimension)
	updatedItemDimension.Length = inEvent.Length
	updatedItemDimension.Width = inEvent.Width
	updatedItemDimension.Height = inEvent.Height
	updatedItemDimension.UnitOfMeasurement = inEvent.UnitOfMeasurement

	updatedItem := new(store.Item)
	updatedItem.Amount = inEvent.Amount
	updatedItem.Color = inEvent.Color
	updatedItem.Description = inEvent.Description
	updatedItem.Name = inEvent.Name
	updatedItem.Photo = inEvent.Photo
	updatedItem.Price = inEvent.Price
	updatedItem.Rating = intendedItem.Rating
	updatedItem.Storefront = existedStorefront
	updatedItem.Unit = inEvent.Unit
	updatedItem.Uuid = intendedItem.Uuid

	updatedItem.PriceDuration = inEvent.PriceDuration
	updatedItem.ItemDimension = updatedItemDimension

	return updatedItem
}

func (usi *UpdateStorefrontItem) addOrSubtractTotalItem(newAmount int32, oldAmount int32, storefrontBefore *store.Storefront) (storefrontAfter *store.Storefront) {
	storefrontAfterTemp := *storefrontBefore

	if newAmount > oldAmount {
		diff := newAmount - oldAmount
		storefrontAfterTemp.TotalItem += diff
	} else if newAmount < oldAmount {
		diff := oldAmount - newAmount
		storefrontAfterTemp.TotalItem -= diff
	}
	return &storefrontAfterTemp
}
