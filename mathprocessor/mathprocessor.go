package mathprocessor

import (
	"time"

	"github.com/STulling/Biermuur_go/audio"
	"github.com/STulling/Biermuur_go/audio/processing"
	"github.com/STulling/Biermuur_go/displaycontroller"
	"github.com/STulling/Biermuur_go/globals"
)

var (
	// ToCalculate
	// Buffer of 64 samples, theoretically shouldn't get filled if
	// the pipeline is keeping up.
	// I just have this buffer if the timer is acting up
	ToCalculate = make(chan [][2]float64, 64)
	NewStuff    = make(chan bool, 10)
	prevTime    = time.Now()
)

func RunCalculationPipe(sampleRate int) {
	for {
		<-NewStuff
		prevTime = time.Now()
		data := <-ToCalculate
		rms, tone := processing.ProcessBlock(data)
		displaycontroller.ToDisplay <- [2]float64{rms, tone}
		cycles := 1
		for cycles < 2*globals.AUDIOSYNC {
			time.Sleep(time.Second/time.Duration(sampleRate/globals.BLOCKSIZE) - time.Now().Sub(prevTime))
			prevTime = time.Now()
			data := <-ToCalculate
			if audio.MusicQueue.Size() != 0 {
				rms, tone = processing.ProcessBlock(data)
			} else {
				rms, tone = 0, 0
			}
			displaycontroller.ToDisplay <- [2]float64{rms, tone}
			cycles++
		}
	}
}
