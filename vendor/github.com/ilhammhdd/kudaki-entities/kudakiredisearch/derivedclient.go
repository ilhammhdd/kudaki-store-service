package kudakiredisearch

import "github.com/RediSearch/redisearch-go/redisearch"

type DerivedClient int

const (
	CartCheckout DerivedClient = iota
	CartItem
	ItemStorefront
	StorefrontUser
	ProfileUser
	MountainRecomendedgears
	MountainReview
)

func (dc DerivedClient) Name() string {
	return []string{
		"cart_checkout",
		"cart_item",
		"item_storefront",
		"storefront_user",
		"profile_user",
		"mountain_recomendedgears",
		"mountain_review",
	}[dc]
}

func (dc DerivedClient) Schema() *redisearch.Schema {
	return []*redisearch.Schema{
		// CartCheckoutSchema
		CombineSchemas([]*redisearch.Schema{Cart.Schema(), Checkout.Schema()}),
		// CartItemSchema
		CombineSchemas([]*redisearch.Schema{Cart.Schema(), Item.Schema()}).
			AddField(redisearch.NewTextField("cart_item_uuid")).
			AddField(redisearch.NewSortableNumericField("cart_item_id")).
			AddField(redisearch.NewSortableNumericField("cart_item_total_amount")).
			AddField(redisearch.NewSortableNumericField("cart_item_total_price")),
		// ItemStorefront
		CombineSchemas([]*redisearch.Schema{Item.Schema(), Storefront.Schema()}),
		// StorefrontUser
		CombineSchemas([]*redisearch.Schema{Storefront.Schema(), User.Schema()}),
		// ProfileUser
		CombineSchemas([]*redisearch.Schema{Profile.Schema(), User.Schema()}),
		// MountainRecomendedgears
		CombineSchemas([]*redisearch.Schema{Mountain.Schema(), RecomendedGear.Schema()}),
		// MountainReview
		Mountain.Schema().
			AddField(redisearch.NewSortableNumericField("review_id")).
			AddField(redisearch.NewTextField("review_uuid")).
			AddField(redisearch.NewSortableNumericField("review_difficulty")).
			AddField(redisearch.NewTextField("review_comment")),
	}[dc]
}
