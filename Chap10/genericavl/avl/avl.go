package avl

import (
	"image/color"
	"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/mitchellh/go-homedir"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	// "constraints"
)

type ordered interface {
	~int | ~float64  | ~string
}

type AVLTree[T OrderedStringer] struct {
	Root *Node[T]
	NumNodes int
}

type Node[T OrderedStringer] struct {
	Value T
	Left *Node[T]
	Right *Node[T]
	Ht int
}

type OrderedStringer interface {
	ordered
	String() string
}

// Methods
func (avl *AVLTree[T]) Insert(newValue T) {
	if avl.Search(newValue) == false {
		avl.Root = insertNode(avl.Root, newValue)
		avl.NumNodes += 1
	}
}

func (avl *AVLTree[T]) Delete(value T) {
	if avl.Search(value) == true {
		avl.Root = deleteNode(avl.Root, value)
		avl.NumNodes -= 1
	}
}

func (avl *AVLTree[T]) Search(value T) bool {
    return search(avl.Root, value)
}

func (avl *AVLTree[T]) Height() int {
	return avl.Root.Height()
}

func (avl *AVLTree[T]) InOrderTraverse(f func(T)) {
    inOrderTraverse(avl.Root, f)
}

func (avl *AVLTree[T]) Min() *T {
    node := avl.Root
    if node == nil {
        return nil
    }
    for {
        if node.Left == nil {
            return &node.Value
        }
        node = node.Left
    }
}

func (avl *AVLTree[T]) Max() *T {
    node := avl.Root
    if node == nil {
        return nil
    }
    for {
        if node.Right == nil {
            return &node.Value
        }
        node = node.Right
    }
}

func (n *Node[T]) balanceFactor() int {
  	if n == nil {
    	return 0
  	}
  	return n.Left.Height() - n.Right.Height()
}

func (n *Node[T]) Height() int {
	if n == nil {
		return 0
	} else {
		return n.Ht
	}
}

func (n *Node[T]) updateHeight() {
  	max := func (a, b int) int {
    	if a > b {
      		return a
    	}
    	return b
  	}
  	n.Ht = max(n.Left.Height(), n.Right.Height()) + 1
}

// Internal functions
func newNode[T OrderedStringer](val T) *Node[T] {
  	return &Node[T] {
    	Value:  val,
    	Left:   nil,
    	Right:  nil,
    	Ht: 1,
  	}
}

func search[T OrderedStringer](n *Node[T], value T) bool {
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

func insertNode[T OrderedStringer](node *Node[T], val T) *Node[T] {
  	// if there's no node, create one
  	if node == nil {
    	return newNode(val)
  	}
  	// if value is greater than current node's value, insert to the right
  	if val > node.Value {
    	right := insertNode(node.Right, val)
    	node.Right = right
  	}
  	// if value is less than current node's value, insert to the left
  	if val < node.Value {
    	left:= insertNode(node.Left, val)
    		node.Left = left
  	}
  	return rotateInsert(node, val)
}

func rightRotate[T OrderedStringer](x *Node[T]) *Node[T] {
  	y := x.Left
  	t := y.Right

  	y.Right = x
  	x.Left = t

  	x.updateHeight()
  	y.updateHeight()

  	return y
}

func leftRotate[T OrderedStringer](x *Node[T]) *Node[T] {
  y := x.Right
  t := y.Left

  y.Left = x
  x.Right = t

  x.updateHeight()
  y.updateHeight()

  return y
}

func rotateInsert[T OrderedStringer](node *Node[T], val T) *Node[T] {
  node.updateHeight()

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

func inOrderTraverse[T OrderedStringer](n *Node[T], op func(T)) {
    if n != nil {
        inOrderTraverse(n.Left, op)
        op(n.Value)
        inOrderTraverse(n.Right, op)
    }
}

func largest[T OrderedStringer](node *Node[T]) *Node[T] {
  	if node == nil {
    	return nil
  	}

  	if node.Right == nil {
    	return node
  	}
  	return largest(node.Right)
}

func rotateDelete[T OrderedStringer](node *Node[T]) *Node[T] {
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

func deleteNode[T OrderedStringer](node *Node[T], val T) *Node[T] {
	if node == nil {
    	return nil
  	}

  	if val > node.Value {
    	right  := deleteNode(node.Right, val)
    	node.Right = right
  	} else if val < node.Value {
    	left := deleteNode(node.Left, val)
    	node.Left = left
  	} else { 
    	if node.Left != nil && node.Right != nil {
      		// has 2 children, find the successor
      		successor := largest(node.Left)
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


// Logic for drawing tree
type NodePair struct {
	Val1, Val2 string
}

type NodePos struct {
	Val string
	YPos int
	XPos int
}

var data []NodePos
var endPoints []NodePair // Used to plot lines

func PrepareDrawTree[T OrderedStringer](tree AVLTree[T]) {	
	prepareToDraw(tree)
}

func FindXY(val interface{}) (int, int) {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return data[i].XPos, data[i].YPos
		}
	}
	return -1, -1
}

func FindX(val interface{}) int {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return i
		}
	}
	return -1
}

func SetXValues() {
	for index := 0; index < len(data); index++ {
		xValue := FindX(data[index].Val)
		data[index].XPos = xValue
	}
}

func prepareToDraw[T OrderedStringer](tree AVLTree[T]) {
	inorderLevel(tree.Root, 1)
	SetXValues()
	getEndPoints(tree.Root, nil)
}

func inorderLevel[T OrderedStringer](node *Node[T], level int) {
	if node != nil {
		inorderLevel(node.Left, level + 1)
		data = append(data, NodePos{node.Value.String(), 100 - level, -1})
		inorderLevel(node.Right, level + 1)
	}
}

func getEndPoints[T OrderedStringer](node *Node[T], parent *Node[T]) {
	if node != nil {
		if parent != nil {
			endPoints = append(endPoints, NodePair{node.Value.String(), parent.Value.String()})
		}
		getEndPoints(node.Left, node)
		getEndPoints(node.Right, node)
	}
}

var path string

func DrawGraph(a fyne.App, w fyne.Window) {
	image := canvas.NewImageFromResource(theme.FyneLogo())
	image = canvas.NewImageFromFile(path + "tree.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.Close()
	w.Show()
}

func ShowTreeGraph[T OrderedStringer](myTree AVLTree[T]) {
	PrepareDrawTree(myTree)
	myApp := app.New()
	myWindow := myApp.NewWindow("Tree")
	myWindow.Resize(fyne.NewSize(1000, 600))
	path, _ := homedir.Dir()
	path += "/Desktop//"

	nodePts := make(plotter.XYs, myTree.NumNodes)
	for i := 0; i < len(data); i++ {
		nodePts[i].Y = float64(data[i].YPos)
		nodePts[i].X = float64(data[i].XPos)
	}
	nodePtsData := nodePts
	p := plot.New()
	p.Add(plotter.NewGrid())
	nodePoints, err := plotter.NewScatter(nodePtsData)
	if err != nil {
		log.Panic(err)
	}
	nodePoints.Shape = draw.CircleGlyph{}
	nodePoints.Color = color.RGBA{G: 255, A: 255}
	nodePoints.Radius = vg.Points(12)

	// Plot lines
	for index := 0; index < len(endPoints); index++ {
		val1 := endPoints[index].Val1
		x1, y1 := FindXY(val1)
		val2 := endPoints[index].Val2
		x2, y2 := FindXY(val2)
		pts := plotter.XYs{{X: float64(x1), Y: float64(y1)}, {X: float64(x2), Y: float64(y2)}}
		line, err := plotter.NewLine(pts)
		if err != nil {
			log.Panic(err)
		}
		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			log.Panic(err)
		}
		p.Add(line, scatter)
	}
	
	p.Add(nodePoints)

	// Add Labels
	for index := 0; index < len(data); index++ {
		x := float64(data[index].XPos) - 0.075 
		y := float64(data[index].YPos) - 0.02
		str := data[index].Val
		label, err := plotter.NewLabels(plotter.XYLabels {
			XYs: []plotter.XY {
				{X: x ,Y: y},
			},
			Labels: []string{str},
			},)
		if err != nil {
			log.Fatalf("could not creates labels plotter: %+v", err)
		}
		p.Add(label)
	}

	path, _ = homedir.Dir()
	path += "/Desktop/GoDS/"
	err = p.Save(1000, 600, "tree.png")
	if err != nil {
		log.Panic(err)
	}

	DrawGraph(myApp, myWindow)

	myWindow.ShowAndRun()
}