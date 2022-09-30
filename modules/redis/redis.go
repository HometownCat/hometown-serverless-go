package redis

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ctx         context.Context = context.Background()
	Rdb                         = GetClient()
	ExpiredTime time.Duration   = time.Hour * 24 * 30
)

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER_HOST"),
		Password: "",
		DB:       0,
	})
	return client
}

func GetContext() *context.Context {
	return &ctx
}

func GetData(key *string, getType *string) (*interface{}, error) {
	var result interface{}
	var err error

	switch *getType {
	case "string":
		result, err = Rdb.Do(ctx, "get", *key).Text()
	case "int":
		result, err = Rdb.Do(ctx, "get", *key).Int()
	case "int64":
		result, err = Rdb.Do(ctx, "get", *key).Int64()
	case "slice":
		result, err = Rdb.Do(ctx, "get", *key).Slice()
	case "bool":
		result, err = Rdb.Do(ctx, "get", *key).Bool()
	default:
		result, err = Rdb.Do(ctx, "get", *key).Result()
	}

	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

func SetData(key *string, value *string, availableTime *time.Duration) error {
	_, err := Rdb.Set(ctx, *key, *value, *availableTime).Result()
	return err
}

func Close() {
	if Rdb != nil {
		Rdb.Close()
	}
}
