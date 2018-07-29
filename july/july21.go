package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

// July21 ...
func July21() {
	const (
		outFileName   = "out/july21.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
	)

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		surface.SetLineWidth(0.5)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)
		amt1 := blmath.LerpSin(percent, 0, 1)
		amt2 := blmath.LerpSin(percent+0.17, 0, 1)
		offset := percent * math.Pi * 2

		// offset := blmath.LerpSin(percent, 0, math.Pi*2)
		for i := 0.0; i < 400.0; i++ {
			y0 := math.Sin(i*0.2*amt1+offset)*60.0*amt1 + 70
			surface.LineTo(i, y0)
		}
		surface.Stroke()
		for i := 0.0; i < 400.0; i++ {
			y0 := math.Sin(i*0.13+offset*20)*60.0*amt2 + 330
			surface.LineTo(i, y0)
		}
		surface.Stroke()
		surface.Stroke()
		for i := 0.0; i < 400.0; i++ {
			y0 := math.Sin(i*0.2*amt1+offset) * 30.0 * amt1
			y0 += math.Sin(i*0.13+offset*20) * 30.0 * amt2
			y0 += 200
			surface.LineTo(i, y0)
		}
		surface.Stroke()
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
