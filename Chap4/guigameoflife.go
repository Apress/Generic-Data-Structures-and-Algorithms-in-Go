package main

import (
	"math/rand"
	"time"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	rows int
	cols int
	rect *canvas.Rectangle
	// Holds rectangle objects
	segments []fyne.CanvasObject

)

type Grid [][]bool

var grid Grid
var newGrid Grid

func (g *Grid) initializeGrid(r, c int) {
	rows = r 
	cols = c
	*g = make([][]bool, rows)
	for row := 0; row < rows; row++ {
    	(*g)[row] = make([]bool, cols)
    }
}

func Copy(target [][]bool, source [][]bool) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			target[row][col] = source[row][col]
		}
	}
}

func (g Grid) bringAlive(row, col int) {
	g[row][col] = true
}

func (g Grid) kill(row, col int) {
	g[row][col] = false
}

func (g Grid) numberLiveNeighbors(row, col int) int {
	result := 0
	if row > 0 && g[row - 1][col] == true {
		result++
	}
	if row > 0 && col < cols - 1 && g[row - 1][col + 1] == true {
			result += 1
		}
	if col < cols - 1 && g[row][col + 1] == true {
		result += 1
	}
	if row < rows - 1 && col < cols -1  && g[row + 1][col + 1] == true {
		result += 1
	}
	if row < rows - 1 && g[row + 1][col] == true {
		result += 1
	}
	if row < rows - 1 && col > 0 && g[row + 1][col - 1] == true {
		result += 1
	}
	if col > 0 && g[row][col - 1] == true {
		result += 1
	}
	if row >  0 && col > 0 && g[row - 1][col - 1] == true {
		result += 1
	}
	return result
}


func (g Grid) evolveGrid() {
	Copy(newGrid, g)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveN := g.numberLiveNeighbors(row, col)
			// Rules 1 and 2
			if g[row][col] == true && (liveN < 2 || liveN >= 4) { 
				newGrid[row][col] = false 
			}
			// Rule 4
			if g[row][col] == false && liveN == 3 { 
				newGrid[row][col] = true 
			}
		}
	}
	Copy(g, newGrid) 
}

func output() *fyne.Container {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[col][row] == false {
				rect = canvas.NewRectangle(&color.RGBA{B: 200, R: 200, G: 200, A: 255})
			} else  {
				rect = canvas.NewRectangle(&color.RGBA{B: 0, R: 255, G: 0, A: 255})
			}
			rect.Resize(fyne.NewSize(10, 10))
			rect.Move(fyne.NewPos(float32(row * 11), float32(col * 11)))
			segments = append(segments, rect)
		}
	}
	return container.NewWithoutLayout(segments...)
}

func main() {
	grid.initializeGrid(25, 25)
	newGrid.initializeGrid(25, 25)
	
	for numberCritters := 0; numberCritters < 4; numberCritters++ {
		r := 5  + rand.Intn(10)
		c := 5 + rand.Intn(10)
		grid.bringAlive(r, c)
		grid.bringAlive(r + 1, c)
		grid.bringAlive(r + 1, c + 1)
		grid.bringAlive(r - 1, c)
		grid.bringAlive(r - 2, c - 1)
	}
	
	a := app.New()
	w := a.NewWindow("GAME OF LIFE - Hit Any Key To Quit")
	w.Resize(fyne.NewSize(300, 300))
	w.SetFixedSize(true)

	go func() {
		for ; ; {
			container := output()
			w.SetContent(container)
			time.Sleep(1  * time.Second)
			grid.evolveGrid()
		}
	}()

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) { 
		// Shuts down simulation
        w.Close()
    })
	
	w.ShowAndRun()
}

