package processing

import (
	"math"
	"math/cmplx"

	"github.com/STulling/Biermuur_go/displaycontroller"
	"github.com/mjibson/go-dsp/fft"
)

func ProcessBlock(block [][2]float64) {

	c1 := make(chan float64)
	c2 := make(chan float64)

	var result [2]float64

	go calcRMS(block, c1)
	go calcFFT(block, c2)

	for i := 0; i < 2; i++ {
		select {
		case rms := <-c1:
			result[0] = rms
		case fft := <-c2:
			result[1] = fft
		}
	}

	displaycontroller.ToDisplay <- result
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

func calcFFT(block [][2]float64, c chan float64) {
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
	c <- math.Max(0., math.Min(color*10., 255.)) / 255.
}
