package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func July13() {
	const (
		outFileName   = "out/july13.gif"
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

	makeCircle := func(surface *blgo.Surface) {
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
		surface.Save()
		surface.Translate(c.x, c.y)
		surface.FillCircle(0, 0, c.radius)
		surface.Restore()
	}
	render := func(surface *blgo.Surface, percent float64) {
		amt := blmath.LerpSin(percent, 0, 1)
		angle := percent * math.Pi * 2.0
		surface.SetLineWidth(0.5)
		surface.SetSourceRGBA(1, 1, 1, blmath.Map(amt, 0, 1, 1, 0.5))
		circles = make([]*Circle, 0)
		surface.ClearRGB(0, 0, 0)
		random.Seed(5)
		for i := 0.0; i < 1000; i++ {
			surface.Save()
			surface.Translate(width/2, height*0.5+height*0.25*amt-i*0.2*amt)
			surface.Scale(1, blmath.Map(amt, 0, 1, 1, 0.25))
			surface.Rotate(angle)
			makeCircle(surface)
			surface.Restore()
		}
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
