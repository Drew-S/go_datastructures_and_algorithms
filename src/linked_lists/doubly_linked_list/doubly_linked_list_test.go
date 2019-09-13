package doublylinkedlist

import (
	"testing"
)

func Test_Init(t *testing.T) {
	l := DoublyLinkedList{}
	if l.head != nil {
		t.Errorf("Should: head == nil not '%v'", l.head)
	}
	if l.tail != nil {
		t.Errorf("Should: tail == nil not '%v'", l.tail)
	}
}

func Test_Insert1(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(1, "test")
	if l.head == nil {
		t.Errorf("Should: head != nil")
	}
	if (*l.head).next == nil {
		t.Errorf("Should: head.next != nil")
	}
	if *(*(*l.head).next).value != "test" {
		t.Errorf("Should: head.next == 'test' got '%v'", *(*(*l.head).next).value)
	}
}
func Test_Insert2(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(1, "test")
	l.Insert(3, "test2")
	l.Insert(2, "test1")
	if l.head == nil {
		t.Errorf("Should: head != nil")
	}
	if l.head.next == nil {
		t.Errorf("Should: head.next != nil")
	}
	if *l.head.next.next.value != "test1" {
		t.Errorf("Should: head.next.next == 'test1' got '%v'", *l.head.next.next.value)
	}
}
func Test_Insert3(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(1, "test")
	l.Insert(2, "test1")
	if l.head == nil {
		t.Errorf("Should: head != nil")
	}
	if (*l.head).next == nil {
		t.Errorf("Should: head.next != nil")
	}
	if *(*(*(*l.head).next).next).value != "test1" {
		t.Errorf("Should: head.next.next == 'test1' got '%v'", *(*(*l.head).next).value)
	}
	if l.len != 2 {
		t.Errorf("Should have 2 elements, got %v", l.len)
	}
	l.Insert(2, "test2")
	if *(*(*(*l.head).next).next).value != "test2" {
		t.Errorf("Should: head.next.next == 'test2' got '%v'", *(*(*l.head).next).value)
	}
	if l.len != 2 {
		t.Errorf("Should have 2 elements, got %v", l.len)
	}
}

func Test_Pop(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(0, "test1")
	l.Insert(1, "test2")
	l.Insert(2, "test3")
	if l.len != 3 {
		t.Errorf("Should have len of 3 not '%d'", l.len)
	}
	i, err := l.Pop()
	if err != nil {
		t.Errorf("Expected no error got '%v'", err)
	}
	if i != "test3" {
		t.Errorf("Expected the popped value to be 'test3' not '%v'", i)
	}
	if l.len != 2 {
		t.Errorf("Should have len of 2 not '%d'", l.len)
	}
}

func Test_Pop_Error(t *testing.T) {
	l := DoublyLinkedList{}
	_, err := l.Pop()
	if err == nil {
		t.Errorf("Expected error to be thrown, not nil")
	}
}

func Test_Shift(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(0, "test1")
	l.Insert(1, "test2")
	l.Insert(2, "test3")
	if l.len != 3 {
		t.Errorf("Should have len of 3 not '%d'", l.len)
	}
	i, err := l.Shift()
	if err != nil {
		t.Errorf("Expected no error got '%v'", err)
	}
	if i != "test1" {
		t.Errorf("Expected the popped value to be 'test1' not '%v'", i)
	}
	if l.len != 2 {
		t.Errorf("Should have len of 2 not '%d'", l.len)
	}
}

func Test_Shift_Error(t *testing.T) {
	l := DoublyLinkedList{}
	_, err := l.Shift()
	if err == nil {
		t.Errorf("Expected error to be thrown, not nil")
	}
}

func Test_Iter(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(1, 1)
	l.Insert(2, 2)
	l.Insert(3, 3)
	l.Insert(4, 4)
	l.Insert(5, 5)
	var i int = 1
	for j := range l.Iter() {
		if *j != i {
			t.Errorf("Expected at index [%d] for '%v == %d'", i, *j, i)
		}
		i++
	}
}

func Test_IterReverse(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(1, 5)
	l.Insert(2, 4)
	l.Insert(3, 3)
	l.Insert(4, 2)
	l.Insert(5, 1)
	var i int = 1
	for j := range l.IterReverse() {
		if *j != i {
			t.Errorf("Expected at index [%d] for '%v == %d'", i-1, *j, i)
		}
		i++
	}
}

func Test_Len(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(0, "test1")
	l.Insert(1, "test2")
	if l.Len() != 2 {
		t.Errorf("Expected Len() == 2 not '%d'", l.Len())
	}
}

func Test_Remove(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(0, "test1")
	l.Insert(1, "test2")
	l.Insert(2, "test3")
	l.Insert(3, "test4")
	l.Insert(4, "test5")
	item, err := l.Remove(2)
	if err != nil {
		t.Errorf("Expected err to be nil, got '%v'", err)
	}
	if item != "test3" {
		t.Errorf("Expected to get 'test3' back, got '%v'", item)
	}
	for i := range l.Iter() {
		if *i == "test3" {
			t.Errorf("Expected 'test3' to no longer be in the list")
		}
	}
	if l.len != 4 {
		t.Errorf("Expected length to be 4, not %d", l.len)
	}
}

func Test_Remove_Error(t *testing.T) {
	l := DoublyLinkedList{}
	_, err := l.Remove(12)
	if err == nil {
		t.Errorf("Expected err to exist, got '%v'", err)
	}
	l.Insert(0, 0)
	_, err = l.Remove(12)
	if err == nil {
		t.Errorf("Expected err to exist, got '%v'", err)
	}
}

func Test_Find(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(0, "test1")
	l.Insert(1, "test2")
	i, err := l.Find("test2")
	if err != nil {
		t.Errorf("There should be no error, got '%v'", err)
	}
	if i != 1 {
		t.Errorf("Expected to get key 1 not '%d'", i)
	}
}

func Test_Find_Error(t *testing.T) {
	l := DoublyLinkedList{}
	_, err := l.Find(12)
	if err == nil {
		t.Errorf("Expected err to exist, got '%v'", err)
	}
	l.Insert(0, 0)
	_, err = l.Find(12)
	if err == nil {
		t.Errorf("Expected err to exist, got '%v'", err)
	}
}

func Test_ToArray(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(0, "test1")
	l.Insert(1, "test2")
	a, err := l.ToArray()
	if err != nil {
		t.Errorf("Expected error to be nil got '%v'", err)
	}
	if *a[0] != "test1" {
		t.Errorf("Expected arr[0] to be 'test1' got '%v'", *a[0])
	}
	if *a[1] != "test2" {
		t.Errorf("Expected arr[1] to be 'test2' got '%v'", *a[1])
	}
}

func Test_ToArray_Error(t *testing.T) {
	l := DoublyLinkedList{}
	_, err := l.ToArray()
	if err == nil {
		t.Errorf("Expected err to exist, got '%v'", err)
	}
}

func Test_String(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(0, "test1")
	l.Insert(1, "test2")
	var output string = "DoublyLinkedList{ head<->test1<->test2<->tail }"
	if l.String() != output {
		t.Errorf("Expected '%s' got '%v'", output, l.String())
	}
}
