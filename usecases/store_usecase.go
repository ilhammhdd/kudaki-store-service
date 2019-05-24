package usecases

import (
	"database/sql"
	"log"
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

	if len(itemIDs) > 0 {
		sir.Items = items
		sir.First = int32(*itemIDs[0])
		sir.Last = int32(*itemIDs[len(itemIDs)-1])
	}

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
		if errorkit.ErrorHandled(err) {
			return nil, nil, err
		}

		itemIDs = append(itemIDs, &itemID)
		items.Items = append(items.Items, &item)
	}

	return &items, itemIDs, nil
}

type StorefrontItemUpdate struct {
	In  *events.UpdateStorefrontItemRequested
	DBO DBOperation
}

func (s StorefrontItemUpdate) checkItemOwnership(siu *events.StorefrontItemUpdated) bool {
	row, err := s.DBO.QueryRow("SELECT id FROM items WHERE storefront_uuid = (SELECT uuid FROM storefronts WHERE user_uuid = ?) AND uuid = ?;", s.In.User.Uuid, s.In.Item.Uuid)
	errorkit.ErrorHandled(err)

	var itemID int64
	if err = row.Scan(&itemID); err == sql.ErrNoRows {
		log.Println("item with the given uuid doesn't belong to the authenticated user")

		siu.EventStatus.Errors = []string{err.Error()}
		siu.EventStatus.Timestamp = ptypes.TimestampNow()
		siu.EventStatus.HttpCode = http.StatusInternalServerError

		return false
	}

	return true
}

func (s StorefrontItemUpdate) updateItemInDB(siu *events.StorefrontItemUpdated) error {
	commandErr := s.DBO.Command(
		"UPDATE items SET name=?,amount=?,unit=?,price=?,description=?,photo=? WHERE uuid=?;",
		s.In.Item.Name,
		s.In.Item.Amount,
		s.In.Item.Unit,
		s.In.Item.Price,
		s.In.Item.Description,
		s.In.Item.Photo,
		s.In.Item.Uuid)

	if errorkit.ErrorHandled(commandErr) {
		siu.EventStatus.Errors = []string{commandErr.Error()}
		siu.EventStatus.HttpCode = http.StatusInternalServerError
		siu.EventStatus.Timestamp = ptypes.TimestampNow()

		return commandErr
	}

	return commandErr
}

func (s StorefrontItemUpdate) Update() *events.StorefrontItemUpdated {

	var siu events.StorefrontItemUpdated
	siu.EventStatus = &events.Status{}
	siu.Uid = s.In.Uid
	siu.User = s.In.User

	if !s.checkItemOwnership(&siu) {
		return &siu
	}

	if updateErr := s.updateItemInDB(&siu); updateErr != nil {
		return &siu
	}

	siu.Item = s.In.Item
	siu.EventStatus.HttpCode = http.StatusOK
	siu.EventStatus.Timestamp = ptypes.TimestampNow()

	return &siu
}

type ItemsRetrieval struct {
	In  *events.RetrieveItemsRequested
	DBO DBOperation
}

func (ir ItemsRetrieval) Retrieve() *events.ItemsRetrieved {

	var out events.ItemsRetrieved
	out.Uid = ir.In.Uid
	out.EventStatus = &events.Status{
		HttpCode:  http.StatusOK,
		Timestamp: ptypes.TimestampNow(),
	}

	items, ids, err := ir.retrieveFromDB()
	if err != nil {
		out.EventStatus.Errors = []string{err.Error()}
		out.EventStatus.HttpCode = http.StatusInternalServerError

		return &out
	}

	if len(ids) > 0 {
		out.First = int32(*ids[0])
		out.Items = items
		out.Last = int32(*ids[len(ids)-1])
	}

	return &out
}

func (ir ItemsRetrieval) retrieveFromDB() (*store.Items, []*int64, error) {
	var items store.Items
	items.Items = []*store.Item{}
	var itemIDs []*int64

	rows, err := ir.DBO.Query("SELECT id,uuid,name,amount,unit,price,description,photo,rating FROM items WHERE id >= ? LIMIT ?", ir.In.From, ir.In.Limit)
	if errorkit.ErrorHandled(err) {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var itemID int64
		var item store.Item

		err = rows.Scan(&itemID, &item.Uuid, &item.Name, &item.Amount, &item.Unit, &item.Price, &item.Description, &item.Photo, &item.Rating)
		if errorkit.ErrorHandled(err) {
			return nil, nil, err
		}

		itemIDs = append(itemIDs, &itemID)
		items.Items = append(items.Items, &item)
	}

	return &items, itemIDs, nil
}

type ItemRetrieval struct {
	DBO DBOperation
	In  *events.RetrieveItemRequested
}

func (ir ItemRetrieval) Retrieve() *events.ItemRetrieved {

	var out events.ItemRetrieved
	out.EventStatus = &events.Status{}
	out.Uid = ir.In.Uid
	out.User = ir.In.User
	out.EventStatus.HttpCode = http.StatusOK

	// retrieve item with the given uuid from db
	row, err := ir.DBO.QueryRow("SELECT uuid,name,amount,unit,price,description,photo,rating FROM items WHERE uuid=?;", ir.In.ItemUuid)
	errorkit.ErrorHandled(err)

	var item store.Item
	err = row.Scan(&item.Uuid, &item.Name, &item.Amount, &item.Unit, &item.Price, &item.Description, &item.Photo, &item.Rating)
	if err == sql.ErrNoRows {
		out.EventStatus.Errors = []string{"item with the given ID doesn't exists"}
		out.EventStatus.HttpCode = http.StatusNotFound
		out.EventStatus.Timestamp = ptypes.TimestampNow()

		return &out
	}
	// retrieve item with the given uuid from db

	out.Item = &item
	out.EventStatus.Timestamp = ptypes.TimestampNow()

	return &out
}

type ItemsSearch struct {
	DBO DBOperation
	In  *events.SearchItemsRequested
}

func (is ItemsSearch) Search() *events.ItemsSearched {

	// create out event
	var out events.ItemsSearched
	out.EventStatus = &events.Status{}
	out.EventStatus.HttpCode = http.StatusOK
	out.Uid = is.In.Uid
	out.User = is.In.User
	out.Limit = is.In.Limit
	out.Keyword = is.In.Keyword
	// create out event

	// fulltext search mysql
	items, itemIDs, err := is.retrieveFromDB()
	errorkit.ErrorHandled(err)
	// fulltext search mysql

	// cram items to out event if result set > 0
	if len(itemIDs) > 0 {
		out.First = uint64(*itemIDs[0])
		out.Items = items
		out.Last = uint64(*itemIDs[len(itemIDs)-1])
	}
	// cram items to out event if result set > 0

	return &out
}

func (is ItemsSearch) retrieveFromDB() (*store.Items, []*int64, error) {
	var items store.Items
	items.Items = []*store.Item{}
	var itemIDs []*int64
	log.Printf("item search : from = %d, limit = %d", is.In.From, is.In.Limit)
	rows, err := is.DBO.Query("SELECT * FROM items WHERE MATCH(name,description) AGAINST(?) AND id >= ? LIMIT ?;", is.In.Keyword, is.In.From, is.In.Limit)
	if errorkit.ErrorHandled(err) {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var itemID int64
		var item store.Item
		item.Storefront = &store.Storefront{}

		err = rows.Scan(&itemID, &item.Uuid, &item.Storefront.Uuid, &item.Name, &item.Amount, &item.Unit, &item.Price, &item.Description, &item.Photo, &item.Rating)
		if errorkit.ErrorHandled(err) {
			return nil, nil, err
		}

		itemIDs = append(itemIDs, &itemID)
		items.Items = append(items.Items, &item)
	}

	return &items, itemIDs, nil
}
