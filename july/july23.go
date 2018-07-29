package july

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/util"
)

// July23 ...
func July23() {
	const (
		outFileName   = "out/july23.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		gridSize      = 1.0
		rows          = int(height) / int(gridSize)
		cols          = int(width) / int(gridSize)
	)

	makeGrid := func(rows, cols int) [][]int {
		grid := make([][]int, cols)
		for i := 0; i < cols; i++ {
			grid[i] = make([]int, rows)
		}
		return grid
	}

	drawGrid := func(grid [][]int, surface *blgo.Surface) {
		for x, col := range grid {
			for y, cell := range col {
				if cell == 1 {
					surface.FillRectangle(float64(x+(y-rows/2)/2), float64(cols/2)+0.866*float64(y-cols/2), 1, 1)
				}
			}
		}
	}

	checkCell := func(grid [][]int, x, y int) bool {
		sum := 0
		yy := y - 1
		sum += grid[x][yy]
		sum += grid[x+1][yy]
		yy = y
		sum += grid[x-1][yy]
		sum += grid[x+1][yy]
		yy = y + 1
		sum += grid[x-1][yy]
		sum += grid[x][yy]
		return sum == 1 || sum == 2
	}

	copyGrid := func(grid, temp [][]int) {
		for x, col := range grid {
			for y, cell := range col {
				if cell == 1 {
					grid[x][y] = 1
				}
			}
		}
	}

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		surface.ClearRGB(0, 0, 0)
		numIter := int(percent * 200)
		grid := makeGrid(rows, cols)
		grid[rows/2][cols/2] = 1

		for iter := 0; iter < numIter; iter++ {
			surface.Save()
			surface.ClearRGB(1, 1, 1)
			surface.SetLineWidth(0.25)
			surface.Scale(gridSize, gridSize)
			drawGrid(grid, surface)
			temp := makeGrid(rows, cols)
			for x := 1; x < rows-1; x++ {
				for y := 1; y < cols-1; y++ {
					result := checkCell(grid, x, y)
					if result {
						temp[x][y] = 1
					}
				}
			}
			grid = temp
			surface.Restore()
			copyGrid(grid, temp)
		}
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
