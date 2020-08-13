package maze

import (
	"math/rand"
	"time"
)

//GenerateDepthFirst generates a maze using an iterative randomized depth-first search algorithm
func (m *Maze) GenerateDepthFirst() {

	rand.Seed(time.Now().UTC().UnixNano())
	cellStack := new(Stack)
	current, _ := m.At(0, 0)

	current.visited = true

	for m.existsUnvisited() {
		next := m.getUnvisited(current)
		if next != nil {
			cellStack.Push(current)
			m.BreakWallBetween(current, next)
			next.visited = true
			current = next
		} else if !cellStack.IsEmpty() {
			current = cellStack.Pop()
		}
	}
}

//GenerateAB generates the maze using the Aldous-Broder algorithm
func (m *Maze) GenerateAB() {

	rand.Seed(time.Now().UTC().UnixNano())

	current, err := m.At(rand.Intn(m.rows), rand.Intn(m.cols))
	if err != nil {
		panic(err)
	}
	for m.existsUnvisited() {
		c := m.getRandom(current)
		if !c.visited {
			m.BreakWallBetween(current, c)
			c.visited = true
		}
		current = c
	}
}

func (m *Maze) getRandom(cell *Cell) *Cell {

	var unvisited []*Cell
	if c, err := m.At(cell.row+1, cell.col); err == nil {
		unvisited = append(unvisited, c)
	}
	if c, err := m.At(cell.row-1, cell.col); err == nil {
		unvisited = append(unvisited, c)
	}
	if c, err := m.At(cell.row, cell.col+1); err == nil {
		unvisited = append(unvisited, c)
	}
	if c, err := m.At(cell.row, cell.col-1); err == nil {
		unvisited = append(unvisited, c)
	}

	if len(unvisited) == 0 {
		return nil
	}
	i := rand.Intn(len(unvisited))
	return unvisited[i]
}

func (m *Maze) getUnvisited(cell *Cell) *Cell {

	var unvisited []*Cell
	if c, err := m.At(cell.row+1, cell.col); err == nil && !c.visited {
		unvisited = append(unvisited, c)
	}
	if c, err := m.At(cell.row-1, cell.col); err == nil && !c.visited {
		unvisited = append(unvisited, c)
	}
	if c, err := m.At(cell.row, cell.col+1); err == nil && !c.visited {
		unvisited = append(unvisited, c)
	}
	if c, err := m.At(cell.row, cell.col-1); err == nil && !c.visited {
		unvisited = append(unvisited, c)
	}

	if len(unvisited) == 0 {
		return nil
	}
	i := rand.Intn(len(unvisited))
	return unvisited[i]
}

func (m *Maze) existsUnvisited() bool {
	for i := range m.arr {
		for j := range m.arr[i] {
			if m.arr[i][j].visited == false {
				return true
			}
		}
	}
	return false
}
