package usecases

import (
	"database/sql"
	"net/http"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/store"
)

func AddStorefrontItem(in *events.AddStorefrontItemRequested, dbo DBOperation) *events.StorefrontItemAdded {

	row, err := dbo.QueryRow("SELECT uuid,user_uuid,total_item,rating FROM storefronts WHERE user_uuid=?", in.Item.Storefront.UserUuid)
	errorkit.ErrorHandled(err)

	var storeFront store.Storefront
	var newStoreFrontUUID string
	scanErr := row.Scan(&storeFront.Uuid, &storeFront.UserUuid, &storeFront.TotalItem, &storeFront.Rating)
	if scanErr == sql.ErrNoRows {
		newStoreFrontUUID = uuid.New().String()
		cmdErr := dbo.Command("INSERT INTO storefronts(uuid,user_uuid,total_item,rating) VALUES(?,?,?,?)", newStoreFrontUUID, in.Item.Storefront.UserUuid, 0, 0)
		errorkit.ErrorHandled(cmdErr)

		cmdErr = dbo.Command(
			"INSERT INTO items(uuid,storefront_uuid,name,amount,unit,price,description,photo,rating) VALUES (?,?,?,?,?,?,?,?,?)",
			uuid.New().String(), newStoreFrontUUID, in.Item.Name, in.Item.Amount, in.Item.Unit, in.Item.Price, in.Item.Description, in.Item.Photo, in.Item.Rating)
		errorkit.ErrorHandled(cmdErr)
	} else {
		cmdErr := dbo.Command(
			"INSERT INTO items(uuid,storefront_uuid,name,amount,unit,price,description,photo,rating) VALUES (?,?,?,?,?,?,?,?,?)",
			uuid.New().String(), storeFront.Uuid, in.Item.Name, in.Item.Amount, in.Item.Unit, in.Item.Price, in.Item.Description, in.Item.Photo, in.Item.Rating)
		errorkit.ErrorHandled(cmdErr)
	}

	var storefrontItemAdded events.StorefrontItemAdded
	storefrontItemAdded.EventStatus = &events.Status{
		HttpCode: http.StatusOK,
	}
	storefrontItemAdded.Item = in.Item
	storefrontItemAdded.Uid = in.Uid

	return &storefrontItemAdded
}

type StorefrontItemDeletion struct {
	In  *events.DeleteStorefrontItemRequested
	DBO DBOperation
}

func (sid *StorefrontItemDeletion) Delete() *events.StorefrontItemDeleted {

	var storefrontItemDeletedEvent events.StorefrontItemDeleted
	storefrontItemDeletedEvent.EventStatus = &events.Status{}
	storefrontItemDeletedEvent.Uid = sid.In.Uid
	storefrontItemDeletedEvent.EventStatus.Timestamp = ptypes.TimestampNow()

	item, err := sid.retrieveItem()
	if err == sql.ErrNoRows {
		storefrontItemDeletedEvent.EventStatus.Errors = []string{"item not found"}
		storefrontItemDeletedEvent.EventStatus.HttpCode = http.StatusBadRequest
		return &storefrontItemDeletedEvent
	}

	storefrontItemDeletedEvent.Item = item

	err = sid.DBO.Command("DELETE FROM items WHERE uuid = ?", item.Uuid)
	if errorkit.ErrorHandled(err) {
		storefrontItemDeletedEvent.EventStatus.Errors = []string{err.Error()}
		storefrontItemDeletedEvent.EventStatus.HttpCode = http.StatusInternalServerError
	}
	storefrontItemDeletedEvent.EventStatus.HttpCode = http.StatusOK

	return &storefrontItemDeletedEvent
}

func (sid *StorefrontItemDeletion) retrieveItem() (*store.Item, error) {
	var item store.Item
	item.Storefront = &store.Storefront{}
	var itemID int64
	var storefrontID int64
	row, err := sid.DBO.QueryRow("SELECT * FROM items JOIN storefronts ON storefronts.uuid = items.storefront_uuid WHERE items.uuid = ? AND storefront_uuid = (SELECT uuid FROM storefronts WHERE user_uuid = ?);", sid.In.Item.Uuid, sid.In.Item.Storefront.UserUuid)
	errorkit.ErrorHandled(err)

	err = row.Scan(&itemID, &item.Uuid, &item.Storefront.Uuid, &item.Name, &item.Amount, &item.Unit, &item.Price, &item.Description, &item.Photo, &item.Rating, &storefrontID, &item.Storefront.Uuid, &item.Storefront.UserUuid, &item.Storefront.TotalItem, &item.Storefront.Rating)

	return &item, err
}

type StorefrontItemsRetrieval struct {
	In  *events.RetrieveStorefrontItemsRequested
	DBO DBOperation
}

func (sirr StorefrontItemsRetrieval) Retrieve() *events.StorefrontItemsRetrieved {

	var sir events.StorefrontItemsRetrieved
	sir.EventStatus = &events.Status{}
	sir.Uid = sirr.In.Uid
	sir.Limit = sirr.In.Limit

	items, itemIDs, err := sirr.retrieveFromDB()
	if err != nil {
		sir.EventStatus.Errors = []string{err.Error()}
		sir.EventStatus.HttpCode = http.StatusInternalServerError
		sir.EventStatus.Timestamp = ptypes.TimestampNow()

		return &sir
	}

	sir.Items = items
	sir.First = int32(*itemIDs[0])
	sir.Last = int32(*itemIDs[len(itemIDs)-1])
	sir.EventStatus.HttpCode = http.StatusOK
	sir.EventStatus.Timestamp = ptypes.TimestampNow()

	return &sir
}

func (sirr StorefrontItemsRetrieval) retrieveFromDB() (*store.Items, []*int64, error) {
	var items store.Items
	items.Items = []*store.Item{}
	var itemIDs []*int64

	rows, err := sirr.DBO.Query("SELECT id,uuid,name,amount,unit,price,description,photo,rating FROM items WHERE id >= ? AND storefront_uuid = (SELECT uuid FROM storefronts WHERE user_uuid = ?) LIMIT ?", sirr.In.From, sirr.In.User.Uuid, sirr.In.Limit)
	if errorkit.ErrorHandled(err) {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var itemID int64
		var item store.Item

		err = rows.Scan(&itemID, &item.Uuid, &item.Name, &item.Amount, &item.Unit, &item.Price, &item.Description, &item.Photo, &item.Rating)
		errorkit.ErrorHandled(err)

		itemIDs = append(itemIDs, &itemID)
		items.Items = append(items.Items, &item)
	}

	return &items, itemIDs, nil
}
