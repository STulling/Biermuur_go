package audio

import (
	"fmt"

	"github.com/faiface/beep"
)

type AudioPlayer struct {
	callback    func(float32, float32)
	blocksize   int
	buffersize  int
	effectqueue chan float32
	audioqueue  chan float32
	volume      float32
	queue       Queue
}

const (
	blockSize = 1024
)

func CreateAudioPlayer() *AudioPlayer {
	p := new(AudioPlayer)
	p.buffersize = 1024
	p.effectqueue = make(chan float32, p.buffersize)
	p.audioqueue = make(chan float32, p.buffersize)
	p.volume = 1
	p.queue = Queue{}
	return p
}

func (audioPlayer *AudioPlayer) Start() {

	fmt.Println("Playing.")

	sr := beep.SampleRate(44100)

	Init(sr, 10*blockSize)
	Play()
	select {}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
