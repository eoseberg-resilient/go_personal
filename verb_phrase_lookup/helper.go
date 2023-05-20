package main

import (
	"errors"
	"fmt"
	"sort"
)

var ValueMissingError = errors.New("Value does not exist")

type Map[Key ~string, Value ValueType] interface {
	Push(Key, Value)
	//Pop(Key) (Value, error)
	Get(Key) Value
	Has(Key) bool
	Keys() []string
	Len() int
	AsString() string
}

func (mmap *MapImpl[Key, Value]) Push(k Key, v Value) {
	_, has := mmap.internalMap[k]
	if !has {
		mmap.keys = append(mmap.keys, k)
	}
	newIndx := len(mmap.keys) - 1
	mmap.internalMap[k] = indexedValue[Key, Value] {
		Value: v,
		Index: newIndx,
	}
}

func (mmap *MapImpl[Key, Value]) Get(k Key) Value {
	if !mmap.Has(k) {
		panic(ValueMissingError)
	}
	val, _ := mmap.internalMap[k]
	return val.Value
}

func (mmap *MapImpl[Key, Value]) Has(k Key) bool {
	_, has := mmap.internalMap[k]
	return has
}

func (mmap MapImpl[Key, Value]) Keys() []string {
	keys := make([]string,0)
	for k, _ := range mmap.internalMap {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	return keys
}

func (mmap *MapImpl[Key, Value]) Len() int {
	return len(mmap.keys)
}

func (mmap *MapImpl[Key, Value]) AsString() string {
	ret := "";
	for k, v := range mmap.internalMap {
		ret += fmt.Sprint("K: " , k, ",  Val: ", v)
		ret += "\n"
	}
	return ret
}

type ValueType interface{}
type indexedValue[Key ~string, Value ValueType] struct {
	Value Value
	Index int
}
type MapImpl[Key ~string, Value ValueType] struct {
	internalMap map[Key]indexedValue[Key, Value]
	keys 		[]Key
}

func NewMap[Key ~string, Value ValueType]() Map[Key, Value] {
	return &MapImpl[Key, Value]{
		internalMap: map[Key]indexedValue[Key, Value]{},
		keys: []Key{},
	}
}



func ToArray(m map[int]string) *[]string {
	ret := make([]string, 1)
	for _, v := range m {
		ret = append(ret, v)
	}
	return &ret
}
