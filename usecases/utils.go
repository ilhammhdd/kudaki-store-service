package usecases

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/golang/protobuf/ptypes"

	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-storefront-service/entities/aggregates/store"
)

type RedisClient interface {
	Name() string
	Schema() *redisearch.Schema
}

type RedisearchSanitizer interface {
	Set(string)
	Sanitize() string
	UnSanitize() string
}

func IntendedItemExists(dbo DBOperator, storefront *store.Storefront, itemUUID string) (*store.Item, bool) {
	row, err := dbo.QueryRow("SELECT name,amount,unit,price,description,photo,rating FROM kudaki_store.items WHERE storefront_uuid=? AND uuid=?;", storefront.Uuid, itemUUID)
	errorkit.ErrorHandled(err)

	var intendedItem store.Item
	if row.Scan(
		&intendedItem.Name,
		&intendedItem.Amount,
		&intendedItem.Unit,
		&intendedItem.Price,
		&intendedItem.Description,
		&intendedItem.Photo,
		&intendedItem.Rating) == sql.ErrNoRows {
		return nil, false
	}
	intendedItem.Storefront = storefront
	intendedItem.Uuid = itemUUID

	return &intendedItem, true
}

func StorefrontExists(usr *user.User, dbo DBOperator) (*store.Storefront, bool) {
	row, err := dbo.QueryRow("SELECT id,uuid,user_uuid,total_item,rating,created_at FROM kudaki_store.storefronts WHERE user_uuid=?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var storefront store.Storefront
	var storefrontCreatedAt int64
	if row.Scan(
		&storefront.Id,
		&storefront.Uuid,
		&storefront.UserUuid,
		&storefront.TotalItem,
		&storefront.Rating,
		&storefrontCreatedAt) == sql.ErrNoRows {
		return nil, false
	}

	createdAtTimestamp, err := ptypes.TimestampProto(time.Unix(storefrontCreatedAt, 0))
	errorkit.ErrorHandled(err)
	storefront.CreatedAt = createdAtTimestamp

	return &storefront, true
}

func userInStorefrontIndexExists(storefrontClient RedisClient, sanitizer RedisearchSanitizer, usr *user.User) (storefrontUUID string, ok bool) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), storefrontClient.Name())
	client.CreateIndex(storefrontClient.Schema())

	sanitizer.Set(usr.Uuid)
	rawQuery := fmt.Sprintf(`@user_uuid:"%s"`, sanitizer.Sanitize())
	docs, total, err := client.Search(redisearch.NewQuery(rawQuery))
	errorkit.ErrorHandled(err)

	if total != 0 {
		sanitizer.Set(docs[0].Properties["storefront_uuid"].(string))
		return sanitizer.UnSanitize(), true
	}
	return "", false
}

func GetUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.UserRole(user.UserRole_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}
