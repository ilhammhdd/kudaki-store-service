package adapters

import (
	"log"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-store-service/usecases"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-entities/events"
)

func AddStorefrontItem(partition int32, offset int64, key string, value []byte, dbo usecases.DBOperation) (string, []byte, error) {
	var asir events.AddStorefrontItemRequested

	var unmarshalErr = proto.Unmarshal(value, &asir)
	if unmarshalErr == nil {
		log.Printf("consumed AddStorefronItemRequested : partition = %d, offset = %d, key = %s", partition, offset, key)

		storefrontItemAdded := usecases.AddStorefrontItem(&asir, dbo)
		storefrontItemAddedBytes, marshalErr := proto.Marshal(storefrontItemAdded)
		errorkit.ErrorHandled(marshalErr)

		return storefrontItemAdded.Uuid, storefrontItemAddedBytes, nil
	}
	return "", nil, unmarshalErr
}
