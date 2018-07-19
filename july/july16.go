package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func July16() {
	const (
		outFileName   = "out/july16.gif"
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
		random.Seed(0)
		surface.ClearRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)
		surface.SetSourceRGBA(0, 0, 0, blmath.LerpSin(percent, 1, 0.5))

		for x := -200.0; x < 200.0; x += 50.0 {
			for y := -200.0; y < 200.0; y += 50.0 {
				surface.Translate(0, blmath.LerpSin(percent, 0, random.FloatRange(-50, 50)))
				surface.Save()
				surface.Scale(1, blmath.LerpSin(percent, 1, 0.01))
				surface.Rotate(blmath.LerpSin(percent, 0, random.FloatRange(-math.Pi/4, math.Pi/4)))
				surface.FillRectangle(x, y, 40, 40)
				surface.Restore()
			}
		}
		surface.Restore()
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
