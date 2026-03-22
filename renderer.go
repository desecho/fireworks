package main

import "github.com/gdamore/tcell/v2"

// Render draws the current world state to the screen.
func Render(s tcell.Screen, w *World) {
	s.Clear()

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack)

	// Draw trails first (behind everything) — dim
	for i := range w.Trails {
		t := &w.Trails[i]
		if !t.Active {
			continue
		}
		ix, iy := int(t.X), int(t.Y)
		if ix >= 0 && ix < w.Width && iy >= 0 && iy < w.Height {
			color := fadeColor(t.Color, t.Life)
			style := defStyle.Foreground(color)
			s.SetContent(ix, iy, t.Char, nil, style)
		}
	}

	// Draw particles — brighter, on top of trails
	for i := range w.Particles {
		p := &w.Particles[i]
		if !p.Active {
			continue
		}
		ix, iy := int(p.X), int(p.Y)
		if ix >= 0 && ix < w.Width && iy >= 0 && iy < w.Height {
			color := fadeColor(p.Color, p.Life)
			style := defStyle.Foreground(color).Bold(p.Life > 0.6)
			s.SetContent(ix, iy, p.Char, nil, style)
		}
	}

	// Draw rockets — bright and bold
	for _, r := range w.Rockets {
		ix, iy := int(r.X), int(r.Y)
		if ix >= 0 && ix < w.Width && iy >= 0 && iy < w.Height {
			style := defStyle.Foreground(r.Color).Bold(true)
			s.SetContent(ix, iy, r.Char, nil, style)
			// Draw a spark above the rocket
			if iy-1 >= 0 {
				sparkStyle := defStyle.Foreground(tcell.NewRGBColor(255, 255, 200)).Bold(true)
				s.SetContent(ix, iy-1, '^', nil, sparkStyle)
			}
		}
	}

	s.Show()
}
