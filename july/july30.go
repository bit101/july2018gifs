package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
	cairo "github.com/bit101/go-cairo"
)

// July30 ...
func July30() {
	const (
		outFileName   = "out/july30.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
	)
	size := 20.0
	w := 3.0

	drawCell := func(mode int, surface *blgo.Surface) {
		if mode == 0 {
			surface.Arc(-size/2, -size/2, size/2+w, 0, math.Pi/2)
			surface.Stroke()
			surface.Arc(-size/2, -size/2, size/2-w, 0, math.Pi/2)
			surface.Stroke()
			surface.Arc(size/2, size/2, size/2+w, math.Pi, math.Pi*1.5)
			surface.Stroke()
			surface.Arc(size/2, size/2, size/2-w, math.Pi, math.Pi*1.5)
			surface.Stroke()
		}
		if mode == 1 {
			surface.MoveTo(-w, -size/2)
			surface.LineTo(-w, size/2)
			surface.MoveTo(w, -size/2)
			surface.LineTo(w, size/2)

			surface.MoveTo(-size/2, -w)
			surface.LineTo(-w, -w)
			surface.MoveTo(-size/2, w)
			surface.LineTo(-w, w)

			surface.MoveTo(size/2, -w)
			surface.LineTo(w, -w)
			surface.MoveTo(size/2, w)
			surface.LineTo(w, w)
			surface.Stroke()
		}
	}

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		w = blmath.LerpSin(percent, 1, 10)
		random.Seed(0)
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.5)
		// surface.Grid(0, 0, width, height, size, size)
		// surface.SetLineWidth(w)
		surface.SetLineCap(cairo.LineCapRound)
		for x := 0.0; x < width; x += size {
			for y := 0.0; y < height; y += size {
				mode := random.IntRange(0, 1)
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
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
