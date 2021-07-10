package bplustree

import (
	"github.com/DevourTech/devourKV/internal/apis"
	"github.com/DevourTech/devourKV/internal/errors"
)

const (
	//indexFilePath  = "./db/index.db"
	//recordFilePath = "./db/records.db"
)

// Node is the structure used to represent a BPlusTree node
type Node struct {
	// keys denote the list of keys
	keys []uint64

	// children is the list of pointers that point to the child nodes
	children []*Node

	// n is the number of children
	n int

	// sibling is the pointer to the next node. This is only applicable for a leaf node
	sibling *Node

	// leaf is the flag that denotes whether a node is a leaf
	leaf bool

	// offset is the position (or offset) of a node stored in the file
	offset int
}

// BPlusTree is the structure used to represent a B+ Tree
type BPlusTree struct {
	// root is the pointer to the root node
	root *Node

	// t is the degree of the tree
	t int
}

// Construct creates a tree either from the file or an entirely new tree
func Construct(degree int) apis.Tree {
	// TODO: Check if an entry exists in the file for the root.
	r := &Node{
		keys:     nil,
		children: nil,
		n:        0,
		sibling:  nil,
		leaf:     true,
		offset:   0,
	}
	return &BPlusTree{root: r, t: degree}
}

// Search returns the value for a specified key
func (tree *BPlusTree) Search(key uint64) (interface{}, error) {
	return tree.search(tree.root, key)
}

// search searches for key in node x
func (tree *BPlusTree) search(x *Node, key uint64) (interface{}, error) {
	if !x.leaf {
		i := 0
		for i < x.n && key > x.keys[i] {
			i++
		}
		return tree.search(x.children[i], key)
	}

	i := 0
	for i < x.n && key > x.keys[i] {
		i++
	}

	// i reaches the end of the keys OR key doesn't match the key at i => Key not found
	if i == x.n || key != x.keys[i] {
		return nil, errors.ErrKeyNotFound
	}

	// TODO: Key found. Get the value from the file. Decode and return it.
	return nil, nil
}

// Insert inserts a key in the tree
func (tree *BPlusTree) Insert(key uint64, value interface{}) error {
	panic("implement me")
}

// Delete removes the key from the tree
func (tree *BPlusTree) Delete(key uint64) error {
	panic("implement me")
}
