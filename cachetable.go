package gocache

import (
	"log"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex

	name string
	items map[interface{}]*CacheItem

	cleanupTimer *time.Timer

	cleanupInterval time.Duration

	logger *log.Logger

	loadData func(key interface{}, args ...interface{}) *CacheItem

	addItem []func(item *CacheItem)

	aboutToDeleteItem []func(item *CacheItem)
}

func (table *CacheTable) Count() int {
	table.RLock()
	defer table.RUnlock()
	return len(table.items())
}

func (table *CacheTable) Foreach(trans func(key interface{}, item *CacheItem)) {
	table.RLock()
	defer table.RUnlock()
	for k,v range table.items {
		trans(k,v)
	}
}


func (table *CacheTable) SetDataLoader(f func(interface{}, ...interface{}) *CacheItem)   {
	table.Lock()
	defer table.UnLock()
	table.loadData = f
}


func (table *CacheTable)  SetAddedItemCallback(f func(*CacheItem))  {
	if len(table.addedItem) >0 {
		table.RemoveAddedItemCallbacks()
	}
	table.Lock()
	defer table.UnLock()
	table.addedItem = append(table.addedItem,f)
}


func (table *CacheTable)  RemoveAddedItemCallbacks() {
	table.Lock()
	defer table.UnLock()
	table.addedItem = nil
}

func (table *CacheTable)   SetAboutToDeleteItemCallback(f func(*CacheItem))  {
	if len(table.aboutToDeleteItem) > 0 {
		table.RemoveAboutToDeleteItemCallback()
	}
	table.Lock()
	defer table.UnLock()
	table.aboutToDeleteItem = append(table.aboutToDeleteItem,f)
}