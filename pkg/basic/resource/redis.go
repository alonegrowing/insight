package resource

import (
	"errors"

	"github.com/alonegrowing/insight/pkg/config"
	"github.com/alonegrowing/insight/pkg/sea/redis"
)

var defaultRedis map[string]*redis.Redis

func init() {
	NewRedis(config.ServiceConfig.Redis)
}

func NewRedis(configs []redis.Config) {
	if defaultRedis == nil {
		defaultRedis = make(map[string]*redis.Redis)
	}
	for _, conf := range configs {
		client, err := redis.NewRedis(&conf)
		if err != nil || client == nil {
			continue
		}
		defaultRedis[conf.Name] = client
	}
	return
}

func GetRedis(service string) (*redis.Redis, error) {
	if client, ok := defaultRedis[service]; ok {
		return client, nil
	}
	return nil, errors.New("RedisClientNotInit")
}
