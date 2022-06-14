package main

import (
	bst"example.com/binarysearchtree"
	"math/rand"
	"time"
	"fmt"
)

// Satisfies OrderedStringer because of ~int
// Also satisfies OrderedStringer because of String() method below
type Number int  

func (num Number) String() string {
	return fmt.Sprintf("%d", num)
}

type Float float64

func (num Float) String() string {
	return fmt.Sprintf("%0.1f", num)
}

func inorderOperator(val Float) {
	fmt.Println(val.String())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Generate a random search tree
	randomSearchTree := bst.BinarySearchTree[Float]{nil, 0}
	for i := 0; i < 30; i++ {
		rn := 1.0 + 99.0 * rand.Float64()
		randomSearchTree.Insert(Float(rn))
	}
	time.Sleep(3 * time.Second)
	bst.ShowTreeGraph(randomSearchTree)
	randomSearchTree.InOrderTraverse(inorderOperator)
	min := randomSearchTree.Min()
	max, _ := randomSearchTree.Max()
	fmt.Printf("\nMinimum value in random search tree is %0.1f  \nMaximum value in random search tree is %0.1f", *min, *max)
	
	start := time.Now()
	tree := bst.BinarySearchTree[Number]{nil, 0}
	for val := 0; val < 100_000; val++ {
		tree.Insert(Number(val))
	}
	elapsed := time.Since(start)
	_, ht := tree.Max()
	fmt.Printf("\nTime to build BST tree with 100,000 nodes in sequential order: %s. Height of tree: %d", elapsed, ht)
}
/* Output
1.2
4.4
6.9
7.7
13.8
14.7
17.3
17.9
20.8
21.2
24.6
25.0
25.1
30.2
33.6
33.9
38.0
46.5
47.0
56.1
56.5
57.2
57.4
60.7
70.5
72.6
75.5
83.3
92.1
94.5

Minimum value in random search tree is 1.2  
Maximum value in random search tree is 94.5
Time to build BST tree with 100,000 nodes in sequential order: 35.645312291s. Height of tree: 100000
*/
