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

		return storefrontItemAdded.Uid, storefrontItemAddedBytes, nil
	}
	return "", nil, unmarshalErr
}

func DeleteStorefrontItem(partition int32, offset int64, key string, value []byte, dbo usecases.DBOperation) (string, []byte, error) {

	var dsir events.DeleteStorefrontItemRequested

	unmarshallErr := proto.Unmarshal(value, &dsir)
	if unmarshallErr == nil {
		log.Printf("consumed DeleteStorefrontItemRequested : partition = %d, offset = %d, key = %s", partition, offset, key)

		sid := &usecases.StorefrontItemDeletion{
			DBO: dbo,
			In:  &dsir,
		}
		storefrontItemDeletedEvent := sid.Delete()
		resultedEvent, err := proto.Marshal(storefrontItemDeletedEvent)
		errorkit.ErrorHandled(err)

		return storefrontItemDeletedEvent.Uid, resultedEvent, nil
	}

	return "", nil, unmarshallErr
}

type StorefrontItemsRetrieval struct {
	Partition int32
	Offset    int64
	Key       string
	Value     []byte
}

func (s StorefrontItemsRetrieval) Retrieve() (string, []byte) {

	var sirr events.RetrieveStorefrontItemsRequested

	unmarshalErr := proto.Unmarshal(s.Value, &sirr)
	if unmarshalErr == nil {
		log.Printf("consumed RetrieveStorefrontItemsRequested : partition = %d, offset = %d, key = %s", s.Partition, s.Offset, s.Key)

	}

	return "", nil
}
