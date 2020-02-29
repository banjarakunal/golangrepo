package main

import "fmt"

//Node represents bst's attributes
type Node struct {
	val   int
	left  *Node
	right *Node
}

//Bst represents binary search tree
/* type Bst struct {
	root *Node
}
*/
//Insert inserts records in binary search tree
func (root *Node) insert(val int) {
	n := &Node{
		val:   val,
		left:  nil,
		right: nil,
	}
	if root.val > n.val {
		if root.left == nil {
			root.left = n
		} else {
			root.left.insert(val)
		}
	} else {
		if root.right == nil {
			root.right = n
		} else {
			root.right.insert(val)
		}
	}
}

func preOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.val, "->")
	preOrderTraversal(n.left)
	preOrderTraversal(n.right)
}

func inOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	preOrderTraversal(n.left)
	fmt.Print(n.val, "->")
	preOrderTraversal(n.right)
}

func postOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	preOrderTraversal(n.left)
	preOrderTraversal(n.right)
	fmt.Print(n.val, "->")
}

func (root *Node) search(key int) bool {
	if root == nil {
		return false
	}
	if root.val > key {
		return root.search(key)
	} else if root.val < key {
		return root.search(key)
	}
	return true
}

/* func (bst *Bst) delete(key int) *Bst {
	if bst.root.val > key {
		bst.root = bst.root.left
		return bst.delete(key)
	} else if bst.root.val < key {
		bst.root = bst.root.right
		return bst.delete(key)
	} else {
		if bst.root.left == nil {
			bst.root = bst.root.right
			return bst
		} else if bst.root.right == nil {
			bst.root = bst.root.left
			return bst
		}
		bst.root = bst.root.right
		minVal := bst.min()
		bst.root.val = minVal

		bst = bst.delete(minVal)
	}
	return bst
} */

/* func (bst *Bst) min() int {
	if bst.root.left == nil {
		return bst.root.val
	}
	bst.root = bst.root.left
	return bst.min()
}

func (bst *Bst) max() int {
	if bst.root.right == nil {
		return bst.root.val
	}
	bst.root = bst.root.right
	return bst.max()
} */

func main() {
	bst := Node{
		val:   8,
		left:  nil,
		right: nil,
	}
	bst.insert(4)
	bst.insert(10)
	bst.insert(2)
	bst.insert(6)
	bst.insert(1)
	bst.insert(3)
	bst.insert(5)
	bst.insert(7)
	bst.insert(9)
	bst.insert(11)
	/* fmt.Println("\nPre order Traversal.....")
	preOrderTraversal(&bst) */
	fmt.Println("\nIn order Traversal..... ")
	inOrderTraversal(&bst)
	/* fmt.Println("\nPost order Traversal..... ")
	postOrderTraversal(&bst)  */
	/* fmt.Println("Serach 10 ", bst.search(10))
	fmt.Println("Serach 100 ", bst.search(100)) */

	//fmt.Println("\nMin value", bst.min())
	//fmt.Println("\nMax value", bst.max())
	/* t := bst.delete(4)
	fmt.Println("\nDekte", t)
	inOrderTraversal(t.root) */

}
