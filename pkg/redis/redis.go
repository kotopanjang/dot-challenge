package redis

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

type IRedisStorage interface {
	Set(key string, dataset any) error
	GetByKey(key string, convertValue any) error
	Del(key string) error
}

type RedisStorage struct {
	client *redis.Client
}

func NewRedis(addr, password string) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	result := client.Ping(context.Background())
	_, err := result.Result()
	if err != nil {
		log.Fatal(err)
	}

	return &RedisStorage{
		client: client,
	}
}

func NewTestRedis() *RedisStorage {
	redisServer, _ := miniredis.Run()
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	return &RedisStorage{
		client: redisClient,
	}
}

func (r *RedisStorage) Set(key string, dataset any) error {
	data, err := json.Marshal(dataset)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, data, 0).Err()
}

func (r *RedisStorage) GetByKey(key string, convertValue any) error {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println(err.Error())
			return err
		}
		convertValue = nil
		return errors.New("no data in redis")
	}

	if val != "" {
		err = json.Unmarshal([]byte(val), &convertValue)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

func (r *RedisStorage) Del(key string) error {
	return r.client.Del(context.Background(), key).Err()
}
