package usecases

import "github.com/golang/protobuf/proto"

type RetrieveItems struct {
	DBO DBOperator
}

func (ri RetrieveItems) Handle(in proto.Message) (out proto.Message) {
	
	return nil
}
