package xcache

import (
	"errors"
	"golang.org/x/sync/singleflight"
	"sync/atomic"
)

var (
	ErrUnrecognizedKid = errors.New("unrecognized kid")
)

type CacheClient struct {
	keys  atomic.Value // map[string]string
	group singleflight.Group
}

func NewCacheClient() *CacheClient {
	c := &CacheClient{}
	m := make(map[string]string)
	c.keys.Store(m)
	return c
}

func (c *CacheClient) Get(kid string, fetch func() (map[string]string, error)) (string, error) {
	key, ok := c.query(kid) //缓存里有直接返回
	if ok {
		return key, nil
	}
	if _, err, _ := c.group.Do(kid, func() (interface{}, error) { //缓存里没有执行 fetch 进行加载
		var keys map[string]string
		var err error
		keys, err = fetch()
		if err != nil {
			return nil, err
		}
		c.keys.Store(keys)
		return nil, nil
	},
	); err != nil {
		return "", err
	}
	key, ok = c.query(kid) //再次去缓存获取一次
	if !ok {
		return "", ErrUnrecognizedKid
	}
	return key, nil
}

func (c *CacheClient) query(kid string) (string, bool) {
	keys := c.keys.Load().(map[string]string)
	key, ok := keys[kid]
	return key, ok
}
