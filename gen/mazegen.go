package gen

import (
	"math/rand"
	"time"

	"yey007.github.io/software/goastar/maze"
	"yey007.github.io/software/goastar/util"
)

//GenerateDepthFirst generates a maze using an iterative randomized depth-first search algorithm
func GenerateDepthFirst(m *maze.Maze) {

	rand.Seed(time.Now().UTC().UnixNano())
	cellStack := new(util.Stack)
	current, _ := m.At(0, 0)

	current.Visited = true

	for existsUnvisited(m) {
		next := getUnvisited(m, current)
		if next != nil {
			cellStack.Push(current)
			m.BreakWallBetween(current, next)
			next.Visited = true
			current = next
		} else if !cellStack.IsEmpty() {
			current = cellStack.Pop()
		}
	}
}

//GenerateAB generates the maze using the Aldous-Broder algorithm
func GenerateAB(m *maze.Maze) {

	rand.Seed(time.Now().UTC().UnixNano())

	current, err := m.At(rand.Intn(m.Rows), rand.Intn(m.Cols))
	if err != nil {
		panic(err)
	}
	for existsUnvisited(m) {
		c := getRandom(m, current)
		if !c.Visited {
			m.BreakWallBetween(current, c)
			c.Visited = true
		}
		current = c
	}
}

func getRandom(m *maze.Maze, cell *maze.Cell) *maze.Cell {

	neighbors := m.GetNeighbors(cell)
	i := rand.Intn(len(neighbors))
	return neighbors[i]
}

func getUnvisited(m *maze.Maze, cell *maze.Cell) *maze.Cell {

	neighbors := m.GetNeighbors(cell)
	var unvisited []*maze.Cell

	for _, n := range neighbors {
		if !n.Visited {
			unvisited = append(unvisited, n)
		}
	}

	if len(unvisited) == 0 {
		return nil
	}
	i := rand.Intn(len(unvisited))
	return unvisited[i]
}

func existsUnvisited(m *maze.Maze) bool {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if c, _ := m.At(i, j); c.Visited == false {
				return true
			}
		}
	}
	return false
}
