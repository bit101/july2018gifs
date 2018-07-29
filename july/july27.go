package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
	cairo "github.com/bit101/go-cairo"
)

// July27 ...
func July27() {
	const (
		outFileName   = "out/july27.gif"
		timeInSeconds = 5
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		size          = 20.0
	)

	drawCell := func(mode int, surface *blgo.Surface) {
		if mode == 0 {
			surface.Arc(-size/2, -size/2, size/2, 0, math.Pi/2)
			surface.Stroke()
			surface.Arc(size/2, size/2, size/2, math.Pi, math.Pi*1.5)
			surface.Stroke()
		}
		if mode == 1 {
			surface.MoveTo(0, -size/2)
			surface.LineTo(0, size/2)
			surface.Stroke()
			surface.FillCircle(-size/2, 0, 2.5)
			surface.FillCircle(size/2, 0, 2.5)
		}
		if mode == 2 {
			surface.MoveTo(0, -size/2)
			surface.LineTo(0, size/2)
			surface.MoveTo(-size/2, 0)
			surface.LineTo(size/2, 0)
			surface.Stroke()
		}
		if mode == 3 {
			surface.Arc(-size/2, -size/2, size/2, 0, math.Pi/2)
			surface.Stroke()
			surface.FillCircle(0, size/2, 2.5)
			surface.FillCircle(size/2, 0, 2.5)
		}
	}

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		random.Seed(0)
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.25)
		surface.Save()
		surface.Translate(-percent*width, 0)
		surface.Grid(0, 0, width*2, height, size, size)
		surface.SetLineWidth(5)
		surface.SetLineCap(cairo.LineCapRound)
		for x := 0.0; x < width; x += size {
			for y := 0.0; y < height; y += size {
				mode := random.IntRange(0, 3)
				r := math.Floor(random.FloatRange(0, 4)) * math.Pi / 2
				surface.Save()
				surface.Translate(x+size/2, y+size/2)
				surface.Rotate(r)
				drawCell(mode, surface)
				surface.Restore()
				surface.Save()
				surface.Translate(x+size/2+width, y+size/2)
				surface.Rotate(r)
				drawCell(mode, surface)
				surface.Restore()
			}
		}
		surface.Restore()
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
