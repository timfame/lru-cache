package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	lru := NewLRUCache(5)
	key, value := "test", 100500
	lru.Insert(key, value)

	assert.True(t, lru.Exists(key))
	assert.Equal(t, value, lru.Get(key))
	assert.Equal(t, 1, lru.Size())
}

func TestMultipleInsertBeforeLimit(t *testing.T) {
	size := 10
	lru := NewLRUCache(size)
	keys := make([]string, 0, size)
	values := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		key, value := fmt.Sprintf("test%d", i), i
		keys = append(keys, key)
		values = append(values, value)
		lru.Insert(key, value)
	}

	assert.Equal(t, size, lru.Size())
	for i := 0; i < size * 10; i++ {
		pos := rand.Intn(size)
		key, value := keys[pos], values[pos]
		assert.True(t, lru.Exists(key))
		assert.Equal(t, value, lru.Get(key))
	}
}

func TestMultipleInsertAfterLimit(t *testing.T) {
	size := 10
	afterLimit := 5
	lru := NewLRUCache(size)
	keys := make([]string, 0, size + afterLimit)
	values := make([]interface{}, 0, size + afterLimit)
	for i := 0; i < size + afterLimit; i++ {
		key, value := fmt.Sprintf("test%d", i), i
		keys = append(keys, key)
		values = append(values, value)
		lru.Insert(key, value)
	}

	assert.Equal(t, size, lru.Size())
	for i := 0; i < size * 10; i++ {
		pos := rand.Intn(size)
		key, value := keys[pos], values[pos]
		if pos < afterLimit {
			assert.False(t, lru.Exists(key))
			assert.Nil(t, lru.Get(key))
		} else {
			assert.True(t, lru.Exists(key))
			assert.Equal(t, value, lru.Get(key))
		}
	}
}

func TestUpdate(t *testing.T) {
	size := 10
	lru := NewLRUCache(size)
	keys := make([]string, 0, size)
	values := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		key, value := "test0", 0
		keys = append(keys, key)
		values = append(values, value)
		lru.Insert(key, value)
	}

	assert.Equal(t, 1, lru.Size())
	assert.True(t, lru.Exists(keys[0]))
	assert.Equal(t, values[0], lru.Get(keys[0]))
}
