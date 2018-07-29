package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/geom"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

// July25 ...
func July25() {
	const (
		outFileName   = "out/july25.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 300.0
	)

	size := 20
	randSize := 20.0
	points0 := make([]*geom.Point, size)
	points1 := make([]*geom.Point, size)
	for i := range points0 {
		a := float64(i) / float64(len(points0)) * math.Pi * 2
		r := 100.0
		points0[i] = geom.NewPoint(
			math.Cos(a)*r+random.FloatRange(-randSize, randSize),
			height/2+math.Sin(a)*r+random.FloatRange(-randSize, randSize),
		)
		points1[i] = geom.NewPoint(
			width/2+math.Cos(a)*r/2+random.FloatRange(-randSize, randSize),
			height/2+math.Sin(a)*r/2+random.FloatRange(-randSize, randSize),
		)
	}

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		random.Seed(1)
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.25)
		for x := 0.0; x < width; x++ {
			p3 := make([]*geom.Point, size)
			for i := range points0 {
				p := geom.LerpPoint(
					blmath.LerpSin(percent+x/width+random.FloatRange(-0.1, 0.1), 0, 1),
					points0[i],
					points1[i],
				)
				r := 5.0
				p.X += random.FloatRange(-r, r)
				p.Y += random.FloatRange(-r, r)
				p3[i] = &p
			}
			surface.Save()
			surface.Translate(x, 0)
			surface.MultiLoop(p3)
			surface.Restore()
			surface.Stroke()
		}
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
