package processing

import (
	"encoding/binary"
	"math"
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

var (
	better = make([]float64, 50)
	fblock = make([]float64, 256)
)

func ProcessBlock(block []byte) (float64, float64) {

	c1 := make(chan float64)

	for i := range fblock {
		fblock[i] = float64(binary.BigEndian.Uint16([]byte{block[1], block[2]})) / math.Pow(2, 16)
	}

	go calcRMS(fblock, c1)
	tone := calcFFT(fblock)
	rms := 0.

	select {
	case x := <-c1:
		rms = x
	}

	return rms, tone
}

func calcRMS(block []float64, c chan float64) {
	sum := 0.
	for _, x := range block {
		num := x
		sum += num * num
	}
	sum /= float64(len(block))
	sum = math.Sqrt(float64(sum))
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
