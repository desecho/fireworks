package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	s.Clear()

	w, h := s.Size()
	world := NewWorld(w, h)

	// Quit channel
	quit := make(chan struct{})

	// Event goroutine
	go func() {
		for {
			ev := s.PollEvent()
			switch ev.(type) {
			case *tcell.EventKey:
				close(quit)
				return
			case *tcell.EventResize:
				nw, nh := s.Size()
				world.Resize(nw, nh)
				s.Sync()
			}
		}
	}()

	// Main loop ~30 FPS
	ticker := time.NewTicker(33 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			s.Fini()
			return
		case <-ticker.C:
			world.Update()
			Render(s, world)
		}
	}
}
