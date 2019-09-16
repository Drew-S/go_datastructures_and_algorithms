package doublylinkedlist

import (
	"errors"
	"fmt"
)

type keyValuePair struct {
	key   int
	value interface{}
}

type link struct {
	key   int
	value interface{}
	next  *link
	prev  *link
}

// DoublyLinkedList holds a list of any mix of items
type DoublyLinkedList struct {
	head *link
	tail *link
	len  int
}

// Insert an item into the list as key, value
func (d *DoublyLinkedList) Insert(key int, value interface{}) {
	if d.head == nil { // if list empty create head, append after head
		d.tail = &link{key: -1}
		d.head = &link{key: -1, prev: d.tail, next: &link{key: key, value: value, next: d.tail}}
		d.head.next.prev = d.head
		d.tail.next = d.head
		d.tail.prev = d.head.next
		d.len++

	} else { // find the correct spot
		var current *link = d.head.next
		for *current != *d.tail {
			if current.key == key { // keys match update value
				current.value = value
				return

			} else if current.key > key { // current key larger, place before
				var prev *link = current.prev
				prev.next = &link{key: key, value: value, prev: prev, next: current}
				current.prev = prev.next
				d.len++
				return

			}
			current = current.next
		}
		current = current.prev
		current.next = &link{key: key, value: value, prev: current, next: d.tail}
		d.tail.prev = current.next
		d.len++
	}
}

// Pop, removes and returns the last item from the list
func (d *DoublyLinkedList) Pop() (interface{}, error) {
	if d.tail == nil {
		return nil, errors.New("Is empty")
	}
	var value interface{} = d.tail.prev.value
	var prev *link = d.tail.prev.prev
	prev.next = d.tail
	d.tail.prev = prev
	d.len--
	return value, nil
}

// Shift, removes and returns the first item from the list
func (d *DoublyLinkedList) Shift() (interface{}, error) {
	if d.head == nil {
		return nil, errors.New("Is empty")
	}
	var value interface{} = d.head.next.value
	var next *link = d.head.next.next
	d.head.next = next
	next.prev = d.head
	d.len--
	return value, nil
}

// Iter loops over the entire list, head to tail
func (d *DoublyLinkedList) Iter() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		var current *link = d.head
		for *current.next != *d.tail {
			current = current.next
			ch <- &current.value
		}
		close(ch)
	}()
	return ch
}

// IterReverse loops over the entire list, head to tail
func (d *DoublyLinkedList) IterReverse() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		var current *link = d.tail
		for *current.prev != *d.head {
			current = current.prev
			ch <- &current.value
		}
		close(ch)
	}()
	return ch
}

// Len returns the number of items in the list
func (d *DoublyLinkedList) Len() int {
	return d.len
}

// Remove an item from the list by its key
func (d *DoublyLinkedList) Remove(key int) (interface{}, error) {
	if d.head == nil {
		return nil, errors.New("Empty list")
	}
	var current *link = d.head
	for *current.next != *d.tail {
		if current.next.key == key {
			var value interface{} = current.next.value
			current.next = current.next.next
			d.len--
			return value, nil
		}
		current = current.next
	}
	return nil, errors.New("Value not found")
}

// Find the first matching key for a value in the list
func (d *DoublyLinkedList) Find(value interface{}) (int, error) {
	if d.head == nil {
		return -1, errors.New("Is empty")
	}
	var current *link = d.head
	for *current.next != *d.tail {
		current = current.next
		if current.value == value {
			return current.key, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Cannot find '%v'", value))
}

// ToArray returns an array of items in order of head<->item<->...<->tail
func (d *DoublyLinkedList) ToArray() ([]keyValuePair, error) {
	if d.head == nil {
		return []keyValuePair{}, errors.New("Is empty")
	}
	ret := make([]keyValuePair, d.len)
	var current *link = d.head
	var i int
	for *current.next != *d.tail {
		current = current.next
		ret[i] = keyValuePair{
			key:   current.key,
			value: current.value,
		}
		i++
	}
	return ret, nil
}

// Sprint returns a string representation of the list
func (d *DoublyLinkedList) Sprint() string {
	var str string = "DoublyLinkedList{ "
	var current *link = d.head
	str += "head<->"
	for *current.next != *d.tail {
		current = current.next
		str += fmt.Sprintf("%v<->", current.value)
	}
	str += "tail }"
	return str
}
