package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func July09() {
	const (
		outFileName   = "out/july09.gif"
		timeInSeconds = 10
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
	)

	type Circle struct {
		x      float64
		y      float64
		radius float64
	}

	var circles []*Circle

	checkCircle := func(c *Circle) bool {
		dist := math.Hypot(c.x, c.y)
		if dist+c.radius > width/2 {
			return true
		}
		for _, c2 := range circles {
			if math.Hypot(c.x-c2.x, c.y-c2.y) < c.radius+c2.radius {
				return true
			}
		}
		return false
	}

	makeCircle := func(surface *blgo.Surface, amt float64) {
		var c *Circle
		for {
			x := random.FloatRange(-width/2, width/2)
			y := random.FloatRange(-height/2, height/2)
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
		surface.FillCircle(c.x, c.y, c.radius*amt)
	}
	render := func(surface *blgo.Surface, percent float64) {
		amt := blmath.LerpSin(percent, 0, 1)
		surface.SetLineWidth(0.5)
		surface.SetSourceRGB(1, 1, 1)
		circles = make([]*Circle, 0)
		surface.ClearRGB(0, 0, 0)
		surface.Save()
		surface.Translate(width/2, height/2)
		random.Seed(1)
		for i := 0.0; i < 1000; i++ {
			makeCircle(surface, amt)
		}
		surface.Restore()
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
