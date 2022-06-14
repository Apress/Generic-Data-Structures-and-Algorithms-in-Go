package main

import bt"example.com/binarytree"


func main() {
	root := bt.Node{"A", nil, nil}
	nodeB := bt.Node{"B",nil, nil}
	nodeC := bt.Node{"C", nil, nil}
	nodeD := bt.Node{"D", nil, nil}
	nodeE := bt.Node{"E", nil, nil}
	nodeF := bt.Node{"F",nil, nil}
	nodeG := bt.Node{"G", nil, nil}
	nodeH := bt.Node{"H", nil, nil}
	nodeI := bt.Node{"I", nil, nil}
	nodeJ := bt.Node{"J", nil, nil}
	nodeK := bt.Node{"K", nil, nil}
	nodeL := bt.Node{"L", nil, nil}
	nodeM := bt.Node{"M", nil, nil}
	nodeN := bt.Node{"N", nil, nil}
	nodeO := bt.Node{"O", nil, nil}
	nodeP := bt.Node{"P", nil, nil}
	nodeQ := bt.Node{"Q", nil, nil}
	nodeR := bt.Node{"R", nil, nil}


	root.Left = &nodeB
	root.Right = &nodeC
	nodeB.Left = &nodeD
	nodeD.Right = &nodeH
	nodeD.Left = &nodeE
	nodeE.Left = &nodeF
	nodeE.Right = &nodeG
	nodeC.Right = &nodeI
	nodeC.Left = &nodeJ
	nodeI.Right = &nodeK 
	nodeK.Left = &nodeL
	nodeL.Left = &nodeM
	nodeL.Right = &nodeN
	nodeN.Right = &nodeO
	nodeO.Left = &nodeP
	nodeO.Right = &nodeQ
	nodeM.Left = &nodeR
	myTree := bt.BinaryTree{&root, 18}
	bt.ShowTreeGraph(myTree)
}