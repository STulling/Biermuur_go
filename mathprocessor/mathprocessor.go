package mathprocessor

import (
	"fmt"
	"time"

	"github.com/STulling/Biermuur_go/audio/processing"
	"github.com/STulling/Biermuur_go/displaycontroller"
)

var (
	ToCalculate chan [][2]float64 = make(chan [][2]float64, 0)
	prevTime                      = time.Now()
)

func RunCalculationPipe() {
	for {
		data := <-ToCalculate
		fmt.Println(time.Since(prevTime))
		prevTime = time.Now()
		rms, tone := processing.ProcessBlock(data)
		displaycontroller.ToDisplay <- [2]float64{rms, tone}
	}
}
