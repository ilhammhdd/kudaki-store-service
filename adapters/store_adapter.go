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

func (s StorefrontItemsRetrieval) Retrieve(dbo usecases.DBOperation) (string, []byte) {

	var rsir events.RetrieveStorefrontItemsRequested

	unmarshalErr := proto.Unmarshal(s.Value, &rsir)
	if unmarshalErr == nil {
		log.Printf("consumed RetrieveStorefrontItemsRequested : partition = %d, offset = %d, key = %s", s.Partition, s.Offset, s.Key)
		sirr := usecases.StorefrontItemsRetrieval{
			DBO: dbo,
			In:  &rsir,
		}
		sir := sirr.Retrieve()
		sirByte, err := proto.Marshal(sir)
		errorkit.ErrorHandled(err)

		return sir.Uid, sirByte
	}

	return "", nil
}

type StorefrontItemUpdate struct {
	Partition int32
	Offset    int64
	Key       string
	Message   []byte
}

func (s StorefrontItemUpdate) Update(dbo usecases.DBOperation) (key string, msg []byte, err error) {

	var usir events.UpdateStorefrontItemRequested

	if unmarshalErr := proto.Unmarshal(s.Message, &usir); unmarshalErr == nil {
		log.Printf("consumed UpdateStorefrontItemRequested : partition = %d, offset = %d, key = %s", s.Partition, s.Offset, s.Key)

		siu := usecases.StorefrontItemUpdate{
			DBO: dbo,
			In:  &usir,
		}
		siud := siu.Update()
		siudByte, marshallErr := proto.Marshal(siud)
		errorkit.ErrorHandled(marshallErr)

		return siud.Uid, siudByte, nil
	} else {
		return "", nil, unmarshalErr
	}
}

type ItemsRetrieval struct {
	Partition int32
	Offset    int64
	Key       string
	Message   []byte
}

func (ir ItemsRetrieval) Retrieve(dbo usecases.DBOperation) (key string, msg []byte, err error) {

	var rir events.RetrieveItemsRequested

	if unmarshalErr := proto.Unmarshal(ir.Message, &rir); unmarshalErr == nil {
		log.Printf("consumed RetrieveItemsRequested : partition = %d, offset = %d, key = %s", ir.Partition, ir.Offset, ir.Key)

		irUsecase := usecases.ItemsRetrieval{
			DBO: dbo,
			In:  &rir,
		}
		ird := irUsecase.Retrieve()
		irdByte, marshallErr := proto.Marshal(ird)
		errorkit.ErrorHandled(marshallErr)

		return ird.Uid, irdByte, nil
	} else {
		return "", nil, unmarshalErr
	}
}

type ItemRetrieval struct {
	Partition int32
	Offset    int64
	Key       string
	Message   []byte
}

func (ir ItemRetrieval) Retrieve(dbo usecases.DBOperation) (key string, msg []byte, err error) {

	var in events.RetrieveItemRequested

	unmarshalErr := proto.Unmarshal(ir.Message, &in)
	if unmarshalErr != nil {
		return "", nil, unmarshalErr
	}

	log.Printf("consumed RetrieveItemRequested : parition = %d, offset = %d, key = %s", ir.Partition, ir.Offset, ir.Key)

	irUsecase := usecases.ItemRetrieval{
		DBO: dbo,
		In:  &in,
	}
	ird := irUsecase.Retrieve()

	irdByte, err := proto.Marshal(ird)
	errorkit.ErrorHandled(err)

	return ird.Uid, irdByte, nil
}

type ItemsSearch struct {
	Partition int32
	Offset    int64
	Key       string
	Message   *[]byte
}

func (is ItemsSearch) Search(dbo usecases.DBOperation) (string, *[]byte, error) {

	var sir events.SearchItemsRequested

	unmarshalErr := proto.Unmarshal(*is.Message, &sir)
	if unmarshalErr != nil {
		return "", nil, unmarshalErr
	}

	log.Printf("consumed SearchItemRequested : parition = %d, offset = %d, key = %s", is.Partition, is.Offset, is.Key)

	isUsecase := usecases.ItemsSearch{
		DBO: dbo,
		In:  &sir,
	}

	isd := isUsecase.Search()
	isdByte, err := proto.Marshal(isd)
	errorkit.ErrorHandled(err)

	return isd.Uid, &isdByte, nil
}
