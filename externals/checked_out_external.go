package externals

import "github.com/golang/protobuf/proto"

type CheckedOut struct{}

func (co *CheckedOut) Work() interface{} {

	return nil
}

func (co *CheckedOut) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {

}
