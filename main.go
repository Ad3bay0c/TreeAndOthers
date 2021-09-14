package main

import (
	"fmt"
	"github.com/Ad3bay0c/TreeImplementation/tree"
)

func main() {

	//var n int
	//fmt.Scanf("%d", &n)
	//fmt.Printf("N value is .....%d\n", n)
	//
	//for i := 0; i < 3; i++ {
	//	var val, left, right int
	//	fmt.Scanf("%d %d %d", &val, &left, &right)
	//	fmt.Printf("Val %d, Left %d, Right %d\n", val, left, right)
	//}

	t := &tree.Tree{Value: 25}
	t.AddToTree(27)
	t.AddToTree(26)
	t.AddToTree(1)
	t.AddToTree(13)
	t.AddToTree(30)
	t.PrintInorder()
	fmt.Println()
	t.PrintPreOrder()
	fmt.Println()
	t.PrintPostOrder()
	fmt.Println()

	t.DeleteFromTree(1)
	t.PrintInorder()
	fmt.Println()

}