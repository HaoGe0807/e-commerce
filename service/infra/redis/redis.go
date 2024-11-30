package redis

import (
	"e-commerce/service/infra/log"
	"gopkg.in/redis.v5"
)

var redisClients map[string]*redis.Client

func init() {
	log.Debug("Init all Redis databases")
	redisClients = make(map[string]*redis.Client)
}

func NewClient(name, host, port, password string, dbnum int) *redis.Client {
	addr := host + ":" + port
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbnum,
	})
	redisClients[name] = client
	log.Info("Connect to redis database successfully database:", addr, " dbnum: ", dbnum)

	return client
}

func GetClient(name string) *redis.Client {
	return redisClients[name]
}

func Close() {
	for _, client := range redisClients {
		client.Close()
	}
}