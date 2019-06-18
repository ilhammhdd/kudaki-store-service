package eventdriven

import (
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/store"
	"github.com/ilhammhdd/kudaki-externals/mysql"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"
	"github.com/ilhammhdd/kudaki-store-service/adapters"
	"github.com/ilhammhdd/kudaki-store-service/usecases"
)

type CartItemAdded struct{}

func (cia *CartItemAdded) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseStat *usecases.UsecaseHandlerStatus) {
	if !usecaseStat.Ok {
		return
	}

	in := inEvent.(*events.CartItemAdded)

	cia.updateItem(in.CartItem.Item)
	cia.reIndexItem(in.CartItem.Item)
	cia.updateStorefront(in.CartItem.Item.Storefront)
	cia.reIndexStorefront(in.CartItem.Item.Storefront)
}

func (cia *CartItemAdded) Work() interface{} {
	usecase := &usecases.CartItemAdded{DBO: mysql.NewDBOperation()}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: cia,
		eventDrivenAdapter:  new(adapters.CartItemAdded),
		eventDrivenUsecase:  usecase,
		eventName:           events.RentalTopic_CART_ITEM_ADDED.String(),
		inTopics:            []string{events.RentalTopic_CART_ITEM_ADDED.String()}}

	edde.handle()
	return nil
}

func (cia *CartItemAdded) updateItem(item *store.Item) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE items SET amount=? WHERE uuid=?", item.Amount, item.Uuid)
	errorkit.ErrorHandled(err)
}

func (cia *CartItemAdded) reIndexItem(item *store.Item) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
	client.CreateIndex(kudakiredisearch.Item.Schema())

	sanitizer := new(kudakiredisearch.RedisearchText)
	sanitizer.Set(item.Uuid)
	doc := redisearch.NewDocument(sanitizer.Sanitize(), 1.0)
	doc.Set("item_amount", item.Amount)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

func (cia *CartItemAdded) updateStorefront(storefront *store.Storefront) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?", storefront.TotalItem, storefront.Uuid)
	errorkit.ErrorHandled(err)
}

func (cia *CartItemAdded) reIndexStorefront(storefront *store.Storefront) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
	client.CreateIndex(kudakiredisearch.Storefront.Schema())

	sanitizer := new(kudakiredisearch.RedisearchText)
	sanitizer.Set(storefront.Uuid)
	doc := redisearch.NewDocument(sanitizer.Sanitize(), 1.0)
	doc.Set("storefront_total_item", storefront.TotalItem)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

type CartItemDeleted struct {
}

func (cid *CartItemDeleted) Work() interface{} {
	adapter := &adapters.CartItemDeleted{
		Sanitizer: new(kudakiredisearch.RedisearchText)}
	usecase := &usecases.CartItemDeleted{
		DBO: mysql.NewDBOperation()}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: cid,
		eventDrivenAdapter:  adapter,
		eventDrivenUsecase:  usecase,
		eventName:           events.RentalTopic_CART_ITEM_DELETED.String(),
		inTopics:            []string{events.RentalTopic_CART_ITEM_DELETED.String()}}

	edde.handle()
	return nil
}

func (cid *CartItemDeleted) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseStat *usecases.UsecaseHandlerStatus) {
	if !usecaseStat.Ok {
		return
	}

	in := inEvent.(*events.CartItemDeleted)

	cid.updateItem(in.CartItem.Item)
	cid.reIndexItem(in.CartItem.Item)
	cid.updateStorefront(in.CartItem.Item.Storefront)
	cid.reIndexStorefront(in.CartItem.Item.Storefront)
}

func (cid *CartItemDeleted) updateItem(item *store.Item) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE items SET amount=? WHERE uuid=?;", item.Amount, item.Uuid)
	errorkit.ErrorHandled(err)
}

func (cid *CartItemDeleted) reIndexItem(item *store.Item) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
	client.CreateIndex(kudakiredisearch.Item.Schema())

	sanitizer := new(kudakiredisearch.RedisearchText)
	sanitizer.Set(item.Uuid)
	doc := redisearch.NewDocument(sanitizer.Sanitize(), 1.0)
	doc.Set("item_amount", item.Amount)
	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

func (cid *CartItemDeleted) updateStorefront(storefront *store.Storefront) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?;", storefront.TotalItem, storefront.Uuid)
	errorkit.ErrorHandled(err)
}

func (cid *CartItemDeleted) reIndexStorefront(storefront *store.Storefront) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
	client.CreateIndex(kudakiredisearch.Storefront.Schema())

	sanitizer := new(kudakiredisearch.RedisearchText)
	sanitizer.Set(storefront.Uuid)
	doc := redisearch.NewDocument(sanitizer.Sanitize(), 1.0)
	doc.Set("storefront_total_item", storefront.TotalItem)
	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}
