package redis

import (
	"fmt"
	"sync"
)

type Manager struct {
	redisMap map[string]*Redis
	sync.RWMutex
}

func NewManager(c []Config) (*Manager, error) {
	m := &Manager{
		redisMap: map[string]*Redis{},
	}
	for _, config := range c {
		r, err := NewRedis(&config)
		if err == nil {
			m.redisMap[config.Name] = r
		} else {
			return nil, fmt.Errorf("redis: init redis: %s error: %s", config.Name, err)
		}
	}
	return m, nil
}

func (m *Manager) Add(name string, r *Redis) {
	m.Lock()
	defer m.Unlock()
	m.redisMap[name] = r
}

func (m *Manager) Get(name string) *Redis {
	m.RLock()
	defer m.RUnlock()
	return m.redisMap[name]
}
