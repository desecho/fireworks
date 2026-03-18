package main

import "github.com/gdamore/tcell/v2"

// Render draws the current world state to the screen.
func Render(s tcell.Screen, w *World) {
	s.Clear()

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack)

	// Draw trails first (behind everything)
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

	// Draw particles
	for i := range w.Particles {
		p := &w.Particles[i]
		if !p.Active {
			continue
		}
		ix, iy := int(p.X), int(p.Y)
		if ix >= 0 && ix < w.Width && iy >= 0 && iy < w.Height {
			color := fadeColor(p.Color, p.Life)
			style := defStyle.Foreground(color)
			s.SetContent(ix, iy, p.Char, nil, style)
		}
	}

	// Draw rockets
	for _, r := range w.Rockets {
		ix, iy := int(r.X), int(r.Y)
		if ix >= 0 && ix < w.Width && iy >= 0 && iy < w.Height {
			style := defStyle.Foreground(r.Color)
			s.SetContent(ix, iy, r.Char, nil, style)
		}
	}

	s.Show()
}
