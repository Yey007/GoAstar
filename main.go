package main

import (
	"fmt"
	"time"
	"yey007.github.io/software/goastar/solve"

	"yey007.github.io/software/goastar/gen"

	"yey007.github.io/software/goastar/maze"
)

func main() {
	start := time.Now()
	m := maze.NewMaze(10, 10)
	gen.GenerateDepthFirst(m)
	s, _ := m.At(0, 0)
	e, _ := m.At(9, 9)
	solve.AStar(m, s, e)
	m.Draw(0)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
