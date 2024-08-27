package redis

import (
	"context"
	"fmt"
	"strconv"
	"user_service/userproto"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisClient{Client: rdb}
}

func (r *RedisClient) SetHash(key string, values map[string]interface{}) error {
	return r.Client.HMSet(ctx, key, values).Err()
}

func (r *RedisClient) VerifyEmail(ctx context.Context, email string, usercode int64) (*userproto.User, error) {
	code, err := r.Client.HGet(ctx, email, "code").Int64()
	if err != nil {
		fmt.Println(code, err)
		return nil, fmt.Errorf("error HGET:%v", err)
	}
	if code == usercode {
		result, err := r.Client.HGetAll(ctx, email).Result()
		if err != nil {
			return nil, fmt.Errorf("error HGETALL: %v", err)
		}

		age, err := strconv.ParseInt(result["age"], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("error converting age to int32: %v", err)
		}

		return &userproto.User{
			Username: result["userName"],
			Age:      int32(age), 
			Email:    result["email"],
			Password: result["password"],
		}, nil
	}
	return nil, err
}

func (r *RedisClient) GetHash(key string) (map[string]string, error) {
	return r.Client.HGetAll(ctx, key).Result()
}

func (r *RedisClient) Delete(key string) error {
	return r.Client.Del(ctx, key).Err()
}
