package processing

import (
	"math"
	"math/cmplx"

	"github.com/STulling/Biermuur_go/globals"
	"github.com/gonum/matrix/mat64"
	"github.com/mjibson/go-dsp/fft"
)

var (
	better    = make([]float64, 50)
	fblock    = make([]float64, globals.BLOCKSIZE)
	buffer    = make([]float64, 21)
	rmsBuffer = make([]float64, 10)
	indices   = []float64{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max_rms = 0.
)

func ProcessBlock(block [][2]float64) (float64, float64) {

	c1 := make(chan float64)

	for i := range fblock {
		fblock[i] = block[i][0]
	}

	go calcRMS(block, c1)
	tone := calcFFT(fblock)
	tone = denoise(tone)

	rmsBuffer = rmsBuffer[1:]
	rms := <-c1
	max_rms = math.Max(rms, max_rms)
	rms = rms / max_rms
	max_rms *= 0.99
	rmsBuffer = append(rmsBuffer, rms)

	return rmsBuffer[0], tone
}

func vandermonde(a []float64, degree int) *mat64.Dense {
	x := mat64.NewDense(len(a), degree+1, nil)
	for i := range a {
		for j, p := 0, 1.; j <= degree; j, p = j+1, p*a[i] {
			x.Set(i, j, p)
		}
	}
	return x
}

func denoise(tone float64) float64 {
	buffer = buffer[1:]
	buffer = append(buffer, tone)
	a := vandermonde(indices, 2)
	b := mat64.NewDense(len(buffer), 1, buffer)
	c := mat64.NewDense(3, 1, nil)

	qr := new(mat64.QR)
	qr.Factorize(a)

	err := c.SolveQR(qr, false, b)
	if err != nil {
		panic(err)
	}
	return c.At(0, 0)
}

func calcRMS(block [][2]float64, c chan float64) {
	sum := 0.
	for i := 0; i < len(block); i++ {
		sum += block[i][0] * block[i][0]
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
