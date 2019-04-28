package usecases

import (
	"database/sql"
	"net/http"

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
	}

	cmdErr := dbo.Command(
		"INSERT INTO items(uuid,storefront_uuid,name,amount,unit,price,description,photo,rating) VALUES (?,?,?,?,?,?,?,?,?)",
		uuid.New().String(), storeFront.Uuid, in.Item.Name, in.Item.Amount, in.Item.Unit, in.Item.Price, in.Item.Description, in.Item.Photo, in.Item.Rating)
	errorkit.ErrorHandled(cmdErr)

	var storefrontItemAdded events.StorefrontItemAdded
	storefrontItemAdded.EventStatus = &events.Status{
		HttpCode: http.StatusOK,
	}
	storefrontItemAdded.Item = in.Item
	storefrontItemAdded.Uuid = in.Uuid

	return &storefrontItemAdded
}
