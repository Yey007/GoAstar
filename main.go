package main

import (
	"fmt"
	"time"

	"yey007.github.io/software/goastar/maze"
)

func main() {
	start := time.Now()
	m := maze.NewMaze(10, 10)
	m.GenerateDepthFirst()
	m.Draw(0)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
