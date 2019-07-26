package usecases

import "github.com/golang/protobuf/proto"

type RetrieveItems struct {
	DBO DBOperator
}

func (ri RetrieveItems) Handle(in proto.Message) (out proto.Message) {
	ri.DBO.Query("SELECT i.id,i.uuid,i.storefront_uuid,i.name,i.amount,i.unit,i.price,i.description,i.photo,i.rating,i.created_at,sf.id,sf.uuid,sf.user_uuid,sf.total_item,sf.rating,sf.created_at,u.id,u.uuid,u.email,u.role,u.phone_number,u.account_type,u.created_at,p.id,p.user_uuid,p.uuid,p.full_name,p.photo,p.reputation,p.created_at FROM kudaki_store.items i JOIN kudaki_store.storefronts sf ON sf.uuid=i.storefront_uuid JOIN kudaki_user.users u ON u.uuid=sf.user_uuid JOIN kudaki_user.profiles p ON u.uuid=p.user_uuid;")
	defer rows.Close()
	return nil
}
