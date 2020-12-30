package redis

import "errors"

type Config struct {
	Name           string `json:"name"`
	Addr           string `json:"addr"`
	Password       string `json:"password"`
	MaxIdle        int    `json:"max_idle"`
	MaxActive      int    `json:"max_active"`
	IdleTimeout    int    `json:"idle_timeout"`
	ConnectTimeout int    `json:"connect_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	Database       int    `json:"database"`
	SlowTime       int    `json:"slow_time"`
	Retry          int    `json:"retry"`
}

func (o *Config) isValid() error {
	if o.Name == "" || o.Addr == "" || o.Database < 0 || o.MaxIdle < 100 || o.MaxActive < 100 || o.IdleTimeout < 100 ||
		o.ReadTimeout < 50 || o.WriteTimeout < 50 || o.ConnectTimeout <= 50 || o.SlowTime <= 100 || o.Retry < 0 {
		return errors.New("redis: Name not allowed empty string")
	}
	return nil
}
