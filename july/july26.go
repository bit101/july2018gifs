package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

// July26 ...
func July26() {
	const (
		outFileName   = "out/july26.gif"
		timeInSeconds = 3
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		size          = 20.0
	)

	render := func(surface *blgo.Surface, percent float64) {
		grid := func(s *blgo.Surface, w, h, tx, ty float64) {
			t := blmath.LerpSin(percent+tx, 0.125, 0.5)
			tt := blmath.LerpSin(percent+ty, 0.125, 0.5)
			s.Save()
			s.Translate(w/2, h/2)
			s.Rotate(percent * math.Pi)
			s.FillEllipse(0, 0, w*t, h*tt)
			s.Restore()
		}

		surface.ClearRGB(1, 1, 1)
		RenderGrid(surface, 40, 40, grid)
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}

// RenderGridCallback defines the callback function for RenderGrid
type RenderGridCallback func(s *blgo.Surface, w, h, tx, ty float64)

// RenderGrid allows you to render multiple drawings in a grid via a RenderGridCallback
func RenderGrid(s *blgo.Surface, cellW, cellH float64, callback RenderGridCallback) {
	for x := 0.0; x < s.Width; x += cellW {
		for y := 0.0; y < s.Height; y += cellH {
			s.Save()
			s.Translate(x, y)
			callback(s, cellW, cellH, x/s.Width, y/s.Height)
			s.Restore()
		}
	}
}
