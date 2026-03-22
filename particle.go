package main

import "github.com/gdamore/tcell/v2"

// Particle represents a single particle with position, velocity, and life.
type Particle struct {
	X, Y     float64
	VX, VY   float64
	Life     float64 // 1.0 -> 0.0
	Decay    float64
	Char     rune
	Color    tcell.Color
	Active   bool
	Shimmer  bool    // twinkle effect
	GravMult float64 // per-particle gravity multiplier (0 = default 1.0)
}

// Update applies physics to the particle each tick.
func (p *Particle) Update(gravity, drag float64) {
	if !p.Active {
		return
	}
	gm := p.GravMult
	if gm == 0 {
		gm = 1.0
	}
	p.VY += gravity * gm
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
func charForLife(life float64, shimmer bool) rune {
	if shimmer && int(life*100)%6 < 2 {
		return '#'
	}
	switch {
	case life > 0.85:
		return '@'
	case life > 0.65:
		return '*'
	case life > 0.45:
		return 'o'
	case life > 0.25:
		return '+'
	case life > 0.10:
		return ':'
	default:
		return '.'
	}
}
