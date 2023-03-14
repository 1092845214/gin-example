package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCli struct {
	Client *redis.Client
}

func NewRedisCli(addr, password string, db int) (cli *redis.Client, err error) {

	cli = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	err = cli.Ping(context.Background()).Err()
	return
}

func (cli *RedisCli) GetStr(key string) (string, error) {
	r, err := cli.Client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return r, nil
	}
	return r, err
}

func (cli *RedisCli) Set(key, value string, ttl time.Duration) (err error) {
	return cli.Client.SetEX(context.Background(), key, value, ttl).Err()
}

func (cli *RedisCli) Delete(key ...string) (err error) {
	return cli.Client.Del(context.Background(), key...).Err()
}
