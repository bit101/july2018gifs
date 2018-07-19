package july

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/util"
)

func July15() {
	const (
		outFileName   = "out/july15.gif"
		timeInSeconds = 4
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		res           = 1.0
		scale         = 0.01
	)

	render := func(surface *blgo.Surface, percent float64) {
		for x := 0.0; x < width; x += res {
			for y := 0.0; y < height; y += res {
				v := blmath.Map(noise.PerlinOct(x*scale, y*scale, 0, 3, blmath.LerpSin(percent, 0.0, 1)), -0.5, 0.5, 0, 1)
				surface.SetSourceRGB(v, v, v)
				surface.FillRectangle(x, y, 1, 1)
			}
		}
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
