package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func July14() {
	const (
		outFileName   = "out/july14.gif"
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

	checkCircle := func(c *Circle) (bool, *Circle) {
		dist := math.Hypot(c.x, c.y)
		if dist+c.radius > width/2 {
			return true, nil
		}
		for _, c2 := range circles {
			if math.Hypot(c.x-c2.x, c.y-c2.y) < c.radius+c2.radius {
				return true, c2
			}
		}
		return false, nil
	}

	makeCircle := func(surface *blgo.Surface, amt float64) {
		var c *Circle
		var c3 *Circle
		for {
			x := random.FloatRange(-width/2, width/2)
			y := random.FloatRange(-height/2, height/2)
			r := 4.0
			c = &Circle{x, y, r}
			hit, _ := checkCircle(c)
			if !hit {
				break
			}
		}
		for {
			hit, c2 := checkCircle(c)
			if hit {
				c3 = c2
				break
			}
			c.radius = c.radius + 1
		}
		circles = append(circles, c)
		surface.StrokeCircle(c.x, c.y, c.radius-1)

		if c3 != nil {
			surface.MoveTo(c3.x, c3.y)
			surface.LineTo(c3.x+(c.x-c3.x)*(1.0-amt), c3.y+(c.y-c3.y)*(1.0-amt))
			surface.Stroke()
		}
	}
	render := func(surface *blgo.Surface, percent float64) {
		surface.SetSourceRGB(1, 1, 0)
		surface.SetLineWidth(0.5)
		amt := blmath.LerpSin(percent, 0, 1)
		circles = make([]*Circle, 0)
		surface.ClearRGB(0, 0, 0)
		random.Seed(9)
		for i := 0.0; i < 200; i++ {
			surface.Save()
			surface.Translate(width/2, height*0.5)
			makeCircle(surface, amt)
			surface.Restore()
		}
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
