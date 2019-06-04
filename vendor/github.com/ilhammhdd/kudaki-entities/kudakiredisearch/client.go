package kudakiredisearch

import "github.com/RediSearch/redisearch-go/redisearch"

type RedisClient interface {
	Name() string
	Schema() *redisearch.Schema
}

type Client int

const (
	Cart Client = iota
	Checkout
	Item
	Storefront
	User
	Profile
	Mountain
	RecomendedGear
)

func (c Client) Name() string {
	return []string{
		"cart",
		"checkout",
		"item",
		"storefront",
		"user",
		"profile",
		"mountain",
		"recomended_gear",
	}[c]
}

func (c Client) Schema() *redisearch.Schema {
	return []*redisearch.Schema{
		// Cart
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("cart_id")).
			AddField(redisearch.NewTextField("cart_uuid")).
			AddField(redisearch.NewSortableNumericField("cart_total_price")).
			AddField(redisearch.NewSortableNumericField("cart_total_items")).
			AddField(redisearch.NewNumericField("cart_open")),
		// Checkout
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("checkout_id")).
			AddField(redisearch.NewTextField("checkout_uuid")).
			AddField(redisearch.NewSortableNumericField("checkout_issued_at")),
		// Item
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("item_id")).
			AddField(redisearch.NewTextField("item_uuid")).
			AddField(redisearch.NewTextField("item_name")).
			AddField(redisearch.NewSortableNumericField("item_amount")).
			AddField(redisearch.NewTextField("item_unit")).
			AddField(redisearch.NewSortableNumericField("item_price")).
			AddField(redisearch.NewTextField("item_description")).
			AddField(redisearch.NewTextField("item_photo")).
			AddField(redisearch.NewSortableNumericField("item_rating")),
		// Storefront
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("storefront_id")).
			AddField(redisearch.NewTextField("storefront_uuid")).
			AddField(redisearch.NewSortableNumericField("storefront_total_item")).
			AddField(redisearch.NewSortableNumericField("storefront_rating")),
		// User
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("user_id")).
			AddField(redisearch.NewTextField("user_uuid")).
			AddField(redisearch.NewTextField("user_email")).
			AddField(redisearch.NewTextField("user_password")).
			AddField(redisearch.NewTextField("user_token")).
			AddField(redisearch.NewTagField("user_role")).
			AddField(redisearch.NewTextField("user_phone_number")).
			AddField(redisearch.NewTagField("user_account_type")),
		// Profile
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("profile_id")).
			AddField(redisearch.NewTextField("profile_uuid")).
			AddField(redisearch.NewTextField("profile_full_name")).
			AddField(redisearch.NewTextField("profile_photo")).
			AddField(redisearch.NewSortableNumericField("profile_reputation")),
		// Mountain
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("mountain_id")).
			AddField(redisearch.NewTextField("mountain_uuid")).
			AddField(redisearch.NewTextField("mountain_name")).
			AddField(redisearch.NewSortableNumericField("mountain_height")).
			AddField(redisearch.NewSortableNumericField("mountain_latitude")).
			AddField(redisearch.NewSortableNumericField("mountain_longitude")).
			AddField(redisearch.NewSortableNumericField("mountain_difficulty")).
			AddField(redisearch.NewTextField("mountain_description")),
		// RecomendedGears
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("recomended_gear_id")).
			AddField(redisearch.NewTextField("recomended_gear_uuid")).
			AddField(redisearch.NewTextField("recomended_gear_photo")).
			AddField(redisearch.NewTextField("recomended_gear_name")).
			AddField(redisearch.NewSortableNumericField("recomended_gear_total")).
			AddField(redisearch.NewTextField("recomended_gear_unit")).
			AddField(redisearch.NewSortableNumericField("recomended_gear_person")).
			AddField(redisearch.NewTextField("recomended_gear_description")),
	}[c]
}
