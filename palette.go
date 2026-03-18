package main

import "github.com/gdamore/tcell/v2"

// Color palettes for fireworks
var palettes = [][]tcell.Color{
	// Classic Red
	{
		tcell.NewRGBColor(255, 80, 80),
		tcell.NewRGBColor(255, 40, 40),
		tcell.NewRGBColor(200, 20, 20),
		tcell.NewRGBColor(150, 10, 10),
	},
	// Gold
	{
		tcell.NewRGBColor(255, 215, 0),
		tcell.NewRGBColor(255, 180, 0),
		tcell.NewRGBColor(200, 140, 0),
		tcell.NewRGBColor(150, 100, 0),
	},
	// Blue Ice
	{
		tcell.NewRGBColor(100, 180, 255),
		tcell.NewRGBColor(60, 140, 255),
		tcell.NewRGBColor(30, 100, 220),
		tcell.NewRGBColor(10, 60, 160),
	},
	// Green
	{
		tcell.NewRGBColor(80, 255, 80),
		tcell.NewRGBColor(40, 220, 40),
		tcell.NewRGBColor(20, 170, 20),
		tcell.NewRGBColor(10, 120, 10),
	},
	// Purple
	{
		tcell.NewRGBColor(200, 100, 255),
		tcell.NewRGBColor(160, 60, 220),
		tcell.NewRGBColor(120, 30, 180),
		tcell.NewRGBColor(80, 10, 130),
	},
	// Rainbow
	{
		tcell.NewRGBColor(255, 50, 50),
		tcell.NewRGBColor(255, 200, 50),
		tcell.NewRGBColor(50, 255, 50),
		tcell.NewRGBColor(50, 150, 255),
		tcell.NewRGBColor(200, 50, 255),
	},
	// Silver
	{
		tcell.NewRGBColor(220, 220, 230),
		tcell.NewRGBColor(180, 180, 190),
		tcell.NewRGBColor(130, 130, 140),
		tcell.NewRGBColor(80, 80, 90),
	},
}

// fadeColor fades a color toward black based on life (0.0-1.0).
// Uses life² for quadratic dimming (more natural).
func fadeColor(c tcell.Color, life float64) tcell.Color {
	if life <= 0 {
		return tcell.ColorBlack
	}
	factor := life * life // quadratic fade
	r, g, b := c.RGB()
	return tcell.NewRGBColor(
		int32(float64(r)*factor),
		int32(float64(g)*factor),
		int32(float64(b)*factor),
	)
}
