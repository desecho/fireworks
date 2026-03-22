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
	Crossette
	Kamuro
	DoubleRing
	Palm
)

const numPatterns = 9

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
			VX:     (rng.Float64() - 0.5) * 0.6,
			VY:     -(1.8 + rng.Float64()*1.2),
			Life:   1.0,
			Decay:  0,
			Char:   '|',
			Color:  tcell.NewRGBColor(255, 220, 140),
			Active: true,
		},
		FuseLife: float64(h/3) + rng.Float64()*float64(h/4),
		Palette:  palettes[rng.Intn(len(palettes))],
		Pattern:  ExplosionPattern(rng.Intn(numPatterns)),
	}
}

// Explode spawns particles based on the rocket's explosion pattern.
func (r *Rocket) Explode(rng *rand.Rand) []Particle {
	switch r.Pattern {
	case Ring:
		return r.explodeRing(rng, 1)
	case Willow:
		return r.explodeWillow(rng)
	case Chrysanthemum:
		return r.explodeChrysanthemum(rng)
	case Peony:
		return r.explodePeony(rng)
	case Crossette:
		return r.explodeCrossette(rng)
	case Kamuro:
		return r.explodeKamuro(rng)
	case DoubleRing:
		return r.explodeRing(rng, 2)
	case Palm:
		return r.explodePalm(rng)
	default: // Burst
		return r.explodeBurst(rng)
	}
}

// --- Explosion implementations ---

func (r *Rocket) explodeBurst(rng *rand.Rand) []Particle {
	count := 100 + rng.Intn(40)
	particles := make([]Particle, 0, count)
	for i := 0; i < count; i++ {
		angle := rng.Float64() * 2 * math.Pi
		speed := 0.5 + rng.Float64()*2.5
		color := r.Palette[rng.Intn(len(r.Palette))]
		particles = append(particles, Particle{
			X: r.X, Y: r.Y,
			VX: math.Cos(angle) * speed, VY: math.Sin(angle) * speed,
			Life: 0.8 + rng.Float64()*0.2, Decay: 0.012 + rng.Float64()*0.008,
			Char: '@', Color: color, Active: true,
			Shimmer: rng.Float64() < 0.3,
		})
	}
	return particles
}

func (r *Rocket) explodeRing(rng *rand.Rand, rings int) []Particle {
	particles := make([]Particle, 0, 80*rings)
	for ring := 0; ring < rings; ring++ {
		count := 50 + rng.Intn(30)
		speed := 1.2 + float64(ring)*1.0 + rng.Float64()*0.5
		for i := 0; i < count; i++ {
			angle := (float64(i) / float64(count)) * 2 * math.Pi
			jitter := (rng.Float64() - 0.5) * 0.15
			color := r.Palette[rng.Intn(len(r.Palette))]
			particles = append(particles, Particle{
				X: r.X, Y: r.Y,
				VX: math.Cos(angle)*(speed+jitter) * 1.8,
				VY: math.Sin(angle) * (speed + jitter),
				Life: 0.7 + rng.Float64()*0.3, Decay: 0.010 + rng.Float64()*0.005,
				Char: '*', Color: color, Active: true,
			})
		}
	}
	return particles
}

func (r *Rocket) explodeWillow(rng *rand.Rand) []Particle {
	count := 80 + rng.Intn(30)
	particles := make([]Particle, 0, count)
	for i := 0; i < count; i++ {
		angle := rng.Float64() * 2 * math.Pi
		speed := 1.0 + rng.Float64()*2.5
		color := r.Palette[rng.Intn(len(r.Palette))]
		particles = append(particles, Particle{
			X: r.X, Y: r.Y,
			VX: math.Cos(angle) * speed, VY: math.Sin(angle) * speed,
			Life: 1.0, Decay: 0.004 + rng.Float64()*0.003,
			Char: '*', Color: color, Active: true,
			GravMult: 2.0, // heavy droop for willow effect
		})
	}
	return particles
}

func (r *Rocket) explodeChrysanthemum(rng *rand.Rand) []Particle {
	count := 120 + rng.Intn(40)
	particles := make([]Particle, 0, count)
	for i := 0; i < count; i++ {
		angle := rng.Float64() * 2 * math.Pi
		speed := 1.5 + rng.Float64()*3.0
		color := r.Palette[rng.Intn(len(r.Palette))]
		particles = append(particles, Particle{
			X: r.X, Y: r.Y,
			VX: math.Cos(angle) * speed * 1.5,
			VY: math.Sin(angle) * speed,
			Life: 0.9 + rng.Float64()*0.1, Decay: 0.005 + rng.Float64()*0.003,
			Char: '@', Color: color, Active: true,
			Shimmer: true,
		})
	}
	return particles
}

func (r *Rocket) explodePeony(rng *rand.Rand) []Particle {
	particles := r.explodeBurst(rng)
	// Add bright white/gold glitter core
	glitterCount := 30 + rng.Intn(20)
	for i := 0; i < glitterCount; i++ {
		angle := rng.Float64() * 2 * math.Pi
		speed := 0.2 + rng.Float64()*1.2
		particles = append(particles, Particle{
			X: r.X, Y: r.Y,
			VX: math.Cos(angle) * speed, VY: math.Sin(angle) * speed,
			Life: 1.0, Decay: 0.006 + rng.Float64()*0.004,
			Char: '@', Active: true, Shimmer: true,
			Color: tcell.NewRGBColor(255, 255, 240),
		})
	}
	return particles
}

func (r *Rocket) explodeCrossette(rng *rand.Rand) []Particle {
	// 4-6 main arms that each burst into mini-explosions (simulated with clusters)
	arms := 4 + rng.Intn(3)
	particles := make([]Particle, 0, arms*25)
	for a := 0; a < arms; a++ {
		angle := (float64(a) / float64(arms)) * 2 * math.Pi
		cx := r.X + math.Cos(angle)*3
		cy := r.Y + math.Sin(angle)*2
		subCount := 20 + rng.Intn(10)
		color := r.Palette[a%len(r.Palette)]
		for i := 0; i < subCount; i++ {
			sa := rng.Float64() * 2 * math.Pi
			speed := 0.5 + rng.Float64()*1.5
			particles = append(particles, Particle{
				X: cx, Y: cy,
				VX: math.Cos(angle)*2.0 + math.Cos(sa)*speed,
				VY: math.Sin(angle)*1.5 + math.Sin(sa)*speed,
				Life: 0.7 + rng.Float64()*0.3, Decay: 0.012 + rng.Float64()*0.006,
				Char: '*', Color: color, Active: true,
				Shimmer: rng.Float64() < 0.4,
			})
		}
	}
	return particles
}

func (r *Rocket) explodeKamuro(rng *rand.Rand) []Particle {
	// Dense golden cascade that droops heavily
	count := 140 + rng.Intn(40)
	particles := make([]Particle, 0, count)
	gold := []tcell.Color{
		tcell.NewRGBColor(255, 230, 100),
		tcell.NewRGBColor(255, 200, 60),
		tcell.NewRGBColor(255, 180, 0),
		tcell.NewRGBColor(255, 255, 180),
	}
	for i := 0; i < count; i++ {
		angle := rng.Float64() * 2 * math.Pi
		speed := 1.0 + rng.Float64()*2.0
		particles = append(particles, Particle{
			X: r.X, Y: r.Y,
			VX: math.Cos(angle) * speed * 1.2,
			VY: math.Sin(angle) * speed,
			Life: 1.0, Decay: 0.003 + rng.Float64()*0.003,
			Char: '*', Color: gold[rng.Intn(len(gold))], Active: true,
			GravMult: 1.8,
			Shimmer:  true,
		})
	}
	return particles
}

func (r *Rocket) explodePalm(rng *rand.Rand) []Particle {
	// Thick rising tendrils that arc outward
	arms := 6 + rng.Intn(4)
	particles := make([]Particle, 0, arms*20)
	for a := 0; a < arms; a++ {
		baseAngle := (float64(a) / float64(arms)) * 2 * math.Pi
		color := r.Palette[a%len(r.Palette)]
		count := 15 + rng.Intn(10)
		for i := 0; i < count; i++ {
			t := float64(i) / float64(count)
			speed := 1.5 + t*2.5
			jitter := (rng.Float64() - 0.5) * 0.3
			particles = append(particles, Particle{
				X: r.X, Y: r.Y,
				VX: math.Cos(baseAngle+jitter) * speed * 1.5,
				VY: math.Sin(baseAngle+jitter)*speed - 0.5, // bias upward
				Life: 0.7 + t*0.3, Decay: 0.006 + rng.Float64()*0.004,
				Char: '*', Color: color, Active: true,
				GravMult: 1.5,
			})
		}
	}
	return particles
}
