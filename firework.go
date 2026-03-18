package main

import (
	"math"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

// ExplosionPattern determines how particles are spawned on explosion.
type ExplosionPattern int

const (
	Burst ExplosionPattern = iota
	Ring
	Willow
	Chrysanthemum
	Peony
)

// Rocket is a firework that flies up then explodes into particles.
type Rocket struct {
	Particle
	FuseLife float64
	Palette  []tcell.Color
	Pattern  ExplosionPattern
}

// NewRocket creates a rocket launching from the bottom of the screen.
func NewRocket(w, h int, rng *rand.Rand) *Rocket {
	x := float64(w/4) + rng.Float64()*float64(w/2)
	return &Rocket{
		Particle: Particle{
			X:      x,
			Y:      float64(h - 1),
			VX:     (rng.Float64() - 0.5) * 0.8,
			VY:     -(1.5 + rng.Float64()*1.0),
			Life:   1.0,
			Decay:  0,
			Char:   '|',
			Color:  tcell.NewRGBColor(255, 200, 100),
			Active: true,
		},
		FuseLife: float64(h/3) + rng.Float64()*float64(h/4),
		Palette:  palettes[rng.Intn(len(palettes))],
		Pattern:  ExplosionPattern(rng.Intn(5)),
	}
}

// Explode spawns particles based on the rocket's explosion pattern.
func (r *Rocket) Explode(rng *rand.Rand) []Particle {
	switch r.Pattern {
	case Ring:
		return r.explodeRing(rng)
	case Willow:
		return r.explodeBurst(rng, 70, 1.0, 3.0, 0.008, 1.5)
	case Chrysanthemum:
		return r.explodeBurst(rng, 90, 1.5, 3.5, 0.005, 1.0)
	case Peony:
		return r.explodePeony(rng)
	default: // Burst
		return r.explodeBurst(rng, 80, 0.5, 2.5, 0.015, 1.0)
	}
}

func (r *Rocket) explodeBurst(rng *rand.Rand, count int, minSpd, maxSpd, decay, gravMult float64) []Particle {
	particles := make([]Particle, 0, count)
	for i := 0; i < count; i++ {
		angle := rng.Float64() * 2 * math.Pi
		speed := minSpd + rng.Float64()*(maxSpd-minSpd)
		color := r.Palette[rng.Intn(len(r.Palette))]
		p := Particle{
			X:      r.X,
			Y:      r.Y,
			VX:     math.Cos(angle) * speed,
			VY:     math.Sin(angle) * speed,
			Life:   0.8 + rng.Float64()*0.2,
			Decay:  decay + rng.Float64()*0.008,
			Char:   '✦',
			Color:  color,
			Active: true,
		}
		_ = gravMult // stored in pattern, used by engine
		particles = append(particles, p)
	}
	return particles
}

func (r *Rocket) explodeRing(rng *rand.Rand) []Particle {
	count := 40 + rng.Intn(20)
	particles := make([]Particle, 0, count)
	speed := 1.5 + rng.Float64()*1.0
	for i := 0; i < count; i++ {
		angle := (float64(i) / float64(count)) * 2 * math.Pi
		color := r.Palette[rng.Intn(len(r.Palette))]
		p := Particle{
			X:      r.X,
			Y:      r.Y,
			VX:     math.Cos(angle) * speed,
			VY:     math.Sin(angle) * speed,
			Life:   0.8 + rng.Float64()*0.2,
			Decay:  0.012 + rng.Float64()*0.005,
			Char:   '✦',
			Color:  color,
			Active: true,
		}
		particles = append(particles, p)
	}
	return particles
}

func (r *Rocket) explodePeony(rng *rand.Rand) []Particle {
	// Dense burst + glitter
	particles := r.explodeBurst(rng, 90, 0.5, 2.5, 0.015, 1.0)
	// Add glitter particles
	glitterCount := 20 + rng.Intn(15)
	for i := 0; i < glitterCount; i++ {
		angle := rng.Float64() * 2 * math.Pi
		speed := 0.3 + rng.Float64()*1.0
		p := Particle{
			X:      r.X,
			Y:      r.Y,
			VX:     math.Cos(angle) * speed,
			VY:     math.Sin(angle) * speed,
			Life:   1.0,
			Decay:  0.008 + rng.Float64()*0.005,
			Char:   '✦',
			Color:  tcell.NewRGBColor(220, 220, 230),
			Active: true,
		}
		particles = append(particles, p)
	}
	return particles
}
