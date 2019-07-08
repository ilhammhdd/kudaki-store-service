package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases/events"
)

type RetrieveStorefrontItems struct {
	Result RetrieveStorefrontItemsResultSchema
}

func (rsi *RetrieveStorefrontItems) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveStorefrontItems
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rsi *RetrieveStorefrontItems) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.StorefrontItemsRetrieved)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveStorefrontItemsResultSchema struct {
	Storefront *usecases.StorefrontTemp `json:"storefront"`
	Items      []*usecases.ItemTemp     `json:"items"`
}

func (rsi *RetrieveStorefrontItems) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rsi.Result = RetrieveStorefrontItemsResultSchema{
		Storefront: i[0].(*usecases.StorefrontTemp),
		Items:      i[1].([]*usecases.ItemTemp),
	}

	return rsi
}

func (rsi *RetrieveStorefrontItems) ParseToResult() []byte {
	result, err := json.Marshal(rsi.Result)
	errorkit.ErrorHandled(err)

	return result
}
