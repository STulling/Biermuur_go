package audio

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type AudioPlayer struct {
	callback    func(float32, float32)
	blocksize   int
	buffersize  int
	effectqueue chan float32
	audioqueue  chan float32
	queue       chan Song
	volume      float32
}

func CreateAudioPlayer() *AudioPlayer {
	p := new(AudioPlayer)
	p.blocksize = 1024
	p.buffersize = 1024
	p.effectqueue = make(chan float32, p.buffersize)
	p.audioqueue = make(chan float32, p.buffersize)
	p.volume = 1
	return p
}

func (audioPlayer *AudioPlayer) Play(file string) {

	song := Load(file, audioPlayer.blocksize)

	fmt.Println("Playing.  Press Ctrl-C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	format := beep.Format{
		SampleRate:  14400,
		NumChannels: 2,
		Precision:   2,
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	playback := Playback{
		song:    song,
		current: 0,
	}
	for i := 0; i < len(song.audio)/audioPlayer.blocksize; i++ {
		speaker.Play(playback)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
