package main

import "math/rand"

const (
	gravity     = 0.04
	drag        = 0.98
	maxParticles = 800
)

// World holds all game state.
type World struct {
	Width, Height int
	Rockets       []*Rocket
	Particles     []Particle
	Trails        []Particle
	RNG           *rand.Rand
	tickCount     int
	spawnInterval int
}

// NewWorld creates a new world with the given dimensions.
func NewWorld(w, h int) *World {
	rng := rand.New(rand.NewSource(rand.Int63()))
	return &World{
		Width:         w,
		Height:        h,
		Rockets:       make([]*Rocket, 0, 16),
		Particles:     make([]Particle, 0, maxParticles),
		Trails:        make([]Particle, 0, maxParticles),
		RNG:           rng,
		spawnInterval: 30 + rng.Intn(30),
	}
}

// Resize updates world dimensions.
func (w *World) Resize(width, height int) {
	w.Width = width
	w.Height = height
}

// Update advances the simulation by one tick.
func (w *World) Update() {
	w.tickCount++

	// Maybe spawn a new rocket
	if w.tickCount >= w.spawnInterval && len(w.Particles) < maxParticles {
		w.Rockets = append(w.Rockets, NewRocket(w.Width, w.Height, w.RNG))
		w.tickCount = 0
		w.spawnInterval = 30 + w.RNG.Intn(30)
	}

	// Update rockets
	aliveRockets := w.Rockets[:0]
	for _, r := range w.Rockets {
		r.Particle.Update(gravity*0.5, 1.0) // rockets have less gravity, no drag
		r.FuseLife--
		if r.FuseLife <= 0 || r.Y < 2 {
			// Explode
			newParticles := r.Explode(w.RNG)
			w.Particles = append(w.Particles, newParticles...)
		} else {
			// Spawn rocket trail
			if len(w.Trails) < maxParticles {
				trail := Particle{
					X: r.X, Y: r.Y,
					Life: 0.4, Decay: 0.06,
					Char: '.',
					Color: fadeColor(r.Color, 0.4),
					Active: true,
				}
				w.Trails = append(w.Trails, trail)
			}
			aliveRockets = append(aliveRockets, r)
		}
	}
	w.Rockets = aliveRockets

	// Update particles
	for i := range w.Particles {
		p := &w.Particles[i]
		if !p.Active {
			continue
		}
		p.Update(gravity, drag)
		p.Char = charForLife(p.Life)

		// Check bounds
		if p.X < 0 || p.X >= float64(w.Width) || p.Y < 0 || p.Y >= float64(w.Height) {
			p.Active = false
			continue
		}

		// Spawn trail
		if len(w.Trails) < maxParticles && p.Life > 0.2 {
			trail := Particle{
				X: p.X, Y: p.Y,
				Life: 0.3, Decay: 0.05,
				Char:  '.',
				Color: fadeColor(p.Color, p.Life*0.3),
				Active: true,
			}
			w.Trails = append(w.Trails, trail)
		}
	}

	// Update trails
	for i := range w.Trails {
		t := &w.Trails[i]
		if !t.Active {
			continue
		}
		t.Life -= t.Decay
		if t.Life <= 0 {
			t.Active = false
		}
	}

	// Compact particles (remove dead ones)
	w.Particles = compact(w.Particles)
	w.Trails = compact(w.Trails)
}

// compact removes inactive particles in-place.
func compact(ps []Particle) []Particle {
	n := 0
	for i := range ps {
		if ps[i].Active {
			ps[n] = ps[i]
			n++
		}
	}
	return ps[:n]
}
