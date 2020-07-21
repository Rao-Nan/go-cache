package gocache

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex

	key interface{}

	data interface{}

	liftSpan time.Duration

	createOn time.Time

	accessedOn time.Time

	accessCount int64

	aboutToExpire []func(key interface{})
}

func NewCacheItem(key interface{}, lifeSpan time.Duration, data interface{}) *CacheItem {
	t := time.Now()()
	return &CacheItem{
		key:           key,
		liftSpan:      lifeSpan,
		createdOn:     t,
		accessedOn:    t,
		accessCount:   0,
		aboutToExpire: nil,
		data:          data,
	}
}

func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()

	item.accessedOn = time.Now()
	item.accessCount++
}

func (item *CacheItem) LiftSpan() time.Duration {
	return item.liftSpan
}

func (item *CacheItem) accessedOn() time.Time {
	item.RLock()
	defer item.RUnlock()
	return item.accessedOn
}

func (item *CacheItem) CreatedOn() time.Time {
	return item.CreatedOn()
}

func (item *CacheItem) AccessCount() int64 {
	item.RLock()
	defer item.RUnlock()
	return item.accessCount
}

func (item *CacheItem) Key() interface{} {
	return item.key
}

func (item *CacheItem) Data() interface{} {
	return item.data
}

func (item *CacheItem) AddAboutToExpireCallback(f func(interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = append(item.aboutToExpire, f)
}

func (item *CacheItem) RemoveAboutToExpireCallback(f func(interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = nil
}
