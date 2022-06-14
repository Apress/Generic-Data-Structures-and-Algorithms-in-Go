package binarytree

import (
	"fmt"
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
)

type BinaryTree struct {
	Root *Node
	NumNodes int
}

type Node struct {
	Value string
	Left *Node
	Right *Node
}

type nodePair struct {
	Val1, Val2 string
}

type nodePos struct {
	Val string
	YPos int
	XPos int
}

var data []nodePos    		// Used to get node positions (Val, XPos, YPos) of each node
var endPoints []nodePair 	// Used to plot lines

func prepareDrawTree(tree BinaryTree) {	
	prepareToDraw(tree)
	fmt.Printf("\nslice of endPoints: %v", endPoints)
	fmt.Printf("\nslice of data: %v", data)
}

func findXY(val string) (int, int) {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return data[i].XPos, data[i].YPos
		}
	}
	return -1, -1
}

func findX(val string) int {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return i
		}
	}
	return -1
}

func setXValues() {
	for index := 0; index < len(data); index++ {
		xValue := findX(data[index].Val)
		data[index].XPos = xValue
	}
}

func prepareToDraw(tree BinaryTree) {
	inorderLevel(tree.Root, 1)
	setXValues()
	getEndPoints(tree.Root, nil)
}

func inorderLevel(node *Node, level int) {
	if node != nil {
		inorderLevel(node.Left, level + 1)
		data = append(data, nodePos{node.Value, 100 - level, -1})
		inorderLevel(node.Right, level + 1)
	}
}

func getEndPoints(node *Node, parent *Node) {
	if node != nil {
		if parent != nil {
			endPoints = append(endPoints, nodePair{node.Value, parent.Value})
		}
		getEndPoints(node.Left, node)
		getEndPoints(node.Right, node)
	}
}

var path string

func drawGraph(a fyne.App, w fyne.Window) {
	image := canvas.NewImageFromResource(theme.FyneLogo())
	image = canvas.NewImageFromFile(path + "tree.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	// w.Close()
	w.Show()
}

func ShowTreeGraph(myTree BinaryTree) {
	prepareDrawTree(myTree)
	myApp := app.New()
	myWindow := myApp.NewWindow("Binary Tree")
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
		x1, y1 := findXY(val1)
		val2 := endPoints[index].Val2
		x2, y2 := findXY(val2)
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
		x := float64(data[index].XPos) - 0.05
		y := float64(data[index].YPos) - 0.02
		str := data[index].Val
		label, err := plotter.NewLabels(plotter.XYLabels {
			XYs: []plotter.XY {
				{X: x ,Y: y},
			},
			Labels: []string{str},
			},)
		if err != nil {
			log.Fatalf("Could not creates labels plotter: %+v", err)
		}
		p.Add(label)
	}

	path, _ = homedir.Dir()
	path += "/Desktop/GoDS/"
	err = p.Save(1000, 600, "tree.png")
	if err != nil {
		log.Panic(err)
	}

	drawGraph(myApp, myWindow)

	myWindow.ShowAndRun()
}