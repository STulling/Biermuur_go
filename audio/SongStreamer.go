package audio

type Playback struct {
	song    *Song
	current int
}

func (p Playback) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0] = float64(p.song.audio[2*p.current])
		samples[i][1] = float64(p.song.audio[2*p.current+1])
		p.current += 1
	}
	return len(samples), true
}

func (n Playback) Err() error {
	return nil
}
