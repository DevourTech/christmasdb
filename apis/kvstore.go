package apis

type KVStore interface {
	Get(key string) (interface{}, error)
	Store(key string, value interface{}) error
	Delete(key string) error
}
