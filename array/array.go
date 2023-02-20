package array

import "fmt"

type Array struct {
	length int
	items  []interface{}
}

func (array *Array) GetLength() (length int) {
	length = array.length
	return
}

func (array *Array) Get(index int) (val interface{}) {
	val = array.items[index]
	return
}

func (array *Array) Push(val interface{}) {
	array.items = append(array.items, val)
	array.length++
}

func (array *Array) Pop() {
	new := Array{}
	for i := 0; i <= array.length; i++ {
		if i < array.length-1 {
			new.Push(array.items[i])
		}
	}
	array.items = new.items
	array.length = new.length
}

func (array *Array) PopByIndex(index int) {
	if index >= array.length {
		fmt.Println("Index must be less than array length")
		return
	}
	new := Array{}
	for i := 0; i < array.length; i++ {
		if i != index {
			new.Push(array.items[i])
		}
	}
	array.items = new.items
	array.length = new.length
}

func (array *Array) Items() (items []interface{}) {
	items = array.items
	return
}

func (array *Array) Reverse() (items []interface{}) {
	new := Array{}
	for i := array.length - 1; i >= 0; i-- {
		new.Push(array.items[i])
	}
	array.items = new.items
	items = array.Items()
	return
}

func (array *Array) Merge(new Array) (items []interface{}) {
	for i := 0; i < new.length; i++ {
		array.Push(new.items[i])
	}
	items = array.Items()
	return
}
