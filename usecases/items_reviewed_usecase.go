package usecases

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type ItemsReviewed struct {
	DBO DBOperator
}

func (ir *ItemsReviewed) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.ItemsReviewed)
	log.Println(inEvent)

	storefrontWithCount := ir.retrieveStorefronts(inEvent)
	ir.recalculateStorefrontRating(storefrontWithCount, inEvent.ItemReviews)

	return &UsecaseHandlerResponse{
		Data: storefrontWithCount,
		Ok:   true}
}

type StorefrontWithCount struct {
	store.Storefront
	CountedTotalReview int32
}

func (ir *ItemsReviewed) retrieveStorefronts(inEvent *events.ItemsReviewed) []*StorefrontWithCount {
	var storefronts []*StorefrontWithCount
	length := len(inEvent.ItemReviews)
	var sfUuids []interface{}
	// ----------------------------------------
	subQuery := "SELECT COUNT(ir_i.id) FROM kudaki_store.item_reviews ir_i JOIN kudaki_store.items i_i ON ir_i.item_uuid = i_i.uuid WHERE i_i.storefront_uuid "
	inOperatorSubQuery := "IN("
	for i := 0; i < length; i++ {
		sfUuids = append(sfUuids, inEvent.ItemReviews[i].Item.Storefront.Uuid)
		if i == length-1 {
			inOperatorSubQuery += "?"
		} else {
			inOperatorSubQuery += "?,"
		}
	}
	inOperatorSubQuery += ")"
	subQuery += inOperatorSubQuery
	// ----------------------------------------
	query := "SELECT (" + subQuery + "),id,uuid,user_uuid,total_item,rating,total_raw_rating,created_at FROM kudaki_store.storefronts WHERE uuid "
	inOperator := "IN("
	for i := 0; i < length; i++ {
		sfUuids = append(sfUuids, inEvent.ItemReviews[i].Item.Storefront.Uuid)
		if i == length-1 {
			inOperator += "?"
		} else {
			inOperator += "?,"
		}
	}
	inOperator += ")"
	query += inOperator
	log.Println("query : ", query)

	rows, err := ir.DBO.Query(query, sfUuids...)
	errorkit.ErrorHandled(err)
	defer rows.Close()
	var createdAt int64

	for rows.Next() {
		var storefront StorefrontWithCount
		rows.Scan(
			&storefront.CountedTotalReview,
			&storefront.Id,
			&storefront.Uuid,
			&storefront.UserUuid,
			&storefront.TotalItem,
			&storefront.Rating,
			&storefront.TotalRawRating,
			&createdAt)
		createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
		errorkit.ErrorHandled(err)

		storefront.CreatedAt = createdAtProto
		storefronts = append(storefronts, &storefront)
	}

	log.Println("storefronts : ", storefronts)

	return storefronts
}

func (ir *ItemsReviewed) recalculateStorefrontRating(storefronts []*StorefrontWithCount, itemReviews []*store.ItemReview) {
	for i := 0; i < len(itemReviews); i++ {
		for j := 0; j < len(storefronts); j++ {
			if itemReviews[i].Item.Storefront.Uuid == storefronts[j].Uuid {
				// (*storefronts[j]).CountedTotalReview++
				(*storefronts[j]).TotalRawRating += itemReviews[i].Rating
				(*storefronts[j]).Rating = storefronts[j].TotalRawRating / float64(storefronts[j].CountedTotalReview)
			}
		}
	}

	log.Println(storefronts)
}
