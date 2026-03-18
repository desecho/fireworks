package main

import "github.com/gdamore/tcell/v2"

// Particle represents a single particle with position, velocity, and life.
type Particle struct {
	X, Y   float64
	VX, VY float64
	Life   float64 // 1.0 → 0.0
	Decay  float64
	Char   rune
	Color  tcell.Color
	Active bool
}

// Update applies physics to the particle each tick.
func (p *Particle) Update(gravity, drag float64) {
	if !p.Active {
		return
	}
	p.VY += gravity
	p.VX *= drag
	p.VY *= drag
	p.X += p.VX
	p.Y += p.VY
	p.Life -= p.Decay
	if p.Life <= 0 {
		p.Active = false
	}
}

// charForLife returns the display character based on remaining life.
func charForLife(life float64) rune {
	switch {
	case life > 0.8:
		return '✦'
	case life > 0.5:
		return '*'
	case life > 0.2:
		return '·'
	default:
		return '.'
	}
}
