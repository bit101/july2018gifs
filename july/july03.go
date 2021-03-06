package july

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/color"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/util"
)

// July03 ...
func July03() {
	const (
		timeInSeconds = 5
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		outFileName   = "out/july03.gif"
		width         = 400.0
		height        = 300.0
		count         = 20.0
		size          = 20.0
		scale         = 0.015
		h             = 70.0
	)

	type Point struct {
		x float64
		y float64
		z float64
		c color.Color
	}

	points := make([]Point, int(count*count))
	i := 0
	for x := 0.0; x < count; x += 1.0 {
		for y := 0.0; y < count; y += 1.0 {
			points[i].x = width/2 + (x-y)/2*size
			points[i].y = 40 + (x+y)/4*size
			points[i].c = color.RandomGreyRange(0.75, 1.0)
			i++
		}
	}

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		amt := blmath.LerpSin(percent, 0, 1)
		surface.ClearRGB(0.25, 0.25, 0.25)
		surface.SetLineWidth(0.25)
		for x := 0; x < int(count-1); x++ {
			for y := 0; y < int(count-1); y++ {
				i := y*int(count) + x
				p0 := points[i]
				p1 := points[i+1]
				p2 := points[i+1+int(count)]
				p3 := points[i+int(count)]
				// do you even optimize your code, bro?
				surface.MoveTo(p0.x, p0.y-noise.Perlin(p0.x*scale, p0.y*scale, amt)*h)
				surface.LineTo(p1.x, p1.y-noise.Perlin(p1.x*scale, p1.y*scale, amt)*h)
				surface.LineTo(p2.x, p2.y-noise.Perlin(p2.x*scale, p2.y*scale, amt)*h)
				surface.LineTo(p3.x, p3.y-noise.Perlin(p3.x*scale, p3.y*scale, amt)*h)
				surface.LineTo(p0.x, p0.y-noise.Perlin(p0.x*scale, p0.y*scale, amt)*h)
				g := color.Grey(blmath.Map(noise.Perlin(p0.x*scale, p0.y*scale, amt), -0.5, 0.5, 0, 1))
				surface.SetSourceColor(g)
				surface.FillPreserve()
				surface.SetSourceRGB(0, 0, 0)
				surface.Stroke()
			}
		}
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
