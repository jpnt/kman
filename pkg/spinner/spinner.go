package spinner

import (
	"fmt"
	"time"
)

type Spinner struct {
	frames []string
	done   chan bool
}

func NewSpinner() *Spinner {
	return &Spinner{
		frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		done:   make(chan bool),
	}
}

func (s *Spinner) Start(message string) {
	go func() {
		fmt.Print(message)
		for {
			select {
			case <-s.done:
				return
			default:
				for _, frame := range s.frames {
					fmt.Printf("\r%s %s", frame, message)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.done <- true
	fmt.Print("\r\033[K") // Clear the spinner line
	fmt.Println("Done.")
}
