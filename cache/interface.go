package cache

type ICache interface {
	Insert(key string, value interface{})
	Get(key string) interface{}
	Exists(key string) bool
	Size() int
}
