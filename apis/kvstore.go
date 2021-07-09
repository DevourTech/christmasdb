package apis

type KVStore interface {
	Get(key interface{}) (interface{}, error)
	Store(key interface{}, value interface{}) error
	Delete(key interface{}) error
}
