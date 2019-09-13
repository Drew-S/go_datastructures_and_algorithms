package doublylinkedlist

import (
	"errors"
	"fmt"
)

type link struct {
	key   int
	value *interface{}
	next  *link
	prev  *link
}

// Doubly Linked List of elements
// head <-> item 1 <-> item 2 <-> ... <-> tail
// Can store any type of item.
// Example:
//   list := DoublyLinkedList{}
type DoublyLinkedList struct {
	head *link
	tail *link
	len  int
}

// Insert an item to the end of the list, just before the tail
// Example:
//    list.Insert(2, "item")
func (d *DoublyLinkedList) Insert(key int, value interface{}) {
	if d.head == nil { // if list empty create head, append after head
		d.tail = &link{key: -1}
		d.head = &link{key: -1, prev: d.tail, next: &link{key: key, value: &value, next: d.tail}}
		d.head.next.prev = d.head
		d.tail.next = d.head
		d.tail.prev = d.head.next
		d.len++

	} else { // find the correct spot
		var current *link = d.head.next
		for *current != *d.tail {
			if current.key == key { // keys match update value
				current.value = &value
				return

			} else if current.key > key { // current key larger, place before
				var prev *link = current.prev
				prev.next = &link{key: key, value: &value, prev: prev, next: current}
				current.prev = prev.next
				d.len++
				return

			}
			current = current.next
		}
		current = current.prev
		current.next = &link{key: key, value: &value, prev: current, next: d.tail}
		d.tail.prev = current.next
		d.len++
	}
}

// Remove the last item from the list
// Example:
//   el, err := list.Pop()
//   if err != nil {
//     // log error
//   }
func (d *DoublyLinkedList) Pop() (interface{}, error) {
	if d.tail == nil {
		return nil, errors.New("Is empty")
	}
	var value interface{} = *d.tail.prev.value
	var prev *link = d.tail.prev.prev
	prev.next = d.tail
	d.tail.prev = prev
	d.len--
	return value, nil
}

// Remove an item from the front of the list
// Example:
//    el, err := list.Shift()
//   if err != nil {
//     // log error
//   }
func (d *DoublyLinkedList) Shift() (interface{}, error) {
	if d.head == nil {
		return nil, errors.New("Is empty")
	}
	var value interface{} = *d.head.next.value
	var next *link = d.head.next.next
	d.head.next = next
	next.prev = d.head
	d.len--
	return value, nil
}

// Iterate over the list, head to tail
// Example:
//   for i := range list.Iter() {
//     // do stuff
//   }
func (d *DoublyLinkedList) Iter() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		var current *link = d.head
		for *current.next != *d.tail {
			current = current.next
			ch <- current.value
		}
		close(ch)
	}()
	return ch
}

// Iterate over the list, tail to head
// Example:
//   for i := range list.IterReverse() {
//     // do stuff
//   }
func (d *DoublyLinkedList) IterReverse() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		var current *link = d.tail
		for *current.prev != *d.head {
			current = current.prev
			ch <- current.value
		}
		close(ch)
	}()
	return ch
}

// Get the length of the list (number of items)
// Example:
//   var l int = list.Len()
func (d *DoublyLinkedList) Len() int {
	return d.len
}

// Remove an element by with its key
// Example:
//   list.Remove(1)
func (d *DoublyLinkedList) Remove(key int) (interface{}, error) {
	if d.head == nil {
		return nil, errors.New("Empty list")
	}
	var current *link = d.head
	for *current.next != *d.tail {
		if current.next.key == key {
			var value interface{} = *current.next.value
			current.next = current.next.next
			d.len--
			return value, nil
		}
		current = current.next
	}
	return nil, errors.New("Value not found")
}

// Find the index of a value given the value
// Example:
//   index, err := list.Find("item")
//   if err != nil {
//     // log error
//   }
func (d *DoublyLinkedList) Find(value interface{}) (int, error) {
	if d.head == nil {
		return -1, errors.New("Is empty")
	}
	var current *link = d.head
	for *current.next != *d.tail {
		current = current.next
		if *current.value == value {
			return current.key, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Cannot find '%v'", value))
}

// Return the list as an array
// Example:
//   arr, err := list.ToArray()
func (d *DoublyLinkedList) ToArray() ([]*interface{}, error) {
	if d.head == nil {
		return []*interface{}{}, errors.New("Is empty")
	}
	ret := make([]*interface{}, d.len)
	var current *link = d.head
	var i int
	for *current.next != *d.tail {
		current = current.next
		ret[i] = current.value
		i++
	}
	return ret, nil
}

// Get a string representation of the linked list (works best with basic types)
// Example:
//   var str string = list.String()
func (d *DoublyLinkedList) String() string {
	var str string = "DoublyLinkedList{ "
	var current *link = d.head
	str += "head<->"
	for *current.next != *d.tail {
		current = current.next
		str += fmt.Sprintf("%v<->", *current.value)
	}
	str += "tail }"
	return str
}
