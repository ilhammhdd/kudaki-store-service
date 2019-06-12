package redisearch

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type DerivedUUIDs []string

func (du *DerivedUUIDs) Combine() string {
	var combinedUUIDs string
	for _, UUID := range *du {
		combinedUUIDs = combinedUUIDs + "_" + UUID
	}

	hasher := sha256.New()
	hasher.Write([]byte(combinedUUIDs))
	hashedCombinedUuids := hex.EncodeToString(hasher.Sum(nil))

	return hashedCombinedUuids
}

type DerivedClient int

const (
	CartItem DerivedClient = iota
	MountainReview
)

func (dc DerivedClient) Name() string {
	return []string{
		"cart_item",
	}[dc]
}

func (dc DerivedClient) Schema() *redisearch.Schema {
	return []*redisearch.Schema{
		// CartItemSchema
		CombineSchemas([]*redisearch.Schema{Cart.Schema(), Item.Schema()}).
			AddField(redisearch.NewTextField("cart_item_uuid")).
			AddField(redisearch.NewSortableNumericField("cart_item_total_amount")).
			AddField(redisearch.NewSortableNumericField("cart_item_total_price")),
	}[dc]
}
