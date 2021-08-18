package mathprocessor

import (
	"github.com/STulling/Biermuur_go/audio/processing"
	"github.com/STulling/Biermuur_go/display"
	"github.com/STulling/Biermuur_go/displaycontroller"
)

var (
	ToCalculate chan [][2]float64 = make(chan [][2]float64, 0)
)

func RunCalculationPipe() {
	display.Init()
	for {
		data := <-ToCalculate
		rms, tone := processing.ProcessBlock(data)
		displaycontroller.ToDisplay <- [2]float64{rms, tone}
	}
}
