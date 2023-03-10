package inventory

import (
	"errors"
	"fmt"
)

var (
	ErrTotalAmountNotEnough = errors.New("bank notes and coins are not enough to return")
)

// Node represents a node in the inventory tree.
type Node struct {
	Value  float64
	Amount int
	Left   *Node
	Right  *Node
}

// Inventory represents the inventory for the auto cashier system.
type Inventory struct {
	Root *Node
}

// NewInventory creates a new Inventory instance.
func NewInventory() *Inventory {
	root := &Node{Value: 1000, Amount: 10}
	root.Insert(500, 20)
	root.Insert(100, 15)
	root.Insert(50, 20)
	root.Insert(20, 30)
	root.Insert(10, 20)
	root.Insert(5, 20)
	root.Insert(1, 20)
	root.Insert(0.25, 50)
	return &Inventory{Root: root}
}

func (n *Node) Insert(value float64, amount int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value, Amount: amount}
		} else {
			n.Left.Insert(value, amount)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value, Amount: amount}
		} else {
			n.Right.Insert(value, amount)
		}
	}
}

// TotalAmount returns the total amount of currency available in the subtree rooted at this node.
func (n *Node) TotalAmount() float64 {
	if n == nil {
		return 0
	}
	return n.Value*float64(n.Amount) + n.Left.TotalAmount() + n.Right.TotalAmount()
}

func (i *Inventory) Validate(price float64, payment float64) error {
	change := payment - price
	if i.TotalAmount() <= change {
		return ErrTotalAmountNotEnough
	}
	return nil
}

// Get TotalAmount of inventory
func (i *Inventory) TotalAmount() float64 {
	return i.Root.TotalAmount()
}

// GetInventory returns the current inventory.
func (i *Inventory) GetInventory() []*Node {
	var items []*Node
	traverseInorder(i.Root, &items)
	return items
}

func traverseInorder(node *Node, items *[]*Node) {
	if node != nil {
		traverseInorder(node.Left, items)
		*items = append(*items, node)
		traverseInorder(node.Right, items)
	}
}

// UpdateInventory updates the inventory with a specific value and amount.
func (i *Inventory) UpdateInventory(value float64, amount int) error {
	node := search(i.Root, value)
	if node == nil {
		return fmt.Errorf("value not found")
	}
	node.Amount += amount
	return nil
}

func search(node *Node, value float64) *Node {
	if node == nil || node.Value == value {
		return node
	}
	if value < node.Value {
		return search(node.Left, value)
	}
	return search(node.Right, value)
}
