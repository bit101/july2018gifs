package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

// July31 ...
func July31() {
	const (
		outFileName   = "out/july31.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
	)

	drawCell := func(surface *blgo.Surface, x, y, radius, percent float64) {
		surface.Save()
		surface.Translate(x, y)
		if random.Boolean() {
			surface.Rotate(percent * math.Pi / 3)
		} else {
			surface.Rotate(percent * -math.Pi / 3)
		}
		surface.SetLineWidth(0.5)
		surface.Polygon(0, 0, radius-5, 6, 0)
		surface.SetSourceRGB(1, 1, 1)
		surface.FillPreserve()
		surface.SetSourceRGB(0, 0, 0)
		surface.StrokePreserve()
		surface.Clip()
		surface.SetLineWidth(16)
		mode := random.IntRange(0, 1)
		if mode == 0 {
			surface.Arc(radius, 0, radius/2, math.Pi*2/3, math.Pi*4/3)
			surface.Stroke()
			surface.Rotate(math.Pi * 2 / 3)
			surface.Arc(radius, 0, radius/2, math.Pi*2/3, math.Pi*4/3)
			surface.Stroke()
			surface.Rotate(math.Pi * 2 / 3)
			surface.Arc(radius, 0, radius/2, math.Pi*2/3, math.Pi*4/3)
			surface.Stroke()
		}
		if mode == 1 {

			r := random.IntRange(0, 2)
			if r == 0 {
				surface.Rotate(-math.Pi / 3)
			} else if r == 2 {
				surface.Rotate(math.Pi / 3)
			}
			surface.MoveTo(0, -radius)
			surface.LineTo(0, radius)
			surface.Stroke()

			surface.Arc(radius, 0, radius/2, math.Pi*2/3, math.Pi*4/3)
			surface.Stroke()
			surface.Rotate(math.Pi)
			surface.Arc(radius, 0, radius/2, math.Pi*2/3, math.Pi*4/3)
			surface.Stroke()
		}
		surface.Restore()
	}

	surface := blgo.NewSurface(width, height)
	surface.SetLineWidth(0.5)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		random.Seed(0)
		surface.ClearRGB(0.8, 0.8, 0.8)
		even := true
		spacing := 80.0
		mult := math.Sin(math.Pi / 3)
		radius := spacing / 2 / mult

		for x := 0.0; x < width+spacing; x += spacing * mult {
			even = !even
			for y := 0.0; y < height+spacing; y += spacing {
				yy := y
				if even {
					yy += spacing / 2
				}
				drawCell(surface, x, yy, radius, blmath.LerpSin(percent, 0, 1))
			}
		}
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
