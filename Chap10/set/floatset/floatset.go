package floatset

import (
    "fmt"
    "sort"
    "sync"
)

const (
    Concurrent = 32
)

var max [Concurrent]float64 // Holds the maximum value in each aVL tree

//MakeSet initialize the set
func NewSet() *Set {
    return &Set{
        container: make(map[float64]struct{}),
    }
}

type Set struct {
    container map[float64]struct{}
}

func (c *Set) IsPresent(key float64) bool {
    _, present := c.container[key]
    return present
}

func (c *Set) Add(key float64) {
    c.container[key] = struct{}{}
}

func (c *Set) Remove(key float64) error {
    _, present := c.container[key]
    if !present {
        return fmt.Errorf("Remove Error: Item doesn't exist in set")
    }
    delete(c.container, key)
    return nil
}

func (c *Set) Size() int {
    return len(c.container)
}

type AVLTree struct {
    Root     *Node
    NumNodes int
}

type Node struct {
    Value float64
    Left  *Node
    Right *Node
    Ht    int
}

// Methods
func (avl *AVLTree) Insert(newValue float64) {
    avl.Root = insertNode(avl.Root, newValue)
    avl.NumNodes += 1

}

func (avl *AVLTree) Delete(value float64) {
    avl.Root = deleteNode(avl.Root, value)
    avl.NumNodes -= 1
}

func (avl *AVLTree) Search(value float64) bool {
    return search(avl.Root, value)
}

func (avl *AVLTree) Height() int {
    return avl.Root.Height()
}

func (n *Node) balanceFactor() int {
    if n == nil {
        return 0
    }
    return n.Left.Height() - n.Right.Height()
}

func (n *Node) Height() int {
    if n == nil {
        return 0
    } else {
        return n.Ht
    }
}

func (n *Node) updateHeight() {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    n.Ht = max(n.Left.Height(), n.Right.Height()) + 1
}

// Internal functions
func newNode(val float64) *Node {
    return &Node{
        Value: val,
        Left:  nil,
        Right: nil,
        Ht:    1,
    }
}

func search(n *Node, value float64) bool {
    if n == nil {
        return false
    }
    if value < n.Value {
        return search(n.Left, value)
    }
    if value > n.Value {
        return search(n.Right, value)
    }
    return true
}

func insertNode(node *Node, val float64) *Node {
    if node == nil {
        return newNode(val)
    }
    if val > node.Value {
        right := insertNode(node.Right, val)
        node.Right = right
    }
    if val < node.Value {
        left := insertNode(node.Left, val)
        node.Left = left
    }
    return rotateInsert(node, val)
}

func rightRotate(x *Node) *Node {
    y := x.Left
    t := y.Right

    y.Right = x
    x.Left = t

    x.updateHeight()
    y.updateHeight()

    return y
}

func leftRotate(x *Node) *Node {
    y := x.Right
    t := y.Left

    y.Left = x
    x.Right = t

    x.updateHeight()
    y.updateHeight()

    return y
}

func rotateInsert(node *Node, val float64) *Node {
    node.updateHeight()

    // bFactor will tell you which side the weight is on
    bFactor := node.balanceFactor()

    if bFactor > 1 && val < node.Left.Value {
        return rightRotate(node)
    }

    if bFactor < -1 && val > node.Right.Value {
        return leftRotate(node)
    }

    if bFactor > 1 && val > node.Left.Value {
        node.Left = leftRotate(node.Left)
        return rightRotate(node)
    }

    if bFactor < -1 && val < node.Right.Value {
        node.Right = rightRotate(node.Right)
        return leftRotate(node)
    }
    return node
}

func greatest(node *Node) *Node {
    if node == nil {
        return nil
    }

    if node.Right == nil {
        return node
    }
    return greatest(node.Right)
}

func rotateDelete(node *Node) *Node {
    node.updateHeight()
    bFactor := node.balanceFactor()

    if bFactor > 1 && node.Left.balanceFactor() >= 0 {
        return rightRotate(node)
    }

    if bFactor > 1 && node.Left.balanceFactor() < 0 {
        node.Left = leftRotate(node.Left)
        return rightRotate(node)
    }

    if bFactor < -1 && node.Right.balanceFactor() <= 0 {
        return leftRotate(node)
    }

    if bFactor < -1 && node.Right.balanceFactor() > 0 {
        node.Right = rightRotate(node.Right)
        return leftRotate(node)
    }
    return node
}

func deleteNode(node *Node, val float64) *Node {
    if node == nil {
        return nil
    }

    if val > node.Value {
        right := deleteNode(node.Right, val)
        node.Right = right
    } else if val < node.Value {
        left := deleteNode(node.Left, val)
        node.Left = left
    } else {
        if node.Left != nil && node.Right != nil {
            // has 2 children, find the successor
            successor := greatest(node.Left)
            value := successor.Value

            // remove the successor
            left := deleteNode(node.Left, value)
            node.Left = left

            // copy the successor value to the current node
            node.Value = value
        } else if node.Left != nil || node.Right != nil {
            // has 1 child
            // move the child position to the current node
            if node.Left != nil {
                node = node.Left
            } else {
                node = node.Right
            }
        } else if node.Left == nil && node.Right == nil {
            // has no child
            // simply remove the node
            node = nil
        }
    }
    if node == nil {
        return nil
    }
    return rotateDelete(node)
}

var concurrrentSet [Concurrent]AVLTree

func BuildConcurrentSet(dataSet []float64) {
    // Use concurrent processing to construct concurrent avl trees
    var wg sync.WaitGroup
    sort.Float64s(dataSet)
    segment := len(dataSet) / Concurrent
    for treeNumber := 0; treeNumber < Concurrent; treeNumber++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            startVal := segment * num
            for j := startVal; j < startVal+segment; j++ {
                concurrrentSet[num].Insert(dataSet[j])
            }
            max[num] = dataSet[startVal+segment-1]
        }(treeNumber)
    }
    wg.Wait()
}

func IsPresent(val float64) bool {
    // Determine which AVL tree val is in
    treeNumber := 0
    for ; treeNumber < len(max); treeNumber++ {
        if val <= max[treeNumber] {
            break
        }
    }
    return concurrrentSet[treeNumber].Search(val)
}
