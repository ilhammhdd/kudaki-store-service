package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type RetrieveStorefrontItems struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rsi *RetrieveStorefrontItems) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rsi.initInOutEvent(in)

	storefront := rsi.retrieveStorefrontResult(outEvent.Requester)

	items := rsi.retrieveItemsResult(inEvent, storefront)
	rsi.ResultSchemer.SetResultSources(storefront, items)
	outEvent.Result = rsi.ResultSchemer.ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rsi *RetrieveStorefrontItems) initInOutEvent(in proto.Message) (inEvent *events.RetrieveStorefrontItems, outEvent *events.StorefrontItemsRetrieved) {
	inEvent = in.(*events.RetrieveStorefrontItems)

	outEvent = new(events.StorefrontItemsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveStorefrontItems = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rsi *RetrieveStorefrontItems) retrieveStorefront(usr *user.User) *store.Storefront {
	row, err := rsi.DBO.QueryRow("SELECT id,uuid,user_uuid,total_item,rating,created_at FROM kudaki_store.storefronts WHERE user_uuid=?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var storefront store.Storefront
	var createdAt int64
	if row.Scan(
		&storefront.Id,
		&storefront.Uuid,
		&storefront.UserUuid,
		&storefront.TotalItem,
		&storefront.Rating,
		&createdAt) == sql.ErrNoRows {
		return nil
	}
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)

	storefront.CreatedAt = createdAtProto
	return &storefront
}

type StorefrontTemp struct {
	store.Storefront
	RatingT    float64 `json:"rating"`
	CreatedAtT int64   `json:"created_at"`
}

type ItemTemp struct {
	store.Item
	PriceDurationT     string  `json:"price_duration"`
	StorefrontUuidT    string  `json:"storefront_uuid"`
	LengthT            int32   `json:"length"`
	WidthT             int32   `json:"width"`
	HeightT            int32   `json:"height"`
	UnitofMeasurementT string  `json:"unit_of_measurement"`
	RatingT            float64 `json:"rating"`
	CreatedAtT         int64   `json:"created_at"`
}

func (rsi *RetrieveStorefrontItems) retrieveStorefrontResult(usr *user.User) *StorefrontTemp {
	row, err := rsi.DBO.QueryRow("SELECT id,uuid,user_uuid,total_item,rating,created_at FROM kudaki_store.storefronts WHERE user_uuid=?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var storefront StorefrontTemp
	if row.Scan(
		&storefront.Id,
		&storefront.Uuid,
		&storefront.UserUuid,
		&storefront.TotalItem,
		&storefront.RatingT,
		&storefront.CreatedAtT) == sql.ErrNoRows {
		return nil
	}

	return &storefront
}

func (rsi *RetrieveStorefrontItems) retrieveItemsResult(inEvent *events.RetrieveStorefrontItems, storefront *StorefrontTemp) []*ItemTemp {
	var storefrontItems []*ItemTemp

	rows, err := rsi.DBO.Query("SELECT i.id,i.uuid,i.storefront_uuid,i.name,i.amount,i.unit,i.price,i.price_duration,i.description,i.photo,i.rating,i.length,i.width,i.height,i.color,i.unit_of_measurement,i.created_at FROM (SELECT id FROM kudaki_store.items WHERE storefront_uuid = ? LIMIT ?, ?) i_ids JOIN kudaki_store.items i ON i.id = i_ids.id;", storefront.Uuid, inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	for rows.Next() {
		var item ItemTemp
		var priceDuration string

		rows.Scan(
			&item.Id,
			&item.Uuid,
			&item.StorefrontUuidT,
			&item.Name,
			&item.Amount,
			&item.Unit,
			&item.Price,
			&item.PriceDurationT,
			&item.Description,
			&item.Photo,
			&item.Rating,
			&item.LengthT,
			&item.WidthT,
			&item.HeightT,
			&item.Color,
			&item.UnitofMeasurementT,
			&item.CreatedAtT)

		item.PriceDuration = store.PriceDuration(store.PriceDuration_value[priceDuration])

		storefrontItems = append(storefrontItems, &item)
	}

	return storefrontItems
}
