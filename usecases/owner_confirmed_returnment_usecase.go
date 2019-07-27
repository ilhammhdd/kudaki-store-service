package usecases

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type OwnerConfirmedReturnment struct {
	DBO DBOperator
}

func (ocr *OwnerConfirmedReturnment) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.OwnerConfirmedReturnment)

	items := ocr.retrieveItemsToBeUpdated(inEvent)
	ocr.addItemQuantity(items)

	return &UsecaseHandlerResponse{
		Data: items,
		Ok:   true}
}

func (ocr *OwnerConfirmedReturnment) retrieveItemsToBeUpdated(inEvent *events.OwnerConfirmedReturnment) []*OwnerConfirmedReturnmentUpdateItem {
	rows, err := ocr.DBO.Query("SELECT ci.total_item AS ci_total_item, i.uuid AS i_uuid, i.amount AS i_amount, sf.uuid AS sf_uuid, sf.total_item AS sf_total_item FROM kudaki_order.owner_orders oo JOIN kudaki_order.orders o ON oo.order_uuid = o.uuid JOIN kudaki_rental.carts c ON o.cart_uuid = o.cart_uuid JOIN kudaki_rental.cart_items ci ON c.uuid = ci.cart_uuid JOIN kudaki_store.items i ON ci.item_uuid = i.uuid JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid WHERE oo.uuid = ?;",
		inEvent.OwnerOrder.Uuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var ocruis []*OwnerConfirmedReturnmentUpdateItem
	for rows.Next() {
		var ocrui OwnerConfirmedReturnmentUpdateItem
		err = rows.Scan(
			&ocrui.CartItemTotalItem,
			&ocrui.ItemUuid,
			&ocrui.ItemAmount,
			&ocrui.StorefrontUuid,
			&ocrui.StorefrontTotalItem)
		errorkit.ErrorHandled(err)

		ocruis = append(ocruis, &ocrui)
	}

	return ocruis
}

type OwnerConfirmedReturnmentUpdateItem struct {
	CartItemTotalItem   int32
	ItemUuid            string
	ItemAmount          int32
	StorefrontUuid      string
	StorefrontTotalItem int32
}

func (ocr *OwnerConfirmedReturnment) addItemQuantity(items []*OwnerConfirmedReturnmentUpdateItem) {
	for i := 0; i < len(items); i++ {
		(*items[i]).ItemAmount += items[i].CartItemTotalItem
		(*items[i]).StorefrontTotalItem += items[i].CartItemTotalItem
	}
}
