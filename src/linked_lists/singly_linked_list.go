package singlylinkedlist

import (
	"fmt"
	"errors"
)

type link struct {
	object *interface{}
	next *link
}

// Singly Linked List of elements
// head -> item 1 -> item 2 -> ... -> tail
// Can store any type of item.
// Example:
//   list := SinglyLinkedList{}
type SinglyLinkedList struct {
	head *link
	len int
}

// Append an item to the end of the list, just before the tail
// Example:
//    list.Append("item")
func (s *SinglyLinkedList) Append(object interface{}) {
	if s.head == nil {
		s.head = &link{ next: &link{ object: &object }}
	} else {
		var current *link = s.head
		for (*current).next != nil {
			current = (*current).next
		}
		(*current).next = &link{ object: &object }
	}
	s.len++
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
	var prev *link = current
	for (*current).next != nil {
		prev = current
		current = (*current).next
	}
	(*prev).next = nil
	s.len--
	return *(*current).object, nil
}

// Add an item to the front of the list
// Example:
//    list.Unshift("item")
func (s *SinglyLinkedList) Unshift(object interface{}) {
	if s.head == nil {
		s.head = &link{ next: &link{ object: &object }}
	} else {
		(*s.head).next = &link{ object: &object, next: (*s.head).next }
	}
	s.len++
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
	return *(*current).object, nil
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
			ch <- (*current).object
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

// Remove an element by its index (starting at 0)
// Example:
//   list.Remove(1)
func (s *SinglyLinkedList) Remove(index int) {
	if s.head == nil {
		return
	}
	var current *link = s.head
	var i int
	for (*current).next != nil {
		if index == i {
			(*current).next = (*(*current).next).next
			s.len--
			return
		}
		current = (*current).next
		i++
	}
}

// Find the index of a value given the value
// Example:
//   index, err := list.Find("item")
//   if err != nil {
//     // log error
//   }
func (s *SinglyLinkedList) Find(object interface{}) (int, error) {
	if s.head == nil {
		return -1, errors.New("Is empty")
	}
	var current *link = s.head
	var i int = -1
	for (*current).next != nil {
		i++
		current = (*current).next
		if *(*current).object == object {
			return i, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Cannot find '%v'", object))
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
		ret[i] = *(*current).object
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
		str += fmt.Sprintf("%v->", *(*current).object) 
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