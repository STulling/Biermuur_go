package processing

import (
	"math"
	"math/cmplx"

	"github.com/STulling/Biermuur_go/globals"
	"github.com/gonum/matrix/mat64"
	"github.com/mjibson/go-dsp/fft"
)

var (
	better     = make([]float64, 50)
	fblock     = make([]float64, globals.BLOCKSIZE)
	buffer     = make([]float64, 21)
	rmsBuffer  = make(chan float64, globals.AUDIOSYNC+globals.BUFFERAMOUNT)
	toneBuffer = make(chan float64, globals.BUFFERAMOUNT)
	indices    = make([][]float64, 21)
	maxRms     = 1.
	i          = 0
)

func InitBuffers() {
	for i := 0; i < globals.AUDIOSYNC+globals.BUFFERAMOUNT; i++ {
		rmsBuffer <- 0
	}
	for i := 0; i < globals.BUFFERAMOUNT; i++ {
		toneBuffer <- 0
	}
	for i := 0; i < 21; i++ {
		indices[i] = make([]float64, 21)
		for j := 0; j < 21; j++ {
			val := float64(-10 + i + j)
			if val > 10 {
				val -= 20
			}
			indices[i][j] = val
		}
	}
}

func ProcessBlock(block [][2]float64) (float64, float64) {

	c1 := make(chan float64)

	for i := range fblock {
		fblock[i] = block[i][0]
	}

	go calcRMS(block, c1)
	tone := calcFFT(fblock)
	tone = denoise(tone)

	currRms := <-rmsBuffer
	rms := <-c1
	maxRms = math.Max(rms, maxRms)
	rmsBuffer <- rms
	ret := math.Min(1., currRms/maxRms)
	maxRms *= 0.99

	currTone := <-toneBuffer
	toneBuffer <- tone

	return ret, currTone
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
	buffer[i] = tone
	a := vandermonde(indices[i], 2)
	b := mat64.NewDense(len(buffer), 1, buffer)
	c := mat64.NewDense(3, 1, nil)

	qr := new(mat64.QR)
	qr.Factorize(a)

	err := c.SolveQR(qr, false, b)
	if err != nil {
		panic(err)
	}

	i++
	if i > 20 {
		i = 0
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
