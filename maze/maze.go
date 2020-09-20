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
	Rows int
	Cols int
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
	m.Rows = rows
	m.Cols = cols
	return m
}

//Draw draws the current state of the maze to an image
func (m *Maze) Draw(n int) {
	scaleFactor := 10
	imgPtrX, imgPtrY := 0, 0
	base := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{m.Cols*scaleFactor + 1, m.Rows*scaleFactor + 1}})

	os.Mkdir("out", os.ModeDir)
	f, err := os.Create("out/maze" + strconv.Itoa(n) + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < m.Rows*scaleFactor+1; i++ {
		for j := 0; j < m.Cols*scaleFactor+1; j++ {
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
			if c.InPath {
				base.Set(imgPtrX + (scaleFactor / 2), imgPtrY + (scaleFactor / 2), color.RGBA{
					R: 200,
					G: 0,
					B: 0,
				})
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
	if first.Row == second.Row {
		if first.Col == second.Col-1 { //if second is to the left first
			first.SetWall(2)
			second.SetWall(0)
			return nil
		} else if first.Col == second.Col+1 { //if second is to the right first
			first.SetWall(0)
			second.SetWall(2)
			return nil
		}
	} else if first.Col == second.Col {
		if first.Row == second.Row-1 { //if second is above first
			first.SetWall(3)
			second.SetWall(1)
			return nil
		} else if first.Row == second.Row+1 { //if second is below first
			first.SetWall(1)
			second.SetWall(3)
			return nil
		}
	}
	return errors.New("Impossible wall set encountered")
}

//BreakWallBetween breaks the wall between two cell
func (m *Maze) BreakWallBetween(first, second *Cell) error {
	if first.Row == second.Row {
		if first.Col == second.Col-1 { //if second is to the left first
			first.BreakWall(2)
			second.BreakWall(0)
			return nil
		} else if first.Col == second.Col+1 { //if second is to the right first
			first.BreakWall(0)
			second.BreakWall(2)
			return nil
		}
	} else if first.Col == second.Col {
		if first.Row == second.Row-1 { //if second is above first
			first.BreakWall(3)
			second.BreakWall(1)
			return nil
		} else if first.Row == second.Row+1 { //if second is below first
			first.BreakWall(1)
			second.BreakWall(3)
			return nil
		}
	}
	return errors.New("Impossible wall set encountered")
}

//HasWallBetween sets the wall between two cell
func (m *Maze) HasWallBetween(first, second *Cell) (bool, error) {
	if first.Row == second.Row {
		if first.Col == second.Col-1 { //if second is to the left first
			return first.HasWall(2) && second.HasWall(0), nil
		} else if first.Col == second.Col+1 { //if second is to the right first
			return first.HasWall(0) && second.HasWall(2), nil
		}
	} else if first.Col == second.Col {
		if first.Row == second.Row-1 { //if second is above first
			return first.HasWall(3) && second.HasWall(1), nil
		} else if first.Row == second.Row+1 { //if second is below first
			return first.HasWall(1) && second.HasWall(3), nil
		}
	}
	return false, errors.New("Impossible wall check encountered")
}

//At gets a reference to the cell at the given coordinates
func (m *Maze) At(row, col int) (*Cell, error) {
	if row >= 0 && col >= 0 && row < m.Rows && col < m.Cols {
		return &m.arr[row][col], nil
	}
	return nil, errors.New("Index out of bounds")
}

//GetNeighbors gets the neighbors of a given cell
func (m *Maze) GetNeighbors(c *Cell) []*Cell {

	var cells []*Cell
	if c, err := m.At(c.Row+1, c.Col); err == nil {
		cells = append(cells, c)
	}
	if c, err := m.At(c.Row-1, c.Col); err == nil {
		cells = append(cells, c)
	}
	if c, err := m.At(c.Row, c.Col+1); err == nil {
		cells = append(cells, c)
	}
	if c, err := m.At(c.Row, c.Col-1); err == nil {
		cells = append(cells, c)
	}
	return cells
}
