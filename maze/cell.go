package maze

//Cell stores data about a cell, ie. where there are walls and where there is space
type Cell struct {
	walls    []bool //0 to 3 clockwise, starting on the left
	visited  bool
	row, col int
}

//NewCell creates a new maze with the given size
func NewCell(row, col int) *Cell {
	c := new(Cell)
	c.walls = make([]bool, 4)
	for i := range c.walls {
		c.walls[i] = true
	}
	c.visited = false
	c.row = row
	c.col = col
	return c
}

//SetWall sets a side of the cell to be a wall
func (c *Cell) SetWall(wall int) {
	c.walls[wall] = true
}

//BreakWall sets a side of the cell to be air
func (c *Cell) BreakWall(wall int) {
	c.walls[wall] = false
}

//HasWall returns whether the cell has a wall in a certain direction
func (c *Cell) HasWall(wall int) bool {
	return c.walls[wall]
}
