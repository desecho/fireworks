package main

import "github.com/gdamore/tcell/v2"

// Color palettes for fireworks — vibrant, multi-hue combinations
var palettes = [][]tcell.Color{
	// Molten Copper → Pink
	{
		tcell.NewRGBColor(255, 120, 40),
		tcell.NewRGBColor(255, 80, 60),
		tcell.NewRGBColor(255, 50, 90),
		tcell.NewRGBColor(255, 170, 80),
		tcell.NewRGBColor(255, 200, 140),
	},
	// Electric Cyan → Violet
	{
		tcell.NewRGBColor(0, 255, 255),
		tcell.NewRGBColor(80, 200, 255),
		tcell.NewRGBColor(140, 140, 255),
		tcell.NewRGBColor(200, 80, 255),
		tcell.NewRGBColor(180, 220, 255),
	},
	// Emerald Flame
	{
		tcell.NewRGBColor(50, 255, 100),
		tcell.NewRGBColor(100, 255, 60),
		tcell.NewRGBColor(180, 255, 40),
		tcell.NewRGBColor(255, 255, 80),
		tcell.NewRGBColor(200, 255, 200),
	},
	// Cherry Blossom
	{
		tcell.NewRGBColor(255, 130, 170),
		tcell.NewRGBColor(255, 180, 200),
		tcell.NewRGBColor(255, 220, 230),
		tcell.NewRGBColor(255, 100, 140),
		tcell.NewRGBColor(255, 200, 220),
	},
	// Royal Gold
	{
		tcell.NewRGBColor(255, 215, 0),
		tcell.NewRGBColor(255, 240, 100),
		tcell.NewRGBColor(255, 180, 0),
		tcell.NewRGBColor(255, 255, 180),
		tcell.NewRGBColor(255, 200, 60),
	},
	// Deep Sea
	{
		tcell.NewRGBColor(0, 100, 255),
		tcell.NewRGBColor(0, 180, 255),
		tcell.NewRGBColor(80, 220, 255),
		tcell.NewRGBColor(0, 60, 200),
		tcell.NewRGBColor(160, 240, 255),
	},
	// Aurora
	{
		tcell.NewRGBColor(0, 255, 120),
		tcell.NewRGBColor(0, 200, 255),
		tcell.NewRGBColor(120, 0, 255),
		tcell.NewRGBColor(255, 0, 200),
		tcell.NewRGBColor(80, 255, 200),
	},
	// Crimson Velvet
	{
		tcell.NewRGBColor(255, 20, 40),
		tcell.NewRGBColor(255, 60, 20),
		tcell.NewRGBColor(200, 0, 40),
		tcell.NewRGBColor(255, 100, 60),
		tcell.NewRGBColor(255, 160, 120),
	},
	// Neon Spectrum
	{
		tcell.NewRGBColor(255, 0, 80),
		tcell.NewRGBColor(255, 255, 0),
		tcell.NewRGBColor(0, 255, 200),
		tcell.NewRGBColor(255, 100, 0),
		tcell.NewRGBColor(0, 200, 255),
		tcell.NewRGBColor(200, 0, 255),
	},
	// White Titanium
	{
		tcell.NewRGBColor(255, 255, 255),
		tcell.NewRGBColor(220, 230, 255),
		tcell.NewRGBColor(200, 210, 240),
		tcell.NewRGBColor(240, 240, 255),
		tcell.NewRGBColor(180, 200, 255),
	},
}

// fadeColor fades a color toward black based on life (0.0-1.0).
func fadeColor(c tcell.Color, life float64) tcell.Color {
	if life <= 0 {
		return tcell.ColorBlack
	}
	factor := life * life // quadratic fade — stays bright longer, then drops fast
	r, g, b := c.RGB()
	return tcell.NewRGBColor(
		int32(float64(r)*factor),
		int32(float64(g)*factor),
		int32(float64(b)*factor),
	)
}

// lerpColor blends two colors by t (0.0 = a, 1.0 = b).
func lerpColor(a, b tcell.Color, t float64) tcell.Color {
	if t <= 0 {
		return a
	}
	if t >= 1 {
		return b
	}
	r1, g1, b1 := a.RGB()
	r2, g2, b2 := b.RGB()
	return tcell.NewRGBColor(
		int32(float64(r1)*(1-t)+float64(r2)*t),
		int32(float64(g1)*(1-t)+float64(g2)*t),
		int32(float64(b1)*(1-t)+float64(b2)*t),
	)
}
