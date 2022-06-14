package main

import (
	avl "example.com/avl"
	"fmt"
	"math/rand"
	"time"
)

func inorderOperator(val Float) {
	// val *= val
	fmt.Println(val.String())
}

// Satisfies OrderedStringer because of ~float64
// Also satisfies OrderedStringer because of String() method below
type Float float64

func (num Float) String() string {
	return fmt.Sprintf("%0.1f", num)
}

type Integer int

func (num Integer) String() string {
	return fmt.Sprintf("%d", num)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var data [100_000]int
	for i := 0; i < 100_000; i++ {
		data[i] = rand.Intn(1_000_000)
	}

	// Generate a random search tree
	randomSearchTree := avl.AVLTree[Float]{nil, 0}
	for i := 0; i < 30; i++ {
		rn := 1.0 + 99.0 * rand.Float64()
		randomSearchTree.Insert(Float(rn))
	}
	time.Sleep(3 * time.Second)
	avl.ShowTreeGraph(randomSearchTree)

	randomSearchTree.InOrderTraverse(inorderOperator)
	min := randomSearchTree.Min()
	max := randomSearchTree.Max()
	fmt.Printf("\nMinimum value in tree is %0.1f  Maximum value in tree is %0.1f", *min, *max)
	
	start := time.Now()
	tree := avl.AVLTree[Integer]{nil, 0}
	for i := 0; i < 100_000; i++ {
		tree.Insert(Integer(data[i]))
	}
	elapsed := time.Since(start)
	fmt.Printf("\nInsertion time for AVL tree: %s.  Height of tree: %d", elapsed, tree.Height())

	start = time.Now()
	for i := 0; i < 100_000; i++ {
		_ = tree.Search(Integer(i))
	}
	elapsed = time.Since(start)
	fmt.Println("\nSearch time for AVL tree: ", elapsed)

}
/* Output
3.9
7.2
8.7
10.3
11.8
13.5
16.8
16.9
22.7
33.1
36.8
45.4
45.6
46.5
53.3
54.6
55.7
58.8
61.8
61.9
71.5
77.8
78.2
91.1
91.4
91.4
96.3
97.3
97.5
98.2

Minimum value in tree is 3.9  Maximum value in tree is 98.2
Time to build AVL tree with a million nodes: 391.125839ms.  Height of tree: 20
Time to search AVL tree with a million nodes:  106.542256ms
*/
