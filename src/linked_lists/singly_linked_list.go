package singlylinkedlist

import (
	"errors"
	"fmt"
)

type link struct {
	key   int
	value *interface{}
	next  *link
}

// Singly Linked List of elements
// head -> item 1 -> item 2 -> ... -> tail
// Can store any type of item.
// Example:
//   list := SinglyLinkedList{}
type SinglyLinkedList struct {
	head *link
	len  int
}

// Insert an item to the end of the list, just before the tail
// Example:
//    list.Insert(2, "item")
func (s *SinglyLinkedList) Insert(key int, value interface{}) {
	if s.head == nil {
		s.head = &link{key: -1, next: &link{key: key, value: &value}}
		s.len++
	} else {
		var current *link = s.head
		for (*current).next != nil {
			if (*current).key == key {
				(*current).value = &value
				return
			} else if (*(*current).next).key > key {
				(*current).next = &link{key: key, value: &value, next: (*current).next}
				s.len++
				return
			}
			current = (*current).next
		}
		(*current).next = &link{key: key, value: &value}
		s.len++
	}
}

// Remove the last item from the list
// Example:
//   el, err := list.Pop()
//   if err != nil {
//     // log error
//   }
func (s *SinglyLinkedList) Pop() (interface{}, error) {
	if s.head == nil {
		return nil, errors.New("Is empty")
	}
	var current *link = s.head
	for (*(*current).next).next != nil {
		current = (*current).next
	}
	var value interface{} = *(*(*current).next).value
	(*current).next = nil
	s.len--
	return value, nil
}

// Remove an item from the front of the list
// Example:
//    el, err := list.Shift()
//   if err != nil {
//     // log error
//   }
func (s *SinglyLinkedList) Shift() (interface{}, error) {
	if s.head == nil {
		return nil, errors.New("Is empty")
	}
	var current *link = (*s.head).next
	(*s.head).next = (*current).next
	s.len--
	return *(*current).value, nil
}

// Iterate over the list,
// Example:
//   for i := range list.Iter() {
//     // do stuff
//   }
func (s *SinglyLinkedList) Iter() chan *interface{} {
	ch := make(chan *interface{})
	go func() {
		var current *link = s.head
		for (*current).next != nil {
			current = (*current).next
			ch <- (*current).value
		}
		close(ch)
	}()
	return ch
}

// Get the length of the list (number of items)
// Example:
//   var l int = list.Len()
func (s *SinglyLinkedList) Len() int {
	return s.len
}

// Remove an element by with its key
// Example:
//   list.Remove(1)
func (s *SinglyLinkedList) Remove(key int) {
	if s.head == nil {
		return
	}
	var current *link = s.head
	for (*current).next != nil {
		if (*(*current).next).key == key {
			(*current).next = (*(*current).next).next
			s.len--
			return
		}
		current = (*current).next
	}
}

// Find the index of a value given the value
// Example:
//   index, err := list.Find("item")
//   if err != nil {
//     // log error
//   }
func (s *SinglyLinkedList) Find(value interface{}) (int, error) {
	if s.head == nil {
		return -1, errors.New("Is empty")
	}
	var current *link = s.head
	for (*current).next != nil {
		current = (*current).next
		if *(*current).value == value {
			return (*current).key, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Cannot find '%v'", value))
}

// Return the list as an array
// Example:
//   arr, err := list.ToArray()
func (s *SinglyLinkedList) ToArray() ([]interface{}, error) {
	if s.head == nil {
		return []interface{}{}, errors.New("Is empty")
	}
	ret := make([]interface{}, s.len)
	var current *link = s.head
	var i int
	for (*current).next != nil {
		current = (*current).next
		ret[i] = *(*current).value
		i++
	}
	return ret, nil
}

// Get a string representation of the linked list (works best with basic types)
// Example:
//   var str string = list.String()
func (s *SinglyLinkedList) String() string {
	var str string = "SinglyLinkedList{ "
	var current *link = s.head
	str += "head->"
	for (*current).next != nil {
		current = (*current).next
		str += fmt.Sprintf("%v->", *(*current).value)
	}
	str += "tail }"
	return str
}

// Prints the linked list to the terminal same as String() above
// Example:
//   list.Println()
func (s *SinglyLinkedList) Println() {
	fmt.Println(s.String())
}
