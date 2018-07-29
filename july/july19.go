package july

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/geom"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

// July19 ...
func July19() {
	const (
		outFileName   = "out/july19.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 450.0
		height        = 300.0
		res           = 1.0
		scale         = 0.01
	)
	points := make([]*geom.Point, 20)
	points2 := make([]*geom.Point, 20)
	index := 0
	random.Seed(0)
	for i := 0.0; i < 10; i++ {
		points[index] = geom.NewPoint(i*50, 100)
		points[index+1] = geom.NewPoint(i*50, 200)
		points2[index] = geom.NewPoint(points[index].X+random.FloatRange(-100, 100), points[index].Y+random.FloatRange(-100, 100))
		points2[index+1] = geom.NewPoint(points[index+1].X+random.FloatRange(-100, 100), points[index+1].Y+random.FloatRange(-100, 100))
		index += 2
	}

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		random.Seed(0)
		surface.SetLineWidth(0.5)
		surface.ClearRGB(1, 1, 1)
		t := blmath.LerpSin(percent, 0, 1)

		for i := 0; i < 18; i++ {
			for j := 0; j < 1000; j++ {
				pA := geom.LerpPoint(t, points[i], points2[i])
				pB := geom.LerpPoint(t, points[i+1], points2[i+1])
				pC := geom.LerpPoint(t, points[i+2], points2[i+2])
				p := geom.RandomPointInTriangle(&pA, &pB, &pC)
				surface.FillCircle(p.X, p.Y, 0.5)
			}
		}
		// surface.Points(points, 5)
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
