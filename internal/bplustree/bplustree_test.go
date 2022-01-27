package bplustree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	n := Node{}
	typ := reflect.TypeOf(n)
	fmt.Println(typ.Size())
}
