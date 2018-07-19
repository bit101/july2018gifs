package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func July12() {
	const (
		outFileName   = "out/july12.gif"
		timeInSeconds = 6
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		scale         = 0.05
	)

	type Circle struct {
		x      float64
		y      float64
		radius float64
	}

	var circles []*Circle

	checkCircle := func(c *Circle) bool {
		if c.x+c.radius > width/2 ||
			c.x-c.radius < -height/2 ||
			c.y+c.radius > height/2 ||
			c.y-c.radius < -height/2 {
			return true
		}
		for _, c2 := range circles {
			if math.Hypot(c.x-c2.x, c.y-c2.y) < c.radius+c2.radius {
				return true
			}
		}
		return false
	}

	makeCircle := func(surface *blgo.Surface, percent float64) {
		var c *Circle
		for {
			amt := blmath.LerpSin(percent, 0, 200)
			x := random.FloatRange(-width/2-amt, width/2+amt)
			y := random.FloatRange(-height/2-amt, height/2+amt)
			r := 1.0
			c = &Circle{x, y, r}
			if !checkCircle(c) {
				break
			}
		}
		for {
			if checkCircle(c) {
				break
			}
			c.radius = c.radius + 1
		}
		circles = append(circles, c)
		surface.FillCircle(c.x, c.y, c.radius-1)
	}
	render := func(surface *blgo.Surface, percent float64) {
		surface.SetLineWidth(0.5)
		surface.SetSourceRGB(1, 1, 1)
		circles = make([]*Circle, 0)
		surface.ClearRGB(0, 0, 0)
		surface.Save()
		surface.Translate(width/2, height/2)
		random.Seed(2)
		for i := 0.0; i < 1000; i++ {
			makeCircle(surface, percent)
		}
		surface.Restore()
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
