package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/color"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

// July11 ...
func July11() {
	const (
		outFileName   = "out/july11.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		scale         = 0.05
	)

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.25)
		surface.Grid(0, 0, width, height, 20, 20)
		random.Seed(1)
		for j := 0; j < 30; j++ {
			x := random.FloatRange(0, width)
			y := random.FloatRange(0, height)
			z := random.FloatRange(0, 5)
			c := color.HSV(random.FloatRange(0, 60), 0.25, 1)
			n := blmath.LerpSin(percent+random.Float(), 20, 300)
			for i := 0.0; i < n; i++ {
				r := blmath.LerpSin(i/n+math.Pi/4, 0, 10)
				surface.Circle(x, y, r)
				surface.SetSourceColor(c)
				surface.FillPreserve()
				surface.SetSourceRGB(0, 0, 0)
				surface.Stroke()
				a := blmath.Map(noise.Perlin(x*scale, y*scale, z), 0, 1, 0, math.Pi*2)
				x += math.Cos(a) * 2
				y += math.Sin(a) * 2
			}
		}
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
