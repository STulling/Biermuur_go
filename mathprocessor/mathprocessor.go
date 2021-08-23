package mathprocessor

import (
	"github.com/STulling/Biermuur_go/audio/processing"
	"github.com/STulling/Biermuur_go/displaycontroller"
	"github.com/STulling/Biermuur_go/globals"
	"time"
)

var (
	// ToCalculate
	// Buffer of 64 samples, theoretically shouldn't get filled if
	// the pipeline is keeping up.
	// I just have this buffer if the timer is acting up
	ToCalculate = make(chan [][2]float64, 64)
	NewStuff = make(chan bool, 1)
	prevTime    = time.Now()
)

func RunCalculationPipe(sampleRate int) {
	for {
		<-NewStuff
		ticker := time.NewTicker(time.Second / time.Duration(sampleRate/globals.BLOCKSIZE))
		cycles := 0
		for cycles < 2*globals.AUDIOSYNC {
			<-ticker.C
			data := <-ToCalculate
			rms, tone := processing.ProcessBlock(data)
			displaycontroller.ToDisplay <- [2]float64{rms, tone}
			cycles++
		}
	}
}
