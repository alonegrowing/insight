package redis

import (
	"errors"
	"sync"

	"github.com/garyburd/redigo/redis"
)

type Pipeline struct {
	conn redis.Conn
	*sync.Mutex
	isClose bool
}

func (r *Redis) NewPipelining() (*Pipeline, error) {
	p := &Pipeline{}
	client := r.pool.Get()
	err := client.Err()
	if err != nil {
		return nil, err
	}
	p.conn = client
	p.Mutex = &sync.Mutex{}
	return p, nil
}

func (p *Pipeline) Send(cmd string, args ...interface{}) error {
	p.Lock()
	defer p.Unlock()
	if p.isClose {
		return errors.New("PipeliningClosed")
	}
	return p.conn.Send(cmd, args...)
}

func (p *Pipeline) Flush() error {
	p.Lock()
	defer p.Unlock()
	if p.isClose {
		return errors.New("PipeliningClosed")
	}
	return p.conn.Flush()
}

func (p *Pipeline) Receive() (reply interface{}, err error) {
	p.Lock()
	defer p.Unlock()
	if p.isClose {
		return nil, errors.New("PipeliningClosed")
	}
	return p.conn.Receive()
}

func (p *Pipeline) Close() error {
	p.Lock()
	defer p.Unlock()
	p.isClose = true
	return p.conn.Close()
}
