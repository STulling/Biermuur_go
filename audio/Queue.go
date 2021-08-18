package audio

import (
	"math/rand"

	"github.com/STulling/Biermuur_go/audio/processing"
	"github.com/STulling/Biermuur_go/musicio"
	"github.com/faiface/beep"
)

type Queue struct {
	streamers []beep.Streamer
	Requested bool
	PlayList  []string
}

func (q *Queue) AddSong(name string) {
	streamer := musicio.Load(name)
	q.Add(streamer)
}

func (q *Queue) addRandom() {
	streamer := musicio.Load(q.PlayList[rand.Intn(len(q.PlayList))])
	q.Add(streamer)
}

func (q *Queue) Clear() {
	q.streamers = q.streamers[len(q.streamers):]
}

func (q *Queue) Add(streamers ...beep.Streamer) {
	q.streamers = append(q.streamers, streamers...)
	if q.Requested {
		q.Requested = false
	}
}

func (q *Queue) SetPlaylist(list []string) {
	q.PlayList = list
}

func (q *Queue) Stream(samples [][2]float64) (n int, ok bool) {
	// We use the filled variable to track how many samples we've
	// successfully filled already. We loop until all samples are filled.
	filled := 0
	for filled < len(samples) {
		// There is just one song in the queue so we request the next.
		if len(q.streamers) == 1 && !q.Requested && len(q.PlayList) != 0 {
			q.Requested = true
			go q.addRandom()
		}
		// There are no streamers in the queue, so we stream silence.
		if len(q.streamers) == 0 {
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		// We stream from the first streamer in the queue.
		n, ok := q.streamers[0].Stream(samples[filled:])
		// If it's drained, we pop it from the queue, thus continuing with
		// the next streamer.
		if !ok {
			q.streamers = q.streamers[1:]
		}
		// We update the number of filled samples.
		filled += n
	}
	go processing.ProcessBlock(samples)
	return len(samples), true
}

func (q *Queue) Err() error {
	return nil
}
