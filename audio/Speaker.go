package audio

import (
	"fmt"
	"sync"

	"github.com/STulling/Biermuur_go/mathprocessor"
	"github.com/faiface/beep"
	"github.com/gordonklaus/portaudio"
	"github.com/hajimehoshi/oto"
)

const (
	blockSize = 1024
)

var (
	mu         sync.Mutex
	MusicQueue Queue
	samples    [][2]float64
	buf        []byte
	context    *oto.Context
	player     *oto.Player
	done       chan struct{}
	out        []byte
	stream     portaudio.Stream
)

// Init initializes audio playback through speaker. Must be called before using this package.
//
// The bufferSize argument specifies the number of samples of the speaker's buffer. Bigger
// bufferSize means lower CPU usage and more reliable playback. Lower bufferSize means better
// responsiveness and less delay.
func Init(sampleRate beep.SampleRate, bufferSize int) error {

	MusicQueue = Queue{}
	samples = make([][2]float64, bufferSize)
	out = make([]byte, len(samples)*4)
	buf = make([]byte, len(samples)*4)

	portaudio.Initialize()

	stream, err := portaudio.OpenDefaultStream(0, 2, 44100, len(out), &out)
	if err != nil {
		panic(err)
	}
	err = stream.Start()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			default:
				update()
			case <-done:
				return
			}
		}
	}()

	return nil
}

// Close closes the playback and the driver. In most cases, there is certainly no need to call Close
// even when the program doesn't play anymore, because in properly set systems, the default mixer
// handles multiple concurrent processes. It's only when the default device is not a virtual but hardware
// device, that you'll probably want to manually manage the device from your application.
func Close() {
	if player != nil {
		if done != nil {
			done <- struct{}{}
			done = nil
		}
		player.Close()
		context.Close()
		player = nil
	}
}

// Play starts playing all provided Streamers through the speaker.
func Play() {}

// Clear removes all currently playing Streamers from the speaker.
func Clear() {
	mu.Lock()
	MusicQueue.Clear()
	mu.Unlock()
}

func write() {
	for remaining := int(blockSize); remaining > 0; remaining -= len(out) {
		if len(out) > remaining {
			out = out[:remaining]
		}
		out = append(buf[:remaining], out...)
		//err := binary.Read(audio, binary.BigEndian, out)
		fmt.Println(stream.Info())
		err := stream.Write()
		if err != nil {
			panic(err)
		}
	}
}

// update pulls new data from the playing Streamers and sends it to the speaker. Blocks until the
// data is sent and started playing.
func update() {
	mu.Lock()
	MusicQueue.Stream(samples)
	mu.Unlock()

	for i := range samples {
		for c := range samples[i] {
			val := samples[i][c]
			if val < -1 {
				val = -1
			}
			if val > +1 {
				val = +1
			}
			valInt16 := int16(val * (1<<15 - 1))
			low := byte(valInt16)
			high := byte(valInt16 >> 8)
			buf[i*4+c*2+0] = low
			buf[i*4+c*2+1] = high
		}
	}
	mathprocessor.ToCalculate <- samples
	write()
}
