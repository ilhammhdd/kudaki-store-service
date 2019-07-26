package usecases

import (
	"log"

	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/rental"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type CheckedOut struct {
	DBO DBOperator
}

func (co *CheckedOut) Handle(in proto.Message) (out proto.Message) {

	return nil
}

func (co *CheckedOut) initInOutEvent(in proto.Message) (inEvent *events.CheckedOut, outEvent *events.StorefrontItemsUpdated) {
	inEvent = in.(*events.CheckedOut)

	outEvent = new(events.StorefrontItemsUpdated)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = inEvent.Tenant
	outEvent.Uid = uuid.New().String()

	return
}

// ------------------------------------------------------------------------

func (co *CheckedOut) retrieveItems(cartUuid string) []*ItemTempCheckedOut {
	rows, err := co.DBO.Query("SELECT i.id AS item_id, i.uuid AS item_uuid, i.name, i.photo, i.price, i.price_duration, sf.uuid AS storefront_uuid FROM (SELECT ci_i.id FROM kudaki_rental.cart_items ci_i JOIN kudaki_store.items i_i ON ci_i.item_uuid = i_i.uuid JOIN kudaki_store.storefronts sf_i ON i_i.storefront_uuid = sf_i.uuid WHERE ci_i.cart_uuid = ?) ci_ids JOIN kudaki_rental.cart_items ci ON ci.id = ci_ids.id JOIN kudaki_store.items i ON ci.item_uuid = i.uuid JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid;",
		cartUuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var itemTemps []*ItemTempCheckedOut
	for rows.Next() {
		var itemTemp ItemTempCheckedOut
		var priceDuration string

		err = rows.Scan(
			&itemTemp.Id,
			&itemTemp.Uuid,
			&itemTemp.Amount,
			&itemTemp.StorefrontUuid)
		errorkit.ErrorHandled(err)
		itemTemp.Item.PriceDuration = store.PriceDuration(store.PriceDuration_value[priceDuration])

		itemTemps = append(itemTemps, &itemTemp)
	}

	return itemTemps
}

func (co *CheckedOut) retrieveStorefronts(cartUuid string) []*StorefrontTempCheckedOut {
	rows, err := co.DBO.Query("SELECT sf.id, sf.uuid, sf.total_item FROM (SELECT id FROM kudaki_rental.cart_items ci_i WHERE ci_i.cart_uuid = ?) ci_ids JOIN kudaki_rental.cart_items ci ON ci.id = ci_ids.id JOIN kudaki_store.items i ON ci.item_uuid = i.uuid JOIN kudaki_store.storefronts sf ON sf.uuid = i.storefront_uuid JOIN kudaki_user.users u ON sf.user_uuid = u.uuid JOIN kudaki_user.profiles p ON u.uuid = p.user_uuid GROUP BY sf.uuid;",
		cartUuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var storefrontTemps []*StorefrontTempCheckedOut

	for rows.Next() {
		var storefrontTemp StorefrontTempCheckedOut

		err = rows.Scan(
			&storefrontTemp.Id,
			&storefrontTemp.Uuid,
			&storefrontTemp.TotalItem)
		errorkit.ErrorHandled(err)

		storefrontTemps = append(storefrontTemps, &storefrontTemp)
	}

	return storefrontTemps
}

func (rci *CheckedOut) pasteItemToStorefront(storefrontTemps []*StorefrontTempCheckedOut, itemTemps []*ItemTempCheckedOut) {
	for i := 0; i < len(itemTemps); i++ {
		for j := 0; j < len(storefrontTemps); j++ {
			if itemTemps[i].StorefrontUuid == storefrontTemps[j].Uuid {
				log.Println("storefront uuid equals : ", itemTemps[i].StorefrontUuid, " == ", storefrontTemps[j].Uuid)
				(*storefrontTemps[j]).ItemTemps = append((*storefrontTemps[j]).ItemTemps, itemTemps[i])
			}
		}
	}
}

type StorefrontTempCheckedOut struct {
	store.Storefront
	ItemTemps []*ItemTempCheckedOut `json:"cart_items"`
}

type ItemTempCheckedOut struct {
	store.Item
	StorefrontUuid string
}

// ------------------------------------------------

func (co *CheckedOut) retrieveCartItems(cartUuid string) []*CartItemTempCheckedOut {
	rows, err := co.DBO.Query("SELECT ci.id, ci.uuid, ci.total_item, i.uuid AS item_uuid, sf.uuid AS storefront_uuid FROM (SELECT ci_i.id FROM kudaki_rental.cart_items ci_i JOIN kudaki_store.items i_i ON ci_i.item_uuid = i_i.uuid JOIN kudaki_store.storefronts sf_i ON i_i.storefront_uuid = sf_i.uuid WHERE ci_i.cart_uuid = ? LIMIT ?, ?) ci_ids JOIN kudaki_rental.cart_items ci ON ci.id = ci_ids.id JOIN kudaki_store.items i ON ci.item_uuid = i.uuid JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid;",
		cartUuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var cartItemTemps []*CartItemTempCheckedOut
	for rows.Next() {
		var cartItemTemp CartItemTempCheckedOut

		err = rows.Scan(
			&cartItemTemp.Id,
			&cartItemTemp.Uuid,
			&cartItemTemp.TotalItem,
			&cartItemTemp.ItemUuid,
			&cartItemTemp.StorefrontUuid)
		errorkit.ErrorHandled(err)

		cartItemTemps = append(cartItemTemps, &cartItemTemp)
	}

	return cartItemTemps
}

type CartItemTempCheckedOut struct {
	rental.CartItem
	StorefrontUuid string
}

func (co *CheckedOut) subtractItemQuantity(storefronts []*StorefrontTempCheckedOut, cartItems []*CartItemTempCheckedOut) {
	var updatedItems []*store.Item

	for i := 0; i < len(cartItems); i++ {
		for j := 0; j < len(storefronts); j++ {
			updatedStorefrontCheckedOut := *storefronts[j]
			if cartItems[i].StorefrontUuid == storefronts[j].Uuid {
				updatedStorefrontCheckedOut.TotalItem -= cartItems[i].TotalItem
			}
			for k := 0; k < len(storefronts[j].ItemTemps); k++ {
				updatedItemCheckedOut := *storefronts[j].ItemTemps[k]
				updatedItemCheckedOut.Amount -= cartItems[i].TotalItem

				updatedItems = append(updatedItems, &updatedItemCheckedOut.Item)
			}
		}
	}

}
