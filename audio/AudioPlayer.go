package audio

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/faiface/beep/mp3"
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

func CreateAudioPlayer() *AudioPlayer {
	p := new(AudioPlayer)
	p.blocksize = 1024
	p.buffersize = 1024
	p.effectqueue = make(chan float32, p.buffersize)
	p.audioqueue = make(chan float32, p.buffersize)
	p.volume = 1
	p.queue = Queue{}
	return p
}

func (audioPlayer *AudioPlayer) Play(file string) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, _ := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	song := SongStreamer{
		streamer: streamer,
		current:  0,
	}

	audioPlayer.queue.Add(song)

	fmt.Println("Playing.  Press Ctrl-C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	Play(&audioPlayer.queue)
	select {}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
