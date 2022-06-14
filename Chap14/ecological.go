package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"sync"
)

const (
	numRows int = 50
	numCols int = 50
	MAKERELREPRO int = 4
	MAKERELSTARVE int = 10000000
	MAKERELLIFE int = 30
	TUNAREPRO int = 8
	TUNASTARVE int = 11
	TUNALIFE int = 18
	SHARKREPRO int = 15
	SHARKSTARVE int = 25
	SHARKLIFE int = 30
)

var (
	quit 	   bool
	contain    *fyne.Container
	rect       *canvas.Rectangle
	mutex = &sync.Mutex{}
	// Holds rectangle objects
	segments  = make([]fyne.CanvasObject, numRows * numCols)  
)

type Location struct {
	x       int
	y       int
	critter MarineLife
}

type MarineLife interface {
	Move()
	Reproduce(l Location)
	Starve() bool
	LifeOver() bool
}

type Tuna struct {
	repro int // Moves til reproduction
	starv int // Movew til starvation 
	life  int // Moves til life over
	x, y int  // Set x to -1, y = -1 if dead
}

type Shark struct {
	repro int 
	starv int 
	life  int 
	x, y int  // Set x to -1, y = -1 if dead
}

type Mackerel struct {
	repro int 
	starv int 
	life  int 
	x, y int  // Set x to -1, y = -1 if dead
}

var locations [numRows][numCols]Location

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func distanceOfOne(x1, y1, x2, y2 float64) bool {
	return (math.Abs(x2-x1) == 0 && math.Abs(y2-y1) == 1) ||
		(math.Abs(x2-x1) == 1 && math.Abs(y2-y1) == 0) ||
		(math.Abs(x2-x1) == 1 && math.Abs(y2-y1) == 1)

}

func initializeLocations() {
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			locations[row][col] = Location{col, row, nil}
		}
	}
}

func findRandomCritter(x int, y int, critter MarineLife) (bool, Location) {
	// Send in nil for critter to get random empty location
	result := []Location{}
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			d := distanceOfOne(float64(x), float64(y), float64(c), float64(r))
			if d == true && reflect.TypeOf(locations[c][r].critter) == reflect.TypeOf(critter) {
				result = append(result, Location{r, c, critter})
			}
		}
	}
	if len(result) == 0 {
		return false, Location{}
	} else {
		return true, result[rand.Intn(len(result))]
	}
}

func (tuna *Tuna) Move() {
	for ; quit == false ; {
		if tuna.x == -1 { // Tuna no longer alive
			break
		}
		mutex.Lock()
		tuna.repro -= 1
		tuna.starv -= 1
		tuna.life -= 1
		if tuna.LifeOver() || tuna.Starve() {
			locations[tuna.y][tuna.x].critter = nil
			tuna.x = -1
			tuna.y = -1
			mutex.Unlock()
			break
		}
		// Find random neighbor that is a Mackerel
		found, newLoc := findRandomCritter(tuna.x, tuna.y, new(Mackerel))
		if found == true {
			fmt.Printf("\nTuna Move from <%d, %d> to <%d, %d>", tuna.x, tuna.y, newLoc.x, newLoc.y)
			tuna.starv = TUNASTARVE
			// Must stop go routine for makerek that was eaten
			eatenMackerel := locations[newLoc.y][newLoc.x].critter.(*Mackerel) // Type assertion
			eatenMackerel.x = -1
			eatenMackerel.y = -1
			fmt.Printf("\nEaten mackerel = %v", eatenMackerel)
			tuna.Reproduce(newLoc)
		}
		found, newLoc = findRandomCritter(tuna.x, tuna.y, nil)
		if found == true {
			fmt.Printf("\nTuna Move from <%d, %d> to <%d, %d>", tuna.x, tuna.y, newLoc.x, newLoc.y)
			tuna.Reproduce(newLoc)
		}
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(500) + 500) * time.Millisecond)
	}
}

func (shark *Shark) Move() {
	for ; quit == false ; {
		if shark.x == -1 { // Shark no longer alive
			break
		}
		mutex.Lock()
		shark.repro -= 1
		shark.starv -= 1
		shark.life -= 1
		if shark.LifeOver() || shark.Starve() {
			locations[shark.y][shark.x].critter = nil
			shark.x = -1
			shark.y = -1
			mutex.Unlock()
			break
		}
		// Find random neighbor that has no critter
		found, newLoc := findRandomCritter(shark.x, shark.y, new(Tuna))
		if found == true {
			fmt.Printf("\nShark Move from <%d, %d> to <%d, %d>", shark.x, shark.y, newLoc.x, newLoc.y)
			shark.starv = SHARKSTARVE
			// Must stop go routine for tuna that was eaten
			eatenTuna := locations[newLoc.y][newLoc.x].critter.(*Tuna) // Type assertion
			eatenTuna.x = -1
			eatenTuna.y = -1
			fmt.Printf("\nEaten tuna = %v", eatenTuna)
			shark.Reproduce(newLoc)
		} else {
			found, newLoc = findRandomCritter(shark.x, shark.y, nil)
			if found == true {
				fmt.Printf("\nShark Move from <%d, %d> to <%d, %d>", shark.x, shark.y, newLoc.x, newLoc.y)
				shark.Reproduce(newLoc)
			}
		}
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(500) + 500) * time.Millisecond)
	}
}

func (mackerel *Mackerel) Move() {
	for ; quit == false ; {
		if mackerel.x == -1 { // mackerel has been killed
			break
		}
		mutex.Lock()
		mackerel.repro -= 1
		mackerel.starv -= 1
		mackerel.life -= 1
		if mackerel.LifeOver() || mackerel.Starve() {
			locations[mackerel.y][mackerel.x].critter = nil
			mackerel.x = -1
			mackerel.y = -1
			mutex.Unlock()
			break
		}
		// Find random neighbor that has no critter
		found, newLoc := findRandomCritter(mackerel.x, mackerel.y, nil)
		if found == true {
			fmt.Printf("\nMackerel Move from <%d, %d> to <%d, %d>", mackerel.x, mackerel.y, newLoc.x, newLoc.y)
			mackerel.Reproduce(newLoc)
		} 
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(500) + 500) * time.Millisecond)
	}
}

func (tuna Tuna) Starve() bool {
	return tuna.starv <= 0
}

func (tuna Tuna) LifeOver() bool {
	return tuna.life <= 0
}

func (shark Shark) Starve() bool {
	return shark.starv <= 0
}

func (shark Shark) LifeOver() bool {
	return shark.life <= 0
}

func (mackerel Mackerel) Starve() bool {
	return mackerel.starv <= 0
}

func (mackerel Mackerel) LifeOver() bool {
	return mackerel.life <= 0
}

func (tuna *Tuna) Reproduce(l Location) {
	if tuna.x == -1 {
		return
	}
	if tuna.repro <= 0 {
		newTuna := new(Tuna)
		newTuna.repro = TUNAREPRO
		newTuna.starv = TUNASTARVE
		newTuna.life = TUNALIFE
		newTuna.x = tuna.x 
		newTuna.y = tuna.y
		locations[tuna.y][tuna.x].critter = newTuna // add newTuna to old location
		go newTuna.Move()
	} else {
		locations[tuna.y][tuna.x].critter = nil // remove tuna from previous location
	}
	tuna.x = l.x // assign tuna to new location
	tuna.y = l.y
	locations[l.y][l.x].critter = tuna // add tuna to new location
}

func (shark *Shark) Reproduce(l Location) {
	if shark.x == -1 {
		return
	}
	if shark.repro <= 0 {
		newShark := new(Shark)
		newShark.repro = SHARKREPRO
		newShark.starv = SHARKSTARVE
		newShark.life = SHARKLIFE
		newShark.x = shark.x 
		newShark.y = shark.y
		locations[shark.y][shark.x].critter = newShark // add newShark to old location
		go newShark.Move()
	} else {
		locations[shark.y][shark.x].critter = nil // remove shark from previous location
	}
	shark.x = l.x // assign shark to new location
	shark.y = l.y
	locations[l.y][l.x].critter = shark // add shark to new location
}

func (mackerel *Mackerel) Reproduce(l Location) {
	if mackerel.x == -1 {
		return
	}
	if mackerel.repro <= 0 {
		newMackerel := new(Mackerel)
		newMackerel.repro = MAKERELREPRO
		newMackerel.starv = MAKERELSTARVE // Never starves
		newMackerel.life = MAKERELLIFE
		newMackerel.x = mackerel.x 
		newMackerel.y = mackerel.y
		locations[mackerel.y][mackerel.x].critter = newMackerel // add newMackerel to old location
		go newMackerel.Move()
	} else {
		locations[mackerel.y][mackerel.x].critter = nil // remove mackerel from previous location
	}
	mackerel.x = l.x // assign mackerel to new location
	mackerel.y = l.y
	locations[l.y][l.x].critter = mackerel // add mackerel to new location
}

func output() *fyne.Container {
	for col := 0; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			if locations[col][row].critter == nil {
				rect = canvas.NewRectangle(&color.RGBA{B: 200, R: 200, G: 200, A: 255})
			} else  if reflect.TypeOf(locations[col][row].critter) == reflect.TypeOf(new(Tuna)) {
				rect = canvas.NewRectangle(&color.RGBA{B: 255, R: 0, G: 0, A: 255})
			} else if reflect.TypeOf(locations[col][row].critter) == reflect.TypeOf(new(Shark)) {
				rect = canvas.NewRectangle(&color.RGBA{B: 0, R: 255, G: 0, A: 255})
			} else if reflect.TypeOf(locations[col][row].critter) == reflect.TypeOf(new(Mackerel)) {
				rect = canvas.NewRectangle(&color.RGBA{B: 0, R: 0, G: 255, A: 255})
			}
			rect.Resize(fyne.NewSize(10, 10))
			rect.Move(fyne.NewPos(float32(col * 11), float32(row * 11)))
			// segments = append(segments, rect)
			segments[col +  numCols * row] = rect
		}
	}
	return container.NewWithoutLayout(segments...)
}

func main() {
	quit = false
	a := app.New()
	w := a.NewWindow("Ecological Simulation - Type Any Key To Quit")
	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)

	initializeLocations()

	
	newTuna := new(Tuna)
	newTuna.repro = TUNAREPRO
	newTuna.starv = TUNASTARVE
	newTuna.life = TUNALIFE
	newTuna.x = 15
	newTuna.y = 15
	locations[15][15].critter = newTuna
	go newTuna.Move()
	
	newTuna = new(Tuna)
	newTuna.repro = TUNAREPRO 
	newTuna.starv = TUNASTARVE
	newTuna.life = TUNALIFE
	newTuna.x = 19
	newTuna.y = 19
	locations[19][19].critter = newTuna
	go newTuna.Move()

	newTuna = new(Tuna)
	newTuna.repro = TUNAREPRO
	newTuna.starv = TUNASTARVE
	newTuna.life = TUNALIFE
	newTuna.x = 4
	newTuna.y = 4
	locations[4][4].critter = newTuna
	go newTuna.Move()

	newShark := new(Shark)
	newShark.repro = SHARKREPRO
	newShark.starv = SHARKSTARVE
	newShark.life = SHARKLIFE
	newShark.x = 11
	newShark.y = 11
	locations[11][11].critter = newShark
	go newShark.Move()

	newShark = new(Shark)
	newShark.repro = SHARKREPRO
	newShark.starv = SHARKSTARVE
	newShark.life = SHARKLIFE
	newShark.x = 16
	newShark.y = 16
	locations[16][16].critter = newShark
	go newShark.Move()

	newMackerel := new(Mackerel)
	newMackerel.repro = MAKERELREPRO
	newMackerel.starv = MAKERELSTARVE
	newMackerel.life = MAKERELLIFE
	newMackerel.x = 2
	newMackerel.y = 2
	locations[2][2].critter = newMackerel
	go newMackerel.Move()

	newMackerel = new(Mackerel)
	newMackerel.repro = MAKERELREPRO
	newMackerel.starv = MAKERELSTARVE
	newMackerel.life = MAKERELLIFE
	newMackerel.x = 13
	newMackerel.y = 8
	locations[8][13].critter = newMackerel
	go newMackerel.Move()

	newMackerel = new(Mackerel)
	newMackerel.repro = MAKERELREPRO
	newMackerel.starv = MAKERELSTARVE
	newMackerel.life = MAKERELLIFE
	newMackerel.x = 16
	newMackerel.y = 16
	locations[16][16].critter = newMackerel
	go newMackerel.Move()

	newMackerel = new(Mackerel)
	newMackerel.repro = MAKERELREPRO
	newMackerel.starv = MAKERELSTARVE
	newMackerel.life = MAKERELLIFE
	newMackerel.x = 28
	newMackerel.y = 28
	locations[28][28].critter = newMackerel
	go newMackerel.Move()
	
	go func() {
		for ; ; {
			mutex.Lock()
			contain := output()
			mutex.Unlock()
			w.SetContent(contain)
			time.Sleep(1000  * time.Millisecond)
		}
	}()

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) { // Shuts down simulation
		quit = true
        w.Close()
    })
	
	w.ShowAndRun()
}
