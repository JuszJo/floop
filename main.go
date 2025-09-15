package main

import "fmt"

type Signal int

const (
	CONTINUE Signal = iota
	DONE
)

type RunnerExiter interface {
	Tick(fn func() Signal) Signal
	Exit(fn func())
}

func FeedbackLoop(re RunnerExiter, work func() Signal, onExit func()) {
	state := CONTINUE

	for state == CONTINUE {
		signal := re.Tick(work)
		if signal == DONE {
			state = DONE
		}
	}

	re.Exit(onExit)
}

type SimpleRunnerExiter struct{}

func (s *SimpleRunnerExiter) Tick(fn func() Signal) Signal {
	return fn()
}

func (s *SimpleRunnerExiter) Exit(fn func()) {
	fn()
}

type Goodbye struct{}

func (g *Goodbye) Exit() {
	fmt.Println("Exiting gracefully")
}

// --- Main ---
func main() {
	re := &SimpleRunnerExiter{}

	count := 3

	FeedbackLoop(
		re,
		func() Signal {
			if count <= 0 {
				return DONE
			}
			fmt.Println("Running:", count)
			count--
			return CONTINUE
		},
		func() {
			fmt.Println("Exiting gracefully")
		},
	)
}