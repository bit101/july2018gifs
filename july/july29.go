package july

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
	cairo "github.com/bit101/go-cairo"
)

// July29 ...
func July29() {
	const (
		outFileName   = "out/july29.gif"
		timeInSeconds = 10
		fps           = 30
		frames        = timeInSeconds * fps
		framesDir     = "frames"
		width         = 400.0
		height        = 400.0
		size          = 10.0
		w             = 3.0
	)

	drawCell := func(mode int, surface *blgo.Surface) {
		if mode == 0 {
			surface.Arc(-size/2, -size/2, size/2, 0, math.Pi/2)
			surface.Stroke()
			surface.Arc(size/2, size/2, size/2, math.Pi, math.Pi*1.5)
			surface.Stroke()
		}
		if mode == 1 {
			surface.MoveTo(0, -size/2)
			surface.LineTo(0, size/2)
			surface.Stroke()
			surface.FillCircle(-size/2, 0, w/2)
			surface.FillCircle(size/2, 0, w/2)
		}
		if mode == 2 {
			surface.MoveTo(0, -size/2)
			surface.LineTo(0, size/2)
			surface.MoveTo(-size/2, 0)
			surface.LineTo(size/2, 0)
			surface.Stroke()
		}
		if mode == 3 {
			surface.Arc(-size/2, -size/2, size/2, 0, math.Pi/2)
			surface.Stroke()
			surface.FillCircle(0, size/2, w/2)
			surface.FillCircle(size/2, 0, w/2)
		}
	}

	render := func(surface *blgo.Surface, percent float64) {
		random.Seed(0)
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.25)
		// surface.Grid(0, 0, width, height, size, size)
		surface.SetLineWidth(w)
		surface.SetLineCap(cairo.LINE_CAP_ROUND)
		r0 := 0.0
		drawMode := "fill"
		if percent < 0.5 {
			r0 = percent * 2 * 300
		} else {
			drawMode = "erase"
			r0 = (percent - 0.5) * 2 * 300
		}
		for x := 0.0; x < width; x += size {
			for y := 0.0; y < height; y += size {
				dfc := math.Hypot(x-width/2+size/2, y-height/2+size/2)
				mode := random.IntRange(0, 3)
				r := math.Floor(random.FloatRange(0, 4)) * math.Pi / 2
				if (drawMode == "fill" && dfc > r0) || (drawMode == "erase" && dfc < r0) {
					mode = 1
				}
				surface.Save()
				surface.Translate(x+size/2, y+size/2)
				surface.Rotate(r)
				drawCell(mode, surface)
				surface.Restore()
				surface.Save()
				surface.Translate(x+size/2+width, y+size/2)
				surface.Rotate(r)
				drawCell(mode, surface)
				surface.Restore()
			}
		}
	}

	animation := anim.NewAnimation(width, height, frames)
	animation.Render(framesDir, "frame", render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}
