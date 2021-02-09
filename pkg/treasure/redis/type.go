package redis

import (
	"github.com/garyburd/redigo/redis"
)

func Bool(reply interface{}, err error) (interface{}, error) {
	return redis.Bool(reply, err)
}

func ByteSlices(reply interface{}, err error) (interface{}, error) {
	return redis.ByteSlices(reply, err)
}

func Bytes(reply interface{}, err error) (interface{}, error) {
	return redis.Bytes(reply, err)
}

func Float64(reply interface{}, err error) (interface{}, error) {
	return redis.Float64(reply, err)
}

func Int(reply interface{}, err error) (interface{}, error) {
	return redis.Int(reply, err)
}

func Int64(reply interface{}, err error) (interface{}, error) {
	return redis.Int64(reply, err)
}

func Int64Map(reply interface{}, err error) (interface{}, error) {
	return redis.Int64Map(reply, err)
}

func IntMap(reply interface{}, err error) (interface{}, error) {
	return redis.IntMap(reply, err)
}

func Ints(reply interface{}, err error) (interface{}, error) {
	return redis.Ints(reply, err)
}

func String(reply interface{}, err error) (interface{}, error) {
	return redis.String(reply, err)
}

func StringMap(reply interface{}, err error) (interface{}, error) {
	return redis.StringMap(reply, err)
}

func Strings(reply interface{}, err error) (interface{}, error) {
	return redis.Strings(reply, err)
}

func Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}
