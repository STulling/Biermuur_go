package audio

import (
	"github.com/faiface/beep"
)

type SongStreamer struct {
	streamer beep.Streamer
	current  int
}

func (p SongStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	n, ok = p.streamer.Stream(samples)
	//fmt.Println(n)
	return n, ok
}

func (n SongStreamer) Err() error {
	return nil
}
