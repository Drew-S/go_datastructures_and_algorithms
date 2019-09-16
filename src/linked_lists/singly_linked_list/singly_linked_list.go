package singlylinkedlist

import (
	"errors"
	"fmt"
)

type link struct {
	key   int
	value interface{}
	next  *link
}

type keyValuePair struct {
	key int
	value interface{}
}

// SinglyLinkedList holds a list of any mix of items
type SinglyLinkedList struct {
	head *link
	len  int
}

// Insert an item into the list as key, value
func (s *SinglyLinkedList) Insert(key int, value interface{}) {
	if s.head == nil { // if list empty create head, append after head
		s.head = &link{key: -1, next: &link{key: key, value: value}}
		s.len++

	} else { // find the correct spot
		var current *link = s.head
		for current.next != nil {
			if (*current.next).key == key { // keys match update value
				(*current.next).value = value
				return

			} else if (*(*current).next).key > key { // next key larger, place before
				current.next = &link{key: key, value: value, next: current.next}
				s.len++
				return
			}
			current = (*current).next
		}
		(*current).next = &link{key: key, value: value}
		s.len++
	}
}

// Pop, removes and returns the last item from the list
func (s *SinglyLinkedList) Pop() (interface{}, error) {
	if s.head == nil {
		return nil, errors.New("Is empty")
	}
	var current *link = s.head
	for current.next.next != nil {
		current = current.next
	}
	var value interface{} = current.next.value
	current.next = nil
	s.len--
	return value, nil
}

// Shift, removes and returns the first item from the list
func (s *SinglyLinkedList) Shift() (interface{}, error) {
	if s.head == nil {
		return nil, errors.New("Is empty")
	}
	var value interface{} = s.head.next.value
	s.head.next = s.head.next.next
	s.len--
	return value, nil
}

// Iter loops over the entire list
func (s *SinglyLinkedList) Iter() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		var current *link = s.head
		for current.next != nil {
			current = current.next
			ch <- &current.value
		}
		close(ch)
	}()
	return ch
}

// Len returns the number of items in the list
func (s *SinglyLinkedList) Len() int {
	return s.len
}

// Remove an item from the list by its key
func (s *SinglyLinkedList) Remove(key int) (interface{}, error) {
	if s.head == nil {
		return nil, errors.New("Empty list")
	}
	var current *link = s.head
	for current.next != nil {
		if current.next.key == key {
			var value interface{} = current.next.value
			current.next = current.next.next
			s.len--
			return value, nil
		}
		current = current.next
	}
	return nil, errors.New("Value not found")
}

// Find the first matching key for a value in the list
func (s *SinglyLinkedList) Find(value interface{}) (int, error) {
	if s.head == nil {
		return -1, errors.New("Is empty")
	}
	var current *link = s.head
	for current.next != nil {
		current = current.next
		if current.value == value {
			return current.key, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Cannot find '%v'", value))
}

// ToArray returns an array of items in order of head->item->...->tail
func (s *SinglyLinkedList) ToArray() ([]keyValuePair, error) {
	if s.head == nil {
		return []keyValuePair{}, errors.New("Is empty")
	}
	ret := make([]keyValuePair, s.len)
	var current *link = s.head
	var i int
	for current.next != nil {
		current = current.next
		ret[i] = keyValuePair{
			key: current.key,
			value: current.value,
		}
		i++
	}
	return ret, nil
}

// Sprint returns a string representation of the list
func (s *SinglyLinkedList) Sprint() string {
	var str string = "SinglyLinkedList{ "
	var current *link = s.head
	str += "head->"
	for current.next != nil {
		current = current.next
		str += fmt.Sprintf("%v->", current.value)
	}
	str += "tail }"
	return str
}
