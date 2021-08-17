package audio

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

const SIZEOF_INT16 = 2

type Song struct {
	audio       []float32
	rms_buffer  []float32
	tone_buffer []float32
}

func readAll(d *mp3.Decoder) []float32 {
	bytes, _ := ioutil.ReadAll(d)
	data := make([]float32, len(bytes)/SIZEOF_INT16)
	for i := range data {
		data[i] = math.Float32frombits(uint32(binary.LittleEndian.Uint16(bytes[i*SIZEOF_INT16 : (i+1)*SIZEOF_INT16])))
	}
	return data
}

func createBuffers(song *Song, blocksize int) {
	size := len(song.audio) / blocksize

	rms_buffer := make([]float32, size)
	tone_buffer := make([]float32, size)

	for i := 0; i < size; i++ {
		sum := float32(0.)
		for x := 0; x <= blocksize; x++ {
			num := float32(song.audio[i*blocksize+x])
			sum += num * num
		}
		sum /= float32(blocksize)
		sum = float32(math.Sqrt(float64(sum)))
		rms_buffer[i] = sum
		//fft.FFTReal()
	}

	song.rms_buffer = rms_buffer
	song.tone_buffer = tone_buffer
}

func Load(name string, blocksize int) *Song {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatal(err)
	}

	song := new(Song)
	song.audio = readAll(d)
	createBuffers(song, blocksize)
	return song
}
