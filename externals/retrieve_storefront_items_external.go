package externals

import (
	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type RetrieveStorefrontItems struct{}

func (rsi *RetrieveStorefrontItems) Work() interface{} {
	adapter := new(adapters.RetrieveStorefrontItems)

	usecase := usecases.RetrieveStorefrontItems{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: &usecase,
		eventName:          events.StorefrontServiceCommandTopic_RETRIEVE_STOREFRONT_ITEMS.String(),
		inTopics:           []string{events.StorefrontServiceCommandTopic_RETRIEVE_STOREFRONT_ITEMS.String()},
		outTopic:           events.StorefrontServiceEventTopic_STOREFRONT_ITEMS_RETRIEVED.String()}

	ede.handle()
	return nil
}
