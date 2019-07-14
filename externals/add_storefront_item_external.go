package externals

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
)

type AddStorefrontItem struct{}

func (asi *AddStorefrontItem) Work() interface{} {
	usecase := &usecases.AddStorefrontItem{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: asi,
		eventDrivenAdapter:  new(adapters.AddStorefrontItem),
		eventDrivenUsecase:  usecase,
		eventName:           events.StorefrontServiceCommandTopic_ADD_STOREFRONT_ITEM.String(),
		inTopics:            []string{events.StorefrontServiceCommandTopic_ADD_STOREFRONT_ITEM.String()},
		outTopic:            events.StorefrontServiceEventTopic_STOREFRONT_ITEM_ADDED.String()}

	ede.handle()
	return nil
}

func (asi *AddStorefrontItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.StorefrontItemAdded)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	if out.Storefront == nil {
		newStorefront := asi.initStorefront(out.Requester, out.Item)
		out.Item.Storefront = newStorefront
		out.Storefront = newStorefront
	} else {
		out.Item.Storefront = out.Storefront
	}

	asi.upsertStorefront(out.Storefront)
	asi.insertItem(out.Item)
}

func (asi *AddStorefrontItem) initStorefront(usr *user.User, item *store.Item) *store.Storefront {
	return &store.Storefront{
		Rating:    0.0,
		TotalItem: item.Amount,
		UserUuid:  usr.Uuid,
		Uuid:      uuid.New().String()}
}

func (asi *AddStorefrontItem) upsertStorefront(storefront *store.Storefront) {
	log.Println("storefront : ", storefront)
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	log.Println("dbo : ", dbo)
	_, err := dbo.Command("INSERT INTO kudaki_store.storefronts(uuid,user_uuid,total_item,rating,created_at) VALUES(?,?,?,?,UNIX_TIMESTAMP()) ON DUPLICATE KEY UPDATE total_item=?;",
		storefront.Uuid,
		storefront.UserUuid,
		storefront.TotalItem,
		storefront.Rating,
		storefront.TotalItem)
	errorkit.ErrorHandled(err)
}

func (asi *AddStorefrontItem) insertItem(item *store.Item) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_store.items(uuid,storefront_uuid,name,amount,unit,price,price_duration,description,photo,rating,length,width,height,color,unit_of_measurement,created_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,UNIX_TIMESTAMP());",
		item.Uuid,
		item.Storefront.Uuid,
		item.Name,
		item.Amount,
		item.Unit,
		item.Price,
		item.PriceDuration.String(),
		item.Description,
		item.Photo,
		item.Rating,
		item.ItemDimension.Length,
		item.ItemDimension.Width,
		item.ItemDimension.Height,
		item.Color,
		item.ItemDimension.UnitOfMeasurement.String())
	errorkit.ErrorHandled(err)
}
