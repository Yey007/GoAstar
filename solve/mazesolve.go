package solve

import (
	"math"
	"yey007.github.io/software/goastar/maze"
)

//AStar runs the A* algorithm on a given maze with given start and end points
func AStar(m *maze.Maze, start *maze.Cell, end *maze.Cell) {
	open := []*maze.Cell{start}
	cameFrom := make(map[*maze.Cell]*maze.Cell)
	fCosts := make(map[*maze.Cell]int)
	gCosts := make(map[*maze.Cell]int)

	cameFrom[start] = start

	gCosts[start] = 0
	fCosts[start] = computeHCost(start, end) //gCost will be 0

	for len(open) != 0 {
		current := findMinFCost(fCosts)

		if current == end {
			path := reconstruct_path(cameFrom, current)
			for _, c := range path {
				c.InPath = true
			}
		}

		open = remove(open, current)
		neighbors := getWalkableNeighbors(m, current)

		for _, n := range neighbors {

			gCosts[current] = computeGCost(start, current, cameFrom)
			gCosts[n] = computeGCost(start, n, cameFrom)

			tentative_gScore := gCosts[current] + 10
			if tentative_gScore < gCosts[n] {
				cameFrom[n] = current
				gCosts[n] = tentative_gScore
				fCosts[n] = gCosts[n] + computeHCost(n, end)
				if !in(open, n) {
					open = append(open, n)
				}
			}
		}
	}

}

func reconstruct_path(cameFrom map[*maze.Cell]*maze.Cell, current *maze.Cell) []*maze.Cell {
	totalpath := []*maze.Cell{current}
	for cameFrom[current] != nil {
		current = cameFrom[current]
		totalpath = append([]*maze.Cell{current}, totalpath...)
	}
	return totalpath
}

func getWalkableNeighbors(m *maze.Maze, c *maze.Cell) []*maze.Cell {
	list := m.GetNeighbors(c)
	var walkable []*maze.Cell
	for _, n := range list {
		if w, err := m.HasWallBetween(c, n); !w && err == nil {
			walkable = append(walkable, n)
		}
	}
	return walkable
}

func findMinFCost(fCosts map[*maze.Cell]int) *maze.Cell {

	var initial *maze.Cell
	for k, _ := range fCosts {
		initial = k
		break
	}

	minCost := fCosts[initial]
	minCell := initial
	for k, v := range fCosts {
		if v < minCost {
			minCost = v
			minCell = k
		}
	}

	return minCell
}

func computeGCost(start *maze.Cell, cell *maze.Cell, cameFrom map[*maze.Cell]*maze.Cell) int {
	totalcost := 0
	current := cell
	for cameFrom[current] != start {
		current = cameFrom[current]
		totalcost += 10 //arbitrary value
	}
	return totalcost
}

func computeHCost(cell *maze.Cell, end *maze.Cell)int {
	//Vetical dist + horizontal dist
	return int(math.Abs(float64(cell.Col - end.Col)) + math.Abs(float64(cell.Row - end.Col)))
}

func remove(s []*maze.Cell, cell *maze.Cell) []*maze.Cell {

	i := 0
	for index, c := range s {
		i = index
		if c == cell {
			break
		}
	}

	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func in(s []*maze.Cell, cell *maze.Cell) bool {
	for _, c := range s {
		if c == cell {
			return true
		}
	}
	return false
}
