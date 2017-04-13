package main

import (
	"fmt"
)

var (
	offsets             = []int{0, 0, 0, 0, 0}
	scaleLength float32 = 25.5
	fretCount           = 22
)

// type Tuning struct {
// 	ID      string `json:"id,omitempty"`
// 	Offsets []int  `json:"tuning,omitempty"`
// }

func getTuningOffsets(offset int) {
	for index, semitone := range offsets {
		offsets[index] = semitone + offset
	}
}

func getDropTuning() {
	if !(offsets[0]+2 == offsets[1]) {
		offsets[0] = offsets[0] - 2
	}
	fmt.Print("offsets: ", offsets)
}

func main() {
	// Down two semitones, full step
	getTuningOffsets(-2)
	// Drop tuning - low strings power chord, a 5th
	getDropTuning()
}
