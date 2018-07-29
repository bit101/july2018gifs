package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/util"
)

// July22 ...
func July22() {
	const (
		outFileName   = "out/july22.gif"
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
		surface.Save()
		surface.Translate(width/2, height/2)

		offset := percent * math.Pi * 2

		for i := 0.0; i < math.Pi*2; i += 0.01 {
			r := math.Sin(i*17-offset*2)*30.0 + 168
			x0 := math.Cos(i) * r
			y0 := math.Sin(i) * r
			surface.LineTo(x0, y0)
		}
		surface.SetSourceRGB(1, 1, 1)
		surface.Fill()

		for i := 0.0; i < math.Pi*2; i += 0.01 {
			r := math.Sin(i*20+offset*3) * 10.0
			r += math.Sin(i*17-offset*2) * 30.0
			r += 100
			x0 := math.Cos(i) * r
			y0 := math.Sin(i) * r
			surface.LineTo(x0, y0)
		}
		surface.SetSourceRGB(0, 0, 0)
		surface.Fill()

		for i := 0.0; i < math.Pi*2; i += 0.01 {
			r := math.Sin(i*20+offset*3)*10.0 + 50
			x0 := math.Cos(i) * r
			y0 := math.Sin(i) * r
			surface.LineTo(x0, y0)
		}
		surface.SetSourceRGB(1, 1, 1)
		surface.Fill()

		surface.Restore()
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
