package binarysearchtree

import (
	"fmt"
	"testing"
)

//       4  len=7
//     /   \
//    2     6
//   / \   / \
//  1   3 5   7
func buildTestTree() BinarySearchTree {
	b := BinarySearchTree{}
	b.Insert(4, 4)
	b.Insert(2, 2)
	b.Insert(3, 3)
	b.Insert(6, 6)
	b.Insert(1, 1)
	b.Insert(5, 5)
	b.Insert(7, 7)
	return b
}

func Test_Init(t *testing.T) {
	b := BinarySearchTree{}
	if b.root != nil {
		t.Errorf("Expected root to be nil, got '%v'", b.root)
	}
	if b.len != 0 {
		t.Errorf("Expected len to be 0, got %v", b.len)
	}
}

func Test_Insert1(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(4, 4)
	if b.root == nil {
		t.Errorf("Expected root to not be nil, got '%v'", b.root)
	}
	if b.root.value != 4 {
		t.Errorf("Expected root.value == 'test', got '%v'", b.root.value)
	}
	if b.len != 1 {
		t.Errorf("Expected len to be 1, got '%v'", b.len)
	}
}

func Test_Insert2(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(4, "test")
	b.Insert(5, "test2")
	b.Insert(1, "test3")
	if b.root == nil {
		t.Errorf("Expected root to not be nil, got '%v'", b.root)
	}
	if b.root.value != "test" {
		t.Errorf("Expected root.value == 'test', got '%v'", b.root.value)
	}
	if b.root.right.value != "test2" {
		t.Errorf("Expected root.right.value == 'test2', got '%v'", b.root.right.value)
	}
	if b.root.left.value != "test3" {
		t.Errorf("Expected root.left.value == 'test3', got '%v'", b.root.left.value)
	}
	if b.len != 3 {
		t.Errorf("Expected len to be 3, got '%v'", b.len)
	}
}

func Test_Insert3(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(4, "test")
	b.Insert(5, "test2")
	b.Insert(1, "test3")
	b.Insert(2, "test4")
	if b.root == nil {
		t.Errorf("Expected root to not be nil, got '%v'", b.root)
	}
	if b.root.value != "test" {
		t.Errorf("Expected root.value == 'test', got '%v'", b.root.value)
	}
	if b.root.right.value != "test2" {
		t.Errorf("Expected root.right.value == 'test2', got '%v'", b.root.right.value)
	}
	if b.root.left.value != "test3" {
		t.Errorf("Expected root.left.value == 'test3', got '%v'", b.root.left.value)
	}
	if b.root.left.right.value != "test4" {
		t.Errorf("Expected root.left.right.value == 'test3', got '%v'", b.root.left.right.value)
	}
	if b.len != 4 {
		t.Errorf("Expected len to be 4, got '%v'", b.len)
	}
}

func Test_Insert4(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(1, 1)
	if b.len != 1 {
		t.Errorf("Expected len to be 1, got '%v'", b.len)
	}
	if b.root.value != 1 {
		t.Errorf("Expected root.value to be 1, got '%v'", b.root.value)
	}
	b.Insert(1, 3)
	if b.len != 1 {
		t.Errorf("Expected len to be 1, got '%v'", b.len)
	}
	if b.root.value != 3 {
		t.Errorf("Expected root.value to be 3, got '%v'", b.root.value)
	}
}

//      4  len=7     4  len=6
//    /   \        /   \
//   2    >6  ->  2     5
//  / \   / \    / \     \
// 1   3 5   7  1   3     7
func Test_Delete1(t *testing.T) {
	b := buildTestTree()
	v, err := b.Delete(6)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 6 {
		t.Errorf("Expected the deleted value to be 'test', got '%v'", v)
	}
	if b.len != 6 {
		t.Errorf("Expected len to be 6, got '%v'", b.len)
	}
	if b.root.value != 4 {
		t.Errorf("Expected root.value to be 4, got '%v'", b.root.value)
	}
	if b.root.left.value != 2 {
		t.Errorf("Expected root.left.value to be 2, got '%v'", b.root.left.value)
	}
	if b.root.left.left.value != 1 {
		t.Errorf("Expected root.left.left.value to be 1, got '%v'", b.root.left.left.value)
	}
	if b.root.left.right.value != 3 {
		t.Errorf("Expected root.left.right.value to be 3, got '%v'", b.root.left.right.value)
	}
	if b.root.right.value != 5 {
		t.Errorf("Expected root.right.value to be 5, got '%v'", b.root.right.value)
	}
	if b.root.right.right.value != 7 {
		t.Errorf("Expected root.right.right.value to be 7, got '%v'", b.root.right.right.value)
	}
}

//      4  len=8       4  len=7
//    /   \          /   \
//   2    >5   ->   2     7
//  / \     \      / \   / \
// 1   3   6-7-8  1   3 6   8
func Test_Delete2(t *testing.T) {
	b := buildTestTree()
	v, err := b.Delete(6)
	b.Insert(8, 8)
	b.Insert(6, 6)
	v, err = b.Delete(5)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 5 {
		t.Errorf("Expected the deleted value to be 5, got '%v'", v)
	}
	if b.len != 7 {
		t.Errorf("Expected len to be 7, got '%v'", b.len)
	}
	if b.root.value != 4 {
		t.Errorf("Expected root.value to be 4, got '%v'", b.root.value)
	}
	if b.root.left.value != 2 {
		t.Errorf("Expected root.left.value to be 2, got '%v'", b.root.left.value)
	}
	if b.root.left.left.value != 1 {
		t.Errorf("Expected root.left.left.value to be 1, got '%v'", b.root.left.left.value)
	}
	if b.root.left.right.value != 3 {
		t.Errorf("Expected root.left.right.value to be 3, got '%v'", b.root.left.right.value)
	}
	if b.root.right.value != 7 {
		t.Errorf("Expected root.right.value to be 7, got '%v'", b.root.right.value)
	}
	if b.root.right.right.value != 8 {
		t.Errorf("Expected root.right.right.value to be 8, got '%v'", b.root.right.right.value)
	}
	if b.root.right.left.value != 6 {
		t.Errorf("Expected root.right.left.value to be 6, got '%v'", b.root.right.left.value)
	}
}

//     >4  len=7     5  len=6
//    /   \        /   \
//   2     6  ->  2     6
//  / \   / \    / \     \
// 1   3 5   7  1   3     7
func Test_Delete3(t *testing.T) {
	b := buildTestTree()
	v, err := b.Delete(4)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 4 {
		t.Errorf("Expected the deleted value to be 4, got '%v'", v)
	}
	if b.root.value != 5 {
		t.Errorf("Expected root.value to be 5, got '%v'", b.root.value)
	}
	if b.root.right.value != 6 {
		t.Errorf("Expected root.right.value to be 6, got '%v'", b.root.right.value)
	}
	if b.root.right.right.value != 7 {
		t.Errorf("Expected root.right.right.value to be 7, got '%v'", b.root.right.right.value)
	}
	if b.root.left.value != 2 {
		t.Errorf("Expected root.left.value to be 2, got '%v'", b.root.left.value)
	}
	if b.root.left.left.value != 1 {
		t.Errorf("Expected root.left.left.value to be 1, got '%v'", b.root.left.left.value)
	}
	if b.root.left.right.value != 3 {
		t.Errorf("Expected root.left.right.value to be 3, got '%v'", b.root.left.right.value)
	}
}

//      4  len=7     4  len=6
//    /   \        /   \
//   2     6  ->  2     6
//  / \   / \    / \     \
// 1   3 5<  7  1   3     7
func Test_Delete4(t *testing.T) {
	b := buildTestTree()
	v, err := b.Delete(5)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 5 {
		t.Errorf("Expected the deleted value to be 5, got '%v'", v)
	}
	if b.root.value != 4 {
		t.Errorf("Expected root.value to be 4, got '%v'", b.root.value)
	}
	if b.root.left.value != 2 {
		t.Errorf("Expected root.left.value to be 2, got '%v'", b.root.left.value)
	}
	if b.root.left.left.value != 1 {
		t.Errorf("Expected root.left.left.value to be 1, got '%v'", b.root.left.left.value)
	}
	if b.root.left.right.value != 3 {
		t.Errorf("Expected root.left.right.value to be 3, got '%v'", b.root.left.right.value)
	}
	if b.root.right.value != 6 {
		t.Errorf("Expected root.right.value to be 6, got '%v'", b.root.right.value)
	}
	if b.root.right.right.value != 7 {
		t.Errorf("Expected root.right.right.value to be 7, got '%v'", b.root.right.right.value)
	}
}

// >4  len=1 ->     len=0
func Test_Delete5(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(4, 4)
	v, err := b.Delete(4)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 4 {
		t.Errorf("Expected deleted value to be 4, got '%v'", v)
	}
	if b.len != 0 {
		t.Errorf("Expected len to be 0, got '%v'", b.len)
	}
}

// 4  len=2 4  len=1
//  \   ->
//  >5
func Test_Delete6(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(4, 4)
	b.Insert(5, 5)
	v, err := b.Delete(5)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 5 {
		t.Errorf("Expected deleted value to be 5, got '%v'", v)
	}
	if b.root.value != 4 {
		t.Errorf("Expected root.value to be 4, got '%v'", b.root.value)
	}
	if b.len != 1 {
		t.Errorf("Expected len to be 1, got '%v'", b.len)
	}
}

//     4  len=3    4  len=2
//    /           /
//  >2      ->   1
//  /
// 1
func Test_Delete7(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(4, 4)
	b.Insert(2, 2)
	b.Insert(1, 1)
	v, err := b.Delete(2)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 2 {
		t.Errorf("Expected deleted value to be 2, got '%v'", v)
	}
	if b.root.value != 4 {
		t.Errorf("Expected root.value to be 4, got '%v'", b.root.value)
	}
	if b.root.left.value != 1 {
		t.Errorf("Expected root.left.value to be 1, got '%v'", b.root.left.value)
	}
	if b.len != 2 {
		t.Errorf("Expected len to be 2, got '%v'", b.len)
	}
}

//     5  len=5    5  len=4
//    /           /
//  >2   ->      3
//  / \         / \
// 1   3-4     1   4
func Test_Delete8(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(5, 5)
	b.Insert(2, 2)
	b.Insert(3, 3)
	b.Insert(1, 1)
	b.Insert(4, 4)
	v, err := b.Delete(2)
	fmt.Println(b.Sprint())
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 2 {
		t.Errorf("Expected deleted value to be 2, got '%v'", v)
	}
	if b.root.value != 5 {
		t.Errorf("Expected root.value to be 5, got '%v'", b.root.value)
	}
	if b.root.left.value != 3 {
		t.Errorf("Expected root.left.value to be 3, got '%v'", b.root.left.value)
	}
	if b.root.left.left.value != 1 {
		t.Errorf("Expected root.left.left.value to be 1, got '%v'", b.root.left.left.value)
	}
	if b.root.left.right.value != 4 {
		t.Errorf("Expected root.left.right.value to be 4, got '%v'", b.root.left.right.value)
	}
	if b.len != 4 {
		t.Errorf("Expected len to be 4, got '%v'", b.len)
	}
}

// 5  len=3   5
//  \          \
//  >7    ->    6
//  /
// 6
func Test_Delete9(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(5, 5)
	b.Insert(7, 7)
	b.Insert(6, 6)
	v, err := b.Delete(7)
	if err != nil {
		t.Errorf("Expected error to be nil, got '%v'", err)
	}
	if v != 7 {
		t.Errorf("Expected deleted value to be 7, got '%v'", v)
	}
	if b.root.value != 5 {
		t.Errorf("Expected root.value to be 5, got '%v'", b.root.value)
	}
	if b.root.right.value != 6 {
		t.Errorf("Expected root.right.value to be 6, got '%v'", b.root.right.value)
	}
	if b.len != 2 {
		t.Errorf("Expected len to be 2, got '%v'", b.len)
	}
}

func Test_Delete_Error1(t *testing.T) {
	b := BinarySearchTree{}
	_, err := b.Delete(4)
	if err == nil {
		t.Errorf("Expected err to exist, got nil")
	}
}
func Test_Delete_Error2(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(3, 3)
	_, err := b.Delete(4)
	if err == nil {
		t.Errorf("Expected err to exist, got nil")
	}
	_, err = b.Delete(2)
	if err == nil {
		t.Errorf("Expected err to exist, got nil")
	}
}
func Test_IterPreorder(t *testing.T) {
	l := []int{4, 2, 1, 3, 6, 5, 7}
	b := buildTestTree()
	var j int
	for i := range b.IterPreorder() {
		if *i != l[j] {
			t.Errorf("Expected to get '%d', got '%v'", l[j], *i)
		}
		j++
	}
}

func Test_IterPostorder(t *testing.T) {
	l := []int{1, 3, 2, 5, 7, 6, 4}
	b := buildTestTree()
	var j int
	for i := range b.IterPostorder() {
		if *i != l[j] {
			t.Errorf("Expected to get '%d', got '%v'", l[j], *i)
		}
		j++
	}
}

func Test_IterInorder(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7}
	b := buildTestTree()
	var j int
	for i := range b.IterInorder() {
		if *i != l[j] {
			t.Errorf("Expected to get '%d', got '%v'", l[j], *i)
		}
		j++
	}
}

func Test_Sprint(t *testing.T) {
	b := buildTestTree()
	var str string = "BinarySearchTree{ 1->2->3->4->5->6->7 }"
	if b.Sprint() != str {
		t.Errorf("Expected to get '%s', got '%v'", str, b.Sprint())
	}
}
