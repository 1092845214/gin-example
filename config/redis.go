package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/yangkaiyue/gin-exp/global"
)

type RedisCli struct {
	cli *redis.Client
}

func InitRedis() (err error) {

	rds := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	err = rds.Ping(context.Background()).Err()
	global.RDS = rds
	return
}

func (cli *RedisCli) GetStr(key string) (string, error) {
	r, err := global.RDS.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return r, nil
	}
	return r, err
}

func (cli *RedisCli) Set(key, value string) (err error) {
	return global.RDS.Set(context.Background(), key, value, viper.GetDuration("redis.ttl")).Err()
}

func (cli *RedisCli) Delete(key ...string) (err error) {
	return global.RDS.Del(context.Background(), key...).Err()
}
