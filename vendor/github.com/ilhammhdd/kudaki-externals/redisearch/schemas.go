package redisearch

import (
	"github.com/RediSearch/redisearch-go/redisearch"
)

func CombineSchemas(schemas []*redisearch.Schema) *redisearch.Schema {
	for i := 1; i < len(schemas); i++ {
		for _, field := range schemas[i].Fields {
			schemas[0].AddField(field)
		}
	}

	return schemas[0]
}
