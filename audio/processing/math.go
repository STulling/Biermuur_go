package processing

import (
	"github.com/STulling/Biermuur_go/globals"
	"math"
	"math/cmplx"
	"github.com/mjibson/go-dsp/fft"
)

var (
	better = make([]float64, 50)
	fblock = make([]float64, globals.BLOCKSIZE)
	buffer = make([]float64, 10)
)

func ProcessBlock(block [][2]float64) (float64, float64) {

	c1 := make(chan float64)

	for i := range fblock {
		fblock[i] = block[i][0]
	}

	go calcRMS(fblock, c1)
	tone := calcFFT(fblock)
	tone = denoise(tone)

	rms := <-c1

	return rms, tone
}

func denoise(tone float64) float64 {
	buffer = buffer[1:]
	buffer = append(buffer, tone)
	sum := 0.
	for _, val := range buffer {
		sum += val
	}
	sum /= float64(len(buffer))
	return sum
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
