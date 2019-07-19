package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type ItemsReviewed struct{}

func (ir *ItemsReviewed) Work() interface{} {
	uc := usecases.ItemsReviewed{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: ir,
		eventDrivenAdapter:  new(adapters.ItemsReviewed),
		eventDrivenUsecase:  &uc,
		eventName:           events.ItemReviewServiceEventTopic_ITEMS_REVIEWED.String(),
		inTopics:            []string{events.ItemReviewServiceEventTopic_ITEMS_REVIEWED.String()}}

	edde.handle()
	return nil
}

func (ir *ItemsReviewed) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	storefronts := usecaseRes.Data.([]*usecases.StorefrontWithCount)

	dbo := mysql.NewDBOperation(mysql.CommandDB)

	for i := 0; i < len(storefronts); i++ {
		_, err := dbo.Command("UPDATE kudaki_store.storefronts SET rating = ?,total_raw_rating = ? WHERE uuid=?;",
			storefronts[i].Rating, storefronts[i].TotalRawRating, storefronts[i].Uuid)
		errorkit.ErrorHandled(err)
	}
}
