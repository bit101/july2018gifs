package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/color"
	"github.com/bit101/blgo/util"
)

// July07 ...
func July07() {
	const (
		outFileName   = "out/july07.gif"
		timeInSeconds = 4
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		count         = 100.0
		size          = 15.0
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
			points[i].y = (x+y)/4*size - 120
			i++
		}
	}

	dist := func(x0, y0, x1, y1 float64) float64 {
		dx, dy := x1-x0, y1-y0
		return math.Hypot(dx, dy)
	}

	surface := blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(func(percent float64) {
		amt := -percent * math.Pi * 2
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
				h := 2000 / dist(p0.x, p0.y, width/2, height/2)
				surface.MoveTo(p0.x, p0.y+math.Sin(dist(p0.x, p0.y, width/2, height/2)*0.05+amt)*h)
				surface.LineTo(p1.x, p1.y+math.Sin(dist(p1.x, p1.y, width/2, height/2)*0.05+amt)*h)
				surface.LineTo(p2.x, p2.y+math.Sin(dist(p2.x, p2.y, width/2, height/2)*0.05+amt)*h)
				surface.LineTo(p3.x, p3.y+math.Sin(dist(p3.x, p3.y, width/2, height/2)*0.05+amt)*h)
				surface.LineTo(p0.x, p0.y+math.Sin(dist(p0.x, p0.y, width/2, height/2)*0.05+amt)*h)
				s := math.Sin(dist(p0.x, p0.y, width/2, height/2)*0.05 + amt)
				g := blmath.Map(s, -1.0, 1.0, 1.0, 0.25)
				surface.SetSourceColor(color.RGB(g, 0, 1-g))
				surface.FillPreserve()
				surface.SetSourceRGB(0, 0, 0)
				surface.Stroke()
			}
		}
	})
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
