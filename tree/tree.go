package tree

import "fmt"

type Tree struct {
	Value	int
	Left	*Tree
	Right	*Tree
}

func (t Tree) PrintInorder(){
	//Left
	if t.Left != nil {
		t.Left.PrintInorder()
	}
	//Root
	fmt.Print( t.Value)
	fmt.Print(", ")
	//Right
	if t.Right != nil {
		t.Right.PrintInorder()
	}
}

func (t Tree) PrintPreOrder() {
	// Root
	fmt.Print(t.Value)
	fmt.Print(",")
	//Left
	if t.Left != nil {
		t.Left.PrintPreOrder()
	}
	//Right
	if t.Right != nil {
		t.Right.PrintPreOrder()
	}
}

func (t Tree) PrintPostOrder() {
	//Left
	if t.Left != nil {
		t.Left.PrintPostOrder()
	}
	//Right
	if t.Right != nil {
		t.Right.PrintPostOrder()
	}
	//Root
	fmt.Print(t.Value)
	fmt.Print(",")
}
func (t *Tree) AddToTree(value int){

	if value < t.Value {
		if t.Left == nil {
			t.Left = &Tree{Value: value}
		} else {
			t.Left.AddToTree(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &Tree{Value: value}
		}else {
			t.Right.AddToTree(value)
		}
	}
}

func (t *Tree) DeleteFromTree(val int) *Tree {
	if t == nil {
		return nil
	}
	if val < t.Value {
		t.Left = t.Left.DeleteFromTree(val)
		return t
	}
	if val > t.Value {
		t.Right = t.Right.DeleteFromTree(val)
		return t
	}
	if t.Left == nil && t.Right == nil {
		t = nil
		return t
	}
	if t.Right == nil {
		t = t.Left
		return t
	}
	if t.Left == nil {
		t = t.Right
		return t
	}
	temp := t.Right
	for  {
		if temp != nil && temp.Left != nil {
			temp = temp.Left
		} else {
			break
		}
	}
	t.Value = temp.Value

	t.Right = t.Right.DeleteFromTree(t.Value)

	return t
}
