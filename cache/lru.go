package cache

import (
	"container/list"
	"github.com/timfame/lru-cache.git/assert"
)

type LRUCache struct {
	storage               map[string]interface{}
	size                  int
	leastRecentlyUsedKeys *list.List
	maxCacheSize          int
}

func NewLRUCache(maxCacheSize int) ICache {
	assert.Assert(maxCacheSize > 0, "Max size of cache cannot be <= 0")
	return &LRUCache{
		storage:               make(map[string]interface{}),
		leastRecentlyUsedKeys: list.New(),
		size:                  0,
		maxCacheSize:          maxCacheSize,
	}
}

func (l *LRUCache) Insert(key string, value interface{}) {
	assert.Assert(value != nil, "Value must not be nil")

	currentSize := l.size

	if _, ok := l.storage[key]; ok {
		l.storage[key] = value
		l.updateKeyInList(key)
		currentSize--
	} else {
		if l.size == l.maxCacheSize {
			lastElem := l.leastRecentlyUsedKeys.Back()
			l.leastRecentlyUsedKeys.Remove(lastElem)
			delete(l.storage, lastElem.Value.(string))
			l.size--
		}
		l.storage[key] = value
		l.leastRecentlyUsedKeys.PushFront(key)
		l.size++
	}

	assert.Assert(l.size == len(l.storage), "Current size of cache must be equal to size of storage")
	assert.Assert(l.size == l.leastRecentlyUsedKeys.Len(), "Current size of cache must be equal to len of keys list")
	assert.Assert(l.size <= l.maxCacheSize, "Size cannot be greater than MaxCacheSize")
}

func (l *LRUCache) Get(key string) interface{} {
	currentSize := l.size
	defer func() {
		assert.Assert(currentSize == l.size, "Get cannot change cache size")
	}()
	if stored, ok := l.storage[key]; ok {
		l.updateKeyInList(key)
		return stored
	}
	return nil
}

func (l *LRUCache) Exists(key string) bool {
	return l.Get(key) != nil
}

func (l *LRUCache) Size() int {
	return l.size
}

func (l *LRUCache) updateKeyInList(key string) {
	l.removeKeyFromList(key)
	l.leastRecentlyUsedKeys.PushFront(key)
}

func (l *LRUCache) removeKeyFromList(key string) {
	currentSize := l.leastRecentlyUsedKeys.Len()
	if currentSize == 0 {
		return
	}
	var elem *list.Element
	for elem = l.leastRecentlyUsedKeys.Front(); elem != nil && elem.Value != key; elem = elem.Next() {}
	l.leastRecentlyUsedKeys.Remove(elem)

	assert.Assert(currentSize - 1 == l.leastRecentlyUsedKeys.Len(), "Remove key must decrease size of list")
}
