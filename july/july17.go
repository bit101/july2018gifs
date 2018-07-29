package july

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/geom"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

// July17 ...
func July17() {
	const (
		outFileName   = "out/july17.gif"
		timeInSeconds = 4
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		res           = 1.0
		scale         = 0.01
	)

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		random.Seed(0)
		surface.SetLineWidth(0.5)
		surface.ClearRGB(1, 1, 1)
		p0 := geom.NewPoint(blmath.LerpSin(percent, 50, 100), blmath.LerpSin(percent, 50, 350))
		p1 := geom.NewPoint(blmath.LerpSin(percent, 50, 350), blmath.LerpSin(percent, 350, 50))
		p2 := geom.NewPoint(blmath.LerpSin(percent, 350, 50), blmath.LerpSin(percent, 300, 100))

		surface.StrokeCircle(p0.X, p0.Y, 10)
		surface.StrokeCircle(p1.X, p1.Y, 10)
		surface.StrokeCircle(p2.X, p2.Y, 10)
		for i := 0; i < 500; i++ {
			p := geom.RandomPointInTriangle(p0, p1, p2)
			surface.FillCircle(p.X, p.Y, 2)
			surface.MoveTo(p.X, p.Y+random.FloatRange(5, 100))
			surface.LineTo(p.X, p.Y-random.FloatRange(5, 100))
			surface.Stroke()

		}

		surface.Save()
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
