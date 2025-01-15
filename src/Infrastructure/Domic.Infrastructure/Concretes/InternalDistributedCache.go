package InfrastructureConcrete

import (
	"Domic.Domain/Commons/Contracts"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type InternalDistributedCache struct {
	serializer DomainCommonContract.ISerializer
	context    context.Context
	client     *redis.Client
}

func (redis *InternalDistributedCache) Set(object interface{}, key string, timeout time.Duration) bool {
	result := redis.client.Set(redis.context, key, object, timeout)

	if result.Err() != nil {
		return false
	}

	return true
}

func (redis *InternalDistributedCache) Get(key string, target interface{}) error {
	result := redis.client.Get(redis.context, key)

	object, err := result.Result()

	if err != nil {
		return err
	}

	return redis.serializer.Deserialize(object, &target)
}

func (redis *InternalDistributedCache) Delete(key string) bool {
	result := redis.client.Del(redis.context, key)

	if result.Err() != nil {
		return false
	}

	return true
}

func NewDistributedCache(serializer DomainCommonContract.ISerializer, connection string, password string) *InternalDistributedCache {
	return &InternalDistributedCache{
		serializer: serializer,
		context:    context.Background(),
		client: redis.NewClient(&redis.Options{
			Addr:     connection,
			Password: password,
			DB:       0,
		}),
	}
}
