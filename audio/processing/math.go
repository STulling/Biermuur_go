package processing

import (
	"math"
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

func ProcessBlock(block [][2]float64) (float64, float64) {

	c1 := make(chan float64)

	go calcRMS(block, c1)
	//tone := calcFFT(block)
	rms := 0.

	select {
	case x := <-c1:
		rms = x
	}

	return rms, 0
}

func calcRMS(block [][2]float64, c chan float64) {
	sum := 0.
	for _, x := range block {
		num := x[0]
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

func calcFFT(block [][2]float64) float64 {
	channel := make([]float64, len(block))
	for i, x := range block {
		channel[i] = x[0]
	}
	fft := fft.FFTReal(channel[:])
	better := make([]float64, 50)
	for i, val := range fft[11 : 11+50] {
		better[i] = cmplx.Abs(val)
	}
	color := float64(argmax(better))
	return math.Max(0., math.Min(color*10., 255.)) / 255.
}
