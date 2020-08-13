package maze

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

//Maze stores data about a maze, ie. where there are walls and where there is space
type Maze struct {
	rows int
	cols int
	arr  [][]Cell
}

//NewMaze creates a new maze with the given size
func NewMaze(rows, cols int) *Maze {
	m := new(Maze)
	m.arr = make([][]Cell, rows)
	for i := range m.arr {
		m.arr[i] = make([]Cell, cols)
		for j := range m.arr[i] {
			m.arr[i][j] = *NewCell(i, j)
		}
	}
	m.rows = rows
	m.cols = cols
	return m
}

//Draw draws the current state of the maze to an image
func (m *Maze) Draw(i int) {
	scaleFactor := 10
	imgPtrX, imgPtrY := 0, 0
	base := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{m.cols*scaleFactor + 1, m.rows*scaleFactor + 1}})

	f, err := os.Create("maze" + strconv.Itoa(i) + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < m.rows*scaleFactor+1; i++ {
		for j := 0; j < m.cols*scaleFactor+1; j++ {
			base.Set(j, i, color.Black) //these are swapped by intention, this function is weird
		}
	}

	for i := range m.arr {
		for j := range m.arr[i] {
			c := m.arr[i][j]
			if c.HasWall(0) {
				for k := 0; k < scaleFactor; k++ {
					base.Set(imgPtrX, k+imgPtrY, color.White)
				}
			}
			if c.HasWall(1) {
				for k := 0; k < scaleFactor; k++ {
					base.Set(k+imgPtrX, imgPtrY, color.White)
				}
			}
			if c.HasWall(2) {
				for k := 0; k < scaleFactor+1; k++ { //+1 makes it look gooder
					base.Set(imgPtrX+scaleFactor, k+imgPtrY, color.White)
				}
			}
			if c.HasWall(3) {
				for k := 0; k < scaleFactor; k++ {
					base.Set(k+imgPtrX, imgPtrY+scaleFactor, color.White)
				}
			}
			imgPtrX += scaleFactor
		}
		imgPtrX = 0
		imgPtrY += scaleFactor
	}

	png.Encode(f, base)
}

//SetWallBetween sets the wall between two cell
func (m *Maze) SetWallBetween(first, second *Cell) error {
	if first.row == second.row {
		if first.col == second.col-1 { //if second is to the left first
			first.SetWall(2)
			second.SetWall(0)
			return nil
		} else if first.col == second.col+1 { //if second is to the right first
			first.SetWall(0)
			second.SetWall(2)
			return nil
		}
	} else if first.col == second.col {
		if first.row == second.row-1 { //if second is above first
			first.SetWall(3)
			second.SetWall(1)
			return nil
		} else if first.row == second.row+1 { //if second is below first
			first.SetWall(1)
			second.SetWall(3)
			return nil
		}
	}
	return errors.New("Impossible wall set encountered")
}

//BreakWallBetween breaks the wall between two cell
func (m *Maze) BreakWallBetween(first, second *Cell) error {
	if first.row == second.row {
		if first.col == second.col-1 { //if second is to the left first
			first.BreakWall(2)
			second.BreakWall(0)
			return nil
		} else if first.col == second.col+1 { //if second is to the right first
			first.BreakWall(0)
			second.BreakWall(2)
			return nil
		}
	} else if first.col == second.col {
		if first.row == second.row-1 { //if second is above first
			first.BreakWall(3)
			second.BreakWall(1)
			return nil
		} else if first.row == second.row+1 { //if second is below first
			first.BreakWall(1)
			second.BreakWall(3)
			return nil
		}
	}
	return errors.New("Impossible wall set encountered")
}

//At gets a reference to the cell at the given coordinates
func (m *Maze) At(row, col int) (*Cell, error) {
	if row >= 0 && col >= 0 && row < m.rows && col < m.cols {
		return &m.arr[row][col], nil
	}
	return nil, errors.New("Index out of bounds")
}
