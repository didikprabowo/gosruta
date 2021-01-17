package bts

import (
	"fmt"
	"log"
)

type (
	Tree struct {
		Root *Node `json:"root"`
	}

	Node struct {
		Value byte
		Left  *Node
		Right *Node
	}

	// BinarySearchTree API references
	BinarySearchTree interface {
		Insert(v float64)
		Get() *Tree
		BulkInsert([]float64)
		Find(value float64) (node *Node, err error)
		Min() *Node
		Max() *Node
		Delete(value float64)
	}
	// NodeTree API references
	NodeTree interface {
		ToValue() (v byte)
	}
)

// NewBinarySearchTree is func for define
func NewBinarySearchTree() BinarySearchTree {
	return &Tree{}
}

// Insert is func for used insert single value
func (t *Tree) Insert(v float64) {
	t.insertChild(t.Root, v)
}

// BulkInsert is func for bulk insert data
func (t *Tree) BulkInsert(values []float64) {
	for i := range values {
		t.Insert(values[i])
	}
}

// Get is func for used get all nodes
func (t *Tree) Get() *Tree {
	return t
}

// Find is func go find value within array
func (t *Tree) Find(value float64) (node *Node, err error) {

	bv := byte(value)

	node = t.findChild(bv, t.Root)
	if node == nil {
		err = fmt.Errorf("Data not found")
		return node, err
	}
	return node, err
}

func (t *Tree) Delete(value float64) {
	t.deleteChild(value, t.Root)
}

// ToValue is func for used get value of node
func (n *Node) ToValue() (v byte) {
	return n.Value
}

func (t *Tree) deleteChild(value float64, in *Node) (node *Node) {

	if in == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("something when wrongs %s", r)
		}
	}()

	vb := byte(value)

	if vb < in.Value {
		fmt.Println("SINI 1", value, in.Left)
		in.Left = t.deleteChild(value, in.Left)
	} else if vb > in.Value {
		in.Right = t.deleteChild(value, in.Right)
	} else {

		switch {
		case in.Left == nil:
			return in.Right
		case in.Right == nil:
			return in.Left
		}

		min := t.minChild(in.Right)
		in.Value = min.Value
		in.Right = t.deleteChild(float64(min.Value), in.Right)
	}

	return in
}

// Min is func for used mininal value of node
func (t *Tree) Min() (n *Node) {
	return t.minChild(t.Root)
}

// Max is func for used maximal value of node
func (t *Tree) Max() *Node {
	n := t.maxChild(t.Root)
	return n
}

func (t *Tree) minChild(n *Node) *Node {
	if n.Left == nil {
		return n
	}

	return t.minChild(n.Left)
}

func (t *Tree) maxChild(n *Node) *Node {
	if n.Right == nil {
		return n
	}

	return t.maxChild(n.Right)
}

// findChild is func for find value at tree with recursice mode
func (t *Tree) findChild(value byte, in *Node) (node *Node) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("something when wrongs %s", r)
		}
	}()

	if in == nil {
		return nil
	}

	if in.Value == value {
		return t.Root
	}

	switch {
	case value <= in.Value:
		return t.findChild(value, in.Left)
	}

	return t.findChild(value, in.Right)
}

// insertChild is func for insert each a node
func (t *Tree) insertChild(in *Node, value float64) (n *Node) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("something when wrongs %s", r)
		}
	}()

	vb := byte(value)

	if t.Root == nil {
		t.Root = &Node{
			Value: vb,
			Left:  nil,
			Right: nil,
		}
		return t.Root
	}

	switch {
	case in == nil:
		return &Node{Value: vb}
	case vb <= in.Value:
		in.Left = t.insertChild(in.Left, value)
	case vb >= in.Value:
		in.Right = t.insertChild(in.Right, value)
	}

	return in
}
