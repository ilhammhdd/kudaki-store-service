package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
)

type CartItemAdded struct{}

func (cia *CartItemAdded) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	// if !usecaseRes.Ok {
	// 	return
	// }

	// in := inEvent.(*events.CartItemAdded)

	// cia.updateItem(in.CartItem.Item)
	// cia.updateStorefront(in.CartItem.Item.Storefront)
}

func (cia *CartItemAdded) Work() interface{} {
	// usecase := &usecases.CartItemAdded{DBO: mysql.NewDBOperation()}

	// edde := EventDrivenDownstreamExternal{
	// 	PostUsecaseExecutor: cia,
	// 	eventDrivenAdapter:  new(adapters.CartItemAdded),
	// 	eventDrivenUsecase:  usecase,
	// 	eventName:           events.RentalTopic_CART_ITEM_ADDED.String(),
	// 	inTopics:            []string{events.RentalTopic_CART_ITEM_ADDED.String()}}

	// edde.handle()
	return nil
}

func (cia *CartItemAdded) updateItem(item *store.Item) {
	// dbo := mysql.NewDBOperation()
	// _, err := dbo.Command("UPDATE items SET amount=? WHERE uuid=?", item.Amount, item.Uuid)
	// errorkit.ErrorHandled(err)

	/* grpcConn, err := grpc.Dial(os.Getenv("STORE_REPO_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
	errorkit.ErrorHandled(err)
	defer grpcConn.Close()

	itemRepoClient := kudakigrpc.NewItemRepoClient(grpcConn)
	_, err = itemRepoClient.UpdateItem(context.Background(), item)
	errorkit.ErrorHandled(err) */
}

func (cia *CartItemAdded) updateStorefront(storefront *store.Storefront) {
	// dbo := mysql.NewDBOperation()
	// _, err := dbo.Command("UPDATE storefronts SET total_item=? WHERE uuid=?", storefront.TotalItem, storefront.Uuid)
	// errorkit.ErrorHandled(err)

	/* grpcConn, err := grpc.Dial(os.Getenv("STORE_REPO_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
	errorkit.ErrorHandled(err)
	defer grpcConn.Close()

	storeRepoClient := kudakigrpc.NewStoreRepoClient(grpcConn)
	_, err = storeRepoClient.UpsertStorefront(context.Background(), storefront)
	errorkit.ErrorHandled(err) */
}
