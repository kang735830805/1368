package pkg

import (
	"context"
	"errors"
	"sync"
	"time"
)

//3. 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口。

type Request struct {
	key string `json:"key"`
	value string `json:"value"`
	expire int `json:"expire" expire:"300"`
}

type CacheValue struct {
	value string `json:"value"`
	expire int `json:"expire" expire:"300"`
}

type CacheMap struct {
	cacheMap sync.Map
}

var DefaultExpire = 500

func NewCache(key, value string, expire int) *CacheMap {
	cache := &CacheMap{}
	val := &CacheValue{value: value, expire: expire}
	cache.cacheMap.Store(key, val)
	return cache
}

type Cache interface {
	create(ctx context.Context, cache *Request) error
	get(ctx context.Context, key *string) (string ,error)
	update(ctx context.Context, cache *Request) (string, error)
	delete(ctx context.Context, key *string)
}

func (c *CacheMap) expireMap(key string)  {
	val, _ := c.cacheMap.Load(key)
	value := val.(CacheValue)
	temp := time.Duration(value.expire)

	timer := time.NewTimer(temp)
	select {
	case <- timer.C:
		c.cacheMap.Delete(key)
	}
	timer.Stop()
}

func (c *CacheMap) create(ctx context.Context, request *Request) error {
	if request == nil {
		return errors.New("参数为空")
	}
	if request.key == "" {
		return errors.New("key 不能为空")
	}

	c.cacheMap.Store(request.key, &CacheValue{request.value, request.expire})

	go c.expireMap(request.key)

	return nil
}

func (c *CacheMap) get(ctx context.Context, key *string) (string ,error) {
	if val, ok := c.cacheMap.Load(key); !ok {
		return "", nil
	} else {
		value := val.(CacheValue)
		return value.value, nil
	}
}

func (c *CacheMap) update(ctx context.Context, request *Request) (error) {
	if request == nil {
		return errors.New("参数为空")
	}
	if request.key == "" {
		return errors.New("key 不能为空")
	}

	c.cacheMap.Store(request.key, &CacheValue{request.value, request.expire})

	go c.expireMap(request.key)

	return nil
}

func (c *CacheMap) delete(key *string) {
	c.cacheMap.Delete(key)
}