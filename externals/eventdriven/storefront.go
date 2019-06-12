package eventdriven

import (
	"log"
	"net/http"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/store"
	"github.com/ilhammhdd/kudaki-externals/mysql"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"
	"github.com/ilhammhdd/kudaki-store-service/adapters"
	"github.com/ilhammhdd/kudaki-store-service/usecases"
)

type AddStorefrontItem struct{}

func (asi *AddStorefrontItem) Work() interface{} {
	usecase := &usecases.AddStorefrontItem{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: asi,
		eventDrivenAdapter:  new(adapters.AddStorefrontItem),
		eventDrivenUsecase:  usecase,
		eventName:           events.StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED.String(),
		inTopics:            []string{events.StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED.String()},
		outTopic:            events.StoreTopic_STOREFRONT_ITEM_ADDED.String()}

	ede.handle()
	return nil
}

func (asi *AddStorefrontItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.StorefrontItemAdded)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	if out.Storefront == nil {
		newStorefront := asi.initStorefront(out.User, out.Item)
		out.Item.Storefront = newStorefront
		out.Storefront = newStorefront
	}

	asi.upsertStorefront(out.Storefront)
	asi.indexStorefront(out.Storefront)
	asi.insertItem(out.Item)
	asi.indexItem(out.Item)
}

func (asi *AddStorefrontItem) initStorefront(usr *user.User, item *store.Item) *store.Storefront {
	return &store.Storefront{
		Rating:    0.0,
		TotalItem: item.Amount,
		User:      usr,
		Uuid:      uuid.New().String()}
}

func (asi *AddStorefrontItem) upsertStorefront(storefront *store.Storefront) {
	dbo := mysql.NewDBOperation()

	_, err := dbo.Command("INSERT INTO storefronts(uuid,user_uuid,total_item,rating) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE total_item=?;", storefront.Uuid, storefront.User.Uuid, storefront.TotalItem, storefront.Rating, storefront.TotalItem)
	errorkit.ErrorHandled(err)
}

func (asi *AddStorefrontItem) insertItem(item *store.Item) {
	dbo := mysql.NewDBOperation()

	_, err := dbo.Command("INSERT INTO items(uuid,storefront_uuid,name,amount,unit,price,description,photo,rating) VALUES(?,?,?,?,?,?,?,?,?);",
		item.Uuid, item.Storefront.Uuid, item.Name, item.Amount, item.Unit, item.Price, item.Description, item.Photo, item.Rating)
	errorkit.ErrorHandled(err)
}

func (asi *AddStorefrontItem) indexItem(item *store.Item) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
	client.CreateIndex(kudakiredisearch.Item.Schema())

	sanitizedItemUUID := kudakiredisearch.RedisearchText(item.Uuid).Sanitize()
	doc := redisearch.NewDocument(sanitizedItemUUID, 1.0)
	doc.Set("item_uuid", sanitizedItemUUID)
	doc.Set("storefront_uuid", kudakiredisearch.RedisearchText(item.Storefront.Uuid).Sanitize())
	doc.Set("item_name", item.Name)
	doc.Set("item_amount", item.Amount)
	doc.Set("item_unit", item.Unit)
	doc.Set("item_price", item.Price)
	doc.Set("item_description", item.Description)
	doc.Set("item_photo", item.Photo)
	doc.Set("item_rating", item.Rating)

	err := client.IndexOptions(redisearch.DefaultIndexingOptions, doc)
	errorkit.ErrorHandled(err)
}

func (asi *AddStorefrontItem) indexStorefront(storefront *store.Storefront) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
	client.CreateIndex(kudakiredisearch.Storefront.Schema())

	sanitizedStorefrontUUID := kudakiredisearch.RedisearchText(storefront.Uuid).Sanitize()
	doc := redisearch.NewDocument(sanitizedStorefrontUUID, 1.0)
	doc.Set("storefront_uuid", sanitizedStorefrontUUID)
	doc.Set("user_uuid", kudakiredisearch.RedisearchText(storefront.User.Uuid).Sanitize())
	doc.Set("storefront_total_item", storefront.TotalItem)
	doc.Set("storefront_rating", storefront.Rating)

	err := client.IndexOptions(redisearch.IndexingOptions{Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

type UpdateStorefrontItem struct{}

func (usi *UpdateStorefrontItem) Work() interface{} {
	usecase := &usecases.UpdateStorefrontItem{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: usi,
		eventDrivenAdapter:  new(adapters.UpdateStorefrontItem),
		eventDrivenUsecase:  usecase,
		eventName:           events.StoreTopic_UPDATE_STOREFRONT_ITEM_REQUESTED.String(),
		inTopics:            []string{events.StoreTopic_UPDATE_STOREFRONT_ITEM_REQUESTED.String()},
		outTopic:            events.StoreTopic_STOREFRONT_ITEM_UPDATED.String()}

	ede.handle()
	return nil
}

func (usi *UpdateStorefrontItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.StorefrontItemUpdated)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	usi.updateStorefront(out.Storefront)
	usi.reIndexStorefront(out.Storefront)
	usi.updateItem(out.Item)
	usi.reIndexItem(out.Item)
}

func (usi *UpdateStorefrontItem) updateStorefront(updatedStorefront *store.Storefront) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?;", updatedStorefront.TotalItem, updatedStorefront.Uuid)
	errorkit.ErrorHandled(err)
}

func (usi *UpdateStorefrontItem) reIndexStorefront(updatedStorefront *store.Storefront) {
	log.Println(updatedStorefront)
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
	client.CreateIndex(kudakiredisearch.Storefront.Schema())

	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(updatedStorefront.Uuid).Sanitize(), 1.0)
	doc.Set("storefront_total_item", updatedStorefront.TotalItem)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

func (usi *UpdateStorefrontItem) updateItem(updatedItem *store.Item) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE items SET name=?,amount=?,unit=?,price=?,description=?,photo=? WHERE uuid=?;",
		updatedItem.Name,
		updatedItem.Amount,
		updatedItem.Unit,
		updatedItem.Price,
		updatedItem.Description,
		updatedItem.Photo,
		updatedItem.Uuid)
	errorkit.ErrorHandled(err)
}

func (usi *UpdateStorefrontItem) reIndexItem(updatedItem *store.Item) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
	client.CreateIndex(kudakiredisearch.Item.Schema())

	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(updatedItem.Uuid).Sanitize(), 1.0)
	doc.Set("item_name", updatedItem.Name)
	doc.Set("item_amount", updatedItem.Amount)
	doc.Set("item_unit", updatedItem.Unit)
	doc.Set("item_price", updatedItem.Price)
	doc.Set("item_photo", updatedItem.Photo)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

type DeleteStorefrontItem struct{}

func (dsi *DeleteStorefrontItem) Work() interface{} {
	usecase := &usecases.DeleteStorefrontItem{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: dsi,
		eventDrivenAdapter:  new(adapters.DeleteStorefrontItem),
		eventDrivenUsecase:  usecase,
		eventName:           events.StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED.String(),
		inTopics:            []string{events.StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED.String()},
		outTopic:            events.StoreTopic_STOREFRONT_ITEM_DELETED.String()}

	ede.handle()
	return nil
}

func (dsi *DeleteStorefrontItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.StorefrontItemDeleted)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dsi.deleteItemFromDB(out.Item, out.Storefront)
	dsi.deleteItemDoc(out.Item)
	dsi.updateStorefront(out.Storefront)
	dsi.reIndexStorefront(out.Storefront)
}

func (dsi *DeleteStorefrontItem) deleteItemFromDB(item *store.Item, storefront *store.Storefront) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("DELETE FROM items WHERE uuid=? AND storefront_uuid=?;", item.Uuid, storefront.Uuid)
	errorkit.ErrorHandled(err)
}

func (dsi *DeleteStorefrontItem) deleteItemDoc(item *store.Item) {
	host := redisearch.NewSingleHostPool(os.Getenv("REDISEARCH_SERVER"))
	defer host.Close()
	conn := host.Get()
	defer conn.Close()
	_, err := conn.Do("FT.DEL", kudakiredisearch.Item.Name(), kudakiredisearch.RedisearchText(item.Uuid).Sanitize(), "DD")
	errorkit.ErrorHandled(err)
}

func (dsi *DeleteStorefrontItem) updateStorefront(storefront *store.Storefront) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?;", storefront.TotalItem, storefront.Uuid)
	errorkit.ErrorHandled(err)
}

func (dsi *DeleteStorefrontItem) reIndexStorefront(storefront *store.Storefront) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
	client.CreateIndex(kudakiredisearch.Storefront.Schema())

	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(storefront.Uuid).Sanitize(), 1.0)
	doc.Set("storefront_total_item", storefront.TotalItem)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}
