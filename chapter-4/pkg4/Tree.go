package pkg4

import "fmt"
import "math"
import "sort"

type Node struct {
	Key int
	Left, Right, Parent, Peer *Node
}

func NewNode(key int) (n *Node) {
	n = new(Node)
	n.Key = key
	return
}

func NewBST(keys ...int) (root *Node) {
	if len(keys) == 0 {
		return
	}

	root = NewNode(keys[0])
	for _, i := range keys[1:] {
		root.InsertBST(i)
	}

	return
}

func max(n ...int) (m int) {
	m = 1 >> 32
	for _, i := range n {
		if i > m {
			m = i
		}
	}

	return
}

func (root *Node) Height(diameter *int) int {
	if root == nil {
		return 0
	}

	lheight := root.Left.Height(diameter)
	rheight := root.Right.Height(diameter)

	if diameter != nil {
		thisDiag := 1 + lheight + rheight
		if thisDiag > *diameter {
			*diameter = thisDiag
		}
	}

	return 1 + max(lheight, rheight)
}

func (root *Node) Diameter() (diam int) {
	if root == nil {
		return 0
	}

	root.Height(&diam)
	return
}

func (root *Node) ConnectPeers() {
	if root == nil {
		return
	}

	pos, length := 0, 1
	queue := make([]*Node, int(math.Pow(2, float64(root.Height(nil))))-1)
	queue[pos] = root

	for pos < length {
		count := length - pos
		fmt.Print("{ ")
		for ; count > 0; count-- {
			fmt.Print(queue[pos].Key, " ")

			if count > 1 {
				queue[pos].Peer = queue[pos+1]
			}

			if queue[pos].Left != nil {
				queue[length] = queue[pos].Left
				length++
			}

			if queue[pos].Right != nil {
				queue[length] = queue[pos].Right
				length++
			}

			pos++
		}
		fmt.Print("}")
	}
}

func (root *Node) convertToDLL(origin bool) *Node {
	if root == nil {
		return root
	}

	if root.Left != nil {
		// Get left subtree converted.
		left := root.Left.convertToDLL(false)

		// Adjust root and its left subtree relation.
		for ; left.Right != nil; left = left.Right {}
		left.Right = root
		root.Left = left
	}

	if root.Right != nil {
		// Get right subtree converted.
		right := root.Right.convertToDLL(false)

		// Adjust root and its right subtree relation.
		for ; right.Left != nil; right = right.Left {}
		right.Left = root
		root.Right = right
	}

	if origin {
		// Move root pointer to the beginning of the list.
		for ; root.Left != nil; root = root.Left {}
	}

	return root
}

func (root *Node) maxWidthByOrder(width *[]int, level int, mwidth *int) int {
	if root == nil {
		return 0
	}

	if width == nil {
		width = new([]int)
		*width = make([]int, root.Height(nil))
	}

	if mwidth == nil {
		mwidth = new(int)
		*mwidth = 0
	}

	(*width)[level]++
	if (*width)[level] > *mwidth {
		*mwidth = (*width)[level]
	}

	if root.Left != nil {
		root.Left.maxWidthByOrder(width, level+1, mwidth)
	}

	if root.Right != nil {
		root.Right.maxWidthByOrder(width, level+1, mwidth)
	}

	return *mwidth
}

func (root *Node) maxWidthByQueue() (mwidth int) {
	if root == nil {
		return
	}

	type nheight struct {
		n *Node
		h int
	}

	height := root.Height(nil)
	hwidth := make([]int, height)
	tqueue := make([]nheight, int(math.Pow(2, float64(height)))-1)

	pos, length := 0, 1
	tqueue[pos].n = root
	tqueue[pos].h = 0

	for pos = 0; pos < length; pos++ {
		// Increase the width of this tree height.
		hwidth[tqueue[pos].h]++
		if hwidth[tqueue[pos].h] > mwidth {
			mwidth = hwidth[tqueue[pos].h]
		}

		if tqueue[pos].n.Left != nil {
			tqueue[length].n = tqueue[pos].n.Left
			tqueue[length].h = tqueue[pos].h+1
			length++
		}

		if tqueue[pos].n.Right != nil {
			tqueue[length].n = tqueue[pos].n.Right
			tqueue[length].h = tqueue[pos].h+1
			length++
		}
	}

	return
}

func (root *Node) SearchLCA(key1, key2 int) (lca int) {
	if key1 > key2 {
		key1, key2 = key2, key1
	}

	for root != nil {
		if key2 < root.Key {
			root = root.Left
		} else if key1 > root.Key {
			root = root.Right
		} else {
			lca = root.Key
			break
		}
	}

	return
}

func (root *Node) sumPair(sum int) (key1, key2 int, found bool) {
	if root == nil {
		return
	}

	length1, length2 := 0, 0
	curr1, curr2 := root, root
	done1, done2 := false, false
	stack1 := make([]*Node, root.Height(nil))
	stack2 := make([]*Node, root.Height(nil))

	for !found {
		for !done1 {
			if curr1 != nil {
				stack1[length1] = curr1
				length1++
				curr1 = curr1.Left
			} else {
				if length1 != 0 {
					length1--
					key1 = stack1[length1].Key
					curr1 = stack1[length1].Right

				}
				done1 = true
			}
		}

		for !done2 {
			if curr2 != nil {
				stack2[length2] = curr2
				length2++
				curr2 = curr2.Right
			} else {
				if length2 != 0 {
					length2--
					key2 = stack2[length2].Key
					curr2 = stack2[length2].Left
				}
				done2 = true
			}
		}

		if key1 >= key2 {
			break
		} else if key1 + key2 == sum {
			found = true
		} else if key1 + key2 < sum {
			done1 = false
		} else if key1 + key2 > sum {
			done2 = false
		}
	}

	return
}

func (root *Node) correction(first, middle, last, prev **Node) {
	if root == nil {
		return
	}

	swap := false
	if first == nil {
		first, middle, last = new(*Node), new(*Node), new(*Node)
		prev = new(*Node)
		swap = true
	}

	root.Left.correction(first, middle, last, prev)

	if *prev != nil && (*prev).Key > root.Key {
		if *first == nil {
			*first = *prev
			*middle = root
		} else {
			*last = root
		}
	}
	*prev = root

	root.Right.correction(first, middle, last, prev)

	if !swap {
		return
	}

	if *first != nil && *last != nil {
		(*first).Key, (*last).Key = (*last).Key, (*first).Key
	} else if *first != nil && *middle != nil {
		(*first).Key, (*middle).Key = (*middle).Key, (*first).Key
	}
}

func (root *Node) successor() (*Node, *Node) {
        if root == nil || root.Left == nil {
                return root, nil
        }

        for ; root.Left.Left != nil; root = root.Left {}

        return root.Left, root
}

func (root *Node) SearchBST(key int) (bool, *Node, *Node) {
	if root == nil {
		return false, nil, nil
	} else if key == root.Key {
		return true, root, nil
	}

	if key < root.Key {
		if root.Left == nil {
			return false, nil, root
		} else if root.Left.Key == key {
			return true, root.Left, root
		}
		return root.Left.SearchBST(key)
	} else {
		if root.Right == nil {
			return false, nil, root
		} else if root.Right.Key == key {
			return true, root.Right, root
		}
		return root.Right.SearchBST(key)
	}
}

func (root *Node) InsertBST(key int) (bool, *Node) {
	present, n, p := root.SearchBST(key)

	if present || p == nil {
		return false, n
	}

	if key < p.Key {
		p.Left = NewNode(key)
		p.Left.Parent = p
		return true, p.Left
	} else {
		p.Right = NewNode(key)
		p.Right.Parent = p
		return true, p.Right
	}
}

func (root *Node) DeleteBST(key int) (bool, *Node) {
	if root == nil {
		return false, nil
	} else if key == root.Key {
		if root.Left == nil {
			return true, root.Right
		} else if root.Right == nil {
			return true, root.Left
		}

		succ, psucc := root.Right.successor()
		succ.Left = root.Left
		if psucc != nil {
			psucc.Left = succ.Right
		}

		if root.Right != succ {
			succ.Right = root.Right
		}
		root.Left, root.Right = nil, nil
		return true, succ
	}

	if key < root.Key {
		ok, newroot := root.Left.DeleteBST(key)
		if ok {
			root.Left = newroot
		}
	} else if key > root.Key {
		ok, newroot := root.Right.DeleteBST(key)
		if ok {
			root.Right = newroot
		}
	}

	return true, root
}

func (root *Node) PrintInorder() {
	if root == nil {
		return
	}

	root.Left.PrintInorder()
	fmt.Print(root.Key, " ")
	root.Right.PrintInorder()
}

func (root *Node) PrintPreorder() {
	if root == nil {
		return
	}

	fmt.Print(root.Key, " ")
	root.Left.PrintPreorder()
	root.Right.PrintPreorder()
}

func (root *Node) PrintPostorder() {
	if root == nil {
		return
	}

	root.Left.PrintPostorder()
	root.Right.PrintPostorder()
	fmt.Print(root.Key, " ")
}

func (root *Node) PrintLevelorder() {
	if root == nil {
		return
	}

	queue := make([]*Node, int(math.Pow(2, float64(root.Height(nil))))-1)
	queue[0] = root
	pos, length := 0, 1

	for pos < length {
		fmt.Print(queue[pos].Key, " ")

		if queue[pos].Left != nil {
			queue[length] = queue[pos].Left
			length++
		}

		if queue[pos].Right != nil {
			queue[length] = queue[pos].Right
			length++
		}

		pos++
	}
}

func (root *Node) PrintDLL() {
	for ; root != nil; root = root.Right {
		fmt.Print(root.Key, " ")
	}
}

func (root *Node) PrintPeers() {
	for ; root != nil; root = root.Peer {
		fmt.Print(root.Key, " ")
	}
}

func (root *Node) PrintAncestors(key int) bool {
	if root == nil {
		return false
	}

	if root.Key == key {
		return true
	}

	if root.Left.PrintAncestors(key) || root.Right.PrintAncestors(key) {
		fmt.Print(root.Key, " ")
		return true
	}

	return false
}

func (root *Node) PrintVerticals(verticals *map[int][]int, distance int) {
	if root == nil {
		return
	}

	print := false
	if verticals == nil {
		verticals = new(map[int][]int)
		*verticals = make(map[int][]int)
		print = true
	}

	(*verticals)[distance] = append((*verticals)[distance], root.Key)
	root.Left.PrintVerticals(verticals, distance-1)
	root.Right.PrintVerticals(verticals, distance+1)

	if !print {
		return
	}

	var keys []int
	for k := range (*verticals) {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range(keys) {
		fmt.Print((*verticals)[key], " ")
	}
}
