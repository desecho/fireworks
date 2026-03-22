package main

import "math/rand"

const (
	gravity      = 0.035
	drag         = 0.985
	maxParticles = 3000
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
		Rockets:       make([]*Rocket, 0, 32),
		Particles:     make([]Particle, 0, maxParticles),
		Trails:        make([]Particle, 0, maxParticles),
		RNG:           rng,
		spawnInterval: 15 + rng.Intn(20),
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

	// Maybe spawn a new rocket — occasionally spawn 2-3 at once for volleys
	if w.tickCount >= w.spawnInterval && len(w.Particles) < maxParticles {
		count := 1
		if w.RNG.Float64() < 0.25 {
			count = 2 + w.RNG.Intn(2) // volley of 2-3
		}
		for i := 0; i < count; i++ {
			w.Rockets = append(w.Rockets, NewRocket(w.Width, w.Height, w.RNG))
		}
		w.tickCount = 0
		w.spawnInterval = 15 + w.RNG.Intn(25)
	}

	// Update rockets
	aliveRockets := w.Rockets[:0]
	for _, r := range w.Rockets {
		r.Particle.Update(gravity*0.4, 1.0)
		r.FuseLife--
		if r.FuseLife <= 0 || r.Y < 2 {
			newParticles := r.Explode(w.RNG)
			w.Particles = append(w.Particles, newParticles...)
		} else {
			// Rocket trail — sparks flying off
			if len(w.Trails) < maxParticles {
				w.Trails = append(w.Trails, Particle{
					X: r.X + (w.RNG.Float64()-0.5)*0.5,
					Y: r.Y,
					VX: (w.RNG.Float64() - 0.5) * 0.3,
					VY: 0.2 + w.RNG.Float64()*0.3,
					Life: 0.5 + w.RNG.Float64()*0.3, Decay: 0.06 + w.RNG.Float64()*0.04,
					Char:   '.',
					Color:  fadeColor(r.Color, 0.6),
					Active: true,
				})
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
		p.Char = charForLife(p.Life, p.Shimmer)

		if p.X < 0 || p.X >= float64(w.Width) || p.Y < 0 || p.Y >= float64(w.Height) {
			p.Active = false
			continue
		}

		// Spawn trail behind active particles
		if len(w.Trails) < maxParticles && p.Life > 0.15 && w.RNG.Float64() < 0.6 {
			w.Trails = append(w.Trails, Particle{
				X: p.X, Y: p.Y,
				Life: p.Life * 0.4, Decay: 0.04 + w.RNG.Float64()*0.03,
				Char:   '.',
				Color:  fadeColor(p.Color, p.Life*0.35),
				Active: true,
			})
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
