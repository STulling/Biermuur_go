package audio

import (
	"sync"

	"github.com/STulling/Biermuur_go/globals"
	"github.com/STulling/Biermuur_go/mathprocessor"

	"github.com/STulling/Biermuur_go/audio/oto"
	"github.com/faiface/beep"
	"github.com/pkg/errors"
)

const (
	blockSize = 1024
)

var (
	mu           sync.Mutex
	MusicQueue   Queue
	samples      [][2]float64
	sentSamples  [][2]float64
	buf          []byte
	context      *oto.Context
	player       *oto.Player
	done         chan struct{}
	syncBuffer   chan [][2]float64
	BufferAmount int
)

// Init initializes audio playback through speaker. Must be called before using this package.
//
// The bufferSize argument specifies the number of samples of the speaker's buffer. Bigger
// bufferSize means lower CPU usage and more reliable playback. Lower bufferSize means better
// responsiveness and less delay.
func Init(sampleRate beep.SampleRate, bufferAmount int) error {
	mu.Lock()
	defer mu.Unlock()

	BufferAmount = bufferAmount

	Close()

	MusicQueue = Queue{}

	numBytes := BufferAmount * globals.BUFFERSIZE
	samples = make([][2]float64, BufferAmount*globals.BLOCKSIZE)
	sentSamples = make([][2]float64, BufferAmount*globals.BLOCKSIZE)
	syncBuffer = make(chan [][2]float64, BufferAmount+globals.AUDIOSYNC)
	for i := 0; i < globals.AUDIOSYNC; i++ {
		syncBuffer <- make([][2]float64, globals.BLOCKSIZE)
	}
	buf = make([]byte, numBytes)

	var err error
	context, err = oto.NewContext(int(sampleRate), globals.CHANNELS, globals.BITDEPTH, numBytes)
	if err != nil {
		return errors.Wrap(err, "failed to initialize speaker")
	}
	player = context.NewPlayer()

	done = make(chan struct{})

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

// Lock locks the speaker. While locked, speaker won't pull new data from the playing Stramers. Lock
// if you want to modify any currently playing Streamers to avoid race conditions.
//
// Always lock speaker for as little time as possible, to avoid playback glitches.
func Lock() {
	mu.Lock()
}

// Unlock unlocks the speaker. Call after modifying any currently playing Streamer.
func Unlock() {
	mu.Unlock()
}

// Play starts playing all provided Streamers through the speaker.
func Play() {}

// Clear removes all currently playing Streamers from the speaker.
func Clear() {
	mu.Lock()
	MusicQueue.Clear()
	mu.Unlock()
}

// update pulls new data from the playing Streamers and sends it to the speaker. Blocks until the
// data is sent and started playing.
func update() {
	mu.Lock()
	MusicQueue.Stream(samples)
	mu.Unlock()
	mathprocessor.NewStuff <- true
	for written := 0; written < len(samples); written += globals.BLOCKSIZE {
		cpy := make([][2]float64, globals.BLOCKSIZE)
		copy(cpy, samples[written:written+globals.BLOCKSIZE])
		mathprocessor.ToCalculate <- cpy
		syncBuffer <- cpy
	}

	for i := 0; i < BufferAmount; i++ {
		copy(sentSamples[i*blockSize:(i+1)*blockSize], <-syncBuffer)
	}

	for i := range sentSamples {
		for c := range sentSamples[i] {
			val := sentSamples[i][c]
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

	player.Write(buf)
}
