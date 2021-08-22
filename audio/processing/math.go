package processing

import (
	"encoding/binary"
	"github.com/STulling/Biermuur_go/globals"
	"math"
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

var (
	better = make([]float64, 50)
	fblock = make([]float64, globals.BLOCKSIZE)
)

func ProcessBlock(block []byte) (float64, float64) {

	c1 := make(chan float64)

	for i := range fblock {
		fblock[i] = float64(binary.LittleEndian.Uint16([]byte{block[i*4], block[i*4 + 1]})) / math.Pow(2, 16)
	}

	go calcRMS(fblock, c1)
	tone := calcFFT(fblock)

	rms := <-c1

	return rms, tone
}

func calcRMS(block []float64, c chan float64) {
	sum := 0.
	for i := 0; i < len(block); i++ {
		sum += block[i] * block[i]
	}
	sum /= float64(len(block))
	sum = math.Sqrt(sum)
	c <- sum
}

func argmax(list []float64) int {
	index := 0
	max := list[0]
	for i, val := range list {
		if val > max {
			max = val
			index = i
		}
	}
	return index
}

func calcFFT(block []float64) float64 {
	fft := fft.FFTReal(block)
	for i, val := range fft[11 : 11+50] {
		better[i] = cmplx.Abs(val)
	}
	color := float64(argmax(better))
	return math.Max(0., math.Min(color*10., 255.)) / 255.
}
