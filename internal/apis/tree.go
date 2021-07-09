package apis

// Tree is an interface that specifies CRUD operations for DBMS
type Tree interface {
	Search(key uint64) (interface{}, error)
	Insert(key uint64, value interface{}) error
	Delete(key uint64) error
}
