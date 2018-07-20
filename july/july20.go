package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/util"
)

func July20() {
	const (
		outFileName   = "out/july20.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
	)
	render := func(surface *blgo.Surface, percent float64) {
		surface.SetLineWidth(0.5)
		surface.ClearRGB(1, 1, 1)

		// offset := blmath.LerpSin(percent, 0, math.Pi*2)
		offset := percent * math.Pi * 2
		for i := 0.0; i < 400.0; i++ {
			y0 := math.Sin(i*0.05+offset)*60.0 + 70
			surface.LineTo(i, y0)
		}
		surface.Stroke()
		for i := 0.0; i < 400.0; i++ {
			y0 := math.Sin(i*0.13-offset*10)*30.0 + 350
			surface.LineTo(i, y0)
		}
		surface.Stroke()
		surface.Stroke()
		for i := 0.0; i < 400.0; i++ {
			y0 := math.Sin(i*0.05+offset) * 60.0
			y0 += math.Sin(i*0.13-offset*10) * 30.0
			y0 += 200
			surface.LineTo(i, y0)
		}
		surface.Stroke()
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}