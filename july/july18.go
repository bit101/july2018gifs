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

func July18() {
	const (
		outFileName   = "out/july18.gif"
		timeInSeconds = 10
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		res           = 1.0
		scale         = 0.01
	)

	render := func(surface *blgo.Surface, percent float64) {
		random.Seed(0)
		surface.SetLineWidth(0.5)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)
		a := percent * math.Pi * 2
		p0 := geom.NewPoint(math.Cos(a+math.Pi*2)*190, math.Sin(a+math.Pi*2)*100)
		p1 := geom.NewPoint(math.Cos(a+math.Pi*2/3)*190, math.Sin(a+math.Pi*2/3)*100)
		p2 := geom.NewPoint(math.Cos(a+math.Pi*4/3)*190, math.Sin(a+math.Pi*4/3)*100)

		surface.StrokeEllipse(p0.X, p0.Y, 10, 5)
		surface.StrokeEllipse(p1.X, p1.Y, 10, 5)
		surface.StrokeEllipse(p2.X, p2.Y, 10, 5)
		for i := 0; i < 500; i++ {
			p := geom.RandomPointInTriangle(p0, p1, p2)
			amt := blmath.LerpSin(percent, 0, 1)
			surface.StrokeEllipse(p.X, p.Y, 2*(1.1-amt), (1.1 - amt))
			surface.MoveTo(p.X, p.Y+random.FloatRange(5, 100)*amt)
			surface.LineTo(p.X, p.Y-random.FloatRange(5, 100)*amt)
			surface.Stroke()

		}

		surface.Restore()
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
