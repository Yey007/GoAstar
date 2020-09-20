package maze

//Cell stores data about a cell, ie. where there are Walls and where there is space
type Cell struct {
	Walls    []bool //0 to 3 clockwise, starting on the left
	Visited  bool
	InPath 	 bool
	Row, Col int
	Parent   *Cell
}

//NewCell creates a new maze with the given size
func NewCell(row, col int) *Cell {
	c := new(Cell)
	c.Walls = make([]bool, 4)
	for i := range c.Walls {
		c.Walls[i] = true
	}
	c.Visited = false
	c.InPath = false
	c.Row = row
	c.Col = col
	c.Parent = nil
	return c
}

//SetWall sets a side of the cell to be a wall
func (c *Cell) SetWall(wall int) {
	c.Walls[wall] = true
}

//BreakWall sets a side of the cell to be air
func (c *Cell) BreakWall(wall int) {
	c.Walls[wall] = false
}

//HasWall returns whether the cell has a wall in a certain direction
func (c *Cell) HasWall(wall int) bool {
	return c.Walls[wall]
}
